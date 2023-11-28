package events

import (
	"context"

	"kassa360/kassa360_go_dynamic_service/config"

	"kassa360/kassa360_go_dynamic_service/pkg/event"

	"github.com/saidamir98/udevs_pkg/logger"
)

// PubsubServer ...
type PubsubServer struct {
	cfg   config.Config
	log   logger.LoggerI
	kafka *event.Kafka
}

type EventParams struct {
	Cfg   config.Config
	Log   logger.LoggerI
	Kafka *event.Kafka
}

// New ...
func New(params EventParams) (*PubsubServer, error) {
	kafka := event.NewKafka(context.Background(), params.Cfg, params.Log)

	return &PubsubServer{
		cfg:   params.Cfg,
		log:   params.Log,
		kafka: kafka,
	}, nil
}

// Run ...
// func (s *PubsubServer) Run(ctx context.Context, storage storage.StorageI, cl client.ServiceManagerI) {
// 	waybillService := waybill.New(
// 		s.cfg,
// 		s.log,
// 		s.kafka,
// 		service.NewWaybillService(s.log, storage, cl, s.cfg),
// 	)

// 	waybillService.RegisterConsumers()

// 	s.kafka.RunConsumers(ctx)
// }
