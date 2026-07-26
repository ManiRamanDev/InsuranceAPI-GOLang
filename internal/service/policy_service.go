package service

import "insurance-api/internal/models"

type PolicyService interface {
	Create(policy *models.Policy) error
	GetByID(id uint) (*models.Policy, error)
	GetAll() ([]models.Policy, error)
	Update(policy *models.Policy) error
	Delete(id uint) error
}
