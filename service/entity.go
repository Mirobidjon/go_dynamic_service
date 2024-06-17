package service

import (
	"context"
	"fmt"

	"github.com/mirobidjon/go_dynamic_service/config"
	pb "github.com/mirobidjon/go_dynamic_service/genproto/dynamic_service"
	"github.com/mirobidjon/go_dynamic_service/pkg/helper"
	"github.com/mirobidjon/go_dynamic_service/storage"

	log "github.com/saidamir98/udevs_pkg/logger"
	"github.com/spf13/cast"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/structpb"
)

type entityRepo struct {
	log  log.LoggerI
	stg  storage.StorageI
	conf *config.Config
	pb.UnimplementedEntityServiceServer
}

func NewEntityService(log log.LoggerI, storage storage.StorageI, conf *config.Config) *entityRepo {
	return &entityRepo{
		log:  log,
		stg:  storage,
		conf: conf,
	}
}

func (s *entityRepo) Create(ctx context.Context, entity *pb.Entity) (*pb.Entity, error) {
	s.log.Info("Create", log.Any("req", entity))

	if entity.Slug == "" {
		return nil, helper.HandleError(s.log, nil, "Slug is required", entity, codes.InvalidArgument)
	}

	body := entity.Data.AsMap()
	body["_id"] = nil

	group, err := s.stg.Group().GetFullGroup(ctx, &pb.GetByIdRequest{XId: entity.Slug})
	if err != nil {
		return nil, helper.HandleError(s.log, err, "error while getting group", entity, codes.Internal)
	}

	err = helper.CheckData(body, group, &entity.Location)
	if err != nil {
		return nil, helper.HandleError(s.log, err, "error while checking body ", entity, codes.Internal)
	}

	err = s.stg.Entity().Create(ctx, entity.Slug, body)
	if err != nil {
		return nil, helper.HandleError(s.log, err, "error while creating "+entity.Slug, entity, codes.Internal)
	}
	entity.XId = cast.ToString(body["_id"])

	entity.Data, err = helper.ToProtoStruct(body)
	if err != nil {
		return nil, helper.HandleError(s.log, err, "error while converting to proto struct", entity, codes.Internal)
	}

	return entity, nil
}

func (s *entityRepo) UpdatePatch(ctx context.Context, entity *pb.Entity) (*pb.Entity, error) {
	s.log.Info("UpdatePatch", log.Any("req", entity))

	body := entity.Data.AsMap()

	if !helper.IsValidObjectId(body["_id"]) {
		return nil, helper.HandleError(s.log, fmt.Errorf("invalid object id"), "error while checking body ", entity, codes.InvalidArgument)
	}

	if entity.Slug == "" {
		return nil, helper.HandleError(s.log, fmt.Errorf("slug is empty"), "error while checking body ", entity, codes.InvalidArgument)
	}

	group, err := s.stg.Group().GetFullGroup(ctx, &pb.GetByIdRequest{XId: entity.Slug})
	if err != nil {
		return nil, helper.HandleError(s.log, err, "error while getting group", entity, codes.Internal)
	}

	err = helper.CheckDataForPatch(body, group, &entity.Location)
	if err != nil {
		return nil, helper.HandleError(s.log, err, "error while checking body ", entity, codes.Internal)
	}

	err = s.stg.Entity().Update(ctx, entity.Slug, entity.XId, body)
	if err != nil {
		return nil, helper.HandleError(s.log, err, "error while updating "+entity.Slug, entity, codes.Internal)
	}

	entity.Data, err = helper.ToProtoStruct(body)
	if err != nil {
		return nil, helper.HandleError(s.log, err, "error while converting to proto struct", entity, codes.Internal)
	}

	return entity, nil
}

func (s *entityRepo) Update(ctx context.Context, entity *pb.Entity) (*pb.Entity, error) {
	s.log.Info("Update", log.Any("req", entity))

	body := entity.Data.AsMap()

	if !helper.IsValidObjectId(body["_id"]) {
		return nil, helper.HandleError(s.log, fmt.Errorf("invalid object id"), "error while checking body ", entity, codes.InvalidArgument)
	}

	if entity.Slug == "" {
		return nil, helper.HandleError(s.log, fmt.Errorf("slug is empty"), "error while checking body ", entity, codes.InvalidArgument)
	}

	group, err := s.stg.Group().GetFullGroup(ctx, &pb.GetByIdRequest{XId: entity.Slug})
	if err != nil {
		return nil, helper.HandleError(s.log, err, "error while getting group", entity, codes.Internal)
	}

	err = helper.CheckData(body, group, &entity.Location)
	if err != nil {
		return nil, helper.HandleError(s.log, err, "error while checking body ", entity, codes.Internal)
	}

	err = s.stg.Entity().Update(ctx, entity.Slug, entity.XId, body)
	if err != nil {
		return nil, helper.HandleError(s.log, err, "error while updating "+entity.Slug, entity, codes.Internal)
	}

	entity.Data, err = helper.ToProtoStruct(body)
	if err != nil {
		return nil, helper.HandleError(s.log, err, "error while converting to proto struct", entity, codes.Internal)
	}

	return entity, nil
}

func (s *entityRepo) Delete(ctx context.Context, entity *pb.Entity) (*emptypb.Empty, error) {
	s.log.Info("Delete", log.Any("req", entity))

	if !helper.IsValidObjectId(entity.XId) {
		return nil, helper.HandleError(s.log, fmt.Errorf("invalid object id"), "error while checking body ", entity, codes.InvalidArgument)
	}

	if entity.Slug == "" {
		return nil, helper.HandleError(s.log, fmt.Errorf("slug is empty"), "error while checking body ", entity, codes.InvalidArgument)
	}

	err := s.stg.Entity().Delete(ctx, entity.Slug, entity.XId)
	if err != nil {
		return nil, helper.HandleError(s.log, err, "error while deleting "+entity.Slug, entity, codes.Internal)
	}

	return &emptypb.Empty{}, nil
}

func (s *entityRepo) GetById(ctx context.Context, req *pb.GetByPk) (*pb.Entity, error) {
	s.log.Info("Get", log.Any("req", req))

	if !helper.IsValidObjectId(req.XId) {
		return nil, helper.HandleError(s.log, fmt.Errorf("invalid object id"), "error while checking body ", req, codes.InvalidArgument)
	}

	if req.Slug == "" {
		return nil, helper.HandleError(s.log, fmt.Errorf("slug is empty"), "error while checking body ", req, codes.InvalidArgument)
	}

	body, err := s.stg.Entity().Get(ctx, req.Slug, req.XId)
	if err != nil {
		return nil, helper.HandleError(s.log, err, "error while getting "+req.Slug, req, codes.Internal)
	}

	entity := &pb.Entity{
		Slug: req.Slug,
		XId:  req.XId,
		Data: &structpb.Struct{},
	}

	body["created_at"], _ = helper.ToLocationTime(cast.ToString(body["created_at"]), req.Location)
	body["updated_at"], _ = helper.ToLocationTime(cast.ToString(body["updated_at"]), req.Location)

	entity.Data, err = helper.ToProtoStruct(body)
	if err != nil {
		return nil, helper.HandleError(s.log, err, "error while converting to proto struct", entity, codes.Internal)
	}

	return entity, nil
}

func (s *entityRepo) GetAll(ctx context.Context, req *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	var (
		res = &pb.GetAllResponse{}
	)
	s.log.Info("GetAll", log.Any("req", req))

	if req.Slug == "" {
		return nil, helper.HandleError(s.log, fmt.Errorf("slug is required"), "error while checking body ", req, codes.InvalidArgument)
	}

	data := req.Data.AsMap()

	group, err := s.stg.Group().GetFullGroup(ctx, &pb.GetByIdRequest{XId: req.Slug})
	if err != nil {
		return nil, helper.HandleError(s.log, err, "error while getting group", req, codes.Internal)
	}

	filter, err := s.stg.Entity().QueryFilter(data, group, req.Search, req.Location)
	if err != nil {
		return nil, helper.HandleError(s.log, err, "error while getting filter", req, codes.Internal)
	}

	entities, err := s.stg.Entity().List(ctx, req.Slug, req.Order, req.Sort, req.Limit, req.Offset, filter)
	if err != nil {
		return nil, helper.HandleError(s.log, err, "error while getting "+req.Slug, req, codes.Internal)
	}

	count, err := s.stg.Entity().Count(ctx, req.Slug, filter)
	if err != nil {
		return nil, helper.HandleError(s.log, err, "error while getting count "+req.Slug, req, codes.Internal)
	}

	for _, item := range entities {
		item["created_at"], _ = helper.ToLocationTime(cast.ToString(item["created_at"]), req.Location)
		item["updated_at"], _ = helper.ToLocationTime(cast.ToString(item["updated_at"]), req.Location)

		entity, err := helper.ToProtoStruct(item)
		if err != nil {
			return nil, helper.HandleError(s.log, err, "error while converting to proto struct", entity, codes.Internal)
		}

		res.Entities = append(res.Entities, entity)
	}

	res.Count = int32(count)

	return res, nil
}

func (s *entityRepo) GetJoin(ctx context.Context, req *pb.GetJoinRequest) (*pb.GetAllResponse, error) {
	var (
		res = &pb.GetAllResponse{}
	)
	s.log.Info("GetAll", log.Any("req", req))

	if req.Slug == "" {
		return nil, helper.HandleError(s.log, fmt.Errorf("slug is required"), "error while checking body ", req, codes.InvalidArgument)
	}

	data := req.Data.AsMap()
	group, err := s.stg.Group().GetFullGroup(ctx, &pb.GetByIdRequest{XId: req.Slug})
	if err != nil {
		return nil, helper.HandleError(s.log, err, "error while getting group", req, codes.Internal)
	}

	filter, err := s.stg.Entity().JoinQueryFilter(data, group, req.Search, req.Location)
	if err != nil {
		return nil, helper.HandleError(s.log, err, "error while getting filter", req, codes.Internal)
	}

	entities, err := s.stg.Entity().JoinList(ctx, req.Slug, req.Order, req.Sort, req.Limit, req.Offset, filter, req.Aggregate)
	if err != nil {
		return nil, helper.HandleError(s.log, err, "error while getting "+req.Slug, req, codes.Internal)
	}

	count, err := s.stg.Entity().JoinCount(ctx, req.Slug, filter)
	if err != nil {
		return nil, helper.HandleError(s.log, err, "error while getting count "+req.Slug, req, codes.Internal)
	}

	for _, item := range entities {
		item["created_at"], _ = helper.ToLocationTime(cast.ToString(item["created_at"]), req.Location)
		item["updated_at"], _ = helper.ToLocationTime(cast.ToString(item["updated_at"]), req.Location)

		entity, err := helper.ToProtoStruct(item)
		if err != nil {
			return nil, helper.HandleError(s.log, err, "error while converting to proto struct", entity, codes.Internal)
		}

		res.Entities = append(res.Entities, entity)
	}

	res.Count = int32(count)

	return res, nil
}
