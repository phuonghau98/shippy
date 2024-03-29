// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/vessel/vessel.proto

package vessel

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type Specification struct {
	Capacity             int32    `protobuf:"varint,1,opt,name=capacity,proto3" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,2,opt,name=max_weight,json=maxWeight,proto3" json:"max_weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Specification) Reset()         { *m = Specification{} }
func (m *Specification) String() string { return proto.CompactTextString(m) }
func (*Specification) ProtoMessage()    {}
func (*Specification) Descriptor() ([]byte, []int) {
	return fileDescriptor_04ef66862bb50716, []int{0}
}

func (m *Specification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Specification.Unmarshal(m, b)
}
func (m *Specification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Specification.Marshal(b, m, deterministic)
}
func (m *Specification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Specification.Merge(m, src)
}
func (m *Specification) XXX_Size() int {
	return xxx_messageInfo_Specification.Size(m)
}
func (m *Specification) XXX_DiscardUnknown() {
	xxx_messageInfo_Specification.DiscardUnknown(m)
}

var xxx_messageInfo_Specification proto.InternalMessageInfo

func (m *Specification) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Specification) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

type Response struct {
	Vessel               *Vessel   `protobuf:"bytes,1,opt,name=vessel,proto3" json:"vessel,omitempty"`
	Vessels              []*Vessel `protobuf:"bytes,2,rep,name=vessels,proto3" json:"vessels,omitempty"`
	Created              bool      `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_04ef66862bb50716, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetVessel() *Vessel {
	if m != nil {
		return m.Vessel
	}
	return nil
}

func (m *Response) GetVessels() []*Vessel {
	if m != nil {
		return m.Vessels
	}
	return nil
}

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

type Vessel struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Capacity             int32    `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,3,opt,name=max_weight,json=maxWeight,proto3" json:"max_weight,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Available            bool     `protobuf:"varint,5,opt,name=available,proto3" json:"available,omitempty"`
	OwnerId              string   `protobuf:"bytes,6,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vessel) Reset()         { *m = Vessel{} }
func (m *Vessel) String() string { return proto.CompactTextString(m) }
func (*Vessel) ProtoMessage()    {}
func (*Vessel) Descriptor() ([]byte, []int) {
	return fileDescriptor_04ef66862bb50716, []int{2}
}

func (m *Vessel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vessel.Unmarshal(m, b)
}
func (m *Vessel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vessel.Marshal(b, m, deterministic)
}
func (m *Vessel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vessel.Merge(m, src)
}
func (m *Vessel) XXX_Size() int {
	return xxx_messageInfo_Vessel.Size(m)
}
func (m *Vessel) XXX_DiscardUnknown() {
	xxx_messageInfo_Vessel.DiscardUnknown(m)
}

var xxx_messageInfo_Vessel proto.InternalMessageInfo

func (m *Vessel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Vessel) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Vessel) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

func (m *Vessel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Vessel) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *Vessel) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

func init() {
	proto.RegisterType((*Specification)(nil), "vessel.Specification")
	proto.RegisterType((*Response)(nil), "vessel.Response")
	proto.RegisterType((*Vessel)(nil), "vessel.Vessel")
}

func init() { proto.RegisterFile("proto/vessel/vessel.proto", fileDescriptor_04ef66862bb50716) }

var fileDescriptor_04ef66862bb50716 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0x4b, 0x4b, 0xf3, 0x50,
	0x10, 0xfd, 0x92, 0xb6, 0x79, 0xcc, 0x47, 0x8a, 0x0c, 0x08, 0xb7, 0x45, 0x21, 0x64, 0x21, 0x59,
	0x48, 0x85, 0xba, 0x73, 0x27, 0x82, 0xa0, 0xcb, 0x14, 0x74, 0x59, 0x6e, 0xef, 0x1d, 0xf5, 0x42,
	0x5e, 0x24, 0x21, 0x6d, 0xff, 0x8d, 0x3f, 0x55, 0x98, 0x24, 0x95, 0x96, 0xe2, 0x6a, 0xe6, 0x9c,
	0x79, 0x9e, 0x19, 0x98, 0x95, 0x55, 0xd1, 0x14, 0x77, 0x2d, 0xd5, 0x35, 0xa5, 0xbd, 0x59, 0x30,
	0x87, 0x4e, 0x87, 0xa2, 0x57, 0x08, 0x56, 0x25, 0x29, 0xf3, 0x61, 0x94, 0x6c, 0x4c, 0x91, 0xe3,
	0x1c, 0x3c, 0x25, 0x4b, 0xa9, 0x4c, 0xb3, 0x17, 0x56, 0x68, 0xc5, 0x93, 0xe4, 0x80, 0xf1, 0x1a,
	0x20, 0x93, 0xbb, 0xf5, 0x96, 0xcc, 0xe7, 0x57, 0x23, 0x6c, 0x8e, 0xfa, 0x99, 0xdc, 0xbd, 0x33,
	0x11, 0xb5, 0xe0, 0x25, 0x54, 0x97, 0x45, 0x5e, 0x13, 0xde, 0x40, 0x3f, 0x81, 0x9b, 0xfc, 0x5f,
	0x4e, 0x17, 0xfd, 0xf8, 0x37, 0x36, 0x49, 0x1f, 0xc5, 0x18, 0xdc, 0xce, 0xab, 0x85, 0x1d, 0x8e,
	0xce, 0x24, 0x0e, 0x61, 0x14, 0xe0, 0xaa, 0x8a, 0x64, 0x43, 0x5a, 0x8c, 0x42, 0x2b, 0xf6, 0x92,
	0x01, 0x46, 0xdf, 0x16, 0x38, 0x5d, 0x36, 0x4e, 0xc1, 0x36, 0x9a, 0x47, 0xfa, 0x89, 0x6d, 0xf4,
	0x91, 0x1a, 0xfb, 0x4f, 0x35, 0xa3, 0x13, 0x35, 0x88, 0x30, 0xce, 0x65, 0x46, 0x62, 0xcc, 0xcd,
	0xd8, 0xc7, 0x2b, 0xf0, 0x65, 0x2b, 0x4d, 0x2a, 0x37, 0x29, 0x89, 0x09, 0x6f, 0xf1, 0x4b, 0xe0,
	0x0c, 0xbc, 0x62, 0x9b, 0x53, 0xb5, 0x36, 0x5a, 0x38, 0x5c, 0xe5, 0x32, 0x7e, 0xd1, 0xcb, 0x3d,
	0x04, 0xdd, 0x86, 0x2b, 0xaa, 0x5a, 0xa3, 0x08, 0x1f, 0x20, 0x78, 0x36, 0xb9, 0x7e, 0x3c, 0x14,
	0x5f, 0x0e, 0xba, 0x8f, 0xde, 0x31, 0xbf, 0x18, 0xe8, 0xe1, 0xb2, 0xd1, 0x3f, 0xbc, 0x05, 0xe7,
	0x89, 0xa5, 0xe3, 0xc9, 0xb1, 0xce, 0x65, 0x6f, 0x1c, 0x7e, 0xf8, 0xfd, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xfe, 0xb4, 0x90, 0x46, 0x0d, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for VesselService service

type VesselServiceClient interface {
	FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error)
	Create(ctx context.Context, in *Vessel, opts ...client.CallOption) (*Response, error)
}

type vesselServiceClient struct {
	c           client.Client
	serviceName string
}

func NewVesselServiceClient(serviceName string, c client.Client) VesselServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "vessel"
	}
	return &vesselServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *vesselServiceClient) FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "VesselService.FindAvailable", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vesselServiceClient) Create(ctx context.Context, in *Vessel, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "VesselService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VesselService service

type VesselServiceHandler interface {
	FindAvailable(context.Context, *Specification, *Response) error
	Create(context.Context, *Vessel, *Response) error
}

func RegisterVesselServiceHandler(s server.Server, hdlr VesselServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&VesselService{hdlr}, opts...))
}

type VesselService struct {
	VesselServiceHandler
}

func (h *VesselService) FindAvailable(ctx context.Context, in *Specification, out *Response) error {
	return h.VesselServiceHandler.FindAvailable(ctx, in, out)
}

func (h *VesselService) Create(ctx context.Context, in *Vessel, out *Response) error {
	return h.VesselServiceHandler.Create(ctx, in, out)
}
