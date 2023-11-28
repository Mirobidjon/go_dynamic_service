package helper

import (
	"fmt"
	"kassa360/kassa360_go_dynamic_service/config"
	pb "kassa360/kassa360_go_dynamic_service/genproto/dynamic_service"
	"kassa360/kassa360_go_dynamic_service/models"
	"regexp"
	"strings"
	"time"

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
		if valueType != models.FieldTypeNumber {
			return false
		}

		if cast.ToInt(value) == 0 {
			return false
		}
	case string:
		if valueType != models.FieldTypeText && valueType != models.FieldTypeDate &&
			valueType != models.FieldTypeDateTime && valueType != models.FieldTypeSelect &&
			valueType != models.FieldTypeRadio && valueType != models.FieldTypeObjectID &&
			valueType != models.FieldTypeUuid {
			return false
		}

	case primitive.ObjectID:
		if valueType != models.FieldTypeObjectID {
			return false
		}

	case uuid.UUID:
		if valueType != models.FieldTypeUuid {
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

func IsValidUUID(id interface{}) bool {
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

func IsValidTime(value interface{}) bool {
	valueStr := cast.ToString(value)
	_, err := time.Parse(config.TimeStampLayout, valueStr)
	return err == nil
}

func IsValidDate(value interface{}) bool {
	valueStr := cast.ToString(value)
	_, err := time.Parse(config.DateLayout, valueStr)
	return err == nil
}

func IsValidObjectId(value interface{}) bool {
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
