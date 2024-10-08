package helper

import (
	"fmt"
	"reflect"

	pb "github.com/mirobidjon/go_dynamic_service/genproto/dynamic_service"
	"github.com/mirobidjon/go_dynamic_service/model"

	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CheckDataForPatch(obj interface{}, group *pb.Group, location *string) error {
	fieldMap := make(map[string]interface{}, 0)

	if group.GroupType == model.GroupTypeObject {
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
	} else if group.GroupType == model.GroupTypeArray {
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
					if err := CheckOnlyOneField(data, v, location, field, 0); err != nil {
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

	if group.GroupType == model.GroupTypeObject {
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
	} else if group.GroupType == model.GroupTypeArray {
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
			if err := CheckOnlyOneField(data, v, location, field, 0); err != nil {
				return err
			}
		}
	}

	return nil
}

func CheckField(data map[string]interface{}, group *pb.Group, location *string) error {
	for _, v := range group.Fields {
		field := group.Slug + "." + v.Slug

		if err := CheckOnlyOneField(data, v, location, field, 0); err != nil {
			return err
		}
	}

	return nil
}

func CheckOnlyOneField(data map[string]interface{}, v *pb.Field, location *string, field string, index int) error {
	var err error

	if v.IsArray && index < 1 {
		if reflect.TypeOf(data[v.Slug]).Kind() != reflect.Slice {
			return fmt.Errorf("%s is not valid (not array)", field)
		}

		arr := cast.ToSlice(data[v.Slug])

		for i := range arr {
			if err := CheckOnlyOneField(cast.ToStringMap(arr[i]), v, location, field, i+1); err != nil {
				return err
			}
		}

		clear(arr)
		return nil
	}

	if _, ok := data[v.Slug]; ok || v.IsRequired {
		if _, ok := data[v.Slug]; !ok {
			return fmt.Errorf("%s is required (empty)", field)
		}

		if data[v.Slug] == EmptyValue(&v.FieldType) || data[v.Slug] == nil {
			data[v.Slug] = GetDefaultValue(&v.FieldType, location, v.DefaultValue)

			if data[v.Slug] == EmptyValue(&v.FieldType) {
				return fmt.Errorf("%s is required (empty)", field)
			}
		}

		if v.FieldType == model.FieldTypePoint {
			if IsValidGeoPoint(data[v.Slug]) {
				return fmt.Errorf("%s is not valid (point)", field)
			}
		}

		if v.FieldType == model.FieldTypePolygon {
			if IsValidGeoPolygon(data[v.Slug]) {
				return fmt.Errorf("%s is not valid (polygon)", field)
			}
		}

		if v.FieldType == model.FieldTypeNumber {
			num := cast.ToInt32(data[v.Slug])

			if v.Max != 0 || v.Min != 0 {
				if num > v.Max || num < v.Min {
					return fmt.Errorf("%s is not valid (min, max)", field)
				}
			}
		}

		if v.FieldType == model.FieldTypeFloat {
			num := cast.ToFloat64(data[v.Slug])

			if v.Max != 0 || v.Min != 0 {
				if num > float64(v.Max) || num < float64(v.Min) {
					return fmt.Errorf("%s is not valid (min, max)", field)
				}
			}
		}

		if v.FieldType == model.FieldTypeText {
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

	if v.FieldType == model.FieldTypeObjectID {
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
	case model.FieldTypeNumber:
		return number

	case model.FieldTypeFloat:
		return numberFloat

	case model.FieldTypeText:
		return defaultValue

	case model.FieldTypeBool:
		return boolean

	case model.FieldTypeDate:
		if defaultValue == "TIME_NOW" {
			return DateNowWithLocation(*location)
		}
		return defaultValue

	case model.FieldTypeDateTime:
		if defaultValue == "TIME_NOW" {
			return TimeNowWithLocation(*location)
		}
		return defaultValue

	case model.FieldTypeRadio, model.FieldTypeSelect, model.FieldTypeFile:
		return defaultValue

	case model.FieldTypeObjectID:
		if defaultValue == "GENERATE" {
			return GenerateID()
		}

		return defaultValue

	case model.FieldTypeUuid:
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
	case model.FieldTypeNumber, model.FieldTypeFloat:
		return 0

	case model.FieldTypeText:
		return ""

	case model.FieldTypeBool:
		return false

	case model.FieldTypeDate:
		return ""

	case model.FieldTypeDateTime:
		return ""

	case model.FieldTypeObjectID, model.FieldTypeUuid:
		return ""

	case model.FieldTypeRadio, model.FieldTypeSelect, model.FieldTypeFile:
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
