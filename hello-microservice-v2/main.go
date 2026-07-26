package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var (
	breakLiveness  bool
	breakReadiness bool
	mu             sync.Mutex
)

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/healthz", healthHandler)

	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", nil)
}

// Hello handler
func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

// Health handler (controls liveness & readiness)
func healthHandler(w http.ResponseWriter, r *http.Request) {
	breakType := r.URL.Query().Get("break")
	timeStr := r.URL.Query().Get("time")

	// If a break is requested
	if breakType != "" && timeStr != "" {
		seconds, err := strconv.Atoi(timeStr)
		if err == nil && seconds > 0 {
			go triggerBreak(breakType, seconds)
			fmt.Fprintf(w, "Triggered %s break for %d seconds\n", breakType, seconds)
			return
		}
	}

	// Default health check
	mu.Lock()
	defer mu.Unlock()

	if breakLiveness {
		http.Error(w, "Liveness probe failed", http.StatusInternalServerError)
		return
	}

	if breakReadiness {
		http.Error(w, "Readiness probe failed", http.StatusServiceUnavailable)
		return
	}

	fmt.Fprint(w, "OK")
}

// Trigger liveness/readiness break
func triggerBreak(breakType string, seconds int) {
	mu.Lock()
	if breakType == "liveness" {
		breakLiveness = true
	} else if breakType == "readiness" {
		breakReadiness = true
	}
	mu.Unlock()

	time.Sleep(time.Duration(seconds) * time.Second)

	mu.Lock()
	if breakType == "liveness" {
		breakLiveness = false
	} else if breakType == "readiness" {
		breakReadiness = false
	}
	mu.Unlock()
}
