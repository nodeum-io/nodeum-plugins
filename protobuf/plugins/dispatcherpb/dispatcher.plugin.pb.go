// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.2
// source: protobuf/dispatcher.plugin.proto

package dispatcherpb

import (
	requestpb "github.com/nodeum-io/nodeum-plugins/protobuf/types/requestpb"
	taskpb "github.com/nodeum-io/nodeum-plugins/protobuf/types/taskpb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OnStartNewTaskRequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task     *taskpb.Task     `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	Metadata *structpb.Struct `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *OnStartNewTaskRequestResponse) Reset() {
	*x = OnStartNewTaskRequestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_dispatcher_plugin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnStartNewTaskRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnStartNewTaskRequestResponse) ProtoMessage() {}

func (x *OnStartNewTaskRequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_dispatcher_plugin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnStartNewTaskRequestResponse.ProtoReflect.Descriptor instead.
func (*OnStartNewTaskRequestResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_dispatcher_plugin_proto_rawDescGZIP(), []int{0}
}

func (x *OnStartNewTaskRequestResponse) GetTask() *taskpb.Task {
	if x != nil {
		return x.Task
	}
	return nil
}

func (x *OnStartNewTaskRequestResponse) GetMetadata() *structpb.Struct {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type OnNewRequestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Execution *taskpb.Execution  `protobuf:"bytes,1,opt,name=execution,proto3" json:"execution,omitempty"`
	Request   *requestpb.Request `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *OnNewRequestRequest) Reset() {
	*x = OnNewRequestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_dispatcher_plugin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnNewRequestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnNewRequestRequest) ProtoMessage() {}

func (x *OnNewRequestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_dispatcher_plugin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnNewRequestRequest.ProtoReflect.Descriptor instead.
func (*OnNewRequestRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_dispatcher_plugin_proto_rawDescGZIP(), []int{1}
}

func (x *OnNewRequestRequest) GetExecution() *taskpb.Execution {
	if x != nil {
		return x.Execution
	}
	return nil
}

func (x *OnNewRequestRequest) GetRequest() *requestpb.Request {
	if x != nil {
		return x.Request
	}
	return nil
}

type OnNewRequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *requestpb.Request `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	Discard bool               `protobuf:"varint,3,opt,name=discard,proto3" json:"discard,omitempty"`
}

func (x *OnNewRequestResponse) Reset() {
	*x = OnNewRequestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_dispatcher_plugin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnNewRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnNewRequestResponse) ProtoMessage() {}

func (x *OnNewRequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_dispatcher_plugin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnNewRequestResponse.ProtoReflect.Descriptor instead.
func (*OnNewRequestResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_dispatcher_plugin_proto_rawDescGZIP(), []int{2}
}

func (x *OnNewRequestResponse) GetRequest() *requestpb.Request {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *OnNewRequestResponse) GetDiscard() bool {
	if x != nil {
		return x.Discard
	}
	return false
}

var File_protobuf_dispatcher_plugin_proto protoreflect.FileDescriptor

var file_protobuf_dispatcher_plugin_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x69, 0x73, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0f, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f,
	0x64, 0x65, 0x75, 0x6d, 0x2d, 0x69, 0x6f, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2d, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e,
	0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2d, 0x69, 0x6f, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2d,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x7f, 0x0a, 0x1d, 0x4f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4e, 0x65, 0x77, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x29, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x12, 0x33, 0x0a, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x83, 0x01, 0x0a, 0x13, 0x4f, 0x6e, 0x4e, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e,
	0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x64, 0x0a, 0x14, 0x4f, 0x6e, 0x4e, 0x65, 0x77,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x32, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x69, 0x73, 0x63, 0x61, 0x72, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x69, 0x73, 0x63, 0x61, 0x72, 0x64, 0x32, 0xe1, 0x01,
	0x0a, 0x10, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x12, 0x70, 0x0a, 0x0e, 0x4f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4e, 0x65, 0x77,
	0x54, 0x61, 0x73, 0x6b, 0x12, 0x2e, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4e, 0x65,
	0x77, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x2e, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4e, 0x65,
	0x77, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x0c, 0x4f, 0x6e, 0x4e, 0x65, 0x77, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4f, 0x6e, 0x4e, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6e, 0x6f, 0x64,
	0x65, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4f, 0x6e, 0x4e,
	0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2d, 0x69, 0x6f, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d,
	0x2d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_dispatcher_plugin_proto_rawDescOnce sync.Once
	file_protobuf_dispatcher_plugin_proto_rawDescData = file_protobuf_dispatcher_plugin_proto_rawDesc
)

func file_protobuf_dispatcher_plugin_proto_rawDescGZIP() []byte {
	file_protobuf_dispatcher_plugin_proto_rawDescOnce.Do(func() {
		file_protobuf_dispatcher_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_dispatcher_plugin_proto_rawDescData)
	})
	return file_protobuf_dispatcher_plugin_proto_rawDescData
}

var file_protobuf_dispatcher_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protobuf_dispatcher_plugin_proto_goTypes = []interface{}{
	(*OnStartNewTaskRequestResponse)(nil), // 0: nodeum.protobuf.OnStartNewTaskRequestResponse
	(*OnNewRequestRequest)(nil),           // 1: nodeum.protobuf.OnNewRequestRequest
	(*OnNewRequestResponse)(nil),          // 2: nodeum.protobuf.OnNewRequestResponse
	(*taskpb.Task)(nil),                   // 3: nodeum.protobuf.Task
	(*structpb.Struct)(nil),               // 4: google.protobuf.Struct
	(*taskpb.Execution)(nil),              // 5: nodeum.protobuf.Execution
	(*requestpb.Request)(nil),             // 6: nodeum.protobuf.Request
}
var file_protobuf_dispatcher_plugin_proto_depIdxs = []int32{
	3, // 0: nodeum.protobuf.OnStartNewTaskRequestResponse.task:type_name -> nodeum.protobuf.Task
	4, // 1: nodeum.protobuf.OnStartNewTaskRequestResponse.metadata:type_name -> google.protobuf.Struct
	5, // 2: nodeum.protobuf.OnNewRequestRequest.execution:type_name -> nodeum.protobuf.Execution
	6, // 3: nodeum.protobuf.OnNewRequestRequest.request:type_name -> nodeum.protobuf.Request
	6, // 4: nodeum.protobuf.OnNewRequestResponse.request:type_name -> nodeum.protobuf.Request
	0, // 5: nodeum.protobuf.DispatcherPlugin.OnStartNewTask:input_type -> nodeum.protobuf.OnStartNewTaskRequestResponse
	1, // 6: nodeum.protobuf.DispatcherPlugin.OnNewRequest:input_type -> nodeum.protobuf.OnNewRequestRequest
	0, // 7: nodeum.protobuf.DispatcherPlugin.OnStartNewTask:output_type -> nodeum.protobuf.OnStartNewTaskRequestResponse
	2, // 8: nodeum.protobuf.DispatcherPlugin.OnNewRequest:output_type -> nodeum.protobuf.OnNewRequestResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_protobuf_dispatcher_plugin_proto_init() }
func file_protobuf_dispatcher_plugin_proto_init() {
	if File_protobuf_dispatcher_plugin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_dispatcher_plugin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnStartNewTaskRequestResponse); i {
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
		file_protobuf_dispatcher_plugin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnNewRequestRequest); i {
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
		file_protobuf_dispatcher_plugin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnNewRequestResponse); i {
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
			RawDescriptor: file_protobuf_dispatcher_plugin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_dispatcher_plugin_proto_goTypes,
		DependencyIndexes: file_protobuf_dispatcher_plugin_proto_depIdxs,
		MessageInfos:      file_protobuf_dispatcher_plugin_proto_msgTypes,
	}.Build()
	File_protobuf_dispatcher_plugin_proto = out.File
	file_protobuf_dispatcher_plugin_proto_rawDesc = nil
	file_protobuf_dispatcher_plugin_proto_goTypes = nil
	file_protobuf_dispatcher_plugin_proto_depIdxs = nil
}
