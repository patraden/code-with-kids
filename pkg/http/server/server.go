package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/patraden/code-with-kids/pkg/http/server/internal/health"
	"github.com/patraden/code-with-kids/pkg/http/server/internal/middleware"
	"github.com/patraden/code-with-kids/pkg/http/server/internal/request"
	"github.com/patraden/code-with-kids/pkg/http/server/internal/response"
)

// Server represents an HTTP server with common functionality
type Server struct {
	router *chi.Mux
	server *http.Server
	config *Config
}

// Config holds server configuration
type Config struct {
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
	Host         string
}

// DefaultConfig returns a default server configuration
func DefaultConfig() *Config {
	return &Config{
		Port:         8080,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
		Host:         "0.0.0.0",
	}
}

// New creates a new HTTP server with the given configuration
func New(config *Config) *Server {
	if config == nil {
		config = DefaultConfig()
	}

	router := chi.NewRouter()

	// Add common middleware
	router.Use(middleware.Logging)
	router.Use(middleware.CORS)
	router.Use(middleware.Recovery)

	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", config.Host, config.Port),
		Handler:      router,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		IdleTimeout:  config.IdleTimeout,
	}

	return &Server{
		router: router,
		server: server,
		config: config,
	}
}

// Router returns the underlying chi router for adding routes
func (s *Server) Router() *chi.Mux {
	return s.router
}

// AddRoute adds a route with the given method, path, and handler
func (s *Server) AddRoute(method, path string, handler http.HandlerFunc) {
	s.router.MethodFunc(method, path, handler)
}

// AddGET adds a GET route
func (s *Server) AddGET(path string, handler http.HandlerFunc) {
	s.AddRoute(http.MethodGet, path, handler)
}

// AddPOST adds a POST route
func (s *Server) AddPOST(path string, handler http.HandlerFunc) {
	s.AddRoute(http.MethodPost, path, handler)
}

// AddPUT adds a PUT route
func (s *Server) AddPUT(path string, handler http.HandlerFunc) {
	s.AddRoute(http.MethodPut, path, handler)
}

// AddDELETE adds a DELETE route
func (s *Server) AddDELETE(path string, handler http.HandlerFunc) {
	s.AddRoute(http.MethodDelete, path, handler)
}

// AddPATCH adds a PATCH route
func (s *Server) AddPATCH(path string, handler http.HandlerFunc) {
	s.AddRoute(http.MethodPatch, path, handler)
}

// Start starts the server and blocks until it's stopped
func (s *Server) Start() error {
	fmt.Printf("Server starting on %s\n", s.server.Addr)
	return s.server.ListenAndServe()
}

// StartWithGracefulShutdown starts the server with graceful shutdown handling
func (s *Server) StartWithGracefulShutdown() error {
	// Create a channel to listen for errors coming from the listener
	serverErrors := make(chan error, 1)

	// Start the server in a goroutine
	go func() {
		fmt.Printf("Server starting on %s\n", s.server.Addr)
		serverErrors <- s.server.ListenAndServe()
	}()

	// Create a channel to listen for an interrupt or terminate signal from the OS
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Blocking select waiting for either a server error or a shutdown signal
	select {
	case err := <-serverErrors:
		return fmt.Errorf("error starting server: %w", err)

	case sig := <-shutdown:
		fmt.Printf("Start shutdown... Signal: %v\n", sig)

		// Give outstanding requests a deadline for completion
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		// Gracefully shutdown the server
		if err := s.server.Shutdown(ctx); err != nil {
			fmt.Printf("Could not stop server gracefully: %v\n", err)
			if err := s.server.Close(); err != nil {
				return fmt.Errorf("could not force close server: %w", err)
			}
		}
	}

	return nil
}

// Stop gracefully stops the server
func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

// GetServer returns the underlying http.Server
func (s *Server) GetServer() *http.Server {
	return s.server
}

// AddHealthRoutes adds common health check routes to the server
func (s *Server) AddHealthRoutes() {
	s.AddGET("/health", health.HealthCheckHandler)
	s.AddGET("/ready", health.ReadinessHandler)
	s.AddGET("/live", health.LivenessHandler)
	s.AddGET("/info", health.InfoHandler)
}

// Response helpers - re-export from internal package for convenience
func SuccessResponse(w http.ResponseWriter, data interface{}, message string) {
	response.Success(w, data, message)
}

func ErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	response.Error(w, statusCode, message)
}

func BadRequest(w http.ResponseWriter, message string) {
	response.BadRequest(w, message)
}

func Unauthorized(w http.ResponseWriter, message string) {
	response.Unauthorized(w, message)
}

func Forbidden(w http.ResponseWriter, message string) {
	response.Forbidden(w, message)
}

func NotFound(w http.ResponseWriter, message string) {
	response.NotFound(w, message)
}

func InternalServerError(w http.ResponseWriter, message string) {
	response.InternalServerError(w, message)
}

func Created(w http.ResponseWriter, data interface{}, message string) {
	response.Created(w, data, message)
}

func NoContent(w http.ResponseWriter) {
	response.NoContent(w)
}

// Request helpers - re-export from internal package for convenience
func ParseJSON(r *http.Request, v interface{}) error {
	return request.ParseJSON(r, v)
}

func GetQueryParam(r *http.Request, key string) string {
	return request.GetQueryParam(r, key)
}

func GetQueryParamInt(r *http.Request, key string) (int, error) {
	return request.GetQueryParamInt(r, key)
}

func GetQueryParamBool(r *http.Request, key string) (bool, error) {
	return request.GetQueryParamBool(r, key)
}

func GetPathParam(r *http.Request, key string) string {
	return request.GetPathParam(r, key)
}

func GetHeader(r *http.Request, key string) string {
	return request.GetHeader(r, key)
}

func GetAuthorizationHeader(r *http.Request) string {
	return request.GetAuthorizationHeader(r)
}

func GetContentType(r *http.Request) string {
	return request.GetContentType(r)
}

func IsJSONRequest(r *http.Request) bool {
	return request.IsJSONRequest(r)
}

func ValidateRequiredFields(data map[string]interface{}, required []string) []string {
	return request.ValidateRequiredFields(data, required)
}

func GetUserAgent(r *http.Request) string {
	return request.GetUserAgent(r)
}

func GetClientIP(r *http.Request) string {
	return request.GetClientIP(r)
}
