package repository

import (
	"insurance-api/internal/models"

	"gorm.io/gorm"
)

type customerPolicyRepository struct {
	db *gorm.DB
}

func NewCustomerPolicyRepository(db *gorm.DB) CustomerPolicyRepository {
	return &customerPolicyRepository{
		db: db,
	}
}

func (r *customerPolicyRepository) Create(customerPolicy *models.CustomerPolicy) error {
	return r.db.Create(customerPolicy).Error
}

func (r *customerPolicyRepository) GetByID(id uint) (*models.CustomerPolicy, error) {
	var customerPolicy models.CustomerPolicy

	err := r.db.
		Preload("Customer").
		Preload("Policy").
		First(&customerPolicy, id).Error

	if err != nil {
		return nil, err
	}

	return &customerPolicy, nil
}

func (r *customerPolicyRepository) GetAll() ([]models.CustomerPolicy, error) {
	var customerPolicies []models.CustomerPolicy

	err := r.db.
		Preload("Customer").
		Preload("Policy").
		Find(&customerPolicies).Error

	if err != nil {
		return nil, err
	}

	return customerPolicies, nil
}

func (r *customerPolicyRepository) Update(customerPolicy *models.CustomerPolicy) error {
	return r.db.Save(customerPolicy).Error
}

func (r *customerPolicyRepository) Delete(id uint) error {
	return r.db.Delete(&models.CustomerPolicy{}, id).Error
}

func (r *customerPolicyRepository) GetByCustomerID(customerID uint) ([]models.CustomerPolicy, error) {
	var customerPolicies []models.CustomerPolicy

	err := r.db.
		Preload("Policy").
		Where("customer_id = ?", customerID).
		Find(&customerPolicies).Error

	if err != nil {
		return nil, err
	}

	return customerPolicies, nil
}

func (r *customerPolicyRepository) GetByPolicyID(policyID uint) ([]models.CustomerPolicy, error) {
	var customerPolicies []models.CustomerPolicy

	err := r.db.
		Preload("Customer").
		Where("policy_id = ?", policyID).
		Find(&customerPolicies).Error

	if err != nil {
		return nil, err
	}

	return customerPolicies, nil
}
