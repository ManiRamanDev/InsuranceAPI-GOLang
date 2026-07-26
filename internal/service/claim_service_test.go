package service

import (
	"reflect"
	"testing"
	"time"

	"insurance-api/internal/models"
)

func TestClaimService_AllMethods(t *testing.T) {
	service := NewClaimService(&MockClaimRepository{
		CreateFn: func(claim *models.Claim) error {
			claim.ID = 41
			return nil
		},
		GetByIDFn: func(id uint) (*models.Claim, error) {
			return &models.Claim{ID: id, CustomerPolicyID: 2, Reason: "Accident"}, nil
		},
		GetAllFn: func() ([]models.Claim, error) {
			return []models.Claim{{ID: 1, CustomerPolicyID: 2, Reason: "Accident"}}, nil
		},
		UpdateFn: func(claim *models.Claim) error {
			if claim.ID != 41 {
				t.Fatalf("expected id 41, got %d", claim.ID)
			}
			return nil
		},
		DeleteFn: func(id uint) error {
			if id != 8 {
				t.Fatalf("expected id 8, got %d", id)
			}
			return nil
		},
		GetByCustomerPolicyIDFn: func(customerPolicyID uint) ([]models.Claim, error) {
			return []models.Claim{{ID: 3, CustomerPolicyID: customerPolicyID, Reason: "Fire"}}, nil
		},
	})

	claimDate := time.Date(2026, time.July, 27, 0, 0, 0, 0, time.UTC)
	input := &models.Claim{CustomerPolicyID: 2, ClaimAmount: 1000, Reason: "Accident", ClaimDate: claimDate}
	if err := service.Create(input); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if input.ID != 41 {
		t.Fatalf("expected ID 41, got %d", input.ID)
	}

	byID, err := service.GetByID(5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if byID.ID != 5 || byID.CustomerPolicyID != 2 {
		t.Fatalf("unexpected claim: %+v", byID)
	}

	all, err := service.GetAll()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	wantAll := []models.Claim{{ID: 1, CustomerPolicyID: 2, Reason: "Accident"}}
	if !reflect.DeepEqual(all, wantAll) {
		t.Fatalf("expected %+v, got %+v", wantAll, all)
	}

	if err := service.Update(&models.Claim{ID: 41, CustomerPolicyID: 2, Reason: "Updated"}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if err := service.Delete(8); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	byCustomerPolicy, err := service.GetByCustomerPolicyID(12)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(byCustomerPolicy) != 1 || byCustomerPolicy[0].CustomerPolicyID != 12 {
		t.Fatalf("unexpected claims: %+v", byCustomerPolicy)
	}
}