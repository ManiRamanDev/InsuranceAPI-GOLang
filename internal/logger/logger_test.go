package logger

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"insurance-api/internal/config"
)

func TestNewAndClose(t *testing.T) {
	dir := t.TempDir()
	logConfig := config.LogConfig{Directory: dir, FileName: "app.log", Level: config.DefaultLogLevel}

	l, err := New(logConfig)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	l.Debug("debug message")
	l.Info("info message")
	l.Warn("warn message")
	l.Error("error message", errors.New("boom"))
	l.Audit("create_customer", "customer_id", 1)

	if err := l.Close(); err != nil {
		t.Fatalf("unexpected close error: %v", err)
	}

	if _, err := os.Stat(filepath.Join(dir, "app.log")); err != nil {
		t.Fatalf("expected log file to exist: %v", err)
	}
}

func TestNewRejectsUnsupportedLevel(t *testing.T) {
	_, err := New(config.LogConfig{Directory: t.TempDir(), FileName: "app.log", Level: "TRACE"})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestCloseOnZeroValueLogger(t *testing.T) {
	var l Logger
	if err := l.Close(); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
}
