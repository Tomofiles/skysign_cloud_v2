package main

import (
	"context"
	"flag"
	"net"
	"os"

	fopegrpc "flight-operation/pkg/flightoperation/adapters/grpc"
	foperm "flight-operation/pkg/flightoperation/adapters/rabbitmq"
	fopeapp "flight-operation/pkg/flightoperation/app"
	fplgrpc "flight-operation/pkg/flightplan/adapters/grpc"
	fplrm "flight-operation/pkg/flightplan/adapters/rabbitmq"
	fplapp "flight-operation/pkg/flightplan/app"
	frepgrpc "flight-operation/pkg/flightreport/adapters/grpc"
	freprm "flight-operation/pkg/flightreport/adapters/rabbitmq"
	frepapp "flight-operation/pkg/flightreport/app"

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
	port = flag.String("port", "5001", "flight-operation port")
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

	db, err := cpg.NewPostgresqlConnection("flight-operation")
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

	fplApp := fplapp.NewApplication(ctx, txm, psm)
	fopeApp := fopeapp.NewApplication(ctx, txm, psm)
	frepApp := frepapp.NewApplication(ctx, txm, psm)

	foperm.SubscribeEventHandler(ctx, psm, fopeApp)
	freprm.SubscribeEventHandler(ctx, psm, frepApp)

	fplrm.SubscribeEventPublisher(psm)
	foperm.SubscribeEventPublisher(psm)

	fplgrpc.SubscribeGrpcServer(s, fplApp)
	fopegrpc.SubscribeGrpcServer(s, fopeApp)
	frepgrpc.SubscribeGrpcServer(s, frepApp)

	glog.Info("start flight-operation server")
	return s.Serve(listen)
}

func main() {
	if err := run(); err != nil {
		glog.Error(err)
		os.Exit(1)
	}
}
