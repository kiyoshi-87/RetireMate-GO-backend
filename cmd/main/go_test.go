// healthcheck_test.go
package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/kiyoshi-87/RetireMate-GO-backend/pkg/controllers"
)

func TestHealthCheck(t *testing.T) {
	// Create a request
	req, err := http.NewRequest("GET", "/healthcheck", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Call the health check handler with the created request and recorder
	controllers.HealthCheck(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expectedBody := `{"status":"ok"}`

	//Log purpose -- IGNORE LATER
	// t.Logf("Actual Body (Bytes): %v", rr.Body.Bytes())
	// t.Logf("Expected Body (Bytes): %v", []byte(expectedBody))

	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expectedBody) {
		t.Errorf("Handler returned unexpected body: got %v, want %v", rr.Body.String(), expectedBody)
	}

}
