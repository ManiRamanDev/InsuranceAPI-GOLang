package config

import (
	"encoding/base64"
	"testing"
)

func TestLoadAPIConfig(t *testing.T) {
	tests := []struct {
		name    string
		env     map[string]string
		want    APIConfig
		wantErr bool
	}{
		{
			name: "defaults",
			env: map[string]string{
				"INSURANCE_HTTP_ADDRESS":           "",
				"INSURANCE_HTTPS_ENABLED":          "",
				"INSURANCE_MAX_REQUEST_BODY_BYTES": "",
			},
			want: APIConfig{
				HTTPAddress:         DefaultHTTPAddress,
				HTTPSEnabled:        false,
				MaxRequestBodyBytes: DefaultMaxRequestBodyBytes,
				HTTPPort:            DefaultPort,
			},
		},
		{
			name: "https enabled requires tls files",
			env: map[string]string{
				"INSURANCE_HTTPS_ENABLED": "true",
				"INSURANCE_TLS_CERT_FILE": "",
				"INSURANCE_TLS_KEY_FILE":  "",
			},
			wantErr: true,
		},
		{
			name: "rejects invalid boolean",
			env: map[string]string{
				"INSURANCE_HTTPS_ENABLED": "maybe",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for key, value := range tt.env {
				t.Setenv(key, value)
			}

			cfg, err := LoadAPIConfig()
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if cfg != tt.want {
				t.Fatalf("expected %+v, got %+v", tt.want, cfg)
			}
		})
	}
}

func TestLoadDatabaseConfig(t *testing.T) {
	tests := []struct {
		name    string
		env     map[string]string
		want    DatabaseConfig
		wantErr bool
	}{
		{
			name: "success",
			env: map[string]string{
				"INSURANCE_DB_HOST":   "localhost",
				"INSURANCE_DB_PORT":   "5432",
				"INSURANCE_DB_NAME":   "insurance",
				"INSURANCE_DB_USER":   "postgres",
				"INSURANCE_DB_SECRET": base64.StdEncoding.EncodeToString([]byte("password")),
			},
			want: DatabaseConfig{
				Host:     "localhost",
				Port:     "5432",
				Name:     "insurance",
				User:     "postgres",
				Password: "password",
			},
		},
		{
			name: "invalid secret",
			env: map[string]string{
				"INSURANCE_DB_HOST":   "localhost",
				"INSURANCE_DB_PORT":   "5432",
				"INSURANCE_DB_NAME":   "insurance",
				"INSURANCE_DB_USER":   "postgres",
				"INSURANCE_DB_SECRET": "not-base64",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for key, value := range tt.env {
				t.Setenv(key, value)
			}

			cfg, err := LoadDatabaseConfig()
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if cfg != tt.want {
				t.Fatalf("expected %+v, got %+v", tt.want, cfg)
			}
		})
	}
}

func TestLoadLogConfig(t *testing.T) {
	tests := []struct {
		name    string
		env     map[string]string
		want    LogConfig
		wantErr bool
	}{
		{
			name: "defaults",
			env: map[string]string{
				"INSURANCE_LOG_DIRECTORY": "",
				"INSURANCE_LOG_FILE":      "",
				"INSURANCE_LOG_LEVEL":     "",
			},
			want: LogConfig{Directory: DefaultLogDirectory, FileName: DefaultLogFileName, Level: DefaultLogLevel},
		},
		{
			name: "rejects invalid file path",
			env: map[string]string{
				"INSURANCE_LOG_FILE": "logs/app.log",
			},
			wantErr: true,
		},
		{
			name: "rejects invalid level",
			env: map[string]string{
				"INSURANCE_LOG_FILE":  DefaultLogFileName,
				"INSURANCE_LOG_LEVEL": "TRACE",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for key, value := range tt.env {
				t.Setenv(key, value)
			}

			cfg, err := LoadLogConfig()
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if cfg != tt.want {
				t.Fatalf("expected %+v, got %+v", tt.want, cfg)
			}
		})
	}
}
