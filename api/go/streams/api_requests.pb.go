// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.9
// source: streams/api_requests.proto

package streams

import (
	_ "github.com/chirpstack/chirpstack/api/go/v4/common"
	_ "github.com/chirpstack/chirpstack/api/go/v4/gw"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ApiRequestLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// API service name.
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	// API method name.
	Method string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	// Metadata.
	Metadata map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ApiRequestLog) Reset() {
	*x = ApiRequestLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streams_api_requests_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiRequestLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiRequestLog) ProtoMessage() {}

func (x *ApiRequestLog) ProtoReflect() protoreflect.Message {
	mi := &file_streams_api_requests_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiRequestLog.ProtoReflect.Descriptor instead.
func (*ApiRequestLog) Descriptor() ([]byte, []int) {
	return file_streams_api_requests_proto_rawDescGZIP(), []int{0}
}

func (x *ApiRequestLog) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *ApiRequestLog) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *ApiRequestLog) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_streams_api_requests_proto protoreflect.FileDescriptor

var file_streams_api_requests_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x67, 0x77, 0x2f,
	0x67, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x01, 0x0a, 0x0d, 0x41, 0x70, 0x69,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x40, 0x0a, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b,
	0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x78, 0x0a, 0x19, 0x69,
	0x6f, 0x2e, 0x63, 0x68, 0x69, 0x72, 0x70, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x42, 0x10, 0x41, 0x70, 0x69, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x69, 0x72, 0x70, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x2f, 0x63, 0x68, 0x69, 0x72, 0x70, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x34, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73,
	0xaa, 0x02, 0x12, 0x43, 0x68, 0x69, 0x72, 0x70, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_streams_api_requests_proto_rawDescOnce sync.Once
	file_streams_api_requests_proto_rawDescData = file_streams_api_requests_proto_rawDesc
)

func file_streams_api_requests_proto_rawDescGZIP() []byte {
	file_streams_api_requests_proto_rawDescOnce.Do(func() {
		file_streams_api_requests_proto_rawDescData = protoimpl.X.CompressGZIP(file_streams_api_requests_proto_rawDescData)
	})
	return file_streams_api_requests_proto_rawDescData
}

var file_streams_api_requests_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_streams_api_requests_proto_goTypes = []interface{}{
	(*ApiRequestLog)(nil), // 0: streams.ApiRequestLog
	nil,                   // 1: streams.ApiRequestLog.MetadataEntry
}
var file_streams_api_requests_proto_depIdxs = []int32{
	1, // 0: streams.ApiRequestLog.metadata:type_name -> streams.ApiRequestLog.MetadataEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_streams_api_requests_proto_init() }
func file_streams_api_requests_proto_init() {
	if File_streams_api_requests_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_streams_api_requests_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiRequestLog); i {
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
			RawDescriptor: file_streams_api_requests_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_streams_api_requests_proto_goTypes,
		DependencyIndexes: file_streams_api_requests_proto_depIdxs,
		MessageInfos:      file_streams_api_requests_proto_msgTypes,
	}.Build()
	File_streams_api_requests_proto = out.File
	file_streams_api_requests_proto_rawDesc = nil
	file_streams_api_requests_proto_goTypes = nil
	file_streams_api_requests_proto_depIdxs = nil
}
