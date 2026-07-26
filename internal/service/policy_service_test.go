package service

import (
	"reflect"
	"testing"

	"insurance-api/internal/models"
)

func TestPolicyService_CreateAndGetByID(t *testing.T) {
	service := NewPolicyService(&MockPolicyRepository{
		CreateFn: func(policy *models.Policy) error {
			policy.ID = 21
			return nil
		},
		GetByIDFn: func(id uint) (*models.Policy, error) {
			return &models.Policy{ID: id, PolicyName: "Life"}, nil
		},
	})

	policy := &models.Policy{PolicyName: "Life"}
	if err := service.Create(policy); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if policy.ID != 21 {
		t.Fatalf("expected ID 21, got %d", policy.ID)
	}

	got, err := service.GetByID(5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.ID != 5 || got.PolicyName != "Life" {
		t.Fatalf("unexpected policy: %+v", got)
	}
}

func TestPolicyService_GetAllUpdateDelete(t *testing.T) {
	service := NewPolicyService(&MockPolicyRepository{
		GetAllFn: func() ([]models.Policy, error) {
			return []models.Policy{{ID: 1, PolicyName: "Life"}, {ID: 2, PolicyName: "Health"}}, nil
		},
		UpdateFn: func(policy *models.Policy) error {
			if policy.ID != 2 {
				t.Fatalf("expected id 2, got %d", policy.ID)
			}
			return nil
		},
		DeleteFn: func(id uint) error {
			if id != 3 {
				t.Fatalf("expected id 3, got %d", id)
			}
			return nil
		},
	})

	policies, err := service.GetAll()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []models.Policy{{ID: 1, PolicyName: "Life"}, {ID: 2, PolicyName: "Health"}}
	if !reflect.DeepEqual(policies, want) {
		t.Fatalf("expected %+v, got %+v", want, policies)
	}

	if err := service.Update(&models.Policy{ID: 2, PolicyName: "Updated"}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if err := service.Delete(3); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
