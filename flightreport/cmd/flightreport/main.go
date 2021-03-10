package main

import (
	"context"
	"flag"
	"flightreport/pkg/flightreport/adapters/postgresql"
	"flightreport/pkg/flightreport/adapters/rabbitmq"
	"flightreport/pkg/flightreport/app"
	"flightreport/pkg/flightreport/ports"
	"net"
	"time"

	proto "flightreport/pkg/skysign_proto"

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

	proto.RegisterReportFlightServiceServer(s, &svc)

	glog.Info("start flightreport server")
	return s.Serve(listen)
}

func main() {
	port = flag.String("port", "5001", "flightreport port")
	flag.Parse()
	defer glog.Flush()

	for {
		if err := run(); err != nil {
			glog.Error(err)
			time.Sleep(10 * time.Second)
		}
	}
}
