package service

import "insurance-api/internal/models"

type MockCustomerRepository struct {
	CreateFn  func(*models.Customer) error
	GetByIDFn func(uint) (*models.Customer, error)
	GetAllFn  func() ([]models.Customer, error)
	UpdateFn  func(*models.Customer) error
	DeleteFn  func(uint) error
}

func (m *MockCustomerRepository) Create(c *models.Customer) error {
	return m.CreateFn(c)
}

func (m *MockCustomerRepository) GetByID(id uint) (*models.Customer, error) {
	return m.GetByIDFn(id)
}

func (m *MockCustomerRepository) GetAll() ([]models.Customer, error) {
	return m.GetAllFn()
}

func (m *MockCustomerRepository) Update(c *models.Customer) error {
	return m.UpdateFn(c)
}

func (m *MockCustomerRepository) Delete(id uint) error {
	return m.DeleteFn(id)
}

type MockPolicyRepository struct {
	CreateFn  func(*models.Policy) error
	GetByIDFn func(uint) (*models.Policy, error)
	GetAllFn  func() ([]models.Policy, error)
	UpdateFn  func(*models.Policy) error
	DeleteFn  func(uint) error
}

func (m *MockPolicyRepository) Create(policy *models.Policy) error {
	return m.CreateFn(policy)
}

func (m *MockPolicyRepository) GetByID(id uint) (*models.Policy, error) {
	return m.GetByIDFn(id)
}

func (m *MockPolicyRepository) GetAll() ([]models.Policy, error) {
	return m.GetAllFn()
}

func (m *MockPolicyRepository) Update(policy *models.Policy) error {
	return m.UpdateFn(policy)
}

func (m *MockPolicyRepository) Delete(id uint) error {
	return m.DeleteFn(id)
}

type MockCustomerPolicyRepository struct {
	CreateFn          func(*models.CustomerPolicy) error
	GetByIDFn         func(uint) (*models.CustomerPolicy, error)
	GetAllFn          func() ([]models.CustomerPolicy, error)
	UpdateFn          func(*models.CustomerPolicy) error
	DeleteFn          func(uint) error
	GetByCustomerIDFn func(uint) ([]models.CustomerPolicy, error)
	GetByPolicyIDFn   func(uint) ([]models.CustomerPolicy, error)
}

func (m *MockCustomerPolicyRepository) Create(customerPolicy *models.CustomerPolicy) error {
	return m.CreateFn(customerPolicy)
}

func (m *MockCustomerPolicyRepository) GetByID(id uint) (*models.CustomerPolicy, error) {
	return m.GetByIDFn(id)
}

func (m *MockCustomerPolicyRepository) GetAll() ([]models.CustomerPolicy, error) {
	return m.GetAllFn()
}

func (m *MockCustomerPolicyRepository) Update(customerPolicy *models.CustomerPolicy) error {
	return m.UpdateFn(customerPolicy)
}

func (m *MockCustomerPolicyRepository) Delete(id uint) error {
	return m.DeleteFn(id)
}

func (m *MockCustomerPolicyRepository) GetByCustomerID(customerID uint) ([]models.CustomerPolicy, error) {
	return m.GetByCustomerIDFn(customerID)
}

func (m *MockCustomerPolicyRepository) GetByPolicyID(policyID uint) ([]models.CustomerPolicy, error) {
	return m.GetByPolicyIDFn(policyID)
}

type MockClaimRepository struct {
	CreateFn                func(*models.Claim) error
	GetByIDFn               func(uint) (*models.Claim, error)
	GetAllFn                func() ([]models.Claim, error)
	UpdateFn                func(*models.Claim) error
	DeleteFn                func(uint) error
	GetByCustomerPolicyIDFn func(uint) ([]models.Claim, error)
}

func (m *MockClaimRepository) Create(claim *models.Claim) error {
	return m.CreateFn(claim)
}

func (m *MockClaimRepository) GetByID(id uint) (*models.Claim, error) {
	return m.GetByIDFn(id)
}

func (m *MockClaimRepository) GetAll() ([]models.Claim, error) {
	return m.GetAllFn()
}

func (m *MockClaimRepository) Update(claim *models.Claim) error {
	return m.UpdateFn(claim)
}

func (m *MockClaimRepository) Delete(id uint) error {
	return m.DeleteFn(id)
}

func (m *MockClaimRepository) GetByCustomerPolicyID(customerPolicyID uint) ([]models.Claim, error) {
	return m.GetByCustomerPolicyIDFn(customerPolicyID)
}
