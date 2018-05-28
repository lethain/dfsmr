// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dfsmr.proto

/*
Package dfsmr is a generated protocol buffer package.

It is generated from these files:
	dfsmr.proto

It has these top-level messages:
	Transition
	Node
	DefineRequest
	DefineReply
	MachinesRequest
	MachinesReply
	ReadyRequest
	TaskMessage
	InstancesRequest
	InstancesReply
	ChangesRequest
	ChangesReply
*/
package dfsmr

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

type Transition struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Node string `protobuf:"bytes,2,opt,name=node" json:"node,omitempty"`
}

func (m *Transition) Reset()                    { *m = Transition{} }
func (m *Transition) String() string            { return proto.CompactTextString(m) }
func (*Transition) ProtoMessage()               {}
func (*Transition) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Transition) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Transition) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

type Node struct {
	Id          string        `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Transitions []*Transition `protobuf:"bytes,2,rep,name=transitions" json:"transitions,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Node) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Node) GetTransitions() []*Transition {
	if m != nil {
		return m.Transitions
	}
	return nil
}

type DefineRequest struct {
	Id    string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Nodes []*Node `protobuf:"bytes,2,rep,name=nodes" json:"nodes,omitempty"`
}

func (m *DefineRequest) Reset()                    { *m = DefineRequest{} }
func (m *DefineRequest) String() string            { return proto.CompactTextString(m) }
func (*DefineRequest) ProtoMessage()               {}
func (*DefineRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DefineRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DefineRequest) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type DefineReply struct {
	Success        bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Id             string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	SuccessMessage string `protobuf:"bytes,3,opt,name=successMessage" json:"successMessage,omitempty"`
	ErrorMessage   string `protobuf:"bytes,4,opt,name=errorMessage" json:"errorMessage,omitempty"`
}

func (m *DefineReply) Reset()                    { *m = DefineReply{} }
func (m *DefineReply) String() string            { return proto.CompactTextString(m) }
func (*DefineReply) ProtoMessage()               {}
func (*DefineReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DefineReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *DefineReply) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DefineReply) GetSuccessMessage() string {
	if m != nil {
		return m.SuccessMessage
	}
	return ""
}

func (m *DefineReply) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

type MachinesRequest struct {
}

func (m *MachinesRequest) Reset()                    { *m = MachinesRequest{} }
func (m *MachinesRequest) String() string            { return proto.CompactTextString(m) }
func (*MachinesRequest) ProtoMessage()               {}
func (*MachinesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type MachinesReply struct {
	Machines []*DefineRequest `protobuf:"bytes,1,rep,name=machines" json:"machines,omitempty"`
}

func (m *MachinesReply) Reset()                    { *m = MachinesReply{} }
func (m *MachinesReply) String() string            { return proto.CompactTextString(m) }
func (*MachinesReply) ProtoMessage()               {}
func (*MachinesReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *MachinesReply) GetMachines() []*DefineRequest {
	if m != nil {
		return m.Machines
	}
	return nil
}

type ReadyRequest struct {
	Client  string `protobuf:"bytes,1,opt,name=client" json:"client,omitempty"`
	Machine string `protobuf:"bytes,2,opt,name=machine" json:"machine,omitempty"`
	Node    string `protobuf:"bytes,3,opt,name=node" json:"node,omitempty"`
}

func (m *ReadyRequest) Reset()                    { *m = ReadyRequest{} }
func (m *ReadyRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadyRequest) ProtoMessage()               {}
func (*ReadyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ReadyRequest) GetClient() string {
	if m != nil {
		return m.Client
	}
	return ""
}

func (m *ReadyRequest) GetMachine() string {
	if m != nil {
		return m.Machine
	}
	return ""
}

func (m *ReadyRequest) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

type TaskMessage struct {
	Id             string            `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Machine        string            `protobuf:"bytes,2,opt,name=machine" json:"machine,omitempty"`
	Node           string            `protobuf:"bytes,3,opt,name=node" json:"node,omitempty"`
	StartTime      int32             `protobuf:"varint,4,opt,name=startTime" json:"startTime,omitempty"`
	Starts         int32             `protobuf:"varint,5,opt,name=starts" json:"starts,omitempty"`
	InstanceParams map[string]string `protobuf:"bytes,6,rep,name=instanceParams" json:"instanceParams,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	NodeParams     map[string]string `protobuf:"bytes,7,rep,name=nodeParams" json:"nodeParams,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *TaskMessage) Reset()                    { *m = TaskMessage{} }
func (m *TaskMessage) String() string            { return proto.CompactTextString(m) }
func (*TaskMessage) ProtoMessage()               {}
func (*TaskMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *TaskMessage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TaskMessage) GetMachine() string {
	if m != nil {
		return m.Machine
	}
	return ""
}

func (m *TaskMessage) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *TaskMessage) GetStartTime() int32 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *TaskMessage) GetStarts() int32 {
	if m != nil {
		return m.Starts
	}
	return 0
}

func (m *TaskMessage) GetInstanceParams() map[string]string {
	if m != nil {
		return m.InstanceParams
	}
	return nil
}

func (m *TaskMessage) GetNodeParams() map[string]string {
	if m != nil {
		return m.NodeParams
	}
	return nil
}

type InstancesRequest struct {
}

func (m *InstancesRequest) Reset()                    { *m = InstancesRequest{} }
func (m *InstancesRequest) String() string            { return proto.CompactTextString(m) }
func (*InstancesRequest) ProtoMessage()               {}
func (*InstancesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type InstancesReply struct {
	Instances []*TaskMessage `protobuf:"bytes,1,rep,name=instances" json:"instances,omitempty"`
}

func (m *InstancesReply) Reset()                    { *m = InstancesReply{} }
func (m *InstancesReply) String() string            { return proto.CompactTextString(m) }
func (*InstancesReply) ProtoMessage()               {}
func (*InstancesReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *InstancesReply) GetInstances() []*TaskMessage {
	if m != nil {
		return m.Instances
	}
	return nil
}

type ChangesRequest struct {
}

func (m *ChangesRequest) Reset()                    { *m = ChangesRequest{} }
func (m *ChangesRequest) String() string            { return proto.CompactTextString(m) }
func (*ChangesRequest) ProtoMessage()               {}
func (*ChangesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type ChangesReply struct {
	Command string `protobuf:"bytes,1,opt,name=command" json:"command,omitempty"`
	Client  string `protobuf:"bytes,2,opt,name=client" json:"client,omitempty"`
}

func (m *ChangesReply) Reset()                    { *m = ChangesReply{} }
func (m *ChangesReply) String() string            { return proto.CompactTextString(m) }
func (*ChangesReply) ProtoMessage()               {}
func (*ChangesReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ChangesReply) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *ChangesReply) GetClient() string {
	if m != nil {
		return m.Client
	}
	return ""
}

func init() {
	proto.RegisterType((*Transition)(nil), "Transition")
	proto.RegisterType((*Node)(nil), "Node")
	proto.RegisterType((*DefineRequest)(nil), "DefineRequest")
	proto.RegisterType((*DefineReply)(nil), "DefineReply")
	proto.RegisterType((*MachinesRequest)(nil), "MachinesRequest")
	proto.RegisterType((*MachinesReply)(nil), "MachinesReply")
	proto.RegisterType((*ReadyRequest)(nil), "ReadyRequest")
	proto.RegisterType((*TaskMessage)(nil), "TaskMessage")
	proto.RegisterType((*InstancesRequest)(nil), "InstancesRequest")
	proto.RegisterType((*InstancesReply)(nil), "InstancesReply")
	proto.RegisterType((*ChangesRequest)(nil), "ChangesRequest")
	proto.RegisterType((*ChangesReply)(nil), "ChangesReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DistributedFSMRunner service

type DistributedFSMRunnerClient interface {
	Define(ctx context.Context, in *DefineRequest, opts ...grpc.CallOption) (*DefineReply, error)
	Start(ctx context.Context, in *TaskMessage, opts ...grpc.CallOption) (*TaskMessage, error)
	Changes(ctx context.Context, in *ChangesRequest, opts ...grpc.CallOption) (DistributedFSMRunner_ChangesClient, error)
	Machines(ctx context.Context, in *MachinesRequest, opts ...grpc.CallOption) (*MachinesReply, error)
	Instances(ctx context.Context, in *InstancesRequest, opts ...grpc.CallOption) (*InstancesReply, error)
	Ready(ctx context.Context, in *ReadyRequest, opts ...grpc.CallOption) (*TaskMessage, error)
}

type distributedFSMRunnerClient struct {
	cc *grpc.ClientConn
}

func NewDistributedFSMRunnerClient(cc *grpc.ClientConn) DistributedFSMRunnerClient {
	return &distributedFSMRunnerClient{cc}
}

func (c *distributedFSMRunnerClient) Define(ctx context.Context, in *DefineRequest, opts ...grpc.CallOption) (*DefineReply, error) {
	out := new(DefineReply)
	err := grpc.Invoke(ctx, "/DistributedFSMRunner/Define", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributedFSMRunnerClient) Start(ctx context.Context, in *TaskMessage, opts ...grpc.CallOption) (*TaskMessage, error) {
	out := new(TaskMessage)
	err := grpc.Invoke(ctx, "/DistributedFSMRunner/Start", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributedFSMRunnerClient) Changes(ctx context.Context, in *ChangesRequest, opts ...grpc.CallOption) (DistributedFSMRunner_ChangesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_DistributedFSMRunner_serviceDesc.Streams[0], c.cc, "/DistributedFSMRunner/Changes", opts...)
	if err != nil {
		return nil, err
	}
	x := &distributedFSMRunnerChangesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DistributedFSMRunner_ChangesClient interface {
	Recv() (*ChangesReply, error)
	grpc.ClientStream
}

type distributedFSMRunnerChangesClient struct {
	grpc.ClientStream
}

func (x *distributedFSMRunnerChangesClient) Recv() (*ChangesReply, error) {
	m := new(ChangesReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *distributedFSMRunnerClient) Machines(ctx context.Context, in *MachinesRequest, opts ...grpc.CallOption) (*MachinesReply, error) {
	out := new(MachinesReply)
	err := grpc.Invoke(ctx, "/DistributedFSMRunner/Machines", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributedFSMRunnerClient) Instances(ctx context.Context, in *InstancesRequest, opts ...grpc.CallOption) (*InstancesReply, error) {
	out := new(InstancesReply)
	err := grpc.Invoke(ctx, "/DistributedFSMRunner/Instances", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributedFSMRunnerClient) Ready(ctx context.Context, in *ReadyRequest, opts ...grpc.CallOption) (*TaskMessage, error) {
	out := new(TaskMessage)
	err := grpc.Invoke(ctx, "/DistributedFSMRunner/Ready", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DistributedFSMRunner service

type DistributedFSMRunnerServer interface {
	Define(context.Context, *DefineRequest) (*DefineReply, error)
	Start(context.Context, *TaskMessage) (*TaskMessage, error)
	Changes(*ChangesRequest, DistributedFSMRunner_ChangesServer) error
	Machines(context.Context, *MachinesRequest) (*MachinesReply, error)
	Instances(context.Context, *InstancesRequest) (*InstancesReply, error)
	Ready(context.Context, *ReadyRequest) (*TaskMessage, error)
}

func RegisterDistributedFSMRunnerServer(s *grpc.Server, srv DistributedFSMRunnerServer) {
	s.RegisterService(&_DistributedFSMRunner_serviceDesc, srv)
}

func _DistributedFSMRunner_Define_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DefineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributedFSMRunnerServer).Define(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DistributedFSMRunner/Define",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributedFSMRunnerServer).Define(ctx, req.(*DefineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DistributedFSMRunner_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributedFSMRunnerServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DistributedFSMRunner/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributedFSMRunnerServer).Start(ctx, req.(*TaskMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _DistributedFSMRunner_Changes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ChangesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DistributedFSMRunnerServer).Changes(m, &distributedFSMRunnerChangesServer{stream})
}

type DistributedFSMRunner_ChangesServer interface {
	Send(*ChangesReply) error
	grpc.ServerStream
}

type distributedFSMRunnerChangesServer struct {
	grpc.ServerStream
}

func (x *distributedFSMRunnerChangesServer) Send(m *ChangesReply) error {
	return x.ServerStream.SendMsg(m)
}

func _DistributedFSMRunner_Machines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MachinesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributedFSMRunnerServer).Machines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DistributedFSMRunner/Machines",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributedFSMRunnerServer).Machines(ctx, req.(*MachinesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DistributedFSMRunner_Instances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributedFSMRunnerServer).Instances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DistributedFSMRunner/Instances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributedFSMRunnerServer).Instances(ctx, req.(*InstancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DistributedFSMRunner_Ready_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributedFSMRunnerServer).Ready(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DistributedFSMRunner/Ready",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributedFSMRunnerServer).Ready(ctx, req.(*ReadyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DistributedFSMRunner_serviceDesc = grpc.ServiceDesc{
	ServiceName: "DistributedFSMRunner",
	HandlerType: (*DistributedFSMRunnerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Define",
			Handler:    _DistributedFSMRunner_Define_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _DistributedFSMRunner_Start_Handler,
		},
		{
			MethodName: "Machines",
			Handler:    _DistributedFSMRunner_Machines_Handler,
		},
		{
			MethodName: "Instances",
			Handler:    _DistributedFSMRunner_Instances_Handler,
		},
		{
			MethodName: "Ready",
			Handler:    _DistributedFSMRunner_Ready_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Changes",
			Handler:       _DistributedFSMRunner_Changes_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "dfsmr.proto",
}

func init() { proto.RegisterFile("dfsmr.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 525 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x6d, 0xd3, 0xa6, 0x1f, 0x37, 0x69, 0xd2, 0x9a, 0x09, 0x45, 0x11, 0x12, 0x95, 0x25, 0xa6,
	0xbe, 0x60, 0xd6, 0x01, 0xd2, 0x84, 0xc4, 0x13, 0x03, 0x89, 0x87, 0x4e, 0xa8, 0xeb, 0x03, 0xaf,
	0x5e, 0xe2, 0x6d, 0xd1, 0x1a, 0x67, 0xd8, 0x2e, 0x52, 0x7f, 0x0e, 0x3f, 0x81, 0x7f, 0x88, 0xe2,
	0xd9, 0x6d, 0x1c, 0x21, 0x78, 0xf4, 0xb5, 0xcf, 0x3d, 0xe7, 0xdc, 0x7b, 0x12, 0x08, 0xf2, 0x5b,
	0x59, 0x0a, 0xf2, 0x28, 0x2a, 0x55, 0xe1, 0x53, 0x80, 0x8d, 0xa0, 0x5c, 0x16, 0xaa, 0xa8, 0x38,
	0x02, 0xf0, 0x8a, 0x3c, 0xe9, 0xce, 0xbb, 0x8b, 0x31, 0x0a, 0xa1, 0xcf, 0xab, 0x9c, 0x25, 0x5e,
	0x7d, 0xc2, 0xef, 0xa0, 0x7f, 0x55, 0xe5, 0xcc, 0x79, 0x31, 0x87, 0x40, 0x1d, 0xb0, 0x32, 0xf1,
	0xe6, 0xbd, 0x45, 0x70, 0x1e, 0x90, 0x63, 0x3f, 0xbc, 0x84, 0xc9, 0x25, 0xbb, 0x2d, 0x38, 0x5b,
	0xb3, 0x1f, 0x3b, 0x26, 0x95, 0x03, 0x3f, 0x01, 0xbf, 0x26, 0xb0, 0x40, 0x9f, 0xd4, 0x04, 0xf8,
	0x3b, 0x04, 0x16, 0xf2, 0xb8, 0xdd, 0xa3, 0x18, 0x86, 0x72, 0x97, 0x65, 0x4c, 0x4a, 0x8d, 0x1a,
	0x99, 0x0e, 0x5a, 0x14, 0x7a, 0x0e, 0x91, 0xb9, 0x5c, 0x31, 0x29, 0xe9, 0x1d, 0x4b, 0x7a, 0xa6,
	0x73, 0xc8, 0x84, 0xa8, 0x84, 0xad, 0xf6, 0xb5, 0x85, 0x19, 0xc4, 0x2b, 0x9a, 0xdd, 0x17, 0x9c,
	0x49, 0x23, 0xa7, 0xd6, 0x77, 0x2c, 0xd5, 0x74, 0x73, 0x18, 0x95, 0xa6, 0x90, 0x74, 0xb5, 0xac,
	0x88, 0x38, 0x0e, 0xf0, 0x47, 0x08, 0xd7, 0x8c, 0xe6, 0x7b, 0xeb, 0x28, 0x82, 0x41, 0xb6, 0x2d,
	0x18, 0x57, 0xc6, 0x55, 0x0c, 0x43, 0xd3, 0xc1, 0x88, 0xb4, 0x73, 0xd4, 0xd2, 0xf0, 0x6f, 0x0f,
	0x82, 0x0d, 0x95, 0x0f, 0x46, 0x9a, 0x33, 0x90, 0x7f, 0x43, 0xd1, 0x0c, 0xc6, 0x52, 0x51, 0xa1,
	0x36, 0x45, 0xf9, 0x64, 0xc9, 0xaf, 0xc9, 0x75, 0x49, 0x26, 0xbe, 0x3e, 0x5f, 0x40, 0x54, 0x70,
	0xa9, 0x28, 0xcf, 0xd8, 0x37, 0x2a, 0x68, 0x29, 0x93, 0x81, 0x36, 0x31, 0x27, 0x0d, 0x4e, 0xf2,
	0xd5, 0x79, 0xf2, 0x99, 0x2b, 0xb1, 0x47, 0x67, 0x00, 0x35, 0x95, 0x41, 0x0d, 0x35, 0xea, 0x85,
	0x83, 0xba, 0x3a, 0x5c, 0x6b, 0x44, 0xfa, 0x1e, 0x9e, 0xfd, 0xad, 0x51, 0x00, 0xbd, 0x07, 0xb6,
	0x37, 0x8e, 0x26, 0xe0, 0xff, 0xa4, 0xdb, 0x9d, 0xf1, 0xf3, 0xc1, 0xbb, 0xe8, 0xa6, 0x4b, 0x88,
	0x5b, 0x9d, 0xfe, 0x07, 0xc1, 0x08, 0xa6, 0x96, 0xa9, 0xb1, 0xb9, 0xa8, 0x51, 0xab, 0x57, 0xf7,
	0x12, 0xc6, 0xd6, 0xbb, 0xdd, 0x5d, 0xd8, 0x34, 0x80, 0xa7, 0x10, 0x7d, 0xba, 0xa7, 0xfc, 0xee,
	0xd8, 0xe4, 0x0d, 0x84, 0x87, 0x8a, 0x09, 0x5b, 0x56, 0x95, 0x25, 0xe5, 0x76, 0x23, 0xc7, 0xe5,
	0x6a, 0x35, 0xe7, 0xbf, 0x3c, 0x38, 0xb9, 0x2c, 0xa4, 0x12, 0xc5, 0xcd, 0x4e, 0xb1, 0xfc, 0xcb,
	0xf5, 0x6a, 0xbd, 0xe3, 0x9c, 0x09, 0xb4, 0x80, 0xc1, 0x53, 0x4c, 0x50, 0x2b, 0x2f, 0x69, 0x48,
	0x1a, 0x71, 0xc6, 0x1d, 0xf4, 0x0a, 0xfc, 0xeb, 0x7a, 0x65, 0xc8, 0x11, 0x97, 0xba, 0x52, 0x3b,
	0xe8, 0x35, 0x0c, 0x8d, 0x34, 0x14, 0x13, 0x57, 0x76, 0x3a, 0x21, 0x4d, 0xd5, 0xb8, 0x73, 0xd6,
	0x45, 0x04, 0x46, 0x36, 0xc8, 0x68, 0x4a, 0x5a, 0x31, 0x4f, 0x23, 0xe2, 0xa4, 0x1c, 0x77, 0xd0,
	0x12, 0xc6, 0x87, 0xf1, 0xa1, 0x19, 0x69, 0x8f, 0x37, 0x8d, 0x89, 0x3b, 0x5d, 0xdc, 0x41, 0xa7,
	0xe0, 0xeb, 0xe0, 0xa3, 0x09, 0x69, 0x7e, 0x00, 0x6d, 0xe5, 0x37, 0x03, 0xfd, 0x63, 0x79, 0xfb,
	0x27, 0x00, 0x00, 0xff, 0xff, 0x42, 0xf6, 0x5e, 0xe1, 0x67, 0x04, 0x00, 0x00,
}
