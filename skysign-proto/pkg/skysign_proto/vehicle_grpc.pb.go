// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package skysign_proto

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

// ManageVehicleServiceClient is the client API for ManageVehicleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManageVehicleServiceClient interface {
	ListVehicles(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListVehiclesResponses, error)
	GetVehicle(ctx context.Context, in *GetVehicleRequest, opts ...grpc.CallOption) (*Vehicle, error)
	CreateVehicle(ctx context.Context, in *Vehicle, opts ...grpc.CallOption) (*Vehicle, error)
	UpdateVehicle(ctx context.Context, in *Vehicle, opts ...grpc.CallOption) (*Vehicle, error)
	DeleteVehicle(ctx context.Context, in *DeleteVehicleRequest, opts ...grpc.CallOption) (*Empty, error)
}

type manageVehicleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewManageVehicleServiceClient(cc grpc.ClientConnInterface) ManageVehicleServiceClient {
	return &manageVehicleServiceClient{cc}
}

func (c *manageVehicleServiceClient) ListVehicles(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListVehiclesResponses, error) {
	out := new(ListVehiclesResponses)
	err := c.cc.Invoke(ctx, "/skysign_proto.ManageVehicleService/ListVehicles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manageVehicleServiceClient) GetVehicle(ctx context.Context, in *GetVehicleRequest, opts ...grpc.CallOption) (*Vehicle, error) {
	out := new(Vehicle)
	err := c.cc.Invoke(ctx, "/skysign_proto.ManageVehicleService/GetVehicle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manageVehicleServiceClient) CreateVehicle(ctx context.Context, in *Vehicle, opts ...grpc.CallOption) (*Vehicle, error) {
	out := new(Vehicle)
	err := c.cc.Invoke(ctx, "/skysign_proto.ManageVehicleService/CreateVehicle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manageVehicleServiceClient) UpdateVehicle(ctx context.Context, in *Vehicle, opts ...grpc.CallOption) (*Vehicle, error) {
	out := new(Vehicle)
	err := c.cc.Invoke(ctx, "/skysign_proto.ManageVehicleService/UpdateVehicle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manageVehicleServiceClient) DeleteVehicle(ctx context.Context, in *DeleteVehicleRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/skysign_proto.ManageVehicleService/DeleteVehicle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManageVehicleServiceServer is the server API for ManageVehicleService service.
// All implementations must embed UnimplementedManageVehicleServiceServer
// for forward compatibility
type ManageVehicleServiceServer interface {
	ListVehicles(context.Context, *Empty) (*ListVehiclesResponses, error)
	GetVehicle(context.Context, *GetVehicleRequest) (*Vehicle, error)
	CreateVehicle(context.Context, *Vehicle) (*Vehicle, error)
	UpdateVehicle(context.Context, *Vehicle) (*Vehicle, error)
	DeleteVehicle(context.Context, *DeleteVehicleRequest) (*Empty, error)
	mustEmbedUnimplementedManageVehicleServiceServer()
}

// UnimplementedManageVehicleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedManageVehicleServiceServer struct {
}

func (UnimplementedManageVehicleServiceServer) ListVehicles(context.Context, *Empty) (*ListVehiclesResponses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVehicles not implemented")
}
func (UnimplementedManageVehicleServiceServer) GetVehicle(context.Context, *GetVehicleRequest) (*Vehicle, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVehicle not implemented")
}
func (UnimplementedManageVehicleServiceServer) CreateVehicle(context.Context, *Vehicle) (*Vehicle, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVehicle not implemented")
}
func (UnimplementedManageVehicleServiceServer) UpdateVehicle(context.Context, *Vehicle) (*Vehicle, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateVehicle not implemented")
}
func (UnimplementedManageVehicleServiceServer) DeleteVehicle(context.Context, *DeleteVehicleRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVehicle not implemented")
}
func (UnimplementedManageVehicleServiceServer) mustEmbedUnimplementedManageVehicleServiceServer() {}

// UnsafeManageVehicleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManageVehicleServiceServer will
// result in compilation errors.
type UnsafeManageVehicleServiceServer interface {
	mustEmbedUnimplementedManageVehicleServiceServer()
}

func RegisterManageVehicleServiceServer(s grpc.ServiceRegistrar, srv ManageVehicleServiceServer) {
	s.RegisterService(&ManageVehicleService_ServiceDesc, srv)
}

func _ManageVehicleService_ListVehicles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageVehicleServiceServer).ListVehicles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/skysign_proto.ManageVehicleService/ListVehicles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageVehicleServiceServer).ListVehicles(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManageVehicleService_GetVehicle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVehicleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageVehicleServiceServer).GetVehicle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/skysign_proto.ManageVehicleService/GetVehicle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageVehicleServiceServer).GetVehicle(ctx, req.(*GetVehicleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManageVehicleService_CreateVehicle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Vehicle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageVehicleServiceServer).CreateVehicle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/skysign_proto.ManageVehicleService/CreateVehicle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageVehicleServiceServer).CreateVehicle(ctx, req.(*Vehicle))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManageVehicleService_UpdateVehicle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Vehicle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageVehicleServiceServer).UpdateVehicle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/skysign_proto.ManageVehicleService/UpdateVehicle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageVehicleServiceServer).UpdateVehicle(ctx, req.(*Vehicle))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManageVehicleService_DeleteVehicle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteVehicleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageVehicleServiceServer).DeleteVehicle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/skysign_proto.ManageVehicleService/DeleteVehicle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageVehicleServiceServer).DeleteVehicle(ctx, req.(*DeleteVehicleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ManageVehicleService_ServiceDesc is the grpc.ServiceDesc for ManageVehicleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ManageVehicleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "skysign_proto.ManageVehicleService",
	HandlerType: (*ManageVehicleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListVehicles",
			Handler:    _ManageVehicleService_ListVehicles_Handler,
		},
		{
			MethodName: "GetVehicle",
			Handler:    _ManageVehicleService_GetVehicle_Handler,
		},
		{
			MethodName: "CreateVehicle",
			Handler:    _ManageVehicleService_CreateVehicle_Handler,
		},
		{
			MethodName: "UpdateVehicle",
			Handler:    _ManageVehicleService_UpdateVehicle_Handler,
		},
		{
			MethodName: "DeleteVehicle",
			Handler:    _ManageVehicleService_DeleteVehicle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/vehicle.proto",
}