package waybill

import (
	"kassa360/kassa360_go_dynamic_service/pkg/event"

	"github.com/saidamir98/udevs_pkg/logger"
)

type WaybillService struct {
	// cfg config.Config
	log logger.LoggerI

	kafka *event.Kafka
	// waybill *service.WaybillService
}

// func New(cfg config.Config, log logger.LoggerI, kafka *event.Kafka, waybill *service.WaybillService) *WaybillService {
// 	return &WaybillService{
// 		cfg:   cfg,
// 		log:   log,
// 		kafka: kafka,
// 		// waybill: waybill,
// 	}
// }

func (c *WaybillService) RegisterConsumers() {
	route := event.MONGO_DOCUMENT_SERVICE_TOPIC + "waybill."

	c.kafka.AddConsumer(
		route+"upsert", // consumer name
		route+"upsert", // topic
		route+"upsert", // group id
		c.Upsert,       // handlerFunction
	)

	c.kafka.AddConsumer(
		route+"update_status", // consumer name
		route+"update_status", // topic
		route+"update_status", // group id
		c.UpdateStatus,        // handlerFunction
	)
}
