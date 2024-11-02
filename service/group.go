package service

import (
	"context"

	"github.com/mirobidjon/go_dynamic_service/config"
	"github.com/mirobidjon/go_dynamic_service/model"
	"github.com/mirobidjon/go_dynamic_service/pkg/helper"
	"github.com/mirobidjon/go_dynamic_service/storage"

	pb "github.com/mirobidjon/go_dynamic_service/genproto/dynamic_service"

	log "github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

type groupService struct {
	log  log.LoggerI
	stg  storage.StorageI
	conf *config.Config
	pb.UnimplementedDynamicServiceServer
}

func NewGroupService(log log.LoggerI, storage storage.StorageI, conf *config.Config) *groupService {
	return &groupService{
		log:  log,
		stg:  storage,
		conf: conf,
	}
}

func (s *groupService) CreateGroup(ctx context.Context, req *pb.Group) (*pb.Group, error) {
	s.log.Info("CreateGroup", log.Any("req", req))

	if req.Slug == "" {
		return nil, helper.HandleError(s.log, nil, "Slug is required", req, codes.InvalidArgument)
	}

	if req.GroupType != 1 && req.GroupType != 2 {
		return nil, helper.HandleError(s.log, nil, "GroupType is required", req, codes.InvalidArgument)
	}

	if (req.Slug == "group" || req.Slug == "field") && req.GroupType != 1 && req.ParentId == nil {
		return nil, helper.HandleError(s.log, nil, "Slug is reserved", req, codes.InvalidArgument)
	}

	resp, err := s.stg.Group().GetAllGroup(ctx, &pb.GetAllGroupRequest{Limit: 2, ParentId: req.ParentId.GetValue(), Slug: req.Slug})
	if err != nil {
		return nil, helper.HandleError(s.log, err, "Error while getting group ", req, codes.Internal)
	}

	if len(resp.Groups) > 0 || resp.Count > 0 {
		return nil, helper.HandleError(s.log, nil, "Slug is already exists", req, codes.InvalidArgument)
	}

	err = s.stg.Group().CreateGroup(ctx, req)
	if err != nil {
		return nil, helper.HandleError(s.log, err, "Error while creating group ", req, codes.Internal)
	}

	return req, nil
}

func (s *groupService) UpdateGroup(ctx context.Context, req *pb.Group) (*pb.Group, error) {
	s.log.Info("UpdateGroup", log.Any("req", req))

	if req.Slug == "" {
		return nil, helper.HandleError(s.log, nil, "Slug is required", req, codes.InvalidArgument)
	}

	if req.GroupType != 1 && req.GroupType != 2 {
		return nil, helper.HandleError(s.log, nil, "GroupType is required", req, codes.InvalidArgument)
	}

	if (req.Slug == "group" || req.Slug == "field") && req.GroupType != 1 && req.ParentId == nil {
		return nil, helper.HandleError(s.log, nil, "Slug is reserved", req, codes.InvalidArgument)
	}

	resp, err := s.stg.Group().GetAllGroup(ctx, &pb.GetAllGroupRequest{Limit: 2, ParentId: req.ParentId.GetValue(), Slug: req.Slug})
	if err != nil {
		return nil, helper.HandleError(s.log, err, "Error while getting group ", req, codes.Internal)
	}

	if len(resp.Groups) > 1 || resp.Count > 1 {
		return nil, helper.HandleError(s.log, nil, "Slug is already exists", req, codes.InvalidArgument)
	}

	if len(resp.Groups) == 1 && resp.Groups[0].XId != req.XId {
		return nil, helper.HandleError(s.log, nil, "Slug is already exists", req, codes.InvalidArgument)
	}

	err = s.stg.Group().UpdateGroup(ctx, req)
	if err != nil {
		return nil, helper.HandleError(s.log, err, "Error while updating group ", req, codes.Internal)
	}

	return req, nil
}

func (s *groupService) DeleteGroup(ctx context.Context, req *pb.GetByIdRequest) (*emptypb.Empty, error) {
	s.log.Info("DeleteGroup", log.Any("req", req))

	err := s.stg.Group().DeleteGroup(ctx, req)
	if err != nil {
		return nil, helper.HandleError(s.log, err, "Error while deleting group ", req, codes.Internal)
	}

	return &emptypb.Empty{}, nil
}

func (s *groupService) GetGroupById(ctx context.Context, req *pb.GetByIdRequest) (*pb.Group, error) {
	s.log.Info("GetGroupById", log.Any("req", req))

	resp, err := s.stg.Group().GetGroupById(ctx, req)
	if err != nil {
		return nil, helper.HandleError(s.log, err, "Error while getting group by id: ", req, codes.Internal)
	}

	return resp, nil
}

func (s *groupService) GetAllGroup(ctx context.Context, req *pb.GetAllGroupRequest) (*pb.GetAllGroupResponse, error) {
	s.log.Info("GetAllGroup", log.Any("req", req))

	resp, err := s.stg.Group().GetAllGroup(ctx, req)
	if err != nil {
		return nil, helper.HandleError(s.log, err, "Error while getting all groups ", req, codes.Internal)
	}
	return resp, nil
}

func (s *groupService) CreateField(ctx context.Context, req *pb.Field) (*pb.Field, error) {
	s.log.Info("CreateField", log.Any("req", req))

	// check group exists
	_, err := s.stg.Group().GetGroupById(ctx, &pb.GetByIdRequest{XId: req.GroupId})
	if err != nil {
		return nil, helper.HandleError(s.log, err, "Error while getting group by id ", req, codes.Internal)
	}

	if req.Slug == "_id" && req.FieldType != model.FieldTypeObjectID {
		return nil, helper.HandleError(s.log, nil, "Slug _id is reserved", req, codes.InvalidArgument)
	}

	if req.ValidationFunc != "" {
		ok := helper.CheckValidationFunction(req.ValidationFunc)
		if !ok {
			return nil, helper.HandleError(s.log, err, "Invalid validation function ", req, codes.Internal)
		}
	}

	if err := helper.ValidateRegex(req.ValidationRegex); err != nil {
		return nil, helper.HandleError(s.log, err, "Invalid regex ", req, codes.Internal)
	}

	if req.Slug == "" {
		return nil, helper.HandleError(s.log, nil, "Slug is required", req, codes.InvalidArgument)
	}

	resp, err := s.stg.Group().GetAllField(ctx, &pb.GetAllFieldRequest{GroupId: req.GroupId, Slug: req.Slug})
	if err != nil || resp.Count > 0 {
		return nil, helper.HandleError(s.log, err, "slug is already exists ", req, codes.Internal)
	}

	err = s.stg.Group().CreateField(ctx, req)
	if err != nil {
		return nil, helper.HandleError(s.log, err, "Error while creating field ", req, codes.Internal)
	}

	return req, nil
}

func (s *groupService) UpdateField(ctx context.Context, req *pb.Field) (*pb.Field, error) {
	s.log.Info("UpdateField", log.Any("req", req))

	// check group exists
	_, err := s.stg.Group().GetGroupById(ctx, &pb.GetByIdRequest{XId: req.GroupId})
	if err != nil {
		return nil, helper.HandleError(s.log, err, "Error while getting group by id ", req, codes.Internal)
	}

	if req.Slug == "_id" && req.FieldType != model.FieldTypeObjectID {
		return nil, helper.HandleError(s.log, nil, "Slug _id is reserved", req, codes.InvalidArgument)
	}

	if req.ValidationFunc != "" {
		ok := helper.CheckValidationFunction(req.ValidationFunc)
		if !ok {
			return nil, helper.HandleError(s.log, err, "Invalid validation function ", req, codes.Internal)
		}
	}

	if err := helper.ValidateRegex(req.ValidationRegex); err != nil {
		return nil, helper.HandleError(s.log, err, "Invalid regex ", req, codes.Internal)
	}

	if req.Slug == "" {
		return nil, helper.HandleError(s.log, nil, "Slug is required", req, codes.InvalidArgument)
	}

	err = s.stg.Group().UpdateField(ctx, req)
	if err != nil {
		return nil, helper.HandleError(s.log, err, "Error while updating field ", req, codes.Internal)
	}

	return req, nil
}

func (s *groupService) DeleteField(ctx context.Context, req *pb.GetByIdRequest) (*emptypb.Empty, error) {
	s.log.Info("DeleteField", log.Any("req", req))

	err := s.stg.Group().DeleteField(ctx, req)
	if err != nil {
		return nil, helper.HandleError(s.log, err, "Error while deleting field ", req, codes.Internal)
	}

	return &emptypb.Empty{}, nil
}

func (s *groupService) GetFieldById(ctx context.Context, req *pb.GetByIdRequest) (*pb.Field, error) {
	s.log.Info("GetFieldById", log.Any("req", req))

	resp, err := s.stg.Group().GetFieldById(ctx, req)
	if err != nil {
		return nil, helper.HandleError(s.log, err, "Error while getting field by id ", req, codes.Internal)
	}

	return resp, nil
}

func (s *groupService) GetAllField(ctx context.Context, req *pb.GetAllFieldRequest) (*pb.GetAllFieldResponse, error) {
	s.log.Info("GetAllField", log.Any("req", req))

	resp, err := s.stg.Group().GetAllField(ctx, req)
	if err != nil {
		return nil, helper.HandleError(s.log, err, "Error while getting all fields ", req, codes.Internal)
	}

	return resp, nil
}

func (s *groupService) GetFullGroup(ctx context.Context, req *pb.GetByIdRequest) (*pb.Group, error) {
	s.log.Info("GetFullGroup", log.Any("req", req))

	resp, err := s.stg.Group().GetFullGroup(ctx, req)
	if err != nil {
		return nil, helper.HandleError(s.log, err, "Error while getting full group ", req, codes.Internal)
	}

	return resp, nil
}
