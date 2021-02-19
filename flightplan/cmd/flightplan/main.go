package main

import (
	"context"
	"flag"
	"net"

	"flightplan/pkg/flightplan/api"
	"flightplan/pkg/flightplan/app"
	"flightplan/pkg/flightplan/domain/bridge"
	"flightplan/pkg/flightplan/infra/postgresql"
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

	db, err := postgresql.NewPostgresqlConnection()
	if err != nil {
		panic(err)
	}
	txm := postgresql.NewGormTransactionManager(db)

	application := app.NewApplication(ctx, txm)

	svc := api.NewGrpcServer(application)
	evt := api.NewEventHandler(application)

	bridge.Bind(evt, application)

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
