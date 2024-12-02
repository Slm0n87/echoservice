package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEchoHandler(t *testing.T) {
	// Prepare a request with headers
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("X-Test-Header", "TestValue")
	req.Header.Set("X-Another-Header", "AnotherValue")

	// Create a response recorder to capture the response
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(echoHandler)

	// Serve the HTTP request
	handler.ServeHTTP(recorder, req)

	// Check the status code
	if status := recorder.Code; status != http.StatusOK {
		t.Fatalf("Expected status code %d, got %d", http.StatusOK, status)
	}

	// Parse the response body
	var headers map[string][]string
	if err := json.NewDecoder(recorder.Body).Decode(&headers); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// Verify headers
	expectedHeaders := map[string][]string{
		"X-Test-Header":    {"TestValue"},
		"X-Another-Header": {"AnotherValue"},
	}
	for key, expectedValues := range expectedHeaders {
		actualValues, ok := headers[key]
		if !ok {
			t.Errorf("Missing header: %s", key)
			continue
		}
		for i, expectedValue := range expectedValues {
			if actualValues[i] != expectedValue {
				t.Errorf("Header mismatch for %s: expected %s, got %s", key, expectedValue, actualValues[i])
			}
		}
	}
}
