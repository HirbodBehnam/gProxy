// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v6.30.0
// source: pkg/proto/proxy.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TCPStreamPacket struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          []byte                 `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TCPStreamPacket) Reset() {
	*x = TCPStreamPacket{}
	mi := &file_pkg_proto_proxy_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TCPStreamPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TCPStreamPacket) ProtoMessage() {}

func (x *TCPStreamPacket) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_proxy_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TCPStreamPacket.ProtoReflect.Descriptor instead.
func (*TCPStreamPacket) Descriptor() ([]byte, []int) {
	return file_pkg_proto_proxy_proto_rawDescGZIP(), []int{0}
}

func (x *TCPStreamPacket) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_pkg_proto_proxy_proto protoreflect.FileDescriptor

var file_pkg_proto_proxy_proto_rawDesc = string([]byte{
	0x0a, 0x15, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25,
	0x0a, 0x0f, 0x54, 0x43, 0x50, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x4d, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x16,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x43, 0x50, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54,
	0x43, 0x50, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x00,
	0x28, 0x01, 0x30, 0x01, 0x42, 0x12, 0x5a, 0x10, 0x67, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_pkg_proto_proxy_proto_rawDescOnce sync.Once
	file_pkg_proto_proxy_proto_rawDescData []byte
)

func file_pkg_proto_proxy_proto_rawDescGZIP() []byte {
	file_pkg_proto_proxy_proto_rawDescOnce.Do(func() {
		file_pkg_proto_proxy_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_pkg_proto_proxy_proto_rawDesc), len(file_pkg_proto_proxy_proto_rawDesc)))
	})
	return file_pkg_proto_proxy_proto_rawDescData
}

var file_pkg_proto_proxy_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_proto_proxy_proto_goTypes = []any{
	(*TCPStreamPacket)(nil), // 0: proto.TCPStreamPacket
}
var file_pkg_proto_proxy_proto_depIdxs = []int32{
	0, // 0: proto.ProxyService.Proxy:input_type -> proto.TCPStreamPacket
	0, // 1: proto.ProxyService.Proxy:output_type -> proto.TCPStreamPacket
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_proto_proxy_proto_init() }
func file_pkg_proto_proxy_proto_init() {
	if File_pkg_proto_proxy_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_proto_proxy_proto_rawDesc), len(file_pkg_proto_proxy_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_proxy_proto_goTypes,
		DependencyIndexes: file_pkg_proto_proxy_proto_depIdxs,
		MessageInfos:      file_pkg_proto_proxy_proto_msgTypes,
	}.Build()
	File_pkg_proto_proxy_proto = out.File
	file_pkg_proto_proxy_proto_goTypes = nil
	file_pkg_proto_proxy_proto_depIdxs = nil
}
