package helper

import (
	"encoding/json"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/structpb"
)

func JsonToJson(data interface{}, js interface{}) error {
	body, err := json.Marshal(js)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, data)
}

func ProtoToStruct(data interface{}, m protoreflect.ProtoMessage) error {
	jsonMarshaller := protojson.MarshalOptions{
		AllowPartial:    true,
		EmitUnpopulated: true,
		UseProtoNames:   true,
	}

	js, err := jsonMarshaller.Marshal(m)
	if err != nil {
		return err
	}

	err = json.Unmarshal(js, data)
	return err
}

func StructToProto(m protoreflect.ProtoMessage, data interface{}) error {
	jsonMarshaller := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: true,
	}

	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return jsonMarshaller.Unmarshal(js, m)
}

func StringToStruct(data interface{}, js string) error {
	return json.Unmarshal([]byte(js), data)
}

func StringToProto(m protoreflect.ProtoMessage, s string) error {
	jsonMarshaller := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: true,
	}

	return jsonMarshaller.Unmarshal([]byte(s), m)
}

func ByteToProto(m protoreflect.ProtoMessage, js []byte) error {
	jsonMarshaller := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: true,
	}

	return jsonMarshaller.Unmarshal(js, m)
}

func ProtoToString(m protoreflect.ProtoMessage) (string, error) {
	jsonMarshaller := protojson.MarshalOptions{
		AllowPartial:    true,
		EmitUnpopulated: true,
	}

	js, err := jsonMarshaller.Marshal(m)
	if err != nil {
		return "", err
	}

	return string(js), nil
}

func ToProtoStruct(body interface{}) (entity *structpb.Struct, err error) {
	if body == nil {
		return nil, nil
	}

	entity = &structpb.Struct{}
	js, err := json.Marshal(body)
	if err != nil {
		return
	}

	err = entity.UnmarshalJSON(js)
	return
}

func ByteToStructPb(data []byte) (*structpb.Struct, error) {
	if data == nil {
		return nil, nil
	}

	structInfo := &structpb.Struct{}
	err := structInfo.UnmarshalJSON(data)
	return structInfo, err
}

func StringToStructPb(data string) (*structpb.Struct, error) {
	if data == "" {
		return nil, nil
	}

	structInfo := &structpb.Struct{}
	err := structInfo.UnmarshalJSON([]byte(data))
	return structInfo, err
}

func MarshalToStruct(data interface{}, resp interface{}) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(js, resp)
	if err != nil {
		return err
	}

	return nil
}

func StructPbToString(entity *structpb.Struct) (string, error) {
	js, err := entity.MarshalJSON()
	if err != nil {
		return "", err
	}

	return string(js), nil
}
