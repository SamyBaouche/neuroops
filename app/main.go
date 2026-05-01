package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// isFailing is a simple in-memory flag that controls the behavior of /health.
// false => service reports healthy, true => service reports unhealthy.
// The /fail endpoint flips this flag to true to simulate an application failure.
var isFailing = false

// main registers all HTTP routes and starts the API server on port 8080.
// For Day 1, we intentionally keep the server minimal and stateful so we can
// test basic cloud-native behaviors such as health checks and failure handling.
func main() {
	// Home endpoint: quick status message to confirm the API is running.
	http.HandleFunc("/", homeHandler)
	// Health endpoint: returns healthy/unhealthy depending on isFailing state.
	http.HandleFunc("/health", healthHandler)
	// Readiness endpoint: indicates the app is ready to receive traffic.
	http.HandleFunc("/ready", readyHandler)
	// Failure trigger endpoint: switches app into failure mode.
	http.HandleFunc("/fail", failHandler)
	// Load endpoint: simulates CPU pressure for a short fixed duration.
	http.HandleFunc("/load", loadHandler)

	log.Println("KubePulse API running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// homeHandler handles GET /.
// It is a simple landing endpoint used to verify the service process is alive.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "KubePulse AWS is running")
}

// healthHandler handles GET /health.
// This endpoint represents a liveness-style signal for monitoring systems.
// If failure mode is active (isFailing == true), it returns HTTP 500.
// Otherwise, it returns HTTP 200 with a healthy message.
func healthHandler(w http.ResponseWriter, r *http.Request) {
	if isFailing {
		http.Error(w, "service unhealthy", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "healthy")
}

// readyHandler handles GET /ready.
// This endpoint represents a readiness-style signal and currently returns ready
// unconditionally for Day 1.
func readyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ready")
}

// failHandler handles GET /fail.
// It simulates a failure event by setting isFailing to true.
// After this endpoint is called, /health will begin returning HTTP 500.
// This helps test how monitoring and orchestration react to unhealthy services.
func failHandler(w http.ResponseWriter, r *http.Request) {
	isFailing = true
	http.Error(w, "failure mode activated", http.StatusInternalServerError)
}

// loadHandler handles GET /load.
// It simulates CPU load by running a busy loop for about 5 seconds.
// The loop repeatedly performs arithmetic work to keep the CPU active,
// allowing local stress testing of the service under higher compute pressure.
func loadHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	for time.Since(start) < 5*time.Second {
		_ = 999999 * 999999
	}

	fmt.Fprintln(w, "CPU load simulated for 5 seconds")
}
