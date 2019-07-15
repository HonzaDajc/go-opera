// Code generated by protoc-gen-go. DO NOT EDIT.
// source: checkpoint.proto

package wire

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

type Checkpoint struct {
	SuperFrameN          uint64   `protobuf:"varint,1,opt,name=SuperFrameN,json=superFrameN,proto3" json:"SuperFrameN,omitempty"`
	LastFinishedFrameN   uint32   `protobuf:"varint,2,opt,name=LastFinishedFrameN,json=lastFinishedFrameN,proto3" json:"LastFinishedFrameN,omitempty"`
	LastBlockN           uint64   `protobuf:"varint,3,opt,name=LastBlockN,json=lastBlockN,proto3" json:"LastBlockN,omitempty"`
	Genesis              []byte   `protobuf:"bytes,4,opt,name=Genesis,json=genesis,proto3" json:"Genesis,omitempty"`
	TotalCap             uint64   `protobuf:"varint,5,opt,name=TotalCap,json=totalCap,proto3" json:"TotalCap,omitempty"`
	LastConsensusTime    uint64   `protobuf:"varint,6,opt,name=LastConsensusTime,json=lastConsensusTime,proto3" json:"LastConsensusTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Checkpoint) Reset()         { *m = Checkpoint{} }
func (m *Checkpoint) String() string { return proto.CompactTextString(m) }
func (*Checkpoint) ProtoMessage()    {}
func (*Checkpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bab050ffa824783, []int{0}
}

func (m *Checkpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Checkpoint.Unmarshal(m, b)
}
func (m *Checkpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Checkpoint.Marshal(b, m, deterministic)
}
func (m *Checkpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Checkpoint.Merge(m, src)
}
func (m *Checkpoint) XXX_Size() int {
	return xxx_messageInfo_Checkpoint.Size(m)
}
func (m *Checkpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Checkpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Checkpoint proto.InternalMessageInfo

func (m *Checkpoint) GetSuperFrameN() uint64 {
	if m != nil {
		return m.SuperFrameN
	}
	return 0
}

func (m *Checkpoint) GetLastFinishedFrameN() uint32 {
	if m != nil {
		return m.LastFinishedFrameN
	}
	return 0
}

func (m *Checkpoint) GetLastBlockN() uint64 {
	if m != nil {
		return m.LastBlockN
	}
	return 0
}

func (m *Checkpoint) GetGenesis() []byte {
	if m != nil {
		return m.Genesis
	}
	return nil
}

func (m *Checkpoint) GetTotalCap() uint64 {
	if m != nil {
		return m.TotalCap
	}
	return 0
}

func (m *Checkpoint) GetLastConsensusTime() uint64 {
	if m != nil {
		return m.LastConsensusTime
	}
	return 0
}

func init() {
	proto.RegisterType((*Checkpoint)(nil), "wire.Checkpoint")
}

func init() { proto.RegisterFile("checkpoint.proto", fileDescriptor_9bab050ffa824783) }

var fileDescriptor_9bab050ffa824783 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x89, 0xd6, 0xdd, 0x65, 0x56, 0xc1, 0x9d, 0x53, 0xf0, 0x20, 0xc1, 0x53, 0x0f, 0xb2,
	0x17, 0xdf, 0xc0, 0xc2, 0x7a, 0x91, 0x1e, 0x6a, 0x5f, 0x20, 0xd6, 0xc1, 0x86, 0xa6, 0x49, 0xc8,
	0xa4, 0xf8, 0xbe, 0x3e, 0x89, 0x34, 0x96, 0x0a, 0xee, 0xf1, 0xff, 0xbf, 0x6f, 0x7e, 0x18, 0xb8,
	0xed, 0x7a, 0xea, 0x86, 0xe0, 0x8d, 0x4b, 0xc7, 0x10, 0x7d, 0xf2, 0x58, 0x7c, 0x99, 0x48, 0x0f,
	0xdf, 0x02, 0xa0, 0x5a, 0x11, 0x2a, 0xd8, 0xbf, 0x4d, 0x81, 0xe2, 0x29, 0xea, 0x91, 0x6a, 0x29,
	0x94, 0x28, 0x8b, 0x66, 0xcf, 0x7f, 0x15, 0x1e, 0x01, 0x5f, 0x35, 0xa7, 0x93, 0x71, 0x86, 0x7b,
	0xfa, 0x58, 0xc4, 0x0b, 0x25, 0xca, 0x9b, 0x06, 0xed, 0x19, 0xc1, 0x7b, 0x80, 0xd9, 0x7f, 0xb6,
	0xbe, 0x1b, 0x6a, 0x79, 0x99, 0x07, 0xc1, 0xae, 0x0d, 0x4a, 0xd8, 0xbe, 0x90, 0x23, 0x36, 0x2c,
	0x0b, 0x25, 0xca, 0xeb, 0x66, 0xfb, 0xf9, 0x1b, 0xf1, 0x0e, 0x76, 0xad, 0x4f, 0xda, 0x56, 0x3a,
	0xc8, 0xab, 0x7c, 0xb7, 0x4b, 0x4b, 0xc6, 0x47, 0x38, 0xcc, 0xab, 0x95, 0x77, 0x4c, 0x8e, 0x27,
	0x6e, 0xcd, 0x48, 0x72, 0x93, 0xa5, 0x83, 0xfd, 0x0f, 0xde, 0x37, 0xf9, 0xe3, 0xa7, 0x9f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x0b, 0x4f, 0x4a, 0x9d, 0x05, 0x01, 0x00, 0x00,
}
