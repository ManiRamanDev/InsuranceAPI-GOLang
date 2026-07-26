package dto

import "time"

type ClaimResponse struct {
	ID               uint      `json:"id"`
	CustomerPolicyID uint      `json:"customer_policy_id"`
	ClaimAmount      float64   `json:"claim_amount"`
	Reason           string    `json:"reason"`
	Status           string    `json:"status"`
	ClaimDate        time.Time `json:"claim_date"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
