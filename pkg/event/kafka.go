package event

import (
	"context"
	"fmt"
	"sync"

	"kassa360/kassa360_go_dynamic_service/config"

	"github.com/saidamir98/udevs_pkg/logger"

	"github.com/IBM/sarama"
)

var (
	DOCUMENT_SERVICE_TOPIC       = "v1.document_service."
	MONGO_DOCUMENT_SERVICE_TOPIC = "v1.mongo_document_service."
)

type Kafka struct {
	ctx           context.Context
	log           logger.LoggerI
	cfg           config.Config
	publishers    map[string]*Publisher
	consumers     map[string]*Consumer
	saramaConfig  *sarama.Config
	consumerGroup sarama.ConsumerGroup
	Responses     chan Response
	ready         chan struct{}
	wg            *sync.WaitGroup
}

func NewKafka(ctx context.Context, cfg config.Config, log logger.LoggerI) *Kafka {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Version = sarama.V2_0_0_0
	// saramaConfig.Consumer.Group.Heartbeat.Interval = time.Second * 20
	// saramaConfig.Consumer.Group.Session.Timeout = time.Minute
	consumerGroup, err := sarama.NewConsumerGroup(
		[]string{cfg.KafkaUrlBroker1, cfg.KafkaUrlBroker2, cfg.KafkaUrlBroker3},
		config.ConsumerGroupID,
		saramaConfig,
	)
	if err != nil {
		panic(err)
	}

	kafka := &Kafka{
		ctx:           ctx,
		log:           log,
		cfg:           cfg,
		publishers:    make(map[string]*Publisher),
		consumers:     make(map[string]*Consumer),
		saramaConfig:  saramaConfig,
		consumerGroup: consumerGroup,
		Responses:     make(chan Response),
		ready:         make(chan struct{}),
		wg:            &sync.WaitGroup{},
	}

	kafka.RegisterPublishers()

	return kafka
}

func (k *Kafka) RegisterPublishers() {
	k.AddPublisher(DOCUMENT_SERVICE_TOPIC)
	k.AddPublisher(MONGO_DOCUMENT_SERVICE_TOPIC)
}

// RunConsumers ...
func (r *Kafka) RunConsumers(ctx context.Context) {
	topics := []string{}

	for _, consumer := range r.consumers {
		topics = append(topics, consumer.topic)
		fmt.Println("Key:", consumer.topic, "=>", "consumer:", consumer)
	}
	r.log.Info("topics:", logger.Any("topics:", topics))

	r.wg.Add(1)
	go func() {
		defer r.wg.Done()
		for {
			if err := r.consumerGroup.Consume(r.ctx, topics, r); err != nil {
				r.log.Error("error while consuming", logger.Error(err))
			}
			if r.ctx.Err() != nil {
				return
			}
			r.ready = make(chan struct{})
		}
	}()

	<-r.ready
	r.wg.Wait()
	r.log.Warn("consumer group started")
}

func (r *Kafka) Shutdown(ctx context.Context) error {
	r.log.Warn("shutting down pub-sub server")
	select {
	case <-r.ctx.Done():
		r.log.Warn("terminating: context cancelled")
	default:
	}
	r.wg.Wait()
	r.consumerGroup.Close()

	for _, client := range r.publishers {
		if err := client.sender.Close(ctx); err != nil {
			r.log.Error("could not close sender", logger.Any("topic", client.topic), logger.Error(err))
		}
	}

	return nil
}
