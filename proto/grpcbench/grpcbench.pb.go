// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v4.23.4
// source: grpcbench/grpcbench.proto

package grpcbench

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

type TestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId     string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	SentUnixNano int64  `protobuf:"varint,2,opt,name=sent_unix_nano,json=sentUnixNano,proto3" json:"sent_unix_nano,omitempty"`
	PayloadBytes int32  `protobuf:"varint,3,opt,name=payload_bytes,json=payloadBytes,proto3" json:"payload_bytes,omitempty"`
	Payload      []byte `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *TestRequest) Reset() {
	*x = TestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcbench_grpcbench_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRequest) ProtoMessage() {}

func (x *TestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpcbench_grpcbench_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRequest.ProtoReflect.Descriptor instead.
func (*TestRequest) Descriptor() ([]byte, []int) {
	return file_grpcbench_grpcbench_proto_rawDescGZIP(), []int{0}
}

func (x *TestRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *TestRequest) GetSentUnixNano() int64 {
	if x != nil {
		return x.SentUnixNano
	}
	return 0
}

func (x *TestRequest) GetPayloadBytes() int32 {
	if x != nil {
		return x.PayloadBytes
	}
	return 0
}

func (x *TestRequest) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type TestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerId         string `protobuf:"bytes,1,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	ReceivedUnixNano int64  `protobuf:"varint,2,opt,name=received_unix_nano,json=receivedUnixNano,proto3" json:"received_unix_nano,omitempty"`
	LatencyNano      int64  `protobuf:"varint,3,opt,name=latency_nano,json=latencyNano,proto3" json:"latency_nano,omitempty"`
}

func (x *TestResponse) Reset() {
	*x = TestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcbench_grpcbench_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestResponse) ProtoMessage() {}

func (x *TestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpcbench_grpcbench_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestResponse.ProtoReflect.Descriptor instead.
func (*TestResponse) Descriptor() ([]byte, []int) {
	return file_grpcbench_grpcbench_proto_rawDescGZIP(), []int{1}
}

func (x *TestResponse) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *TestResponse) GetReceivedUnixNano() int64 {
	if x != nil {
		return x.ReceivedUnixNano
	}
	return 0
}

func (x *TestResponse) GetLatencyNano() int64 {
	if x != nil {
		return x.LatencyNano
	}
	return 0
}

type StreamSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerId      string  `protobuf:"bytes,1,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	ReceivedCount int64   `protobuf:"varint,2,opt,name=received_count,json=receivedCount,proto3" json:"received_count,omitempty"`
	TotalBytes    int64   `protobuf:"varint,3,opt,name=total_bytes,json=totalBytes,proto3" json:"total_bytes,omitempty"`
	LatencyMs     float64 `protobuf:"fixed64,4,opt,name=latency_ms,json=latencyMs,proto3" json:"latency_ms,omitempty"`
}

func (x *StreamSummary) Reset() {
	*x = StreamSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcbench_grpcbench_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamSummary) ProtoMessage() {}

func (x *StreamSummary) ProtoReflect() protoreflect.Message {
	mi := &file_grpcbench_grpcbench_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamSummary.ProtoReflect.Descriptor instead.
func (*StreamSummary) Descriptor() ([]byte, []int) {
	return file_grpcbench_grpcbench_proto_rawDescGZIP(), []int{2}
}

func (x *StreamSummary) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *StreamSummary) GetReceivedCount() int64 {
	if x != nil {
		return x.ReceivedCount
	}
	return 0
}

func (x *StreamSummary) GetTotalBytes() int64 {
	if x != nil {
		return x.TotalBytes
	}
	return 0
}

func (x *StreamSummary) GetLatencyMs() float64 {
	if x != nil {
		return x.LatencyMs
	}
	return 0
}

var File_grpcbench_grpcbench_proto protoreflect.FileDescriptor

var file_grpcbench_grpcbench_proto_rawDesc = []byte{
	0x0a, 0x19, 0x67, 0x72, 0x70, 0x63, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x62, 0x65, 0x6e, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x67, 0x72, 0x70,
	0x63, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x22, 0x8f, 0x01, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x6e, 0x69, 0x78,
	0x5f, 0x6e, 0x61, 0x6e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x73, 0x65, 0x6e,
	0x74, 0x55, 0x6e, 0x69, 0x78, 0x4e, 0x61, 0x6e, 0x6f, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x7c, 0x0a, 0x0c, 0x54, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x64, 0x5f, 0x75, 0x6e, 0x69, 0x78, 0x5f, 0x6e, 0x61, 0x6e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x10, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x55, 0x6e, 0x69, 0x78, 0x4e,
	0x61, 0x6e, 0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x6e,
	0x61, 0x6e, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6c, 0x61, 0x74, 0x65, 0x6e,
	0x63, 0x79, 0x4e, 0x61, 0x6e, 0x6f, 0x22, 0x93, 0x01, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x1d, 0x0a,
	0x0a, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x09, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x4d, 0x73, 0x32, 0x8e, 0x01, 0x0a,
	0x0c, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a,
	0x09, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x54, 0x65, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x2e, 0x54,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0a, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x65, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x62, 0x65, 0x6e, 0x63, 0x68, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x2e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x28, 0x01, 0x42, 0x3f, 0x5a,
	0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x61, 0x6b, 0x61,
	0x2d, 0x67, 0x61, 0x77, 0x61, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x62, 0x65, 0x6e, 0x63, 0x68,
	0x74, 0x6f, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62,
	0x65, 0x6e, 0x63, 0x68, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpcbench_grpcbench_proto_rawDescOnce sync.Once
	file_grpcbench_grpcbench_proto_rawDescData = file_grpcbench_grpcbench_proto_rawDesc
)

func file_grpcbench_grpcbench_proto_rawDescGZIP() []byte {
	file_grpcbench_grpcbench_proto_rawDescOnce.Do(func() {
		file_grpcbench_grpcbench_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpcbench_grpcbench_proto_rawDescData)
	})
	return file_grpcbench_grpcbench_proto_rawDescData
}

var file_grpcbench_grpcbench_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_grpcbench_grpcbench_proto_goTypes = []interface{}{
	(*TestRequest)(nil),   // 0: grpcbench.TestRequest
	(*TestResponse)(nil),  // 1: grpcbench.TestResponse
	(*StreamSummary)(nil), // 2: grpcbench.StreamSummary
}
var file_grpcbench_grpcbench_proto_depIdxs = []int32{
	0, // 0: grpcbench.BenchService.UnaryTest:input_type -> grpcbench.TestRequest
	0, // 1: grpcbench.BenchService.StreamTest:input_type -> grpcbench.TestRequest
	1, // 2: grpcbench.BenchService.UnaryTest:output_type -> grpcbench.TestResponse
	2, // 3: grpcbench.BenchService.StreamTest:output_type -> grpcbench.StreamSummary
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpcbench_grpcbench_proto_init() }
func file_grpcbench_grpcbench_proto_init() {
	if File_grpcbench_grpcbench_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpcbench_grpcbench_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestRequest); i {
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
		file_grpcbench_grpcbench_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestResponse); i {
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
		file_grpcbench_grpcbench_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamSummary); i {
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
			RawDescriptor: file_grpcbench_grpcbench_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpcbench_grpcbench_proto_goTypes,
		DependencyIndexes: file_grpcbench_grpcbench_proto_depIdxs,
		MessageInfos:      file_grpcbench_grpcbench_proto_msgTypes,
	}.Build()
	File_grpcbench_grpcbench_proto = out.File
	file_grpcbench_grpcbench_proto_rawDesc = nil
	file_grpcbench_grpcbench_proto_goTypes = nil
	file_grpcbench_grpcbench_proto_depIdxs = nil
}
