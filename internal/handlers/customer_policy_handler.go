package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"insurance-api/internal/dto"
	"insurance-api/internal/models"
	"insurance-api/internal/service"
)

type CustomerPolicyHandler struct {
	service service.CustomerPolicyService
}

func NewCustomerPolicyHandler(service service.CustomerPolicyService) *CustomerPolicyHandler {
	return &CustomerPolicyHandler{
		service: service,
	}
}

// Create Customer Policy
func (h *CustomerPolicyHandler) Create(w http.ResponseWriter, r *http.Request) {

	var req dto.CreateCustomerPolicyRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	customerPolicy := models.CustomerPolicy{
		CustomerID: req.CustomerID,
		PolicyID:   req.PolicyID,
		StartDate:  req.StartDate,
		EndDate:    req.EndDate,
		Status:     req.Status,
	}

	if err := h.service.Create(&customerPolicy); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := dto.CustomerPolicyResponse{
		ID:         customerPolicy.ID,
		CustomerID: customerPolicy.CustomerID,
		PolicyID:   customerPolicy.PolicyID,
		StartDate:  customerPolicy.StartDate,
		EndDate:    customerPolicy.EndDate,
		Status:     customerPolicy.Status,
		CreatedAt:  customerPolicy.CreatedAt,
		UpdatedAt:  customerPolicy.UpdatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// Get All Customer Policies
func (h *CustomerPolicyHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	customerPolicies, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := make([]dto.CustomerPolicyResponse, 0)

	for _, cp := range customerPolicies {
		response = append(response, dto.CustomerPolicyResponse{
			ID:         cp.ID,
			CustomerID: cp.CustomerID,
			PolicyID:   cp.PolicyID,
			StartDate:  cp.StartDate,
			EndDate:    cp.EndDate,
			Status:     cp.Status,
			CreatedAt:  cp.CreatedAt,
			UpdatedAt:  cp.UpdatedAt,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Get Customer Policy By ID
func (h *CustomerPolicyHandler) GetByID(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid customer policy ID", http.StatusBadRequest)
		return
	}

	customerPolicy, err := h.service.GetByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	response := dto.CustomerPolicyResponse{
		ID:         customerPolicy.ID,
		CustomerID: customerPolicy.CustomerID,
		PolicyID:   customerPolicy.PolicyID,
		StartDate:  customerPolicy.StartDate,
		EndDate:    customerPolicy.EndDate,
		Status:     customerPolicy.Status,
		CreatedAt:  customerPolicy.CreatedAt,
		UpdatedAt:  customerPolicy.UpdatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Update Customer Policy
func (h *CustomerPolicyHandler) Update(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid customer policy ID", http.StatusBadRequest)
		return
	}

	var req dto.UpdateCustomerPolicyRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	customerPolicy := models.CustomerPolicy{
		ID:        uint(id),
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		Status:    req.Status,
	}

	if err := h.service.Update(&customerPolicy); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := dto.CustomerPolicyResponse{
		ID:         customerPolicy.ID,
		CustomerID: customerPolicy.CustomerID,
		PolicyID:   customerPolicy.PolicyID,
		StartDate:  customerPolicy.StartDate,
		EndDate:    customerPolicy.EndDate,
		Status:     customerPolicy.Status,
		CreatedAt:  customerPolicy.CreatedAt,
		UpdatedAt:  customerPolicy.UpdatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Delete Customer Policy
func (h *CustomerPolicyHandler) Delete(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid customer policy ID", http.StatusBadRequest)
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
