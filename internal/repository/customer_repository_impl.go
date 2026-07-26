package repository

import (
	models "insurance-api/internal/models"

	"gorm.io/gorm"
)

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{
		db: db,
	}
}

func (r *customerRepository) Create(customer *models.Customer) error {
	return r.db.Create(customer).Error
}

func (r *customerRepository) GetByID(id uint) (*models.Customer, error) {
	var customer models.Customer

	err := r.db.First(&customer, id).Error
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (r *customerRepository) GetAll() ([]models.Customer, error) {
	var customers []models.Customer

	err := r.db.Find(&customers).Error
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (r *customerRepository) Update(customer *models.Customer) error {
	return r.db.Save(customer).Error
}

func (r *customerRepository) Delete(id uint) error {
	return r.db.Delete(&models.Customer{}, id).Error
}
