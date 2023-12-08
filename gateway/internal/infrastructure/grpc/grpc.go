package grpc

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

func ConnectClient(
	ctx context.Context,
	cfg Config,
) (*grpc.ClientConn, error) {
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("validate config: %w", err)
	}

	var conn *grpc.ClientConn
	var err error

	if cfg.UseTLS {
		// TODO: add TLS configuration
	} else {
		conn, err = grpc.DialContext(
			ctx,
			cfg.Host,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithKeepaliveParams(keepalive.ClientParameters{
				Time:    cfg.KeepAliveTime,
				Timeout: cfg.KeepAliveTimeout,
			}),
		)
	}

	if err != nil {
		return nil, fmt.Errorf("dial grpc server: %w", err)
	}

	return conn, nil
}
