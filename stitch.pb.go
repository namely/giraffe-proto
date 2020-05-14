// Code generated by protoc-gen-go. DO NOT EDIT.
// source: namely/giraffe/stitch.proto

package giraffe

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type Stitch struct {
	// Service is the fully-qualified package and
	// service name of the service, e.g. namely.profiles.ProfilesService
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	// Method is the method to call within the service, e.g GetProfile
	Method string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	// Request Parameter is the name of the field within the request type
	// to set. The value of this field will be the value of the annotated
	// stitch option.
	RequestParameter string `protobuf:"bytes,3,opt,name=request_parameter,json=requestParameter,proto3" json:"request_parameter,omitempty"`
	// Field name is the name of the field which will be added to the GraphQL
	// schema
	FieldName string `protobuf:"bytes,4,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty"`
	// Hide field will hide the source field from the schema so only the stitched
	// object is present
	HideField bool `protobuf:"varint,5,opt,name=hide_field,json=hideField,proto3" json:"hide_field,omitempty"`
	// Request Quantifier is an optional quantifier for the method payload (e.g.
	// byEmployment). The value of this field will act as a key for the payload {
	// [requestQuantifier]: { [requestParamter]: <value> }}
	RequestQuantifier    string   `protobuf:"bytes,6,opt,name=request_quantifier,json=requestQuantifier,proto3" json:"request_quantifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Stitch) Reset()         { *m = Stitch{} }
func (m *Stitch) String() string { return proto.CompactTextString(m) }
func (*Stitch) ProtoMessage()    {}
func (*Stitch) Descriptor() ([]byte, []int) {
	return fileDescriptor_93b975b9709c8bb6, []int{0}
}

func (m *Stitch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Stitch.Unmarshal(m, b)
}
func (m *Stitch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Stitch.Marshal(b, m, deterministic)
}
func (m *Stitch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stitch.Merge(m, src)
}
func (m *Stitch) XXX_Size() int {
	return xxx_messageInfo_Stitch.Size(m)
}
func (m *Stitch) XXX_DiscardUnknown() {
	xxx_messageInfo_Stitch.DiscardUnknown(m)
}

var xxx_messageInfo_Stitch proto.InternalMessageInfo

func (m *Stitch) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Stitch) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Stitch) GetRequestParameter() string {
	if m != nil {
		return m.RequestParameter
	}
	return ""
}

func (m *Stitch) GetFieldName() string {
	if m != nil {
		return m.FieldName
	}
	return ""
}

func (m *Stitch) GetHideField() bool {
	if m != nil {
		return m.HideField
	}
	return false
}

func (m *Stitch) GetRequestQuantifier() string {
	if m != nil {
		return m.RequestQuantifier
	}
	return ""
}

var E_Stitch = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*Stitch)(nil),
	Field:         165218,
	Name:          "namely.giraffe.stitch",
	Tag:           "bytes,165218,opt,name=stitch",
	Filename:      "namely/giraffe/stitch.proto",
}

func init() {
	proto.RegisterType((*Stitch)(nil), "namely.giraffe.Stitch")
	proto.RegisterExtension(E_Stitch)
}

func init() { proto.RegisterFile("namely/giraffe/stitch.proto", fileDescriptor_93b975b9709c8bb6) }

var fileDescriptor_93b975b9709c8bb6 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0xc6, 0xa9, 0x7f, 0xba, 0x76, 0x04, 0xd1, 0x1c, 0x96, 0xa0, 0x2c, 0x14, 0x4f, 0x05, 0x31,
	0x05, 0xbd, 0x79, 0xf4, 0xe0, 0xd1, 0xd5, 0x7a, 0xf3, 0x52, 0xb2, 0xed, 0xb4, 0x0d, 0xb4, 0x4d,
	0x37, 0x4d, 0x05, 0x5f, 0xc1, 0xe7, 0xf3, 0x09, 0x7c, 0x12, 0xe9, 0x34, 0x11, 0x3c, 0xce, 0xf7,
	0x9b, 0x7c, 0x19, 0x7e, 0x70, 0xd5, 0xcb, 0x0e, 0xdb, 0xcf, 0xb4, 0x56, 0x46, 0x56, 0x15, 0xa6,
	0xa3, 0x55, 0xb6, 0x68, 0xc4, 0x60, 0xb4, 0xd5, 0xec, 0x6c, 0x81, 0xc2, 0xc1, 0xcb, 0xb8, 0xd6,
	0xba, 0x6e, 0x31, 0x25, 0xba, 0x9b, 0xaa, 0xb4, 0xc4, 0xb1, 0x30, 0x6a, 0xb0, 0xda, 0x2c, 0x2f,
	0xae, 0xbf, 0x03, 0x08, 0xdf, 0xa8, 0x82, 0x71, 0x58, 0x8d, 0x68, 0x3e, 0x54, 0x81, 0x3c, 0x88,
	0x83, 0x24, 0xca, 0xfc, 0xc8, 0xd6, 0x10, 0x76, 0x68, 0x1b, 0x5d, 0xf2, 0x03, 0x02, 0x6e, 0x62,
	0x37, 0x70, 0x61, 0x70, 0x3f, 0xe1, 0x68, 0xf3, 0x41, 0x1a, 0xd9, 0xa1, 0x45, 0xc3, 0x0f, 0x69,
	0xe5, 0xdc, 0x81, 0x17, 0x9f, 0xb3, 0x0d, 0x40, 0xa5, 0xb0, 0x2d, 0xf3, 0xf9, 0x46, 0x7e, 0x44,
	0x5b, 0x11, 0x25, 0xcf, 0xb2, 0xc3, 0x19, 0x37, 0xaa, 0xc4, 0x9c, 0x12, 0x7e, 0x1c, 0x07, 0xc9,
	0x49, 0x16, 0xcd, 0xc9, 0xd3, 0x1c, 0xb0, 0x5b, 0x60, 0xfe, 0xab, 0xfd, 0x24, 0x7b, 0xab, 0x2a,
	0x85, 0x86, 0x87, 0xd4, 0xe2, 0x8f, 0x78, 0xfd, 0x03, 0x0f, 0x5b, 0x08, 0x17, 0x31, 0x6c, 0x23,
	0x16, 0x07, 0xc2, 0x3b, 0x10, 0x54, 0xb8, 0x1d, 0xac, 0xd2, 0xfd, 0xc8, 0x7f, 0xbe, 0x20, 0x0e,
	0x92, 0xd3, 0xbb, 0xb5, 0xf8, 0xaf, 0x4e, 0x2c, 0x52, 0x32, 0x57, 0xf3, 0x18, 0xbd, 0xaf, 0x1c,
	0xd9, 0x85, 0xd4, 0x74, 0xff, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xa3, 0xf5, 0xa9, 0x10, 0x8a, 0x01,
	0x00, 0x00,
}
