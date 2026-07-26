package repository

import "insurance-api/internal/models"

type ClaimRepository interface {
	Create(claim *models.Claim) error
	GetByID(id uint) (*models.Claim, error)
	GetAll() ([]models.Claim, error)
	Update(claim *models.Claim) error
	Delete(id uint) error

	GetByCustomerPolicyID(customerPolicyID uint) ([]models.Claim, error)
}
