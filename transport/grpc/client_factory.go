package grpc

import (
	pb "distributed-algorithms/generated/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientFactory struct{}

func (factory *ClientFactory) NewClient(address string) (*Client, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	gRPCConnection, err := grpc.NewClient(address, opts...)

	if err != nil {
		return nil, err
	}

	return &Client{
		gRPCClient:     pb.NewRaftServiceClient(gRPCConnection),
		gRPCConnection: gRPCConnection,
	}, nil
}
