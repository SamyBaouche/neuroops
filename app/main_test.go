package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func get(t *testing.T, h http.Handler, path string) *httptest.ResponseRecorder {
	t.Helper()
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, path, nil))
	return rec
}

func TestHealthyEndpoints(t *testing.T) {
	isFailing.Store(false)
	mux := newMux()

	for _, path := range []string{"/", "/health", "/ready"} {
		if rec := get(t, mux, path); rec.Code != http.StatusOK {
			t.Errorf("GET %s = %d, want 200", path, rec.Code)
		}
	}
}

func TestUnknownPathIsNotFound(t *testing.T) {
	isFailing.Store(false)
	if rec := get(t, newMux(), "/does-not-exist"); rec.Code != http.StatusNotFound {
		t.Errorf("GET /does-not-exist = %d, want 404", rec.Code)
	}
}

func TestFailureModeFailsProbesAndTraffic(t *testing.T) {
	isFailing.Store(false)
	t.Cleanup(func() { isFailing.Store(false) })
	mux := newMux()

	if rec := get(t, mux, "/fail"); rec.Code != http.StatusInternalServerError {
		t.Fatalf("GET /fail = %d, want 500", rec.Code)
	}

	want := map[string]int{
		"/":       http.StatusServiceUnavailable,
		"/health": http.StatusInternalServerError,
		"/ready":  http.StatusServiceUnavailable,
	}
	for path, code := range want {
		if rec := get(t, mux, path); rec.Code != code {
			t.Errorf("GET %s in failure mode = %d, want %d", path, rec.Code, code)
		}
	}
}

func TestMetricsExposeRequestCounters(t *testing.T) {
	isFailing.Store(false)
	mux := newMux()
	get(t, mux, "/health")

	body := get(t, mux, "/metrics").Body.String()
	for _, want := range []string{
		`neuroops_http_requests_total{code="200",method="get",route="/health"}`,
		`neuroops_http_request_duration_seconds_bucket{method="get",route="/health"`,
		`neuroops_failure_mode 0`,
	} {
		if !strings.Contains(body, want) {
			t.Errorf("/metrics missing %q", want)
		}
	}
}
