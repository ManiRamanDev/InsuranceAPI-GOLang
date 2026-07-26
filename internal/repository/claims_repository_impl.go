package repository

import (
	"insurance-api/internal/models"

	"gorm.io/gorm"
)

type claimRepository struct {
	db *gorm.DB
}

func NewClaimRepository(db *gorm.DB) ClaimRepository {
	return &claimRepository{
		db: db,
	}
}

func (r *claimRepository) Create(claim *models.Claim) error {
	return r.db.Create(claim).Error
}

func (r *claimRepository) GetByID(id uint) (*models.Claim, error) {
	var claim models.Claim

	err := r.db.
		Preload("CustomerPolicy").
		First(&claim, id).Error

	if err != nil {
		return nil, err
	}

	return &claim, nil
}

func (r *claimRepository) GetAll() ([]models.Claim, error) {
	var claims []models.Claim

	err := r.db.
		Preload("CustomerPolicy").
		Find(&claims).Error

	if err != nil {
		return nil, err
	}

	return claims, nil
}

func (r *claimRepository) Update(claim *models.Claim) error {
	return r.db.Save(claim).Error
}

func (r *claimRepository) Delete(id uint) error {
	return r.db.Delete(&models.Claim{}, id).Error
}

func (r *claimRepository) GetByCustomerPolicyID(customerPolicyID uint) ([]models.Claim, error) {
	var claims []models.Claim

	err := r.db.
		Where("customer_policy_id = ?", customerPolicyID).
		Find(&claims).Error

	if err != nil {
		return nil, err
	}

	return claims, nil
}
