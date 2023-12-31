// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// BackendClient is the client API for Backend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BackendClient interface {
	ConfigureBackend(ctx context.Context, in *ConfigureBackend_Request, opts ...grpc.CallOption) (*Empty, error)
	DeleteWorkspace(ctx context.Context, in *DeleteWorkspace_Request, opts ...grpc.CallOption) (*Empty, error)
	ListWorkspaces(ctx context.Context, in *ListWorkspaces_Request, opts ...grpc.CallOption) (*ListWorkspaces_Response, error)
	GetStatePayload(ctx context.Context, in *GetStatePayload_Request, opts ...grpc.CallOption) (*GetStatePayload_Response, error)
	PutState(ctx context.Context, in *PutState_Request, opts ...grpc.CallOption) (*Empty, error)
	DeleteState(ctx context.Context, in *DeleteState_Request, opts ...grpc.CallOption) (*Empty, error)
	LockState(ctx context.Context, in *StateLock_Request, opts ...grpc.CallOption) (*StateLock_Response, error)
	UnlockState(ctx context.Context, in *StateUnlock_Request, opts ...grpc.CallOption) (*Empty, error)
}

type backendClient struct {
	cc grpc.ClientConnInterface
}

func NewBackendClient(cc grpc.ClientConnInterface) BackendClient {
	return &backendClient{cc}
}

func (c *backendClient) ConfigureBackend(ctx context.Context, in *ConfigureBackend_Request, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/Backend/ConfigureBackend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) DeleteWorkspace(ctx context.Context, in *DeleteWorkspace_Request, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/Backend/DeleteWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) ListWorkspaces(ctx context.Context, in *ListWorkspaces_Request, opts ...grpc.CallOption) (*ListWorkspaces_Response, error) {
	out := new(ListWorkspaces_Response)
	err := c.cc.Invoke(ctx, "/Backend/ListWorkspaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) GetStatePayload(ctx context.Context, in *GetStatePayload_Request, opts ...grpc.CallOption) (*GetStatePayload_Response, error) {
	out := new(GetStatePayload_Response)
	err := c.cc.Invoke(ctx, "/Backend/GetStatePayload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) PutState(ctx context.Context, in *PutState_Request, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/Backend/PutState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) DeleteState(ctx context.Context, in *DeleteState_Request, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/Backend/DeleteState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) LockState(ctx context.Context, in *StateLock_Request, opts ...grpc.CallOption) (*StateLock_Response, error) {
	out := new(StateLock_Response)
	err := c.cc.Invoke(ctx, "/Backend/LockState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) UnlockState(ctx context.Context, in *StateUnlock_Request, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/Backend/UnlockState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BackendServer is the server API for Backend service.
// All implementations must embed UnimplementedBackendServer
// for forward compatibility
type BackendServer interface {
	ConfigureBackend(context.Context, *ConfigureBackend_Request) (*Empty, error)
	DeleteWorkspace(context.Context, *DeleteWorkspace_Request) (*Empty, error)
	ListWorkspaces(context.Context, *ListWorkspaces_Request) (*ListWorkspaces_Response, error)
	GetStatePayload(context.Context, *GetStatePayload_Request) (*GetStatePayload_Response, error)
	PutState(context.Context, *PutState_Request) (*Empty, error)
	DeleteState(context.Context, *DeleteState_Request) (*Empty, error)
	LockState(context.Context, *StateLock_Request) (*StateLock_Response, error)
	UnlockState(context.Context, *StateUnlock_Request) (*Empty, error)
	mustEmbedUnimplementedBackendServer()
}

// UnimplementedBackendServer must be embedded to have forward compatible implementations.
type UnimplementedBackendServer struct {
}

func (UnimplementedBackendServer) ConfigureBackend(context.Context, *ConfigureBackend_Request) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigureBackend not implemented")
}
func (UnimplementedBackendServer) DeleteWorkspace(context.Context, *DeleteWorkspace_Request) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWorkspace not implemented")
}
func (UnimplementedBackendServer) ListWorkspaces(context.Context, *ListWorkspaces_Request) (*ListWorkspaces_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorkspaces not implemented")
}
func (UnimplementedBackendServer) GetStatePayload(context.Context, *GetStatePayload_Request) (*GetStatePayload_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatePayload not implemented")
}
func (UnimplementedBackendServer) PutState(context.Context, *PutState_Request) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutState not implemented")
}
func (UnimplementedBackendServer) DeleteState(context.Context, *DeleteState_Request) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteState not implemented")
}
func (UnimplementedBackendServer) LockState(context.Context, *StateLock_Request) (*StateLock_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LockState not implemented")
}
func (UnimplementedBackendServer) UnlockState(context.Context, *StateUnlock_Request) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnlockState not implemented")
}
func (UnimplementedBackendServer) mustEmbedUnimplementedBackendServer() {}

// UnsafeBackendServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BackendServer will
// result in compilation errors.
type UnsafeBackendServer interface {
	mustEmbedUnimplementedBackendServer()
}

func RegisterBackendServer(s grpc.ServiceRegistrar, srv BackendServer) {
	s.RegisterService(&Backend_ServiceDesc, srv)
}

func _Backend_ConfigureBackend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureBackend_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).ConfigureBackend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Backend/ConfigureBackend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).ConfigureBackend(ctx, req.(*ConfigureBackend_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_DeleteWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWorkspace_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).DeleteWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Backend/DeleteWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).DeleteWorkspace(ctx, req.(*DeleteWorkspace_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_ListWorkspaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWorkspaces_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).ListWorkspaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Backend/ListWorkspaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).ListWorkspaces(ctx, req.(*ListWorkspaces_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_GetStatePayload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatePayload_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).GetStatePayload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Backend/GetStatePayload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).GetStatePayload(ctx, req.(*GetStatePayload_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_PutState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutState_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).PutState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Backend/PutState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).PutState(ctx, req.(*PutState_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_DeleteState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteState_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).DeleteState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Backend/DeleteState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).DeleteState(ctx, req.(*DeleteState_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_LockState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StateLock_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).LockState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Backend/LockState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).LockState(ctx, req.(*StateLock_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_UnlockState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StateUnlock_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).UnlockState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Backend/UnlockState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).UnlockState(ctx, req.(*StateUnlock_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Backend_ServiceDesc is the grpc.ServiceDesc for Backend service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Backend_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Backend",
	HandlerType: (*BackendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConfigureBackend",
			Handler:    _Backend_ConfigureBackend_Handler,
		},
		{
			MethodName: "DeleteWorkspace",
			Handler:    _Backend_DeleteWorkspace_Handler,
		},
		{
			MethodName: "ListWorkspaces",
			Handler:    _Backend_ListWorkspaces_Handler,
		},
		{
			MethodName: "GetStatePayload",
			Handler:    _Backend_GetStatePayload_Handler,
		},
		{
			MethodName: "PutState",
			Handler:    _Backend_PutState_Handler,
		},
		{
			MethodName: "DeleteState",
			Handler:    _Backend_DeleteState_Handler,
		},
		{
			MethodName: "LockState",
			Handler:    _Backend_LockState_Handler,
		},
		{
			MethodName: "UnlockState",
			Handler:    _Backend_UnlockState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "plugin.v1.proto",
}
