package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"insurance-api/internal/dto"
	"insurance-api/internal/models"
	"insurance-api/internal/service"
)

type PolicyHandler struct {
	service service.PolicyService
}

func NewPolicyHandler(service service.PolicyService) *PolicyHandler {
	return &PolicyHandler{
		service: service,
	}
}

// Create Policy
func (h *PolicyHandler) Create(w http.ResponseWriter, r *http.Request) {

	var req dto.CreatePolicyRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	policy := models.Policy{
		PolicyName:  req.PolicyName,
		Description: req.Description,
		Coverage:    req.Coverage,
		Premium:     req.Premium,
	}

	if err := h.service.Create(&policy); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := dto.PolicyResponse{
		ID:          policy.ID,
		PolicyName:  policy.PolicyName,
		Description: policy.Description,
		Coverage:    policy.Coverage,
		Premium:     policy.Premium,
		CreatedAt:   policy.CreatedAt,
		UpdatedAt:   policy.UpdatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// Get All Policies
func (h *PolicyHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	policies, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := make([]dto.PolicyResponse, 0)

	for _, policy := range policies {
		response = append(response, dto.PolicyResponse{
			ID:          policy.ID,
			PolicyName:  policy.PolicyName,
			Description: policy.Description,
			Coverage:    policy.Coverage,
			Premium:     policy.Premium,
			CreatedAt:   policy.CreatedAt,
			UpdatedAt:   policy.UpdatedAt,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Get Policy By ID
func (h *PolicyHandler) GetByID(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid policy ID", http.StatusBadRequest)
		return
	}

	policy, err := h.service.GetByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	response := dto.PolicyResponse{
		ID:          policy.ID,
		PolicyName:  policy.PolicyName,
		Description: policy.Description,
		Coverage:    policy.Coverage,
		Premium:     policy.Premium,
		CreatedAt:   policy.CreatedAt,
		UpdatedAt:   policy.UpdatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Update Policy
func (h *PolicyHandler) Update(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid policy ID", http.StatusBadRequest)
		return
	}

	var req dto.UpdatePolicyRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	policy := models.Policy{
		ID:          uint(id),
		PolicyName:  req.PolicyName,
		Description: req.Description,
		Coverage:    req.Coverage,
		Premium:     req.Premium,
	}

	if err := h.service.Update(&policy); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := dto.PolicyResponse{
		ID:          policy.ID,
		PolicyName:  policy.PolicyName,
		Description: policy.Description,
		Coverage:    policy.Coverage,
		Premium:     policy.Premium,
		CreatedAt:   policy.CreatedAt,
		UpdatedAt:   policy.UpdatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Delete Policy
func (h *PolicyHandler) Delete(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid policy ID", http.StatusBadRequest)
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
