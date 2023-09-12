// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: agent_grpc.proto

package agentGrpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CmdServiceClient is the client API for CmdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CmdServiceClient interface {
	// execute a remote command
	ExecRemoteCmd(ctx context.Context, opts ...grpc.CallOption) (CmdService_ExecRemoteCmdClient, error)
}

type cmdServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCmdServiceClient(cc grpc.ClientConnInterface) CmdServiceClient {
	return &cmdServiceClient{cc}
}

func (c *cmdServiceClient) ExecRemoteCmd(ctx context.Context, opts ...grpc.CallOption) (CmdService_ExecRemoteCmdClient, error) {
	stream, err := c.cc.NewStream(ctx, &CmdService_ServiceDesc.Streams[0], "/proto.CmdService/ExecRemoteCmd", opts...)
	if err != nil {
		return nil, err
	}
	x := &cmdServiceExecRemoteCmdClient{stream}
	return x, nil
}

type CmdService_ExecRemoteCmdClient interface {
	Send(*ExecRequestMsg) error
	Recv() (*ExecResponseMsg, error)
	grpc.ClientStream
}

type cmdServiceExecRemoteCmdClient struct {
	grpc.ClientStream
}

func (x *cmdServiceExecRemoteCmdClient) Send(m *ExecRequestMsg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *cmdServiceExecRemoteCmdClient) Recv() (*ExecResponseMsg, error) {
	m := new(ExecResponseMsg)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CmdServiceServer is the server API for CmdService service.
// All implementations must embed UnimplementedCmdServiceServer
// for forward compatibility
type CmdServiceServer interface {
	// execute a remote command
	ExecRemoteCmd(CmdService_ExecRemoteCmdServer) error
	mustEmbedUnimplementedCmdServiceServer()
}

// UnimplementedCmdServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCmdServiceServer struct {
}

func (UnimplementedCmdServiceServer) ExecRemoteCmd(CmdService_ExecRemoteCmdServer) error {
	return status.Errorf(codes.Unimplemented, "method ExecRemoteCmd not implemented")
}
func (UnimplementedCmdServiceServer) mustEmbedUnimplementedCmdServiceServer() {}

// UnsafeCmdServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CmdServiceServer will
// result in compilation errors.
type UnsafeCmdServiceServer interface {
	mustEmbedUnimplementedCmdServiceServer()
}

func RegisterCmdServiceServer(s grpc.ServiceRegistrar, srv CmdServiceServer) {
	s.RegisterService(&CmdService_ServiceDesc, srv)
}

func _CmdService_ExecRemoteCmd_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CmdServiceServer).ExecRemoteCmd(&cmdServiceExecRemoteCmdServer{stream})
}

type CmdService_ExecRemoteCmdServer interface {
	Send(*ExecResponseMsg) error
	Recv() (*ExecRequestMsg, error)
	grpc.ServerStream
}

type cmdServiceExecRemoteCmdServer struct {
	grpc.ServerStream
}

func (x *cmdServiceExecRemoteCmdServer) Send(m *ExecResponseMsg) error {
	return x.ServerStream.SendMsg(m)
}

func (x *cmdServiceExecRemoteCmdServer) Recv() (*ExecRequestMsg, error) {
	m := new(ExecRequestMsg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CmdService_ServiceDesc is the grpc.ServiceDesc for CmdService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CmdService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CmdService",
	HandlerType: (*CmdServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ExecRemoteCmd",
			Handler:       _CmdService_ExecRemoteCmd_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "agent_grpc.proto",
}