package config

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

// DatabaseConfig contains the database values needed to establish a connection.
type DatabaseConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

// LoadDatabaseConfig reads, validates, and prepares database configuration.
func LoadDatabaseConfig() (DatabaseConfig, error) {
	host, err := requiredEnv("INSURANCE_DB_HOST")
	if err != nil {
		return DatabaseConfig{}, err
	}

	port, err := requiredEnv("INSURANCE_DB_PORT")
	if err != nil {
		return DatabaseConfig{}, err
	}

	name, err := requiredEnv("INSURANCE_DB_NAME")
	if err != nil {
		return DatabaseConfig{}, err
	}

	user, err := requiredEnv("INSURANCE_DB_USER")
	if err != nil {
		return DatabaseConfig{}, err
	}

	secret, err := requiredEnv("INSURANCE_DB_SECRET")
	if err != nil {
		return DatabaseConfig{}, err
	}

	passwordBytes, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return DatabaseConfig{}, fmt.Errorf("database secret is invalid")
	}

	if len(passwordBytes) == 0 {
		return DatabaseConfig{}, fmt.Errorf("database secret is invalid")
	}

	return DatabaseConfig{
		Host:     host,
		Port:     port,
		Name:     name,
		User:     user,
		Password: string(passwordBytes),
	}, nil
}

func requiredEnv(name string) (string, error) {
	value := strings.TrimSpace(os.Getenv(name))
	if value == "" {
		return "", fmt.Errorf("required environment variable %s is not set", name)
	}

	return value, nil
}
