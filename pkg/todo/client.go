package todo

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-api/pkg/todo/pb"
)

type ServiceClient struct {
	Client pb.TodoServiceClient
}

func NewTodoServiceClient(url string) pb.TodoServiceClient {
	cc, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect todo server", err)
	}

	return pb.NewTodoServiceClient(cc)
}
