package dto

import "time"

type PolicyResponse struct {
	ID          uint      `json:"id"`
	PolicyName  string    `json:"policy_name"`
	PolicyType  string    `json:"policy_type"`
	Coverage    float64   `json:"coverage"`
	Premium     float64   `json:"premium"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
