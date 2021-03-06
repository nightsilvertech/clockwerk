// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ClockwerkClient is the client API for Clockwerk service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClockwerkClient interface {
	GetDummy(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	PostDummy(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteDummy(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	PutDummy(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetSchedulers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Schedulers, error)
	AddScheduler(ctx context.Context, in *Scheduler, opts ...grpc.CallOption) (*Scheduler, error)
	DeleteScheduler(ctx context.Context, in *SelectScheduler, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ToggleScheduler(ctx context.Context, in *SelectToggle, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Backup(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type clockwerkClient struct {
	cc grpc.ClientConnInterface
}

func NewClockwerkClient(cc grpc.ClientConnInterface) ClockwerkClient {
	return &clockwerkClient{cc}
}

func (c *clockwerkClient) GetDummy(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.v1.Clockwerk/GetDummy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clockwerkClient) PostDummy(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.v1.Clockwerk/PostDummy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clockwerkClient) DeleteDummy(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.v1.Clockwerk/DeleteDummy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clockwerkClient) PutDummy(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.v1.Clockwerk/PutDummy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clockwerkClient) GetSchedulers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Schedulers, error) {
	out := new(Schedulers)
	err := c.cc.Invoke(ctx, "/api.v1.Clockwerk/GetSchedulers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clockwerkClient) AddScheduler(ctx context.Context, in *Scheduler, opts ...grpc.CallOption) (*Scheduler, error) {
	out := new(Scheduler)
	err := c.cc.Invoke(ctx, "/api.v1.Clockwerk/AddScheduler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clockwerkClient) DeleteScheduler(ctx context.Context, in *SelectScheduler, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.v1.Clockwerk/DeleteScheduler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clockwerkClient) ToggleScheduler(ctx context.Context, in *SelectToggle, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.v1.Clockwerk/ToggleScheduler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clockwerkClient) Backup(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.v1.Clockwerk/Backup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClockwerkServer is the server API for Clockwerk service.
// All implementations should embed UnimplementedClockwerkServer
// for forward compatibility
type ClockwerkServer interface {
	GetDummy(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	PostDummy(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	DeleteDummy(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	PutDummy(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	GetSchedulers(context.Context, *emptypb.Empty) (*Schedulers, error)
	AddScheduler(context.Context, *Scheduler) (*Scheduler, error)
	DeleteScheduler(context.Context, *SelectScheduler) (*emptypb.Empty, error)
	ToggleScheduler(context.Context, *SelectToggle) (*emptypb.Empty, error)
	Backup(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
}

// UnimplementedClockwerkServer should be embedded to have forward compatible implementations.
type UnimplementedClockwerkServer struct {
}

func (UnimplementedClockwerkServer) GetDummy(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDummy not implemented")
}
func (UnimplementedClockwerkServer) PostDummy(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostDummy not implemented")
}
func (UnimplementedClockwerkServer) DeleteDummy(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDummy not implemented")
}
func (UnimplementedClockwerkServer) PutDummy(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutDummy not implemented")
}
func (UnimplementedClockwerkServer) GetSchedulers(context.Context, *emptypb.Empty) (*Schedulers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSchedulers not implemented")
}
func (UnimplementedClockwerkServer) AddScheduler(context.Context, *Scheduler) (*Scheduler, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddScheduler not implemented")
}
func (UnimplementedClockwerkServer) DeleteScheduler(context.Context, *SelectScheduler) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteScheduler not implemented")
}
func (UnimplementedClockwerkServer) ToggleScheduler(context.Context, *SelectToggle) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToggleScheduler not implemented")
}
func (UnimplementedClockwerkServer) Backup(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Backup not implemented")
}

// UnsafeClockwerkServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClockwerkServer will
// result in compilation errors.
type UnsafeClockwerkServer interface {
	mustEmbedUnimplementedClockwerkServer()
}

func RegisterClockwerkServer(s grpc.ServiceRegistrar, srv ClockwerkServer) {
	s.RegisterService(&Clockwerk_ServiceDesc, srv)
}

func _Clockwerk_GetDummy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClockwerkServer).GetDummy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.Clockwerk/GetDummy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClockwerkServer).GetDummy(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clockwerk_PostDummy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClockwerkServer).PostDummy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.Clockwerk/PostDummy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClockwerkServer).PostDummy(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clockwerk_DeleteDummy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClockwerkServer).DeleteDummy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.Clockwerk/DeleteDummy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClockwerkServer).DeleteDummy(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clockwerk_PutDummy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClockwerkServer).PutDummy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.Clockwerk/PutDummy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClockwerkServer).PutDummy(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clockwerk_GetSchedulers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClockwerkServer).GetSchedulers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.Clockwerk/GetSchedulers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClockwerkServer).GetSchedulers(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clockwerk_AddScheduler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Scheduler)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClockwerkServer).AddScheduler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.Clockwerk/AddScheduler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClockwerkServer).AddScheduler(ctx, req.(*Scheduler))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clockwerk_DeleteScheduler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectScheduler)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClockwerkServer).DeleteScheduler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.Clockwerk/DeleteScheduler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClockwerkServer).DeleteScheduler(ctx, req.(*SelectScheduler))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clockwerk_ToggleScheduler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectToggle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClockwerkServer).ToggleScheduler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.Clockwerk/ToggleScheduler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClockwerkServer).ToggleScheduler(ctx, req.(*SelectToggle))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clockwerk_Backup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClockwerkServer).Backup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.Clockwerk/Backup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClockwerkServer).Backup(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Clockwerk_ServiceDesc is the grpc.ServiceDesc for Clockwerk service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Clockwerk_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.Clockwerk",
	HandlerType: (*ClockwerkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDummy",
			Handler:    _Clockwerk_GetDummy_Handler,
		},
		{
			MethodName: "PostDummy",
			Handler:    _Clockwerk_PostDummy_Handler,
		},
		{
			MethodName: "DeleteDummy",
			Handler:    _Clockwerk_DeleteDummy_Handler,
		},
		{
			MethodName: "PutDummy",
			Handler:    _Clockwerk_PutDummy_Handler,
		},
		{
			MethodName: "GetSchedulers",
			Handler:    _Clockwerk_GetSchedulers_Handler,
		},
		{
			MethodName: "AddScheduler",
			Handler:    _Clockwerk_AddScheduler_Handler,
		},
		{
			MethodName: "DeleteScheduler",
			Handler:    _Clockwerk_DeleteScheduler_Handler,
		},
		{
			MethodName: "ToggleScheduler",
			Handler:    _Clockwerk_ToggleScheduler_Handler,
		},
		{
			MethodName: "Backup",
			Handler:    _Clockwerk_Backup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "clockwerk.proto",
}
