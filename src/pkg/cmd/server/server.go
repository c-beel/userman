package cmd

import (
	"context"
	"github.com/c-beel/userman/src/pkg/service/v1"
	"github.com/c-beel/userman/src/pkg/protocol/grpc"
	"github.com/c-beel/userman/src/configman"
	"flag"
	"log"
)

// RunServer runs gRPC server and HTTP gateway
func RunServer() error {
	ctx := context.Background()

	configFileAddress := flag.String("conf", "config.yaml", "The path to the service config file")
	autoMigrate := flag.Bool("migrate", true, "Auto-migrate models")
	flag.Parse()

	// get configuration
	cfg, err := configman.ImportConfigFromFile(*configFileAddress)
	if err != nil {
		log.Fatalf("Failed to parse config file with error %v", err)
	}

	v1API, err := v1.NewUsermanServer(cfg)
	if err != nil {
		log.Fatalf("failed to start service : %v", err)
	}
	if *autoMigrate {
		log.Println("Starting auto migrate...")
		if err := v1API.AutoMigrate(); err != nil {
			log.Fatalf("failed to auto migrate : %v", err)
		}
		log.Println("Auto migrate done.")
	}
	return grpc.RunServer(ctx, *v1API, cfg.ListenPort)
}
