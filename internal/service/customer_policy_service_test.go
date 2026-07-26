package service

import (
	"reflect"
	"testing"
	"time"

	"insurance-api/internal/models"
)

func TestCustomerPolicyService_AllMethods(t *testing.T) {
	service := NewCustomerPolicyService(&MockCustomerPolicyRepository{
		CreateFn: func(customerPolicy *models.CustomerPolicy) error {
			customerPolicy.ID = 31
			return nil
		},
		GetByIDFn: func(id uint) (*models.CustomerPolicy, error) {
			return &models.CustomerPolicy{ID: id, CustomerID: 1, PolicyID: 2}, nil
		},
		GetAllFn: func() ([]models.CustomerPolicy, error) {
			return []models.CustomerPolicy{{ID: 1, CustomerID: 1, PolicyID: 2}}, nil
		},
		UpdateFn: func(customerPolicy *models.CustomerPolicy) error {
			if customerPolicy.ID != 31 {
				t.Fatalf("expected id 31, got %d", customerPolicy.ID)
			}
			return nil
		},
		DeleteFn: func(id uint) error {
			if id != 9 {
				t.Fatalf("expected id 9, got %d", id)
			}
			return nil
		},
		GetByCustomerIDFn: func(customerID uint) ([]models.CustomerPolicy, error) {
			return []models.CustomerPolicy{{ID: 2, CustomerID: customerID, PolicyID: 4}}, nil
		},
		GetByPolicyIDFn: func(policyID uint) ([]models.CustomerPolicy, error) {
			return []models.CustomerPolicy{{ID: 3, CustomerID: 5, PolicyID: policyID}}, nil
		},
	})

	start := time.Date(2026, time.July, 27, 0, 0, 0, 0, time.UTC)
	input := &models.CustomerPolicy{CustomerID: 1, PolicyID: 2, StartDate: start, EndDate: start.AddDate(1, 0, 0)}
	if err := service.Create(input); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if input.ID != 31 {
		t.Fatalf("expected ID 31, got %d", input.ID)
	}

	byID, err := service.GetByID(7)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if byID.ID != 7 || byID.CustomerID != 1 || byID.PolicyID != 2 {
		t.Fatalf("unexpected customer policy: %+v", byID)
	}

	all, err := service.GetAll()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	wantAll := []models.CustomerPolicy{{ID: 1, CustomerID: 1, PolicyID: 2}}
	if !reflect.DeepEqual(all, wantAll) {
		t.Fatalf("expected %+v, got %+v", wantAll, all)
	}

	if err := service.Update(&models.CustomerPolicy{ID: 31, CustomerID: 1, PolicyID: 2}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if err := service.Delete(9); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	byCustomer, err := service.GetByCustomerID(10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(byCustomer) != 1 || byCustomer[0].CustomerID != 10 {
		t.Fatalf("unexpected customer policies: %+v", byCustomer)
	}

	byPolicy, err := service.GetByPolicyID(11)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(byPolicy) != 1 || byPolicy[0].PolicyID != 11 {
		t.Fatalf("unexpected customer policies: %+v", byPolicy)
	}
}
