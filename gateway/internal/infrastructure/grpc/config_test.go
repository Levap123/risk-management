package grpc

import (
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
				Host:             "localhost:8080",
				KeepAliveTime:    time.Minute * 100,
				KeepAliveTimeout: time.Second * 20,
			},
			noErr: true,
		},
		{
			name: "empty keep alive time",
			config: Config{
				Host:             "localhost:8080",
				KeepAliveTimeout: time.Second * 20,
			},
		},
		{
			name: "empty keep alive timeout",
			config: Config{
				Host:          "localhost:8080",
				KeepAliveTime: time.Minute * 100,
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
