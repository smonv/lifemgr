// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nutrient.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	nutrient.proto
	service.proto

It has these top-level messages:
	Nutrient
	GetNutrientsResponse
*/
package api

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

type Nutrient struct {
	Id    string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Unit  string `protobuf:"bytes,3,opt,name=unit" json:"unit,omitempty"`
	Group string `protobuf:"bytes,4,opt,name=group" json:"group,omitempty"`
}

func (m *Nutrient) Reset()                    { *m = Nutrient{} }
func (m *Nutrient) String() string            { return proto.CompactTextString(m) }
func (*Nutrient) ProtoMessage()               {}
func (*Nutrient) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Nutrient) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Nutrient) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Nutrient) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *Nutrient) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func init() {
	proto.RegisterType((*Nutrient)(nil), "Nutrient")
}

func init() { proto.RegisterFile("nutrient.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 115 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x2b, 0x2d, 0x29,
	0xca, 0x4c, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x8a, 0xe0, 0xe2, 0xf0, 0x83,
	0x8a, 0x08, 0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x31, 0x65,
	0xa6, 0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x81, 0x45, 0xc0, 0x6c, 0x90,
	0x58, 0x69, 0x5e, 0x66, 0x89, 0x04, 0x33, 0x44, 0x0c, 0xc4, 0x16, 0x12, 0xe1, 0x62, 0x4d, 0x2f,
	0xca, 0x2f, 0x2d, 0x90, 0x60, 0x01, 0x0b, 0x42, 0x38, 0x4e, 0xac, 0x51, 0xcc, 0x89, 0x05, 0x99,
	0x49, 0x6c, 0x60, 0x7b, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x77, 0x2d, 0xb9, 0x29, 0x79,
	0x00, 0x00, 0x00,
}
