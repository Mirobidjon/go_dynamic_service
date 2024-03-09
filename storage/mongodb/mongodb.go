package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/mirobidjon/go_dynamic_service/config"
	"github.com/mirobidjon/go_dynamic_service/storage"

	"github.com/jellydator/ttlcache/v3"
	"github.com/saidamir98/udevs_pkg/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type storagePg struct {
	mongoConn  *mongo.Client
	groupRepo  storage.GroupI
	entityRepo storage.EntityI
}

func NewStoragePg(cfg config.Config, log logger.LoggerI) (storage.StorageI, *ttlcache.Cache[string, string]) {
	credential := options.Credential{
		Username:      cfg.MongoUser,
		Password:      cfg.MongoPassword,
		AuthMechanism: "SCRAM-SHA-256",
		AuthSource:    cfg.MongoDatabase,
	}

	log.Info("connecting to mongodb...", logger.Any("url", fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", cfg.MongoUser, cfg.MongoPassword, cfg.MongoHost1, cfg.MongoPort, cfg.MongoDatabase)))

	mongoString := fmt.Sprintf("mongodb://%s:%d", cfg.MongoHost1, cfg.MongoPort)
	mongoConn, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoString).SetAuth(credential))
	if err != nil {
		log.Fatal("error to connect to mongo database", logger.Error(err))
	}

	if err := mongoConn.Ping(context.TODO(), nil); err != nil {
		log.Fatal("Cannot connect to database error -> ", logger.Error(err))
	}
	log.Info("connected to mongodb")

	db := mongoConn.Database(cfg.MongoDatabase)

	cache := ttlcache.New[string, string](
		ttlcache.WithTTL[string, string](time.Duration(cfg.CacheTTL) * time.Minute),
	)

	return &storagePg{
		mongoConn:  mongoConn,
		groupRepo:  NewGroupRepo(db, cache),
		entityRepo: NewEntityRepo(db),
	}, cache
}

func (s *storagePg) Disconnect() error {
	return s.mongoConn.Disconnect(context.Background())
}

func (s *storagePg) Group() storage.GroupI {
	return s.groupRepo
}

func (s *storagePg) Entity() storage.EntityI {
	return s.entityRepo
}
