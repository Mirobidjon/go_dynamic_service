package event

import (
	"encoding/json"
	"fmt"
	"kassa360/kassa360_go_dynamic_service/config"
	"kassa360/kassa360_go_dynamic_service/pkg/helper"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/google/uuid"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (r *Kafka) PushStructEvent(topic, source string, v interface{}) (string, error) {
	var (
		js  []byte
		err error
	)

	event := cloudevents.NewEvent()

	correlationID := uuid.NewString()
	event.SetID(correlationID)
	event.SetSource(source)
	event.SetType(config.StructType)

	if js, err = json.Marshal(v); err != nil {
		return correlationID, fmt.Errorf("error while parse to string proto data: %w", err)
	}

	err = event.SetData(cloudevents.ApplicationJSON, js)
	if err != nil {
		return correlationID, fmt.Errorf("error while setting data to event: %w", err)
	}

	err = r.Push(topic, event)
	if err != nil {
		return correlationID, fmt.Errorf("error while pushing data to topic: %w", err)
	}

	return correlationID, err
}

func (r *Kafka) PushProtoEvent(topic, source string, v protoreflect.ProtoMessage) (string, error) {
	var (
		js  string
		err error
	)

	event := cloudevents.NewEvent()

	correlationID := uuid.NewString()
	event.SetID(correlationID)
	event.SetSource(source)
	event.SetType(config.ProtoType)

	if js, err = helper.ProtoToString(v); err != nil {
		return correlationID, fmt.Errorf("error while parse to string proto data: %w", err)
	}

	err = event.SetData(cloudevents.ApplicationJSON, []byte(js))
	if err != nil {
		return correlationID, fmt.Errorf("error while setting data to event: %w", err)
	}

	err = r.Push(topic, event)
	if err != nil {
		return correlationID, fmt.Errorf("error while pushing data to topic: %w", err)
	}

	return correlationID, nil
}
