package informer

import (
	grpcConf "github.com/strugglehonor/KCS/internal/config/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type grpcClient struct {
	address string
	client *grpc.ClientConn
}

func GetGrpcConn() (*grpcClient, error) {
	conn, err := grpc.Dial(grpcConf.Address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	return &grpcClient{address: grpcConf.Address, client: conn}, nil
}