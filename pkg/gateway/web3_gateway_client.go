package gateway

import (
	"context"

	"github.com/project-nova/backend/pkg/logger"
	"github.com/project-nova/backend/proto/v1/web3_gateway"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Web3GatewayClient interface {
	UploadContent(*web3_gateway.UploadContentReq) (*web3_gateway.UploadContentResp, error)
}

type grpcWeb3GatewayClient struct {
	storageService web3_gateway.StorageServiceClient
}

func NewWeb3GatewayClient(connString string) (Web3GatewayClient, error) {
	logger.Infof("Connecting to web3-gateway server %s", connString)
	conn, err := grpc.Dial(connString, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Errorf("cannot connect to grpc server: %v", err)
		return nil, err
	}
	logger.Infof("Successfully connected to web3-gateway server %s", connString)
	return &grpcWeb3GatewayClient{
		storageService: web3_gateway.NewStorageServiceClient(conn),
	}, nil
}

func (g *grpcWeb3GatewayClient) UploadContent(req *web3_gateway.UploadContentReq) (*web3_gateway.UploadContentResp, error) {
	return g.storageService.UploadContent(context.Background(), req)
}
