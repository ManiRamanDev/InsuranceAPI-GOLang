package database

import (
	"fmt"

	"insurance-api/internal/config"
	"insurance-api/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect establishes and verifies a PostgreSQL database connection.
func Connect(dbConfig config.DatabaseConfig) (*gorm.DB, error) {
	if dbConfig.Host == "" ||
		dbConfig.Port == "" ||
		dbConfig.Name == "" ||
		dbConfig.User == "" ||
		dbConfig.Password == "" {
		return nil, fmt.Errorf("database configuration is incomplete")
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Name,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("open database connection: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("get SQL database connection: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		_ = sqlDB.Close()
		return nil, fmt.Errorf("ping database: %w", err)
	}

	return db, nil
}

// InitSchema creates missing tables and applies compatible schema changes.
func InitSchema(db *gorm.DB) error {
	if db == nil {
		return fmt.Errorf("database instance is nil")
	}

	if err := db.AutoMigrate(
		&models.Customer{},
		&models.Policy{},
		&models.CustomerPolicy{},
		&models.Claim{},
	); err != nil {
		return fmt.Errorf("migrate database schema: %w", err)
	}

	return nil
}

// Close releases the PostgreSQL connection.
func Close(db *gorm.DB) error {
	if db == nil {
		return nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("get SQL database connection: %w", err)
	}

	if err := sqlDB.Close(); err != nil {
		return fmt.Errorf("close database connection: %w", err)
	}

	return nil
}
