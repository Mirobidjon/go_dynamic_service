package event

import (
	"context"
	"errors"
	"time"

	"kassa360/kassa360_go_dynamic_service/pkg/helper"

	"github.com/IBM/sarama"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/saidamir98/udevs_pkg/logger"
)

type Error struct {
	Message string `json:"code"`
	Error   string `json:"message"`
}

type Response struct {
	NoResponse bool        `json:"-"`
	SessionID  string      `json:"session_id,omitempty" swaggerignore:"true"`
	StatusCode int32       `json:"status_code,omitempty"`
	ID         string      `json:"id,omitempty"`
	Error      Error       `json:"error,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

type HandlerFunc func(context.Context, cloudevents.Event) Response

// Consumer ...
type Consumer struct {
	ctx          context.Context
	consumerName string
	topic        string
	handler      HandlerFunc
}

// AddConsumer ...
func (kafka *Kafka) AddConsumer(consumerName, topic, groupID string, handler HandlerFunc) {
	if kafka.consumers[consumerName] != nil {
		panic(errors.New("consumer with the same name already exists: " + consumerName))
	}

	kafka.consumers[consumerName] = &Consumer{
		ctx:          kafka.ctx,
		consumerName: consumerName,
		topic:        topic,
		handler:      handler,
	}
}

// Setup is run at the beginning of a new session, before ConsumeClaim.
func (c *Kafka) Setup(_ sarama.ConsumerGroupSession) error {
	close(c.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
// but before the offsets are committed for the very last time.
func (c *Kafka) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// Once the Messages() channel is closed, the Handler must finish its processing
// loop and exit.
func (c *Kafka) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	errChan := make(chan error, 2)
	for message := range claim.Messages() {
		resp := make(chan Response, 2)

		go func() {
			select {
			case res := <-resp:
				if res.Error.Error != "" {
					c.log.Error("ERROR_KAFKA_TOPIC", logger.String("topic", message.Topic), logger.String("message", res.Error.Message), logger.String("error", res.Error.Error))
				}
				errChan <- nil
			case <-c.ctx.Done():
				c.log.Error("ERROR_KAFKA_CTX_DONE", logger.String("topic", message.Topic), logger.String("time", time.Now().Format("2006-01-02 15:04:05")))
				errChan <- c.ctx.Err()
			case <-time.After(time.Second * 60 * 10):
				c.log.Error("ERROR_KAFKA_TIMEOUT", logger.String("topic", message.Topic), logger.String("time", time.Now().Format("2006-01-02 15:04:05")))
				errChan <- nil
			}
		}()

		event := helper.MessageToEvent(message)

		consumer, ok := c.consumers[message.Topic]
		if ok {
			resp <- consumer.handler(c.ctx, event)
		} else {
			resp <- c.handler(c.ctx, event)
		}

		err := <-errChan
		if err != nil {
			return err
		}
	}
	return nil
}

func (k *Kafka) handler(ctx context.Context, e cloudevents.Event) Response {
	var r Response
	k.log.Info("WRONG_EVENT", logger.Any("event", e))
	return r
}

func HandleResponse(message string, err error, data interface{}) Response {
	if err != nil {
		return Response{
			NoResponse: true,
			Error: Error{
				Message: message,
				Error:   err.Error(),
			},
		}
	}

	return Response{
		NoResponse: true,
		Data:       data,
	}
}
