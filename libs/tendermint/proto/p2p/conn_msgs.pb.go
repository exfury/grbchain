// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/p2p/conn_msgs.proto

package p2p

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type PacketPing struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PacketPing) Reset()         { *m = PacketPing{} }
func (m *PacketPing) String() string { return proto.CompactTextString(m) }
func (*PacketPing) ProtoMessage()    {}
func (*PacketPing) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c680f0b24d73fe7, []int{0}
}
func (m *PacketPing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PacketPing.Unmarshal(m, b)
}
func (m *PacketPing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PacketPing.Marshal(b, m, deterministic)
}
func (m *PacketPing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PacketPing.Merge(m, src)
}
func (m *PacketPing) XXX_Size() int {
	return xxx_messageInfo_PacketPing.Size(m)
}
func (m *PacketPing) XXX_DiscardUnknown() {
	xxx_messageInfo_PacketPing.DiscardUnknown(m)
}

var xxx_messageInfo_PacketPing proto.InternalMessageInfo

type PacketPong struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PacketPong) Reset()         { *m = PacketPong{} }
func (m *PacketPong) String() string { return proto.CompactTextString(m) }
func (*PacketPong) ProtoMessage()    {}
func (*PacketPong) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c680f0b24d73fe7, []int{1}
}
func (m *PacketPong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PacketPong.Unmarshal(m, b)
}
func (m *PacketPong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PacketPong.Marshal(b, m, deterministic)
}
func (m *PacketPong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PacketPong.Merge(m, src)
}
func (m *PacketPong) XXX_Size() int {
	return xxx_messageInfo_PacketPong.Size(m)
}
func (m *PacketPong) XXX_DiscardUnknown() {
	xxx_messageInfo_PacketPong.DiscardUnknown(m)
}

var xxx_messageInfo_PacketPong proto.InternalMessageInfo

type PacketMsg struct {
	ChannelID            int32    `protobuf:"varint,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	EOF                  int32    `protobuf:"varint,2,opt,name=eof,proto3" json:"eof,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PacketMsg) Reset()         { *m = PacketMsg{} }
func (m *PacketMsg) String() string { return proto.CompactTextString(m) }
func (*PacketMsg) ProtoMessage()    {}
func (*PacketMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c680f0b24d73fe7, []int{2}
}
func (m *PacketMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PacketMsg.Unmarshal(m, b)
}
func (m *PacketMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PacketMsg.Marshal(b, m, deterministic)
}
func (m *PacketMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PacketMsg.Merge(m, src)
}
func (m *PacketMsg) XXX_Size() int {
	return xxx_messageInfo_PacketMsg.Size(m)
}
func (m *PacketMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_PacketMsg.DiscardUnknown(m)
}

var xxx_messageInfo_PacketMsg proto.InternalMessageInfo

func (m *PacketMsg) GetChannelID() int32 {
	if m != nil {
		return m.ChannelID
	}
	return 0
}

func (m *PacketMsg) GetEOF() int32 {
	if m != nil {
		return m.EOF
	}
	return 0
}

func (m *PacketMsg) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Packet struct {
	// Types that are valid to be assigned to Sum:
	//	*Packet_PacketPing
	//	*Packet_PacketPong
	//	*Packet_PacketMsg
	Sum                  isPacket_Sum `protobuf_oneof:"sum"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Packet) Reset()         { *m = Packet{} }
func (m *Packet) String() string { return proto.CompactTextString(m) }
func (*Packet) ProtoMessage()    {}
func (*Packet) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c680f0b24d73fe7, []int{3}
}
func (m *Packet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Packet.Unmarshal(m, b)
}
func (m *Packet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Packet.Marshal(b, m, deterministic)
}
func (m *Packet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Packet.Merge(m, src)
}
func (m *Packet) XXX_Size() int {
	return xxx_messageInfo_Packet.Size(m)
}
func (m *Packet) XXX_DiscardUnknown() {
	xxx_messageInfo_Packet.DiscardUnknown(m)
}

var xxx_messageInfo_Packet proto.InternalMessageInfo

type isPacket_Sum interface {
	isPacket_Sum()
}

type Packet_PacketPing struct {
	PacketPing *PacketPing `protobuf:"bytes,1,opt,name=packet_ping,json=packetPing,proto3,oneof" json:"packet_ping,omitempty"`
}
type Packet_PacketPong struct {
	PacketPong *PacketPong `protobuf:"bytes,2,opt,name=packet_pong,json=packetPong,proto3,oneof" json:"packet_pong,omitempty"`
}
type Packet_PacketMsg struct {
	PacketMsg *PacketMsg `protobuf:"bytes,3,opt,name=packet_msg,json=packetMsg,proto3,oneof" json:"packet_msg,omitempty"`
}

func (*Packet_PacketPing) isPacket_Sum() {}
func (*Packet_PacketPong) isPacket_Sum() {}
func (*Packet_PacketMsg) isPacket_Sum()  {}

func (m *Packet) GetSum() isPacket_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *Packet) GetPacketPing() *PacketPing {
	if x, ok := m.GetSum().(*Packet_PacketPing); ok {
		return x.PacketPing
	}
	return nil
}

func (m *Packet) GetPacketPong() *PacketPong {
	if x, ok := m.GetSum().(*Packet_PacketPong); ok {
		return x.PacketPong
	}
	return nil
}

func (m *Packet) GetPacketMsg() *PacketMsg {
	if x, ok := m.GetSum().(*Packet_PacketMsg); ok {
		return x.PacketMsg
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Packet) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Packet_PacketPing)(nil),
		(*Packet_PacketPong)(nil),
		(*Packet_PacketMsg)(nil),
	}
}

func init() {
	proto.RegisterType((*PacketPing)(nil), "tendermint.proto.p2p.PacketPing")
	proto.RegisterType((*PacketPong)(nil), "tendermint.proto.p2p.PacketPong")
	proto.RegisterType((*PacketMsg)(nil), "tendermint.proto.p2p.PacketMsg")
	proto.RegisterType((*Packet)(nil), "tendermint.proto.p2p.Packet")
}

func init() { proto.RegisterFile("proto/p2p/conn_msgs.proto", fileDescriptor_8c680f0b24d73fe7) }

var fileDescriptor_8c680f0b24d73fe7 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4f, 0x83, 0x30,
	0x14, 0xc6, 0x45, 0xdc, 0x0c, 0x8f, 0x79, 0x69, 0x3c, 0x30, 0x2f, 0x10, 0x0e, 0x66, 0x31, 0x0b,
	0x24, 0xf8, 0x0f, 0x18, 0xa6, 0xc6, 0x1d, 0x16, 0x17, 0x8e, 0x5e, 0x08, 0x03, 0x2c, 0x8d, 0xd2,
	0xd7, 0x40, 0x77, 0xf0, 0x6f, 0x35, 0xd9, 0x61, 0x7f, 0x89, 0xa1, 0x9d, 0x03, 0x13, 0xa3, 0xb7,
	0xef, 0xfb, 0xf2, 0xfa, 0x7b, 0x5f, 0x5b, 0x98, 0x8a, 0x06, 0x25, 0x86, 0x22, 0x12, 0x61, 0x8e,
	0x9c, 0xa7, 0x75, 0x4b, 0xdb, 0x40, 0x65, 0xe4, 0x52, 0x96, 0xbc, 0x28, 0x9b, 0x9a, 0x71, 0xa9,
	0x93, 0x40, 0x44, 0xe2, 0xea, 0x5a, 0x56, 0xac, 0x29, 0x52, 0x91, 0x35, 0xf2, 0x23, 0xd4, 0x87,
	0x29, 0x52, 0xec, 0x95, 0x9e, 0xf5, 0x27, 0x00, 0xeb, 0x2c, 0x7f, 0x2b, 0xe5, 0x9a, 0x71, 0x3a,
	0x70, 0xc8, 0xa9, 0x5f, 0x81, 0xa5, 0xdd, 0xaa, 0xa5, 0x64, 0x0e, 0x90, 0x57, 0x19, 0xe7, 0xe5,
	0x7b, 0xca, 0x0a, 0xc7, 0xf0, 0x8c, 0xd9, 0x28, 0xbe, 0xd8, 0xef, 0x5c, 0x6b, 0xa1, 0xd3, 0xe5,
	0x7d, 0x62, 0x1d, 0x06, 0x96, 0x05, 0x99, 0x82, 0x59, 0xe2, 0xab, 0x73, 0xaa, 0xc6, 0xce, 0xf7,
	0x3b, 0xd7, 0x7c, 0x78, 0x7e, 0x4c, 0xba, 0x8c, 0x10, 0x38, 0x2b, 0x32, 0x99, 0x39, 0xa6, 0x67,
	0xcc, 0x26, 0x89, 0xd2, 0xfe, 0xa7, 0x01, 0x63, 0xbd, 0x8a, 0x2c, 0xc0, 0x16, 0x4a, 0xa5, 0x82,
	0x71, 0xaa, 0x16, 0xd9, 0x91, 0x17, 0xfc, 0x76, 0xc9, 0xa0, 0x6f, 0xfe, 0x74, 0x92, 0x80, 0x38,
	0xba, 0x21, 0x04, 0x39, 0x55, 0x35, 0xfe, 0x83, 0xe0, 0x0f, 0x08, 0x72, 0x4a, 0xee, 0xe0, 0xe0,
	0xba, 0xd7, 0x56, 0x75, 0xed, 0xc8, 0xfd, 0x8b, 0xb1, 0x6a, 0x3b, 0x84, 0x25, 0xbe, 0x4d, 0x3c,
	0x02, 0xb3, 0xdd, 0xd6, 0xf1, 0xfc, 0xe5, 0x86, 0x32, 0x59, 0x6d, 0x37, 0x41, 0x8e, 0x75, 0xd8,
	0x03, 0x86, 0xf2, 0xf8, 0xbf, 0x9b, 0xb1, 0x92, 0xb7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x87,
	0x8c, 0x0a, 0x5f, 0xf3, 0x01, 0x00, 0x00,
}
