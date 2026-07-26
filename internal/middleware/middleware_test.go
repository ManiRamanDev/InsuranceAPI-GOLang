package middleware

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"insurance-api/internal/config"
)

func TestContentTypeValidation(t *testing.T) {
	tests := []struct {
		name        string
		method      string
		contentType string
		wantCode    int
		wantCalled  bool
	}{
		{name: "rejects missing json content type", method: http.MethodPost, wantCode: http.StatusUnsupportedMediaType, wantCalled: false},
		{name: "allows json content type", method: http.MethodPut, contentType: "application/json", wantCode: http.StatusOK, wantCalled: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			called := false
			next := http.HandlerFunc(func(http.ResponseWriter, *http.Request) { called = true })
			handler := ContentTypeValidation(next)

			req := httptest.NewRequest(tt.method, "/claims", nil)
			if tt.contentType != "" {
				req.Header.Set("Content-Type", tt.contentType)
			}
			rr := httptest.NewRecorder()

			handler.ServeHTTP(rr, req)

			if rr.Code != tt.wantCode {
				t.Fatalf("expected %d, got %d", tt.wantCode, rr.Code)
			}
			if called != tt.wantCalled {
				t.Fatalf("expected next handler called=%v, got %v", tt.wantCalled, called)
			}
		})
	}
}

func TestRecovery(t *testing.T) {
	t.Setenv("INSURANCE_LOG_DIRECTORY", t.TempDir())
	t.Setenv("INSURANCE_LOG_FILE", "middleware.log")
	t.Setenv("INSURANCE_LOG_LEVEL", config.DefaultLogLevel)

	handler := Recovery(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		panic("boom")
	}))

	req := httptest.NewRequest(http.MethodGet, "/claims", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Fatalf("expected %d, got %d", http.StatusInternalServerError, rr.Code)
	}
	if body := rr.Body.String(); body != "Internal Server Error\n" {
		t.Fatalf("unexpected body %q", body)
	}
}

func TestRequestLogging(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("INSURANCE_LOG_DIRECTORY", dir)
	t.Setenv("INSURANCE_LOG_FILE", "request.log")
	t.Setenv("INSURANCE_LOG_LEVEL", config.DefaultLogLevel)

	called := false
	handler := RequestLogging(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if !called {
		t.Fatal("next handler should be called")
	}
	if rr.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, rr.Code)
	}
	if _, err := os.Stat(filepath.Join(dir, "request.log")); err != nil {
		t.Fatalf("expected log file to be created: %v", err)
	}
}
