package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	// DefaultHTTPAddress is the local address used when no address is configured.
	DefaultHTTPAddress = "localhost"
	DefaultPort        = "8080"
	SecuredPort        = "8443"
	// DefaultMaxRequestBodyBytes limits JSON request bodies to 500 KB.
	DefaultMaxRequestBodyBytes int64 = 500 * 1024

	// API endpoint paths.
	CustomersPath = "/customers"
	PoliciesPath  = "/policies"
	ClaimsPath    = "/claims"

	// CustomerPolicyPath documents the assignment endpoint shape.
	CustomerPolicyPath = "/customers/{customerID}/policies/{policyID}"
)

// APIConfig contains HTTP, HTTPS, and request-limit configuration.
type APIConfig struct {
	HTTPAddress         string
	HTTPSEnabled        bool
	TLSCertificateFile  string
	TLSPrivateKeyFile   string
	MaxRequestBodyBytes int64
	HTTPPort            string
}

// LoadAPIConfig reads API configuration from environment variables.
func LoadAPIConfig() (APIConfig, error) {
	config := APIConfig{
		HTTPAddress:         optionalEnv("INSURANCE_HTTP_ADDRESS", DefaultHTTPAddress),
		MaxRequestBodyBytes: DefaultMaxRequestBodyBytes,
		HTTPPort:            DefaultPort,
	}

	httpsEnabled, err := optionalBoolEnv("INSURANCE_HTTPS_ENABLED", false)
	if err != nil {
		return APIConfig{}, err
	}
	config.HTTPSEnabled = httpsEnabled

	maxRequestBodyBytes, err := optionalPositiveInt64Env(
		"INSURANCE_MAX_REQUEST_BODY_BYTES",
		DefaultMaxRequestBodyBytes,
	)
	if err != nil {
		return APIConfig{}, err
	}
	config.MaxRequestBodyBytes = maxRequestBodyBytes

	if !config.HTTPSEnabled {
		return config, nil
	}

	certificateFile, err := requiredEnv("INSURANCE_TLS_CERT_FILE")
	if err != nil {
		return APIConfig{}, err
	}

	privateKeyFile, err := requiredEnv("INSURANCE_TLS_KEY_FILE")
	if err != nil {
		return APIConfig{}, err
	}

	config.TLSCertificateFile = certificateFile
	config.TLSPrivateKeyFile = privateKeyFile

	return config, nil
}

func optionalEnv(name, defaultValue string) string {
	value := strings.TrimSpace(os.Getenv(name))
	if value == "" {
		return defaultValue
	}

	return value
}

func optionalBoolEnv(name string, defaultValue bool) (bool, error) {
	value := strings.TrimSpace(os.Getenv(name))
	if value == "" {
		return defaultValue, nil
	}

	parsedValue, err := strconv.ParseBool(value)
	if err != nil {
		return false, fmt.Errorf("environment variable %s must be true or false", name)
	}

	return parsedValue, nil
}

func optionalPositiveInt64Env(name string, defaultValue int64) (int64, error) {
	value := strings.TrimSpace(os.Getenv(name))
	if value == "" {
		return defaultValue, nil
	}

	parsedValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil || parsedValue <= 0 {
		return 0, fmt.Errorf("environment variable %s must be a positive number", name)
	}

	return parsedValue, nil
}
