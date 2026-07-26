package service

import (
	"errors"
	"reflect"
	"testing"

	"insurance-api/internal/models"
)

func TestCustomerService_Create(t *testing.T) {
	tests := []struct {
		name    string
		repoErr error
		wantErr bool
	}{
		{name: "success"},
		{name: "repo error", repoErr: errors.New("create failed"), wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotCustomer *models.Customer
			service := NewCustomerService(&MockCustomerRepository{
				CreateFn: func(customer *models.Customer) error {
					gotCustomer = customer
					if tt.repoErr == nil {
						customer.ID = 11
					}
					return tt.repoErr
				},
			})

			customer := &models.Customer{FirstName: "Raj", LastName: "Kumar"}
			err := service.Create(customer)

			if tt.wantErr {
				if err == nil || err.Error() != tt.repoErr.Error() {
					t.Fatalf("expected error %v, got %v", tt.repoErr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if gotCustomer != customer {
				t.Fatalf("service passed a different customer pointer")
			}
			if customer.ID != 11 {
				t.Fatalf("expected ID 11, got %d", customer.ID)
			}
		})
	}
}

func TestCustomerService_GetByID(t *testing.T) {
	expected := &models.Customer{ID: 1, FirstName: "John", LastName: "Doe"}
	service := NewCustomerService(&MockCustomerRepository{
		GetByIDFn: func(id uint) (*models.Customer, error) {
			if id != 1 {
				t.Fatalf("expected id 1, got %d", id)
			}
			return expected, nil
		},
	})

	customer, err := service.GetByID(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(customer, expected) {
		t.Fatalf("expected %+v, got %+v", expected, customer)
	}
}

func TestCustomerService_GetAll(t *testing.T) {
	expected := []models.Customer{{ID: 1, FirstName: "John"}, {ID: 2, FirstName: "Jane"}}
	service := NewCustomerService(&MockCustomerRepository{
		GetAllFn: func() ([]models.Customer, error) {
			return expected, nil
		},
	})

	customers, err := service.GetAll()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(customers, expected) {
		t.Fatalf("expected %+v, got %+v", expected, customers)
	}
}

func TestCustomerService_Update(t *testing.T) {
	service := NewCustomerService(&MockCustomerRepository{
		UpdateFn: func(customer *models.Customer) error {
			if customer.ID != 10 {
				t.Fatalf("expected ID 10, got %d", customer.ID)
			}
			return nil
		},
	})

	err := service.Update(&models.Customer{ID: 10, FirstName: "Updated"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestCustomerService_Delete(t *testing.T) {
	service := NewCustomerService(&MockCustomerRepository{
		DeleteFn: func(id uint) error {
			if id != 9 {
				t.Fatalf("expected id 9, got %d", id)
			}
			return nil
		},
	})

	if err := service.Delete(9); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
