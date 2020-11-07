// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: proto_files/files.proto

package files

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ManyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Powers int32 `protobuf:"varint,2,opt,name=powers,proto3" json:"powers,omitempty"`
}

func (x *ManyRequest) Reset() {
	*x = ManyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_files_files_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManyRequest) ProtoMessage() {}

func (x *ManyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_files_files_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManyRequest.ProtoReflect.Descriptor instead.
func (*ManyRequest) Descriptor() ([]byte, []int) {
	return file_proto_files_files_proto_rawDescGZIP(), []int{0}
}

func (x *ManyRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *ManyRequest) GetPowers() int32 {
	if x != nil {
		return x.Powers
	}
	return 0
}

type ManyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ManyResponse) Reset() {
	*x = ManyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_files_files_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManyResponse) ProtoMessage() {}

func (x *ManyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_files_files_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManyResponse.ProtoReflect.Descriptor instead.
func (*ManyResponse) Descriptor() ([]byte, []int) {
	return file_proto_files_files_proto_rawDescGZIP(), []int{1}
}

func (x *ManyResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_proto_files_files_proto protoreflect.FileDescriptor

var file_proto_files_files_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x22, 0x3d, 0x0a, 0x0b, 0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x77, 0x65, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x22,
	0x26, 0x0a, 0x0c, 0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x47, 0x0a, 0x0c, 0x47, 0x72, 0x65, 0x65, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x43, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x12, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x4d, 0x61,
	0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x2e, 0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_files_files_proto_rawDescOnce sync.Once
	file_proto_files_files_proto_rawDescData = file_proto_files_files_proto_rawDesc
)

func file_proto_files_files_proto_rawDescGZIP() []byte {
	file_proto_files_files_proto_rawDescOnce.Do(func() {
		file_proto_files_files_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_files_files_proto_rawDescData)
	})
	return file_proto_files_files_proto_rawDescData
}

var file_proto_files_files_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_files_files_proto_goTypes = []interface{}{
	(*ManyRequest)(nil),  // 0: files.ManyRequest
	(*ManyResponse)(nil), // 1: files.ManyResponse
}
var file_proto_files_files_proto_depIdxs = []int32{
	0, // 0: files.GreetService.Calculator:input_type -> files.ManyRequest
	1, // 1: files.GreetService.Calculator:output_type -> files.ManyResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_files_files_proto_init() }
func file_proto_files_files_proto_init() {
	if File_proto_files_files_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_files_files_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManyRequest); i {
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
		file_proto_files_files_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManyResponse); i {
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
			RawDescriptor: file_proto_files_files_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_files_files_proto_goTypes,
		DependencyIndexes: file_proto_files_files_proto_depIdxs,
		MessageInfos:      file_proto_files_files_proto_msgTypes,
	}.Build()
	File_proto_files_files_proto = out.File
	file_proto_files_files_proto_rawDesc = nil
	file_proto_files_files_proto_goTypes = nil
	file_proto_files_files_proto_depIdxs = nil
}