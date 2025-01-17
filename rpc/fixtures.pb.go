// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: fixtures.proto

package rpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FixtureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An event to get the default fixture for
	Event string `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *FixtureRequest) Reset() {
	*x = FixtureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fixtures_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FixtureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FixtureRequest) ProtoMessage() {}

func (x *FixtureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fixtures_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FixtureRequest.ProtoReflect.Descriptor instead.
func (*FixtureRequest) Descriptor() ([]byte, []int) {
	return file_fixtures_proto_rawDescGZIP(), []int{0}
}

func (x *FixtureRequest) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

type FixtureResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// default fixture for event
	Fixture string `protobuf:"bytes,1,opt,name=fixture,proto3" json:"fixture,omitempty"`
}

func (x *FixtureResponse) Reset() {
	*x = FixtureResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fixtures_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FixtureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FixtureResponse) ProtoMessage() {}

func (x *FixtureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fixtures_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FixtureResponse.ProtoReflect.Descriptor instead.
func (*FixtureResponse) Descriptor() ([]byte, []int) {
	return file_fixtures_proto_rawDescGZIP(), []int{1}
}

func (x *FixtureResponse) GetFixture() string {
	if x != nil {
		return x.Fixture
	}
	return ""
}

var File_fixtures_proto protoreflect.FileDescriptor

var file_fixtures_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x66, 0x69, 0x78, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x72, 0x70, 0x63, 0x22, 0x26, 0x0a, 0x0e, 0x46, 0x69, 0x78, 0x74, 0x75, 0x72, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x2b, 0x0a,
	0x0f, 0x46, 0x69, 0x78, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x78, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x66, 0x69, 0x78, 0x74, 0x75, 0x72, 0x65, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x2f,
	0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x2d, 0x63, 0x6c, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fixtures_proto_rawDescOnce sync.Once
	file_fixtures_proto_rawDescData = file_fixtures_proto_rawDesc
)

func file_fixtures_proto_rawDescGZIP() []byte {
	file_fixtures_proto_rawDescOnce.Do(func() {
		file_fixtures_proto_rawDescData = protoimpl.X.CompressGZIP(file_fixtures_proto_rawDescData)
	})
	return file_fixtures_proto_rawDescData
}

var file_fixtures_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_fixtures_proto_goTypes = []interface{}{
	(*FixtureRequest)(nil),  // 0: rpc.FixtureRequest
	(*FixtureResponse)(nil), // 1: rpc.FixtureResponse
}
var file_fixtures_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_fixtures_proto_init() }
func file_fixtures_proto_init() {
	if File_fixtures_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fixtures_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FixtureRequest); i {
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
		file_fixtures_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FixtureResponse); i {
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
			RawDescriptor: file_fixtures_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fixtures_proto_goTypes,
		DependencyIndexes: file_fixtures_proto_depIdxs,
		MessageInfos:      file_fixtures_proto_msgTypes,
	}.Build()
	File_fixtures_proto = out.File
	file_fixtures_proto_rawDesc = nil
	file_fixtures_proto_goTypes = nil
	file_fixtures_proto_depIdxs = nil
}
