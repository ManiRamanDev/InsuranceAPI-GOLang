package service

import (
	"insurance-api/internal/models"
	"insurance-api/internal/repository"
)

type claimService struct {
	repo repository.ClaimRepository
}

func NewClaimService(repo repository.ClaimRepository) ClaimService {
	return &claimService{
		repo: repo,
	}
}

func (s *claimService) Create(claim *models.Claim) error {
	return s.repo.Create(claim)
}

func (s *claimService) GetByID(id uint) (*models.Claim, error) {
	return s.repo.GetByID(id)
}

func (s *claimService) GetAll() ([]models.Claim, error) {
	return s.repo.GetAll()
}

func (s *claimService) Update(claim *models.Claim) error {
	return s.repo.Update(claim)
}

func (s *claimService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *claimService) GetByCustomerPolicyID(customerPolicyID uint) ([]models.Claim, error) {
	return s.repo.GetByCustomerPolicyID(customerPolicyID)
}
