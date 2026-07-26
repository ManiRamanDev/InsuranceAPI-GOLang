package dto

import (
	"encoding/json"
	"testing"
	"time"
)

func BenchmarkRequestDTOMarshal(b *testing.B) {
	createdAt := time.Date(2026, time.July, 27, 0, 0, 0, 0, time.UTC)
	cases := map[string]any{
		"customer": CreateCustomerRequest{FirstName: "Raj", LastName: "Kumar", Email: "raj@example.com", PhoneNumber: "123", Address: "Pune", DateOfBirth: createdAt},
		"policy": CreatePolicyRequest{PolicyName: "Life", PolicyType: "Term", Coverage: 100, Premium: 10, Description: "desc"},
		"customer_policy": CreateCustomerPolicyRequest{CustomerID: 1, PolicyID: 2, StartDate: createdAt, EndDate: createdAt.AddDate(1, 0, 0), Status: "ACTIVE"},
		"claim": CreateClaimRequest{CustomerPolicyID: 3, ClaimAmount: 1000, Reason: "Accident", Status: "PENDING", ClaimDate: createdAt},
	}

	for name, payload := range cases {
		b.Run(name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				if _, err := json.Marshal(payload); err != nil {
					b.Fatalf("marshal failed: %v", err)
				}
			}
		})
	}
}
