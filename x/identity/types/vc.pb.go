// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sonr/identity/v1/vc.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Proof represents a credential/presentation proof as defined by the Linked Data Proofs 1.0 specification (https://w3c-ccg.github.io/ld-proofs/).
type Proof struct {
	// Type defines the specific proof type used. For example, an Ed25519Signature2018 type indicates that the proof includes a digital signature produced by an ed25519 cryptographic key.
	Type ProofType `protobuf:"varint,1,opt,name=type,proto3,enum=sonrio.sonr.identity.v1.ProofType" json:"type,omitempty"`
	// ProofPurpose defines the intent for the proof, the reason why an entity created it. Acts as a safeguard to prevent the proof from being misused for a purpose other than the one it was intended for.
	ProofPurpose string `protobuf:"bytes,2,opt,name=proof_purpose,json=proofPurpose,proto3" json:"proof_purpose,omitempty"`
	// VerificationMethod points to the ID that can be used to verify the proof, eg: a public key.
	VerificationMethod string `protobuf:"bytes,3,opt,name=verification_method,json=verificationMethod,proto3" json:"verification_method,omitempty"`
	// Created notes when the proof was created using a iso8601 string
	Created string `protobuf:"bytes,4,opt,name=created,proto3" json:"created,omitempty"`
	// Domain specifies the restricted domain of the proof
	Domain string `protobuf:"bytes,5,opt,name=domain,proto3" json:"domain,omitempty"`
}

func (m *Proof) Reset()         { *m = Proof{} }
func (m *Proof) String() string { return proto.CompactTextString(m) }
func (*Proof) ProtoMessage()    {}
func (*Proof) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a30edc5c71454f0, []int{0}
}
func (m *Proof) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Proof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Proof.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Proof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proof.Merge(m, src)
}
func (m *Proof) XXX_Size() int {
	return m.Size()
}
func (m *Proof) XXX_DiscardUnknown() {
	xxx_messageInfo_Proof.DiscardUnknown(m)
}

var xxx_messageInfo_Proof proto.InternalMessageInfo

func (m *Proof) GetType() ProofType {
	if m != nil {
		return m.Type
	}
	return ProofType_ProofType_UNSPECIFIED
}

func (m *Proof) GetProofPurpose() string {
	if m != nil {
		return m.ProofPurpose
	}
	return ""
}

func (m *Proof) GetVerificationMethod() string {
	if m != nil {
		return m.VerificationMethod
	}
	return ""
}

func (m *Proof) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *Proof) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

// JSONWebSignature2020Proof is a VC proof with a signature according to JsonWebSignature2020
type JSONWebSignature2020Proof struct {
	Proof *Proof `protobuf:"bytes,1,opt,name=proof,proto3" json:"proof,omitempty"`
	Jws   string `protobuf:"bytes,2,opt,name=jws,proto3" json:"jws,omitempty"`
}

func (m *JSONWebSignature2020Proof) Reset()         { *m = JSONWebSignature2020Proof{} }
func (m *JSONWebSignature2020Proof) String() string { return proto.CompactTextString(m) }
func (*JSONWebSignature2020Proof) ProtoMessage()    {}
func (*JSONWebSignature2020Proof) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a30edc5c71454f0, []int{1}
}
func (m *JSONWebSignature2020Proof) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *JSONWebSignature2020Proof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_JSONWebSignature2020Proof.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *JSONWebSignature2020Proof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JSONWebSignature2020Proof.Merge(m, src)
}
func (m *JSONWebSignature2020Proof) XXX_Size() int {
	return m.Size()
}
func (m *JSONWebSignature2020Proof) XXX_DiscardUnknown() {
	xxx_messageInfo_JSONWebSignature2020Proof.DiscardUnknown(m)
}

var xxx_messageInfo_JSONWebSignature2020Proof proto.InternalMessageInfo

func (m *JSONWebSignature2020Proof) GetProof() *Proof {
	if m != nil {
		return m.Proof
	}
	return nil
}

func (m *JSONWebSignature2020Proof) GetJws() string {
	if m != nil {
		return m.Jws
	}
	return ""
}

// VerifiableCredential represents a credential as defined by the Verifiable Credentials Data Model 1.0 specification (https://www.w3.org/TR/vc-data-model/).
type VerifiableCredential struct {
	// ID is the unique identifier for the credential.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Context is a list of URIs that define the context of the credential.
	Context []string `protobuf:"bytes,2,rep,name=context,proto3" json:"context,omitempty"`
	// Type is a list of URIs that define the type of the credential.
	Type []string `protobuf:"bytes,3,rep,name=type,proto3" json:"type,omitempty"`
	// Issuer is the DID of the issuer of the credential.
	Issuer string `protobuf:"bytes,4,opt,name=issuer,proto3" json:"issuer,omitempty"`
	// IssuanceDate is the date the credential was issued.
	IssuanceDate string `protobuf:"bytes,5,opt,name=issuance_date,json=issuanceDate,proto3" json:"issuance_date,omitempty"`
	// ExpirationDate is the date the credential expires.
	ExpirationDate string `protobuf:"bytes,6,opt,name=expiration_date,json=expirationDate,proto3" json:"expiration_date,omitempty"`
	// CredentialSubject is the subject of the credential.
	CredentialSubject map[string]string `protobuf:"bytes,7,rep,name=credential_subject,json=credentialSubject,proto3" json:"credential_subject,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Proof is the proof of the credential.
	Proof map[string]string `protobuf:"bytes,8,rep,name=proof,proto3" json:"proof,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *VerifiableCredential) Reset()         { *m = VerifiableCredential{} }
func (m *VerifiableCredential) String() string { return proto.CompactTextString(m) }
func (*VerifiableCredential) ProtoMessage()    {}
func (*VerifiableCredential) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a30edc5c71454f0, []int{2}
}
func (m *VerifiableCredential) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VerifiableCredential) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VerifiableCredential.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VerifiableCredential) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifiableCredential.Merge(m, src)
}
func (m *VerifiableCredential) XXX_Size() int {
	return m.Size()
}
func (m *VerifiableCredential) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifiableCredential.DiscardUnknown(m)
}

var xxx_messageInfo_VerifiableCredential proto.InternalMessageInfo

func (m *VerifiableCredential) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *VerifiableCredential) GetContext() []string {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *VerifiableCredential) GetType() []string {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *VerifiableCredential) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *VerifiableCredential) GetIssuanceDate() string {
	if m != nil {
		return m.IssuanceDate
	}
	return ""
}

func (m *VerifiableCredential) GetExpirationDate() string {
	if m != nil {
		return m.ExpirationDate
	}
	return ""
}

func (m *VerifiableCredential) GetCredentialSubject() map[string]string {
	if m != nil {
		return m.CredentialSubject
	}
	return nil
}

func (m *VerifiableCredential) GetProof() map[string]string {
	if m != nil {
		return m.Proof
	}
	return nil
}

func init() {
	proto.RegisterType((*Proof)(nil), "sonrio.sonr.identity.v1.Proof")
	proto.RegisterType((*JSONWebSignature2020Proof)(nil), "sonrio.sonr.identity.v1.JSONWebSignature2020Proof")
	proto.RegisterType((*VerifiableCredential)(nil), "sonrio.sonr.identity.v1.VerifiableCredential")
	proto.RegisterMapType((map[string]string)(nil), "sonrio.sonr.identity.v1.VerifiableCredential.CredentialSubjectEntry")
	proto.RegisterMapType((map[string]string)(nil), "sonrio.sonr.identity.v1.VerifiableCredential.ProofEntry")
}

func init() { proto.RegisterFile("sonr/identity/v1/vc.proto", fileDescriptor_5a30edc5c71454f0) }

var fileDescriptor_5a30edc5c71454f0 = []byte{
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0x8e, 0xe3, 0x26, 0xfd, 0x75, 0xdb, 0x5f, 0x80, 0xa5, 0x0a, 0x6e, 0x0e, 0x56, 0x14, 0x0e,
	0xe4, 0x82, 0xdd, 0x06, 0x84, 0x22, 0x8e, 0x25, 0x5c, 0x90, 0x28, 0x55, 0x82, 0x40, 0xe2, 0x12,
	0x6d, 0xec, 0x69, 0xbb, 0x25, 0xf1, 0x5a, 0xbb, 0x6b, 0x13, 0xbf, 0x05, 0x6f, 0x05, 0xc7, 0x1e,
	0x39, 0xa2, 0xe4, 0x41, 0x40, 0x3b, 0x6b, 0x13, 0x04, 0x2d, 0x52, 0x4f, 0x9e, 0xf9, 0xe6, 0x9b,
	0x3f, 0xdf, 0x8c, 0x97, 0x1c, 0x28, 0x91, 0xc8, 0x90, 0xc7, 0x90, 0x68, 0xae, 0x8b, 0x30, 0x3f,
	0x0a, 0xf3, 0x28, 0x48, 0xa5, 0xd0, 0x82, 0x3e, 0x30, 0x21, 0x2e, 0x02, 0xf3, 0x09, 0x2a, 0x46,
	0x90, 0x1f, 0x75, 0x3a, 0x7f, 0xe5, 0x28, 0xc5, 0x6d, 0x52, 0xef, 0x8b, 0x43, 0x1a, 0xa7, 0x52,
	0x88, 0x33, 0xfa, 0x8c, 0x6c, 0xe9, 0x22, 0x05, 0xcf, 0xe9, 0x3a, 0xfd, 0xd6, 0xa0, 0x17, 0xdc,
	0x50, 0x2d, 0x40, 0xf6, 0xdb, 0x22, 0x85, 0x31, 0xf2, 0xe9, 0x43, 0xf2, 0x7f, 0x6a, 0xa0, 0x69,
	0x9a, 0xc9, 0x54, 0x28, 0xf0, 0xea, 0x5d, 0xa7, 0xbf, 0x33, 0xde, 0x43, 0xf0, 0xd4, 0x62, 0x34,
	0x24, 0xf7, 0x73, 0x90, 0xfc, 0x8c, 0x47, 0x4c, 0x73, 0x91, 0x4c, 0x17, 0xa0, 0x2f, 0x44, 0xec,
	0xb9, 0x48, 0xa5, 0xbf, 0x87, 0x5e, 0x63, 0x84, 0x7a, 0x64, 0x3b, 0x92, 0xc0, 0x34, 0xc4, 0xde,
	0x16, 0x92, 0x2a, 0x97, 0xb6, 0x49, 0x33, 0x16, 0x0b, 0xc6, 0x13, 0xaf, 0x81, 0x81, 0xd2, 0xeb,
	0x45, 0xe4, 0xe0, 0xd5, 0xe4, 0xcd, 0xc9, 0x7b, 0x98, 0x4d, 0xf8, 0x79, 0xc2, 0x74, 0x26, 0x61,
	0x70, 0x38, 0x38, 0xb4, 0xe2, 0x9e, 0x92, 0x06, 0xce, 0x83, 0xea, 0x76, 0x07, 0xfe, 0xbf, 0xd5,
	0x8d, 0x2d, 0x99, 0xde, 0x25, 0xee, 0xe5, 0x27, 0x55, 0x0a, 0x32, 0x66, 0xef, 0x87, 0x4b, 0xf6,
	0xdf, 0xe1, 0xb4, 0x6c, 0x36, 0x87, 0x17, 0x12, 0x30, 0x95, 0xcd, 0x69, 0x8b, 0xd4, 0x79, 0x8c,
	0xd5, 0x77, 0xc6, 0x75, 0x6e, 0xe7, 0x17, 0x89, 0x86, 0xa5, 0xf6, 0xea, 0x5d, 0x17, 0xe7, 0xb7,
	0x2e, 0xa5, 0xe5, 0x9e, 0x5d, 0x84, 0xed, 0x0e, 0xdb, 0xa4, 0xc9, 0x95, 0xca, 0x40, 0x96, 0x62,
	0x4b, 0xcf, 0xec, 0xd6, 0x58, 0x2c, 0x89, 0x60, 0x1a, 0x33, 0x0d, 0xa5, 0xe4, 0xbd, 0x0a, 0x1c,
	0x31, 0x0d, 0xf4, 0x11, 0xb9, 0x03, 0xcb, 0x94, 0x4b, 0xbb, 0x59, 0xa4, 0x35, 0x91, 0xd6, 0xda,
	0xc0, 0x48, 0x54, 0x84, 0x46, 0xbf, 0x26, 0x9e, 0xaa, 0x6c, 0x76, 0x09, 0x91, 0xf6, 0xb6, 0xbb,
	0x6e, 0x7f, 0x77, 0x30, 0xba, 0x71, 0x23, 0xd7, 0xc9, 0x0d, 0x36, 0xe6, 0xc4, 0x96, 0x79, 0x99,
	0x68, 0x59, 0x8c, 0xef, 0x45, 0x7f, 0xe2, 0xf4, 0xa4, 0xda, 0xfc, 0x7f, 0xd8, 0x67, 0x78, 0xbb,
	0x3e, 0x78, 0x0e, 0x5b, 0xdb, 0x96, 0xe9, 0x8c, 0x48, 0xfb, 0xfa, 0xe6, 0xe6, 0x5a, 0x1f, 0xa1,
	0x28, 0x6f, 0x60, 0x4c, 0xba, 0x4f, 0x1a, 0x39, 0x9b, 0x67, 0xd5, 0x2f, 0x69, 0x9d, 0xe7, 0xf5,
	0xa1, 0xd3, 0x19, 0x12, 0xb2, 0x29, 0x7d, 0x9b, 0xcc, 0xe3, 0xe3, 0xaf, 0x2b, 0xdf, 0xb9, 0x5a,
	0xf9, 0xce, 0xf7, 0x95, 0xef, 0x7c, 0x5e, 0xfb, 0xb5, 0xab, 0xb5, 0x5f, 0xfb, 0xb6, 0xf6, 0x6b,
	0x1f, 0xfa, 0xe7, 0x5c, 0x5f, 0x64, 0xb3, 0x20, 0x12, 0x8b, 0xd0, 0xa8, 0x7b, 0xcc, 0x05, 0x7e,
	0xc3, 0xe5, 0xe6, 0xed, 0x99, 0x6b, 0xab, 0x59, 0x13, 0xdf, 0xde, 0x93, 0x9f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xf8, 0x1b, 0x97, 0x96, 0xcd, 0x03, 0x00, 0x00,
}

func (m *Proof) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Proof) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Proof) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Domain) > 0 {
		i -= len(m.Domain)
		copy(dAtA[i:], m.Domain)
		i = encodeVarintVc(dAtA, i, uint64(len(m.Domain)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Created) > 0 {
		i -= len(m.Created)
		copy(dAtA[i:], m.Created)
		i = encodeVarintVc(dAtA, i, uint64(len(m.Created)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.VerificationMethod) > 0 {
		i -= len(m.VerificationMethod)
		copy(dAtA[i:], m.VerificationMethod)
		i = encodeVarintVc(dAtA, i, uint64(len(m.VerificationMethod)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ProofPurpose) > 0 {
		i -= len(m.ProofPurpose)
		copy(dAtA[i:], m.ProofPurpose)
		i = encodeVarintVc(dAtA, i, uint64(len(m.ProofPurpose)))
		i--
		dAtA[i] = 0x12
	}
	if m.Type != 0 {
		i = encodeVarintVc(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *JSONWebSignature2020Proof) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JSONWebSignature2020Proof) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *JSONWebSignature2020Proof) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Jws) > 0 {
		i -= len(m.Jws)
		copy(dAtA[i:], m.Jws)
		i = encodeVarintVc(dAtA, i, uint64(len(m.Jws)))
		i--
		dAtA[i] = 0x12
	}
	if m.Proof != nil {
		{
			size, err := m.Proof.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintVc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *VerifiableCredential) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VerifiableCredential) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VerifiableCredential) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Proof) > 0 {
		for k := range m.Proof {
			v := m.Proof[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintVc(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintVc(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintVc(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.CredentialSubject) > 0 {
		for k := range m.CredentialSubject {
			v := m.CredentialSubject[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintVc(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintVc(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintVc(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.ExpirationDate) > 0 {
		i -= len(m.ExpirationDate)
		copy(dAtA[i:], m.ExpirationDate)
		i = encodeVarintVc(dAtA, i, uint64(len(m.ExpirationDate)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.IssuanceDate) > 0 {
		i -= len(m.IssuanceDate)
		copy(dAtA[i:], m.IssuanceDate)
		i = encodeVarintVc(dAtA, i, uint64(len(m.IssuanceDate)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Issuer) > 0 {
		i -= len(m.Issuer)
		copy(dAtA[i:], m.Issuer)
		i = encodeVarintVc(dAtA, i, uint64(len(m.Issuer)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Type) > 0 {
		for iNdEx := len(m.Type) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Type[iNdEx])
			copy(dAtA[i:], m.Type[iNdEx])
			i = encodeVarintVc(dAtA, i, uint64(len(m.Type[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Context) > 0 {
		for iNdEx := len(m.Context) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Context[iNdEx])
			copy(dAtA[i:], m.Context[iNdEx])
			i = encodeVarintVc(dAtA, i, uint64(len(m.Context[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintVc(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintVc(dAtA []byte, offset int, v uint64) int {
	offset -= sovVc(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Proof) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovVc(uint64(m.Type))
	}
	l = len(m.ProofPurpose)
	if l > 0 {
		n += 1 + l + sovVc(uint64(l))
	}
	l = len(m.VerificationMethod)
	if l > 0 {
		n += 1 + l + sovVc(uint64(l))
	}
	l = len(m.Created)
	if l > 0 {
		n += 1 + l + sovVc(uint64(l))
	}
	l = len(m.Domain)
	if l > 0 {
		n += 1 + l + sovVc(uint64(l))
	}
	return n
}

func (m *JSONWebSignature2020Proof) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Proof != nil {
		l = m.Proof.Size()
		n += 1 + l + sovVc(uint64(l))
	}
	l = len(m.Jws)
	if l > 0 {
		n += 1 + l + sovVc(uint64(l))
	}
	return n
}

func (m *VerifiableCredential) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovVc(uint64(l))
	}
	if len(m.Context) > 0 {
		for _, s := range m.Context {
			l = len(s)
			n += 1 + l + sovVc(uint64(l))
		}
	}
	if len(m.Type) > 0 {
		for _, s := range m.Type {
			l = len(s)
			n += 1 + l + sovVc(uint64(l))
		}
	}
	l = len(m.Issuer)
	if l > 0 {
		n += 1 + l + sovVc(uint64(l))
	}
	l = len(m.IssuanceDate)
	if l > 0 {
		n += 1 + l + sovVc(uint64(l))
	}
	l = len(m.ExpirationDate)
	if l > 0 {
		n += 1 + l + sovVc(uint64(l))
	}
	if len(m.CredentialSubject) > 0 {
		for k, v := range m.CredentialSubject {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovVc(uint64(len(k))) + 1 + len(v) + sovVc(uint64(len(v)))
			n += mapEntrySize + 1 + sovVc(uint64(mapEntrySize))
		}
	}
	if len(m.Proof) > 0 {
		for k, v := range m.Proof {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovVc(uint64(len(k))) + 1 + len(v) + sovVc(uint64(len(v)))
			n += mapEntrySize + 1 + sovVc(uint64(mapEntrySize))
		}
	}
	return n
}

func sovVc(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVc(x uint64) (n int) {
	return sovVc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Proof) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Proof: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Proof: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= ProofType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProofPurpose", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProofPurpose = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VerificationMethod", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VerificationMethod = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Created", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Created = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Domain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Domain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *JSONWebSignature2020Proof) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: JSONWebSignature2020Proof: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JSONWebSignature2020Proof: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthVc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Proof == nil {
				m.Proof = &Proof{}
			}
			if err := m.Proof.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Jws", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Jws = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *VerifiableCredential) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: VerifiableCredential: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VerifiableCredential: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Context", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Context = append(m.Context, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = append(m.Type, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issuer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Issuer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IssuanceDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IssuanceDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpirationDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExpirationDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CredentialSubject", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthVc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CredentialSubject == nil {
				m.CredentialSubject = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowVc
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowVc
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthVc
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthVc
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowVc
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthVc
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthVc
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipVc(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthVc
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.CredentialSubject[mapkey] = mapvalue
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthVc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Proof == nil {
				m.Proof = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowVc
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowVc
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthVc
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthVc
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowVc
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthVc
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthVc
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipVc(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthVc
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Proof[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipVc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVc
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowVc
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowVc
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthVc
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVc
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVc
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVc        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVc          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVc = fmt.Errorf("proto: unexpected end of group")
)
