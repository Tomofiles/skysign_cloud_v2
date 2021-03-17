package main

import (
	"action/pkg/action/adapters/postgresql"
	"action/pkg/action/adapters/rabbitmq"
	"action/pkg/action/app"
	"action/pkg/action/ports"
	proto "action/pkg/skysign_proto"
	"context"
	"flag"
	"net"
	"time"

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

	application := app.NewApplication(ctx, txm)

	svc := ports.NewGrpcServer(application)
	evt := ports.NewEventHandler(application)

	psm.SetConsumer(
		ctx,
		ports.CopiedVehicleCreatedEventExchangeName,
		ports.CopiedVehicleCreatedEventQueueName,
		func(event []byte) {
			if err := evt.HandleCopiedVehicleCreatedEvent(ctx, event); err != nil {
				glog.Error(err)
			}
		},
	)
	psm.SetConsumer(
		ctx,
		ports.FlightoperationCompletedEventExchangeName,
		ports.FlightoperationCompletedEventQueueName,
		func(event []byte) {
			if err := evt.HandleFlightoperationCompletedEvent(ctx, event); err != nil {
				glog.Error(err)
			}
		},
	)
	psm.SetConsumer(
		ctx,
		ports.TelemetryUpdatedEventExchangeName,
		ports.TelemetryUpdatedEventQueueName,
		func(event []byte) {
			if err := evt.HandleTelemetryUpdatedEvent(ctx, event); err != nil {
				glog.Error(err)
			}
		},
	)

	proto.RegisterActionServiceServer(s, &svc)

	glog.Info("start action server")
	return s.Serve(listen)
}

func main() {
	port = flag.String("port", "5001", "action port")
	flag.Parse()
	defer glog.Flush()

	for {
		if err := run(); err != nil {
			glog.Error(err)
			time.Sleep(10 * time.Second)
		}
	}
}
