package main

import (
	"context"
	"flag"
	"mission/pkg/mission/adapters/postgresql"
	"mission/pkg/mission/adapters/rabbitmq"
	"mission/pkg/mission/app"
	"mission/pkg/mission/ports"
	proto "mission/pkg/skysign_proto"
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

	application := app.NewApplication(ctx, txm, psm)

	svc := ports.NewGrpcServer(application)
	evt := ports.NewEventHandler(application)

	psm.SetConsumer(
		ctx,
		ports.MissionCopiedWhenFlightplanCopiedEventExchangeName,
		ports.MissionCopiedWhenFlightplanCopiedEventQueueName,
		func(event []byte) {
			if err := evt.HandleMissionCopiedWhenFlightplanCopiedEvent(ctx, event); err != nil {
				glog.Error(err)
			}
		},
	)

	proto.RegisterManageMissionServiceServer(s, &svc)

	glog.Info("start mission server")
	return s.Serve(listen)
}

func main() {
	port = flag.String("port", "5001", "mission port")
	flag.Parse()
	defer glog.Flush()

	for {
		if err := run(); err != nil {
			glog.Error(err)
			time.Sleep(10 * time.Second)
		}
	}
}
