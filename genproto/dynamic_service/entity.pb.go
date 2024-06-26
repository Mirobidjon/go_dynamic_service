// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: entity.proto

package dynamic_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Entity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XId      string           `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id" bson:"_id"`
	Slug     string           `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug" bson:"slug"`
	Data     *structpb.Struct `protobuf:"bytes,3,opt,name=data,proto3" json:"data" bson:"data"`
	Location string           `protobuf:"bytes,4,opt,name=location,proto3" json:"location" bson:"location"`
}

func (x *Entity) Reset() {
	*x = Entity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entity) ProtoMessage() {}

func (x *Entity) ProtoReflect() protoreflect.Message {
	mi := &file_entity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entity.ProtoReflect.Descriptor instead.
func (*Entity) Descriptor() ([]byte, []int) {
	return file_entity_proto_rawDescGZIP(), []int{0}
}

func (x *Entity) GetXId() string {
	if x != nil {
		return x.XId
	}
	return ""
}

func (x *Entity) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *Entity) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Entity) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

type GetAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slug     string             `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug" bson:"slug"`
	Entities []*structpb.Struct `protobuf:"bytes,2,rep,name=entities,proto3" json:"entities" bson:"entities"`
	Count    int32              `protobuf:"varint,3,opt,name=count,proto3" json:"count" bson:"count"`
}

func (x *GetAllResponse) Reset() {
	*x = GetAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllResponse) ProtoMessage() {}

func (x *GetAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_entity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllResponse.ProtoReflect.Descriptor instead.
func (*GetAllResponse) Descriptor() ([]byte, []int) {
	return file_entity_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllResponse) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *GetAllResponse) GetEntities() []*structpb.Struct {
	if x != nil {
		return x.Entities
	}
	return nil
}

func (x *GetAllResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type GetByPk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XId      string `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id" bson:"_id"`
	Slug     string `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug" bson:"slug"`
	Location string `protobuf:"bytes,3,opt,name=location,proto3" json:"location" bson:"location"`
}

func (x *GetByPk) Reset() {
	*x = GetByPk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByPk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByPk) ProtoMessage() {}

func (x *GetByPk) ProtoReflect() protoreflect.Message {
	mi := &file_entity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByPk.ProtoReflect.Descriptor instead.
func (*GetByPk) Descriptor() ([]byte, []int) {
	return file_entity_proto_rawDescGZIP(), []int{2}
}

func (x *GetByPk) GetXId() string {
	if x != nil {
		return x.XId
	}
	return ""
}

func (x *GetByPk) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *GetByPk) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

type GetAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slug     string           `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug" bson:"slug"`
	Limit    int32            `protobuf:"varint,2,opt,name=limit,proto3" json:"limit" bson:"limit"`
	Offset   int32            `protobuf:"varint,3,opt,name=offset,proto3" json:"offset" bson:"offset"`
	Sort     string           `protobuf:"bytes,4,opt,name=sort,proto3" json:"sort" bson:"sort"`
	Order    string           `protobuf:"bytes,5,opt,name=order,proto3" json:"order" bson:"order"`
	Search   string           `protobuf:"bytes,6,opt,name=search,proto3" json:"search" bson:"search"`
	Data     *structpb.Struct `protobuf:"bytes,7,opt,name=data,proto3" json:"data" bson:"data"`
	Location string           `protobuf:"bytes,8,opt,name=location,proto3" json:"location" bson:"location"`
}

func (x *GetAllRequest) Reset() {
	*x = GetAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRequest) ProtoMessage() {}

func (x *GetAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_entity_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRequest.ProtoReflect.Descriptor instead.
func (*GetAllRequest) Descriptor() ([]byte, []int) {
	return file_entity_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllRequest) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *GetAllRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetAllRequest) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

func (x *GetAllRequest) GetOrder() string {
	if x != nil {
		return x.Order
	}
	return ""
}

func (x *GetAllRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *GetAllRequest) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GetAllRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

type GetBySlug struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slug     string `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug" bson:"slug"`
	Location string `protobuf:"bytes,2,opt,name=location,proto3" json:"location" bson:"location"`
}

func (x *GetBySlug) Reset() {
	*x = GetBySlug{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBySlug) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBySlug) ProtoMessage() {}

func (x *GetBySlug) ProtoReflect() protoreflect.Message {
	mi := &file_entity_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBySlug.ProtoReflect.Descriptor instead.
func (*GetBySlug) Descriptor() ([]byte, []int) {
	return file_entity_proto_rawDescGZIP(), []int{4}
}

func (x *GetBySlug) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *GetBySlug) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

type LookUps struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From         string `protobuf:"bytes,1,opt,name=from,proto3" json:"from" bson:"from"`
	LocalField   string `protobuf:"bytes,2,opt,name=localField,proto3" json:"localField" bson:"localField"`
	ForeignField string `protobuf:"bytes,3,opt,name=foreignField,proto3" json:"foreignField" bson:"foreignField"`
	As           string `protobuf:"bytes,4,opt,name=as,proto3" json:"as" bson:"as"`
}

func (x *LookUps) Reset() {
	*x = LookUps{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LookUps) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookUps) ProtoMessage() {}

func (x *LookUps) ProtoReflect() protoreflect.Message {
	mi := &file_entity_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookUps.ProtoReflect.Descriptor instead.
func (*LookUps) Descriptor() ([]byte, []int) {
	return file_entity_proto_rawDescGZIP(), []int{5}
}

func (x *LookUps) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *LookUps) GetLocalField() string {
	if x != nil {
		return x.LocalField
	}
	return ""
}

func (x *LookUps) GetForeignField() string {
	if x != nil {
		return x.ForeignField
	}
	return ""
}

func (x *LookUps) GetAs() string {
	if x != nil {
		return x.As
	}
	return ""
}

type AggregationGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XId         string `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id" bson:"_id"`
	Field       string `protobuf:"bytes,2,opt,name=field,proto3" json:"field" bson:"field"`
	Accumulator string `protobuf:"bytes,3,opt,name=accumulator,proto3" json:"accumulator" bson:"accumulator"`
	Expression  string `protobuf:"bytes,4,opt,name=expression,proto3" json:"expression" bson:"expression"`
}

func (x *AggregationGroup) Reset() {
	*x = AggregationGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AggregationGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggregationGroup) ProtoMessage() {}

func (x *AggregationGroup) ProtoReflect() protoreflect.Message {
	mi := &file_entity_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggregationGroup.ProtoReflect.Descriptor instead.
func (*AggregationGroup) Descriptor() ([]byte, []int) {
	return file_entity_proto_rawDescGZIP(), []int{6}
}

func (x *AggregationGroup) GetXId() string {
	if x != nil {
		return x.XId
	}
	return ""
}

func (x *AggregationGroup) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *AggregationGroup) GetAccumulator() string {
	if x != nil {
		return x.Accumulator
	}
	return ""
}

func (x *AggregationGroup) GetExpression() string {
	if x != nil {
		return x.Expression
	}
	return ""
}

type Aggregate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lookups []*LookUps        `protobuf:"bytes,4,rep,name=lookups,proto3" json:"lookups" bson:"lookups"`
	Group   *AggregationGroup `protobuf:"bytes,5,opt,name=group,proto3" json:"group" bson:"group"`
}

func (x *Aggregate) Reset() {
	*x = Aggregate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Aggregate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Aggregate) ProtoMessage() {}

func (x *Aggregate) ProtoReflect() protoreflect.Message {
	mi := &file_entity_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Aggregate.ProtoReflect.Descriptor instead.
func (*Aggregate) Descriptor() ([]byte, []int) {
	return file_entity_proto_rawDescGZIP(), []int{7}
}

func (x *Aggregate) GetLookups() []*LookUps {
	if x != nil {
		return x.Lookups
	}
	return nil
}

func (x *Aggregate) GetGroup() *AggregationGroup {
	if x != nil {
		return x.Group
	}
	return nil
}

type GetJoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slug      string           `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug" bson:"slug"`
	Limit     int32            `protobuf:"varint,2,opt,name=limit,proto3" json:"limit" bson:"limit"`
	Offset    int32            `protobuf:"varint,3,opt,name=offset,proto3" json:"offset" bson:"offset"`
	Sort      string           `protobuf:"bytes,4,opt,name=sort,proto3" json:"sort" bson:"sort"`
	Order     string           `protobuf:"bytes,5,opt,name=order,proto3" json:"order" bson:"order"`
	Search    string           `protobuf:"bytes,6,opt,name=search,proto3" json:"search" bson:"search"`
	Data      *structpb.Struct `protobuf:"bytes,7,opt,name=data,proto3" json:"data" bson:"data"`
	Location  string           `protobuf:"bytes,8,opt,name=location,proto3" json:"location" bson:"location"`
	Aggregate *Aggregate       `protobuf:"bytes,9,opt,name=aggregate,proto3" json:"aggregate" bson:"aggregate"`
}

func (x *GetJoinRequest) Reset() {
	*x = GetJoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJoinRequest) ProtoMessage() {}

func (x *GetJoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_entity_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJoinRequest.ProtoReflect.Descriptor instead.
func (*GetJoinRequest) Descriptor() ([]byte, []int) {
	return file_entity_proto_rawDescGZIP(), []int{8}
}

func (x *GetJoinRequest) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *GetJoinRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetJoinRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetJoinRequest) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

func (x *GetJoinRequest) GetOrder() string {
	if x != nil {
		return x.Order
	}
	return ""
}

func (x *GetJoinRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *GetJoinRequest) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GetJoinRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *GetJoinRequest) GetAggregate() *Aggregate {
	if x != nil {
		return x.Aggregate
	}
	return nil
}

var File_entity_proto protoreflect.FileDescriptor

var file_entity_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x76, 0x0a, 0x06, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x0f, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x6f, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x33, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x4a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x50, 0x6b, 0x12, 0x0f,
	0x0a, 0x03, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73,
	0x6c, 0x75, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0xdc, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3b,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x42, 0x79, 0x53, 0x6c, 0x75, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x6c, 0x75, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x71, 0x0a, 0x07, 0x4c,
	0x6f, 0x6f, 0x6b, 0x55, 0x70, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x6f,
	0x72, 0x65, 0x69, 0x67, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x66, 0x6f, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x61, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x61, 0x73, 0x22, 0x7b,
	0x0a, 0x10, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x0f, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63,
	0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x61, 0x63, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x65,
	0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x78, 0x0a, 0x09, 0x41,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x6c, 0x6f, 0x6f, 0x6b,
	0x75, 0x70, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64, 0x79, 0x6e, 0x61,
	0x6d, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b,
	0x55, 0x70, 0x73, 0x52, 0x07, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x73, 0x12, 0x37, 0x0a, 0x05,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x64, 0x79,
	0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x97, 0x02, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f,
	0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x2b, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x09, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d,
	0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x52, 0x09, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x32,
	0xe7, 0x03, 0x0a, 0x0d, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3c, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x64, 0x79,
	0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x3c, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x64, 0x79, 0x6e, 0x61,
	0x6d, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x1a, 0x17, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x00, 0x12, 0x41, 0x0a,
	0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x63, 0x68, 0x12, 0x17, 0x2e, 0x64,
	0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x3b, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x64, 0x79, 0x6e,
	0x61, 0x6d, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4b, 0x0a,
	0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x1e, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69,
	0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69,
	0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x18, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x50, 0x6b, 0x1a,
	0x17, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x1f, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1a, 0x5a, 0x18, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_entity_proto_rawDescOnce sync.Once
	file_entity_proto_rawDescData = file_entity_proto_rawDesc
)

func file_entity_proto_rawDescGZIP() []byte {
	file_entity_proto_rawDescOnce.Do(func() {
		file_entity_proto_rawDescData = protoimpl.X.CompressGZIP(file_entity_proto_rawDescData)
	})
	return file_entity_proto_rawDescData
}

var file_entity_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_entity_proto_goTypes = []interface{}{
	(*Entity)(nil),           // 0: dynamic_service.Entity
	(*GetAllResponse)(nil),   // 1: dynamic_service.GetAllResponse
	(*GetByPk)(nil),          // 2: dynamic_service.GetByPk
	(*GetAllRequest)(nil),    // 3: dynamic_service.GetAllRequest
	(*GetBySlug)(nil),        // 4: dynamic_service.GetBySlug
	(*LookUps)(nil),          // 5: dynamic_service.LookUps
	(*AggregationGroup)(nil), // 6: dynamic_service.AggregationGroup
	(*Aggregate)(nil),        // 7: dynamic_service.Aggregate
	(*GetJoinRequest)(nil),   // 8: dynamic_service.GetJoinRequest
	(*structpb.Struct)(nil),  // 9: google.protobuf.Struct
	(*emptypb.Empty)(nil),    // 10: google.protobuf.Empty
}
var file_entity_proto_depIdxs = []int32{
	9,  // 0: dynamic_service.Entity.data:type_name -> google.protobuf.Struct
	9,  // 1: dynamic_service.GetAllResponse.entities:type_name -> google.protobuf.Struct
	9,  // 2: dynamic_service.GetAllRequest.data:type_name -> google.protobuf.Struct
	5,  // 3: dynamic_service.Aggregate.lookups:type_name -> dynamic_service.LookUps
	6,  // 4: dynamic_service.Aggregate.group:type_name -> dynamic_service.AggregationGroup
	9,  // 5: dynamic_service.GetJoinRequest.data:type_name -> google.protobuf.Struct
	7,  // 6: dynamic_service.GetJoinRequest.aggregate:type_name -> dynamic_service.Aggregate
	0,  // 7: dynamic_service.EntityService.Create:input_type -> dynamic_service.Entity
	0,  // 8: dynamic_service.EntityService.Update:input_type -> dynamic_service.Entity
	0,  // 9: dynamic_service.EntityService.UpdatePatch:input_type -> dynamic_service.Entity
	0,  // 10: dynamic_service.EntityService.Delete:input_type -> dynamic_service.Entity
	3,  // 11: dynamic_service.EntityService.GetAll:input_type -> dynamic_service.GetAllRequest
	2,  // 12: dynamic_service.EntityService.GetById:input_type -> dynamic_service.GetByPk
	8,  // 13: dynamic_service.EntityService.GetJoin:input_type -> dynamic_service.GetJoinRequest
	0,  // 14: dynamic_service.EntityService.Create:output_type -> dynamic_service.Entity
	0,  // 15: dynamic_service.EntityService.Update:output_type -> dynamic_service.Entity
	0,  // 16: dynamic_service.EntityService.UpdatePatch:output_type -> dynamic_service.Entity
	10, // 17: dynamic_service.EntityService.Delete:output_type -> google.protobuf.Empty
	1,  // 18: dynamic_service.EntityService.GetAll:output_type -> dynamic_service.GetAllResponse
	0,  // 19: dynamic_service.EntityService.GetById:output_type -> dynamic_service.Entity
	1,  // 20: dynamic_service.EntityService.GetJoin:output_type -> dynamic_service.GetAllResponse
	14, // [14:21] is the sub-list for method output_type
	7,  // [7:14] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_entity_proto_init() }
func file_entity_proto_init() {
	if File_entity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_entity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entity); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_entity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_entity_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByPk); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_entity_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_entity_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBySlug); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_entity_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LookUps); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_entity_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AggregationGroup); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_entity_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Aggregate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_entity_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJoinRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_entity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_entity_proto_goTypes,
		DependencyIndexes: file_entity_proto_depIdxs,
		MessageInfos:      file_entity_proto_msgTypes,
	}.Build()
	File_entity_proto = out.File
	file_entity_proto_rawDesc = nil
	file_entity_proto_goTypes = nil
	file_entity_proto_depIdxs = nil
}
