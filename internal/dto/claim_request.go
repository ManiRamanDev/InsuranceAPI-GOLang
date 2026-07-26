package dto

import "time"

type CreateClaimRequest struct {
	CustomerPolicyID uint      `json:"customer_policy_id"`
	ClaimAmount      float64   `json:"claim_amount"`
	Reason           string    `json:"reason"`
	Status           string    `json:"status"`
	ClaimDate        time.Time `json:"claim_date"`
}

type UpdateClaimRequest struct {
	ClaimAmount float64 `json:"claim_amount"`
	Reason      string  `json:"reason"`
	Status      string  `json:"status"`
}
