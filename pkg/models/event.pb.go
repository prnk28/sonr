// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: event.proto

package proto

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

// [PLUGIN]
// Authorization being sent from
type InviteEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeerId   string `protobuf:"bytes,1,opt,name=peerId,proto3" json:"peerId,omitempty"`
	FilePath string `protobuf:"bytes,2,opt,name=filePath,proto3" json:"filePath,omitempty"`
}

func (x *InviteEvent) Reset() {
	*x = InviteEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteEvent) ProtoMessage() {}

func (x *InviteEvent) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteEvent.ProtoReflect.Descriptor instead.
func (*InviteEvent) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{0}
}

func (x *InviteEvent) GetPeerId() string {
	if x != nil {
		return x.PeerId
	}
	return ""
}

func (x *InviteEvent) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

// [PLUGIN]
// Authorization being sent from
type RespondEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Decision bool `protobuf:"varint,1,opt,name=decision,proto3" json:"decision,omitempty"`
}

func (x *RespondEvent) Reset() {
	*x = RespondEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespondEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespondEvent) ProtoMessage() {}

func (x *RespondEvent) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespondEvent.ProtoReflect.Descriptor instead.
func (*RespondEvent) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{1}
}

func (x *RespondEvent) GetDecision() bool {
	if x != nil {
		return x.Decision
	}
	return false
}

// [PLUGIN]
// When Direction gets updated
type UpdateEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewDirection float64 `protobuf:"fixed64,1,opt,name=newDirection,proto3" json:"newDirection,omitempty"`
}

func (x *UpdateEvent) Reset() {
	*x = UpdateEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEvent) ProtoMessage() {}

func (x *UpdateEvent) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEvent.ProtoReflect.Descriptor instead.
func (*UpdateEvent) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateEvent) GetNewDirection() float64 {
	if x != nil {
		return x.NewDirection
	}
	return 0
}

var File_event_proto protoreflect.FileDescriptor

var file_event_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a,
	0x0b, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x65, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68,
	0x22, 0x2a, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x31, 0x0a, 0x0b,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x6e,
	0x65, 0x77, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0c, 0x6e, 0x65, 0x77, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_event_proto_rawDescOnce sync.Once
	file_event_proto_rawDescData = file_event_proto_rawDesc
)

func file_event_proto_rawDescGZIP() []byte {
	file_event_proto_rawDescOnce.Do(func() {
		file_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_proto_rawDescData)
	})
	return file_event_proto_rawDescData
}

var file_event_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_event_proto_goTypes = []interface{}{
	(*InviteEvent)(nil),  // 0: InviteEvent
	(*RespondEvent)(nil), // 1: RespondEvent
	(*UpdateEvent)(nil),  // 2: UpdateEvent
}
var file_event_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_event_proto_init() }
func file_event_proto_init() {
	if File_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteEvent); i {
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
		file_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespondEvent); i {
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
		file_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEvent); i {
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
			RawDescriptor: file_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_event_proto_goTypes,
		DependencyIndexes: file_event_proto_depIdxs,
		MessageInfos:      file_event_proto_msgTypes,
	}.Build()
	File_event_proto = out.File
	file_event_proto_rawDesc = nil
	file_event_proto_goTypes = nil
	file_event_proto_depIdxs = nil
}
