package helper

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/mirobidjon/go_dynamic_service/config"
	pb "github.com/mirobidjon/go_dynamic_service/genproto/dynamic_service"
	"github.com/mirobidjon/go_dynamic_service/model"

	"github.com/google/uuid"
	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var funcMap = map[string]interface{}{
	"phone":    IsValidPhone,
	"email":    IsValidEmail,
	"uuid":     IsValidUUID,
	"time":     IsValidTime,
	"date":     IsValidDate,
	"objectId": IsValidObjectId,
}

func CheckValidationFunction(name string) bool {
	if _, ok := funcMap[name]; !ok {
		return false
	}

	return true
}

func ValidateRegex(regex string) error {
	if regex == "" {
		return nil
	}

	_, err := regexp.Compile(regex)
	if err != nil {
		return err
	}

	return nil
}

func RegexValidate(regex string, value interface{}) bool {
	if value == nil {
		return false
	}

	if regex == "" {
		return true
	}

	if cast.ToString(value) == "" {
		return false
	}

	r := regexp.MustCompile(regex)
	return r.MatchString(cast.ToString(value))
}

func CheckValidationFunc(funcName string, value interface{}, valueType string) bool {
	switch value.(type) {
	case int, int32, int64, uint, uint32, uint64:
		if valueType != model.FieldTypeNumber {
			return false
		}

		if cast.ToInt(value) == 0 {
			return false
		}
	case string:
		if valueType != model.FieldTypeText && valueType != model.FieldTypeDate &&
			valueType != model.FieldTypeDateTime && valueType != model.FieldTypeSelect &&
			valueType != model.FieldTypeRadio && valueType != model.FieldTypeObjectID &&
			valueType != model.FieldTypeUuid {
			return false
		}

	case primitive.ObjectID:
		if valueType != model.FieldTypeObjectID {
			return false
		}

	case uuid.UUID:
		if valueType != model.FieldTypeUuid {
			return false
		}

	default:
		return false
	}

	for name, i := strings.Split(funcName, ","), 0; i < len(name); i++ {
		if _, ok := funcMap[name[i]]; !ok {
			return false
		}

		if value == nil {
			return false
		}

		if ok := funcMap[name[i]].(func(interface{}) bool)(value); !ok {
			return false
		}
	}

	return true
}

func CheckVariable(value interface{}, field *pb.Field) bool {
	if field.ValidationFunc == "" {
		return true
	}

	return CheckValidationFunc(field.ValidationFunc, value, field.FieldType)
}

func IsValidPhone(phone interface{}) bool {
	phoneStr := cast.ToString(phone)
	r := regexp.MustCompile(`^\+998[0-9]{2}[0-9]{7}$`)
	return r.MatchString(phoneStr)
}

func IsValidEmail(email interface{}) bool {
	emailStr := cast.ToString(email)
	r := regexp.MustCompile(`^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`)
	return r.MatchString(emailStr)
}

func IsValidUUID(id any) bool {
	switch s := id.(type) {
	case string, []byte, []rune:
		uuidStr := cast.ToString(id)
		r := regexp.MustCompile(`^[a-f\d]{8}(-[a-f\d]{4}){4}[a-f\d]{8}$`)
		return r.MatchString(uuidStr)

	case uuid.UUID:
		return s.String() != ""

	default:
		return false
	}
}

func IsValidTime(value any) bool {
	valueStr := cast.ToString(value)
	_, err := time.Parse(config.TimeStampLayout, valueStr)
	return err == nil
}

func IsValidDate(value any) bool {
	valueStr := cast.ToString(value)
	_, err := time.Parse(config.DateLayout, valueStr)
	return err == nil
}

func IsValidObjectId(value any) bool {
	ok := false
	defer func() {
		fmt.Println("IsValidObjectId", value, "ok", ok)
	}()

	switch s := value.(type) {
	case string, []byte, []rune:
		valueStr := cast.ToString(value)
		// r := regexp.MustCompile(`^[0-9a-fA-F]{24}$`)
		// return r.MatchString(valueStr)
		_, err := primitive.ObjectIDFromHex(valueStr)
		ok = err == nil
		return ok

	case primitive.ObjectID:
		ok = s.Hex() != ""
		return ok

	default:
		return ok
	}
}

func IsValidGeoPoint(value any) bool {
	valueStr, err := json.Marshal(value)
	if err != nil {
		return false
	}

	var point model.GeoPoint
	if err := json.Unmarshal(valueStr, &point); err != nil {
		return false
	}

	if point.Type != "Point" {
		return false
	}

	if len(point.Coordinates) != 2 {
		return false
	}

	return true
}

func IsValidGeoPolygon(value any) bool {
	valueStr, err := json.Marshal(value)
	if err != nil {
		return false
	}

	var polygon model.GeoPolygon
	if err := json.Unmarshal(valueStr, &polygon); err != nil {
		return false
	}

	if polygon.Type != "Polygon" {
		return false
	}

	if len(polygon.Coordinates) < 3 {
		return false
	}

	if polygon.Coordinates[0][0] != polygon.Coordinates[len(polygon.Coordinates)-1][0] || polygon.Coordinates[0][1] != polygon.Coordinates[len(polygon.Coordinates)-1][1] {
		return false
	}

	return true
}
