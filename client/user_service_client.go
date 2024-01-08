package client

import (
	"fmt"
	"os"

	"github.com/loak155/microservices-proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IUserGRPCClient interface {
	GetUser(req *pb.GetUserRequest) (*pb.getUserResponse, error)
}

type userGRPCClient struct {
	client pb.UserServiceClient
}

func NewUserGRPCClient() (IUserGRPCClient, error) {
	address := fmt.Sprintf("%s:%s", os.Getenv("USER_SERVICE_HOST"), os.Getenv("USER_SERVICE_PORT"))
	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, err
	}
	client := pb.NewUserServiceClient(conn)

	return &userGRPCClient{client}, nil
}

func (c *userGRPCClient) GetUser(req *pb.GetUserRequest) (*pb.getUserResponse, error) {
	res, err := c.client.GetUser(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
