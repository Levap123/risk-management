package grpc

import (
	"errors"
	"time"
)

type Config struct {
	Host             string
	UseTLS           bool
	KeepAliveTime    time.Duration
	KeepAliveTimeout time.Duration
}

func (c Config) Validate() error {
	if c.Host == "" {
		return errors.New("host is empty")
	}

	if c.KeepAliveTime <= 0 {
		return errors.New("keep alive time is less or equal zero")
	}

	if c.KeepAliveTimeout <= 0 {
		return errors.New("keep alive timeout is less or equal zero")
	}

	return nil
}
