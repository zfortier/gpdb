// Code generated by protoc-gen-go. DO NOT EDIT.
// source: command.proto

/*
Package idl is a generated protocol buffer package.

It is generated from these files:
	command.proto

It has these top-level messages:
	CheckUpgradeStatusRequest
	CheckUpgradeStatusReply
*/
package idl

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CheckUpgradeStatusRequest struct {
}

func (m *CheckUpgradeStatusRequest) Reset()                    { *m = CheckUpgradeStatusRequest{} }
func (m *CheckUpgradeStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*CheckUpgradeStatusRequest) ProtoMessage()               {}
func (*CheckUpgradeStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CheckUpgradeStatusReply struct {
	ProcessList string `protobuf:"bytes,1,opt,name=process_list,json=processList" json:"process_list,omitempty"`
}

func (m *CheckUpgradeStatusReply) Reset()                    { *m = CheckUpgradeStatusReply{} }
func (m *CheckUpgradeStatusReply) String() string            { return proto.CompactTextString(m) }
func (*CheckUpgradeStatusReply) ProtoMessage()               {}
func (*CheckUpgradeStatusReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CheckUpgradeStatusReply) GetProcessList() string {
	if m != nil {
		return m.ProcessList
	}
	return ""
}

func init() {
	proto.RegisterType((*CheckUpgradeStatusRequest)(nil), "idl.CheckUpgradeStatusRequest")
	proto.RegisterType((*CheckUpgradeStatusReply)(nil), "idl.CheckUpgradeStatusReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CommandListener service

type CommandListenerClient interface {
	CheckUpgradeStatus(ctx context.Context, in *CheckUpgradeStatusRequest, opts ...grpc.CallOption) (*CheckUpgradeStatusReply, error)
}

type commandListenerClient struct {
	cc *grpc.ClientConn
}

func NewCommandListenerClient(cc *grpc.ClientConn) CommandListenerClient {
	return &commandListenerClient{cc}
}

func (c *commandListenerClient) CheckUpgradeStatus(ctx context.Context, in *CheckUpgradeStatusRequest, opts ...grpc.CallOption) (*CheckUpgradeStatusReply, error) {
	out := new(CheckUpgradeStatusReply)
	err := grpc.Invoke(ctx, "/idl.CommandListener/CheckUpgradeStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CommandListener service

type CommandListenerServer interface {
	CheckUpgradeStatus(context.Context, *CheckUpgradeStatusRequest) (*CheckUpgradeStatusReply, error)
}

func RegisterCommandListenerServer(s *grpc.Server, srv CommandListenerServer) {
	s.RegisterService(&_CommandListener_serviceDesc, srv)
}

func _CommandListener_CheckUpgradeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUpgradeStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandListenerServer).CheckUpgradeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.CommandListener/CheckUpgradeStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandListenerServer).CheckUpgradeStatus(ctx, req.(*CheckUpgradeStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CommandListener_serviceDesc = grpc.ServiceDesc{
	ServiceName: "idl.CommandListener",
	HandlerType: (*CommandListenerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckUpgradeStatus",
			Handler:    _CommandListener_CheckUpgradeStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "command.proto",
}

func init() { proto.RegisterFile("command.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xce, 0xcf, 0xcd,
	0x4d, 0xcc, 0x4b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x4c, 0xc9, 0x51, 0x92,
	0xe6, 0x92, 0x74, 0xce, 0x48, 0x4d, 0xce, 0x0e, 0x2d, 0x48, 0x2f, 0x4a, 0x4c, 0x49, 0x0d, 0x2e,
	0x49, 0x2c, 0x29, 0x2d, 0x0e, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x51, 0xb2, 0xe1, 0x12, 0xc7,
	0x26, 0x59, 0x90, 0x53, 0x29, 0xa4, 0xc8, 0xc5, 0x53, 0x50, 0x94, 0x9f, 0x9c, 0x5a, 0x5c, 0x1c,
	0x9f, 0x93, 0x59, 0x5c, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0xc4, 0x0d, 0x15, 0xf3, 0xc9,
	0x2c, 0x2e, 0x31, 0x4a, 0xe7, 0xe2, 0x77, 0x86, 0x58, 0x08, 0xe2, 0xa6, 0xe6, 0xa5, 0x16, 0x09,
	0x85, 0x70, 0x09, 0x61, 0x1a, 0x28, 0x24, 0xa7, 0x97, 0x99, 0x92, 0xa3, 0x87, 0xd3, 0x19, 0x52,
	0x32, 0x38, 0xe5, 0x0b, 0x72, 0x2a, 0x95, 0x18, 0x92, 0xd8, 0xc0, 0xfe, 0x31, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x49, 0x5e, 0xf5, 0x92, 0xe0, 0x00, 0x00, 0x00,
}
