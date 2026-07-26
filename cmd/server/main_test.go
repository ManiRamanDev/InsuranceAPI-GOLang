package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterHealthEndpoint(t *testing.T) {
	mux := http.NewServeMux()
	registerHealthEndpoint(mux)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rr := httptest.NewRecorder()

	mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, rr.Code)
	}
	if rr.Body.String() != "Insurance API is running" {
		t.Fatalf("unexpected body %q", rr.Body.String())
	}
}
