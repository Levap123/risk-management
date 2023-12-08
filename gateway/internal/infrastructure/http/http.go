package http

import (
	"fmt"
	"net/http"
	"time"
)

func New(cfg Config) (*http.Server, error) {
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("validate config: %w", err)
	}

	return &http.Server{
		Addr:              cfg.Address(),
		ReadHeaderTimeout: time.Minute,
		Handler:           cfg.handler,
	}, nil
}
