package dto

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestRequestDTOJSONTags(t *testing.T) {
	createdAt := time.Date(2026, time.July, 27, 0, 0, 0, 0, time.UTC)

	customer := CreateCustomerRequest{FirstName: "Raj", LastName: "Kumar", Email: "raj@example.com", PhoneNumber: "123", Address: "Pune", DateOfBirth: createdAt}
	policy := CreatePolicyRequest{PolicyName: "Life", PolicyType: "Term", Coverage: 100, Premium: 10, Description: "desc"}
	customerPolicy := CreateCustomerPolicyRequest{CustomerID: 1, PolicyID: 2, StartDate: createdAt, EndDate: createdAt.AddDate(1, 0, 0), Status: "ACTIVE"}
	claim := CreateClaimRequest{CustomerPolicyID: 3, ClaimAmount: 1000, Reason: "Accident", Status: "PENDING", ClaimDate: createdAt}

	for _, value := range []any{customer, policy, customerPolicy, claim} {
		payload, err := json.Marshal(value)
		if err != nil {
			t.Fatalf("marshal: %v", err)
		}
		if len(payload) == 0 {
			t.Fatal("expected json payload")
		}
	}
}

func TestResponseDTOShapes(t *testing.T) {
	if reflect.TypeOf(CustomerResponse{}).Field(0).Tag.Get("json") != "id" {
		t.Fatal("unexpected customer response tag")
	}
	if reflect.TypeOf(PolicyResponse{}).Field(0).Tag.Get("json") != "id" {
		t.Fatal("unexpected policy response tag")
	}
	if reflect.TypeOf(CustomerPolicyResponse{}).Field(0).Tag.Get("json") != "id" {
		t.Fatal("unexpected customer policy response tag")
	}
	if reflect.TypeOf(ClaimResponse{}).Field(0).Tag.Get("json") != "id" {
		t.Fatal("unexpected claim response tag")
	}
}
