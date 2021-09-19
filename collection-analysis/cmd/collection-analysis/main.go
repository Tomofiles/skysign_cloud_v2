package main

import (
	"context"
	"flag"
	"net"
	"os"

	agrpc "github.com/Tomofiles/skysign_cloud_v2/collection-analysis/pkg/action/adapters/grpc"
	arm "github.com/Tomofiles/skysign_cloud_v2/collection-analysis/pkg/action/adapters/rabbitmq"
	aapp "github.com/Tomofiles/skysign_cloud_v2/collection-analysis/pkg/action/app"

	cpg "github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/adapters/postgresql"
	crm "github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/adapters/rabbitmq"
	cports "github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/ports"

	"github.com/golang/glog"
	"google.golang.org/grpc"
)

var (
	port *string
)

func run() error {
	port = flag.String("port", "5001", "collection-analysis port")
	flag.Parse()
	defer glog.Flush()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	listen, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		return err
	}
	defer listen.Close()
	s := grpc.NewServer(grpc.UnaryInterceptor(cports.LogBodyInterceptor()))

	db, err := cpg.NewPostgresqlConnection("collection-analysis")
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

	arm.SubscribeEventHandler(ctx, psm, aApp)

	agrpc.SubscribeGrpcServer(s, aApp)

	glog.Info("start collection-analysis server")
	return s.Serve(listen)
}

func main() {
	if err := run(); err != nil {
		glog.Error(err)
		os.Exit(1)
	}
}
