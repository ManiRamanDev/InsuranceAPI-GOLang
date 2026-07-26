package service

import (
	"insurance-api/internal/models"
	"insurance-api/internal/repository"
)

type customerService struct {
	repo repository.CustomerRepository
}

func NewCustomerService(repo repository.CustomerRepository) CustomerService {
	return &customerService{
		repo: repo,
	}
}

func (s *customerService) Create(customer *models.Customer) error {
	return s.repo.Create(customer)
}

func (s *customerService) GetByID(id uint) (*models.Customer, error) {
	return s.repo.GetByID(id)
}

func (s *customerService) GetAll() ([]models.Customer, error) {
	return s.repo.GetAll()
}

func (s *customerService) Update(customer *models.Customer) error {
	return s.repo.Update(customer)
}

func (s *customerService) Delete(id uint) error {
	return s.repo.Delete(id)
}
