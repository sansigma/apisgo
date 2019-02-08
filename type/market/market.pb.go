// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sansigmaprotos/type/market.proto

package market

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

type Market struct {
	Symbol               string   `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Exchange             string   `protobuf:"bytes,2,opt,name=exchange,proto3" json:"exchange,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Market) Reset()         { *m = Market{} }
func (m *Market) String() string { return proto.CompactTextString(m) }
func (*Market) ProtoMessage()    {}
func (*Market) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e93b1714c6c096d, []int{0}
}

func (m *Market) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Market.Unmarshal(m, b)
}
func (m *Market) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Market.Marshal(b, m, deterministic)
}
func (m *Market) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Market.Merge(m, src)
}
func (m *Market) XXX_Size() int {
	return xxx_messageInfo_Market.Size(m)
}
func (m *Market) XXX_DiscardUnknown() {
	xxx_messageInfo_Market.DiscardUnknown(m)
}

var xxx_messageInfo_Market proto.InternalMessageInfo

func (m *Market) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *Market) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func init() {
	proto.RegisterType((*Market)(nil), "ssigma.type.Market")
}

func init() { proto.RegisterFile("sansigmaprotos/type/market.proto", fileDescriptor_7e93b1714c6c096d) }

var fileDescriptor_7e93b1714c6c096d = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0x4e, 0xcc, 0x2b,
	0xce, 0x4c, 0xcf, 0x4d, 0x2c, 0x28, 0xca, 0x2f, 0xc9, 0x2f, 0xd6, 0x2f, 0xa9, 0x2c, 0x48, 0xd5,
	0xcf, 0x4d, 0x2c, 0xca, 0x4e, 0x2d, 0xd1, 0x03, 0x0b, 0x09, 0x71, 0x17, 0x83, 0xe5, 0xf5, 0x40,
	0x32, 0x4a, 0x36, 0x5c, 0x6c, 0xbe, 0x60, 0x49, 0x21, 0x31, 0x2e, 0xb6, 0xe2, 0xca, 0xdc, 0xa4,
	0xfc, 0x1c, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x28, 0x4f, 0x48, 0x8a, 0x8b, 0x23, 0xb5,
	0x22, 0x39, 0x23, 0x31, 0x2f, 0x3d, 0x55, 0x82, 0x09, 0x2c, 0x03, 0xe7, 0x3b, 0xe9, 0x47, 0xe9,
	0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xc3, 0x6c, 0xd6, 0x4f, 0x2c,
	0xc8, 0x2c, 0x4e, 0xcf, 0x47, 0xb6, 0xda, 0x1a, 0x42, 0x25, 0xb1, 0x81, 0x9d, 0x60, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x8e, 0x2d, 0xe1, 0x66, 0xa6, 0x00, 0x00, 0x00,
}
