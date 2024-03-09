package helper

import (
	"time"

	"github.com/mirobidjon/go_dynamic_service/config"

	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	uuidGenerator = uuid.NewGenWithOptions(
		uuid.WithEpochFunc(func() time.Time {
			return time.Now().UTC().Add(5 * time.Hour)
		}),
	)
)

func GenerateID() primitive.ObjectID {
	return primitive.NewObjectIDFromTimestamp(time.Now().UTC())
}

func GenerateIdWithTime(date, location string) (primitive.ObjectID, error) {
	loc := getLocation(location)

	t, err := time.ParseInLocation(config.TimeStampLayout, date, loc)
	if err != nil {
		t, err = time.ParseInLocation(config.DateLayout, date, loc)
		if err != nil {
			return primitive.NewObjectID(), err
		}
	}

	t = t.In(time.UTC)

	return primitive.NewObjectIDFromTimestamp(t), nil
}

func GenerateUUID() string {
	id, _ := uuidGenerator.NewV7()
	return id.String()
}

func GenerateUUIDWithTime(date, location string) (string, error) {
	loc := getLocation(location)

	t, err := time.ParseInLocation(config.TimeStampLayout, date, loc)
	if err != nil {
		t, err = time.ParseInLocation(config.DateLayout, date, loc)
		if err != nil {
			return GenerateUUID(), err
		}
	}

	gen := uuid.NewGenWithOptions(
		uuid.WithEpochFunc(func() time.Time {
			return t
		}),
	)

	id, _ := gen.NewV7()
	return id.String(), nil
}
