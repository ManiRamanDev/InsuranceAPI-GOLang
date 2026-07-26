package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"insurance-api/internal/dto"
	"insurance-api/internal/models"
	"insurance-api/internal/service"
)

type ClaimHandler struct {
	service service.ClaimService
}

func NewClaimHandler(service service.ClaimService) *ClaimHandler {
	return &ClaimHandler{
		service: service,
	}
}

// Create Claim
func (h *ClaimHandler) Create(w http.ResponseWriter, r *http.Request) {

	var req dto.CreateClaimRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	claim := models.Claim{
		CustomerPolicyID: req.CustomerPolicyID,
		ClaimAmount:      req.ClaimAmount,
		Reason:           req.Reason,
		Status:           req.Status,
		ClaimDate:        req.ClaimDate,
	}

	if err := h.service.Create(&claim); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := dto.ClaimResponse{
		ID:               claim.ID,
		CustomerPolicyID: claim.CustomerPolicyID,
		ClaimAmount:      claim.ClaimAmount,
		Reason:           claim.Reason,
		Status:           claim.Status,
		ClaimDate:        claim.ClaimDate,
		CreatedAt:        claim.CreatedAt,
		UpdatedAt:        claim.UpdatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// Get All Claims
func (h *ClaimHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	claims, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := make([]dto.ClaimResponse, 0)

	for _, claim := range claims {
		response = append(response, dto.ClaimResponse{
			ID:               claim.ID,
			CustomerPolicyID: claim.CustomerPolicyID,
			ClaimAmount:      claim.ClaimAmount,
			Reason:           claim.Reason,
			Status:           claim.Status,
			ClaimDate:        claim.ClaimDate,
			CreatedAt:        claim.CreatedAt,
			UpdatedAt:        claim.UpdatedAt,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Get Claim By ID
func (h *ClaimHandler) GetByID(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid claim ID", http.StatusBadRequest)
		return
	}

	claim, err := h.service.GetByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	response := dto.ClaimResponse{
		ID:               claim.ID,
		CustomerPolicyID: claim.CustomerPolicyID,
		ClaimAmount:      claim.ClaimAmount,
		Reason:           claim.Reason,
		Status:           claim.Status,
		ClaimDate:        claim.ClaimDate,
		CreatedAt:        claim.CreatedAt,
		UpdatedAt:        claim.UpdatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Update Claim
func (h *ClaimHandler) Update(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid claim ID", http.StatusBadRequest)
		return
	}

	var req dto.UpdateClaimRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	claim := models.Claim{
		ID:          uint(id),
		ClaimAmount: req.ClaimAmount,
		Reason:      req.Reason,
		Status:      req.Status,
	}

	if err := h.service.Update(&claim); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := dto.ClaimResponse{
		ID:               claim.ID,
		CustomerPolicyID: claim.CustomerPolicyID,
		ClaimAmount:      claim.ClaimAmount,
		Reason:           claim.Reason,
		Status:           claim.Status,
		ClaimDate:        claim.ClaimDate,
		CreatedAt:        claim.CreatedAt,
		UpdatedAt:        claim.UpdatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Delete Claim
func (h *ClaimHandler) Delete(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid claim ID", http.StatusBadRequest)
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
