package logger

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"insurance-api/internal/config"
)

// Logger writes structured application and audit logs.
type Logger struct {
	log  *slog.Logger
	file *os.File
}

// New creates a file logger from application configuration.
func New(logConfig config.LogConfig) (*Logger, error) {
	level, err := parseLevel(logConfig.Level)
	if err != nil {
		return nil, err
	}

	if err := os.MkdirAll(logConfig.Directory, 0o755); err != nil {
		return nil, fmt.Errorf("create log directory: %w", err)
	}

	logPath := filepath.Join(logConfig.Directory, logConfig.FileName)
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o640)
	if err != nil {
		return nil, fmt.Errorf("open log file: %w", err)
	}

	handler := slog.NewJSONHandler(file, &slog.HandlerOptions{
		Level: level,
	})

	return &Logger{
		log:  slog.New(handler),
		file: file,
	}, nil
}

// Debug records diagnostic information.
func (l *Logger) Debug(message string, fields ...any) {
	l.log.Debug(message, fields...)
}

// Info records normal application activity.
func (l *Logger) Info(message string, fields ...any) {
	l.log.Info(message, fields...)
}

// Warn records unexpected but recoverable activity.
func (l *Logger) Warn(message string, fields ...any) {
	l.log.Warn(message, fields...)
}

// Error records a failure that needs attention.
func (l *Logger) Error(message string, err error, fields ...any) {
	fields = append(fields, "error", err)
	l.log.Error(message, fields...)
}

// Audit records a business action for later review.
func (l *Logger) Audit(action string, fields ...any) {
	fields = append([]any{"event_type", "audit", "action", action}, fields...)
	l.log.Info("audit event", fields...)
}

// Close releases the log file handle.
func (l *Logger) Close() error {
	if l.file == nil {
		return nil
	}

	return l.file.Close()
}

func parseLevel(level string) (slog.Level, error) {
	switch level {
	case "DEBUG":
		return slog.LevelDebug, nil
	case "INFO":
		return slog.LevelInfo, nil
	case "WARN":
		return slog.LevelWarn, nil
	case "ERROR":
		return slog.LevelError, nil
	default:
		return 0, fmt.Errorf("unsupported log level %q", level)
	}
}
