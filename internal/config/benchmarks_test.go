package config

import (
	"encoding/base64"
	"testing"
)

func BenchmarkLoadAPIConfig(b *testing.B) {
	b.Setenv("INSURANCE_HTTP_ADDRESS", "localhost")
	b.Setenv("INSURANCE_HTTPS_ENABLED", "false")
	b.Setenv("INSURANCE_MAX_REQUEST_BODY_BYTES", "512000")

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := LoadAPIConfig(); err != nil {
			b.Fatalf("load api config failed: %v", err)
		}
	}
}

func BenchmarkLoadDatabaseConfig(b *testing.B) {
	secret := base64.StdEncoding.EncodeToString([]byte("vijay123"))
	b.Setenv("INSURANCE_DB_HOST", "localhost")
	b.Setenv("INSURANCE_DB_PORT", "5432")
	b.Setenv("INSURANCE_DB_NAME", "insurance_api")
	b.Setenv("INSURANCE_DB_USER", "postgres")
	b.Setenv("INSURANCE_DB_SECRET", secret)

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := LoadDatabaseConfig(); err != nil {
			b.Fatalf("load db config failed: %v", err)
		}
	}
}

func BenchmarkLoadLogConfig(b *testing.B) {
	b.Setenv("INSURANCE_LOG_DIRECTORY", "logs")
	b.Setenv("INSURANCE_LOG_FILE", "insurance-api.log")
	b.Setenv("INSURANCE_LOG_LEVEL", "INFO")

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := LoadLogConfig(); err != nil {
			b.Fatalf("load log config failed: %v", err)
		}
	}
}
