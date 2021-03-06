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
	TransitionRequest
	TransitionReply
	RelinquishRequest
	RelinquishReply
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

type TransitionRequest struct {
	Instance   string            `protobuf:"bytes,1,opt,name=instance" json:"instance,omitempty"`
	Transition string            `protobuf:"bytes,2,opt,name=transition" json:"transition,omitempty"`
	NodeParams map[string]string `protobuf:"bytes,3,rep,name=nodeParams" json:"nodeParams,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *TransitionRequest) Reset()                    { *m = TransitionRequest{} }
func (m *TransitionRequest) String() string            { return proto.CompactTextString(m) }
func (*TransitionRequest) ProtoMessage()               {}
func (*TransitionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *TransitionRequest) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

func (m *TransitionRequest) GetTransition() string {
	if m != nil {
		return m.Transition
	}
	return ""
}

func (m *TransitionRequest) GetNodeParams() map[string]string {
	if m != nil {
		return m.NodeParams
	}
	return nil
}

type TransitionReply struct {
	Instance string `protobuf:"bytes,1,opt,name=instance" json:"instance,omitempty"`
	Node     string `protobuf:"bytes,2,opt,name=node" json:"node,omitempty"`
	PrevNode string `protobuf:"bytes,3,opt,name=prevNode" json:"prevNode,omitempty"`
}

func (m *TransitionReply) Reset()                    { *m = TransitionReply{} }
func (m *TransitionReply) String() string            { return proto.CompactTextString(m) }
func (*TransitionReply) ProtoMessage()               {}
func (*TransitionReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *TransitionReply) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

func (m *TransitionReply) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *TransitionReply) GetPrevNode() string {
	if m != nil {
		return m.PrevNode
	}
	return ""
}

type RelinquishRequest struct {
	Instance string `protobuf:"bytes,1,opt,name=instance" json:"instance,omitempty"`
}

func (m *RelinquishRequest) Reset()                    { *m = RelinquishRequest{} }
func (m *RelinquishRequest) String() string            { return proto.CompactTextString(m) }
func (*RelinquishRequest) ProtoMessage()               {}
func (*RelinquishRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *RelinquishRequest) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

type RelinquishReply struct {
	Istance string `protobuf:"bytes,1,opt,name=istance" json:"istance,omitempty"`
	Node    string `protobuf:"bytes,2,opt,name=node" json:"node,omitempty"`
}

func (m *RelinquishReply) Reset()                    { *m = RelinquishReply{} }
func (m *RelinquishReply) String() string            { return proto.CompactTextString(m) }
func (*RelinquishReply) ProtoMessage()               {}
func (*RelinquishReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *RelinquishReply) GetIstance() string {
	if m != nil {
		return m.Istance
	}
	return ""
}

func (m *RelinquishReply) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

type InstancesRequest struct {
}

func (m *InstancesRequest) Reset()                    { *m = InstancesRequest{} }
func (m *InstancesRequest) String() string            { return proto.CompactTextString(m) }
func (*InstancesRequest) ProtoMessage()               {}
func (*InstancesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type InstancesReply struct {
	Instances []*TaskMessage `protobuf:"bytes,1,rep,name=instances" json:"instances,omitempty"`
}

func (m *InstancesReply) Reset()                    { *m = InstancesReply{} }
func (m *InstancesReply) String() string            { return proto.CompactTextString(m) }
func (*InstancesReply) ProtoMessage()               {}
func (*InstancesReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

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
func (*ChangesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

type ChangesReply struct {
	Command string `protobuf:"bytes,1,opt,name=command" json:"command,omitempty"`
	Client  string `protobuf:"bytes,2,opt,name=client" json:"client,omitempty"`
}

func (m *ChangesReply) Reset()                    { *m = ChangesReply{} }
func (m *ChangesReply) String() string            { return proto.CompactTextString(m) }
func (*ChangesReply) ProtoMessage()               {}
func (*ChangesReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

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
	proto.RegisterType((*TransitionRequest)(nil), "TransitionRequest")
	proto.RegisterType((*TransitionReply)(nil), "TransitionReply")
	proto.RegisterType((*RelinquishRequest)(nil), "RelinquishRequest")
	proto.RegisterType((*RelinquishReply)(nil), "RelinquishReply")
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
	Transition(ctx context.Context, in *TransitionRequest, opts ...grpc.CallOption) (*TransitionReply, error)
	Relinquish(ctx context.Context, in *RelinquishRequest, opts ...grpc.CallOption) (*RelinquishReply, error)
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

func (c *distributedFSMRunnerClient) Transition(ctx context.Context, in *TransitionRequest, opts ...grpc.CallOption) (*TransitionReply, error) {
	out := new(TransitionReply)
	err := grpc.Invoke(ctx, "/DistributedFSMRunner/Transition", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributedFSMRunnerClient) Relinquish(ctx context.Context, in *RelinquishRequest, opts ...grpc.CallOption) (*RelinquishReply, error) {
	out := new(RelinquishReply)
	err := grpc.Invoke(ctx, "/DistributedFSMRunner/Relinquish", in, out, c.cc, opts...)
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
	Transition(context.Context, *TransitionRequest) (*TransitionReply, error)
	Relinquish(context.Context, *RelinquishRequest) (*RelinquishReply, error)
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

func _DistributedFSMRunner_Transition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributedFSMRunnerServer).Transition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DistributedFSMRunner/Transition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributedFSMRunnerServer).Transition(ctx, req.(*TransitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DistributedFSMRunner_Relinquish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RelinquishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributedFSMRunnerServer).Relinquish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DistributedFSMRunner/Relinquish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributedFSMRunnerServer).Relinquish(ctx, req.(*RelinquishRequest))
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
		{
			MethodName: "Transition",
			Handler:    _DistributedFSMRunner_Transition_Handler,
		},
		{
			MethodName: "Relinquish",
			Handler:    _DistributedFSMRunner_Relinquish_Handler,
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
	// 663 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xdd, 0x6e, 0xd3, 0x30,
	0x18, 0x6d, 0xda, 0xa6, 0x3f, 0x5f, 0xda, 0xa4, 0x35, 0x13, 0x8a, 0x22, 0x24, 0x22, 0x4b, 0x9b,
	0x7a, 0x83, 0xd9, 0xc6, 0x8f, 0x26, 0x24, 0xae, 0x36, 0x90, 0xb8, 0xd8, 0x84, 0xba, 0x5d, 0x70,
	0xeb, 0xb5, 0xde, 0x66, 0xad, 0x75, 0x36, 0x3b, 0x1d, 0xea, 0x43, 0xf0, 0x16, 0x3c, 0x08, 0x8f,
	0x86, 0xec, 0x39, 0x4d, 0xdc, 0x30, 0x90, 0xb8, 0xb4, 0xe3, 0x73, 0xbe, 0x73, 0x3e, 0x9f, 0xcf,
	0x81, 0x60, 0x7e, 0xa5, 0x96, 0x92, 0xdc, 0xc9, 0x2c, 0xcf, 0xf0, 0x1e, 0xc0, 0x85, 0xa4, 0x42,
	0xf1, 0x9c, 0x67, 0x02, 0x01, 0x34, 0xf9, 0x3c, 0xf6, 0x52, 0x6f, 0xd2, 0x47, 0x03, 0x68, 0x8b,
	0x6c, 0xce, 0xe2, 0xa6, 0x5e, 0xe1, 0x29, 0xb4, 0xcf, 0xb2, 0x39, 0x73, 0x4e, 0xa4, 0x10, 0xe4,
	0x1b, 0xac, 0x8a, 0x9b, 0x69, 0x6b, 0x12, 0x1c, 0x06, 0xa4, 0xc2, 0x37, 0x04, 0x5f, 0xe5, 0x54,
	0xe6, 0x71, 0x2b, 0xf5, 0x26, 0x3d, 0xbd, 0xbc, 0xe2, 0x82, 0x2e, 0xe2, 0xb6, 0x5e, 0xe2, 0x03,
	0x18, 0x9e, 0xb0, 0x2b, 0x2e, 0xd8, 0x94, 0xdd, 0xaf, 0x98, 0xca, 0x1d, 0xf2, 0x1d, 0xf0, 0x75,
	0xf9, 0x82, 0xd6, 0x27, 0xba, 0x3c, 0xfe, 0x06, 0x41, 0x01, 0xb9, 0x5b, 0xac, 0x51, 0x04, 0x5d,
	0xb5, 0x9a, 0xcd, 0x98, 0x52, 0x06, 0xd5, 0xb3, 0x0c, 0x46, 0x32, 0x7a, 0x0e, 0xa1, 0xfd, 0x78,
	0xca, 0x94, 0xa2, 0xd7, 0xcc, 0xa8, 0xd0, 0xcc, 0x03, 0x26, 0x65, 0x26, 0x8b, 0xdd, 0xb6, 0x31,
	0x38, 0x86, 0xe8, 0x94, 0xce, 0x6e, 0xb8, 0x60, 0xca, 0xca, 0xd1, 0xfa, 0xca, 0x2d, 0x5d, 0x2e,
	0x85, 0xde, 0xd2, 0x6e, 0xc4, 0x9e, 0x91, 0x15, 0x12, 0xc7, 0x01, 0xfe, 0x08, 0x83, 0x29, 0xa3,
	0xf3, 0x75, 0xe1, 0x28, 0x84, 0xce, 0x6c, 0xc1, 0x99, 0xc8, 0xad, 0xab, 0x08, 0xba, 0x96, 0xc1,
	0x8a, 0x2c, 0xba, 0x6c, 0xa4, 0xe1, 0x5f, 0x4d, 0x08, 0x2e, 0xa8, 0xba, 0xb5, 0xd2, 0x9c, 0x86,
	0xfc, 0x1d, 0xaa, 0x7b, 0x9b, 0x7d, 0x17, 0x4c, 0x3e, 0xda, 0x41, 0x63, 0xe8, 0x9b, 0xce, 0x5f,
	0xf0, 0x25, 0x8b, 0xfd, 0xd4, 0x9b, 0xf8, 0x5a, 0x8b, 0xd9, 0x52, 0x71, 0xc7, 0xac, 0x8f, 0x20,
	0xe4, 0x42, 0xe5, 0x54, 0xcc, 0xd8, 0x57, 0x2a, 0xe9, 0x52, 0xc5, 0x5d, 0xe3, 0x29, 0x25, 0x15,
	0x09, 0xe4, 0x8b, 0x73, 0xe4, 0x93, 0xc8, 0xe5, 0x1a, 0xed, 0x03, 0xe8, 0xca, 0x16, 0xd5, 0x33,
	0xa8, 0x17, 0x0e, 0xea, 0x6c, 0xf3, 0xd9, 0x20, 0x92, 0x77, 0xf0, 0xec, 0x4f, 0x44, 0x01, 0xb4,
	0x6e, 0xd9, 0xda, 0x1a, 0x1c, 0x82, 0xff, 0x40, 0x17, 0x2b, 0x6b, 0xef, 0x43, 0xf3, 0xc8, 0x4b,
	0x0e, 0x20, 0xda, 0x62, 0xfa, 0x17, 0x04, 0xff, 0xf4, 0x60, 0x5c, 0x26, 0xb0, 0xb8, 0x87, 0x11,
	0xf4, 0x0a, 0xaf, 0x16, 0x8a, 0x00, 0xca, 0xf0, 0xda, 0x8e, 0xbe, 0x77, 0x7c, 0xb5, 0x8c, 0x2f,
	0x4c, 0x6a, 0x6c, 0x35, 0x77, 0xff, 0x21, 0xf3, 0x18, 0xa2, 0x2a, 0xaf, 0x4e, 0x57, 0x5d, 0xa3,
	0x33, 0x82, 0xfa, 0xfb, 0x9d, 0x64, 0x0f, 0x67, 0x65, 0x5c, 0x76, 0x61, 0x3c, 0x65, 0x0b, 0x2e,
	0xee, 0x57, 0x5c, 0xdd, 0x3c, 0x69, 0x15, 0xef, 0x43, 0x54, 0x3d, 0x66, 0x07, 0x87, 0x3f, 0x5d,
	0x0a, 0x23, 0x18, 0x15, 0xd7, 0x55, 0x99, 0x86, 0xb0, 0xb2, 0xa7, 0x49, 0x5e, 0x42, 0xbf, 0xa8,
	0x54, 0xcc, 0xc3, 0xa0, 0x9a, 0x02, 0x3c, 0x82, 0xf0, 0xf8, 0x86, 0x8a, 0xeb, 0x92, 0xe4, 0x35,
	0x0c, 0x36, 0x3b, 0x56, 0xc7, 0x2c, 0x5b, 0x2e, 0xa9, 0x28, 0x52, 0x5e, 0x0e, 0x8c, 0x51, 0x72,
	0xf8, 0xa3, 0x05, 0x3b, 0x27, 0x5c, 0xe5, 0x92, 0x5f, 0xae, 0x72, 0x36, 0xff, 0x7c, 0x7e, 0x3a,
	0x5d, 0x09, 0xc1, 0x24, 0x9a, 0x40, 0xe7, 0x71, 0xf4, 0xd0, 0xd6, 0x0c, 0x26, 0x03, 0x52, 0x79,
	0x22, 0x70, 0x03, 0xed, 0x82, 0x7f, 0xae, 0x73, 0x8f, 0x1c, 0x71, 0x89, 0x2b, 0xb5, 0x81, 0xf6,
	0xc0, 0x37, 0xa3, 0x8b, 0x86, 0xa4, 0x3a, 0xc2, 0xb5, 0x73, 0xaf, 0xa0, 0x6b, 0x2d, 0xa0, 0x88,
	0xb8, 0xf6, 0x92, 0x21, 0xa9, 0xba, 0xc3, 0x8d, 0x7d, 0x0f, 0x11, 0xe8, 0x15, 0x8f, 0x08, 0x1a,
	0x91, 0xad, 0x27, 0x26, 0x09, 0x89, 0xf3, 0xc2, 0xe0, 0x06, 0x3a, 0x80, 0xfe, 0xa6, 0xcd, 0x68,
	0x4c, 0xb6, 0xaf, 0x21, 0x89, 0x88, 0x7b, 0x0b, 0xb8, 0x81, 0xde, 0x3a, 0x6f, 0x38, 0xaa, 0x07,
	0x36, 0x19, 0x91, 0xad, 0xb0, 0x3d, 0xa2, 0xca, 0x54, 0x20, 0x44, 0x6a, 0x49, 0x4a, 0x46, 0x64,
	0x2b, 0x36, 0xb8, 0x71, 0xd9, 0x31, 0xbf, 0x8d, 0x37, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x8f,
	0xd3, 0x8c, 0x56, 0x45, 0x06, 0x00, 0x00,
}
