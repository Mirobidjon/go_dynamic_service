package event

import (
	"context"
	"errors"
	"fmt"

	"github.com/saidamir98/udevs_pkg/logger"

	"github.com/IBM/sarama"
	"github.com/cloudevents/sdk-go/protocol/kafka_sarama/v2"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

// Publisher ...
type Publisher struct {
	topic            string
	cloudEventClient cloudevents.Client
	sender           *kafka_sarama.Sender
}

// AddPublisher ...
func (kafka *Kafka) AddPublisher(topic string) {
	if kafka.publishers[topic] != nil {
		kafka.log.Warn("publisher exists", logger.Error(errors.New("publisher with the same topic already exists: "+topic)))
		return
	}

	// sender, err := kafka_sarama.NewSender(
	// 	[]string{kafka.cfg.KafkaUrlBroker1, kafka.cfg.KafkaUrlBroker2, kafka.cfg.KafkaUrlBroker3}, // Kafka connection url
	// 	kafka.saramaConfig, // Kafka sarama config
	// 	topic,              // Topic
	// )
	// if err != nil {
	// 	panic(err)
	// }

	var sender *kafka_sarama.Sender

	defer sender.Close(context.Background())

	c, err := cloudevents.NewClient(sender, cloudevents.WithTimeNow(), cloudevents.WithUUIDs())
	if err != nil {
		panic(err)
	}

	kafka.publishers[topic] = &Publisher{
		topic:            topic,
		cloudEventClient: c,
		sender:           sender,
	}
}

// Push ...
func (r *Kafka) Push(topic string, e cloudevents.Event) error {
	p := r.publishers[topic]

	if p == nil {
		return errors.New("publisher with that topic doesn't exists: " + topic)
	}

	result := p.cloudEventClient.Send(
		kafka_sarama.WithMessageKey(context.Background(), sarama.StringEncoder(e.ID())),
		e,
	)

	if cloudevents.IsUndelivered(result) {
		return errors.New("failed to publish event")
	}

	fmt.Println("Pushed to ", topic)

	return nil
}
