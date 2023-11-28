package waybill

import (
	"context"
	"kassa360/kassa360_go_dynamic_service/pkg/event"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/saidamir98/udevs_pkg/logger"
)

func (c *WaybillService) Upsert(ctx context.Context, e cloudevents.Event) event.Response {
	// var (
	// 	req pb.Waybill
	// )

	// c.log.Info("Upsert waybill req")

	// err := helper.ByteToProto(&req, e.DataEncoded)
	// if err != nil {
	// 	return event.HandleResponse("unmarshall error", err, &req)
	// }

	// _, err = c.waybill.Upsert(ctx, &req)
	// if err != nil {
	// 	c.log.Error("Upsert waybill", logger.Any("waybill", string(e.DataEncoded)))
	// 	return event.HandleResponse("error while upserting waybill", err, &req)
	// }

	c.log.Info("Upsert Waybill Succsess")

	return event.HandleResponse("success", nil, "OK")
}

func (c *WaybillService) UpdateStatus(ctx context.Context, e cloudevents.Event) event.Response {
	// var (
	// 	req pb.UpdateStatusReq
	// )

	c.log.Info("UpdateStatus Waybill", logger.Any("req", e))

	// err := helper.ByteToProto(&req, e.DataEncoded)
	// if err != nil {
	// 	return event.HandleResponse("unmarshall error", err, &req)
	// }

	// resp, err := c.waybill.UpdateStatus(ctx, &req)
	// if err != nil {
	// 	return event.HandleResponse("error while UpdateStatus waybill", err, &req)
	// }

	// c.log.Info("UpdateStatus Waybill", logger.Any("resp", resp))

	// return event.HandleResponse("success", nil, resp)
	return event.HandleResponse("success", nil, nil)
}
