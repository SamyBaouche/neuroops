// package main tells Go this file builds an executable program (not a reusable library).
package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

// NeuroOps is a small operations-focused API used to exercise Kubernetes
// reliability patterns: health/readiness probes, failure injection, CPU load,
// and the three observability signals (metrics, logs, traces).

// isFailing is the simulated failure state. Once /fail is called, the pod
// serves errors on / and fails both probes until Kubernetes restarts it.
// atomic.Bool because handlers run concurrently.
var isFailing atomic.Bool

// registry holds every metric exposed on /metrics.
var registry = prometheus.NewRegistry()

var (
	// httpRequests counts requests by route, method, and status code.
	// Error-rate and availability SLOs are computed from this counter.
	httpRequests = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "neuroops_http_requests_total",
		Help: "HTTP requests processed, by route, method, and status code.",
	}, []string{"route", "method", "code"})

	// httpDuration records request latency; p50/p99 come from its buckets.
	httpDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "neuroops_http_request_duration_seconds",
		Help:    "HTTP request latency, by route and method.",
		Buckets: []float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.2, 0.5, 1, 2.5, 5, 10},
	}, []string{"route", "method"})

	// failureMode exposes the simulated failure flag (1 = failing).
	failureMode = prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Name: "neuroops_failure_mode",
		Help: "1 when failure mode has been activated through /fail.",
	}, func() float64 {
		if isFailing.Load() {
			return 1
		}
		return 0
	})
)

func init() {
	registry.MustRegister(
		httpRequests,
		httpDuration,
		failureMode,
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
	)
}

// main sets up logging and tracing, starts the HTTP server on :8080, and
// shuts down gracefully on SIGTERM so in-flight requests finish during pod
// restarts and rollouts.
func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	shutdownTracing, err := setupTracing(ctx)
	if err != nil {
		slog.Error("tracing setup failed", "error", err)
		os.Exit(1)
	}

	server := &http.Server{
		Addr:              ":8080",
		Handler:           newMux(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		slog.Info("NeuroOps API running", "port", 8080)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("server failed", "error", err)
			os.Exit(1)
		}
	}()

	<-ctx.Done()
	slog.Info("shutting down")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		slog.Error("graceful shutdown failed", "error", err)
	}
	if err := shutdownTracing(shutdownCtx); err != nil {
		slog.Error("tracing shutdown failed", "error", err)
	}
}

// setupTracing exports spans over OTLP/gRPC when OTEL_EXPORTER_OTLP_ENDPOINT
// is set (e.g. Tempo). Without it, tracing is a no-op so the API still runs
// locally with no collector. Service name comes from OTEL_SERVICE_NAME.
func setupTracing(ctx context.Context) (func(context.Context) error, error) {
	if os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT") == "" {
		return func(context.Context) error { return nil }, nil
	}

	exporter, err := otlptracegrpc.New(ctx)
	if err != nil {
		return nil, err
	}

	provider := sdktrace.NewTracerProvider(sdktrace.WithBatcher(exporter))
	otel.SetTracerProvider(provider)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))
	return provider.Shutdown, nil
}

// newMux wires every route through the same instrumentation chain:
// tracing span -> structured log -> request counter -> latency histogram.
func newMux() *http.ServeMux {
	mux := http.NewServeMux()

	handle(mux, "/{$}", "/", homeHandler)
	handle(mux, "/health", "/health", healthHandler)
	handle(mux, "/ready", "/ready", readyHandler)
	handle(mux, "/fail", "/fail", failHandler)
	handle(mux, "/load", "/load", loadHandler)

	mux.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	return mux
}

// handle registers pattern on mux with metrics, logs, and tracing. route is
// the label value used in metrics and logs, kept low-cardinality on purpose.
func handle(mux *http.ServeMux, pattern, route string, h http.HandlerFunc) {
	labels := prometheus.Labels{"route": route}

	var handler http.Handler = h
	handler = promhttp.InstrumentHandlerDuration(httpDuration.MustCurryWith(labels), handler)
	handler = promhttp.InstrumentHandlerCounter(httpRequests.MustCurryWith(labels), handler)
	handler = withLogging(route, handler)
	handler = otelhttp.NewHandler(handler, route)

	mux.Handle(pattern, handler)
}

// statusRecorder captures the status code written by a handler for logging.
type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

// withLogging emits one JSON log line per request, including the trace ID so
// Loki log lines can be correlated with Tempo traces. Successful probe calls
// are skipped because kubelet sends one every 5 seconds per probe.
func withLogging(route string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(rec, r)

		if (route == "/health" || route == "/ready") && rec.status < 400 {
			return
		}

		level := slog.LevelInfo
		if rec.status >= 500 {
			level = slog.LevelError
		}

		attrs := []any{
			"route", route,
			"method", r.Method,
			"status", rec.status,
			"duration_ms", time.Since(start).Milliseconds(),
		}
		if sc := trace.SpanContextFromContext(r.Context()); sc.HasTraceID() {
			attrs = append(attrs, "trace_id", sc.TraceID().String())
		}
		slog.Log(r.Context(), level, "request", attrs...)
	})
}

// homeHandler handles GET /. It serves errors while failure mode is active,
// which is what users would see from a broken pod.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if isFailing.Load() {
		http.Error(w, "service unhealthy", http.StatusServiceUnavailable)
		return
	}
	fmt.Fprintln(w, "NeuroOps is running")
}

// healthHandler handles GET /health (liveness probe). Two consecutive
// failures make the kubelet restart the container.
func healthHandler(w http.ResponseWriter, r *http.Request) {
	if isFailing.Load() {
		http.Error(w, "service unhealthy", http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "healthy")
}

// readyHandler handles GET /ready (readiness probe). A failing pod reports
// not-ready so the Service stops routing traffic to it before it restarts.
func readyHandler(w http.ResponseWriter, r *http.Request) {
	if isFailing.Load() {
		http.Error(w, "not ready", http.StatusServiceUnavailable)
		return
	}
	fmt.Fprintln(w, "ready")
}

// failHandler handles GET /fail. It switches the pod into failure mode until
// it is restarted, to exercise probes, alerts, and recovery.
func failHandler(w http.ResponseWriter, r *http.Request) {
	isFailing.Store(true)
	slog.Warn("failure mode activated")
	http.Error(w, "failure mode activated", http.StatusInternalServerError)
}

// loadHandler handles GET /load. It burns CPU for about 5 seconds to trigger
// HPA scale-up.
func loadHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	for time.Since(start) < 5*time.Second {
		_ = 999999 * 999999
	}

	fmt.Fprintln(w, "CPU load simulated for 5 seconds")
}
