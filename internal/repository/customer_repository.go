package repository

import (
	"insurance-api/internal/models"
)

type CustomerRepository interface {
	Create(customer *models.Customer) error
	GetByID(id uint) (*models.Customer, error)
	GetAll() ([]models.Customer, error)
	Update(customer *models.Customer) error
	Delete(id uint) error
}
