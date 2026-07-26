package models

import "time"

type Customer struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	FirstName string `gorm:"size:100;not null"`
	LastName  string `gorm:"size:100;not null"`
	Email     string `gorm:"size:255;unique;not null"`
	Phone     string `gorm:"size:15;not null"`
	Address   string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Policies []CustomerPolicy `gorm:"foreignKey:CustomerID"`
}
