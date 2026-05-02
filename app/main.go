// package main tells Go this file builds an executable program (not a reusable library).
package main

import (
	// fmt is used to write simple text responses to HTTP clients.
	"fmt"
	// log is used for startup logging and fatal server errors.
	"log"
	// net/http provides the HTTP server, routing, and response utilities.
	"net/http"
	// time is used by /load to run a CPU busy loop for a fixed duration.
	"time"
)

// KubePulse AWS is a minimal operations-focused API.
// The endpoints below are intentionally simple so they can be used to
// practice DevOps workflows like health checks, failure testing, and load testing.

// isFailing is an in-memory flag that controls the health endpoint behavior.
// false: service is healthy.
// true: service is unhealthy.
//
// In Kubernetes terms, this value affects liveness-style behavior:
// - /health can tell the platform the app is no longer healthy.
// - /ready remains a readiness-style signal (currently always ready on Day 1/2).
//
// Note: this is intentionally simple for learning. In production, state is usually
// managed more safely and often with synchronization primitives for concurrency.
var isFailing = false

// main wires all routes and starts the HTTP server on port 8080.
//
// This service is intentionally small so you can clearly see cloud-native ideas:
// - operational endpoints (/health and /ready)
// - controlled failure simulation (/fail)
// - controlled load simulation (/load)
//
// These patterns are common in Kubernetes workloads where orchestration depends
// on endpoint signals to make restart and traffic-routing decisions.
func main() {
	// Home endpoint: quick confirmation that the process is running.
	http.HandleFunc("/", homeHandler)

	// Health endpoint: liveness-style indicator.
	// When this fails consistently in Kubernetes, pods may be restarted.
	http.HandleFunc("/health", healthHandler)

	// Readiness endpoint: readiness-style indicator.
	// When this fails in Kubernetes, traffic can be temporarily stopped to this pod.
	http.HandleFunc("/ready", readyHandler)

	// Failure simulation endpoint: toggles app into failing mode.
	http.HandleFunc("/fail", failHandler)

	// Load simulation endpoint: keeps CPU busy for a short period.
	http.HandleFunc("/load", loadHandler)

	log.Println("KubePulse API running on port 8080")

	// Start the HTTP server and stop the process if startup/runtime fails.
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// homeHandler handles GET /.
// It returns a human-readable message confirming the API is online.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "KubePulse AWS is running")
}

// healthHandler handles GET /health.
//
// Why it matters in Kubernetes:
// - This endpoint is commonly used by liveness probes.
// - If it reports failures, Kubernetes can restart the container to recover.
//
// Behavior in this project:
// - If isFailing == true, return HTTP 500 (unhealthy).
// - Otherwise return HTTP 200 (healthy).
func healthHandler(w http.ResponseWriter, r *http.Request) {
	if isFailing {
		http.Error(w, "service unhealthy", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "healthy")
}

// readyHandler handles GET /ready.
//
// Why it matters in Kubernetes:
// - This endpoint is commonly used by readiness probes.
// - Readiness controls whether a pod should receive traffic.
//
// Current behavior:
// - Always returns "ready" for this early project phase.
// - Later versions can include dependency checks (DB, cache, API reachability).
func readyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ready")
}

// failHandler handles GET /fail.
//
// This endpoint simulates an application failure by setting isFailing = true.
// After this call, /health starts returning HTTP 500, which mimics a broken app
// state and allows you to test monitoring, alerts, and orchestration behavior.
//
// Returning HTTP 500 here makes the failure activation explicit to callers.
func failHandler(w http.ResponseWriter, r *http.Request) {
	isFailing = true
	http.Error(w, "failure mode activated", http.StatusInternalServerError)
}

// loadHandler handles GET /load.
//
// This endpoint simulates temporary CPU pressure by running a busy loop for
// around 5 seconds. The repeated arithmetic operation is intentionally useless
// work whose purpose is to consume CPU cycles.
//
// This helps test:
// - service responsiveness under load
// - future autoscaling behavior
// - how probes behave during higher resource usage
func loadHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	for time.Since(start) < 5*time.Second {
		_ = 999999 * 999999
	}

	fmt.Fprintln(w, "CPU load simulated for 5 seconds")
}
