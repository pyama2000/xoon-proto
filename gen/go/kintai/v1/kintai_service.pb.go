// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: kintai/v1/kintai_service.proto

package kintaiv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Config int32

const (
	Config_CONFIG_UNSPECIFIED                     Config = 0
	Config_CONFIG_KINTAI_NOTIFICATION_DESTINATION Config = 1
)

// Enum value maps for Config.
var (
	Config_name = map[int32]string{
		0: "CONFIG_UNSPECIFIED",
		1: "CONFIG_KINTAI_NOTIFICATION_DESTINATION",
	}
	Config_value = map[string]int32{
		"CONFIG_UNSPECIFIED":                     0,
		"CONFIG_KINTAI_NOTIFICATION_DESTINATION": 1,
	}
)

func (x Config) Enum() *Config {
	p := new(Config)
	*p = x
	return p
}

func (x Config) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Config) Descriptor() protoreflect.EnumDescriptor {
	return file_kintai_v1_kintai_service_proto_enumTypes[0].Descriptor()
}

func (Config) Type() protoreflect.EnumType {
	return &file_kintai_v1_kintai_service_proto_enumTypes[0]
}

func (x Config) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Config.Descriptor instead.
func (Config) EnumDescriptor() ([]byte, []int) {
	return file_kintai_v1_kintai_service_proto_rawDescGZIP(), []int{0}
}

type GetSummariesRequest_RangeType int32

const (
	GetSummariesRequest_RANGE_TYPE_UNSPECIFIED GetSummariesRequest_RangeType = 0
	GetSummariesRequest_RANGE_TYPE_WEEKLY      GetSummariesRequest_RangeType = 1
	GetSummariesRequest_RANGE_TYPE_MONTHLY     GetSummariesRequest_RangeType = 2
	GetSummariesRequest_RANGE_TYPE_QUARTER     GetSummariesRequest_RangeType = 3
	GetSummariesRequest_RANGE_TYPE_HALF_YEAR   GetSummariesRequest_RangeType = 4
)

// Enum value maps for GetSummariesRequest_RangeType.
var (
	GetSummariesRequest_RangeType_name = map[int32]string{
		0: "RANGE_TYPE_UNSPECIFIED",
		1: "RANGE_TYPE_WEEKLY",
		2: "RANGE_TYPE_MONTHLY",
		3: "RANGE_TYPE_QUARTER",
		4: "RANGE_TYPE_HALF_YEAR",
	}
	GetSummariesRequest_RangeType_value = map[string]int32{
		"RANGE_TYPE_UNSPECIFIED": 0,
		"RANGE_TYPE_WEEKLY":      1,
		"RANGE_TYPE_MONTHLY":     2,
		"RANGE_TYPE_QUARTER":     3,
		"RANGE_TYPE_HALF_YEAR":   4,
	}
)

func (x GetSummariesRequest_RangeType) Enum() *GetSummariesRequest_RangeType {
	p := new(GetSummariesRequest_RangeType)
	*p = x
	return p
}

func (x GetSummariesRequest_RangeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetSummariesRequest_RangeType) Descriptor() protoreflect.EnumDescriptor {
	return file_kintai_v1_kintai_service_proto_enumTypes[1].Descriptor()
}

func (GetSummariesRequest_RangeType) Type() protoreflect.EnumType {
	return &file_kintai_v1_kintai_service_proto_enumTypes[1]
}

func (x GetSummariesRequest_RangeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetSummariesRequest_RangeType.Descriptor instead.
func (GetSummariesRequest_RangeType) EnumDescriptor() ([]byte, []int) {
	return file_kintai_v1_kintai_service_proto_rawDescGZIP(), []int{4, 0}
}

type StartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SlackChannels []*wrapperspb.StringValue `protobuf:"bytes,3,rep,name=slack_channels,json=slackChannels,proto3" json:"slack_channels,omitempty"`
}

func (x *StartRequest) Reset() {
	*x = StartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kintai_v1_kintai_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRequest) ProtoMessage() {}

func (x *StartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kintai_v1_kintai_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartRequest.ProtoReflect.Descriptor instead.
func (*StartRequest) Descriptor() ([]byte, []int) {
	return file_kintai_v1_kintai_service_proto_rawDescGZIP(), []int{0}
}

func (x *StartRequest) GetSlackChannels() []*wrapperspb.StringValue {
	if x != nil {
		return x.SlackChannels
	}
	return nil
}

type StartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartResponse) Reset() {
	*x = StartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kintai_v1_kintai_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartResponse) ProtoMessage() {}

func (x *StartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kintai_v1_kintai_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartResponse.ProtoReflect.Descriptor instead.
func (*StartResponse) Descriptor() ([]byte, []int) {
	return file_kintai_v1_kintai_service_proto_rawDescGZIP(), []int{1}
}

type FinishRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SlackChannels []*wrapperspb.StringValue `protobuf:"bytes,1,rep,name=slack_channels,json=slackChannels,proto3" json:"slack_channels,omitempty"`
	Content       *string                   `protobuf:"bytes,2,opt,name=content,proto3,oneof" json:"content,omitempty"`
}

func (x *FinishRequest) Reset() {
	*x = FinishRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kintai_v1_kintai_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinishRequest) ProtoMessage() {}

func (x *FinishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kintai_v1_kintai_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinishRequest.ProtoReflect.Descriptor instead.
func (*FinishRequest) Descriptor() ([]byte, []int) {
	return file_kintai_v1_kintai_service_proto_rawDescGZIP(), []int{2}
}

func (x *FinishRequest) GetSlackChannels() []*wrapperspb.StringValue {
	if x != nil {
		return x.SlackChannels
	}
	return nil
}

func (x *FinishRequest) GetContent() string {
	if x != nil && x.Content != nil {
		return *x.Content
	}
	return ""
}

type FinishResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FinishResponse) Reset() {
	*x = FinishResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kintai_v1_kintai_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinishResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinishResponse) ProtoMessage() {}

func (x *FinishResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kintai_v1_kintai_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinishResponse.ProtoReflect.Descriptor instead.
func (*FinishResponse) Descriptor() ([]byte, []int) {
	return file_kintai_v1_kintai_service_proto_rawDescGZIP(), []int{3}
}

type GetSummariesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Range:
	//	*GetSummariesRequest_Type
	//	*GetSummariesRequest_DateRange_
	Range isGetSummariesRequest_Range `protobuf_oneof:"range"`
}

func (x *GetSummariesRequest) Reset() {
	*x = GetSummariesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kintai_v1_kintai_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSummariesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSummariesRequest) ProtoMessage() {}

func (x *GetSummariesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kintai_v1_kintai_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSummariesRequest.ProtoReflect.Descriptor instead.
func (*GetSummariesRequest) Descriptor() ([]byte, []int) {
	return file_kintai_v1_kintai_service_proto_rawDescGZIP(), []int{4}
}

func (m *GetSummariesRequest) GetRange() isGetSummariesRequest_Range {
	if m != nil {
		return m.Range
	}
	return nil
}

func (x *GetSummariesRequest) GetType() GetSummariesRequest_RangeType {
	if x, ok := x.GetRange().(*GetSummariesRequest_Type); ok {
		return x.Type
	}
	return GetSummariesRequest_RANGE_TYPE_UNSPECIFIED
}

func (x *GetSummariesRequest) GetDateRange() *GetSummariesRequest_DateRange {
	if x, ok := x.GetRange().(*GetSummariesRequest_DateRange_); ok {
		return x.DateRange
	}
	return nil
}

type isGetSummariesRequest_Range interface {
	isGetSummariesRequest_Range()
}

type GetSummariesRequest_Type struct {
	Type GetSummariesRequest_RangeType `protobuf:"varint,1,opt,name=type,proto3,enum=kintai.v1.GetSummariesRequest_RangeType,oneof"`
}

type GetSummariesRequest_DateRange_ struct {
	DateRange *GetSummariesRequest_DateRange `protobuf:"bytes,2,opt,name=date_range,json=dateRange,proto3,oneof"`
}

func (*GetSummariesRequest_Type) isGetSummariesRequest_Range() {}

func (*GetSummariesRequest_DateRange_) isGetSummariesRequest_Range() {}

type GetSummariesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []*GetSummariesResponse_Content `protobuf:"bytes,1,rep,name=content,proto3" json:"content,omitempty"`
}

func (x *GetSummariesResponse) Reset() {
	*x = GetSummariesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kintai_v1_kintai_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSummariesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSummariesResponse) ProtoMessage() {}

func (x *GetSummariesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kintai_v1_kintai_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSummariesResponse.ProtoReflect.Descriptor instead.
func (*GetSummariesResponse) Descriptor() ([]byte, []int) {
	return file_kintai_v1_kintai_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetSummariesResponse) GetContent() []*GetSummariesResponse_Content {
	if x != nil {
		return x.Content
	}
	return nil
}

type ConfigMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   Config `protobuf:"varint,1,opt,name=key,proto3,enum=kintai.v1.Config" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ConfigMap) Reset() {
	*x = ConfigMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kintai_v1_kintai_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigMap) ProtoMessage() {}

func (x *ConfigMap) ProtoReflect() protoreflect.Message {
	mi := &file_kintai_v1_kintai_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigMap.ProtoReflect.Descriptor instead.
func (*ConfigMap) Descriptor() ([]byte, []int) {
	return file_kintai_v1_kintai_service_proto_rawDescGZIP(), []int{6}
}

func (x *ConfigMap) GetKey() Config {
	if x != nil {
		return x.Key
	}
	return Config_CONFIG_UNSPECIFIED
}

func (x *ConfigMap) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type UpdateConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Configs []*ConfigMap `protobuf:"bytes,1,rep,name=configs,proto3" json:"configs,omitempty"`
}

func (x *UpdateConfigRequest) Reset() {
	*x = UpdateConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kintai_v1_kintai_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateConfigRequest) ProtoMessage() {}

func (x *UpdateConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kintai_v1_kintai_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateConfigRequest.ProtoReflect.Descriptor instead.
func (*UpdateConfigRequest) Descriptor() ([]byte, []int) {
	return file_kintai_v1_kintai_service_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateConfigRequest) GetConfigs() []*ConfigMap {
	if x != nil {
		return x.Configs
	}
	return nil
}

type UpdateConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateConfigResponse) Reset() {
	*x = UpdateConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kintai_v1_kintai_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateConfigResponse) ProtoMessage() {}

func (x *UpdateConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kintai_v1_kintai_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateConfigResponse.ProtoReflect.Descriptor instead.
func (*UpdateConfigResponse) Descriptor() ([]byte, []int) {
	return file_kintai_v1_kintai_service_proto_rawDescGZIP(), []int{8}
}

type GetSummariesRequest_DateRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *GetSummariesRequest_DateRange) Reset() {
	*x = GetSummariesRequest_DateRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kintai_v1_kintai_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSummariesRequest_DateRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSummariesRequest_DateRange) ProtoMessage() {}

func (x *GetSummariesRequest_DateRange) ProtoReflect() protoreflect.Message {
	mi := &file_kintai_v1_kintai_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSummariesRequest_DateRange.ProtoReflect.Descriptor instead.
func (*GetSummariesRequest_DateRange) Descriptor() ([]byte, []int) {
	return file_kintai_v1_kintai_service_proto_rawDescGZIP(), []int{4, 0}
}

func (x *GetSummariesRequest_DateRange) GetStart() *timestamppb.Timestamp {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *GetSummariesRequest_DateRange) GetEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.End
	}
	return nil
}

type GetSummariesResponse_Content struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date    *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Summary string                 `protobuf:"bytes,2,opt,name=summary,proto3" json:"summary,omitempty"`
}

func (x *GetSummariesResponse_Content) Reset() {
	*x = GetSummariesResponse_Content{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kintai_v1_kintai_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSummariesResponse_Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSummariesResponse_Content) ProtoMessage() {}

func (x *GetSummariesResponse_Content) ProtoReflect() protoreflect.Message {
	mi := &file_kintai_v1_kintai_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSummariesResponse_Content.ProtoReflect.Descriptor instead.
func (*GetSummariesResponse_Content) Descriptor() ([]byte, []int) {
	return file_kintai_v1_kintai_service_proto_rawDescGZIP(), []int{5, 0}
}

func (x *GetSummariesResponse_Content) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *GetSummariesResponse_Content) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

var File_kintai_v1_kintai_service_proto protoreflect.FileDescriptor

var file_kintai_v1_kintai_service_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6b, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x69, 0x6e, 0x74,
	0x61, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x09, 0x6b, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x0c,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x0e,
	0x73, 0x6c, 0x61, 0x63, 0x6b, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x0d, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x73, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x22, 0x0f, 0x0a,
	0x0d, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x7f,
	0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x43, 0x0a, 0x0e, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0d, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x73, 0x12, 0x1d, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22,
	0x10, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0xa1, 0x03, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x6b, 0x69, 0x6e, 0x74, 0x61, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x48, 0x00, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x6b, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x61,
	0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x48, 0x00, 0x52, 0x09, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x61, 0x6e, 0x67, 0x65, 0x1a, 0x6b, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x12, 0x2c, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x65, 0x6e,
	0x64, 0x22, 0x88, 0x01, 0x0a, 0x09, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1a, 0x0a, 0x16, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x52,
	0x41, 0x4e, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x57, 0x45, 0x45, 0x4b, 0x4c, 0x59,
	0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x4d, 0x4f, 0x4e, 0x54, 0x48, 0x4c, 0x59, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x41,
	0x4e, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x51, 0x55, 0x41, 0x52, 0x54, 0x45, 0x52,
	0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x48, 0x41, 0x4c, 0x46, 0x5f, 0x59, 0x45, 0x41, 0x52, 0x10, 0x04, 0x42, 0x07, 0x0a, 0x05,
	0x72, 0x61, 0x6e, 0x67, 0x65, 0x22, 0xae, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x6b, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x1a, 0x53, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x22, 0x46, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x4d, 0x61, 0x70, 0x12, 0x23, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x11, 0x2e, 0x6b, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x45,
	0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6b, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x61, 0x70, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x73, 0x22, 0x16, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x4c, 0x0a,
	0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x4f, 0x4e, 0x46, 0x49,
	0x47, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x2a, 0x0a, 0x26, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x4b, 0x49, 0x4e, 0x54, 0x41, 0x49,
	0x5f, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45,
	0x53, 0x54, 0x49, 0x4e, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x32, 0xb4, 0x02, 0x0a, 0x0d,
	0x4b, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a,
	0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x17, 0x2e, 0x6b, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x6b, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x06, 0x46,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x12, 0x18, 0x2e, 0x6b, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x6b, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1e, 0x2e, 0x6b,
	0x69, 0x6e, 0x74, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6b,
	0x69, 0x6e, 0x74, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x51, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x1e, 0x2e, 0x6b, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x6b, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x79, 0x61, 0x6d, 0x61, 0x32, 0x30, 0x30, 0x30, 0x2f, 0x78, 0x6f, 0x6f, 0x6e, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x6b, 0x69, 0x6e,
	0x74, 0x61, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x6b, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kintai_v1_kintai_service_proto_rawDescOnce sync.Once
	file_kintai_v1_kintai_service_proto_rawDescData = file_kintai_v1_kintai_service_proto_rawDesc
)

func file_kintai_v1_kintai_service_proto_rawDescGZIP() []byte {
	file_kintai_v1_kintai_service_proto_rawDescOnce.Do(func() {
		file_kintai_v1_kintai_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_kintai_v1_kintai_service_proto_rawDescData)
	})
	return file_kintai_v1_kintai_service_proto_rawDescData
}

var file_kintai_v1_kintai_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_kintai_v1_kintai_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_kintai_v1_kintai_service_proto_goTypes = []interface{}{
	(Config)(0),                           // 0: kintai.v1.Config
	(GetSummariesRequest_RangeType)(0),    // 1: kintai.v1.GetSummariesRequest.RangeType
	(*StartRequest)(nil),                  // 2: kintai.v1.StartRequest
	(*StartResponse)(nil),                 // 3: kintai.v1.StartResponse
	(*FinishRequest)(nil),                 // 4: kintai.v1.FinishRequest
	(*FinishResponse)(nil),                // 5: kintai.v1.FinishResponse
	(*GetSummariesRequest)(nil),           // 6: kintai.v1.GetSummariesRequest
	(*GetSummariesResponse)(nil),          // 7: kintai.v1.GetSummariesResponse
	(*ConfigMap)(nil),                     // 8: kintai.v1.ConfigMap
	(*UpdateConfigRequest)(nil),           // 9: kintai.v1.UpdateConfigRequest
	(*UpdateConfigResponse)(nil),          // 10: kintai.v1.UpdateConfigResponse
	(*GetSummariesRequest_DateRange)(nil), // 11: kintai.v1.GetSummariesRequest.DateRange
	(*GetSummariesResponse_Content)(nil),  // 12: kintai.v1.GetSummariesResponse.Content
	(*wrapperspb.StringValue)(nil),        // 13: google.protobuf.StringValue
	(*timestamppb.Timestamp)(nil),         // 14: google.protobuf.Timestamp
}
var file_kintai_v1_kintai_service_proto_depIdxs = []int32{
	13, // 0: kintai.v1.StartRequest.slack_channels:type_name -> google.protobuf.StringValue
	13, // 1: kintai.v1.FinishRequest.slack_channels:type_name -> google.protobuf.StringValue
	1,  // 2: kintai.v1.GetSummariesRequest.type:type_name -> kintai.v1.GetSummariesRequest.RangeType
	11, // 3: kintai.v1.GetSummariesRequest.date_range:type_name -> kintai.v1.GetSummariesRequest.DateRange
	12, // 4: kintai.v1.GetSummariesResponse.content:type_name -> kintai.v1.GetSummariesResponse.Content
	0,  // 5: kintai.v1.ConfigMap.key:type_name -> kintai.v1.Config
	8,  // 6: kintai.v1.UpdateConfigRequest.configs:type_name -> kintai.v1.ConfigMap
	14, // 7: kintai.v1.GetSummariesRequest.DateRange.start:type_name -> google.protobuf.Timestamp
	14, // 8: kintai.v1.GetSummariesRequest.DateRange.end:type_name -> google.protobuf.Timestamp
	14, // 9: kintai.v1.GetSummariesResponse.Content.date:type_name -> google.protobuf.Timestamp
	2,  // 10: kintai.v1.KintaiService.Start:input_type -> kintai.v1.StartRequest
	4,  // 11: kintai.v1.KintaiService.Finish:input_type -> kintai.v1.FinishRequest
	6,  // 12: kintai.v1.KintaiService.GetSummaries:input_type -> kintai.v1.GetSummariesRequest
	9,  // 13: kintai.v1.KintaiService.UpdateConfig:input_type -> kintai.v1.UpdateConfigRequest
	3,  // 14: kintai.v1.KintaiService.Start:output_type -> kintai.v1.StartResponse
	5,  // 15: kintai.v1.KintaiService.Finish:output_type -> kintai.v1.FinishResponse
	7,  // 16: kintai.v1.KintaiService.GetSummaries:output_type -> kintai.v1.GetSummariesResponse
	10, // 17: kintai.v1.KintaiService.UpdateConfig:output_type -> kintai.v1.UpdateConfigResponse
	14, // [14:18] is the sub-list for method output_type
	10, // [10:14] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_kintai_v1_kintai_service_proto_init() }
func file_kintai_v1_kintai_service_proto_init() {
	if File_kintai_v1_kintai_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kintai_v1_kintai_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartRequest); i {
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
		file_kintai_v1_kintai_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartResponse); i {
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
		file_kintai_v1_kintai_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinishRequest); i {
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
		file_kintai_v1_kintai_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinishResponse); i {
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
		file_kintai_v1_kintai_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSummariesRequest); i {
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
		file_kintai_v1_kintai_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSummariesResponse); i {
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
		file_kintai_v1_kintai_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigMap); i {
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
		file_kintai_v1_kintai_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateConfigRequest); i {
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
		file_kintai_v1_kintai_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateConfigResponse); i {
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
		file_kintai_v1_kintai_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSummariesRequest_DateRange); i {
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
		file_kintai_v1_kintai_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSummariesResponse_Content); i {
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
	file_kintai_v1_kintai_service_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_kintai_v1_kintai_service_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*GetSummariesRequest_Type)(nil),
		(*GetSummariesRequest_DateRange_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kintai_v1_kintai_service_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kintai_v1_kintai_service_proto_goTypes,
		DependencyIndexes: file_kintai_v1_kintai_service_proto_depIdxs,
		EnumInfos:         file_kintai_v1_kintai_service_proto_enumTypes,
		MessageInfos:      file_kintai_v1_kintai_service_proto_msgTypes,
	}.Build()
	File_kintai_v1_kintai_service_proto = out.File
	file_kintai_v1_kintai_service_proto_rawDesc = nil
	file_kintai_v1_kintai_service_proto_goTypes = nil
	file_kintai_v1_kintai_service_proto_depIdxs = nil
}
