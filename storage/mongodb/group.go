package mongodb

import (
	"context"
	"errors"
	"fmt"
	"strings"

	pd "github.com/mirobidjon/go_dynamic_service/genproto/dynamic_service"
	"github.com/mirobidjon/go_dynamic_service/pkg/helper"
	"github.com/mirobidjon/go_dynamic_service/storage"

	"github.com/jellydator/ttlcache/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type groupRepo struct {
	db    *mongo.Database
	cache *ttlcache.Cache[string, string]
}

func NewGroupRepo(db *mongo.Database, cache *ttlcache.Cache[string, string]) storage.GroupI {
	return &groupRepo{
		db:    db,
		cache: cache,
	}
}

func (r *groupRepo) CreateGroup(ctx context.Context, req *pd.Group) error {
	collection := r.db.Collection("group")

	id := primitive.NewObjectID()

	req.XId = id.Hex()
	req.CreatedAt = helper.TimeNow()
	req.UpdatedAt = helper.TimeNow()

	group := bson.M{
		"_id":          id,
		"name":         req.Name,
		"slug":         req.Slug,
		"description":  req.Description,
		"status":       req.Status,
		"parent_id":    req.ParentId,
		"group_type":   req.GroupType,
		"created_at":   req.CreatedAt,
		"updated_at":   req.UpdatedAt,
		"order_number": req.OrderNumber,
	}

	_, err := collection.InsertOne(ctx, group)

	r.cache.DeleteAll()

	return err
}

func (r *groupRepo) GetGroupById(ctx context.Context, req *pd.GetByIdRequest) (*pd.Group, error) {
	collection := r.db.Collection("group")

	var (
		key   = "_id"
		value interface{}
		group pd.Group
	)

	id, err := primitive.ObjectIDFromHex(req.XId)
	if err != nil {
		key = "slug"
		value = req.XId
	} else {
		value = id
	}

	err = collection.FindOne(ctx, bson.M{key: value}).Decode(&group)
	return &group, err
}

func (r *groupRepo) DeleteGroup(ctx context.Context, req *pd.GetByIdRequest) error {
	collection := r.db.Collection("group")

	var ids []primitive.ObjectID

	arr := strings.Split(req.XId, ",")

	for _, v := range arr {
		id, err := primitive.ObjectIDFromHex(v)
		if err != nil {
			return err
		}
		ids = append(ids, id)
	}

	if len(ids) == 0 {
		return errors.New("id is empty")
	}

	_, err := collection.DeleteMany(ctx, bson.M{"_id": bson.M{"$in": ids}})

	return err
}

func (r *groupRepo) UpdateGroup(ctx context.Context, req *pd.Group) error {
	collection := r.db.Collection("group")

	var (
		id     primitive.ObjectID
		err    error
		upsert = true
	)

	req.UpdatedAt = helper.TimeNow()

	if req.XId != "" {
		id, err = primitive.ObjectIDFromHex(req.XId)
		if err != nil {
			return err
		}
	} else {
		id = primitive.NewObjectID()
	}

	group := bson.M{
		"_id":          id,
		"name":         req.Name,
		"slug":         req.Slug,
		"description":  req.Description,
		"status":       req.Status,
		"parent_id":    req.ParentId,
		"group_type":   req.GroupType,
		"updated_at":   req.UpdatedAt,
		"order_number": req.OrderNumber,
	}

	_, err = collection.UpdateOne(ctx,
		bson.M{"_id": id},
		bson.M{"$set": group},
		&options.UpdateOptions{Upsert: &upsert},
	)

	r.cache.DeleteAll()

	return err
}

func (r *groupRepo) GetAllGroup(ctx context.Context, req *pd.GetAllGroupRequest) (*pd.GetAllGroupResponse, error) {
	collection := r.db.Collection("group")

	var (
		groups []*pd.Group
		filter bson.D
		opts   = options.Find()
	)

	opts.SetLimit(int64(req.Limit))

	if req.Offset > 0 {
		opts.SetSkip(int64(req.Offset))
	}

	if req.ParentId != "" {
		opts.SetSort(bson.M{
			"order_number": 1,
		})
	} else {
		opts.SetSort(bson.M{
			"created_at": -1,
		})
	}

	if req.ParentId != "" {
		filter = append(filter, bson.E{Key: "parent_id", Value: req.ParentId})
	} else {
		filter = append(filter, bson.E{
			Key:   "parent_id",
			Value: primitive.Null{},
		})
	}

	if req.Slug != "" {
		filter = append(filter, bson.E{Key: "slug", Value: req.Slug})
	}

	if req.Search != "" {
		filter = append(filter, bson.E{Key: "name", Value: primitive.Regex{Pattern: req.Search, Options: "i"}})
	}

	if req.GroupType != 0 {
		filter = append(filter, bson.E{Key: "group_type", Value: req.GroupType})
	}

	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &groups); err != nil {
		return nil, err
	}

	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &pd.GetAllGroupResponse{
		Groups: groups,
		Count:  int32(count),
	}, nil
}

func (r *groupRepo) CreateField(ctx context.Context, req *pd.Field) error {
	collection := r.db.Collection("field")

	id := primitive.NewObjectID()
	req.CreatedAt = helper.TimeNow()
	req.UpdatedAt = helper.TimeNow()
	req.XId = id.Hex()

	field := bson.M{
		"_id":              id,
		"name":             req.Name,
		"slug":             req.Slug,
		"description":      req.Description,
		"status":           req.Status,
		"created_at":       req.CreatedAt,
		"updated_at":       req.UpdatedAt,
		"group_id":         req.GroupId,
		"order_number":     req.OrderNumber,
		"field_type":       req.FieldType,
		"placeholder":      req.Placeholder,
		"is_required":      req.IsRequired,
		"select_type":      req.SelectType,
		"validation_regex": req.ValidationRegex,
		"validation_func":  req.ValidationFunc,
		"min":              req.Min,
		"max":              req.Max,
		"default_value":    req.DefaultValue,
		"is_searchable":    req.IsSearchable,
	}

	_, err := collection.InsertOne(ctx, field)

	r.cache.DeleteAll()

	return err
}

func (r *groupRepo) GetFieldById(ctx context.Context, req *pd.GetByIdRequest) (*pd.Field, error) {
	collection := r.db.Collection("field")

	id, err := primitive.ObjectIDFromHex(req.XId)
	if err != nil {
		return nil, err
	}

	var field pd.Field
	err = collection.FindOne(ctx, bson.M{"_id": id}).Decode(&field)
	return &field, err
}

func (r *groupRepo) DeleteField(ctx context.Context, req *pd.GetByIdRequest) error {
	var ids []primitive.ObjectID
	collection := r.db.Collection("field")

	arr := strings.Split(req.XId, ",")

	if len(arr) == 0 {
		return errors.New("id is empty")
	}

	for _, v := range arr {
		id, err := primitive.ObjectIDFromHex(v)
		if err != nil {
			return err
		}
		ids = append(ids, id)
	}

	_, err := collection.DeleteMany(ctx, bson.M{"_id": bson.M{"$in": ids}})

	r.cache.DeleteAll()

	return err
}

func (r *groupRepo) UpdateField(ctx context.Context, req *pd.Field) error {
	collection := r.db.Collection("field")

	var (
		id     primitive.ObjectID
		err    error
		upsert = true
	)

	req.UpdatedAt = helper.TimeNow()

	if req.XId != "" {
		id, err = primitive.ObjectIDFromHex(req.XId)
		if err != nil {
			return err
		}
	} else {
		id = primitive.NewObjectID()
	}

	field := bson.M{
		"_id":              id,
		"name":             req.Name,
		"slug":             req.Slug,
		"description":      req.Description,
		"status":           req.Status,
		"created_at":       req.CreatedAt,
		"updated_at":       req.UpdatedAt,
		"group_id":         req.GroupId,
		"order_number":     req.OrderNumber,
		"field_type":       req.FieldType,
		"placeholder":      req.Placeholder,
		"is_required":      req.IsRequired,
		"select_type":      req.SelectType,
		"validation_regex": req.ValidationRegex,
		"validation_func":  req.ValidationFunc,
		"min":              req.Min,
		"max":              req.Max,
		"default_value":    req.DefaultValue,
		"is_searchable":    req.IsSearchable,
	}

	_, err = collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		&bson.M{"$set": field},
		&options.UpdateOptions{Upsert: &upsert},
	)

	r.cache.DeleteAll()

	return err
}

func (r *groupRepo) GetAllField(ctx context.Context, req *pd.GetAllFieldRequest) (*pd.GetAllFieldResponse, error) {
	collection := r.db.Collection("field")

	var (
		fields []*pd.Field
		opts   = options.Find()
		filter = bson.D{}
	)

	opts.SetLimit(int64(req.Limit))
	if req.Offset > 0 {
		opts.SetSkip(int64(req.Offset))
	}

	opts.SetSort(bson.D{
		bson.E{Key: "order_number", Value: 1},
		bson.E{Key: "created_at", Value: -1},
	})

	if req.GroupId != "" {
		filter = append(filter, bson.E{Key: "group_id", Value: req.GroupId})
	}

	if req.Search != "" {
		filter = append(filter, bson.E{Key: "$or", Value: bson.A{
			bson.D{bson.E{Key: "name", Value: primitive.Regex{Pattern: req.Search, Options: "i"}}},
			bson.D{bson.E{Key: "slug", Value: primitive.Regex{Pattern: req.Search, Options: "i"}}},
			bson.D{bson.E{Key: "placeholder", Value: primitive.Regex{Pattern: req.Search, Options: "i"}}},
		},
		})
	}

	if req.FieldType != "" {
		filter = append(filter, bson.E{Key: "field_type", Value: req.FieldType})
	}

	if req.Slug != "" {
		filter = append(filter, bson.E{Key: "slug", Value: req.Slug})
	}

	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &fields); err != nil {
		return nil, err
	}

	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &pd.GetAllFieldResponse{
		Fields: fields,
		Count:  int32(count),
	}, nil
}

func (r *groupRepo) GetFullGroup(ctx context.Context, req *pd.GetByIdRequest) (*pd.Group, error) {

	item := r.cache.Get(req.XId)
	if item != nil {
		var group = pd.Group{}
		err := helper.StringToProto(&group, item.Value())
		if err != nil {
			return nil, fmt.Errorf("error while converting string to group" + err.Error())
		}

		return &group, nil
	}

	var (
		groups          []*pd.Group
		filter          bson.D
		groupIds        []string
		groupsMap       = make(map[string][]*pd.Group)
		group           *pd.Group
		groupCollection = r.db.Collection("group")
		fieldCollection = r.db.Collection("field")
		key             = "_id"
		value           interface{}
		fieldsMap       = make(map[string][]*pd.Field)
	)

	if req.XId == "" {
		return nil, fmt.Errorf("ID or Slug is required")
	}

	id, err := primitive.ObjectIDFromHex(req.XId)
	if err != nil {
		key = "slug"
		value = req.XId
	} else {
		value = id
	}

	filter = bson.D{
		{Key: "$or",
			Value: bson.A{
				bson.D{{Key: key, Value: value}},
				bson.D{
					{Key: "parent_id",
						Value: bson.D{
							{Key: "$nin",
								Value: bson.A{
									primitive.Null{},
								},
							},
						},
					},
				},
			},
		},
	}

	cursor, err := groupCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &groups); err != nil {
		return nil, err
	}

	if len(groups) == 0 {
		return nil, fmt.Errorf("invalid ID or Slug")
	}

	for _, g := range groups {
		groupIds = append(groupIds, g.XId)

		if g.ParentId != nil {
			groupsMap[g.ParentId.GetValue()] = append(groupsMap[g.ParentId.GetValue()], g)
		} else {
			group = g
		}
	}

	if group == nil {
		return nil, fmt.Errorf("invalid ID or Slug")
	}

	var fields []*pd.Field
	filter = bson.D{
		{Key: "group_id", Value: bson.D{{Key: "$in", Value: groupIds}}},
	}

	cursor, err = fieldCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &fields); err != nil {
		return nil, err
	}

	for _, f := range fields {
		fieldsMap[f.GroupId] = append(fieldsMap[f.GroupId], f)
	}

	group.Children = CollectGroup(groupsMap, fieldsMap, group.XId)
	group.Fields = fieldsMap[group.XId]

	clear(groupsMap)
	clear(fieldsMap)
	clear(groupIds)
	clear(groups)
	clear(fields)

	groupString, err := helper.ProtoToString(group)
	if err != nil {
		return nil, fmt.Errorf("error while converting group to string" + err.Error())
	}

	r.cache.Set(req.XId, groupString, ttlcache.DefaultTTL)

	return group, nil
}

func CollectGroup(groupMap map[string][]*pd.Group, fieldsMap map[string][]*pd.Field, parent_id string) []*pd.Group {
	groups, ok := groupMap[parent_id]

	if !ok {
		return nil
	}

	for _, g := range groups {
		g.Children = CollectGroup(groupMap, fieldsMap, g.XId)
		g.Fields = fieldsMap[g.XId]
	}

	return groups
}
