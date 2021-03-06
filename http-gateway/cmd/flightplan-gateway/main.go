package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	gw "github.com/Tomofiles/skysign_cloud/http-gateway/pkg/skysign_proto"
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
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}),
	}

	mux := runtime.NewServeMux(smOpts...)
	opts := []grpc.DialOption{grpc.WithInsecure()}
	endpoint := fmt.Sprintf(*backendHost + ":" + *backendPort)
	err := gw.RegisterManageFlightplanServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		return err
	}
	err = gw.RegisterAssignAssetsToFlightplanServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
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
