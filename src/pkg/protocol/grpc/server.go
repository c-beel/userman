package grpc

import (
	"net"
	"os"
	"os/signal"
	"context"
	"log"
	"google.golang.org/grpc"
	"github.com/c-beel/userman/src/pkg/api/v1"
	"fmt"
)

func RunServer(ctx context.Context, v1API v1.UsermanServiceServer, port string) error {
	fmt.Println(":" + port)
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// register service
	server := grpc.NewServer()
	v1.RegisterUsermanServiceServer(server, v1API)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	log.Println("starting gRPC server...")
	return server.Serve(listen)
}
