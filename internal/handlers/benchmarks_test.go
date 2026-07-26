package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"insurance-api/internal/models"
)

func BenchmarkCustomerHandler(b *testing.B) {
	createdAt := time.Date(2026, time.July, 27, 0, 0, 0, 0, time.UTC)
	handler := NewCustomerHandler(&mockCustomerService{
		CreateFn: func(customer *models.Customer) error {
			customer.ID = 11
			customer.CreatedAt = createdAt
			customer.UpdatedAt = createdAt
			return nil
		},
		GetAllFn: func() ([]models.Customer, error) {
			return []models.Customer{{ID: 1, FirstName: "Jane", LastName: "Doe", Email: "jane@example.com", Phone: "555"}}, nil
		},
	})

	createBody := `{"first_name":"Raj","last_name":"Kumar","email":"raj@example.com","phone_number":"123","address":"Pune"}`

	b.ReportAllocs()
	b.Run("Create", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			req := httptest.NewRequest(http.MethodPost, "/customers", strings.NewReader(createBody))
			rr := httptest.NewRecorder()
			handler.Create(rr, req)
			if rr.Code != http.StatusCreated {
				b.Fatalf("expected %d, got %d", http.StatusCreated, rr.Code)
			}
		}
	})

	b.Run("GetAll", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			req := httptest.NewRequest(http.MethodGet, "/customers", nil)
			rr := httptest.NewRecorder()
			handler.GetAll(rr, req)
			if rr.Code != http.StatusOK {
				b.Fatalf("expected %d, got %d", http.StatusOK, rr.Code)
			}
		}
	})
}

func BenchmarkPolicyHandler(b *testing.B) {
	createdAt := time.Date(2026, time.July, 27, 0, 0, 0, 0, time.UTC)
	handler := NewPolicyHandler(&mockPolicyService{
		CreateFn: func(policy *models.Policy) error {
			policy.ID = 21
			policy.CreatedAt = createdAt
			policy.UpdatedAt = createdAt
			return nil
		},
		GetAllFn: func() ([]models.Policy, error) {
			return []models.Policy{{ID: 1, PolicyName: "Life", Coverage: 100, Premium: 10}}, nil
		},
	})

	createBody := `{"policy_name":"Life","policy_type":"Term","coverage":100,"premium":10,"description":"desc"}`

	b.ReportAllocs()
	b.Run("Create", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			req := httptest.NewRequest(http.MethodPost, "/policies", strings.NewReader(createBody))
			rr := httptest.NewRecorder()
			handler.Create(rr, req)
			if rr.Code != http.StatusCreated {
				b.Fatalf("expected %d, got %d", http.StatusCreated, rr.Code)
			}
		}
	})

	b.Run("GetAll", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			req := httptest.NewRequest(http.MethodGet, "/policies", nil)
			rr := httptest.NewRecorder()
			handler.GetAll(rr, req)
			if rr.Code != http.StatusOK {
				b.Fatalf("expected %d, got %d", http.StatusOK, rr.Code)
			}
		}
	})
}

func BenchmarkCustomerPolicyHandler(b *testing.B) {
	createdAt := time.Date(2026, time.July, 27, 0, 0, 0, 0, time.UTC)
	handler := NewCustomerPolicyHandler(&mockCustomerPolicyService{
		CreateFn: func(customerPolicy *models.CustomerPolicy) error {
			customerPolicy.ID = 31
			customerPolicy.CreatedAt = createdAt
			customerPolicy.UpdatedAt = createdAt
			return nil
		},
		GetAllFn: func() ([]models.CustomerPolicy, error) {
			return []models.CustomerPolicy{{ID: 1, CustomerID: 1, PolicyID: 2, Status: "ACTIVE"}}, nil
		},
	})

	createBody := `{"customer_id":1,"policy_id":2,"start_date":"2026-07-27T00:00:00Z","end_date":"2027-07-27T00:00:00Z","status":"ACTIVE"}`

	b.ReportAllocs()
	b.Run("Create", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			req := httptest.NewRequest(http.MethodPost, "/customer-policies", strings.NewReader(createBody))
			rr := httptest.NewRecorder()
			handler.Create(rr, req)
			if rr.Code != http.StatusCreated {
				b.Fatalf("expected %d, got %d", http.StatusCreated, rr.Code)
			}
		}
	})

	b.Run("GetAll", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			req := httptest.NewRequest(http.MethodGet, "/customer-policies", nil)
			rr := httptest.NewRecorder()
			handler.GetAll(rr, req)
			if rr.Code != http.StatusOK {
				b.Fatalf("expected %d, got %d", http.StatusOK, rr.Code)
			}
		}
	})
}

func BenchmarkClaimHandler(b *testing.B) {
	createdAt := time.Date(2026, time.July, 27, 0, 0, 0, 0, time.UTC)
	handler := NewClaimHandler(&mockClaimService{
		CreateFn: func(claim *models.Claim) error {
			claim.ID = 41
			claim.CreatedAt = createdAt
			claim.UpdatedAt = createdAt
			return nil
		},
		GetAllFn: func() ([]models.Claim, error) {
			return []models.Claim{{ID: 1, CustomerPolicyID: 2, ClaimAmount: 1000, Reason: "Accident", Status: "PENDING", ClaimDate: createdAt}}, nil
		},
	})

	createBody := `{"customer_policy_id":2,"claim_amount":1000,"reason":"Accident","status":"PENDING","claim_date":"2026-07-27T00:00:00Z"}`

	b.ReportAllocs()
	b.Run("Create", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			req := httptest.NewRequest(http.MethodPost, "/claims", strings.NewReader(createBody))
			rr := httptest.NewRecorder()
			handler.Create(rr, req)
			if rr.Code != http.StatusCreated {
				b.Fatalf("expected %d, got %d", http.StatusCreated, rr.Code)
			}
		}
	})

	b.Run("GetAll", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			req := httptest.NewRequest(http.MethodGet, "/claims", nil)
			rr := httptest.NewRecorder()
			handler.GetAll(rr, req)
			if rr.Code != http.StatusOK {
				b.Fatalf("expected %d, got %d", http.StatusOK, rr.Code)
			}
		}
	})
}
