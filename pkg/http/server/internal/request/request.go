package request

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// ParseJSON parses JSON from request body into the given struct
func ParseJSON(r *http.Request, v interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.Unmarshal(body, v)
}

// GetQueryParam gets a query parameter as string
func GetQueryParam(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}

// GetQueryParamInt gets a query parameter as int
func GetQueryParamInt(r *http.Request, key string) (int, error) {
	val := r.URL.Query().Get(key)
	if val == "" {
		return 0, nil
	}
	return strconv.Atoi(val)
}

// GetQueryParamBool gets a query parameter as bool
func GetQueryParamBool(r *http.Request, key string) (bool, error) {
	val := r.URL.Query().Get(key)
	if val == "" {
		return false, nil
	}
	return strconv.ParseBool(val)
}

// GetPathParam gets a path parameter from the request
func GetPathParam(r *http.Request, key string) string {
	return chi.URLParam(r, key)
}

// GetHeader gets a header value from the request
func GetHeader(r *http.Request, key string) string {
	return r.Header.Get(key)
}

// GetAuthorizationHeader gets the Authorization header
func GetAuthorizationHeader(r *http.Request) string {
	return GetHeader(r, "Authorization")
}

// GetContentType gets the Content-Type header
func GetContentType(r *http.Request) string {
	return GetHeader(r, "Content-Type")
}

// IsJSONRequest checks if the request has JSON content type
func IsJSONRequest(r *http.Request) bool {
	contentType := GetContentType(r)
	return contentType == "application/json"
}

// ValidateRequiredFields validates that all required fields are present in a map
func ValidateRequiredFields(data map[string]interface{}, required []string) []string {
	var missing []string
	for _, field := range required {
		if _, exists := data[field]; !exists {
			missing = append(missing, field)
		}
	}
	return missing
}

// GetUserAgent gets the User-Agent header
func GetUserAgent(r *http.Request) string {
	return GetHeader(r, "User-Agent")
}

// GetClientIP gets the client IP address
func GetClientIP(r *http.Request) string {
	// Check for forwarded headers first
	if ip := GetHeader(r, "X-Forwarded-For"); ip != "" {
		return ip
	}
	if ip := GetHeader(r, "X-Real-IP"); ip != "" {
		return ip
	}

	// Fall back to remote address
	return r.RemoteAddr
}
