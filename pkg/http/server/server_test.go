package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	// Test with default config
	srv := New(nil)
	if srv == nil {
		t.Fatal("Expected server to be created")
	}

	// Test with custom config
	config := &Config{
		Port:         9090,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
		Host:         "127.0.0.1",
	}

	srv2 := New(config)
	if srv2 == nil {
		t.Fatal("Expected server to be created with custom config")
	}
}

func TestAddRoutes(t *testing.T) {
	srv := New(nil)

	// Test adding routes
	srv.AddGET("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test"))
	})

	// Test that route was added
	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()

	srv.Router().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	if w.Body.String() != "test" {
		t.Errorf("Expected body 'test', got '%s'", w.Body.String())
	}
}

func TestResponseHelpers(t *testing.T) {
	w := httptest.NewRecorder()

	// Test SuccessResponse
	SuccessResponse(w, map[string]string{"key": "value"}, "Success")

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	// Test BadRequest
	w2 := httptest.NewRecorder()
	BadRequest(w2, "Bad request")

	if w2.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w2.Code)
	}

	// Test NotFound
	w3 := httptest.NewRecorder()
	NotFound(w3, "Not found")

	if w3.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", w3.Code)
	}
}

func TestRequestHelpers(t *testing.T) {
	// Test GetQueryParam
	req := httptest.NewRequest("GET", "/test?name=john&age=25", nil)

	name := GetQueryParam(req, "name")
	if name != "john" {
		t.Errorf("Expected 'john', got '%s'", name)
	}

	age, err := GetQueryParamInt(req, "age")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if age != 25 {
		t.Errorf("Expected 25, got %d", age)
	}

	// Test GetHeader
	req.Header.Set("Authorization", "Bearer token")
	auth := GetAuthorizationHeader(req)
	if auth != "Bearer token" {
		t.Errorf("Expected 'Bearer token', got '%s'", auth)
	}
}

func TestHealthRoutes(t *testing.T) {
	srv := New(nil)
	srv.AddHealthRoutes()

	// Test health endpoint
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	srv.Router().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	// Test info endpoint
	req2 := httptest.NewRequest("GET", "/info", nil)
	w2 := httptest.NewRecorder()

	srv.Router().ServeHTTP(w2, req2)

	if w2.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w2.Code)
	}
}
