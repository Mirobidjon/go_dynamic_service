package mongodb

import (
	"context"
	"fmt"
	"strings"

	pd "github.com/mirobidjon/go_dynamic_service/genproto/dynamic_service"
	"github.com/mirobidjon/go_dynamic_service/models"
	"github.com/mirobidjon/go_dynamic_service/pkg/helper"
	"github.com/mirobidjon/go_dynamic_service/storage"

	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type entityRepo struct {
	db          *mongo.Database
	collections map[string]*mongo.Collection
}

func NewEntityRepo(db *mongo.Database) storage.EntityI {
	return &entityRepo{
		db:          db,
		collections: make(map[string]*mongo.Collection),
	}
}

func (r *entityRepo) Disconnect() error {
	return r.db.Client().Disconnect(context.Background())
}

func (r *entityRepo) getCollection(slug string) *mongo.Collection {
	if r.collections[slug] == nil {
		r.collections[slug] = r.db.Collection(slug)
	}

	return r.collections[slug]
}

func (r *entityRepo) Create(ctx context.Context, slug string, body map[string]interface{}) error {
	col := r.getCollection(slug)

	body["created_at"] = helper.TimeNow()
	body["updated_at"] = helper.TimeNow()

	_, err := col.InsertOne(ctx, body)
	return err
}

func (r *entityRepo) Update(ctx context.Context, slug string, id string, body map[string]interface{}) error {
	col := r.getCollection(slug)

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid id format")
	}

	body["_id"] = objId
	body["updated_at"] = helper.TimeNow()

	delete(body, "created_at")

	_, err = col.UpdateOne(ctx, bson.M{"_id": objId}, &bson.M{"$set": body})
	return err
}

func (r *entityRepo) Delete(ctx context.Context, slug string, id string) error {
	col := r.getCollection(slug)

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid id format")
	}

	_, err = col.DeleteOne(ctx, bson.M{"_id": objId})
	return err
}

func (r *entityRepo) Get(ctx context.Context, slug string, id string) (map[string]interface{}, error) {
	col := r.getCollection(slug)

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id format")
	}

	var result map[string]interface{}
	err = col.FindOne(ctx, bson.M{"_id": objId}).Decode(&result)
	return result, err
}

func (r *entityRepo) List(ctx context.Context, slug, order, sort string, limit, offset int32, filter bson.D) ([]map[string]interface{}, error) {
	var (
		col     = r.getCollection(slug)
		opts    = options.Find()
		orderBy int32
	)

	if limit <= 0 {
		limit = 10
	}

	if offset < 0 {
		offset = 0
	}

	if order == "" {
		order = "desc"
	}

	if order == "desc" {
		orderBy = -1
	} else {
		orderBy = 1
	}

	if sort == "" {
		sort = "created_at"
	}

	opts.SetLimit(int64(limit))
	opts.SetSkip(int64(offset))
	opts.SetSort(bson.M{
		sort: orderBy,
	})

	fmt.Println(filter)

	var result []map[string]interface{}
	cur, err := col.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var doc map[string]interface{}
		err := cur.Decode(&doc)
		if err != nil {
			return nil, err
		}

		result = append(result, doc)
	}

	return result, nil
}

func (r *entityRepo) JoinList(ctx context.Context, slug, order, sort string, limit, offset int32, filter bson.A, aggregate *pd.Aggregate) ([]map[string]interface{}, error) {
	var (
		col     = r.getCollection(slug)
		orderBy int32
	)

	if limit <= 0 {
		limit = 10
	}

	if order == "" {
		order = "desc"
	}

	if order == "desc" {
		orderBy = -1
	} else {
		orderBy = 1
	}

	if sort == "" && aggregate.GetGroup() == nil {
		sort = "created_at"
	}

	if aggregate.GetGroup() == nil {
		filter = append(filter, bson.M{
			"$limit": limit,
		})

		if offset > 0 {
			filter = append(filter, bson.M{
				"$skip": offset,
			})
		}

		filter = append(filter, bson.M{
			"$sort": bson.M{
				sort: orderBy,
			},
		})
	}

	for _, lookup := range aggregate.GetLookups() {
		filter = append(filter, bson.M{
			"$lookup": bson.M{
				"from":         lookup.From,
				"localField":   lookup.LocalField,
				"foreignField": lookup.ForeignField,
				"as":           lookup.As,
			},
		})
	}

	if aggregate.GetGroup() != nil {
		// field ga hohlagan param berish mumkin, lekin accumulatorga faqat sum, avg, min, max, push, first, last, expression ga esa table da bo'lgan field nomini berish kerak
		filter = append(filter, bson.M{
			"$group": bson.M{
				"_id": "$" + aggregate.Group.XId,
				aggregate.Group.Field: bson.M{
					"$" + aggregate.Group.Accumulator: "$" + aggregate.Group.Expression,
				},
			},
		})
	}

	// fmt.Println(filter)

	var result []map[string]interface{}
	cur, err := col.Aggregate(ctx, filter)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var doc map[string]interface{}
		err := cur.Decode(&doc)
		if err != nil {
			return nil, err
		}

		result = append(result, doc)
	}

	return result, nil
}

func (r *entityRepo) Count(ctx context.Context, slug string, filter bson.D) (int64, error) {
	col := r.getCollection(slug)

	return col.CountDocuments(ctx, filter)
}

func (r *entityRepo) JoinCount(ctx context.Context, slug string, filter bson.A) (int64, error) {
	col := r.getCollection(slug)

	filter = append(filter, bson.M{
		"$count": "count",
	})

	var result []map[string]interface{}
	cur, err := col.Aggregate(ctx, filter)
	if err != nil {
		return 0, err
	}

	for cur.Next(ctx) {
		var doc map[string]interface{}
		err := cur.Decode(&doc)
		if err != nil {
			return 0, err
		}

		result = append(result, doc)
	}

	if len(result) == 0 {
		return 0, nil
	}

	return cast.ToInt64(result[0]["count"]), nil
}

func (r *entityRepo) QueryFilter(req map[string]interface{}, group *pd.Group, search, location string) (bson.D, error) {
	filter := bson.D{}

	fromDate := cast.ToString(req["from_date"])
	if fromDate != "" {
		guid, err := helper.GenerateIdWithTime(fromDate, location)
		if err != nil {
			return nil, fmt.Errorf("%s Invalid from date format (yyyy-mm-dd hh24:mm:ss): %s", fromDate, err.Error())
		}

		filter = append(filter, bson.E{Key: "_id", Value: bson.D{{Key: "$gte", Value: guid}}})
	}

	toDate := cast.ToString(req["to_date"])
	if toDate != "" {
		guid, err := helper.GenerateIdWithTime(toDate, location)
		if err != nil {
			return nil, fmt.Errorf("%s Invalid to date format (yyyy-mm-dd hh24:mm:ss): %s", toDate, err.Error())
		}

		filter = append(filter, bson.E{Key: "_id", Value: bson.D{{Key: "$lte", Value: guid}}})
	}

	filter = makeQueryFilter(req, group, filter, "", search, location)

	return filter, nil
}

func (r *entityRepo) JoinQueryFilter(req map[string]interface{}, group *pd.Group, search, location string) (bson.A, error) {
	filter := bson.A{}

	fromDate := cast.ToString(req["from_date"])
	if fromDate != "" {
		guid, err := helper.GenerateIdWithTime(fromDate, location)
		if err != nil {
			return nil, fmt.Errorf("%s Invalid from date format (yyyy-mm-dd hh24:mm:ss): %s", fromDate, err.Error())
		}

		filter = append(filter, bson.M{
			"$match": bson.M{
				"_id": bson.M{
					"$gte": guid,
				},
			},
		})
	}

	toDate := cast.ToString(req["to_date"])
	if toDate != "" {
		guid, err := helper.GenerateIdWithTime(toDate, location)
		if err != nil {
			return nil, fmt.Errorf("%s Invalid to date format (yyyy-mm-dd hh24:mm:ss): %s", toDate, err.Error())
		}

		filter = append(filter, bson.M{
			"$match": bson.M{
				"_id": bson.M{
					"$lte": guid,
				},
			},
		})
	}

	filter = makeJoinQueryFilter(req, group, filter, "", search, location)

	return filter, nil
}

func makeJoinQueryFilter(req map[string]interface{}, group *pd.Group, filter bson.A, slug, search, location string) bson.A {
	for _, f := range group.Fields {
		slugStr := ""
		if slug != "" {
			slugStr = slug + "." + f.Slug
		} else {
			slugStr = f.Slug
		}

		val, ok := req[slugStr]
		if !ok {
			continue
		}

		if val == "null" {
			filter = append(filter, bson.M{
				"$match": bson.M{
					slugStr: primitive.Null{},
				},
			})
		}

		if f.IsSearchable == 1 && f.FieldType != models.FieldTypeNumber && f.FieldType != models.FieldTypeFloat {
			if cast.ToString(val) == "on_search" {
				filter = append(filter, bson.M{
					"$match": bson.M{
						slugStr: bson.M{
							"$regex":   search,
							"$options": "i",
						},
					},
				})
			} else {
				filter = append(filter, bson.M{
					"$match": bson.M{
						slugStr: bson.M{
							"$regex":   cast.ToString(val),
							"$options": "i",
						},
					},
				})
			}
		}

		if f.IsSearchable == 2 {
			if f.FieldType == models.FieldTypeDate || f.FieldType == models.FieldTypeDateTime {
				val, _ = helper.ToUTC(cast.ToString(val), location)
			}

			if slug == "" && f.Slug == "_id" {
				objIds := strings.Split(cast.ToString(val), ",")
				var objIdsHex []primitive.ObjectID
				for _, objId := range objIds {
					objIdHex, err := primitive.ObjectIDFromHex(objId)
					if err != nil {
						continue
					}

					objIdsHex = append(objIdsHex, objIdHex)
				}

				if len(objIdsHex) > 0 {
					filter = append(filter, bson.M{
						"$match": bson.M{
							"_id": bson.M{
								"$in": objIdsHex,
							},
						},
					})
				}
			} else {
				if f.FieldType == models.FieldTypeObjectID {
					val, _ = helper.ToObjectID(val)
				}

				filter = append(filter, bson.M{
					"$match": bson.M{
						slugStr: val,
					},
				})
			}
		}

		if f.FieldType == models.FieldTypeNumber || f.FieldType == models.FieldTypeFloat {
			numberVal, ok := val.(models.Pair)
			if ok {
				// TODO: add validation to the key
				// if !helper.Check
				filter = append(filter, bson.M{
					"$match": bson.M{
						slugStr: bson.M{
							numberVal.Operator: numberVal.Value,
						},
					},
				})
			}

			filter = append(filter, bson.M{
				"$match": bson.M{
					slugStr: val,
				},
			})
		}
	}

	for _, gr := range group.Children {
		slugStr := ""
		if slug != "" {
			slugStr = slug + "." + gr.Slug
		} else {
			slugStr = gr.Slug
		}

		filter = makeJoinQueryFilter(req, gr, filter, slugStr, search, location)
	}

	return filter
}

func makeQueryFilter(req map[string]interface{}, group *pd.Group, filter bson.D, slug, search, location string) bson.D {
	for _, f := range group.Fields {
		slugStr := ""
		if slug != "" {
			slugStr = slug + "." + f.Slug
		} else {
			slugStr = f.Slug
		}

		val, ok := req[slugStr]
		if !ok {
			continue
		}

		if val == "null" {
			filter = append(filter, bson.E{Key: slugStr, Value: primitive.Null{}})
		}

		if f.IsSearchable == 1 {
			if cast.ToString(val) == "on_search" {
				filter = append(filter, bson.E{Key: slugStr, Value: primitive.Regex{Pattern: search, Options: "i"}})
			} else {
				filter = append(filter, bson.E{Key: slugStr, Value: primitive.Regex{Pattern: cast.ToString(val), Options: "i"}})
			}
		}

		if f.IsSearchable == 2 {
			if f.FieldType == models.FieldTypeDate || f.FieldType == models.FieldTypeDateTime {
				val, _ = helper.ToUTC(cast.ToString(val), location)
			}

			if slug == "" && f.Slug == "_id" {
				objIds := strings.Split(cast.ToString(val), ",")
				var objIdsHex []primitive.ObjectID
				for _, objId := range objIds {
					objIdHex, err := primitive.ObjectIDFromHex(objId)
					if err != nil {
						continue
					}

					objIdsHex = append(objIdsHex, objIdHex)
				}

				if len(objIdsHex) > 0 {
					filter = append(filter, bson.E{Key: "_id", Value: bson.D{{Key: "$in", Value: objIdsHex}}})
				}
			} else {
				if f.FieldType == models.FieldTypeObjectID {
					val, _ = helper.ToObjectID(val)
				}

				filter = append(filter, bson.E{Key: slugStr, Value: val})
			}
		}

		if f.FieldType == models.FieldTypeNumber || f.FieldType == models.FieldTypeFloat {
			numberVal, ok := val.(models.Pair)
			if ok {
				// TODO: add validation to the key
				// if !helper.Check
				filter = append(filter, bson.E{Key: slugStr, Value: bson.D{{Key: numberVal.Operator, Value: numberVal.Value}}})
			}

			filter = append(filter, bson.E{Key: slugStr, Value: val})
		}
	}

	for _, g := range group.Children {
		slugStr := ""
		if slug != "" {
			slugStr = slug + "." + g.Slug
		} else {
			slugStr = g.Slug
		}

		filter = makeQueryFilter(req, g, filter, slugStr, search, location)
	}

	return filter
}
