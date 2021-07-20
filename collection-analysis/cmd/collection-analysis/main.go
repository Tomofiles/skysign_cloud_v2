package main

import (
	"context"
	"flag"
	"net"
	"time"

	arm "collection-analysis/pkg/action/adapters/rabbitmq"
	aapp "collection-analysis/pkg/action/app"
	aports "collection-analysis/pkg/action/ports"
	cpg "collection-analysis/pkg/common/adapters/postgresql"
	crm "collection-analysis/pkg/common/adapters/rabbitmq"
	cports "collection-analysis/pkg/common/ports"
	proto "collection-analysis/pkg/skysign_proto"

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

	aApp := aapp.NewApplication(ctx, txm)

	aevt := aports.NewEventHandler(aApp)

	aports.SubscribeEventHandler(ctx, psm, aevt)

	arm.SubscribePublishHandler(psm)

	asvc := aports.NewGrpcServer(aApp)

	proto.RegisterActionServiceServer(s, &asvc)

	glog.Info("start collection-analysis server")
	return s.Serve(listen)
}

func main() {
	port = flag.String("port", "5001", "collection-analysis port")
	flag.Parse()
	defer glog.Flush()

	for {
		if err := run(); err != nil {
			glog.Error(err)
			time.Sleep(10 * time.Second)
		}
	}
}
