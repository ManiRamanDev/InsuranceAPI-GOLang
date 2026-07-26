package dto

import "time"

type CreateCustomerRequest struct {
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	DateOfBirth time.Time `json:"date_of_birth"`
}

type UpdateCustomerRequest struct {
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	DateOfBirth time.Time `json:"date_of_birth"`
}
