package config

import (
	"fmt"
	"path/filepath"
	"strings"
)

const (
	// DefaultLogDirectory is where application log files are written.
	DefaultLogDirectory = "logs"

	// DefaultLogFileName is the application log file name.
	DefaultLogFileName = "insurance-api.log"

	// DefaultLogLevel controls the minimum log level written by the logger.
	DefaultLogLevel = "INFO"
)

// LogConfig contains file-logging settings.
type LogConfig struct {
	Directory string
	FileName  string
	Level     string
}

// LoadLogConfig reads and validates logging configuration.
func LoadLogConfig() (LogConfig, error) {
	directory := optionalEnv("INSURANCE_LOG_DIRECTORY", DefaultLogDirectory)
	fileName := optionalEnv("INSURANCE_LOG_FILE", DefaultLogFileName)
	level := strings.ToUpper(optionalEnv("INSURANCE_LOG_LEVEL", DefaultLogLevel))

	if filepath.Base(fileName) != fileName {
		return LogConfig{}, fmt.Errorf("INSURANCE_LOG_FILE must contain only a file name")
	}

	switch level {
	case "DEBUG", "INFO", "WARN", "ERROR":
	default:
		return LogConfig{}, fmt.Errorf(
			"INSURANCE_LOG_LEVEL must be DEBUG, INFO, WARN, or ERROR",
		)
	}

	return LogConfig{
		Directory: directory,
		FileName:  fileName,
		Level:     level,
	}, nil
}
