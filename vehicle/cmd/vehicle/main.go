package main

import (
	"context"
	"flag"
	"net"
	"time"
	proto "vehicle/pkg/skysign_proto"
	"vehicle/pkg/vehicle/adapters/postgresql"
	"vehicle/pkg/vehicle/adapters/rabbitmq"
	"vehicle/pkg/vehicle/app"
	"vehicle/pkg/vehicle/ports"

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
		ports.VehicleCopiedWhenFlightplanCopiedEventExchangeName,
		ports.VehicleCopiedWhenFlightplanCopiedEventQueueName,
		func(event []byte) {
			if err := evt.HandleVehicleCopiedWhenFlightplanCopiedEvent(ctx, event); err != nil {
				glog.Error(err)
			}
		},
	)

	proto.RegisterManageVehicleServiceServer(s, &svc)

	glog.Info("start vehicle server")
	return s.Serve(listen)
}

func main() {
	port = flag.String("port", "5001", "vehicle port")
	flag.Parse()
	defer glog.Flush()

	for {
		if err := run(); err != nil {
			glog.Error(err)
			time.Sleep(10 * time.Second)
		}
	}
}
