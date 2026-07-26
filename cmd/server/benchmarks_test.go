package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkRegisterHealthEndpoint(b *testing.B) {
	mux := http.NewServeMux()
	registerHealthEndpoint(mux)

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodGet, "/health", nil)
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			b.Fatalf("expected %d, got %d", http.StatusOK, rr.Code)
		}
	}
}
