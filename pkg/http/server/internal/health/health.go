package health

import (
	"net/http"
	"runtime"
	"time"

	"github.com/patraden/code-with-kids/pkg/http/server/internal/response"
)

// HealthStatus represents the health status of the application
type HealthStatus struct {
	Status    string            `json:"status"`
	Timestamp time.Time         `json:"timestamp"`
	Uptime    string            `json:"uptime"`
	Version   string            `json:"version,omitempty"`
	Memory    MemoryStats       `json:"memory"`
	Runtime   RuntimeStats      `json:"runtime"`
	Services  map[string]string `json:"services,omitempty"`
}

// MemoryStats represents memory usage statistics
type MemoryStats struct {
	Alloc      uint64 `json:"alloc"`
	TotalAlloc uint64 `json:"total_alloc"`
	Sys        uint64 `json:"sys"`
	NumGC      uint32 `json:"num_gc"`
}

// RuntimeStats represents runtime statistics
type RuntimeStats struct {
	Goroutines int `json:"goroutines"`
	Threads    int `json:"threads"`
}

var startTime = time.Now()

// HealthCheckHandler handles health check requests
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	status := HealthStatus{
		Status:    "healthy",
		Timestamp: time.Now(),
		Uptime:    time.Since(startTime).String(),
		Memory: MemoryStats{
			Alloc:      memStats.Alloc,
			TotalAlloc: memStats.TotalAlloc,
			Sys:        memStats.Sys,
			NumGC:      memStats.NumGC,
		},
		Runtime: RuntimeStats{
			Goroutines: runtime.NumGoroutine(),
			Threads:    runtime.GOMAXPROCS(0),
		},
		Services: make(map[string]string),
	}

	response.JSON(w, http.StatusOK, status)
}

// ReadinessHandler handles readiness check requests
func ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	// Add your readiness checks here
	// For example, check database connectivity, external services, etc.

	status := HealthStatus{
		Status:    "ready",
		Timestamp: time.Now(),
		Uptime:    time.Since(startTime).String(),
	}

	response.JSON(w, http.StatusOK, status)
}

// LivenessHandler handles liveness check requests
func LivenessHandler(w http.ResponseWriter, r *http.Request) {
	status := HealthStatus{
		Status:    "alive",
		Timestamp: time.Now(),
		Uptime:    time.Since(startTime).String(),
	}

	response.JSON(w, http.StatusOK, status)
}

// InfoHandler provides basic application information
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	info := map[string]interface{}{
		"name":        "HTTP Server",
		"version":     "1.0.0",
		"description": "A reusable HTTP server package",
		"started_at":  startTime,
		"uptime":      time.Since(startTime).String(),
		"go_version":  runtime.Version(),
		"os":          runtime.GOOS,
		"arch":        runtime.GOARCH,
	}

	response.JSON(w, http.StatusOK, info)
}
