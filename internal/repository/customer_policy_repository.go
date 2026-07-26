package repository

import "insurance-api/internal/models"

type CustomerPolicyRepository interface {
	Create(customerPolicy *models.CustomerPolicy) error
	GetByID(id uint) (*models.CustomerPolicy, error)
	GetAll() ([]models.CustomerPolicy, error)
	Update(customerPolicy *models.CustomerPolicy) error
	Delete(id uint) error

	GetByCustomerID(customerID uint) ([]models.CustomerPolicy, error)
	GetByPolicyID(policyID uint) ([]models.CustomerPolicy, error)
}
