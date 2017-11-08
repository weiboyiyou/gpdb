// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cli_to_hub.proto

/*
Package idl is a generated protocol buffer package.

It is generated from these files:
	cli_to_hub.proto
	command.proto

It has these top-level messages:
	StatusUpgradeRequest
	StatusUpgradeReply
	UpgradeStepStatus
	CheckConfigRequest
	CheckConfigReply
	CheckUpgradeStatusRequest
	CheckUpgradeStatusReply
	FileSysUsage
	CheckDiskUsageRequest
	CheckDiskUsageReply
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

type UpgradeSteps int32

const (
	UpgradeSteps_UNKNOWN_STEP UpgradeSteps = 0
	UpgradeSteps_CHECK_CONFIG UpgradeSteps = 1
	UpgradeSteps_SEGINSTALL   UpgradeSteps = 2
)

var UpgradeSteps_name = map[int32]string{
	0: "UNKNOWN_STEP",
	1: "CHECK_CONFIG",
	2: "SEGINSTALL",
}
var UpgradeSteps_value = map[string]int32{
	"UNKNOWN_STEP": 0,
	"CHECK_CONFIG": 1,
	"SEGINSTALL":   2,
}

func (x UpgradeSteps) String() string {
	return proto.EnumName(UpgradeSteps_name, int32(x))
}
func (UpgradeSteps) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type StepStatus int32

const (
	StepStatus_UNKNOWN_STATUS StepStatus = 0
	StepStatus_PENDING        StepStatus = 1
	StepStatus_RUNNING        StepStatus = 2
	StepStatus_COMPLETE       StepStatus = 3
	StepStatus_FAILED         StepStatus = 4
)

var StepStatus_name = map[int32]string{
	0: "UNKNOWN_STATUS",
	1: "PENDING",
	2: "RUNNING",
	3: "COMPLETE",
	4: "FAILED",
}
var StepStatus_value = map[string]int32{
	"UNKNOWN_STATUS": 0,
	"PENDING":        1,
	"RUNNING":        2,
	"COMPLETE":       3,
	"FAILED":         4,
}

func (x StepStatus) String() string {
	return proto.EnumName(StepStatus_name, int32(x))
}
func (StepStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type StatusUpgradeRequest struct {
}

func (m *StatusUpgradeRequest) Reset()                    { *m = StatusUpgradeRequest{} }
func (m *StatusUpgradeRequest) String() string            { return proto.CompactTextString(m) }
func (*StatusUpgradeRequest) ProtoMessage()               {}
func (*StatusUpgradeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type StatusUpgradeReply struct {
	ListOfUpgradeStepStatuses []*UpgradeStepStatus `protobuf:"bytes,1,rep,name=listOfUpgradeStepStatuses" json:"listOfUpgradeStepStatuses,omitempty"`
}

func (m *StatusUpgradeReply) Reset()                    { *m = StatusUpgradeReply{} }
func (m *StatusUpgradeReply) String() string            { return proto.CompactTextString(m) }
func (*StatusUpgradeReply) ProtoMessage()               {}
func (*StatusUpgradeReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *StatusUpgradeReply) GetListOfUpgradeStepStatuses() []*UpgradeStepStatus {
	if m != nil {
		return m.ListOfUpgradeStepStatuses
	}
	return nil
}

type UpgradeStepStatus struct {
	Step   UpgradeSteps `protobuf:"varint,1,opt,name=step,enum=idl.UpgradeSteps" json:"step,omitempty"`
	Status StepStatus   `protobuf:"varint,2,opt,name=status,enum=idl.StepStatus" json:"status,omitempty"`
}

func (m *UpgradeStepStatus) Reset()                    { *m = UpgradeStepStatus{} }
func (m *UpgradeStepStatus) String() string            { return proto.CompactTextString(m) }
func (*UpgradeStepStatus) ProtoMessage()               {}
func (*UpgradeStepStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UpgradeStepStatus) GetStep() UpgradeSteps {
	if m != nil {
		return m.Step
	}
	return UpgradeSteps_UNKNOWN_STEP
}

func (m *UpgradeStepStatus) GetStatus() StepStatus {
	if m != nil {
		return m.Status
	}
	return StepStatus_UNKNOWN_STATUS
}

type CheckConfigRequest struct {
	DbPort int32 `protobuf:"varint,1,opt,name=dbPort" json:"dbPort,omitempty"`
}

func (m *CheckConfigRequest) Reset()                    { *m = CheckConfigRequest{} }
func (m *CheckConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*CheckConfigRequest) ProtoMessage()               {}
func (*CheckConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CheckConfigRequest) GetDbPort() int32 {
	if m != nil {
		return m.DbPort
	}
	return 0
}

type CheckConfigReply struct {
	ConfigStatus string `protobuf:"bytes,1,opt,name=configStatus" json:"configStatus,omitempty"`
}

func (m *CheckConfigReply) Reset()                    { *m = CheckConfigReply{} }
func (m *CheckConfigReply) String() string            { return proto.CompactTextString(m) }
func (*CheckConfigReply) ProtoMessage()               {}
func (*CheckConfigReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CheckConfigReply) GetConfigStatus() string {
	if m != nil {
		return m.ConfigStatus
	}
	return ""
}

func init() {
	proto.RegisterType((*StatusUpgradeRequest)(nil), "idl.StatusUpgradeRequest")
	proto.RegisterType((*StatusUpgradeReply)(nil), "idl.StatusUpgradeReply")
	proto.RegisterType((*UpgradeStepStatus)(nil), "idl.UpgradeStepStatus")
	proto.RegisterType((*CheckConfigRequest)(nil), "idl.CheckConfigRequest")
	proto.RegisterType((*CheckConfigReply)(nil), "idl.CheckConfigReply")
	proto.RegisterEnum("idl.UpgradeSteps", UpgradeSteps_name, UpgradeSteps_value)
	proto.RegisterEnum("idl.StepStatus", StepStatus_name, StepStatus_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CliToHub service

type CliToHubClient interface {
	StatusUpgrade(ctx context.Context, in *StatusUpgradeRequest, opts ...grpc.CallOption) (*StatusUpgradeReply, error)
	CheckConfig(ctx context.Context, in *CheckConfigRequest, opts ...grpc.CallOption) (*CheckConfigReply, error)
}

type cliToHubClient struct {
	cc *grpc.ClientConn
}

func NewCliToHubClient(cc *grpc.ClientConn) CliToHubClient {
	return &cliToHubClient{cc}
}

func (c *cliToHubClient) StatusUpgrade(ctx context.Context, in *StatusUpgradeRequest, opts ...grpc.CallOption) (*StatusUpgradeReply, error) {
	out := new(StatusUpgradeReply)
	err := grpc.Invoke(ctx, "/idl.CliToHub/StatusUpgrade", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cliToHubClient) CheckConfig(ctx context.Context, in *CheckConfigRequest, opts ...grpc.CallOption) (*CheckConfigReply, error) {
	out := new(CheckConfigReply)
	err := grpc.Invoke(ctx, "/idl.CliToHub/CheckConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CliToHub service

type CliToHubServer interface {
	StatusUpgrade(context.Context, *StatusUpgradeRequest) (*StatusUpgradeReply, error)
	CheckConfig(context.Context, *CheckConfigRequest) (*CheckConfigReply, error)
}

func RegisterCliToHubServer(s *grpc.Server, srv CliToHubServer) {
	s.RegisterService(&_CliToHub_serviceDesc, srv)
}

func _CliToHub_StatusUpgrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusUpgradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliToHubServer).StatusUpgrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.CliToHub/StatusUpgrade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliToHubServer).StatusUpgrade(ctx, req.(*StatusUpgradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CliToHub_CheckConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliToHubServer).CheckConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.CliToHub/CheckConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliToHubServer).CheckConfig(ctx, req.(*CheckConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CliToHub_serviceDesc = grpc.ServiceDesc{
	ServiceName: "idl.CliToHub",
	HandlerType: (*CliToHubServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StatusUpgrade",
			Handler:    _CliToHub_StatusUpgrade_Handler,
		},
		{
			MethodName: "CheckConfig",
			Handler:    _CliToHub_CheckConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cli_to_hub.proto",
}

func init() { proto.RegisterFile("cli_to_hub.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4b, 0x8f, 0x93, 0x50,
	0x14, 0x2e, 0xed, 0xc8, 0x8c, 0x07, 0xac, 0x77, 0x4e, 0xb4, 0x76, 0x66, 0x35, 0x21, 0x31, 0x4e,
	0x26, 0xa6, 0x8b, 0x9a, 0xb8, 0x35, 0x95, 0xde, 0xb6, 0xa4, 0x78, 0x21, 0x3c, 0xe2, 0x92, 0x94,
	0x47, 0x5b, 0x94, 0x08, 0x96, 0xcb, 0xa2, 0x7f, 0xc3, 0x5f, 0x6c, 0xb8, 0x25, 0x91, 0x16, 0x67,
	0x79, 0xbe, 0xd7, 0xe5, 0x7c, 0x1c, 0x20, 0x51, 0x96, 0x06, 0x3c, 0x0f, 0xf6, 0x55, 0x38, 0x29,
	0x0e, 0x39, 0xcf, 0x71, 0x90, 0xc6, 0x99, 0x36, 0x82, 0x37, 0x2e, 0xdf, 0xf0, 0xaa, 0xf4, 0x8b,
	0xdd, 0x61, 0x13, 0x27, 0x4e, 0xf2, 0xbb, 0x4a, 0x4a, 0xae, 0xfd, 0x00, 0xbc, 0xc0, 0x8b, 0xec,
	0x88, 0x1e, 0xdc, 0x65, 0x69, 0xc9, 0xad, 0x6d, 0x83, 0xba, 0x3c, 0x29, 0x4e, 0xb2, 0xa4, 0x1c,
	0x4b, 0x0f, 0x83, 0x47, 0x65, 0x3a, 0x9a, 0xa4, 0x71, 0x36, 0xe9, 0xf0, 0xce, 0xf3, 0x46, 0x2d,
	0x82, 0xdb, 0x0e, 0x8c, 0xef, 0xe1, 0xaa, 0xe4, 0x49, 0x31, 0x96, 0x1e, 0xa4, 0xc7, 0xe1, 0xf4,
	0xf6, 0x32, 0xb5, 0x74, 0x04, 0x8d, 0x1f, 0x40, 0x2e, 0x85, 0x61, 0xdc, 0x17, 0xc2, 0xd7, 0x42,
	0xd8, 0x7a, 0xb7, 0xa1, 0xb5, 0x8f, 0x80, 0xfa, 0x3e, 0x89, 0x7e, 0xea, 0xf9, 0xaf, 0x6d, 0xba,
	0x6b, 0xd6, 0xc4, 0x11, 0xc8, 0x71, 0x68, 0xe7, 0x07, 0x2e, 0xde, 0x79, 0xe1, 0x34, 0x93, 0xf6,
	0x19, 0xc8, 0x99, 0xba, 0x5e, 0x5e, 0x03, 0x35, 0x12, 0xe3, 0x29, 0x59, 0x38, 0x5e, 0x3a, 0x67,
	0xd8, 0xd3, 0x57, 0x50, 0xdb, 0x1f, 0x89, 0x04, 0x54, 0x9f, 0xad, 0x99, 0xf5, 0x9d, 0x05, 0xae,
	0x47, 0x6d, 0xd2, 0xab, 0x11, 0x7d, 0x45, 0xf5, 0x75, 0xa0, 0x5b, 0x6c, 0x61, 0x2c, 0x89, 0x84,
	0x43, 0x00, 0x97, 0x2e, 0x0d, 0xe6, 0x7a, 0x33, 0xd3, 0x24, 0xfd, 0x27, 0x0f, 0xa0, 0xd5, 0x03,
	0xc2, 0xf0, 0x5f, 0xc2, 0xcc, 0xf3, 0x5d, 0xd2, 0x43, 0x05, 0xae, 0x6d, 0xca, 0xe6, 0x06, 0xab,
	0xed, 0x0a, 0x5c, 0x3b, 0x3e, 0x63, 0xf5, 0xd0, 0x47, 0x15, 0x6e, 0x74, 0xeb, 0x9b, 0x6d, 0x52,
	0x8f, 0x92, 0x01, 0x02, 0xc8, 0x8b, 0x99, 0x61, 0xd2, 0x39, 0xb9, 0x9a, 0xfe, 0x91, 0xe0, 0x46,
	0xcf, 0x52, 0x2f, 0x5f, 0x55, 0x21, 0x52, 0x78, 0x75, 0xf6, 0x77, 0xf1, 0xae, 0xa9, 0xad, 0x7b,
	0x09, 0xf7, 0xef, 0xfe, 0x47, 0x15, 0xd9, 0x51, 0xeb, 0xe1, 0x17, 0x50, 0x5a, 0x2d, 0xe1, 0x49,
	0xd9, 0x6d, 0xf9, 0xfe, 0x6d, 0x97, 0x10, 0x01, 0xa1, 0x2c, 0x2e, 0xf1, 0xd3, 0xdf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x5e, 0x3b, 0xe7, 0x5a, 0x9d, 0x02, 0x00, 0x00,
}
