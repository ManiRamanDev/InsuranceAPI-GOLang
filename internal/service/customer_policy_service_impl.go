package service

import (
	"insurance-api/internal/models"
	"insurance-api/internal/repository"
)

type customerPolicyService struct {
	repo repository.CustomerPolicyRepository
}

func NewCustomerPolicyService(repo repository.CustomerPolicyRepository) CustomerPolicyService {
	return &customerPolicyService{
		repo: repo,
	}
}

func (s *customerPolicyService) Create(customerPolicy *models.CustomerPolicy) error {
	return s.repo.Create(customerPolicy)
}

func (s *customerPolicyService) GetByID(id uint) (*models.CustomerPolicy, error) {
	return s.repo.GetByID(id)
}

func (s *customerPolicyService) GetAll() ([]models.CustomerPolicy, error) {
	return s.repo.GetAll()
}

func (s *customerPolicyService) Update(customerPolicy *models.CustomerPolicy) error {
	return s.repo.Update(customerPolicy)
}

func (s *customerPolicyService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *customerPolicyService) GetByCustomerID(customerID uint) ([]models.CustomerPolicy, error) {
	return s.repo.GetByCustomerID(customerID)
}

func (s *customerPolicyService) GetByPolicyID(policyID uint) ([]models.CustomerPolicy, error) {
	return s.repo.GetByPolicyID(policyID)
}
