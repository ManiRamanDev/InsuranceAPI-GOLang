package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkContentTypeValidation(b *testing.B) {
	next := http.HandlerFunc(func(http.ResponseWriter, *http.Request) {})
	handler := ContentTypeValidation(next)

	b.Run("allow_json", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			req := httptest.NewRequest(http.MethodPost, "/claims", nil)
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)
			if rr.Code != http.StatusOK {
				b.Fatalf("expected %d, got %d", http.StatusOK, rr.Code)
			}
		}
	})

	b.Run("reject_non_json", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			req := httptest.NewRequest(http.MethodPost, "/claims", nil)
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)
			if rr.Code != http.StatusUnsupportedMediaType {
				b.Fatalf("expected %d, got %d", http.StatusUnsupportedMediaType, rr.Code)
			}
		}
	})
}
