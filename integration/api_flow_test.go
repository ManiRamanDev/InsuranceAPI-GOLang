package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"insurance-api/internal/database"
	"insurance-api/internal/dto"
	"insurance-api/internal/handlers"
	"insurance-api/internal/repository"
	"insurance-api/internal/routes"
	"insurance-api/internal/service"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type integrationApp struct {
	server          *httptest.Server
	db              *gorm.DB
	customerRepo    repository.CustomerRepository
	policyRepo      repository.PolicyRepository
	customerPolRepo repository.CustomerPolicyRepository
	claimRepo       repository.ClaimRepository
}

func newIntegrationApp(t *testing.T) *integrationApp {
	t.Helper()

	db := newSQLiteDB(t)
	if err := database.InitSchema(db); err != nil {
		t.Fatalf("init schema: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Insurance API is running"))
	})

	customerRepo := repository.NewCustomerRepository(db)
	policyRepo := repository.NewPolicyRepository(db)
	customerPolRepo := repository.NewCustomerPolicyRepository(db)
	claimRepo := repository.NewClaimRepository(db)

	routes.RegisterRoutes(
		mux,
		handlers.NewCustomerHandler(service.NewCustomerService(customerRepo)),
		handlers.NewPolicyHandler(service.NewPolicyService(policyRepo)),
		handlers.NewCustomerPolicyHandler(service.NewCustomerPolicyService(customerPolRepo)),
		handlers.NewClaimHandler(service.NewClaimService(claimRepo)),
	)

	server := httptest.NewServer(mux)

	return &integrationApp{
		server:          server,
		db:              db,
		customerRepo:    customerRepo,
		policyRepo:      policyRepo,
		customerPolRepo: customerPolRepo,
		claimRepo:       claimRepo,
	}
}

func newSQLiteDB(t *testing.T) *gorm.DB {
	t.Helper()

	dbFile := filepath.Join(t.TempDir(), "integration.db")
	dsn := fmt.Sprintf("file:%s?_foreign_keys=1", dbFile)
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("open sqlite db: %v", err)
	}

	return db
}

func (a *integrationApp) close() {
	if a.server != nil {
		a.server.Close()
	}
	if a.db == nil {
		return
	}
	sqlDB, err := a.db.DB()
	if err == nil && sqlDB != nil {
		_ = sqlDB.Close()
	}
}

func (a *integrationApp) doRequest(t *testing.T, method, path string, body any, contentType string) (*http.Response, []byte) {
	t.Helper()

	var reader *bytes.Reader
	if body != nil {
		payload, err := json.Marshal(body)
		if err != nil {
			t.Fatalf("marshal request body: %v", err)
		}
		reader = bytes.NewReader(payload)
	} else {
		reader = bytes.NewReader(nil)
	}

	req, err := http.NewRequest(method, a.server.URL+path, reader)
	if err != nil {
		t.Fatalf("create request: %v", err)
	}
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	resp, err := a.server.Client().Do(req)
	if err != nil {
		t.Fatalf("do request: %v", err)
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("read response body: %v", err)
	}
	return resp, bodyBytes
}

func (a *integrationApp) doRawRequest(t *testing.T, method, path, payload, contentType string) (*http.Response, []byte) {
	t.Helper()

	req, err := http.NewRequest(method, a.server.URL+path, strings.NewReader(payload))
	if err != nil {
		t.Fatalf("create request: %v", err)
	}
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	resp, err := a.server.Client().Do(req)
	if err != nil {
		t.Fatalf("do request: %v", err)
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("read response body: %v", err)
	}
	return resp, bodyBytes
}

func TestHealthEndpoint(t *testing.T) {
	a := newIntegrationApp(t)
	defer a.close()

	resp, body := a.doRequest(t, http.MethodGet, "/health", nil, "")
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}
	if string(body) != "Insurance API is running" {
		t.Fatalf("unexpected body %q", string(body))
	}
}

func TestCustomerFlow(t *testing.T) {
	a := newIntegrationApp(t)
	defer a.close()

	var created dto.CustomerResponse
	resp, body := a.doRequest(t, http.MethodPost, "/customers", map[string]any{
		"first_name":   "Raj",
		"last_name":    "Kumar",
		"email":        "raj@example.com",
		"phone_number": "123456",
		"address":      "Pune",
	}, "application/json")
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, resp.StatusCode)
	}
	if err := json.Unmarshal(body, &created); err != nil {
		t.Fatalf("unmarshal customer create response: %v", err)
	}
	if created.ID == 0 {
		t.Fatal("expected created customer ID")
	}

	stored, err := a.customerRepo.GetByID(created.ID)
	if err != nil {
		t.Fatalf("repo get by id: %v", err)
	}
	if stored.Email != "raj@example.com" {
		t.Fatalf("unexpected stored customer: %+v", stored)
	}

	resp, body = a.doRequest(t, http.MethodGet, fmt.Sprintf("/customers/%d", created.ID), nil, "")
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}
	var got dto.CustomerResponse
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal customer get response: %v", err)
	}
	if got.ID != created.ID {
		t.Fatalf("expected id %d, got %d", created.ID, got.ID)
	}

	resp, body = a.doRequest(t, http.MethodGet, "/customers", nil, "")
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}
	var list []dto.CustomerResponse
	if err := json.Unmarshal(body, &list); err != nil {
		t.Fatalf("unmarshal customer list: %v", err)
	}
	if len(list) != 1 {
		t.Fatalf("expected 1 customer, got %d", len(list))
	}

	resp, _ = a.doRequest(t, http.MethodPut, fmt.Sprintf("/customers/%d", created.ID), map[string]any{
		"first_name":   "Raj",
		"last_name":    "Updated",
		"email":        "raj.updated@example.com",
		"phone_number": "987654",
		"address":      "Mumbai",
	}, "application/json")
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}

	updated, err := a.customerRepo.GetByID(created.ID)
	if err != nil {
		t.Fatalf("repo get updated customer: %v", err)
	}
	if updated.LastName != "Updated" || updated.Email != "raj.updated@example.com" {
		t.Fatalf("unexpected updated customer: %+v", updated)
	}

	resp, _ = a.doRequest(t, http.MethodDelete, fmt.Sprintf("/customers/%d", created.ID), nil, "")
	if resp.StatusCode != http.StatusNoContent {
		t.Fatalf("expected status %d, got %d", http.StatusNoContent, resp.StatusCode)
	}

	if _, err := a.customerRepo.GetByID(created.ID); err == nil {
		t.Fatal("expected repo get by id after delete to fail")
	}

	resp, _ = a.doRequest(t, http.MethodGet, "/customers/abc", nil, "")
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}

	resp, _ = a.doRequest(t, http.MethodGet, "/customers/abc", nil, "")
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status %d for invalid id, got %d", http.StatusBadRequest, resp.StatusCode)
	}
}

func TestPolicyFlow(t *testing.T) {
	a := newIntegrationApp(t)
	defer a.close()

	resp, body := a.doRequest(t, http.MethodPost, "/policies", map[string]any{
		"policy_name": "Life",
		"policy_type": "Term",
		"coverage":    100000,
		"premium":     2500,
		"description": "Base policy",
	}, "application/json")
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, resp.StatusCode)
	}
	var created dto.PolicyResponse
	if err := json.Unmarshal(body, &created); err != nil {
		t.Fatalf("unmarshal policy create response: %v", err)
	}

	resp, _ = a.doRequest(t, http.MethodGet, "/policies", nil, "")
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}

	resp, body = a.doRequest(t, http.MethodGet, fmt.Sprintf("/policies/%d", created.ID), nil, "")
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}
	var got dto.PolicyResponse
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal policy get response: %v", err)
	}
	if got.PolicyName != "Life" {
		t.Fatalf("unexpected policy: %+v", got)
	}

	resp, _ = a.doRequest(t, http.MethodPut, fmt.Sprintf("/policies/%d", created.ID), map[string]any{
		"policy_name": "Life Plus",
		"policy_type": "Term",
		"coverage":    200000,
		"premium":     5000,
		"description": "Updated policy",
	}, "application/json")
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}

	updated, err := a.policyRepo.GetByID(created.ID)
	if err != nil {
		t.Fatalf("repo get policy: %v", err)
	}
	if updated.PolicyName != "Life Plus" || updated.Coverage != 200000 {
		t.Fatalf("unexpected updated policy: %+v", updated)
	}

	resp, _ = a.doRequest(t, http.MethodDelete, fmt.Sprintf("/policies/%d", created.ID), nil, "")
	if resp.StatusCode != http.StatusNoContent {
		t.Fatalf("expected status %d, got %d", http.StatusNoContent, resp.StatusCode)
	}

	if _, err := a.policyRepo.GetByID(created.ID); err == nil {
		t.Fatal("expected repo get by id after delete to fail")
	}

	resp, _ = a.doRequest(t, http.MethodGet, "/policies/abc", nil, "")
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}
}

func TestCustomerPolicyFlow(t *testing.T) {
	a := newIntegrationApp(t)
	defer a.close()

	customerResp, customerBody := a.doRequest(t, http.MethodPost, "/customers", map[string]any{
		"first_name":   "Policy",
		"last_name":    "Customer",
		"email":        "policy.customer@example.com",
		"phone_number": "111111",
		"address":      "Delhi",
	}, "application/json")
	if customerResp.StatusCode != http.StatusCreated {
		t.Fatalf("expected customer create status %d, got %d", http.StatusCreated, customerResp.StatusCode)
	}

	policyResp, policyBody := a.doRequest(t, http.MethodPost, "/policies", map[string]any{
		"policy_name": "Health",
		"policy_type": "Annual",
		"coverage":    50000,
		"premium":     1200,
		"description": "Health cover",
	}, "application/json")
	if policyResp.StatusCode != http.StatusCreated {
		t.Fatalf("expected policy create status %d, got %d", http.StatusCreated, policyResp.StatusCode)
	}

	var createdCustomer dto.CustomerResponse
	var createdPolicy dto.PolicyResponse
	if err := json.Unmarshal(customerBody, &createdCustomer); err != nil {
		t.Fatalf("unmarshal customer create: %v", err)
	}
	if err := json.Unmarshal(policyBody, &createdPolicy); err != nil {
		t.Fatalf("unmarshal policy create: %v", err)
	}

	resp, body := a.doRequest(t, http.MethodPost, "/customer-policies", map[string]any{
		"customer_id": createdCustomer.ID,
		"policy_id":   createdPolicy.ID,
		"start_date":  time.Now().UTC().Format(time.RFC3339),
		"end_date":    time.Now().UTC().AddDate(1, 0, 0).Format(time.RFC3339),
		"status":      "ACTIVE",
	}, "application/json")
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, resp.StatusCode)
	}
	var created dto.CustomerPolicyResponse
	if err := json.Unmarshal(body, &created); err != nil {
		t.Fatalf("unmarshal customer-policy create response: %v", err)
	}

	resp, _ = a.doRequest(t, http.MethodGet, "/customer-policies", nil, "")
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}

	resp, body = a.doRequest(t, http.MethodGet, fmt.Sprintf("/customer-policies/%d", created.ID), nil, "")
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}
	var got dto.CustomerPolicyResponse
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal customer-policy get response: %v", err)
	}
	if got.CustomerID != createdCustomer.ID || got.PolicyID != createdPolicy.ID {
		t.Fatalf("unexpected customer-policy: %+v", got)
	}

	byCustomer, err := a.customerPolRepo.GetByCustomerID(createdCustomer.ID)
	if err != nil {
		t.Fatalf("repo get by customer id: %v", err)
	}
	if len(byCustomer) != 1 {
		t.Fatalf("expected 1 customer-policy by customer, got %d", len(byCustomer))
	}

	byPolicy, err := a.customerPolRepo.GetByPolicyID(createdPolicy.ID)
	if err != nil {
		t.Fatalf("repo get by policy id: %v", err)
	}
	if len(byPolicy) != 1 {
		t.Fatalf("expected 1 customer-policy by policy, got %d", len(byPolicy))
	}

	resp, _ = a.doRequest(t, http.MethodPut, fmt.Sprintf("/customer-policies/%d", created.ID), map[string]any{
		"start_date": time.Now().UTC().AddDate(0, -1, 0).Format(time.RFC3339),
		"end_date":   time.Now().UTC().AddDate(1, 0, 0).Format(time.RFC3339),
		"status":     "INACTIVE",
	}, "application/json")
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}

	updated, err := a.customerPolRepo.GetByID(created.ID)
	if err != nil {
		t.Fatalf("repo get customer-policy: %v", err)
	}
	if updated.Status != "INACTIVE" {
		t.Fatalf("unexpected updated customer-policy: %+v", updated)
	}

	resp, _ = a.doRequest(t, http.MethodDelete, fmt.Sprintf("/customer-policies/%d", created.ID), nil, "")
	if resp.StatusCode != http.StatusNoContent {
		t.Fatalf("expected status %d, got %d", http.StatusNoContent, resp.StatusCode)
	}

	if _, err := a.customerPolRepo.GetByID(created.ID); err == nil {
		t.Fatal("expected repo get by id after delete to fail")
	}

	resp, _ = a.doRequest(t, http.MethodGet, "/customer-policies/abc", nil, "")
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}
}

func TestClaimFlow(t *testing.T) {
	a := newIntegrationApp(t)
	defer a.close()

	customerResp, customerBody := a.doRequest(t, http.MethodPost, "/customers", map[string]any{
		"first_name":   "Claim",
		"last_name":    "Customer",
		"email":        "claim.customer@example.com",
		"phone_number": "222222",
		"address":      "Chennai",
	}, "application/json")
	if customerResp.StatusCode != http.StatusCreated {
		t.Fatalf("expected customer create status %d, got %d", http.StatusCreated, customerResp.StatusCode)
	}
	policyResp, policyBody := a.doRequest(t, http.MethodPost, "/policies", map[string]any{
		"policy_name": "Motor",
		"policy_type": "Annual",
		"coverage":    300000,
		"premium":     9000,
		"description": "Motor cover",
	}, "application/json")
	if policyResp.StatusCode != http.StatusCreated {
		t.Fatalf("expected policy create status %d, got %d", http.StatusCreated, policyResp.StatusCode)
	}

	var createdCustomer dto.CustomerResponse
	var createdPolicy dto.PolicyResponse
	if err := json.Unmarshal(customerBody, &createdCustomer); err != nil {
		t.Fatalf("unmarshal customer create: %v", err)
	}
	if err := json.Unmarshal(policyBody, &createdPolicy); err != nil {
		t.Fatalf("unmarshal policy create: %v", err)
	}

	cpResp, cpBody := a.doRequest(t, http.MethodPost, "/customer-policies", map[string]any{
		"customer_id": createdCustomer.ID,
		"policy_id":   createdPolicy.ID,
		"start_date":  time.Now().UTC().Format(time.RFC3339),
		"end_date":    time.Now().UTC().AddDate(1, 0, 0).Format(time.RFC3339),
		"status":      "ACTIVE",
	}, "application/json")
	if cpResp.StatusCode != http.StatusCreated {
		t.Fatalf("expected customer-policy create status %d, got %d", http.StatusCreated, cpResp.StatusCode)
	}
	var createdCP dto.CustomerPolicyResponse
	if err := json.Unmarshal(cpBody, &createdCP); err != nil {
		t.Fatalf("unmarshal customer-policy create: %v", err)
	}

	resp, body := a.doRequest(t, http.MethodPost, "/claims", map[string]any{
		"customer_policy_id": createdCP.ID,
		"claim_amount":       50000,
		"reason":             "Accident",
		"status":             "PENDING",
		"claim_date":         time.Now().UTC().Format(time.RFC3339),
	}, "application/json")
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, resp.StatusCode)
	}
	var created dto.ClaimResponse
	if err := json.Unmarshal(body, &created); err != nil {
		t.Fatalf("unmarshal claim create response: %v", err)
	}

	resp, _ = a.doRequest(t, http.MethodGet, "/claims", nil, "")
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}

	resp, body = a.doRequest(t, http.MethodGet, fmt.Sprintf("/claims/%d", created.ID), nil, "")
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}
	var got dto.ClaimResponse
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal claim get response: %v", err)
	}
	if got.CustomerPolicyID != createdCP.ID {
		t.Fatalf("unexpected claim: %+v", got)
	}

	byCustomerPolicy, err := a.claimRepo.GetByCustomerPolicyID(createdCP.ID)
	if err != nil {
		t.Fatalf("repo get by customer policy id: %v", err)
	}
	if len(byCustomerPolicy) != 1 {
		t.Fatalf("expected 1 claim by customer policy, got %d", len(byCustomerPolicy))
	}

	resp, _ = a.doRequest(t, http.MethodPut, fmt.Sprintf("/claims/%d", created.ID), map[string]any{
		"claim_amount": 75000,
		"reason":       "Updated accident",
		"status":       "APPROVED",
	}, "application/json")
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}

	updated, err := a.claimRepo.GetByID(created.ID)
	if err != nil {
		t.Fatalf("repo get claim: %v", err)
	}
	if updated.Status != "APPROVED" || updated.ClaimAmount != 75000 {
		t.Fatalf("unexpected updated claim: %+v", updated)
	}

	resp, _ = a.doRequest(t, http.MethodDelete, fmt.Sprintf("/claims/%d", created.ID), nil, "application/json")
	if resp.StatusCode != http.StatusNoContent {
		t.Fatalf("expected status %d, got %d", http.StatusNoContent, resp.StatusCode)
	}

	if _, err := a.claimRepo.GetByID(created.ID); err == nil {
		t.Fatal("expected repo get by id after delete to fail")
	}

	resp, _ = a.doRequest(t, http.MethodPost, "/claims", map[string]any{"reason": "Missing content type"}, "")
	if resp.StatusCode != http.StatusUnsupportedMediaType {
		t.Fatalf("expected status %d, got %d", http.StatusUnsupportedMediaType, resp.StatusCode)
	}

	resp, _ = a.doRawRequest(t, http.MethodPost, "/claims", "{", "application/json")
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status %d for invalid body, got %d", http.StatusBadRequest, resp.StatusCode)
	}

	resp, _ = a.doRequest(t, http.MethodGet, "/claims/abc", nil, "")
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}
}
