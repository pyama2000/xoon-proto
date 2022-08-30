// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: kintai/v1/kintai_service.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type StartRequest_Place int32

const (
	StartRequest_PLACE_UNSPECIFIED StartRequest_Place = 0
	StartRequest_PLACE_HOME        StartRequest_Place = 1
	StartRequest_PLACE_OFFICE      StartRequest_Place = 2
)

// Enum value maps for StartRequest_Place.
var (
	StartRequest_Place_name = map[int32]string{
		0: "PLACE_UNSPECIFIED",
		1: "PLACE_HOME",
		2: "PLACE_OFFICE",
	}
	StartRequest_Place_value = map[string]int32{
		"PLACE_UNSPECIFIED": 0,
		"PLACE_HOME":        1,
		"PLACE_OFFICE":      2,
	}
)

func (x StartRequest_Place) Enum() *StartRequest_Place {
	p := new(StartRequest_Place)
	*p = x
	return p
}

func (x StartRequest_Place) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StartRequest_Place) Descriptor() protoreflect.EnumDescriptor {
	return file_kintai_v1_kintai_service_proto_enumTypes[0].Descriptor()
}

func (StartRequest_Place) Type() protoreflect.EnumType {
	return &file_kintai_v1_kintai_service_proto_enumTypes[0]
}

func (x StartRequest_Place) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StartRequest_Place.Descriptor instead.
func (StartRequest_Place) EnumDescriptor() ([]byte, []int) {
	return file_kintai_v1_kintai_service_proto_rawDescGZIP(), []int{0, 0}
}

type StartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Place             StartRequest_Place        `protobuf:"varint,1,opt,name=place,proto3,enum=kintai.v1.StartRequest_Place" json:"place,omitempty"`
	SlackChannelNames []*wrapperspb.StringValue `protobuf:"bytes,2,rep,name=slack_channel_names,json=slackChannelNames,proto3" json:"slack_channel_names,omitempty"`
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

func (x *StartRequest) GetPlace() StartRequest_Place {
	if x != nil {
		return x.Place
	}
	return StartRequest_PLACE_UNSPECIFIED
}

func (x *StartRequest) GetSlackChannelNames() []*wrapperspb.StringValue {
	if x != nil {
		return x.SlackChannelNames
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

var File_kintai_v1_kintai_service_proto protoreflect.FileDescriptor

var file_kintai_v1_kintai_service_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6b, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x69, 0x6e, 0x74,
	0x61, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x09, 0x6b, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x01, 0x0a, 0x0c,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x05,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x6b, 0x69,
	0x6e, 0x74, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x52, 0x05, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x12, 0x4c, 0x0a, 0x13, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x11, 0x73, 0x6c,
	0x61, 0x63, 0x6b, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22,
	0x40, 0x0a, 0x05, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x4c, 0x41, 0x43,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0e, 0x0a, 0x0a, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x5f, 0x48, 0x4f, 0x4d, 0x45, 0x10, 0x01, 0x12,
	0x10, 0x0a, 0x0c, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x5f, 0x4f, 0x46, 0x46, 0x49, 0x43, 0x45, 0x10,
	0x02, 0x22, 0x0f, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x7f, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x0e, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x5f, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0d, 0x73, 0x6c, 0x61, 0x63, 0x6b,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x12, 0x1d, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x8e, 0x01, 0x0a, 0x0d, 0x4b, 0x69, 0x6e, 0x74, 0x61, 0x69,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x12, 0x17, 0x2e, 0x6b, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6b, 0x69, 0x6e, 0x74,
	0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x06, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x12,
	0x18, 0x2e, 0x6b, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6b, 0x69, 0x6e, 0x74,
	0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x79, 0x61, 0x6d, 0x61, 0x32, 0x30, 0x30, 0x30, 0x2f, 0x78,
	0x6f, 0x6f, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f,
	0x2f, 0x6b, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
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

var file_kintai_v1_kintai_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_kintai_v1_kintai_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_kintai_v1_kintai_service_proto_goTypes = []interface{}{
	(StartRequest_Place)(0),        // 0: kintai.v1.StartRequest.Place
	(*StartRequest)(nil),           // 1: kintai.v1.StartRequest
	(*StartResponse)(nil),          // 2: kintai.v1.StartResponse
	(*FinishRequest)(nil),          // 3: kintai.v1.FinishRequest
	(*FinishResponse)(nil),         // 4: kintai.v1.FinishResponse
	(*wrapperspb.StringValue)(nil), // 5: google.protobuf.StringValue
}
var file_kintai_v1_kintai_service_proto_depIdxs = []int32{
	0, // 0: kintai.v1.StartRequest.place:type_name -> kintai.v1.StartRequest.Place
	5, // 1: kintai.v1.StartRequest.slack_channel_names:type_name -> google.protobuf.StringValue
	5, // 2: kintai.v1.FinishRequest.slack_channels:type_name -> google.protobuf.StringValue
	1, // 3: kintai.v1.KintaiService.Start:input_type -> kintai.v1.StartRequest
	3, // 4: kintai.v1.KintaiService.Finish:input_type -> kintai.v1.FinishRequest
	2, // 5: kintai.v1.KintaiService.Start:output_type -> kintai.v1.StartResponse
	4, // 6: kintai.v1.KintaiService.Finish:output_type -> kintai.v1.FinishResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
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
	}
	file_kintai_v1_kintai_service_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kintai_v1_kintai_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
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
