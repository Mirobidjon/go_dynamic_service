package main

import (
	"context"
	"kassa360/kassa360_go_dynamic_service/api"
	"kassa360/kassa360_go_dynamic_service/config"
	"kassa360/kassa360_go_dynamic_service/grpc"
	"kassa360/kassa360_go_dynamic_service/storage"
	"kassa360/kassa360_go_dynamic_service/storage/mongodb"

	"github.com/saidamir98/udevs_pkg/logger"
	"golang.org/x/sync/errgroup"
)

func main() {
	var (
		loggerLevel string
		cfg         = config.Load()
	)

	switch cfg.Environment {
	case config.DebugMode:
		loggerLevel = logger.LevelDebug
	case config.TestMode:
		loggerLevel = logger.LevelDebug
	default:
		loggerLevel = logger.LevelInfo
	}

	log := logger.NewLogger(cfg.ServiceName, loggerLevel)
	defer func() {
		if err := logger.Cleanup(log); err != nil {
			log.Error("Failed to cleanup logger", logger.Error(err))
		}
	}()

	store, cache := mongodb.NewStoragePg(cfg, log)
	defer func() {
		if err := store.Disconnect(); err != nil {
			log.Error(err.Error())
		}
	}()

	group, _ := errgroup.WithContext(context.Background())
	group.Go(func() error {
		return storage.StartTTLCache(log, cache)
	})

	group.Go(func() error {
		return grpc.StartGRPCServer(cfg, log, store)
	})

	group.Go(func() error {
		return api.StartHTTPServer(cfg, log)
	})

	if err := group.Wait(); err != nil {
		log.Error("error while running services", logger.Error(err))
	}

	log.Info("server stopped")
}
