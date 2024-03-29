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

// ReportFlightServiceClient is the client API for ReportFlightService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReportFlightServiceClient interface {
	ListFlightreports(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListFlightreportsResponses, error)
	GetFlightreport(ctx context.Context, in *GetFlightreportRequest, opts ...grpc.CallOption) (*Flightreport, error)
}

type reportFlightServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReportFlightServiceClient(cc grpc.ClientConnInterface) ReportFlightServiceClient {
	return &reportFlightServiceClient{cc}
}

func (c *reportFlightServiceClient) ListFlightreports(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListFlightreportsResponses, error) {
	out := new(ListFlightreportsResponses)
	err := c.cc.Invoke(ctx, "/skysign_proto.ReportFlightService/ListFlightreports", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportFlightServiceClient) GetFlightreport(ctx context.Context, in *GetFlightreportRequest, opts ...grpc.CallOption) (*Flightreport, error) {
	out := new(Flightreport)
	err := c.cc.Invoke(ctx, "/skysign_proto.ReportFlightService/GetFlightreport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReportFlightServiceServer is the server API for ReportFlightService service.
// All implementations must embed UnimplementedReportFlightServiceServer
// for forward compatibility
type ReportFlightServiceServer interface {
	ListFlightreports(context.Context, *Empty) (*ListFlightreportsResponses, error)
	GetFlightreport(context.Context, *GetFlightreportRequest) (*Flightreport, error)
	mustEmbedUnimplementedReportFlightServiceServer()
}

// UnimplementedReportFlightServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReportFlightServiceServer struct {
}

func (UnimplementedReportFlightServiceServer) ListFlightreports(context.Context, *Empty) (*ListFlightreportsResponses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFlightreports not implemented")
}
func (UnimplementedReportFlightServiceServer) GetFlightreport(context.Context, *GetFlightreportRequest) (*Flightreport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFlightreport not implemented")
}
func (UnimplementedReportFlightServiceServer) mustEmbedUnimplementedReportFlightServiceServer() {}

// UnsafeReportFlightServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReportFlightServiceServer will
// result in compilation errors.
type UnsafeReportFlightServiceServer interface {
	mustEmbedUnimplementedReportFlightServiceServer()
}

func RegisterReportFlightServiceServer(s grpc.ServiceRegistrar, srv ReportFlightServiceServer) {
	s.RegisterService(&ReportFlightService_ServiceDesc, srv)
}

func _ReportFlightService_ListFlightreports_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportFlightServiceServer).ListFlightreports(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/skysign_proto.ReportFlightService/ListFlightreports",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportFlightServiceServer).ListFlightreports(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportFlightService_GetFlightreport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFlightreportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportFlightServiceServer).GetFlightreport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/skysign_proto.ReportFlightService/GetFlightreport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportFlightServiceServer).GetFlightreport(ctx, req.(*GetFlightreportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReportFlightService_ServiceDesc is the grpc.ServiceDesc for ReportFlightService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReportFlightService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "skysign_proto.ReportFlightService",
	HandlerType: (*ReportFlightServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListFlightreports",
			Handler:    _ReportFlightService_ListFlightreports_Handler,
		},
		{
			MethodName: "GetFlightreport",
			Handler:    _ReportFlightService_GetFlightreport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/flightreport.proto",
}
