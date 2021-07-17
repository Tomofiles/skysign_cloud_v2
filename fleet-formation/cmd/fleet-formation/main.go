package main

import (
	"context"
	"flag"
	"net"
	"time"

	"fleet-formation/pkg/fleet/adapters/postgresql"
	"fleet-formation/pkg/fleet/adapters/rabbitmq"
	"fleet-formation/pkg/fleet/app"
	"fleet-formation/pkg/fleet/ports"
	proto "fleet-formation/pkg/skysign_proto"

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
	evt := ports.NewEventHandler(application)

	psm.SetConsumer(
		ctx,
		ports.FleetIDGaveEventExchangeName,
		ports.FleetIDGaveEventQueueName,
		func(event []byte) {
			if err := evt.HandleFleetIDGaveEvent(ctx, event); err != nil {
				glog.Error(err)
			}
		},
	)
	psm.SetConsumer(
		ctx,
		ports.FleetIDRemovedEventExchangeName,
		ports.FleetIDRemovedEventQueueName,
		func(event []byte) {
			if err := evt.HandleFleetIDRemovedEvent(ctx, event); err != nil {
				glog.Error(err)
			}
		},
	)
	psm.SetConsumer(
		ctx,
		ports.FleetCopiedEventExchangeName,
		ports.FleetCopiedEventQueueName,
		func(event []byte) {
			if err := evt.HandleFleetCopiedEvent(ctx, event); err != nil {
				glog.Error(err)
			}
		},
	)

	proto.RegisterAssignAssetsToFleetServiceServer(s, &svc)

	glog.Info("start fleet-formation server")
	return s.Serve(listen)
}

func main() {
	port = flag.String("port", "5001", "fleet-formation port")
	flag.Parse()
	defer glog.Flush()

	for {
		if err := run(); err != nil {
			glog.Error(err)
			time.Sleep(10 * time.Second)
		}
	}
}
