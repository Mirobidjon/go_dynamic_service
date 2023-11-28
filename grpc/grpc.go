package grpc

import (
	"context"
	"fmt"
	"kassa360/kassa360_go_dynamic_service/config"
	pb "kassa360/kassa360_go_dynamic_service/genproto/dynamic_service"
	"kassa360/kassa360_go_dynamic_service/pkg/helper"
	"kassa360/kassa360_go_dynamic_service/service"
	"kassa360/kassa360_go_dynamic_service/storage"
	"net"
	"runtime"

	"github.com/saidamir98/udevs_pkg/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGRPCServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI) error {
	grpcServer := SetUpServer(cfg, log, strg)
	
	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Error("error while listen RPC", logger.Error(err))
		return err
	}

	log.Info("GRPC: Server being started...", logger.String("port", cfg.RPCPort))
	if err := grpcServer.Serve(lis); err != nil {
		log.Error("grpcServer.Serve", logger.Error(err))
		return err
	}

	return nil
}

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI) (grpcServer *grpc.Server) {
	maxMsgSize := 100 * 1024 * 1024 * 1024

	grpcServer = grpc.NewServer(
		grpc.MaxRecvMsgSize(maxMsgSize),
		grpc.MaxSendMsgSize(maxMsgSize),
		grpc.UnaryInterceptor(Middleware),
	)

	group := service.NewGroupService(log, strg, &cfg)
	entity := service.NewEntityService(log, strg, &cfg)

	pb.RegisterDynamicServiceServer(grpcServer, group)
	pb.RegisterEntityServiceServer(grpcServer, entity)

	reflection.Register(grpcServer)
	return
}

func Middleware(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	defer handlePanic()

	return handler(ctx, req)
}

func handlePanic() {
	if r := recover(); r != nil {
		errMsg := fmt.Sprintf("Panic occurred\nerror --> %v\n", r)

		depth := 1
		for {
			_, file, line, ok := runtime.Caller(depth)
			if !ok {
				break
			}

			errMsg += fmt.Sprintf("\n%d --> %s:%d", depth, file, line)
			depth++
		}

		err := helper.SendTelegramMessage(errMsg)
		if err != nil {
			fmt.Printf("failed to send message on tg: %v", err)
		}
	}
}
