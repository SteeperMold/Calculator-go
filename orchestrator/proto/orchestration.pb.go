// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: orchestration.proto

package orchestrator

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

type GetTaskResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ExpressionId  int64                  `protobuf:"varint,1,opt,name=expression_id,json=expressionId,proto3" json:"expression_id,omitempty"`
	NodeId        int64                  `protobuf:"varint,2,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Arg1          float64                `protobuf:"fixed64,3,opt,name=arg1,proto3" json:"arg1,omitempty"`
	Arg2          float64                `protobuf:"fixed64,4,opt,name=arg2,proto3" json:"arg2,omitempty"`
	Operation     string                 `protobuf:"bytes,5,opt,name=operation,proto3" json:"operation,omitempty"`
	OperationTime int32                  `protobuf:"varint,6,opt,name=operation_time,json=operationTime,proto3" json:"operation_time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTaskResponse) Reset() {
	*x = GetTaskResponse{}
	mi := &file_orchestration_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskResponse) ProtoMessage() {}

func (x *GetTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orchestration_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskResponse.ProtoReflect.Descriptor instead.
func (*GetTaskResponse) Descriptor() ([]byte, []int) {
	return file_orchestration_proto_rawDescGZIP(), []int{0}
}

func (x *GetTaskResponse) GetExpressionId() int64 {
	if x != nil {
		return x.ExpressionId
	}
	return 0
}

func (x *GetTaskResponse) GetNodeId() int64 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

func (x *GetTaskResponse) GetArg1() float64 {
	if x != nil {
		return x.Arg1
	}
	return 0
}

func (x *GetTaskResponse) GetArg2() float64 {
	if x != nil {
		return x.Arg2
	}
	return 0
}

func (x *GetTaskResponse) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *GetTaskResponse) GetOperationTime() int32 {
	if x != nil {
		return x.OperationTime
	}
	return 0
}

type PostTaskResult struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ExpressionId  int64                  `protobuf:"varint,1,opt,name=expression_id,json=expressionId,proto3" json:"expression_id,omitempty"`
	NodeId        int64                  `protobuf:"varint,2,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Result        float64                `protobuf:"fixed64,3,opt,name=result,proto3" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PostTaskResult) Reset() {
	*x = PostTaskResult{}
	mi := &file_orchestration_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostTaskResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostTaskResult) ProtoMessage() {}

func (x *PostTaskResult) ProtoReflect() protoreflect.Message {
	mi := &file_orchestration_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostTaskResult.ProtoReflect.Descriptor instead.
func (*PostTaskResult) Descriptor() ([]byte, []int) {
	return file_orchestration_proto_rawDescGZIP(), []int{1}
}

func (x *PostTaskResult) GetExpressionId() int64 {
	if x != nil {
		return x.ExpressionId
	}
	return 0
}

func (x *PostTaskResult) GetNodeId() int64 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

func (x *PostTaskResult) GetResult() float64 {
	if x != nil {
		return x.Result
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_orchestration_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_orchestration_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_orchestration_proto_rawDescGZIP(), []int{2}
}

var File_orchestration_proto protoreflect.FileDescriptor

const file_orchestration_proto_rawDesc = "" +
	"\n" +
	"\x13orchestration.proto\x12\rorchestration\"\xbc\x01\n" +
	"\x0fGetTaskResponse\x12#\n" +
	"\rexpression_id\x18\x01 \x01(\x03R\fexpressionId\x12\x17\n" +
	"\anode_id\x18\x02 \x01(\x03R\x06nodeId\x12\x12\n" +
	"\x04arg1\x18\x03 \x01(\x01R\x04arg1\x12\x12\n" +
	"\x04arg2\x18\x04 \x01(\x01R\x04arg2\x12\x1c\n" +
	"\toperation\x18\x05 \x01(\tR\toperation\x12%\n" +
	"\x0eoperation_time\x18\x06 \x01(\x05R\roperationTime\"f\n" +
	"\x0ePostTaskResult\x12#\n" +
	"\rexpression_id\x18\x01 \x01(\x03R\fexpressionId\x12\x17\n" +
	"\anode_id\x18\x02 \x01(\x03R\x06nodeId\x12\x16\n" +
	"\x06result\x18\x03 \x01(\x01R\x06result\"\a\n" +
	"\x05Empty2\x94\x01\n" +
	"\fOrchestrator\x12A\n" +
	"\tFetchTask\x12\x14.orchestration.Empty\x1a\x1e.orchestration.GetTaskResponse\x12A\n" +
	"\n" +
	"SendResult\x12\x1d.orchestration.PostTaskResult\x1a\x14.orchestration.EmptyB3Z1github.com/SteeperMold/Calculator-go/orchestratorb\x06proto3"

var (
	file_orchestration_proto_rawDescOnce sync.Once
	file_orchestration_proto_rawDescData []byte
)

func file_orchestration_proto_rawDescGZIP() []byte {
	file_orchestration_proto_rawDescOnce.Do(func() {
		file_orchestration_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_orchestration_proto_rawDesc), len(file_orchestration_proto_rawDesc)))
	})
	return file_orchestration_proto_rawDescData
}

var file_orchestration_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_orchestration_proto_goTypes = []any{
	(*GetTaskResponse)(nil), // 0: orchestration.GetTaskResponse
	(*PostTaskResult)(nil),  // 1: orchestration.PostTaskResult
	(*Empty)(nil),           // 2: orchestration.Empty
}
var file_orchestration_proto_depIdxs = []int32{
	2, // 0: orchestration.Orchestrator.FetchTask:input_type -> orchestration.Empty
	1, // 1: orchestration.Orchestrator.SendResult:input_type -> orchestration.PostTaskResult
	0, // 2: orchestration.Orchestrator.FetchTask:output_type -> orchestration.GetTaskResponse
	2, // 3: orchestration.Orchestrator.SendResult:output_type -> orchestration.Empty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_orchestration_proto_init() }
func file_orchestration_proto_init() {
	if File_orchestration_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_orchestration_proto_rawDesc), len(file_orchestration_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orchestration_proto_goTypes,
		DependencyIndexes: file_orchestration_proto_depIdxs,
		MessageInfos:      file_orchestration_proto_msgTypes,
	}.Build()
	File_orchestration_proto = out.File
	file_orchestration_proto_goTypes = nil
	file_orchestration_proto_depIdxs = nil
}
