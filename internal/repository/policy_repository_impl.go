package repository

import (
	"insurance-api/internal/models"

	"gorm.io/gorm"
)

type policyRepository struct {
	db *gorm.DB
}

func NewPolicyRepository(db *gorm.DB) PolicyRepository {
	return &policyRepository{
		db: db,
	}
}

func (r *policyRepository) Create(policy *models.Policy) error {
	return r.db.Create(policy).Error
}

func (r *policyRepository) GetByID(id uint) (*models.Policy, error) {
	var policy models.Policy

	err := r.db.First(&policy, id).Error
	if err != nil {
		return nil, err
	}

	return &policy, nil
}

func (r *policyRepository) GetAll() ([]models.Policy, error) {
	var policies []models.Policy

	err := r.db.Find(&policies).Error
	if err != nil {
		return nil, err
	}

	return policies, nil
}

func (r *policyRepository) Update(policy *models.Policy) error {
	return r.db.Save(policy).Error
}

func (r *policyRepository) Delete(id uint) error {
	return r.db.Delete(&models.Policy{}, id).Error
}
