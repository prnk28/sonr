// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: api.proto

package models

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

// Enum for Type of Transfer
type InviteRequest_TransferType int32

const (
	InviteRequest_Unknown    InviteRequest_TransferType = 0 // Default Value, Would not Transfer
	InviteRequest_File       InviteRequest_TransferType = 1 // Media or Document
	InviteRequest_MultiFiles InviteRequest_TransferType = 2 // Multiple Media or Documents
	InviteRequest_Contact    InviteRequest_TransferType = 3 // Users Contact Card
	InviteRequest_URL        InviteRequest_TransferType = 4 // Website URL Link
)

// Enum value maps for InviteRequest_TransferType.
var (
	InviteRequest_TransferType_name = map[int32]string{
		0: "Unknown",
		1: "File",
		2: "MultiFiles",
		3: "Contact",
		4: "URL",
	}
	InviteRequest_TransferType_value = map[string]int32{
		"Unknown":    0,
		"File":       1,
		"MultiFiles": 2,
		"Contact":    3,
		"URL":        4,
	}
)

func (x InviteRequest_TransferType) Enum() *InviteRequest_TransferType {
	p := new(InviteRequest_TransferType)
	*p = x
	return p
}

func (x InviteRequest_TransferType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InviteRequest_TransferType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[0].Descriptor()
}

func (InviteRequest_TransferType) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[0]
}

func (x InviteRequest_TransferType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InviteRequest_TransferType.Descriptor instead.
func (InviteRequest_TransferType) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1, 0}
}

// Initial Connection Message to Establish Sonr
type ConnectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude    float64      `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`     // Geolocation for OLC
	Longitude   float64      `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`   // Geolocation for OLC
	Profile     *Profile     `protobuf:"bytes,3,opt,name=profile,proto3" json:"profile,omitempty"`         // General Peer Info
	Device      *Device      `protobuf:"bytes,4,opt,name=device,proto3" json:"device,omitempty"`           // Users Device
	Directories *Directories `protobuf:"bytes,5,opt,name=directories,proto3" json:"directories,omitempty"` // Users Available File Paths
	Contact     *Contact     `protobuf:"bytes,6,opt,name=contact,proto3" json:"contact,omitempty"`         // Users Contact Card
}

func (x *ConnectionRequest) Reset() {
	*x = ConnectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionRequest) ProtoMessage() {}

func (x *ConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionRequest.ProtoReflect.Descriptor instead.
func (*ConnectionRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectionRequest) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *ConnectionRequest) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *ConnectionRequest) GetProfile() *Profile {
	if x != nil {
		return x.Profile
	}
	return nil
}

func (x *ConnectionRequest) GetDevice() *Device {
	if x != nil {
		return x.Device
	}
	return nil
}

func (x *ConnectionRequest) GetDirectories() *Directories {
	if x != nil {
		return x.Directories
	}
	return nil
}

func (x *ConnectionRequest) GetContact() *Contact {
	if x != nil {
		return x.Contact
	}
	return nil
}

// Processes Given File and Invites Specified Peer
type InviteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     InviteRequest_TransferType `protobuf:"varint,1,opt,name=type,proto3,enum=InviteRequest_TransferType" json:"type,omitempty"` // General Payload Type
	To       *Peer                      `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`                                      // User thats being Invited
	IsDirect bool                       `protobuf:"varint,3,opt,name=isDirect,proto3" json:"isDirect,omitempty"`                         // Is Users own Device
	// Attached Data
	Url     string                    `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`         // URL to be Sent - Optional
	Contact *Contact                  `protobuf:"bytes,5,opt,name=contact,proto3" json:"contact,omitempty"` // Contact Card to be Sent - Optional
	Files   []*InviteRequest_FileInfo `protobuf:"bytes,6,rep,name=files,proto3" json:"files,omitempty"`     // Files to be Sent - Optional
}

func (x *InviteRequest) Reset() {
	*x = InviteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteRequest) ProtoMessage() {}

func (x *InviteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
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
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *InviteRequest) GetType() InviteRequest_TransferType {
	if x != nil {
		return x.Type
	}
	return InviteRequest_Unknown
}

func (x *InviteRequest) GetTo() *Peer {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *InviteRequest) GetIsDirect() bool {
	if x != nil {
		return x.IsDirect
	}
	return false
}

func (x *InviteRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *InviteRequest) GetContact() *Contact {
	if x != nil {
		return x.Contact
	}
	return nil
}

func (x *InviteRequest) GetFiles() []*InviteRequest_FileInfo {
	if x != nil {
		return x.Files
	}
	return nil
}

// Invitation Message sent on RPC
type AuthInvite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload  Payload       `protobuf:"varint,1,opt,name=payload,proto3,enum=Payload" json:"payload,omitempty"` // Type of Transfer
	From     *Peer         `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`                     // Users Peer Data
	Card     *TransferCard `protobuf:"bytes,3,opt,name=card,proto3" json:"card,omitempty"`                     // Card contains all Data Info, Transfer Info
	IsDirect bool          `protobuf:"varint,4,opt,name=isDirect,proto3" json:"isDirect,omitempty"`            // Is Users own Device
}

func (x *AuthInvite) Reset() {
	*x = AuthInvite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthInvite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthInvite) ProtoMessage() {}

func (x *AuthInvite) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthInvite.ProtoReflect.Descriptor instead.
func (*AuthInvite) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *AuthInvite) GetPayload() Payload {
	if x != nil {
		return x.Payload
	}
	return Payload_UNDEFINED
}

func (x *AuthInvite) GetFrom() *Peer {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *AuthInvite) GetCard() *TransferCard {
	if x != nil {
		return x.Card
	}
	return nil
}

func (x *AuthInvite) GetIsDirect() bool {
	if x != nil {
		return x.IsDirect
	}
	return false
}

// Reply Message sent on RPC
type AuthReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload  Payload       `protobuf:"varint,1,opt,name=payload,proto3,enum=Payload" json:"payload,omitempty"` // Type of Transfer
	From     *Peer         `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`                     // Users Peer Data
	Decision bool          `protobuf:"varint,3,opt,name=decision,proto3" json:"decision,omitempty"`            // Users Decision for the Invitation
	Card     *TransferCard `protobuf:"bytes,4,opt,name=card,proto3" json:"card,omitempty"`                     // Card contains all Data Info, Transfer Info
}

func (x *AuthReply) Reset() {
	*x = AuthReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthReply) ProtoMessage() {}

func (x *AuthReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthReply.ProtoReflect.Descriptor instead.
func (*AuthReply) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *AuthReply) GetPayload() Payload {
	if x != nil {
		return x.Payload
	}
	return Payload_UNDEFINED
}

func (x *AuthReply) GetFrom() *Peer {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *AuthReply) GetDecision() bool {
	if x != nil {
		return x.Decision
	}
	return false
}

func (x *AuthReply) GetCard() *TransferCard {
	if x != nil {
		return x.Card
	}
	return nil
}

// Error Message returned from Core Library
type ErrorMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"` // Generated Error Message
	Method  string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`   // Method Error Occurred
}

func (x *ErrorMessage) Reset() {
	*x = ErrorMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorMessage) ProtoMessage() {}

func (x *ErrorMessage) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorMessage.ProtoReflect.Descriptor instead.
func (*ErrorMessage) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *ErrorMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ErrorMessage) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

// FileInfo Message for Attached Files
type InviteRequest_FileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path         string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`                  // Location of File
	HasThumbnail bool   `protobuf:"varint,2,opt,name=hasThumbnail,proto3" json:"hasThumbnail,omitempty"` // Provided File already has thumbnail
	Thumbpath    string `protobuf:"bytes,3,opt,name=thumbpath,proto3" json:"thumbpath,omitempty"`        // Path of Provided Thumbnail - Optional
	Duration     int32  `protobuf:"varint,4,opt,name=duration,proto3" json:"duration,omitempty"`         // Duration of Video - Optional
}

func (x *InviteRequest_FileInfo) Reset() {
	*x = InviteRequest_FileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteRequest_FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteRequest_FileInfo) ProtoMessage() {}

func (x *InviteRequest_FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteRequest_FileInfo.ProtoReflect.Descriptor instead.
func (*InviteRequest_FileInfo) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1, 0}
}

func (x *InviteRequest_FileInfo) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *InviteRequest_FileInfo) GetHasThumbnail() bool {
	if x != nil {
		return x.HasThumbnail
	}
	return false
}

func (x *InviteRequest_FileInfo) GetThumbpath() string {
	if x != nil {
		return x.Thumbpath
	}
	return ""
}

func (x *InviteRequest_FileInfo) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xe6, 0x01, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x07,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x0b, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x0b, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x22, 0xa3, 0x03, 0x0a,
	0x0d, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x15, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x50, 0x65,
	0x65, 0x72, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x12, 0x22, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x2d, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x1a, 0x7c, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x22, 0x0a, 0x0c, 0x68, 0x61, 0x73, 0x54, 0x68,
	0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x68,
	0x61, 0x73, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x68, 0x75, 0x6d, 0x62, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x74, 0x68, 0x75, 0x6d, 0x62, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4b, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x55, 0x52, 0x4c,
	0x10, 0x04, 0x22, 0x8a, 0x01, 0x0a, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x49, 0x6e, 0x76, 0x69, 0x74,
	0x65, 0x12, 0x22, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x08, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x19, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d,
	0x12, 0x21, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x43, 0x61, 0x72, 0x64, 0x52, 0x04, 0x63,
	0x61, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x22,
	0x89, 0x01, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x22, 0x0a,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x08,
	0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x19, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x05, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x43, 0x61, 0x72, 0x64, 0x52, 0x04, 0x63, 0x61, 0x72, 0x64, 0x22, 0x40, 0x0a, 0x0c, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x3b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_proto_goTypes = []interface{}{
	(InviteRequest_TransferType)(0), // 0: InviteRequest.TransferType
	(*ConnectionRequest)(nil),       // 1: ConnectionRequest
	(*InviteRequest)(nil),           // 2: InviteRequest
	(*AuthInvite)(nil),              // 3: AuthInvite
	(*AuthReply)(nil),               // 4: AuthReply
	(*ErrorMessage)(nil),            // 5: ErrorMessage
	(*InviteRequest_FileInfo)(nil),  // 6: InviteRequest.FileInfo
	(*Profile)(nil),                 // 7: Profile
	(*Device)(nil),                  // 8: Device
	(*Directories)(nil),             // 9: Directories
	(*Contact)(nil),                 // 10: Contact
	(*Peer)(nil),                    // 11: Peer
	(Payload)(0),                    // 12: Payload
	(*TransferCard)(nil),            // 13: TransferCard
}
var file_api_proto_depIdxs = []int32{
	7,  // 0: ConnectionRequest.profile:type_name -> Profile
	8,  // 1: ConnectionRequest.device:type_name -> Device
	9,  // 2: ConnectionRequest.directories:type_name -> Directories
	10, // 3: ConnectionRequest.contact:type_name -> Contact
	0,  // 4: InviteRequest.type:type_name -> InviteRequest.TransferType
	11, // 5: InviteRequest.to:type_name -> Peer
	10, // 6: InviteRequest.contact:type_name -> Contact
	6,  // 7: InviteRequest.files:type_name -> InviteRequest.FileInfo
	12, // 8: AuthInvite.payload:type_name -> Payload
	11, // 9: AuthInvite.from:type_name -> Peer
	13, // 10: AuthInvite.card:type_name -> TransferCard
	12, // 11: AuthReply.payload:type_name -> Payload
	11, // 12: AuthReply.from:type_name -> Peer
	13, // 13: AuthReply.card:type_name -> TransferCard
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	file_data_proto_init()
	file_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionRequest); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthInvite); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthReply); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorMessage); i {
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
		file_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteRequest_FileInfo); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		EnumInfos:         file_api_proto_enumTypes,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
