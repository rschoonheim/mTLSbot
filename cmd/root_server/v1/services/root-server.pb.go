// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: api/root-server.proto

package root_server_v1

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

type RootServerInstallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// storage_path - the path where the root server should store its data.
	StoragePath string `protobuf:"bytes,1,opt,name=storage_path,json=storagePath,proto3" json:"storage_path,omitempty"`
}

func (x *RootServerInstallRequest) Reset() {
	*x = RootServerInstallRequest{}
	mi := &file_api_root_server_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RootServerInstallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RootServerInstallRequest) ProtoMessage() {}

func (x *RootServerInstallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_root_server_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RootServerInstallRequest.ProtoReflect.Descriptor instead.
func (*RootServerInstallRequest) Descriptor() ([]byte, []int) {
	return file_api_root_server_proto_rawDescGZIP(), []int{0}
}

func (x *RootServerInstallRequest) GetStoragePath() string {
	if x != nil {
		return x.StoragePath
	}
	return ""
}

// InstallResponse - response to the Install request.
type RootServerInstallResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RootServerInstallResponse) Reset() {
	*x = RootServerInstallResponse{}
	mi := &file_api_root_server_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RootServerInstallResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RootServerInstallResponse) ProtoMessage() {}

func (x *RootServerInstallResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_root_server_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RootServerInstallResponse.ProtoReflect.Descriptor instead.
func (*RootServerInstallResponse) Descriptor() ([]byte, []int) {
	return file_api_root_server_proto_rawDescGZIP(), []int{1}
}

func (x *RootServerInstallResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_api_root_server_proto protoreflect.FileDescriptor

var file_api_root_server_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f, 0x6f, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x72, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6e,
	0x68, 0x65, 0x69, 0x6d, 0x2e, 0x6d, 0x74, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x6f, 0x6f,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x3d, 0x0a, 0x18, 0x52, 0x6f, 0x6f, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x50, 0x61, 0x74, 0x68, 0x22, 0x35, 0x0a, 0x19, 0x52, 0x6f, 0x6f, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x8f,
	0x01, 0x0a, 0x0a, 0x52, 0x6f, 0x6f, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x80, 0x01,
	0x0a, 0x07, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x12, 0x39, 0x2e, 0x72, 0x73, 0x63, 0x68,
	0x6f, 0x6f, 0x6e, 0x68, 0x65, 0x69, 0x6d, 0x2e, 0x6d, 0x74, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6f, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x72, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6e, 0x68, 0x65,
	0x69, 0x6d, 0x2e, 0x6d, 0x74, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x6f, 0x6f, 0x74, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6f, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x2c, 0x5a, 0x2a, 0x63, 0x6d, 0x64, 0x2f, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3b,
	0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_root_server_proto_rawDescOnce sync.Once
	file_api_root_server_proto_rawDescData = file_api_root_server_proto_rawDesc
)

func file_api_root_server_proto_rawDescGZIP() []byte {
	file_api_root_server_proto_rawDescOnce.Do(func() {
		file_api_root_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_root_server_proto_rawDescData)
	})
	return file_api_root_server_proto_rawDescData
}

var file_api_root_server_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_root_server_proto_goTypes = []any{
	(*RootServerInstallRequest)(nil),  // 0: rschoonheim.mtls.v1.root_server.RootServerInstallRequest
	(*RootServerInstallResponse)(nil), // 1: rschoonheim.mtls.v1.root_server.RootServerInstallResponse
}
var file_api_root_server_proto_depIdxs = []int32{
	0, // 0: rschoonheim.mtls.v1.root_server.RootServer.Install:input_type -> rschoonheim.mtls.v1.root_server.RootServerInstallRequest
	1, // 1: rschoonheim.mtls.v1.root_server.RootServer.Install:output_type -> rschoonheim.mtls.v1.root_server.RootServerInstallResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_root_server_proto_init() }
func file_api_root_server_proto_init() {
	if File_api_root_server_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_root_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_root_server_proto_goTypes,
		DependencyIndexes: file_api_root_server_proto_depIdxs,
		MessageInfos:      file_api_root_server_proto_msgTypes,
	}.Build()
	File_api_root_server_proto = out.File
	file_api_root_server_proto_rawDesc = nil
	file_api_root_server_proto_goTypes = nil
	file_api_root_server_proto_depIdxs = nil
}
