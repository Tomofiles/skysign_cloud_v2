package main

import (
	"context"
	"flag"
	"net"
	"time"

	"flightplan/pkg/flightplan/adapters/postgresql"
	"flightplan/pkg/flightplan/adapters/rabbitmq"
	"flightplan/pkg/flightplan/app"
	"flightplan/pkg/flightplan/ports"
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
	defer listen.Close()
	s := grpc.NewServer(grpc.UnaryInterceptor(ports.LogBodyInterceptor()))

	db, err := postgresql.NewPostgresqlConnection()
	if err != nil {
		return err
	}
	txm := postgresql.NewGormTransactionManager(db)

	conn, err := rabbitmq.NewRabbitMQConnection()
	if err != nil {
		return err
	}
	defer conn.Close()
	psm := rabbitmq.NewPubSubManager(conn)

	application := app.NewApplication(ctx, txm, psm)

	svc := ports.NewGrpcServer(application)

	proto.RegisterManageFlightplanServiceServer(s, &svc)
	proto.RegisterChangeFlightplanServiceServer(s, &svc)
	proto.RegisterExecuteFlightplanServiceServer(s, &svc)

	glog.Info("start flightplan server")
	return s.Serve(listen)
}

func main() {
	port = flag.String("port", "5001", "flightplan port")
	flag.Parse()
	defer glog.Flush()

	for {
		if err := run(); err != nil {
			glog.Error(err)
			time.Sleep(10 * time.Second)
		}
	}
}
