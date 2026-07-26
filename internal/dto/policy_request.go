package dto

type CreatePolicyRequest struct {
	PolicyName  string  `json:"policy_name"`
	PolicyType  string  `json:"policy_type"`
	Coverage    float64 `json:"coverage"`
	Premium     float64 `json:"premium"`
	Description string  `json:"description"`
}

type UpdatePolicyRequest struct {
	PolicyName  string  `json:"policy_name"`
	PolicyType  string  `json:"policy_type"`
	Coverage    float64 `json:"coverage"`
	Premium     float64 `json:"premium"`
	Description string  `json:"description"`
}
