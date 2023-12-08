package http

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestConfig_Validate(t *testing.T) {
	tests := []struct {
		name   string
		config Config
		noErr  bool
	}{
		{
			name: "success",
			config: Config{
				Host:              "localhost:",
				Port:              "7834",
				ReadHeaderTimeout: time.Hour,
				handler:           http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			},
			noErr: true,
		},
		{
			name: "empty host",
			config: Config{
				Port:              "7834",
				ReadHeaderTimeout: time.Hour,
				handler:           http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			},
		},
		{
			name: "empty port",
			config: Config{
				Host:              "localhost:",
				ReadHeaderTimeout: time.Hour,
				handler:           http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			},
		},
		{
			name: "empty read header timeout",
			config: Config{
				Host:    "localhost:",
				Port:    "7834",
				handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			},
		},
		{
			name: "empty handler",
			config: Config{
				Host:              "localhost:",
				Port:              "7834",
				ReadHeaderTimeout: time.Hour,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if tt.noErr {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
