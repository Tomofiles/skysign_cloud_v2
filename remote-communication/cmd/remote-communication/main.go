package main

import (
	"context"
	"flag"
	"net"
	"time"

	cpg "remote-communication/pkg/common/adapters/postgresql"
	crm "remote-communication/pkg/common/adapters/rabbitmq"
	cports "remote-communication/pkg/common/ports"
	rrm "remote-communication/pkg/communication/adapters/rabbitmq"
	rapp "remote-communication/pkg/communication/app"
	rports "remote-communication/pkg/communication/ports"
	"remote-communication/pkg/skysign_proto"

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
	s := grpc.NewServer(grpc.UnaryInterceptor(cports.LogBodyInterceptor()))

	db, err := cpg.NewPostgresqlConnection()
	if err != nil {
		return err
	}
	txm := cpg.NewGormTransactionManager(db)

	conn, err := crm.NewRabbitMQConnection()
	if err != nil {
		return err
	}
	defer conn.Close()
	psm := crm.NewPubSubManager(conn)

	communicationApp := rapp.NewApplication(ctx, txm, psm)

	revt := rports.NewEventHandler(communicationApp)

	rports.SubscribeEventHandler(ctx, psm, revt)

	rrm.SubscribePublishHandler(psm)

	rsvc := rports.NewGrpcServer(communicationApp)

	skysign_proto.RegisterCommunicationUserServiceServer(s, &rsvc)
	skysign_proto.RegisterCommunicationEdgeServiceServer(s, &rsvc)

	glog.Info("start remote-communication server")
	return s.Serve(listen)
}

func main() {
	port = flag.String("port", "5001", "remote-communication port")
	flag.Parse()
	defer glog.Flush()

	for {
		if err := run(); err != nil {
			glog.Error(err)
			time.Sleep(10 * time.Second)
		}
	}
}
