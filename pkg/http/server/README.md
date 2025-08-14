# HTTP Server Package

A reusable HTTP server package with common functionality for building web services and APIs.

## Features

- **Configurable Server**: Easy configuration with sensible defaults
- **Graceful Shutdown**: Proper shutdown handling with signal management
- **Built-in Middleware**: Logging, CORS, and panic recovery
- **Health Checks**: Built-in health, readiness, and liveness endpoints
- **Response Helpers**: Standardized JSON response formatting
- **Request Utilities**: Common request parsing and validation functions
- **Route Management**: Simple API for adding routes
- **Clean Architecture**: Internal packages for better organization and encapsulation

## Installation

First, add the Chi router dependency:

```bash
go get github.com/go-chi/chi/v5
```

## Quick Start

```go
package main

import (
    "log"
    "net/http"
    
    "github.com/patraden/code-with-kids/pkg/http/server"
)

func main() {
    // Create server with default configuration
    srv := server.New(nil)
    
    // Add health check routes
    srv.AddHealthRoutes()
    
    // Add your custom routes
    srv.AddGET("/api/hello", func(w http.ResponseWriter, r *http.Request) {
        server.SuccessResponse(w, map[string]string{
            "message": "Hello, World!",
        }, "Success")
    })
    
    // Start server with graceful shutdown
    log.Fatal(srv.StartWithGracefulShutdown())
}
```

## Configuration

You can customize the server configuration:

```go
config := &server.Config{
    Port:         3000,
    ReadTimeout:  15 * time.Second,
    WriteTimeout: 15 * time.Second,
    IdleTimeout:  60 * time.Second,
    Host:         "localhost",
}

srv := server.New(config)
```

## Available Endpoints

When you call `srv.AddHealthRoutes()`, the following endpoints are automatically added:

- `GET /health` - Health check with detailed system information
- `GET /ready` - Readiness check for Kubernetes deployments
- `GET /live` - Liveness check for Kubernetes deployments
- `GET /info` - Server information and runtime details

## Response Helpers

The package provides standardized response functions:

```go
// Success responses
server.SuccessResponse(w, data, "Operation successful")
server.Created(w, data, "Resource created")
server.NoContent(w)

// Error responses
server.BadRequest(w, "Invalid input")
server.Unauthorized(w, "Authentication required")
server.Forbidden(w, "Access denied")
server.NotFound(w, "Resource not found")
server.InternalServerError(w, "Something went wrong")
```

## Request Utilities

Common request parsing functions:

```go
// Parse JSON body
var user User
if err := server.ParseJSON(r, &user); err != nil {
    server.BadRequest(w, "Invalid JSON")
    return
}

// Get query parameters
name := server.GetQueryParam(r, "name")
age, err := server.GetQueryParamInt(r, "age")
active, err := server.GetQueryParamBool(r, "active")

// Get headers
auth := server.GetAuthorizationHeader(r)
contentType := server.GetContentType(r)
userAgent := server.GetUserAgent(r)
clientIP := server.GetClientIP(r)
```

## Adding Routes

The server provides convenient methods for adding routes:

```go
srv.AddGET("/users", getUsersHandler)
srv.AddPOST("/users", createUserHandler)
srv.AddPUT("/users/{id}", updateUserHandler)
srv.AddDELETE("/users/{id}", deleteUserHandler)
srv.AddPATCH("/users/{id}", patchUserHandler)

// Or use the generic method
srv.AddRoute("GET", "/custom", customHandler)
```

## Middleware

The server automatically includes these middleware:

1. **Logging Middleware**: Logs all requests with method, path, status code, and duration
2. **CORS Middleware**: Adds CORS headers for cross-origin requests
3. **Recovery Middleware**: Recovers from panics and returns 500 errors

## Package Structure

The package is organized using internal packages for better encapsulation:

```
pkg/http/server/
├── server.go                    # Main public API
├── server_test.go              # Tests for public API
├── README.md                   # Documentation
├── example/                    # Usage examples
└── internal/                   # Internal implementation
    ├── middleware/             # Middleware implementations
    ├── response/               # Response utilities
    ├── request/                # Request utilities
    └── health/                 # Health check handlers
```

This structure ensures that internal implementation details are not exposed to external consumers while maintaining a clean public API.

## Example Usage

See the `example/` directory for a complete working example that demonstrates:

- Custom route handlers
- JSON request/response handling
- Error handling
- Health check endpoints

To run the example:

```bash
cd pkg/http/server/example
go run main.go
```

Then visit `http://localhost:3000` to see the API in action.

## API Reference

### Server Methods

- `New(config *Config) *Server` - Create a new server
- `Start() error` - Start the server (blocks)
- `StartWithGracefulShutdown() error` - Start with graceful shutdown
- `Stop(ctx context.Context) error` - Gracefully stop the server
- `Router() *chi.Mux` - Get the underlying router
- `AddRoute(method, path string, handler http.HandlerFunc)` - Add a route
- `AddGET(path string, handler http.HandlerFunc)` - Add GET route
- `AddPOST(path string, handler http.HandlerFunc)` - Add POST route
- `AddPUT(path string, handler http.HandlerFunc)` - Add PUT route
- `AddDELETE(path string, handler http.HandlerFunc)` - Add DELETE route
- `AddPATCH(path string, handler http.HandlerFunc)` - Add PATCH route
- `AddHealthRoutes()` - Add health check endpoints

### Response Functions

- `JSONResponse(w, statusCode, data)` - Send JSON response
- `SuccessResponse(w, data, message)` - Send success response
- `ErrorResponse(w, statusCode, message)` - Send error response
- `BadRequest(w, message)` - Send 400 response
- `Unauthorized(w, message)` - Send 401 response
- `Forbidden(w, message)` - Send 403 response
- `NotFound(w, message)` - Send 404 response
- `InternalServerError(w, message)` - Send 500 response
- `Created(w, data, message)` - Send 201 response
- `NoContent(w)` - Send 204 response

### Request Functions

- `ParseJSON(r, v)` - Parse JSON request body
- `GetQueryParam(r, key)` - Get query parameter as string
- `GetQueryParamInt(r, key)` - Get query parameter as int
- `GetQueryParamBool(r, key)` - Get query parameter as bool
- `GetHeader(r, key)` - Get header value
- `GetAuthorizationHeader(r)` - Get Authorization header
- `GetContentType(r)` - Get Content-Type header
- `IsJSONRequest(r)` - Check if request has JSON content type
- `GetUserAgent(r)` - Get User-Agent header
- `GetClientIP(r)` - Get client IP address
- `ValidateRequiredFields(data, required)` - Validate required fields

## License

This package is part of the code-with-kids project.
