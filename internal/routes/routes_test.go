package routes

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"insurance-api/internal/handlers"
	"insurance-api/internal/models"
)

type noopCustomerService struct{}

func (noopCustomerService) Create(*models.Customer) error                     { return nil }
func (noopCustomerService) GetByID(uint) (*models.Customer, error)            { return &models.Customer{}, nil }
func (noopCustomerService) GetAll() ([]models.Customer, error)                { return nil, nil }
func (noopCustomerService) Update(*models.Customer) error                     { return nil }
func (noopCustomerService) Delete(uint) error                                 { return nil }

type noopPolicyService struct{}

func (noopPolicyService) Create(*models.Policy) error                         { return nil }
func (noopPolicyService) GetByID(uint) (*models.Policy, error)                { return &models.Policy{}, nil }
func (noopPolicyService) GetAll() ([]models.Policy, error)                    { return nil, nil }
func (noopPolicyService) Update(*models.Policy) error                         { return nil }
func (noopPolicyService) Delete(uint) error                                   { return nil }

type noopCustomerPolicyService struct{}

func (noopCustomerPolicyService) Create(*models.CustomerPolicy) error         { return nil }
func (noopCustomerPolicyService) GetByID(uint) (*models.CustomerPolicy, error) { return &models.CustomerPolicy{}, nil }
func (noopCustomerPolicyService) GetAll() ([]models.CustomerPolicy, error)    { return nil, nil }
func (noopCustomerPolicyService) Update(*models.CustomerPolicy) error          { return nil }
func (noopCustomerPolicyService) Delete(uint) error                            { return nil }
func (noopCustomerPolicyService) GetByCustomerID(uint) ([]models.CustomerPolicy, error) {
	return nil, nil
}
func (noopCustomerPolicyService) GetByPolicyID(uint) ([]models.CustomerPolicy, error) {
	return nil, nil
}

type noopClaimService struct{}

func (noopClaimService) Create(*models.Claim) error                            { return nil }
func (noopClaimService) GetByID(uint) (*models.Claim, error)                   { return &models.Claim{}, nil }
func (noopClaimService) GetAll() ([]models.Claim, error)                       { return nil, nil }
func (noopClaimService) Update(*models.Claim) error                            { return nil }
func (noopClaimService) Delete(uint) error                                     { return nil }
func (noopClaimService) GetByCustomerPolicyID(uint) ([]models.Claim, error)    { return nil, nil }

func TestRegisterRoutes(t *testing.T) {
	mux := http.NewServeMux()
	RegisterRoutes(
		mux,
		handlers.NewCustomerHandler(noopCustomerService{}),
		handlers.NewPolicyHandler(noopPolicyService{}),
		handlers.NewCustomerPolicyHandler(noopCustomerPolicyService{}),
		handlers.NewClaimHandler(noopClaimService{}),
	)

	checks := []struct {
		name       string
		method     string
		path       string
		body       string
		contentType string
		wantCode   int
		wantPattern string
	}{
		{name: "customer get all", method: http.MethodGet, path: "/customers", wantCode: http.StatusOK, wantPattern: "GET /customers"},
		{name: "customer create", method: http.MethodPost, path: "/customers", body: `{"first_name":"Raj","last_name":"Kumar","email":"raj@example.com","phone_number":"123","address":"Pune"}`, wantCode: http.StatusCreated, wantPattern: "POST /customers"},
		{name: "customer get by id", method: http.MethodGet, path: "/customers/7", wantCode: http.StatusOK, wantPattern: "GET /customers/{id}"},
		{name: "customer update", method: http.MethodPut, path: "/customers/7", body: `{"first_name":"Jane","last_name":"Doe","email":"jane@example.com","phone_number":"555","address":"Delhi"}`, contentType: "application/json", wantCode: http.StatusOK, wantPattern: "PUT /customers/{id}"},
		{name: "customer delete", method: http.MethodDelete, path: "/customers/7", wantCode: http.StatusNoContent, wantPattern: "DELETE /customers/{id}"},
		{name: "policy get all", method: http.MethodGet, path: "/policies", wantCode: http.StatusOK, wantPattern: "GET /policies"},
		{name: "policy create", method: http.MethodPost, path: "/policies", body: `{"policy_name":"Life","policy_type":"Term","coverage":100,"premium":10,"description":"desc"}`, wantCode: http.StatusCreated, wantPattern: "POST /policies"},
		{name: "policy get by id", method: http.MethodGet, path: "/policies/21", wantCode: http.StatusOK, wantPattern: "GET /policies/{id}"},
		{name: "policy update", method: http.MethodPut, path: "/policies/21", body: `{"policy_name":"Life","policy_type":"Term","coverage":200,"premium":20,"description":"updated"}`, contentType: "application/json", wantCode: http.StatusOK, wantPattern: "PUT /policies/{id}"},
		{name: "policy delete", method: http.MethodDelete, path: "/policies/21", wantCode: http.StatusNoContent, wantPattern: "DELETE /policies/{id}"},
		{name: "customer policy get all", method: http.MethodGet, path: "/customer-policies", wantCode: http.StatusOK, wantPattern: "GET /customer-policies"},
		{name: "customer policy create", method: http.MethodPost, path: "/customer-policies", body: `{"customer_id":1,"policy_id":2,"start_date":"2026-07-27T00:00:00Z","end_date":"2027-07-27T00:00:00Z","status":"ACTIVE"}`, wantCode: http.StatusCreated, wantPattern: "POST /customer-policies"},
		{name: "customer policy get by id", method: http.MethodGet, path: "/customer-policies/31", wantCode: http.StatusOK, wantPattern: "GET /customer-policies/{id}"},
		{name: "customer policy update", method: http.MethodPut, path: "/customer-policies/31", body: `{"start_date":"2026-07-27T00:00:00Z","end_date":"2027-07-27T00:00:00Z","status":"INACTIVE"}`, contentType: "application/json", wantCode: http.StatusOK, wantPattern: "PUT /customer-policies/{id}"},
		{name: "customer policy delete", method: http.MethodDelete, path: "/customer-policies/31", wantCode: http.StatusNoContent, wantPattern: "DELETE /customer-policies/{id}"},
		{name: "claim get all", method: http.MethodGet, path: "/claims", wantCode: http.StatusOK, wantPattern: "GET /claims"},
		{name: "claim create", method: http.MethodPost, path: "/claims", body: `{"customer_policy_id":2,"claim_amount":1000,"reason":"Accident","status":"PENDING","claim_date":"2026-07-27T00:00:00Z"}`, contentType: "application/json", wantCode: http.StatusCreated, wantPattern: "POST /claims"},
		{name: "claim get by id", method: http.MethodGet, path: "/claims/41", wantCode: http.StatusOK, wantPattern: "GET /claims/{id}"},
		{name: "claim update", method: http.MethodPut, path: "/claims/41", body: `{"claim_amount":2500,"reason":"Updated","status":"APPROVED"}`, contentType: "application/json", wantCode: http.StatusOK, wantPattern: "PUT /claims/{id}"},
		{name: "claim delete", method: http.MethodDelete, path: "/claims/41", wantCode: http.StatusNoContent, wantPattern: "DELETE /claims/{id}"},
	}

	for _, tt := range checks {
		t.Run(tt.name, func(t *testing.T) {
			var bodyReader *strings.Reader
			if tt.body != "" {
				bodyReader = strings.NewReader(tt.body)
			} else {
				bodyReader = strings.NewReader("")
			}

			req := httptest.NewRequest(tt.method, tt.path, bodyReader)
			if tt.contentType != "" {
				req.Header.Set("Content-Type", tt.contentType)
			}

			rr := httptest.NewRecorder()
			mux.ServeHTTP(rr, req)

			if rr.Code != tt.wantCode {
				t.Fatalf("expected status %d, got %d", tt.wantCode, rr.Code)
			}

			_, pattern := mux.Handler(req)
			if pattern != tt.wantPattern {
				t.Fatalf("expected pattern %q, got %q", tt.wantPattern, pattern)
			}
		})
	}
}
