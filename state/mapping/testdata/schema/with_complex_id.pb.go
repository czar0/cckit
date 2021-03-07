// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.7.1
// source: with_complex_id.proto

package schema

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type EntityWithComplexId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       *EntityComplexId     `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	SomeDate *timestamp.Timestamp `protobuf:"bytes,2,opt,name=some_date,json=someDate,proto3" json:"some_date,omitempty"`
}

func (x *EntityWithComplexId) Reset() {
	*x = EntityWithComplexId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_with_complex_id_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityWithComplexId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityWithComplexId) ProtoMessage() {}

func (x *EntityWithComplexId) ProtoReflect() protoreflect.Message {
	mi := &file_with_complex_id_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityWithComplexId.ProtoReflect.Descriptor instead.
func (*EntityWithComplexId) Descriptor() ([]byte, []int) {
	return file_with_complex_id_proto_rawDescGZIP(), []int{0}
}

func (x *EntityWithComplexId) GetId() *EntityComplexId {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *EntityWithComplexId) GetSomeDate() *timestamp.Timestamp {
	if x != nil {
		return x.SomeDate
	}
	return nil
}

// EntityComplexId
type EntityComplexId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdPart1 string `protobuf:"bytes,1,opt,name=idPart1,proto3" json:"idPart1,omitempty"`
	IdPart2 string `protobuf:"bytes,2,opt,name=idPart2,proto3" json:"idPart2,omitempty"`
}

func (x *EntityComplexId) Reset() {
	*x = EntityComplexId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_with_complex_id_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityComplexId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityComplexId) ProtoMessage() {}

func (x *EntityComplexId) ProtoReflect() protoreflect.Message {
	mi := &file_with_complex_id_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityComplexId.ProtoReflect.Descriptor instead.
func (*EntityComplexId) Descriptor() ([]byte, []int) {
	return file_with_complex_id_proto_rawDescGZIP(), []int{1}
}

func (x *EntityComplexId) GetIdPart1() string {
	if x != nil {
		return x.IdPart1
	}
	return ""
}

func (x *EntityComplexId) GetIdPart2() string {
	if x != nil {
		return x.IdPart2
	}
	return ""
}

var File_with_complex_id_proto protoreflect.FileDescriptor

var file_with_complex_id_proto_rawDesc = []byte{
	0x0a, 0x15, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f, 0x69,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x77, 0x0a, 0x13, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x78, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x49, 0x64, 0x52, 0x02, 0x49, 0x64,
	0x12, 0x37, 0x0a, 0x09, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x08, 0x73, 0x6f, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x65, 0x22, 0x45, 0x0a, 0x0f, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x69, 0x64, 0x50, 0x61, 0x72, 0x74, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69,
	0x64, 0x50, 0x61, 0x72, 0x74, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x64, 0x50, 0x61, 0x72, 0x74,
	0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x64, 0x50, 0x61, 0x72, 0x74, 0x32,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_with_complex_id_proto_rawDescOnce sync.Once
	file_with_complex_id_proto_rawDescData = file_with_complex_id_proto_rawDesc
)

func file_with_complex_id_proto_rawDescGZIP() []byte {
	file_with_complex_id_proto_rawDescOnce.Do(func() {
		file_with_complex_id_proto_rawDescData = protoimpl.X.CompressGZIP(file_with_complex_id_proto_rawDescData)
	})
	return file_with_complex_id_proto_rawDescData
}

var file_with_complex_id_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_with_complex_id_proto_goTypes = []interface{}{
	(*EntityWithComplexId)(nil), // 0: schema.EntityWithComplexId
	(*EntityComplexId)(nil),     // 1: schema.EntityComplexId
	(*timestamp.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_with_complex_id_proto_depIdxs = []int32{
	1, // 0: schema.EntityWithComplexId.Id:type_name -> schema.EntityComplexId
	2, // 1: schema.EntityWithComplexId.some_date:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_with_complex_id_proto_init() }
func file_with_complex_id_proto_init() {
	if File_with_complex_id_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_with_complex_id_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityWithComplexId); i {
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
		file_with_complex_id_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityComplexId); i {
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
			RawDescriptor: file_with_complex_id_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_with_complex_id_proto_goTypes,
		DependencyIndexes: file_with_complex_id_proto_depIdxs,
		MessageInfos:      file_with_complex_id_proto_msgTypes,
	}.Build()
	File_with_complex_id_proto = out.File
	file_with_complex_id_proto_rawDesc = nil
	file_with_complex_id_proto_goTypes = nil
	file_with_complex_id_proto_depIdxs = nil
}