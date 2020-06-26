package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	gw "github.com/Tomofiles/skysign_cloud/http_gateway/pkg/skysign_proto"
)

var (
	backendHost *string
	backendPort *string
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	endpoint := fmt.Sprintf(*backendHost + ":" + *backendPort)
	err := gw.RegisterVehicleServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		return err
	}
	err = gw.RegisterCommunicationUserServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		return err
	}
	err = gw.RegisterCommunicationVehicleServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":5000", mux)
}

func main() {
	backendHost = flag.String("backend_host", "localhost", "backend host")
	backendPort = flag.String("backend_port", "5001", "backend port")
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
