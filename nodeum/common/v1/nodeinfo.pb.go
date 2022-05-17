// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: nodeum/common/v1/nodeinfo.proto

package commonv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NodeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path     string                 `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Mode     uint32                 `protobuf:"varint,2,opt,name=mode,proto3" json:"mode,omitempty"`
	Uid      uint32                 `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
	Gid      uint32                 `protobuf:"varint,4,opt,name=gid,proto3" json:"gid,omitempty"`
	Size     int64                  `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	Atime    *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=atime,proto3" json:"atime,omitempty"`
	Mtime    *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=mtime,proto3" json:"mtime,omitempty"`
	Ctime    *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=ctime,proto3" json:"ctime,omitempty"`
	Metadata map[string]string      `protobuf:"bytes,9,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *NodeInfo) Reset() {
	*x = NodeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodeum_common_v1_nodeinfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeInfo) ProtoMessage() {}

func (x *NodeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_nodeum_common_v1_nodeinfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeInfo.ProtoReflect.Descriptor instead.
func (*NodeInfo) Descriptor() ([]byte, []int) {
	return file_nodeum_common_v1_nodeinfo_proto_rawDescGZIP(), []int{0}
}

func (x *NodeInfo) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *NodeInfo) GetMode() uint32 {
	if x != nil {
		return x.Mode
	}
	return 0
}

func (x *NodeInfo) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *NodeInfo) GetGid() uint32 {
	if x != nil {
		return x.Gid
	}
	return 0
}

func (x *NodeInfo) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *NodeInfo) GetAtime() *timestamppb.Timestamp {
	if x != nil {
		return x.Atime
	}
	return nil
}

func (x *NodeInfo) GetMtime() *timestamppb.Timestamp {
	if x != nil {
		return x.Mtime
	}
	return nil
}

func (x *NodeInfo) GetCtime() *timestamppb.Timestamp {
	if x != nil {
		return x.Ctime
	}
	return nil
}

func (x *NodeInfo) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_nodeum_common_v1_nodeinfo_proto protoreflect.FileDescriptor

var file_nodeum_common_v1_nodeinfo_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x10, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x03, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x67,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x67, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x12, 0x30, 0x0a, 0x05, 0x61, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x61, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x6d, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05,
	0x6d, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x63, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x05, 0x63, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x44, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6e, 0x6f, 0x64, 0x65,
	0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a,
	0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2d,
	0x69, 0x6f, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x73, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_nodeum_common_v1_nodeinfo_proto_rawDescOnce sync.Once
	file_nodeum_common_v1_nodeinfo_proto_rawDescData = file_nodeum_common_v1_nodeinfo_proto_rawDesc
)

func file_nodeum_common_v1_nodeinfo_proto_rawDescGZIP() []byte {
	file_nodeum_common_v1_nodeinfo_proto_rawDescOnce.Do(func() {
		file_nodeum_common_v1_nodeinfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_nodeum_common_v1_nodeinfo_proto_rawDescData)
	})
	return file_nodeum_common_v1_nodeinfo_proto_rawDescData
}

var file_nodeum_common_v1_nodeinfo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_nodeum_common_v1_nodeinfo_proto_goTypes = []interface{}{
	(*NodeInfo)(nil),              // 0: nodeum.common.v1.NodeInfo
	nil,                           // 1: nodeum.common.v1.NodeInfo.MetadataEntry
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_nodeum_common_v1_nodeinfo_proto_depIdxs = []int32{
	2, // 0: nodeum.common.v1.NodeInfo.atime:type_name -> google.protobuf.Timestamp
	2, // 1: nodeum.common.v1.NodeInfo.mtime:type_name -> google.protobuf.Timestamp
	2, // 2: nodeum.common.v1.NodeInfo.ctime:type_name -> google.protobuf.Timestamp
	1, // 3: nodeum.common.v1.NodeInfo.metadata:type_name -> nodeum.common.v1.NodeInfo.MetadataEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_nodeum_common_v1_nodeinfo_proto_init() }
func file_nodeum_common_v1_nodeinfo_proto_init() {
	if File_nodeum_common_v1_nodeinfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nodeum_common_v1_nodeinfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeInfo); i {
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
			RawDescriptor: file_nodeum_common_v1_nodeinfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nodeum_common_v1_nodeinfo_proto_goTypes,
		DependencyIndexes: file_nodeum_common_v1_nodeinfo_proto_depIdxs,
		MessageInfos:      file_nodeum_common_v1_nodeinfo_proto_msgTypes,
	}.Build()
	File_nodeum_common_v1_nodeinfo_proto = out.File
	file_nodeum_common_v1_nodeinfo_proto_rawDesc = nil
	file_nodeum_common_v1_nodeinfo_proto_goTypes = nil
	file_nodeum_common_v1_nodeinfo_proto_depIdxs = nil
}