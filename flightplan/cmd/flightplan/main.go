package main

import (
	"context"
	"flag"
	"net"

	"flightplan/pkg/flightplan/api"
	"flightplan/pkg/flightplan/app"
	proto "flightplan/pkg/skysign_proto"

	"github.com/golang/glog"
	"google.golang.org/grpc"
)

var (
	port *string
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	listen, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		return err
	}
	s := grpc.NewServer()

	application := app.NewApplication(ctx)

	svc := api.NewGrpcServer(application)

	proto.RegisterManageFlightplanServiceServer(s, &svc)
	proto.RegisterAssignAssetsToFlightplanServiceServer(s, &svc)

	return s.Serve(listen)
}

func main() {
	port = flag.String("port", "5001", "flightplan port")
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
