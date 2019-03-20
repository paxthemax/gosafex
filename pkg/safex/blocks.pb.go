// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blocks.proto

package safex

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BlockHeader struct {
	Depth                uint64   `protobuf:"varint,1,opt,name=depth,proto3" json:"depth,omitempty"`
	Hash                 string   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	MajorVersion         uint32   `protobuf:"varint,3,opt,name=major_version,json=majorVersion,proto3" json:"major_version,omitempty"`
	MinorVersion         uint32   `protobuf:"varint,4,opt,name=minor_version,json=minorVersion,proto3" json:"minor_version,omitempty"`
	PrevHash             string   `protobuf:"bytes,5,opt,name=prev_hash,json=prevHash,proto3" json:"prev_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockHeader) Reset()         { *m = BlockHeader{} }
func (m *BlockHeader) String() string { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()    {}
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_blocks_cac37187167ffd84, []int{0}
}
func (m *BlockHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockHeader.Unmarshal(m, b)
}
func (m *BlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockHeader.Marshal(b, m, deterministic)
}
func (dst *BlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeader.Merge(dst, src)
}
func (m *BlockHeader) XXX_Size() int {
	return xxx_messageInfo_BlockHeader.Size(m)
}
func (m *BlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeader proto.InternalMessageInfo

func (m *BlockHeader) GetDepth() uint64 {
	if m != nil {
		return m.Depth
	}
	return 0
}

func (m *BlockHeader) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *BlockHeader) GetMajorVersion() uint32 {
	if m != nil {
		return m.MajorVersion
	}
	return 0
}

func (m *BlockHeader) GetMinorVersion() uint32 {
	if m != nil {
		return m.MinorVersion
	}
	return 0
}

func (m *BlockHeader) GetPrevHash() string {
	if m != nil {
		return m.PrevHash
	}
	return ""
}

type Block struct {
	Header               *BlockHeader   `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Txs                  []*Transaction `protobuf:"bytes,2,rep,name=txs,proto3" json:"txs,omitempty"`
	MinerTx              *Transaction   `protobuf:"bytes,3,opt,name=miner_tx,json=minerTx,proto3" json:"miner_tx,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_blocks_cac37187167ffd84, []int{1}
}
func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (dst *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(dst, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetTxs() []*Transaction {
	if m != nil {
		return m.Txs
	}
	return nil
}

func (m *Block) GetMinerTx() *Transaction {
	if m != nil {
		return m.MinerTx
	}
	return nil
}

type Blocks struct {
	Block                []*Block `protobuf:"bytes,1,rep,name=block,proto3" json:"block,omitempty"`
	Status               bool     `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Untrusted            bool     `protobuf:"varint,3,opt,name=untrusted,proto3" json:"untrusted,omitempty"`
	Error                string   `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Blocks) Reset()         { *m = Blocks{} }
func (m *Blocks) String() string { return proto.CompactTextString(m) }
func (*Blocks) ProtoMessage()    {}
func (*Blocks) Descriptor() ([]byte, []int) {
	return fileDescriptor_blocks_cac37187167ffd84, []int{2}
}
func (m *Blocks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blocks.Unmarshal(m, b)
}
func (m *Blocks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blocks.Marshal(b, m, deterministic)
}
func (dst *Blocks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blocks.Merge(dst, src)
}
func (m *Blocks) XXX_Size() int {
	return xxx_messageInfo_Blocks.Size(m)
}
func (m *Blocks) XXX_DiscardUnknown() {
	xxx_messageInfo_Blocks.DiscardUnknown(m)
}

var xxx_messageInfo_Blocks proto.InternalMessageInfo

func (m *Blocks) GetBlock() []*Block {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *Blocks) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *Blocks) GetUntrusted() bool {
	if m != nil {
		return m.Untrusted
	}
	return false
}

func (m *Blocks) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*BlockHeader)(nil), "safex.BlockHeader")
	proto.RegisterType((*Block)(nil), "safex.Block")
	proto.RegisterType((*Blocks)(nil), "safex.Blocks")
}

func init() { proto.RegisterFile("blocks.proto", fileDescriptor_blocks_cac37187167ffd84) }

var fileDescriptor_blocks_cac37187167ffd84 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x4d, 0x4e, 0xc3, 0x30,
	0x10, 0x85, 0xe5, 0xe6, 0x87, 0x64, 0x9a, 0x6e, 0x46, 0x08, 0x45, 0xc0, 0x22, 0x0a, 0x2c, 0x22,
	0x24, 0xb2, 0x80, 0x1b, 0xb0, 0xea, 0xda, 0xaa, 0xd8, 0x46, 0x6e, 0x63, 0x94, 0x00, 0xb5, 0x23,
	0xdb, 0xa9, 0x72, 0x04, 0x4e, 0xc1, 0x59, 0x51, 0x26, 0x41, 0xcd, 0xa2, 0x3b, 0xcf, 0x9b, 0x4f,
	0x7a, 0x6f, 0x9e, 0x21, 0xd9, 0x7f, 0xeb, 0xc3, 0x97, 0x2d, 0x3b, 0xa3, 0x9d, 0xc6, 0xc0, 0x8a,
	0x0f, 0x39, 0xdc, 0xa2, 0x33, 0x42, 0x59, 0x71, 0x70, 0xad, 0x56, 0xf3, 0x2a, 0xff, 0x65, 0xb0,
	0x7e, 0x1b, 0xd9, 0xad, 0x14, 0xb5, 0x34, 0x78, 0x0d, 0x41, 0x2d, 0x3b, 0xd7, 0xa4, 0x2c, 0x63,
	0x85, 0xcf, 0xa7, 0x01, 0x11, 0xfc, 0x46, 0xd8, 0x26, 0x5d, 0x65, 0xac, 0x88, 0x39, 0xbd, 0xf1,
	0x01, 0x36, 0x47, 0xf1, 0xa9, 0x4d, 0x75, 0x92, 0xc6, 0xb6, 0x5a, 0xa5, 0x5e, 0xc6, 0x8a, 0x0d,
	0x4f, 0x48, 0x7c, 0x9f, 0x34, 0x82, 0x5a, 0xb5, 0x80, 0xfc, 0x19, 0x1a, 0xc5, 0x7f, 0xe8, 0x0e,
	0xe2, 0xce, 0xc8, 0x53, 0x45, 0x16, 0x01, 0x59, 0x44, 0xa3, 0xb0, 0x15, 0xb6, 0xc9, 0x7f, 0x18,
	0x04, 0x14, 0x10, 0x9f, 0x20, 0x6c, 0x28, 0x24, 0x65, 0x5b, 0xbf, 0x60, 0x49, 0x67, 0x95, 0x8b,
	0xf8, 0x7c, 0x26, 0xf0, 0x11, 0x3c, 0x37, 0xd8, 0x74, 0x95, 0x79, 0x0b, 0x70, 0x77, 0x3e, 0x9f,
	0x8f, 0x6b, 0x7c, 0x86, 0xe8, 0xd8, 0x2a, 0x69, 0x2a, 0x37, 0x50, 0xfa, 0xcb, 0xe8, 0x15, 0x31,
	0xbb, 0x21, 0x1f, 0x20, 0x24, 0x2f, 0x8b, 0x39, 0x04, 0x54, 0x70, 0xca, 0xc8, 0x20, 0x59, 0x26,
	0xe1, 0xd3, 0x0a, 0x6f, 0x20, 0xb4, 0x4e, 0xb8, 0xde, 0x52, 0x6b, 0x11, 0x9f, 0x27, 0xbc, 0x87,
	0xb8, 0x57, 0xce, 0xf4, 0xd6, 0xc9, 0x9a, 0x5c, 0x23, 0x7e, 0x16, 0xc6, 0xfe, 0xa5, 0x31, 0xda,
	0x50, 0x51, 0x31, 0x9f, 0x86, 0x7d, 0x48, 0x9f, 0xf5, 0xfa, 0x17, 0x00, 0x00, 0xff, 0xff, 0x7e,
	0x5b, 0x01, 0x51, 0xd7, 0x01, 0x00, 0x00,
}
