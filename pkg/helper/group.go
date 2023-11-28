package helper

import (
	"fmt"
	pb "kassa360/kassa360_go_dynamic_service/genproto/dynamic_service"
	"kassa360/kassa360_go_dynamic_service/models"
	"reflect"

	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CheckDataForPatch(obj interface{}, group *pb.Group, location *string) error {
	fieldMap := make(map[string]interface{}, 0)

	if group.GroupType == models.GroupTypeObject {
		var data map[string]interface{}
		switch obj.(type) {
		case map[string]interface{}:
			data = cast.ToStringMap(obj)
		default:
			return fmt.Errorf("invalid data type: %s", group.Slug)
		}

		for _, val := range group.Fields {
			fieldMap[val.Slug] = true
		}

		for _, val := range group.Children {
			fieldMap[val.Slug] = true
		}

		for k := range data {
			if _, ok := fieldMap[k]; !ok {
				return fmt.Errorf("invalid field: %s", group.Slug+"."+k)
			}
		}

		if err := CheckFieldForPatch(data, group, location); err != nil {
			return err
		}

		for i := range group.Fields {
			if val, ok := data[group.Fields[i].Slug]; ok && EmptyValue(&group.Fields[i].FieldType) == val {
				delete(data, group.Fields[i].Slug)
			}
		}
		for _, v := range group.Children {
			if err := CheckDataForPatch(data[v.Slug], v, location); err != nil {
				return fmt.Errorf("%s.%s", group.Slug, err.Error())
			}
		}
	} else if group.GroupType == models.GroupTypeArray {
		var arr []interface{}
		if reflect.TypeOf(obj).Kind() == reflect.Slice {
			arr = cast.ToSlice(obj)
		} else {
			return fmt.Errorf("invalid data type: %s", group.Slug)
		}

		for _, val := range group.Fields {
			fieldMap[val.Slug] = true
		}

		for _, val := range group.Children {
			fieldMap[val.Slug] = true
		}

		for _, v := range arr {
			var data map[string]interface{}

			switch v.(type) {
			case map[string]interface{}:
				data = cast.ToStringMap(v)
			default:
				return fmt.Errorf("invalid data type: %s", group.Slug)
			}

			for k := range data {
				if _, ok := fieldMap[k]; !ok {
					return fmt.Errorf("invalid field: %s", group.Slug+"."+k)
				}
			}

			for _, v := range group.Fields {
				field := group.Slug + "." + v.Slug

				if _, ok := data[v.Slug]; ok {
					if err := CheckOnlyOneField(data, v, location, field); err != nil {
						return err
					}
				}
			}

			for _, v := range group.Children {
				if err := CheckDataForPatch(data[v.Slug], v, location); err != nil {
					return fmt.Errorf("%s.%s", group.Slug, err.Error())
				}
			}
		}

		clear(arr)
	} else {
		return fmt.Errorf("invalid group type: %s", group.Slug)
	}

	clear(fieldMap)

	return nil
}

func CheckData(obj interface{}, group *pb.Group, location *string) error {
	fieldMap := make(map[string]interface{}, 0)

	if group.GroupType == models.GroupTypeObject {
		var data map[string]interface{}
		switch obj.(type) {
		case map[string]interface{}:
			data = cast.ToStringMap(obj)
		default:
			return fmt.Errorf("invalid data type: %s", group.Slug)
		}

		for _, val := range group.Fields {
			fieldMap[val.Slug] = true
		}

		for _, val := range group.Children {
			fieldMap[val.Slug] = true
		}

		for k := range data {
			if _, ok := fieldMap[k]; !ok {
				return fmt.Errorf("invalid field: %s", group.Slug+"."+k)
			}
		}

		if err := CheckField(data, group, location); err != nil {
			return err
		}

		for _, v := range group.Children {
			if err := CheckData(data[v.Slug], v, location); err != nil {
				return fmt.Errorf("%s.%s", group.Slug, err.Error())
			}
		}
	} else if group.GroupType == models.GroupTypeArray {
		var arr []interface{}
		if reflect.TypeOf(obj).Kind() == reflect.Slice {
			arr = cast.ToSlice(obj)
		} else {
			return fmt.Errorf("invalid data type: %s", group.Slug)
		}

		for _, val := range group.Fields {
			fieldMap[val.Slug] = true
		}

		for _, val := range group.Children {
			fieldMap[val.Slug] = true
		}

		for _, v := range arr {
			var data map[string]interface{}

			switch v.(type) {
			case map[string]interface{}:
				data = cast.ToStringMap(v)
			default:
				return fmt.Errorf("invalid data type: %s", group.Slug)
			}

			for k := range data {
				if _, ok := fieldMap[k]; !ok {
					return fmt.Errorf("invalid field: %s", group.Slug+"."+k)
				}
			}

			if err := CheckField(data, group, location); err != nil {
				return err
			}

			for _, v := range group.Children {
				if err := CheckData(data[v.Slug], v, location); err != nil {
					return fmt.Errorf("%s.%s", group.Slug, err.Error())
				}
			}
		}

		clear(arr)
	} else {
		return fmt.Errorf("invalid group type: %s", group.Slug)
	}

	clear(fieldMap)

	return nil
}

func CheckFieldForPatch(data map[string]interface{}, group *pb.Group, location *string) error {
	for _, v := range group.Fields {
		field := group.Slug + "." + v.Slug

		if _, ok := data[v.Slug]; ok {
			if err := CheckOnlyOneField(data, v, location, field); err != nil {
				return err
			}
		}
	}

	return nil
}

func CheckField(data map[string]interface{}, group *pb.Group, location *string) error {
	for _, v := range group.Fields {
		field := group.Slug + "." + v.Slug

		if err := CheckOnlyOneField(data, v, location, field); err != nil {
			return err
		}
	}

	return nil
}

func CheckOnlyOneField(data map[string]interface{}, v *pb.Field, location *string, field string) error {
	var err error

	if v.IsRequired {
		if _, ok := data[v.Slug]; !ok {
			return fmt.Errorf("%s is required (empty)", field)
		}

		if data[v.Slug] == EmptyValue(&v.FieldType) || data[v.Slug] == nil {
			data[v.Slug] = GetDefaultValue(&v.FieldType, location, v.DefaultValue)

			if data[v.Slug] == EmptyValue(&v.FieldType) {
				return fmt.Errorf("%s is required (empty)", field)
			}
		}

		if v.FieldType == models.FieldTypeNumber {
			num := cast.ToInt32(data[v.Slug])

			if v.Max != 0 || v.Min != 0 {
				if num > v.Max || num < v.Min {
					return fmt.Errorf("%s is not valid (min, max)", field)
				}
			}
		}

		if v.FieldType == models.FieldTypeFloat {
			num := cast.ToFloat64(data[v.Slug])

			if v.Max != 0 || v.Min != 0 {
				if num > float64(v.Max) || num < float64(v.Min) {
					return fmt.Errorf("%s is not valid (min, max)", field)
				}
			}
		}

		if v.FieldType == models.FieldTypeText {
			val := cast.ToString(data[v.Slug])

			if v.Max != 0 || v.Min != 0 {
				if int32(len(val)) > v.Max || int32(len(val)) < v.Min {
					return fmt.Errorf("%s is not valid (min, max)", field)
				}
			}
		}

		if v.ValidationRegex != "" {
			if !RegexValidate(v.ValidationRegex, data[v.Slug]) {
				return fmt.Errorf("%s is not valid with regex", field)
			}
		}

		if v.ValidationFunc != "" {
			if !CheckVariable(data[v.Slug], v) {
				return fmt.Errorf("%s is not valid", field)
			}
		}
	} else {
		if obj, ok := data[v.Slug]; !ok || obj == nil || obj == "" || obj == 0 || obj == EmptyValue(&v.FieldType) {
			data[v.Slug] = GetDefaultValue(&v.FieldType, location, v.DefaultValue)
		}
	}

	if v.FieldType == models.FieldTypeObjectID {
		data[v.Slug], err = ToObjectID(data[v.Slug])
		if err != nil {
			return fmt.Errorf("%s is not valid", field)
		}
	}

	return nil
}

func GetDefaultValue(fieldType, location *string, defaultValue string) interface{} {
	number := cast.ToInt32(defaultValue)
	numberFloat := cast.ToFloat64(defaultValue)
	boolean := cast.ToBool(defaultValue)

	if defaultValue == "null" {
		return nil
	}

	switch *fieldType {
	case models.FieldTypeNumber:
		return number

	case models.FieldTypeFloat:
		return numberFloat

	case models.FieldTypeText:
		return defaultValue

	case models.FieldTypeBool:
		return boolean

	case models.FieldTypeDate:
		if defaultValue == "TIME_NOW" {
			return DateNowWithLocation(*location)
		}
		return defaultValue

	case models.FieldTypeDateTime:
		if defaultValue == "TIME_NOW" {
			return TimeNowWithLocation(*location)
		}
		return defaultValue

	case models.FieldTypeRadio, models.FieldTypeSelect, models.FieldTypeFile:
		return defaultValue

	case models.FieldTypeObjectID:
		if defaultValue == "GENERATE" {
			return GenerateID()
		}

		return defaultValue

	case models.FieldTypeUuid:
		if defaultValue == "GENERATE" {
			return GenerateUUID()
		}

		return defaultValue

	default:
		return nil
	}
}

func EmptyValue(filedType *string) interface{} {
	switch *filedType {
	case models.FieldTypeNumber, models.FieldTypeFloat:
		return 0

	case models.FieldTypeText:
		return ""

	case models.FieldTypeBool:
		return false

	case models.FieldTypeDate:
		return ""

	case models.FieldTypeDateTime:
		return ""

	case models.FieldTypeObjectID, models.FieldTypeUuid:
		return ""

	case models.FieldTypeRadio, models.FieldTypeSelect, models.FieldTypeFile:
		return ""

	default:
		return nil
	}
}

func ToObjectID(id interface{}) (primitive.ObjectID, error) {
	switch s := id.(type) {
	case primitive.ObjectID:
		return s, nil
	case string:
		return primitive.ObjectIDFromHex(s)
	case []byte:
		return primitive.ObjectIDFromHex(string(s))
	case []rune:
		return primitive.ObjectIDFromHex(string(s))
	default:
		return primitive.NilObjectID, fmt.Errorf("invalid id type")
	}
}
