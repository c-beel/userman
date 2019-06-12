package cmd

import (
	"context"
	"fmt"
	"github.com/c-beel/userman/src/pkg/service/v1"
	"github.com/c-beel/userman/src/pkg/protocol/grpc"
	"github.com/c-beel/userman/src/configman"
	"os"
)

// RunServer runs gRPC server and HTTP gateway
func RunServer() error {
	ctx := context.Background()

	// get configuration
	cfg := configman.Config{
		GRPCPort:          os.Getenv("GRPCPort"),
		GoogleOAuthAPIKey: os.Getenv("GoogleOAuthAPIKey"),
		DBAddress:         os.Getenv("DBAddress"),
	}

	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}

	v1API, err := v1.NewUsermanServer(cfg)
	if err != nil {
		return fmt.Errorf("failed to start service : %v", err)
	}
	if err := v1API.AutoMigrate(); err != nil {
		return fmt.Errorf("failed to auto migrate : %v", err)
	}

	return grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}
