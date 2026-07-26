package service

import (
	"insurance-api/internal/models"
	"insurance-api/internal/repository"
)

type policyService struct {
	repo repository.PolicyRepository
}

func NewPolicyService(repo repository.PolicyRepository) PolicyService {
	return &policyService{
		repo: repo,
	}
}

func (s *policyService) Create(policy *models.Policy) error {
	return s.repo.Create(policy)
}

func (s *policyService) GetByID(id uint) (*models.Policy, error) {
	return s.repo.GetByID(id)
}

func (s *policyService) GetAll() ([]models.Policy, error) {
	return s.repo.GetAll()
}

func (s *policyService) Update(policy *models.Policy) error {
	return s.repo.Update(policy)
}

func (s *policyService) Delete(id uint) error {
	return s.repo.Delete(id)
}
