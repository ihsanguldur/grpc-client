package auth

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-api/pkg/auth/pb"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func NewServiceClient(url string) pb.AuthServiceClient {
	cc, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect auth server", err)
	}

	return pb.NewAuthServiceClient(cc)
}
