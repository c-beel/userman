package main

import (
	"flag"
	"log"
	"time"
	"context"
	"google.golang.org/grpc"
	"github.com/c-beel/userman/src/pkg/api/v1"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

func main() {
	// get configuration
	address := flag.String("server", "localhost:8080", "gRPC server in format host:port")
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

	req2 := v1.GetUserByIdRequest{
		Api:     apiVersion,
		Id:      1,
		IdToken: "..-----",
	}

	req := v1.UpdateUserRequest{
		Api: apiVersion,
		User: &v1.User{
			Username:  "ex@mple",
			FirstName: "example",
			LastName:  "exam",
			Email:     "example@example.com",
		},
		IdToken: "..-----",
	}

	res2, err := c.GetUserById(ctx, &req2)
	if err != nil {
		log.Fatalf("Get failed: %v", err)
	}
	log.Printf("Get result: <%+v>\n\n", res2)

	res, err := c.UpdateUser(ctx, &req)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create result: <%+v>\n\n", res)

	res2, err = c.GetUserById(ctx, &req2)
	if err != nil {
		log.Fatalf("Get failed: %v", err)
	}
	log.Printf("Get result: <%+v>\n\n", res2)
}
