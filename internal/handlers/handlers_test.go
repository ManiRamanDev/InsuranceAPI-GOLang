package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"insurance-api/internal/dto"
	"insurance-api/internal/models"
)

type mockCustomerService struct {
	CreateFn  func(*models.Customer) error
	GetByIDFn func(uint) (*models.Customer, error)
	GetAllFn  func() ([]models.Customer, error)
	UpdateFn  func(*models.Customer) error
	DeleteFn  func(uint) error
}

func (m *mockCustomerService) Create(customer *models.Customer) error {
	if m.CreateFn != nil {
		return m.CreateFn(customer)
	}
	return nil
}

func (m *mockCustomerService) GetByID(id uint) (*models.Customer, error) {
	if m.GetByIDFn != nil {
		return m.GetByIDFn(id)
	}
	return &models.Customer{}, nil
}

func (m *mockCustomerService) GetAll() ([]models.Customer, error) {
	if m.GetAllFn != nil {
		return m.GetAllFn()
	}
	return nil, nil
}

func (m *mockCustomerService) Update(customer *models.Customer) error {
	if m.UpdateFn != nil {
		return m.UpdateFn(customer)
	}
	return nil
}

func (m *mockCustomerService) Delete(id uint) error {
	if m.DeleteFn != nil {
		return m.DeleteFn(id)
	}
	return nil
}

type mockPolicyService struct {
	CreateFn  func(*models.Policy) error
	GetByIDFn func(uint) (*models.Policy, error)
	GetAllFn  func() ([]models.Policy, error)
	UpdateFn  func(*models.Policy) error
	DeleteFn  func(uint) error
}

func (m *mockPolicyService) Create(policy *models.Policy) error {
	if m.CreateFn != nil {
		return m.CreateFn(policy)
	}
	return nil
}

func (m *mockPolicyService) GetByID(id uint) (*models.Policy, error) {
	if m.GetByIDFn != nil {
		return m.GetByIDFn(id)
	}
	return &models.Policy{}, nil
}

func (m *mockPolicyService) GetAll() ([]models.Policy, error) {
	if m.GetAllFn != nil {
		return m.GetAllFn()
	}
	return nil, nil
}

func (m *mockPolicyService) Update(policy *models.Policy) error {
	if m.UpdateFn != nil {
		return m.UpdateFn(policy)
	}
	return nil
}

func (m *mockPolicyService) Delete(id uint) error {
	if m.DeleteFn != nil {
		return m.DeleteFn(id)
	}
	return nil
}

type mockCustomerPolicyService struct {
	CreateFn          func(*models.CustomerPolicy) error
	GetByIDFn         func(uint) (*models.CustomerPolicy, error)
	GetAllFn          func() ([]models.CustomerPolicy, error)
	UpdateFn          func(*models.CustomerPolicy) error
	DeleteFn          func(uint) error
	GetByCustomerIDFn func(uint) ([]models.CustomerPolicy, error)
	GetByPolicyIDFn   func(uint) ([]models.CustomerPolicy, error)
}

func (m *mockCustomerPolicyService) Create(customerPolicy *models.CustomerPolicy) error {
	if m.CreateFn != nil {
		return m.CreateFn(customerPolicy)
	}
	return nil
}

func (m *mockCustomerPolicyService) GetByID(id uint) (*models.CustomerPolicy, error) {
	if m.GetByIDFn != nil {
		return m.GetByIDFn(id)
	}
	return &models.CustomerPolicy{}, nil
}

func (m *mockCustomerPolicyService) GetAll() ([]models.CustomerPolicy, error) {
	if m.GetAllFn != nil {
		return m.GetAllFn()
	}
	return nil, nil
}

func (m *mockCustomerPolicyService) Update(customerPolicy *models.CustomerPolicy) error {
	if m.UpdateFn != nil {
		return m.UpdateFn(customerPolicy)
	}
	return nil
}

func (m *mockCustomerPolicyService) Delete(id uint) error {
	if m.DeleteFn != nil {
		return m.DeleteFn(id)
	}
	return nil
}

func (m *mockCustomerPolicyService) GetByCustomerID(customerID uint) ([]models.CustomerPolicy, error) {
	if m.GetByCustomerIDFn != nil {
		return m.GetByCustomerIDFn(customerID)
	}
	return nil, nil
}

func (m *mockCustomerPolicyService) GetByPolicyID(policyID uint) ([]models.CustomerPolicy, error) {
	if m.GetByPolicyIDFn != nil {
		return m.GetByPolicyIDFn(policyID)
	}
	return nil, nil
}

type mockClaimService struct {
	CreateFn                func(*models.Claim) error
	GetByIDFn               func(uint) (*models.Claim, error)
	GetAllFn                func() ([]models.Claim, error)
	UpdateFn                func(*models.Claim) error
	DeleteFn                func(uint) error
	GetByCustomerPolicyIDFn func(uint) ([]models.Claim, error)
}

func (m *mockClaimService) Create(claim *models.Claim) error {
	if m.CreateFn != nil {
		return m.CreateFn(claim)
	}
	return nil
}

func (m *mockClaimService) GetByID(id uint) (*models.Claim, error) {
	if m.GetByIDFn != nil {
		return m.GetByIDFn(id)
	}
	return &models.Claim{}, nil
}

func (m *mockClaimService) GetAll() ([]models.Claim, error) {
	if m.GetAllFn != nil {
		return m.GetAllFn()
	}
	return nil, nil
}

func (m *mockClaimService) Update(claim *models.Claim) error {
	if m.UpdateFn != nil {
		return m.UpdateFn(claim)
	}
	return nil
}

func (m *mockClaimService) Delete(id uint) error {
	if m.DeleteFn != nil {
		return m.DeleteFn(id)
	}
	return nil
}

func (m *mockClaimService) GetByCustomerPolicyID(customerPolicyID uint) ([]models.Claim, error) {
	if m.GetByCustomerPolicyIDFn != nil {
		return m.GetByCustomerPolicyIDFn(customerPolicyID)
	}
	return nil, nil
}

func TestCustomerHandler(t *testing.T) {
	createdAt := time.Date(2026, time.July, 27, 0, 0, 0, 0, time.UTC)
	updatedAt := createdAt.Add(time.Hour)

	t.Run("Create", func(t *testing.T) {
		handler := NewCustomerHandler(&mockCustomerService{
			CreateFn: func(customer *models.Customer) error {
				customer.ID = 11
				customer.CreatedAt = createdAt
				customer.UpdatedAt = updatedAt
				return nil
			},
		})

		req := httptest.NewRequest(http.MethodPost, "/customers", strings.NewReader(`{"first_name":"Raj","last_name":"Kumar","email":"raj@example.com","phone_number":"123","address":"Pune"}`))
		rr := httptest.NewRecorder()

		handler.Create(rr, req)

		if rr.Code != http.StatusCreated {
			t.Fatalf("expected %d, got %d", http.StatusCreated, rr.Code)
		}

		var resp dto.CustomerResponse
		if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
			t.Fatalf("decode response: %v", err)
		}
		if resp.ID != 11 || resp.FirstName != "Raj" || resp.Email != "raj@example.com" {
			t.Fatalf("unexpected response: %+v", resp)
		}
	})

	t.Run("Create invalid body", func(t *testing.T) {
		handler := NewCustomerHandler(&mockCustomerService{})
		req := httptest.NewRequest(http.MethodPost, "/customers", strings.NewReader(`{`))
		rr := httptest.NewRecorder()

		handler.Create(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Fatalf("expected %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("GetAll", func(t *testing.T) {
		handler := NewCustomerHandler(&mockCustomerService{
			GetAllFn: func() ([]models.Customer, error) {
				return []models.Customer{{ID: 1, FirstName: "Jane", LastName: "Doe", Email: "jane@example.com", Phone: "555", Address: "Delhi", CreatedAt: createdAt, UpdatedAt: updatedAt}}, nil
			},
		})
		req := httptest.NewRequest(http.MethodGet, "/customers", nil)
		rr := httptest.NewRecorder()

		handler.GetAll(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected %d, got %d", http.StatusOK, rr.Code)
		}

		var resp []dto.CustomerResponse
		if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
			t.Fatalf("decode response: %v", err)
		}
		if len(resp) != 1 || resp[0].FirstName != "Jane" || resp[0].PhoneNumber != "555" {
			t.Fatalf("unexpected response: %+v", resp)
		}
	})

	t.Run("GetByID", func(t *testing.T) {
		handler := NewCustomerHandler(&mockCustomerService{
			GetByIDFn: func(id uint) (*models.Customer, error) {
				if id != 7 {
					t.Fatalf("expected id 7, got %d", id)
				}
				return &models.Customer{ID: 7, FirstName: "Jane", LastName: "Doe", Email: "jane@example.com", Phone: "555", Address: "Delhi", CreatedAt: createdAt, UpdatedAt: updatedAt}, nil
			},
		})
		req := httptest.NewRequest(http.MethodGet, "/customers/7", nil)
		req.SetPathValue("id", "7")
		rr := httptest.NewRecorder()

		handler.GetByID(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected %d, got %d", http.StatusOK, rr.Code)
		}
	})

	t.Run("Update", func(t *testing.T) {
		handler := NewCustomerHandler(&mockCustomerService{
			UpdateFn: func(customer *models.Customer) error {
				if customer.ID != 7 || customer.FirstName != "Jane" {
					t.Fatalf("unexpected customer: %+v", customer)
				}
				customer.UpdatedAt = updatedAt
				return nil
			},
		})
		req := httptest.NewRequest(http.MethodPut, "/customers/7", strings.NewReader(`{"first_name":"Jane","last_name":"Doe","email":"jane@example.com","phone_number":"555","address":"Delhi"}`))
		req.SetPathValue("id", "7")
		rr := httptest.NewRecorder()

		handler.Update(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected %d, got %d", http.StatusOK, rr.Code)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		handler := NewCustomerHandler(&mockCustomerService{
			DeleteFn: func(id uint) error {
				if id != 7 {
					t.Fatalf("expected id 7, got %d", id)
				}
				return nil
			},
		})
		req := httptest.NewRequest(http.MethodDelete, "/customers/7", nil)
		req.SetPathValue("id", "7")
		rr := httptest.NewRecorder()

		handler.Delete(rr, req)

		if rr.Code != http.StatusNoContent {
			t.Fatalf("expected %d, got %d", http.StatusNoContent, rr.Code)
		}
	})

	t.Run("GetByID invalid id", func(t *testing.T) {
		handler := NewCustomerHandler(&mockCustomerService{})
		req := httptest.NewRequest(http.MethodGet, "/customers/abc", nil)
		req.SetPathValue("id", "abc")
		rr := httptest.NewRecorder()

		handler.GetByID(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Fatalf("expected %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("Delete service error", func(t *testing.T) {
		handler := NewCustomerHandler(&mockCustomerService{
			DeleteFn: func(uint) error { return errors.New("delete failed") },
		})
		req := httptest.NewRequest(http.MethodDelete, "/customers/7", nil)
		req.SetPathValue("id", "7")
		rr := httptest.NewRecorder()

		handler.Delete(rr, req)

		if rr.Code != http.StatusInternalServerError {
			t.Fatalf("expected %d, got %d", http.StatusInternalServerError, rr.Code)
		}
	})
}

func TestPolicyHandler(t *testing.T) {
	createdAt := time.Date(2026, time.July, 27, 0, 0, 0, 0, time.UTC)
	updatedAt := createdAt.Add(time.Hour)

	handler := NewPolicyHandler(&mockPolicyService{
		CreateFn: func(policy *models.Policy) error {
			policy.ID = 21
			policy.CreatedAt = createdAt
			policy.UpdatedAt = updatedAt
			return nil
		},
		GetByIDFn: func(id uint) (*models.Policy, error) {
			return &models.Policy{ID: id, PolicyName: "Life", Description: "desc", Coverage: 100, Premium: 10, CreatedAt: createdAt, UpdatedAt: updatedAt}, nil
		},
		GetAllFn: func() ([]models.Policy, error) {
			return []models.Policy{{ID: 1, PolicyName: "Life", Description: "desc", Coverage: 100, Premium: 10, CreatedAt: createdAt, UpdatedAt: updatedAt}}, nil
		},
		UpdateFn: func(policy *models.Policy) error {
			if policy.ID != 21 {
				t.Fatalf("unexpected policy: %+v", policy)
			}
			policy.UpdatedAt = updatedAt
			return nil
		},
		DeleteFn: func(id uint) error {
			if id != 21 {
				t.Fatalf("expected id 21, got %d", id)
			}
			return nil
		},
	})

	createReq := httptest.NewRequest(http.MethodPost, "/policies", strings.NewReader(`{"policy_name":"Life","policy_type":"Term","coverage":100,"premium":10,"description":"desc"}`))
	createRR := httptest.NewRecorder()
	handler.Create(createRR, createReq)
	if createRR.Code != http.StatusCreated {
		t.Fatalf("expected %d, got %d", http.StatusCreated, createRR.Code)
	}

	getAllRR := httptest.NewRecorder()
	handler.GetAll(getAllRR, httptest.NewRequest(http.MethodGet, "/policies", nil))
	if getAllRR.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, getAllRR.Code)
	}

	getReq := httptest.NewRequest(http.MethodGet, "/policies/21", nil)
	getReq.SetPathValue("id", "21")
	getRR := httptest.NewRecorder()
	handler.GetByID(getRR, getReq)
	if getRR.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, getRR.Code)
	}

	updateReq := httptest.NewRequest(http.MethodPut, "/policies/21", strings.NewReader(`{"policy_name":"Life","policy_type":"Term","coverage":200,"premium":20,"description":"updated"}`))
	updateReq.SetPathValue("id", "21")
	updateRR := httptest.NewRecorder()
	handler.Update(updateRR, updateReq)
	if updateRR.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, updateRR.Code)
	}

	deleteReq := httptest.NewRequest(http.MethodDelete, "/policies/21", nil)
	deleteReq.SetPathValue("id", "21")
	deleteRR := httptest.NewRecorder()
	handler.Delete(deleteRR, deleteReq)
	if deleteRR.Code != http.StatusNoContent {
		t.Fatalf("expected %d, got %d", http.StatusNoContent, deleteRR.Code)
	}

	invalidReq := httptest.NewRequest(http.MethodGet, "/policies/abc", nil)
	invalidReq.SetPathValue("id", "abc")
	invalidRR := httptest.NewRecorder()
	handler.GetByID(invalidRR, invalidReq)
	if invalidRR.Code != http.StatusBadRequest {
		t.Fatalf("expected %d, got %d", http.StatusBadRequest, invalidRR.Code)
	}
}

func TestCustomerPolicyHandler(t *testing.T) {
	createdAt := time.Date(2026, time.July, 27, 0, 0, 0, 0, time.UTC)
	updatedAt := createdAt.Add(time.Hour)
	startDate := createdAt
	endDate := createdAt.AddDate(1, 0, 0)

	handler := NewCustomerPolicyHandler(&mockCustomerPolicyService{
		CreateFn: func(customerPolicy *models.CustomerPolicy) error {
			customerPolicy.ID = 31
			customerPolicy.CreatedAt = createdAt
			customerPolicy.UpdatedAt = updatedAt
			return nil
		},
		GetByIDFn: func(id uint) (*models.CustomerPolicy, error) {
			return &models.CustomerPolicy{ID: id, CustomerID: 1, PolicyID: 2, StartDate: startDate, EndDate: endDate, Status: "ACTIVE", CreatedAt: createdAt, UpdatedAt: updatedAt}, nil
		},
		GetAllFn: func() ([]models.CustomerPolicy, error) {
			return []models.CustomerPolicy{{ID: 1, CustomerID: 1, PolicyID: 2, StartDate: startDate, EndDate: endDate, Status: "ACTIVE", CreatedAt: createdAt, UpdatedAt: updatedAt}}, nil
		},
		UpdateFn: func(customerPolicy *models.CustomerPolicy) error {
			if customerPolicy.ID != 31 {
				t.Fatalf("unexpected customer policy: %+v", customerPolicy)
			}
			customerPolicy.UpdatedAt = updatedAt
			return nil
		},
		DeleteFn: func(id uint) error {
			if id != 31 {
				t.Fatalf("expected id 31, got %d", id)
			}
			return nil
		},
		GetByCustomerIDFn: func(customerID uint) ([]models.CustomerPolicy, error) {
			return []models.CustomerPolicy{{ID: 4, CustomerID: customerID, PolicyID: 2, StartDate: startDate, EndDate: endDate, Status: "ACTIVE"}}, nil
		},
		GetByPolicyIDFn: func(policyID uint) ([]models.CustomerPolicy, error) {
			return []models.CustomerPolicy{{ID: 5, CustomerID: 1, PolicyID: policyID, StartDate: startDate, EndDate: endDate, Status: "ACTIVE"}}, nil
		},
	})

	createReq := httptest.NewRequest(http.MethodPost, "/customer-policies", strings.NewReader(`{"customer_id":1,"policy_id":2,"start_date":"2026-07-27T00:00:00Z","end_date":"2027-07-27T00:00:00Z","status":"ACTIVE"}`))
	createRR := httptest.NewRecorder()
	handler.Create(createRR, createReq)
	if createRR.Code != http.StatusCreated {
		t.Fatalf("expected %d, got %d", http.StatusCreated, createRR.Code)
	}

	getAllRR := httptest.NewRecorder()
	handler.GetAll(getAllRR, httptest.NewRequest(http.MethodGet, "/customer-policies", nil))
	if getAllRR.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, getAllRR.Code)
	}

	getReq := httptest.NewRequest(http.MethodGet, "/customer-policies/31", nil)
	getReq.SetPathValue("id", "31")
	getRR := httptest.NewRecorder()
	handler.GetByID(getRR, getReq)
	if getRR.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, getRR.Code)
	}

	updateReq := httptest.NewRequest(http.MethodPut, "/customer-policies/31", strings.NewReader(`{"start_date":"2026-07-27T00:00:00Z","end_date":"2027-07-27T00:00:00Z","status":"INACTIVE"}`))
	updateReq.SetPathValue("id", "31")
	updateRR := httptest.NewRecorder()
	handler.Update(updateRR, updateReq)
	if updateRR.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, updateRR.Code)
	}

	deleteReq := httptest.NewRequest(http.MethodDelete, "/customer-policies/31", nil)
	deleteReq.SetPathValue("id", "31")
	deleteRR := httptest.NewRecorder()
	handler.Delete(deleteRR, deleteReq)
	if deleteRR.Code != http.StatusNoContent {
		t.Fatalf("expected %d, got %d", http.StatusNoContent, deleteRR.Code)
	}
}

func TestClaimHandler(t *testing.T) {
	createdAt := time.Date(2026, time.July, 27, 0, 0, 0, 0, time.UTC)
	updatedAt := createdAt.Add(time.Hour)
	claimDate := createdAt

	handler := NewClaimHandler(&mockClaimService{
		CreateFn: func(claim *models.Claim) error {
			claim.ID = 41
			claim.CreatedAt = createdAt
			claim.UpdatedAt = updatedAt
			return nil
		},
		GetByIDFn: func(id uint) (*models.Claim, error) {
			return &models.Claim{ID: id, CustomerPolicyID: 2, ClaimAmount: 1000, Reason: "Accident", Status: "PENDING", ClaimDate: claimDate, CreatedAt: createdAt, UpdatedAt: updatedAt}, nil
		},
		GetAllFn: func() ([]models.Claim, error) {
			return []models.Claim{{ID: 1, CustomerPolicyID: 2, ClaimAmount: 1000, Reason: "Accident", Status: "PENDING", ClaimDate: claimDate, CreatedAt: createdAt, UpdatedAt: updatedAt}}, nil
		},
		UpdateFn: func(claim *models.Claim) error {
			if claim.ID != 41 {
				t.Fatalf("unexpected claim: %+v", claim)
			}
			claim.UpdatedAt = updatedAt
			return nil
		},
		DeleteFn: func(id uint) error {
			if id != 41 {
				t.Fatalf("expected id 41, got %d", id)
			}
			return nil
		},
		GetByCustomerPolicyIDFn: func(customerPolicyID uint) ([]models.Claim, error) {
			return []models.Claim{{ID: 2, CustomerPolicyID: customerPolicyID, ClaimAmount: 500, Reason: "Fire", Status: "PENDING", ClaimDate: claimDate}}, nil
		},
	})

	createReq := httptest.NewRequest(http.MethodPost, "/claims", strings.NewReader(`{"customer_policy_id":2,"claim_amount":1000,"reason":"Accident","status":"PENDING","claim_date":"2026-07-27T00:00:00Z"}`))
	createReq.Header.Set("Content-Type", "application/json")
	createRR := httptest.NewRecorder()
	handler.Create(createRR, createReq)
	if createRR.Code != http.StatusCreated {
		t.Fatalf("expected %d, got %d", http.StatusCreated, createRR.Code)
	}

	getAllRR := httptest.NewRecorder()
	handler.GetAll(getAllRR, httptest.NewRequest(http.MethodGet, "/claims", nil))
	if getAllRR.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, getAllRR.Code)
	}

	getReq := httptest.NewRequest(http.MethodGet, "/claims/41", nil)
	getReq.SetPathValue("id", "41")
	getRR := httptest.NewRecorder()
	handler.GetByID(getRR, getReq)
	if getRR.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, getRR.Code)
	}

	updateReq := httptest.NewRequest(http.MethodPut, "/claims/41", strings.NewReader(`{"claim_amount":2500,"reason":"Updated","status":"APPROVED"}`))
	updateReq.Header.Set("Content-Type", "application/json")
	updateReq.SetPathValue("id", "41")
	updateRR := httptest.NewRecorder()
	handler.Update(updateRR, updateReq)
	if updateRR.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, updateRR.Code)
	}

	deleteReq := httptest.NewRequest(http.MethodDelete, "/claims/41", nil)
	deleteReq.SetPathValue("id", "41")
	deleteRR := httptest.NewRecorder()
	handler.Delete(deleteRR, deleteReq)
	if deleteRR.Code != http.StatusNoContent {
		t.Fatalf("expected %d, got %d", http.StatusNoContent, deleteRR.Code)
	}

	byCustomerPolicy, err := handler.service.GetByCustomerPolicyID(2)
	if err != nil || len(byCustomerPolicy) != 1 {
		t.Fatalf("unexpected claim result: %v %+v", err, byCustomerPolicy)
	}
}
