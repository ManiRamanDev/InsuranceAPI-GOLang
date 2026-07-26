package dto

import "time"

type CreateCustomerPolicyRequest struct {
	CustomerID uint      `json:"customer_id"`
	PolicyID   uint      `json:"policy_id"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	Status     string    `json:"status"`
}

type UpdateCustomerPolicyRequest struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Status    string    `json:"status"`
}
