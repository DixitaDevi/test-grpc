// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package stetpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Person struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreatePersonRequest struct {
	Person               *Person  `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePersonRequest) Reset()         { *m = CreatePersonRequest{} }
func (m *CreatePersonRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePersonRequest) ProtoMessage()    {}
func (*CreatePersonRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *CreatePersonRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePersonRequest.Unmarshal(m, b)
}
func (m *CreatePersonRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePersonRequest.Marshal(b, m, deterministic)
}
func (m *CreatePersonRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePersonRequest.Merge(m, src)
}
func (m *CreatePersonRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePersonRequest.Size(m)
}
func (m *CreatePersonRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePersonRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePersonRequest proto.InternalMessageInfo

func (m *CreatePersonRequest) GetPerson() *Person {
	if m != nil {
		return m.Person
	}
	return nil
}

type CreatePersonResponse struct {
	Person               *Person  `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePersonResponse) Reset()         { *m = CreatePersonResponse{} }
func (m *CreatePersonResponse) String() string { return proto.CompactTextString(m) }
func (*CreatePersonResponse) ProtoMessage()    {}
func (*CreatePersonResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *CreatePersonResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePersonResponse.Unmarshal(m, b)
}
func (m *CreatePersonResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePersonResponse.Marshal(b, m, deterministic)
}
func (m *CreatePersonResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePersonResponse.Merge(m, src)
}
func (m *CreatePersonResponse) XXX_Size() int {
	return xxx_messageInfo_CreatePersonResponse.Size(m)
}
func (m *CreatePersonResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePersonResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePersonResponse proto.InternalMessageInfo

func (m *CreatePersonResponse) GetPerson() *Person {
	if m != nil {
		return m.Person
	}
	return nil
}

type ReadPersonRequest struct {
	PersonId             string   `protobuf:"bytes,1,opt,name=person_id,json=personId,proto3" json:"person_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadPersonRequest) Reset()         { *m = ReadPersonRequest{} }
func (m *ReadPersonRequest) String() string { return proto.CompactTextString(m) }
func (*ReadPersonRequest) ProtoMessage()    {}
func (*ReadPersonRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *ReadPersonRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadPersonRequest.Unmarshal(m, b)
}
func (m *ReadPersonRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadPersonRequest.Marshal(b, m, deterministic)
}
func (m *ReadPersonRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadPersonRequest.Merge(m, src)
}
func (m *ReadPersonRequest) XXX_Size() int {
	return xxx_messageInfo_ReadPersonRequest.Size(m)
}
func (m *ReadPersonRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadPersonRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadPersonRequest proto.InternalMessageInfo

func (m *ReadPersonRequest) GetPersonId() string {
	if m != nil {
		return m.PersonId
	}
	return ""
}

type ReadPersonResponse struct {
	Person               *Person  `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadPersonResponse) Reset()         { *m = ReadPersonResponse{} }
func (m *ReadPersonResponse) String() string { return proto.CompactTextString(m) }
func (*ReadPersonResponse) ProtoMessage()    {}
func (*ReadPersonResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *ReadPersonResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadPersonResponse.Unmarshal(m, b)
}
func (m *ReadPersonResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadPersonResponse.Marshal(b, m, deterministic)
}
func (m *ReadPersonResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadPersonResponse.Merge(m, src)
}
func (m *ReadPersonResponse) XXX_Size() int {
	return xxx_messageInfo_ReadPersonResponse.Size(m)
}
func (m *ReadPersonResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadPersonResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadPersonResponse proto.InternalMessageInfo

func (m *ReadPersonResponse) GetPerson() *Person {
	if m != nil {
		return m.Person
	}
	return nil
}

type UpdatePersonRequest struct {
	Person               *Person  `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePersonRequest) Reset()         { *m = UpdatePersonRequest{} }
func (m *UpdatePersonRequest) String() string { return proto.CompactTextString(m) }
func (*UpdatePersonRequest) ProtoMessage()    {}
func (*UpdatePersonRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{5}
}

func (m *UpdatePersonRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePersonRequest.Unmarshal(m, b)
}
func (m *UpdatePersonRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePersonRequest.Marshal(b, m, deterministic)
}
func (m *UpdatePersonRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePersonRequest.Merge(m, src)
}
func (m *UpdatePersonRequest) XXX_Size() int {
	return xxx_messageInfo_UpdatePersonRequest.Size(m)
}
func (m *UpdatePersonRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePersonRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePersonRequest proto.InternalMessageInfo

func (m *UpdatePersonRequest) GetPerson() *Person {
	if m != nil {
		return m.Person
	}
	return nil
}

type UpdatePersonResponse struct {
	Person               *Person  `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePersonResponse) Reset()         { *m = UpdatePersonResponse{} }
func (m *UpdatePersonResponse) String() string { return proto.CompactTextString(m) }
func (*UpdatePersonResponse) ProtoMessage()    {}
func (*UpdatePersonResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{6}
}

func (m *UpdatePersonResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePersonResponse.Unmarshal(m, b)
}
func (m *UpdatePersonResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePersonResponse.Marshal(b, m, deterministic)
}
func (m *UpdatePersonResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePersonResponse.Merge(m, src)
}
func (m *UpdatePersonResponse) XXX_Size() int {
	return xxx_messageInfo_UpdatePersonResponse.Size(m)
}
func (m *UpdatePersonResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePersonResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePersonResponse proto.InternalMessageInfo

func (m *UpdatePersonResponse) GetPerson() *Person {
	if m != nil {
		return m.Person
	}
	return nil
}

type DeletePersonRequest struct {
	PersonId             string   `protobuf:"bytes,1,opt,name=person_id,json=personId,proto3" json:"person_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePersonRequest) Reset()         { *m = DeletePersonRequest{} }
func (m *DeletePersonRequest) String() string { return proto.CompactTextString(m) }
func (*DeletePersonRequest) ProtoMessage()    {}
func (*DeletePersonRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{7}
}

func (m *DeletePersonRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePersonRequest.Unmarshal(m, b)
}
func (m *DeletePersonRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePersonRequest.Marshal(b, m, deterministic)
}
func (m *DeletePersonRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePersonRequest.Merge(m, src)
}
func (m *DeletePersonRequest) XXX_Size() int {
	return xxx_messageInfo_DeletePersonRequest.Size(m)
}
func (m *DeletePersonRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePersonRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePersonRequest proto.InternalMessageInfo

func (m *DeletePersonRequest) GetPersonId() string {
	if m != nil {
		return m.PersonId
	}
	return ""
}

type DeletePersonResponse struct {
	PersonId             string   `protobuf:"bytes,1,opt,name=person_id,json=personId,proto3" json:"person_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePersonResponse) Reset()         { *m = DeletePersonResponse{} }
func (m *DeletePersonResponse) String() string { return proto.CompactTextString(m) }
func (*DeletePersonResponse) ProtoMessage()    {}
func (*DeletePersonResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{8}
}

func (m *DeletePersonResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePersonResponse.Unmarshal(m, b)
}
func (m *DeletePersonResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePersonResponse.Marshal(b, m, deterministic)
}
func (m *DeletePersonResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePersonResponse.Merge(m, src)
}
func (m *DeletePersonResponse) XXX_Size() int {
	return xxx_messageInfo_DeletePersonResponse.Size(m)
}
func (m *DeletePersonResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePersonResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePersonResponse proto.InternalMessageInfo

func (m *DeletePersonResponse) GetPersonId() string {
	if m != nil {
		return m.PersonId
	}
	return ""
}

func init() {
	proto.RegisterType((*Person)(nil), "proto.Person")
	proto.RegisterType((*CreatePersonRequest)(nil), "proto.CreatePersonRequest")
	proto.RegisterType((*CreatePersonResponse)(nil), "proto.CreatePersonResponse")
	proto.RegisterType((*ReadPersonRequest)(nil), "proto.ReadPersonRequest")
	proto.RegisterType((*ReadPersonResponse)(nil), "proto.ReadPersonResponse")
	proto.RegisterType((*UpdatePersonRequest)(nil), "proto.UpdatePersonRequest")
	proto.RegisterType((*UpdatePersonResponse)(nil), "proto.UpdatePersonResponse")
	proto.RegisterType((*DeletePersonRequest)(nil), "proto.DeletePersonRequest")
	proto.RegisterType((*DeletePersonResponse)(nil), "proto.DeletePersonResponse")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x69, 0xd0, 0xd0, 0xbe, 0xb6, 0x82, 0x9b, 0x1e, 0x62, 0x7a, 0x91, 0x80, 0xe0, 0x41,
	0x8a, 0xa4, 0x47, 0xeb, 0xc1, 0x3f, 0x10, 0xbc, 0x49, 0x8a, 0x17, 0x2f, 0x92, 0xba, 0x73, 0x08,
	0x68, 0xb2, 0x66, 0x57, 0xbf, 0x94, 0x5f, 0x52, 0xd8, 0x5d, 0xd6, 0x44, 0x97, 0x62, 0xf1, 0x94,
	0x30, 0x33, 0xbf, 0x79, 0x33, 0x6f, 0x16, 0x53, 0x49, 0xed, 0x47, 0xf5, 0x4c, 0x0b, 0xd1, 0x36,
	0xaa, 0x61, 0xfb, 0xfa, 0x93, 0x9e, 0x21, 0xbc, 0xa7, 0x56, 0x36, 0x35, 0x3b, 0x40, 0x50, 0xf1,
	0x78, 0x70, 0x3c, 0x38, 0x1d, 0x15, 0x41, 0xc5, 0x19, 0xc3, 0x5e, 0x5d, 0xbe, 0x52, 0x1c, 0xe8,
	0x88, 0xfe, 0x4f, 0x57, 0x88, 0x6e, 0x5a, 0x2a, 0x15, 0x19, 0xa6, 0xa0, 0xb7, 0x77, 0x92, 0x8a,
	0x9d, 0x20, 0x14, 0x3a, 0xa0, 0xf1, 0x71, 0x36, 0x35, 0x1a, 0x0b, 0x5b, 0x65, 0x93, 0xe9, 0x25,
	0x66, 0x7d, 0x5a, 0x8a, 0xa6, 0x96, 0xf4, 0x57, 0xfc, 0x1c, 0x87, 0x05, 0x95, 0xbc, 0x2f, 0x3d,
	0xc7, 0xc8, 0xa4, 0x9f, 0xdc, 0xf0, 0x43, 0x13, 0xb8, 0xe3, 0xe9, 0x05, 0x58, 0x97, 0xd8, 0x4d,
	0x6e, 0x85, 0xe8, 0x41, 0xf0, 0x7f, 0xec, 0xda, 0xa7, 0x77, 0x13, 0xcf, 0x10, 0xdd, 0xd2, 0x0b,
	0xfd, 0x14, 0xdf, 0xba, 0xed, 0x12, 0xb3, 0x3e, 0x63, 0x25, 0xb7, 0x41, 0xd9, 0x67, 0x80, 0xf1,
	0x5a, 0x91, 0x5a, 0x9b, 0xc7, 0xc1, 0x72, 0x4c, 0xba, 0x37, 0x62, 0x89, 0x9d, 0xcf, 0x73, 0xf6,
	0x64, 0xee, 0xcd, 0x59, 0xd5, 0x2b, 0xe0, 0xdb, 0x7b, 0x16, 0xdb, 0xd2, 0x5f, 0x07, 0x4c, 0x8e,
	0x3c, 0x19, 0xdb, 0x22, 0xc7, 0xa4, 0xeb, 0xa1, 0x9b, 0xc5, 0x73, 0x16, 0x37, 0x8b, 0xd7, 0xf4,
	0x1c, 0x93, 0xae, 0x33, 0xae, 0x91, 0xc7, 0x62, 0xd7, 0xc8, 0x67, 0xe5, 0xf5, 0xf0, 0x31, 0x94,
	0x8a, 0x94, 0xd8, 0x6c, 0x42, 0x5d, 0xb5, 0xfc, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x39, 0xb3,
	0x59, 0x56, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StetServiceClient is the client API for StetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StetServiceClient interface {
	CreatePerson(ctx context.Context, in *CreatePersonRequest, opts ...grpc.CallOption) (*CreatePersonResponse, error)
	ReadPerson(ctx context.Context, in *ReadPersonRequest, opts ...grpc.CallOption) (*ReadPersonResponse, error)
	UpdatePerson(ctx context.Context, in *UpdatePersonRequest, opts ...grpc.CallOption) (*UpdatePersonResponse, error)
	DeletePerson(ctx context.Context, in *DeletePersonRequest, opts ...grpc.CallOption) (*DeletePersonResponse, error)
}

type stetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStetServiceClient(cc grpc.ClientConnInterface) StetServiceClient {
	return &stetServiceClient{cc}
}

func (c *stetServiceClient) CreatePerson(ctx context.Context, in *CreatePersonRequest, opts ...grpc.CallOption) (*CreatePersonResponse, error) {
	out := new(CreatePersonResponse)
	err := c.cc.Invoke(ctx, "/proto.StetService/CreatePerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stetServiceClient) ReadPerson(ctx context.Context, in *ReadPersonRequest, opts ...grpc.CallOption) (*ReadPersonResponse, error) {
	out := new(ReadPersonResponse)
	err := c.cc.Invoke(ctx, "/proto.StetService/ReadPerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stetServiceClient) UpdatePerson(ctx context.Context, in *UpdatePersonRequest, opts ...grpc.CallOption) (*UpdatePersonResponse, error) {
	out := new(UpdatePersonResponse)
	err := c.cc.Invoke(ctx, "/proto.StetService/UpdatePerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stetServiceClient) DeletePerson(ctx context.Context, in *DeletePersonRequest, opts ...grpc.CallOption) (*DeletePersonResponse, error) {
	out := new(DeletePersonResponse)
	err := c.cc.Invoke(ctx, "/proto.StetService/DeletePerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StetServiceServer is the server API for StetService service.
type StetServiceServer interface {
	CreatePerson(context.Context, *CreatePersonRequest) (*CreatePersonResponse, error)
	ReadPerson(context.Context, *ReadPersonRequest) (*ReadPersonResponse, error)
	UpdatePerson(context.Context, *UpdatePersonRequest) (*UpdatePersonResponse, error)
	DeletePerson(context.Context, *DeletePersonRequest) (*DeletePersonResponse, error)
}

// UnimplementedStetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStetServiceServer struct {
}

func (*UnimplementedStetServiceServer) CreatePerson(ctx context.Context, req *CreatePersonRequest) (*CreatePersonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePerson not implemented")
}
func (*UnimplementedStetServiceServer) ReadPerson(ctx context.Context, req *ReadPersonRequest) (*ReadPersonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadPerson not implemented")
}
func (*UnimplementedStetServiceServer) UpdatePerson(ctx context.Context, req *UpdatePersonRequest) (*UpdatePersonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePerson not implemented")
}
func (*UnimplementedStetServiceServer) DeletePerson(ctx context.Context, req *DeletePersonRequest) (*DeletePersonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePerson not implemented")
}

func RegisterStetServiceServer(s *grpc.Server, srv StetServiceServer) {
	s.RegisterService(&_StetService_serviceDesc, srv)
}

func _StetService_CreatePerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StetServiceServer).CreatePerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.StetService/CreatePerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StetServiceServer).CreatePerson(ctx, req.(*CreatePersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StetService_ReadPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadPersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StetServiceServer).ReadPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.StetService/ReadPerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StetServiceServer).ReadPerson(ctx, req.(*ReadPersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StetService_UpdatePerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StetServiceServer).UpdatePerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.StetService/UpdatePerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StetServiceServer).UpdatePerson(ctx, req.(*UpdatePersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StetService_DeletePerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StetServiceServer).DeletePerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.StetService/DeletePerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StetServiceServer).DeletePerson(ctx, req.(*DeletePersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.StetService",
	HandlerType: (*StetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePerson",
			Handler:    _StetService_CreatePerson_Handler,
		},
		{
			MethodName: "ReadPerson",
			Handler:    _StetService_ReadPerson_Handler,
		},
		{
			MethodName: "UpdatePerson",
			Handler:    _StetService_UpdatePerson_Handler,
		},
		{
			MethodName: "DeletePerson",
			Handler:    _StetService_DeletePerson_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
