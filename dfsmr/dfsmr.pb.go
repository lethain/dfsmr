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
	Start       bool          `protobuf:"varint,3,opt,name=start" json:"start,omitempty"`
	Final       bool          `protobuf:"varint,4,opt,name=final" json:"final,omitempty"`
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

func (m *Node) GetStart() bool {
	if m != nil {
		return m.Start
	}
	return false
}

func (m *Node) GetFinal() bool {
	if m != nil {
		return m.Final
	}
	return false
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
	Owner          string            `protobuf:"bytes,4,opt,name=owner" json:"owner,omitempty"`
	StartTime      int32             `protobuf:"varint,5,opt,name=startTime" json:"startTime,omitempty"`
	Starts         int32             `protobuf:"varint,6,opt,name=starts" json:"starts,omitempty"`
	InstanceParams map[string]string `protobuf:"bytes,7,rep,name=instanceParams" json:"instanceParams,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	NodeParams     map[string]string `protobuf:"bytes,8,rep,name=nodeParams" json:"nodeParams,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
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

func (m *TaskMessage) GetOwner() string {
	if m != nil {
		return m.Owner
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
	Ready(ctx context.Context, in *ReadyRequest, opts ...grpc.CallOption) (*TaskMessage, error)
	Changes(ctx context.Context, in *ChangesRequest, opts ...grpc.CallOption) (DistributedFSMRunner_ChangesClient, error)
	Machines(ctx context.Context, in *MachinesRequest, opts ...grpc.CallOption) (*MachinesReply, error)
	Instances(ctx context.Context, in *InstancesRequest, opts ...grpc.CallOption) (*InstancesReply, error)
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

func (c *distributedFSMRunnerClient) Ready(ctx context.Context, in *ReadyRequest, opts ...grpc.CallOption) (*TaskMessage, error) {
	out := new(TaskMessage)
	err := grpc.Invoke(ctx, "/DistributedFSMRunner/Ready", in, out, c.cc, opts...)
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

// Server API for DistributedFSMRunner service

type DistributedFSMRunnerServer interface {
	Define(context.Context, *DefineRequest) (*DefineReply, error)
	Start(context.Context, *TaskMessage) (*TaskMessage, error)
	Ready(context.Context, *ReadyRequest) (*TaskMessage, error)
	Changes(*ChangesRequest, DistributedFSMRunner_ChangesServer) error
	Machines(context.Context, *MachinesRequest) (*MachinesReply, error)
	Instances(context.Context, *InstancesRequest) (*InstancesReply, error)
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
			MethodName: "Ready",
			Handler:    _DistributedFSMRunner_Ready_Handler,
		},
		{
			MethodName: "Machines",
			Handler:    _DistributedFSMRunner_Machines_Handler,
		},
		{
			MethodName: "Instances",
			Handler:    _DistributedFSMRunner_Instances_Handler,
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
	// 554 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x6d, 0xd2, 0xa6, 0x1f, 0x37, 0x69, 0xd2, 0x9a, 0x09, 0x45, 0x11, 0x12, 0x91, 0x25, 0xa6,
	0xbe, 0x60, 0xd6, 0x21, 0xa4, 0x09, 0x89, 0x27, 0x06, 0x12, 0x0f, 0x9d, 0x50, 0xd7, 0x07, 0x5e,
	0xbd, 0xc4, 0xdd, 0xa2, 0x35, 0xce, 0xb0, 0x53, 0x50, 0x7f, 0x0e, 0xff, 0x84, 0x9f, 0x86, 0xec,
	0x39, 0x6d, 0x5c, 0x10, 0x7b, 0xf4, 0x8d, 0xcf, 0xb9, 0xe7, 0xdc, 0x73, 0x1d, 0xf0, 0xf3, 0xb5,
	0x2c, 0x05, 0x79, 0x10, 0x55, 0x5d, 0xe1, 0x53, 0x80, 0x95, 0xa0, 0x5c, 0x16, 0x75, 0x51, 0x71,
	0x04, 0xe0, 0x16, 0x79, 0xec, 0xa4, 0xce, 0x6c, 0x84, 0x02, 0xe8, 0xf1, 0x2a, 0x67, 0xb1, 0xab,
	0x4e, 0x78, 0x09, 0xbd, 0xab, 0x2a, 0x67, 0xd6, 0x8d, 0x14, 0xfc, 0x7a, 0x8f, 0x95, 0xb1, 0x9b,
	0x76, 0x67, 0xfe, 0xb9, 0x4f, 0x5a, 0x7c, 0x63, 0xf0, 0x64, 0x4d, 0x45, 0x1d, 0x77, 0x53, 0x67,
	0x36, 0x54, 0xc7, 0x75, 0xc1, 0xe9, 0x26, 0xee, 0xa9, 0x23, 0x9e, 0xc3, 0xf8, 0x92, 0xad, 0x0b,
	0xce, 0x96, 0xec, 0xfb, 0x96, 0xc9, 0xda, 0x22, 0x3f, 0x01, 0x4f, 0xb5, 0x6f, 0x68, 0x3d, 0xa2,
	0xda, 0xe3, 0x6f, 0xe0, 0x37, 0x90, 0x87, 0xcd, 0x0e, 0x45, 0x30, 0x90, 0xdb, 0x2c, 0x63, 0x52,
	0x6a, 0xd4, 0xd0, 0x30, 0x68, 0xc9, 0xe8, 0x39, 0x84, 0xe6, 0xe3, 0x82, 0x49, 0x49, 0x6f, 0x99,
	0x56, 0xa1, 0x98, 0x03, 0x26, 0x44, 0x25, 0x9a, 0x6a, 0x4f, 0x1b, 0x9c, 0x42, 0xb4, 0xa0, 0xd9,
	0x5d, 0xc1, 0x99, 0x34, 0x72, 0x94, 0xbe, 0x43, 0x49, 0xb5, 0x4b, 0x61, 0x58, 0x9a, 0x42, 0xec,
	0x68, 0x59, 0x21, 0xb1, 0x1c, 0xe0, 0x0f, 0x10, 0x2c, 0x19, 0xcd, 0x77, 0x8d, 0xa3, 0x10, 0xfa,
	0xd9, 0xa6, 0x60, 0xbc, 0x36, 0xae, 0x22, 0x18, 0x18, 0x06, 0x23, 0xb2, 0x99, 0xb2, 0x96, 0x86,
	0x7f, 0xbb, 0xe0, 0xaf, 0xa8, 0xbc, 0x37, 0xd2, 0xac, 0x81, 0xfc, 0x1f, 0xaa, 0x66, 0x5b, 0xfd,
	0xe4, 0x4c, 0x3c, 0xda, 0x41, 0x53, 0x18, 0xe9, 0xc9, 0xaf, 0x8a, 0x92, 0xc5, 0x5e, 0xea, 0xcc,
	0x3c, 0xa5, 0x45, 0x97, 0x64, 0xdc, 0xd7, 0xe7, 0x0b, 0x08, 0x0b, 0x2e, 0x6b, 0xca, 0x33, 0xf6,
	0x95, 0x0a, 0x5a, 0xca, 0x78, 0xa0, 0x3d, 0xa5, 0xa4, 0x25, 0x81, 0x7c, 0xb1, 0xae, 0x7c, 0xe2,
	0xb5, 0xd8, 0xa1, 0x33, 0x00, 0xd5, 0xd9, 0xa0, 0x86, 0x1a, 0xf5, 0xc2, 0x42, 0x5d, 0xed, 0x3f,
	0x6b, 0x44, 0xf2, 0x0e, 0x9e, 0xfd, 0x8b, 0xc8, 0x87, 0xee, 0x3d, 0xdb, 0x19, 0x83, 0x63, 0xf0,
	0x7e, 0xd0, 0xcd, 0xd6, 0xd8, 0x7b, 0xef, 0x5e, 0x38, 0xc9, 0x1c, 0xa2, 0x23, 0xa6, 0xa7, 0x20,
	0x18, 0xc1, 0xa4, 0xe9, 0xd4, 0x0a, 0x32, 0x6c, 0xd5, 0x54, 0x92, 0x2f, 0x61, 0xd4, 0x78, 0x6f,
	0xa2, 0x0c, 0xda, 0x06, 0xf0, 0x04, 0xc2, 0x8f, 0x77, 0x94, 0xdf, 0x1e, 0x48, 0xde, 0x40, 0xb0,
	0xaf, 0x98, 0xdd, 0xcb, 0xaa, 0xb2, 0xa4, 0xbc, 0x09, 0xe8, 0x90, 0xb5, 0x56, 0x73, 0xfe, 0xcb,
	0x85, 0x93, 0xcb, 0x42, 0xd6, 0xa2, 0xb8, 0xd9, 0xd6, 0x2c, 0xff, 0x7c, 0xbd, 0x58, 0x6e, 0x39,
	0x67, 0x02, 0xcd, 0xa0, 0xff, 0xb8, 0x35, 0xe8, 0x68, 0x7d, 0x92, 0x80, 0xb4, 0xb6, 0x1b, 0x77,
	0xd0, 0x2b, 0xf0, 0xae, 0x55, 0x64, 0xc8, 0x12, 0x97, 0xd8, 0x52, 0x3b, 0xe8, 0x14, 0x3c, 0xbd,
	0x75, 0x68, 0x4c, 0xda, 0xdb, 0xf7, 0xd7, 0xbd, 0xd7, 0x30, 0x30, 0x16, 0x50, 0x44, 0x6c, 0x7b,
	0xc9, 0x98, 0xb4, 0xdd, 0xe1, 0xce, 0x99, 0x83, 0x08, 0x0c, 0x9b, 0xfd, 0x47, 0x13, 0x72, 0xf4,
	0x3a, 0x92, 0x90, 0x58, 0x8f, 0x03, 0x77, 0xd0, 0x1c, 0x46, 0xfb, 0x31, 0xa3, 0x29, 0x39, 0x8e,
	0x21, 0x89, 0x88, 0x9d, 0x02, 0xee, 0xdc, 0xf4, 0xf5, 0x5f, 0xe8, 0xed, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xbc, 0x03, 0x32, 0x86, 0x94, 0x04, 0x00, 0x00,
}
