package helper

import (
	"encoding/json"
	"fmt"
	"kassa360/kassa360_go_dynamic_service/config"
	"time"

	"github.com/IBM/sarama"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/google/uuid"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func MessageToEvent(message *sarama.ConsumerMessage) cloudevents.Event {
	event := cloudevents.NewEvent()

	for _, header := range message.Headers {
		if x := string(header.Key); x == "ce_id" {
			event.SetID(string(header.Value))
		} else if x == "ce_source" {
			event.SetSource(string(header.Value))
		} else if x == "ce_type" {
			event.SetType(string(header.Value))
		} else if x == "ce_time" {
			t, _ := time.Parse("2006-01-02T15:04:05.999999999Z", string(header.Value))
			event.SetTime(t)
		} else if x == "content-type" {
			event.SetDataContentType(string(header.Value))
		} else if x == "ce_specversion" {
			event.SetSpecVersion(string(header.Value))
		} else {
			fmt.Println("not equal: ", x, string(header.Value))
		}
	}

	if err := event.SetData(cloudevents.ApplicationJSON, message.Value); err != nil {
		fmt.Println("error while setting data to event: ", err)
	}

	return event
}

func CreateStructEvent(source string, v interface{}) (cloudevents.Event, uuid.UUID, error) {
	var (
		js  []byte
		err error
	)

	event := cloudevents.NewEvent()

	correlationID, _ := uuid.NewRandom()
	event.SetID(correlationID.String())
	event.SetSource(source)
	event.SetType(config.StructType)

	if js, err = json.Marshal(v); err != nil {
		return event, correlationID, err
	}

	err = event.SetData(cloudevents.ApplicationJSON, js)

	return event, correlationID, err
}

func CreateProtoEvent(source string, v protoreflect.ProtoMessage) (cloudevents.Event, uuid.UUID, error) {
	var (
		js  string
		err error
	)

	event := cloudevents.NewEvent()

	correlationID, _ := uuid.NewRandom()
	event.SetID(correlationID.String())
	event.SetSource(source)
	event.SetType(config.ProtoType)

	if js, err = ProtoToString(v); err != nil {
		return event, correlationID, err
	}

	err = event.SetData(cloudevents.ApplicationJSON, []byte(js))

	return event, correlationID, err
}
