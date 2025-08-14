package main

import (
	"log"
	"net/http"
	"time"

	"github.com/patraden/code-with-kids/pkg/http/server"
)

// Example types for future use
// type Example struct {
//     ID   string `json:"id"`
//     Name string `json:"name"`
// }

func main() {
	// Create server with custom configuration
	config := &server.Config{
		Port:         8888,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
		Host:         "localhost",
	}

	srv := server.New(config)

	// Add health check routes
	srv.AddHealthRoutes()

	// Add custom routes (if needed in the future)
	// srv.AddGET("/api/example", exampleHandler)

	// Serve static files
	srv.Router().Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("pkg/http/server/example/static"))))

	// Add a simple welcome route that redirects to the HTML page
	srv.AddGET("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "pkg/http/server/example/static/index.html")
	})

	log.Println("Starting server on http://localhost:8888")
	log.Println("Available endpoints:")
	log.Println("  GET  / - Welcome message")
	log.Println("  GET  /health - Health check")
	log.Println("  GET  /ready - Readiness check")
	log.Println("  GET  /live - Liveness check")
	log.Println("  GET  /info - Server info")
	log.Println("  GET  /static/* - Static files")

	// Start server with graceful shutdown
	if err := srv.StartWithGracefulShutdown(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

// Example handlers for future use
// func exampleHandler(w http.ResponseWriter, r *http.Request) {
//     server.SuccessResponse(w, map[string]string{
//         "message": "This is an example endpoint",
//     }, "Example response")
// }
