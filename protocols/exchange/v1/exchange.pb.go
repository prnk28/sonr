// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: protocols/exchange/v1/exchange.proto

// Package exchange defines interfaces and types for exchange between two nodes in the network.

package exchange

import (
	common "github.com/sonr-io/core/common"
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

type MailboxMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                 // ID is the Message ID
	Payload   *common.Payload  `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`                       // Payload is the message data
	From      *common.Profile  `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`                             // Users Peer Data
	To        *common.Profile  `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`                                 // Receivers Peer Data
	Metadata  *common.Metadata `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`                     // Metadata
	CreatedAt int64            `protobuf:"varint,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"` // Timestamp
}

func (x *MailboxMessage) Reset() {
	*x = MailboxMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocols_exchange_v1_exchange_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailboxMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailboxMessage) ProtoMessage() {}

func (x *MailboxMessage) ProtoReflect() protoreflect.Message {
	mi := &file_protocols_exchange_v1_exchange_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailboxMessage.ProtoReflect.Descriptor instead.
func (*MailboxMessage) Descriptor() ([]byte, []int) {
	return file_protocols_exchange_v1_exchange_proto_rawDescGZIP(), []int{0}
}

func (x *MailboxMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MailboxMessage) GetPayload() *common.Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *MailboxMessage) GetFrom() *common.Profile {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *MailboxMessage) GetTo() *common.Profile {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *MailboxMessage) GetMetadata() *common.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *MailboxMessage) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

// Invitation Message sent on RPC
type InviteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload  *common.Payload  `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`   // Attached Data
	From     *common.Peer     `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`         // Users Peer Data
	To       *common.Peer     `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`             // Receivers Peer Data
	Metadata *common.Metadata `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"` // Metadata
}

func (x *InviteRequest) Reset() {
	*x = InviteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocols_exchange_v1_exchange_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteRequest) ProtoMessage() {}

func (x *InviteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protocols_exchange_v1_exchange_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteRequest.ProtoReflect.Descriptor instead.
func (*InviteRequest) Descriptor() ([]byte, []int) {
	return file_protocols_exchange_v1_exchange_proto_rawDescGZIP(), []int{1}
}

func (x *InviteRequest) GetPayload() *common.Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *InviteRequest) GetFrom() *common.Peer {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *InviteRequest) GetTo() *common.Peer {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *InviteRequest) GetMetadata() *common.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// Reply Message sent on RPC
type InviteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Decision bool             `protobuf:"varint,1,opt,name=decision,proto3" json:"decision,omitempty"` // Success
	From     *common.Peer     `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`          // Users Peer Data
	To       *common.Peer     `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`              // Receivers Peer Data
	Metadata *common.Metadata `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`  // Metadata
}

func (x *InviteResponse) Reset() {
	*x = InviteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocols_exchange_v1_exchange_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteResponse) ProtoMessage() {}

func (x *InviteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protocols_exchange_v1_exchange_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteResponse.ProtoReflect.Descriptor instead.
func (*InviteResponse) Descriptor() ([]byte, []int) {
	return file_protocols_exchange_v1_exchange_proto_rawDescGZIP(), []int{2}
}

func (x *InviteResponse) GetDecision() bool {
	if x != nil {
		return x.Decision
	}
	return false
}

func (x *InviteResponse) GetFrom() *common.Peer {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *InviteResponse) GetTo() *common.Peer {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *InviteResponse) GetMetadata() *common.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_protocols_exchange_v1_exchange_proto protoreflect.FileDescriptor

var file_protocols_exchange_v1_exchange_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2f, 0x65, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x73, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x11, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xde, 0x01, 0x0a, 0x0e, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x23, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x1f, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x2c, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0xa8, 0x01, 0x0a, 0x0d, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x20, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x12, 0x1c, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x02, 0x74,
	0x6f, 0x12, 0x2c, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x9a, 0x01, 0x0a, 0x0e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d,
	0x12, 0x1c, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x2c,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x22, 0x5a, 0x20,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x2d,
	0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocols_exchange_v1_exchange_proto_rawDescOnce sync.Once
	file_protocols_exchange_v1_exchange_proto_rawDescData = file_protocols_exchange_v1_exchange_proto_rawDesc
)

func file_protocols_exchange_v1_exchange_proto_rawDescGZIP() []byte {
	file_protocols_exchange_v1_exchange_proto_rawDescOnce.Do(func() {
		file_protocols_exchange_v1_exchange_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocols_exchange_v1_exchange_proto_rawDescData)
	})
	return file_protocols_exchange_v1_exchange_proto_rawDescData
}

var file_protocols_exchange_v1_exchange_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protocols_exchange_v1_exchange_proto_goTypes = []interface{}{
	(*MailboxMessage)(nil),  // 0: protocols.exchange.v1.MailboxMessage
	(*InviteRequest)(nil),   // 1: protocols.exchange.v1.InviteRequest
	(*InviteResponse)(nil),  // 2: protocols.exchange.v1.InviteResponse
	(*common.Payload)(nil),  // 3: common.Payload
	(*common.Profile)(nil),  // 4: common.Profile
	(*common.Metadata)(nil), // 5: common.Metadata
	(*common.Peer)(nil),     // 6: common.Peer
}
var file_protocols_exchange_v1_exchange_proto_depIdxs = []int32{
	3,  // 0: protocols.exchange.v1.MailboxMessage.payload:type_name -> common.Payload
	4,  // 1: protocols.exchange.v1.MailboxMessage.from:type_name -> common.Profile
	4,  // 2: protocols.exchange.v1.MailboxMessage.to:type_name -> common.Profile
	5,  // 3: protocols.exchange.v1.MailboxMessage.metadata:type_name -> common.Metadata
	3,  // 4: protocols.exchange.v1.InviteRequest.payload:type_name -> common.Payload
	6,  // 5: protocols.exchange.v1.InviteRequest.from:type_name -> common.Peer
	6,  // 6: protocols.exchange.v1.InviteRequest.to:type_name -> common.Peer
	5,  // 7: protocols.exchange.v1.InviteRequest.metadata:type_name -> common.Metadata
	6,  // 8: protocols.exchange.v1.InviteResponse.from:type_name -> common.Peer
	6,  // 9: protocols.exchange.v1.InviteResponse.to:type_name -> common.Peer
	5,  // 10: protocols.exchange.v1.InviteResponse.metadata:type_name -> common.Metadata
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_protocols_exchange_v1_exchange_proto_init() }
func file_protocols_exchange_v1_exchange_proto_init() {
	if File_protocols_exchange_v1_exchange_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocols_exchange_v1_exchange_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailboxMessage); i {
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
		file_protocols_exchange_v1_exchange_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteRequest); i {
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
		file_protocols_exchange_v1_exchange_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteResponse); i {
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
			RawDescriptor: file_protocols_exchange_v1_exchange_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protocols_exchange_v1_exchange_proto_goTypes,
		DependencyIndexes: file_protocols_exchange_v1_exchange_proto_depIdxs,
		MessageInfos:      file_protocols_exchange_v1_exchange_proto_msgTypes,
	}.Build()
	File_protocols_exchange_v1_exchange_proto = out.File
	file_protocols_exchange_v1_exchange_proto_rawDesc = nil
	file_protocols_exchange_v1_exchange_proto_goTypes = nil
	file_protocols_exchange_v1_exchange_proto_depIdxs = nil
}