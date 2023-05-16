// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/privval/msgs.proto

package privval

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	keys "github.com/exfury/grbchain/libs/tendermint/proto/crypto/keys"
	types "github.com/exfury/grbchain/libs/tendermint/proto/types"
	math "math"
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

type RemoteSignerError struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoteSignerError) Reset()         { *m = RemoteSignerError{} }
func (m *RemoteSignerError) String() string { return proto.CompactTextString(m) }
func (*RemoteSignerError) ProtoMessage()    {}
func (*RemoteSignerError) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ec52cc5e378f9a4, []int{0}
}
func (m *RemoteSignerError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoteSignerError.Unmarshal(m, b)
}
func (m *RemoteSignerError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoteSignerError.Marshal(b, m, deterministic)
}
func (m *RemoteSignerError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteSignerError.Merge(m, src)
}
func (m *RemoteSignerError) XXX_Size() int {
	return xxx_messageInfo_RemoteSignerError.Size(m)
}
func (m *RemoteSignerError) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteSignerError.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteSignerError proto.InternalMessageInfo

func (m *RemoteSignerError) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RemoteSignerError) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// PubKeyRequest requests the consensus public key from the remote signer.
type PubKeyRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PubKeyRequest) Reset()         { *m = PubKeyRequest{} }
func (m *PubKeyRequest) String() string { return proto.CompactTextString(m) }
func (*PubKeyRequest) ProtoMessage()    {}
func (*PubKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ec52cc5e378f9a4, []int{1}
}
func (m *PubKeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PubKeyRequest.Unmarshal(m, b)
}
func (m *PubKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PubKeyRequest.Marshal(b, m, deterministic)
}
func (m *PubKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubKeyRequest.Merge(m, src)
}
func (m *PubKeyRequest) XXX_Size() int {
	return xxx_messageInfo_PubKeyRequest.Size(m)
}
func (m *PubKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PubKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PubKeyRequest proto.InternalMessageInfo

// PubKeyResponse is a response message containing the public key.
type PubKeyResponse struct {
	PubKey               keys.PublicKey     `protobuf:"bytes,1,opt,name=pub_key,json=pubKey,proto3" json:"pub_key"`
	Error                *RemoteSignerError `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *PubKeyResponse) Reset()         { *m = PubKeyResponse{} }
func (m *PubKeyResponse) String() string { return proto.CompactTextString(m) }
func (*PubKeyResponse) ProtoMessage()    {}
func (*PubKeyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ec52cc5e378f9a4, []int{2}
}
func (m *PubKeyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PubKeyResponse.Unmarshal(m, b)
}
func (m *PubKeyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PubKeyResponse.Marshal(b, m, deterministic)
}
func (m *PubKeyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubKeyResponse.Merge(m, src)
}
func (m *PubKeyResponse) XXX_Size() int {
	return xxx_messageInfo_PubKeyResponse.Size(m)
}
func (m *PubKeyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PubKeyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PubKeyResponse proto.InternalMessageInfo

func (m *PubKeyResponse) GetPubKey() keys.PublicKey {
	if m != nil {
		return m.PubKey
	}
	return keys.PublicKey{}
}

func (m *PubKeyResponse) GetError() *RemoteSignerError {
	if m != nil {
		return m.Error
	}
	return nil
}

// SignVoteRequest is a request to sign a vote
type SignVoteRequest struct {
	Vote                 types.Vote `protobuf:"bytes,1,opt,name=vote,proto3" json:"vote"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SignVoteRequest) Reset()         { *m = SignVoteRequest{} }
func (m *SignVoteRequest) String() string { return proto.CompactTextString(m) }
func (*SignVoteRequest) ProtoMessage()    {}
func (*SignVoteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ec52cc5e378f9a4, []int{3}
}
func (m *SignVoteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignVoteRequest.Unmarshal(m, b)
}
func (m *SignVoteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignVoteRequest.Marshal(b, m, deterministic)
}
func (m *SignVoteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignVoteRequest.Merge(m, src)
}
func (m *SignVoteRequest) XXX_Size() int {
	return xxx_messageInfo_SignVoteRequest.Size(m)
}
func (m *SignVoteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignVoteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignVoteRequest proto.InternalMessageInfo

func (m *SignVoteRequest) GetVote() types.Vote {
	if m != nil {
		return m.Vote
	}
	return types.Vote{}
}

// SignedVoteResponse is a response containing a signed vote or an error
type SignVoteResponse struct {
	Vote                 types.Vote         `protobuf:"bytes,1,opt,name=vote,proto3" json:"vote"`
	Error                *RemoteSignerError `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SignVoteResponse) Reset()         { *m = SignVoteResponse{} }
func (m *SignVoteResponse) String() string { return proto.CompactTextString(m) }
func (*SignVoteResponse) ProtoMessage()    {}
func (*SignVoteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ec52cc5e378f9a4, []int{4}
}
func (m *SignVoteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignVoteResponse.Unmarshal(m, b)
}
func (m *SignVoteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignVoteResponse.Marshal(b, m, deterministic)
}
func (m *SignVoteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignVoteResponse.Merge(m, src)
}
func (m *SignVoteResponse) XXX_Size() int {
	return xxx_messageInfo_SignVoteResponse.Size(m)
}
func (m *SignVoteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignVoteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignVoteResponse proto.InternalMessageInfo

func (m *SignVoteResponse) GetVote() types.Vote {
	if m != nil {
		return m.Vote
	}
	return types.Vote{}
}

func (m *SignVoteResponse) GetError() *RemoteSignerError {
	if m != nil {
		return m.Error
	}
	return nil
}

// SignProposalRequest is a request to sign a proposal
type SignProposalRequest struct {
	Proposal             types.Proposal `protobuf:"bytes,1,opt,name=proposal,proto3" json:"proposal"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SignProposalRequest) Reset()         { *m = SignProposalRequest{} }
func (m *SignProposalRequest) String() string { return proto.CompactTextString(m) }
func (*SignProposalRequest) ProtoMessage()    {}
func (*SignProposalRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ec52cc5e378f9a4, []int{5}
}
func (m *SignProposalRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignProposalRequest.Unmarshal(m, b)
}
func (m *SignProposalRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignProposalRequest.Marshal(b, m, deterministic)
}
func (m *SignProposalRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignProposalRequest.Merge(m, src)
}
func (m *SignProposalRequest) XXX_Size() int {
	return xxx_messageInfo_SignProposalRequest.Size(m)
}
func (m *SignProposalRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignProposalRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignProposalRequest proto.InternalMessageInfo

func (m *SignProposalRequest) GetProposal() types.Proposal {
	if m != nil {
		return m.Proposal
	}
	return types.Proposal{}
}

// SignedProposalResponse is response containing a signed proposal or an error
type SignedProposalResponse struct {
	Proposal             types.Proposal     `protobuf:"bytes,1,opt,name=proposal,proto3" json:"proposal"`
	Error                *RemoteSignerError `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SignedProposalResponse) Reset()         { *m = SignedProposalResponse{} }
func (m *SignedProposalResponse) String() string { return proto.CompactTextString(m) }
func (*SignedProposalResponse) ProtoMessage()    {}
func (*SignedProposalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ec52cc5e378f9a4, []int{6}
}
func (m *SignedProposalResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedProposalResponse.Unmarshal(m, b)
}
func (m *SignedProposalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedProposalResponse.Marshal(b, m, deterministic)
}
func (m *SignedProposalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedProposalResponse.Merge(m, src)
}
func (m *SignedProposalResponse) XXX_Size() int {
	return xxx_messageInfo_SignedProposalResponse.Size(m)
}
func (m *SignedProposalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedProposalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignedProposalResponse proto.InternalMessageInfo

func (m *SignedProposalResponse) GetProposal() types.Proposal {
	if m != nil {
		return m.Proposal
	}
	return types.Proposal{}
}

func (m *SignedProposalResponse) GetError() *RemoteSignerError {
	if m != nil {
		return m.Error
	}
	return nil
}

// PingRequest is a request to confirm that the connection is alive.
type PingRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ec52cc5e378f9a4, []int{7}
}
func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

// PingResponse is a response to confirm that the connection is alive.
type PingResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ec52cc5e378f9a4, []int{8}
}
func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResponse.Unmarshal(m, b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
}
func (m *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(m, src)
}
func (m *PingResponse) XXX_Size() int {
	return xxx_messageInfo_PingResponse.Size(m)
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RemoteSignerError)(nil), "tendermint.proto.privval.RemoteSignerError")
	proto.RegisterType((*PubKeyRequest)(nil), "tendermint.proto.privval.PubKeyRequest")
	proto.RegisterType((*PubKeyResponse)(nil), "tendermint.proto.privval.PubKeyResponse")
	proto.RegisterType((*SignVoteRequest)(nil), "tendermint.proto.privval.SignVoteRequest")
	proto.RegisterType((*SignVoteResponse)(nil), "tendermint.proto.privval.SignVoteResponse")
	proto.RegisterType((*SignProposalRequest)(nil), "tendermint.proto.privval.SignProposalRequest")
	proto.RegisterType((*SignedProposalResponse)(nil), "tendermint.proto.privval.SignedProposalResponse")
	proto.RegisterType((*PingRequest)(nil), "tendermint.proto.privval.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "tendermint.proto.privval.PingResponse")
}

func init() { proto.RegisterFile("proto/privval/msgs.proto", fileDescriptor_9ec52cc5e378f9a4) }

var fileDescriptor_9ec52cc5e378f9a4 = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xd1, 0xca, 0xd3, 0x30,
	0x1c, 0xc5, 0xad, 0x6c, 0x53, 0xff, 0x75, 0x9b, 0x56, 0xd0, 0x32, 0x14, 0x4b, 0x2f, 0x74, 0x20,
	0xa4, 0x32, 0xc1, 0x7b, 0x07, 0x0a, 0x63, 0x37, 0xa5, 0x82, 0xa0, 0x37, 0x63, 0x6d, 0xff, 0x74,
	0x61, 0x6b, 0x13, 0x93, 0x74, 0xd0, 0x87, 0xf0, 0x09, 0xbc, 0xf0, 0x75, 0x7c, 0x0a, 0x9f, 0x45,
	0x9a, 0xa4, 0x5f, 0xf7, 0x31, 0x76, 0xf3, 0xb1, 0xbb, 0xe4, 0xe4, 0x7f, 0x4e, 0xce, 0x2f, 0xb4,
	0xe0, 0x73, 0xc1, 0x14, 0x8b, 0xb8, 0xa0, 0xc7, 0xe3, 0xf6, 0x10, 0x95, 0xb2, 0x90, 0x44, 0x4b,
	0x9e, 0xaf, 0xb0, 0xca, 0x51, 0x94, 0xb4, 0x52, 0x46, 0x21, 0x76, 0x68, 0xf6, 0x46, 0xed, 0xa8,
	0xc8, 0x37, 0x7c, 0x2b, 0x54, 0x13, 0x19, 0x7f, 0xc1, 0x0a, 0xd6, 0xaf, 0xcc, 0xfc, 0xec, 0x95,
	0x51, 0x32, 0xd1, 0x70, 0xc5, 0xa2, 0x3d, 0x36, 0x32, 0x52, 0x0d, 0x47, 0x7b, 0xc1, 0xec, 0x85,
	0x39, 0xd6, 0xd2, 0xe9, 0x41, 0xb8, 0x82, 0xa7, 0x09, 0x96, 0x4c, 0xe1, 0x57, 0x5a, 0x54, 0x28,
	0x3e, 0x0b, 0xc1, 0x84, 0xe7, 0xc1, 0x20, 0x63, 0x39, 0xfa, 0x4e, 0xe0, 0xcc, 0x87, 0x89, 0x5e,
	0x7b, 0x01, 0xb8, 0x39, 0xca, 0x4c, 0x50, 0xae, 0x28, 0xab, 0xfc, 0xfb, 0x81, 0x33, 0x7f, 0x94,
	0x9c, 0x4a, 0xe1, 0x14, 0xc6, 0x71, 0x9d, 0xae, 0xb1, 0x49, 0xf0, 0x67, 0x8d, 0x52, 0x85, 0xbf,
	0x1d, 0x98, 0x74, 0x8a, 0xe4, 0xac, 0x92, 0xe8, 0x7d, 0x81, 0x07, 0xbc, 0x4e, 0x37, 0x7b, 0x6c,
	0x74, 0xb8, 0xbb, 0x78, 0x4b, 0xce, 0xd0, 0x0d, 0x03, 0x69, 0x19, 0x48, 0x5c, 0xa7, 0x07, 0x9a,
	0xad, 0xb1, 0x59, 0x0e, 0xfe, 0xfe, 0x7b, 0x7d, 0x2f, 0x19, 0x71, 0x9d, 0xe7, 0x7d, 0x82, 0x21,
	0xb6, 0x55, 0x75, 0x0f, 0x77, 0xf1, 0x8e, 0x5c, 0x7a, 0x40, 0x72, 0x46, 0x97, 0x18, 0x67, 0xb8,
	0x82, 0x69, 0xab, 0x7e, 0x63, 0x0a, 0x6d, 0x61, 0xef, 0x23, 0x0c, 0x8e, 0x4c, 0xa1, 0xad, 0xf6,
	0xf2, 0x3c, 0xd4, 0xbc, 0x5c, 0x6b, 0xb1, 0x7d, 0xf4, 0x7c, 0xf8, 0xcb, 0x81, 0x27, 0x7d, 0x96,
	0x45, 0xbd, 0x63, 0xd8, 0x35, 0xd0, 0xbe, 0xc3, 0xb3, 0x56, 0x8d, 0x05, 0xe3, 0x4c, 0x6e, 0x0f,
	0x1d, 0xde, 0x12, 0x1e, 0x72, 0x2b, 0xd9, 0x56, 0xc1, 0xa5, 0x56, 0x9d, 0xd5, 0x36, 0xbb, 0xf1,
	0x85, 0x7f, 0x1c, 0x78, 0xae, 0x6f, 0xcc, 0xfb, 0x74, 0x0b, 0x7c, 0x85, 0xf8, 0x6b, 0xc0, 0x8f,
	0xc1, 0x8d, 0x69, 0x55, 0x74, 0x1f, 0xe1, 0x04, 0x1e, 0x9b, 0xad, 0x69, 0xb9, 0x7c, 0xff, 0x83,
	0x14, 0x54, 0xed, 0xea, 0x94, 0x64, 0xac, 0x8c, 0xfa, 0xf8, 0xd3, 0xe5, 0xad, 0xff, 0x34, 0x1d,
	0xe9, 0xed, 0x87, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x00, 0x13, 0x1c, 0xbf, 0x03, 0x00,
	0x00,
}