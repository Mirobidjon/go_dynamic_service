package client

import (
	"github.com/mirobidjon/go_dynamic_service/config"
	pd "github.com/mirobidjon/go_dynamic_service/genproto/dynamic_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManagerI interface {
	DynamicService() pd.DynamicServiceClient
	EntityService() pd.EntityServiceClient
}

type grpcClients struct {
	dynamicService pd.DynamicServiceClient
	entityService  pd.EntityServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {
	maxMsgSize := 10 * 1024 * 1024 * 1024
	connDynamicService, err := grpc.Dial(
		cfg.DynamicServiceHost+cfg.DynamicGRPCPort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxMsgSize), grpc.MaxCallSendMsgSize(maxMsgSize)),
	)
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		dynamicService: pd.NewDynamicServiceClient(connDynamicService),
		entityService:  pd.NewEntityServiceClient(connDynamicService),
	}, nil
}

func (g *grpcClients) DynamicService() pd.DynamicServiceClient {
	return g.dynamicService
}

func (g *grpcClients) EntityService() pd.EntityServiceClient {
	return g.entityService
}
