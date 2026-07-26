package models

import "time"

type CustomerPolicy struct {
	ID         uint      `gorm:"primaryKey;autoIncrement"`
	CustomerID uint      `gorm:"not null"`
	PolicyID   uint      `gorm:"not null"`
	StartDate  time.Time `gorm:"not null"`
	EndDate    time.Time `gorm:"not null"`
	Status     string    `gorm:"size:30;default:'ACTIVE'"`
	CreatedAt  time.Time
	UpdatedAt  time.Time

	Customer Customer `gorm:"foreignKey:CustomerID"`
	Policy   Policy   `gorm:"foreignKey:PolicyID"`
	Claims   []Claim  `gorm:"foreignKey:CustomerPolicyID"`
}
