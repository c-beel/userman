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
		IdToken: "eyJhbGciOiJSUzI1NiIsImtpZCI6ImM3ZjUyMmQwMzIyODRkMjUyYmVlNGZkODA1NjBjZWZhMGZiNjBjMzkiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJhY2NvdW50cy5nb29nbGUuY29tIiwiYXpwIjoiODY2NzQ4MzQxMjAtdjcyMXJ2OXFxaWswNWI3dWNmbnFrYjRlZ25xZzlmdjguYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJhdWQiOiI4NjY3NDgzNDEyMC12NzIxcnY5cXFpazA1Yjd1Y2ZucWtiNGVnbnFnOWZ2OC5hcHBzLmdvb2dsZXVzZXJjb250ZW50LmNvbSIsInN1YiI6IjExMzM2NzQwNTUxMDQ0NzU1NzM3MyIsImVtYWlsIjoiYW1vb21hamlkOTlAZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImF0X2hhc2giOiJqRWRKbmNLeGlXUmY3d3owMEVPcVJnIiwibmFtZSI6Ik1hamlkIEdhcm9vc2kiLCJwaWN0dXJlIjoiaHR0cHM6Ly9saDUuZ29vZ2xldXNlcmNvbnRlbnQuY29tLy1Jc1FfZnp0QUp5cy9BQUFBQUFBQUFBSS9BQUFBQUFBQUJjWS9ueWZPb1lnUGItay9zOTYtYy9waG90by5qcGciLCJnaXZlbl9uYW1lIjoiTWFqaWQiLCJmYW1pbHlfbmFtZSI6Ikdhcm9vc2kiLCJsb2NhbGUiOiJlbi1HQiIsImlhdCI6MTU1OTU5MjI0NywiZXhwIjoxNTU5NTk1ODQ3LCJqdGkiOiJjYzhhOWQxODc1ZGQxY2JjNTBjMjM3OGUwMThmYjU0OTQ5MTEzZjZlIn0.rwok2_3q_wYwu31swLadAcK8MdJ35Dtg7QyhPo5bdGH8qopnh15TFPojFPojyWiC_-rxo5kjG-WV4iuZni0p4tLMbLhKs_NtapJ-NhTRvRiVe1Cg4zBXPWMk5JcqaLKx6bsXeNwRHhDFGmuVy0DRCPrco9vDkWnK6CUARf2i-0sKxXULNPrkJKLtZWtQZHkBPw3MUfig820eDP1aLucrCBcoHtWA3BAKoW7dcMSdvB3PW0ZIjd9XKoDbw_9K2pCzf5sMfMin8itcoHgXXLZM0XcwCEIaK8Gel9GrpkJKHw4mAeHXhuvUbMvliPQ-sTBbHkjddZnvIjikvRFBi4WuWQ",
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
