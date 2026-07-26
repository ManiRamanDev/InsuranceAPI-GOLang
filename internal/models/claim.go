package models

import "time"

type Claim struct {
	ID               uint      `gorm:"primaryKey;autoIncrement"`
	CustomerPolicyID uint      `gorm:"not null"`
	ClaimAmount      float64   `gorm:"not null"`
	Reason           string    `gorm:"size:500;not null"`
	Status           string    `gorm:"size:30;default:'PENDING'"`
	ClaimDate        time.Time `gorm:"not null"`
	CreatedAt        time.Time
	UpdatedAt        time.Time

	CustomerPolicy CustomerPolicy `gorm:"foreignKey:CustomerPolicyID"`
}
