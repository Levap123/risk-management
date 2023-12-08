package http

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Host              string
	Port              string
	ReadHeaderTimeout time.Duration
	handler           http.Handler
}

func (c Config) Validate() error {
	if c.Host == "" {
		return errors.New("host is empty")
	}

	if c.Port == "" {
		return errors.New("port is empty")
	}

	if c.ReadHeaderTimeout <= 0 {
		return errors.New("read header timeout is less or equal 0")
	}

	if c.handler == nil {
		return errors.New("handler is empty")
	}

	return nil
}

func (c Config) Address() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}
