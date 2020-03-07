// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core/node_accessor_protocol.proto

package core

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Carrier struct {
	Packet               []*Packet `protobuf:"bytes,1,rep,name=packet,proto3" json:"packet,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Carrier) Reset()         { *m = Carrier{} }
func (m *Carrier) String() string { return proto.CompactTextString(m) }
func (*Carrier) ProtoMessage()    {}
func (*Carrier) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6da5911fa8c514f, []int{0}
}

func (m *Carrier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Carrier.Unmarshal(m, b)
}
func (m *Carrier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Carrier.Marshal(b, m, deterministic)
}
func (m *Carrier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Carrier.Merge(m, src)
}
func (m *Carrier) XXX_Size() int {
	return xxx_messageInfo_Carrier.Size(m)
}
func (m *Carrier) XXX_DiscardUnknown() {
	xxx_messageInfo_Carrier.DiscardUnknown(m)
}

var xxx_messageInfo_Carrier proto.InternalMessageInfo

func (m *Carrier) GetPacket() []*Packet {
	if m != nil {
		return m.Packet
	}
	return nil
}

type Head struct {
	DstNid               *NodeID  `protobuf:"bytes,1,opt,name=dst_nid,json=dstNid,proto3" json:"dst_nid,omitempty"`
	SrcNid               *NodeID  `protobuf:"bytes,2,opt,name=src_nid,json=srcNid,proto3" json:"src_nid,omitempty"`
	Mode                 uint32   `protobuf:"varint,3,opt,name=mode,proto3" json:"mode,omitempty"`
	Channel              uint32   `protobuf:"varint,4,opt,name=channel,proto3" json:"channel,omitempty"`
	ModuleChannel        uint32   `protobuf:"varint,6,opt,name=module_channel,json=moduleChannel,proto3" json:"module_channel,omitempty"`
	CommandId            uint32   `protobuf:"varint,5,opt,name=command_id,json=commandId,proto3" json:"command_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Head) Reset()         { *m = Head{} }
func (m *Head) String() string { return proto.CompactTextString(m) }
func (*Head) ProtoMessage()    {}
func (*Head) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6da5911fa8c514f, []int{1}
}

func (m *Head) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Head.Unmarshal(m, b)
}
func (m *Head) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Head.Marshal(b, m, deterministic)
}
func (m *Head) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Head.Merge(m, src)
}
func (m *Head) XXX_Size() int {
	return xxx_messageInfo_Head.Size(m)
}
func (m *Head) XXX_DiscardUnknown() {
	xxx_messageInfo_Head.DiscardUnknown(m)
}

var xxx_messageInfo_Head proto.InternalMessageInfo

func (m *Head) GetDstNid() *NodeID {
	if m != nil {
		return m.DstNid
	}
	return nil
}

func (m *Head) GetSrcNid() *NodeID {
	if m != nil {
		return m.SrcNid
	}
	return nil
}

func (m *Head) GetMode() uint32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

func (m *Head) GetChannel() uint32 {
	if m != nil {
		return m.Channel
	}
	return 0
}

func (m *Head) GetModuleChannel() uint32 {
	if m != nil {
		return m.ModuleChannel
	}
	return 0
}

func (m *Head) GetCommandId() uint32 {
	if m != nil {
		return m.CommandId
	}
	return 0
}

type Packet struct {
	// Enable head if index is 0.
	Head *Head `protobuf:"bytes,1,opt,name=head,proto3" json:"head,omitempty"`
	// ID is common to sequence of packets.
	Id uint32 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// index decreses by 1 in a sequance of packets.
	Index                uint32   `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
	Content              []byte   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Packet) Reset()         { *m = Packet{} }
func (m *Packet) String() string { return proto.CompactTextString(m) }
func (*Packet) ProtoMessage()    {}
func (*Packet) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6da5911fa8c514f, []int{2}
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

func (m *Packet) GetHead() *Head {
	if m != nil {
		return m.Head
	}
	return nil
}

func (m *Packet) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Packet) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Packet) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type ICE struct {
	LocalNid             *NodeID  `protobuf:"bytes,1,opt,name=local_nid,json=localNid,proto3" json:"local_nid,omitempty"`
	RemoteNid            *NodeID  `protobuf:"bytes,2,opt,name=remote_nid,json=remoteNid,proto3" json:"remote_nid,omitempty"`
	Ice                  string   `protobuf:"bytes,3,opt,name=ice,proto3" json:"ice,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ICE) Reset()         { *m = ICE{} }
func (m *ICE) String() string { return proto.CompactTextString(m) }
func (*ICE) ProtoMessage()    {}
func (*ICE) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6da5911fa8c514f, []int{3}
}

func (m *ICE) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ICE.Unmarshal(m, b)
}
func (m *ICE) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ICE.Marshal(b, m, deterministic)
}
func (m *ICE) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ICE.Merge(m, src)
}
func (m *ICE) XXX_Size() int {
	return xxx_messageInfo_ICE.Size(m)
}
func (m *ICE) XXX_DiscardUnknown() {
	xxx_messageInfo_ICE.DiscardUnknown(m)
}

var xxx_messageInfo_ICE proto.InternalMessageInfo

func (m *ICE) GetLocalNid() *NodeID {
	if m != nil {
		return m.LocalNid
	}
	return nil
}

func (m *ICE) GetRemoteNid() *NodeID {
	if m != nil {
		return m.RemoteNid
	}
	return nil
}

func (m *ICE) GetIce() string {
	if m != nil {
		return m.Ice
	}
	return ""
}

type Offer struct {
	PrimeNid             *NodeID  `protobuf:"bytes,1,opt,name=prime_nid,json=primeNid,proto3" json:"prime_nid,omitempty"`
	SecondNid            *NodeID  `protobuf:"bytes,2,opt,name=second_nid,json=secondNid,proto3" json:"second_nid,omitempty"`
	Sdp                  string   `protobuf:"bytes,3,opt,name=sdp,proto3" json:"sdp,omitempty"`
	Type                 uint32   `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Offer) Reset()         { *m = Offer{} }
func (m *Offer) String() string { return proto.CompactTextString(m) }
func (*Offer) ProtoMessage()    {}
func (*Offer) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6da5911fa8c514f, []int{4}
}

func (m *Offer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Offer.Unmarshal(m, b)
}
func (m *Offer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Offer.Marshal(b, m, deterministic)
}
func (m *Offer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Offer.Merge(m, src)
}
func (m *Offer) XXX_Size() int {
	return xxx_messageInfo_Offer.Size(m)
}
func (m *Offer) XXX_DiscardUnknown() {
	xxx_messageInfo_Offer.DiscardUnknown(m)
}

var xxx_messageInfo_Offer proto.InternalMessageInfo

func (m *Offer) GetPrimeNid() *NodeID {
	if m != nil {
		return m.PrimeNid
	}
	return nil
}

func (m *Offer) GetSecondNid() *NodeID {
	if m != nil {
		return m.SecondNid
	}
	return nil
}

func (m *Offer) GetSdp() string {
	if m != nil {
		return m.Sdp
	}
	return ""
}

func (m *Offer) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

type OfferSuccess struct {
	Status               uint32   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	SecondNid            *NodeID  `protobuf:"bytes,2,opt,name=second_nid,json=secondNid,proto3" json:"second_nid,omitempty"`
	Sdp                  string   `protobuf:"bytes,3,opt,name=sdp,proto3" json:"sdp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OfferSuccess) Reset()         { *m = OfferSuccess{} }
func (m *OfferSuccess) String() string { return proto.CompactTextString(m) }
func (*OfferSuccess) ProtoMessage()    {}
func (*OfferSuccess) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6da5911fa8c514f, []int{5}
}

func (m *OfferSuccess) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OfferSuccess.Unmarshal(m, b)
}
func (m *OfferSuccess) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OfferSuccess.Marshal(b, m, deterministic)
}
func (m *OfferSuccess) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OfferSuccess.Merge(m, src)
}
func (m *OfferSuccess) XXX_Size() int {
	return xxx_messageInfo_OfferSuccess.Size(m)
}
func (m *OfferSuccess) XXX_DiscardUnknown() {
	xxx_messageInfo_OfferSuccess.DiscardUnknown(m)
}

var xxx_messageInfo_OfferSuccess proto.InternalMessageInfo

func (m *OfferSuccess) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *OfferSuccess) GetSecondNid() *NodeID {
	if m != nil {
		return m.SecondNid
	}
	return nil
}

func (m *OfferSuccess) GetSdp() string {
	if m != nil {
		return m.Sdp
	}
	return ""
}

type OfferFailure struct {
	Status               uint32   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	PrimeNid             *NodeID  `protobuf:"bytes,2,opt,name=prime_nid,json=primeNid,proto3" json:"prime_nid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OfferFailure) Reset()         { *m = OfferFailure{} }
func (m *OfferFailure) String() string { return proto.CompactTextString(m) }
func (*OfferFailure) ProtoMessage()    {}
func (*OfferFailure) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6da5911fa8c514f, []int{6}
}

func (m *OfferFailure) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OfferFailure.Unmarshal(m, b)
}
func (m *OfferFailure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OfferFailure.Marshal(b, m, deterministic)
}
func (m *OfferFailure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OfferFailure.Merge(m, src)
}
func (m *OfferFailure) XXX_Size() int {
	return xxx_messageInfo_OfferFailure.Size(m)
}
func (m *OfferFailure) XXX_DiscardUnknown() {
	xxx_messageInfo_OfferFailure.DiscardUnknown(m)
}

var xxx_messageInfo_OfferFailure proto.InternalMessageInfo

func (m *OfferFailure) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *OfferFailure) GetPrimeNid() *NodeID {
	if m != nil {
		return m.PrimeNid
	}
	return nil
}

func init() {
	proto.RegisterType((*Carrier)(nil), "colonio.NodeAccessorProtocol.Carrier")
	proto.RegisterType((*Head)(nil), "colonio.NodeAccessorProtocol.Head")
	proto.RegisterType((*Packet)(nil), "colonio.NodeAccessorProtocol.Packet")
	proto.RegisterType((*ICE)(nil), "colonio.NodeAccessorProtocol.ICE")
	proto.RegisterType((*Offer)(nil), "colonio.NodeAccessorProtocol.Offer")
	proto.RegisterType((*OfferSuccess)(nil), "colonio.NodeAccessorProtocol.OfferSuccess")
	proto.RegisterType((*OfferFailure)(nil), "colonio.NodeAccessorProtocol.OfferFailure")
}

func init() { proto.RegisterFile("core/node_accessor_protocol.proto", fileDescriptor_d6da5911fa8c514f) }

var fileDescriptor_d6da5911fa8c514f = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcf, 0x8a, 0x13, 0x41,
	0x10, 0xc6, 0x99, 0xfc, 0x99, 0x98, 0xda, 0xcd, 0x2a, 0xcd, 0x22, 0x83, 0x28, 0xc4, 0x41, 0x21,
	0x17, 0x23, 0xee, 0x82, 0x27, 0x2f, 0x1a, 0xff, 0xe5, 0xb2, 0x2e, 0xe3, 0x49, 0x2f, 0x43, 0xdb,
	0x5d, 0xcb, 0x36, 0xce, 0x74, 0x0d, 0xdd, 0x1d, 0xd0, 0x93, 0xbe, 0x84, 0xef, 0xe6, 0xe3, 0x48,
	0xd7, 0xf4, 0x08, 0x1e, 0x4c, 0x3c, 0x78, 0x09, 0x55, 0x9d, 0x5f, 0x57, 0x7d, 0xdf, 0xd7, 0x0c,
	0xdc, 0x57, 0xe4, 0xf0, 0xb1, 0x25, 0x8d, 0xb5, 0x54, 0x0a, 0xbd, 0x27, 0x57, 0x77, 0x8e, 0x02,
	0x29, 0x6a, 0xd6, 0x5c, 0x88, 0xbb, 0x8a, 0x1a, 0xb2, 0x86, 0xd6, 0x17, 0xa4, 0xf1, 0x79, 0x82,
	0x2e, 0x13, 0x73, 0xe7, 0x26, 0x0f, 0x88, 0x3f, 0x3d, 0x5e, 0xbe, 0x81, 0xd9, 0x46, 0x3a, 0x67,
	0xd0, 0x89, 0x67, 0x90, 0x77, 0x52, 0x7d, 0xc6, 0x50, 0x64, 0xcb, 0xf1, 0xea, 0xe8, 0xec, 0xc1,
	0x7a, 0xdf, 0xa8, 0xf5, 0x25, 0xb3, 0x55, 0xba, 0x53, 0xfe, 0xcc, 0x60, 0xf2, 0x16, 0xa5, 0x16,
	0x8f, 0x60, 0xa6, 0x7d, 0xa8, 0xad, 0xd1, 0x45, 0xb6, 0xcc, 0x56, 0x47, 0x67, 0xa7, 0xbf, 0xe7,
	0xf0, 0xde, 0x38, 0x6c, 0xfb, 0xb2, 0xca, 0xb5, 0x0f, 0x17, 0x86, 0x71, 0xef, 0x14, 0xe3, 0xa3,
	0x7d, 0xb8, 0x77, 0x2a, 0xe2, 0x02, 0x26, 0x2d, 0x69, 0x2c, 0xc6, 0xcb, 0x6c, 0xb5, 0xa8, 0xb8,
	0x16, 0x05, 0xcc, 0xd4, 0xb5, 0xb4, 0x16, 0x9b, 0x62, 0xc2, 0xc7, 0x43, 0x2b, 0x1e, 0xc2, 0x49,
	0x4b, 0x7a, 0xd7, 0x60, 0x3d, 0x00, 0x39, 0x03, 0x8b, 0xfe, 0x74, 0x93, 0xb0, 0x7b, 0x00, 0x8a,
	0xda, 0x56, 0x5a, 0x5d, 0x1b, 0x5d, 0x4c, 0x19, 0x99, 0xa7, 0x93, 0xad, 0x2e, 0xbf, 0x67, 0x90,
	0xf7, 0x6e, 0xc5, 0x53, 0x98, 0x5c, 0xa3, 0x1c, 0x9c, 0x95, 0xfb, 0x13, 0x8a, 0x71, 0x54, 0xcc,
	0x8b, 0x13, 0x18, 0x25, 0x83, 0x8b, 0x6a, 0x64, 0xb4, 0x38, 0x85, 0xa9, 0xb1, 0x1a, 0xbf, 0x24,
	0x1f, 0x7d, 0xc3, 0x46, 0xc8, 0x06, 0xb4, 0x81, 0x8d, 0x1c, 0x57, 0x43, 0x5b, 0x7e, 0x83, 0xf1,
	0x76, 0xf3, 0x4a, 0x3c, 0x81, 0x79, 0x43, 0x4a, 0x36, 0x07, 0xd3, 0xbd, 0xc1, 0x58, 0x0c, 0xec,
	0x1c, 0xc0, 0x61, 0x4b, 0x01, 0x0f, 0x46, 0x3c, 0xef, 0xb9, 0x78, 0xe9, 0x16, 0x8c, 0x8d, 0xea,
	0x43, 0x9e, 0x57, 0xb1, 0x2c, 0x7f, 0x64, 0x30, 0x7d, 0x77, 0x75, 0x85, 0x2e, 0x6a, 0xe8, 0x9c,
	0x69, 0xf1, 0xb0, 0x06, 0xc6, 0x92, 0x06, 0x8f, 0x8a, 0xac, 0x3e, 0xac, 0xa1, 0xe7, 0x92, 0x06,
	0xaf, 0xbb, 0x41, 0x83, 0xd7, 0x5d, 0x7c, 0xfb, 0xf0, 0xb5, 0xc3, 0xf4, 0xc8, 0x5c, 0x97, 0x2d,
	0x1c, 0xb3, 0xac, 0xf7, 0x3b, 0x4e, 0x5f, 0xdc, 0x86, 0xdc, 0x07, 0x19, 0x76, 0x9e, 0xa5, 0x2d,
	0xaa, 0xd4, 0xfd, 0x27, 0x09, 0xe5, 0x87, 0xb4, 0xee, 0xb5, 0x34, 0xcd, 0xce, 0xe1, 0x5f, 0xd7,
	0xfd, 0x11, 0xd2, 0xe8, 0x5f, 0x42, 0x7a, 0x91, 0x7f, 0x9c, 0xc4, 0x3f, 0x3e, 0xe5, 0xfc, 0x61,
	0x9e, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x01, 0xa3, 0xa0, 0x4f, 0xec, 0x03, 0x00, 0x00,
}