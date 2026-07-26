package database

import (
	"testing"

	"insurance-api/internal/config"
)

func TestConnectRejectsIncompleteConfig(t *testing.T) {
	_, err := Connect(config.DatabaseConfig{})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestInitSchemaRejectsNilDB(t *testing.T) {
	if err := InitSchema(nil); err == nil {
		t.Fatal("expected error")
	}
}

func TestCloseNilDB(t *testing.T) {
	if err := Close(nil); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
}
