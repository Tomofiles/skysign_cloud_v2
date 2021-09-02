package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"

	gw "github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
)

var (
	backendHost *string
	backendPort *string
	port        *string
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	smOpts := []runtime.ServeMuxOption{
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames:   true,
				EmitUnpopulated: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		}),
	}

	mux := runtime.NewServeMux(smOpts...)
	opts := []grpc.DialOption{grpc.WithInsecure()}
	endpoint := fmt.Sprintf(*backendHost + ":" + *backendPort)
	err := gw.RegisterManageFlightplanServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		return err
	}
	err = gw.RegisterChangeFlightplanServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		return err
	}
	err = gw.RegisterExecuteFlightplanServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		return err
	}
	err = gw.RegisterOperateFlightServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		return err
	}
	err = gw.RegisterReportFlightServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":"+*port, mux)
}

func main() {
	backendHost = flag.String("backend_host", "localhost", "backend host")
	backendPort = flag.String("backend_port", "5001", "backend port")
	port = flag.String("port", "5000", "http gateway port")
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
