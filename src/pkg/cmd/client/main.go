package main

import (
	"flag"
	"log"
	"time"
	"context"
	"google.golang.org/grpc"
	"github.com/c-beel/userman/src/pkg/api/v1"
)

func main() {
	// get configuration
	address := flag.String("server", "localhost:8000", "gRPC server in format host:port")
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := v1.NewUsermanServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	req2 := v1.ReadUserRequest{
		Uid: 1,
	}

	req := v1.CreateUserRequest{
		User: &v1.User{
			Username:  "example",
			Nickname:  "Example",
			FirstName: "exam",
			LastName:  "ple",
			Email:     "example@google.com",
		},
	}

	res2, err := c.ReadUser(ctx, &req2)
	if err != nil {
		log.Println("Get failed: %v", err)
	}
	log.Printf("Get result: <%+v>\n\n", res2)

	res, err := c.CreateUser(ctx, &req)
	if err != nil {
		log.Println("Create failed: %v", err)
	}
	log.Printf("Create result: <%+v>\n\n", res)

	res2, err = c.ReadUser(ctx, &req2)
	if err != nil {
		log.Println("Get failed: %v", err)
	}
	log.Printf("Get result: <%+v>\n\n", res2)
}
