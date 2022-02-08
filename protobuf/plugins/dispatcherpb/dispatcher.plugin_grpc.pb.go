// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.2
// source: protobuf/dispatcher.plugin.proto

package dispatcherpb

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

// DispatcherPluginClient is the client API for DispatcherPlugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DispatcherPluginClient interface {
	// Before the task is started, you can change the task and/or the trigger metadata
	OnStartNewTask(ctx context.Context, in *OnStartNewTaskRequestResponse, opts ...grpc.CallOption) (*OnStartNewTaskRequestResponse, error)
	// Before a request is dispatched, you can change it or discard it
	OnNewRequest(ctx context.Context, in *OnNewRequestRequest, opts ...grpc.CallOption) (*OnNewRequestResponse, error)
}

type dispatcherPluginClient struct {
	cc grpc.ClientConnInterface
}

func NewDispatcherPluginClient(cc grpc.ClientConnInterface) DispatcherPluginClient {
	return &dispatcherPluginClient{cc}
}

func (c *dispatcherPluginClient) OnStartNewTask(ctx context.Context, in *OnStartNewTaskRequestResponse, opts ...grpc.CallOption) (*OnStartNewTaskRequestResponse, error) {
	out := new(OnStartNewTaskRequestResponse)
	err := c.cc.Invoke(ctx, "/nodeum.protobuf.DispatcherPlugin/OnStartNewTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherPluginClient) OnNewRequest(ctx context.Context, in *OnNewRequestRequest, opts ...grpc.CallOption) (*OnNewRequestResponse, error) {
	out := new(OnNewRequestResponse)
	err := c.cc.Invoke(ctx, "/nodeum.protobuf.DispatcherPlugin/OnNewRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DispatcherPluginServer is the server API for DispatcherPlugin service.
// All implementations must embed UnimplementedDispatcherPluginServer
// for forward compatibility
type DispatcherPluginServer interface {
	// Before the task is started, you can change the task and/or the trigger metadata
	OnStartNewTask(context.Context, *OnStartNewTaskRequestResponse) (*OnStartNewTaskRequestResponse, error)
	// Before a request is dispatched, you can change it or discard it
	OnNewRequest(context.Context, *OnNewRequestRequest) (*OnNewRequestResponse, error)
	mustEmbedUnimplementedDispatcherPluginServer()
}

// UnimplementedDispatcherPluginServer must be embedded to have forward compatible implementations.
type UnimplementedDispatcherPluginServer struct {
}

func (UnimplementedDispatcherPluginServer) OnStartNewTask(context.Context, *OnStartNewTaskRequestResponse) (*OnStartNewTaskRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnStartNewTask not implemented")
}
func (UnimplementedDispatcherPluginServer) OnNewRequest(context.Context, *OnNewRequestRequest) (*OnNewRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnNewRequest not implemented")
}
func (UnimplementedDispatcherPluginServer) mustEmbedUnimplementedDispatcherPluginServer() {}

// UnsafeDispatcherPluginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DispatcherPluginServer will
// result in compilation errors.
type UnsafeDispatcherPluginServer interface {
	mustEmbedUnimplementedDispatcherPluginServer()
}

func RegisterDispatcherPluginServer(s grpc.ServiceRegistrar, srv DispatcherPluginServer) {
	s.RegisterService(&DispatcherPlugin_ServiceDesc, srv)
}

func _DispatcherPlugin_OnStartNewTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnStartNewTaskRequestResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherPluginServer).OnStartNewTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodeum.protobuf.DispatcherPlugin/OnStartNewTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherPluginServer).OnStartNewTask(ctx, req.(*OnStartNewTaskRequestResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatcherPlugin_OnNewRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnNewRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherPluginServer).OnNewRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodeum.protobuf.DispatcherPlugin/OnNewRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherPluginServer).OnNewRequest(ctx, req.(*OnNewRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DispatcherPlugin_ServiceDesc is the grpc.ServiceDesc for DispatcherPlugin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DispatcherPlugin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nodeum.protobuf.DispatcherPlugin",
	HandlerType: (*DispatcherPluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OnStartNewTask",
			Handler:    _DispatcherPlugin_OnStartNewTask_Handler,
		},
		{
			MethodName: "OnNewRequest",
			Handler:    _DispatcherPlugin_OnNewRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/dispatcher.plugin.proto",
}
