// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: proto/flightplan_execute.proto

/*
Package skysign_proto is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package skysign_proto

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_ExecuteFlightplanService_ExecuteFlightplan_0(ctx context.Context, marshaler runtime.Marshaler, client ExecuteFlightplanServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ExecuteFlightplanRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "id")
	}

	protoReq.Id, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", err)
	}

	msg, err := client.ExecuteFlightplan(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ExecuteFlightplanService_ExecuteFlightplan_0(ctx context.Context, marshaler runtime.Marshaler, server ExecuteFlightplanServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ExecuteFlightplanRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "id")
	}

	protoReq.Id, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", err)
	}

	msg, err := server.ExecuteFlightplan(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterExecuteFlightplanServiceHandlerServer registers the http handlers for service ExecuteFlightplanService to "mux".
// UnaryRPC     :call ExecuteFlightplanServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterExecuteFlightplanServiceHandlerFromEndpoint instead.
func RegisterExecuteFlightplanServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server ExecuteFlightplanServiceServer) error {

	mux.Handle("POST", pattern_ExecuteFlightplanService_ExecuteFlightplan_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/skysign_proto.ExecuteFlightplanService/ExecuteFlightplan", runtime.WithHTTPPathPattern("/api/v1/flightplans/{id}/execute"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_ExecuteFlightplanService_ExecuteFlightplan_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ExecuteFlightplanService_ExecuteFlightplan_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterExecuteFlightplanServiceHandlerFromEndpoint is same as RegisterExecuteFlightplanServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterExecuteFlightplanServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterExecuteFlightplanServiceHandler(ctx, mux, conn)
}

// RegisterExecuteFlightplanServiceHandler registers the http handlers for service ExecuteFlightplanService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterExecuteFlightplanServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterExecuteFlightplanServiceHandlerClient(ctx, mux, NewExecuteFlightplanServiceClient(conn))
}

// RegisterExecuteFlightplanServiceHandlerClient registers the http handlers for service ExecuteFlightplanService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "ExecuteFlightplanServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "ExecuteFlightplanServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "ExecuteFlightplanServiceClient" to call the correct interceptors.
func RegisterExecuteFlightplanServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client ExecuteFlightplanServiceClient) error {

	mux.Handle("POST", pattern_ExecuteFlightplanService_ExecuteFlightplan_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/skysign_proto.ExecuteFlightplanService/ExecuteFlightplan", runtime.WithHTTPPathPattern("/api/v1/flightplans/{id}/execute"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ExecuteFlightplanService_ExecuteFlightplan_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ExecuteFlightplanService_ExecuteFlightplan_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_ExecuteFlightplanService_ExecuteFlightplan_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3, 2, 4}, []string{"api", "v1", "flightplans", "id", "execute"}, ""))
)

var (
	forward_ExecuteFlightplanService_ExecuteFlightplan_0 = runtime.ForwardResponseMessage
)