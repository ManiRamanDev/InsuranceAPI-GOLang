package models

import "time"

type Policy struct {
	ID          uint    `gorm:"primaryKey;autoIncrement"`
	PolicyName  string  `gorm:"size:100;not null"`
	Description string  `gorm:"size:500"`
	Coverage    float64 `gorm:"not null"`
	Premium     float64 `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Customers []CustomerPolicy `gorm:"foreignKey:PolicyID"`
}
