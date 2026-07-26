package service

import (
	"testing"
	"time"

	"insurance-api/internal/models"
)

func BenchmarkCustomerService(b *testing.B) {
	svc := NewCustomerService(&MockCustomerRepository{
		CreateFn: func(c *models.Customer) error {
			c.ID = 1
			return nil
		},
		GetByIDFn: func(id uint) (*models.Customer, error) {
			return &models.Customer{ID: id, FirstName: "Raj"}, nil
		},
		GetAllFn: func() ([]models.Customer, error) {
			return []models.Customer{{ID: 1, FirstName: "Raj"}}, nil
		},
		UpdateFn: func(*models.Customer) error { return nil },
		DeleteFn: func(uint) error { return nil },
	})

	b.ReportAllocs()
	b.Run("Create", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			customer := &models.Customer{FirstName: "Raj", LastName: "Kumar"}
			if err := svc.Create(customer); err != nil {
				b.Fatalf("create failed: %v", err)
			}
		}
	})

	b.Run("GetByID", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if _, err := svc.GetByID(uint(i + 1)); err != nil {
				b.Fatalf("get by id failed: %v", err)
			}
		}
	})

	b.Run("GetAll", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if _, err := svc.GetAll(); err != nil {
				b.Fatalf("get all failed: %v", err)
			}
		}
	})

	b.Run("Update", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if err := svc.Update(&models.Customer{ID: 1, FirstName: "Updated"}); err != nil {
				b.Fatalf("update failed: %v", err)
			}
		}
	})

	b.Run("Delete", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if err := svc.Delete(1); err != nil {
				b.Fatalf("delete failed: %v", err)
			}
		}
	})
}

func BenchmarkPolicyService(b *testing.B) {
	svc := NewPolicyService(&MockPolicyRepository{
		CreateFn: func(p *models.Policy) error {
			p.ID = 1
			return nil
		},
		GetByIDFn: func(id uint) (*models.Policy, error) {
			return &models.Policy{ID: id, PolicyName: "Life"}, nil
		},
		GetAllFn: func() ([]models.Policy, error) {
			return []models.Policy{{ID: 1, PolicyName: "Life"}}, nil
		},
		UpdateFn: func(*models.Policy) error { return nil },
		DeleteFn: func(uint) error { return nil },
	})

	b.ReportAllocs()
	b.Run("Create", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			policy := &models.Policy{PolicyName: "Life", Coverage: 100, Premium: 10}
			if err := svc.Create(policy); err != nil {
				b.Fatalf("create failed: %v", err)
			}
		}
	})

	b.Run("GetByID", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if _, err := svc.GetByID(uint(i + 1)); err != nil {
				b.Fatalf("get by id failed: %v", err)
			}
		}
	})

	b.Run("GetAll", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if _, err := svc.GetAll(); err != nil {
				b.Fatalf("get all failed: %v", err)
			}
		}
	})

	b.Run("Update", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if err := svc.Update(&models.Policy{ID: 1, PolicyName: "Updated", Coverage: 200, Premium: 20}); err != nil {
				b.Fatalf("update failed: %v", err)
			}
		}
	})

	b.Run("Delete", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if err := svc.Delete(1); err != nil {
				b.Fatalf("delete failed: %v", err)
			}
		}
	})
}

func BenchmarkCustomerPolicyService(b *testing.B) {
	start := time.Date(2026, time.July, 27, 0, 0, 0, 0, time.UTC)
	svc := NewCustomerPolicyService(&MockCustomerPolicyRepository{
		CreateFn: func(cp *models.CustomerPolicy) error {
			cp.ID = 1
			return nil
		},
		GetByIDFn: func(id uint) (*models.CustomerPolicy, error) {
			return &models.CustomerPolicy{ID: id, CustomerID: 1, PolicyID: 2}, nil
		},
		GetAllFn: func() ([]models.CustomerPolicy, error) {
			return []models.CustomerPolicy{{ID: 1, CustomerID: 1, PolicyID: 2}}, nil
		},
		UpdateFn: func(*models.CustomerPolicy) error { return nil },
		DeleteFn: func(uint) error { return nil },
		GetByCustomerIDFn: func(id uint) ([]models.CustomerPolicy, error) {
			return []models.CustomerPolicy{{ID: 1, CustomerID: id, PolicyID: 2}}, nil
		},
		GetByPolicyIDFn: func(id uint) ([]models.CustomerPolicy, error) {
			return []models.CustomerPolicy{{ID: 1, CustomerID: 2, PolicyID: id}}, nil
		},
	})

	b.ReportAllocs()
	b.Run("Create", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			cp := &models.CustomerPolicy{CustomerID: 1, PolicyID: 2, StartDate: start, EndDate: start.AddDate(1, 0, 0), Status: "ACTIVE"}
			if err := svc.Create(cp); err != nil {
				b.Fatalf("create failed: %v", err)
			}
		}
	})

	b.Run("GetByID", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if _, err := svc.GetByID(uint(i + 1)); err != nil {
				b.Fatalf("get by id failed: %v", err)
			}
		}
	})

	b.Run("GetAll", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if _, err := svc.GetAll(); err != nil {
				b.Fatalf("get all failed: %v", err)
			}
		}
	})

	b.Run("Update", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if err := svc.Update(&models.CustomerPolicy{ID: 1, StartDate: start, EndDate: start.AddDate(1, 0, 0), Status: "INACTIVE"}); err != nil {
				b.Fatalf("update failed: %v", err)
			}
		}
	})

	b.Run("Delete", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if err := svc.Delete(1); err != nil {
				b.Fatalf("delete failed: %v", err)
			}
		}
	})

	b.Run("GetByCustomerID", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if _, err := svc.GetByCustomerID(1); err != nil {
				b.Fatalf("get by customer failed: %v", err)
			}
		}
	})

	b.Run("GetByPolicyID", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if _, err := svc.GetByPolicyID(2); err != nil {
				b.Fatalf("get by policy failed: %v", err)
			}
		}
	})
}

func BenchmarkClaimService(b *testing.B) {
	claimDate := time.Date(2026, time.July, 27, 0, 0, 0, 0, time.UTC)
	svc := NewClaimService(&MockClaimRepository{
		CreateFn: func(c *models.Claim) error {
			c.ID = 1
			return nil
		},
		GetByIDFn: func(id uint) (*models.Claim, error) {
			return &models.Claim{ID: id, CustomerPolicyID: 2, Reason: "Accident"}, nil
		},
		GetAllFn: func() ([]models.Claim, error) {
			return []models.Claim{{ID: 1, CustomerPolicyID: 2, Reason: "Accident"}}, nil
		},
		UpdateFn: func(*models.Claim) error { return nil },
		DeleteFn: func(uint) error { return nil },
		GetByCustomerPolicyIDFn: func(id uint) ([]models.Claim, error) {
			return []models.Claim{{ID: 1, CustomerPolicyID: id, Reason: "Fire"}}, nil
		},
	})

	b.ReportAllocs()
	b.Run("Create", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			claim := &models.Claim{CustomerPolicyID: 2, ClaimAmount: 1000, Reason: "Accident", ClaimDate: claimDate}
			if err := svc.Create(claim); err != nil {
				b.Fatalf("create failed: %v", err)
			}
		}
	})

	b.Run("GetByID", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if _, err := svc.GetByID(uint(i + 1)); err != nil {
				b.Fatalf("get by id failed: %v", err)
			}
		}
	})

	b.Run("GetAll", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if _, err := svc.GetAll(); err != nil {
				b.Fatalf("get all failed: %v", err)
			}
		}
	})

	b.Run("Update", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if err := svc.Update(&models.Claim{ID: 1, ClaimAmount: 2000, Reason: "Updated", Status: "APPROVED"}); err != nil {
				b.Fatalf("update failed: %v", err)
			}
		}
	})

	b.Run("Delete", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if err := svc.Delete(1); err != nil {
				b.Fatalf("delete failed: %v", err)
			}
		}
	})

	b.Run("GetByCustomerPolicyID", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if _, err := svc.GetByCustomerPolicyID(2); err != nil {
				b.Fatalf("get by customer policy failed: %v", err)
			}
		}
	})
}
