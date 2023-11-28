package storage

import (
	"context"
	pd "kassa360/kassa360_go_dynamic_service/genproto/dynamic_service"

	"github.com/jellydator/ttlcache/v3"
	"github.com/saidamir98/udevs_pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
)

func StartTTLCache(log logger.LoggerI, cache *ttlcache.Cache[string, string]) error {
	log.Info("TTLCache: started automatic cleanup process of expired items")
	cache.Start()
	return nil
}

type StorageI interface {
	Disconnect() error
	Group() GroupI
	Entity() EntityI
}

type GroupI interface {
	CreateGroup(ctx context.Context, req *pd.Group) error
	GetGroupById(ctx context.Context, req *pd.GetByIdRequest) (*pd.Group, error)
	DeleteGroup(ctx context.Context, req *pd.GetByIdRequest) error
	UpdateGroup(ctx context.Context, req *pd.Group) error
	GetAllGroup(ctx context.Context, req *pd.GetAllGroupRequest) (*pd.GetAllGroupResponse, error)
	CreateField(ctx context.Context, req *pd.Field) error
	GetFieldById(ctx context.Context, req *pd.GetByIdRequest) (*pd.Field, error)
	DeleteField(ctx context.Context, req *pd.GetByIdRequest) error
	UpdateField(ctx context.Context, req *pd.Field) error
	GetAllField(ctx context.Context, req *pd.GetAllFieldRequest) (*pd.GetAllFieldResponse, error)
	GetFullGroup(ctx context.Context, req *pd.GetByIdRequest) (*pd.Group, error)
}

type EntityI interface {
	Create(ctx context.Context, slug string, body map[string]interface{}) error
	Update(ctx context.Context, slug string, id string, body map[string]interface{}) error
	Delete(ctx context.Context, slug string, id string) error
	Get(ctx context.Context, slug string, id string) (map[string]interface{}, error)
	List(ctx context.Context, slug, order, sort string, limit, offset int32, filter bson.D) ([]map[string]interface{}, error)
	JoinList(ctx context.Context, slug, order, sort string, limit, offset int32, filter bson.A, lookups []*pd.LookUps) ([]map[string]interface{}, error)
	Count(ctx context.Context, slug string, filter bson.D) (int64, error)
	JoinCount(ctx context.Context, slug string, filter bson.A) (int64, error)
	QueryFilter(req map[string]interface{}, group *pd.Group, search, location string) (bson.D, error)
	JoinQueryFilter(req map[string]interface{}, group *pd.Group, search, location string) (bson.A, error)
}

type StoragePgI interface {
	CloseDB()
}

type EntityPgI interface {
}
