package routes

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"insurance-api/internal/handlers"
)

func benchmarkMux() *http.ServeMux {
	mux := http.NewServeMux()
	RegisterRoutes(
		mux,
		handlers.NewCustomerHandler(noopCustomerService{}),
		handlers.NewPolicyHandler(noopPolicyService{}),
		handlers.NewCustomerPolicyHandler(noopCustomerPolicyService{}),
		handlers.NewClaimHandler(noopClaimService{}),
	)
	return mux
}

func BenchmarkRegisterRoutes(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		mux := http.NewServeMux()
		RegisterRoutes(
			mux,
			handlers.NewCustomerHandler(noopCustomerService{}),
			handlers.NewPolicyHandler(noopPolicyService{}),
			handlers.NewCustomerPolicyHandler(noopCustomerPolicyService{}),
			handlers.NewClaimHandler(noopClaimService{}),
		)
	}
}

func BenchmarkRouteDispatchCustomerCreate(b *testing.B) {
	mux := benchmarkMux()
	body := `{"first_name":"Raj","last_name":"Kumar","email":"raj@example.com","phone_number":"123","address":"Pune"}`

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodPost, "/customers", strings.NewReader(body))
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)
		if rr.Code != http.StatusCreated {
			b.Fatalf("expected %d, got %d", http.StatusCreated, rr.Code)
		}
	}
}

func BenchmarkRouteDispatchCustomerGetAll(b *testing.B) {
	mux := benchmarkMux()

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodGet, "/customers", nil)
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			b.Fatalf("expected %d, got %d", http.StatusOK, rr.Code)
		}
	}
}
