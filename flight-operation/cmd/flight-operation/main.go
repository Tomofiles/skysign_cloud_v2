package main

import (
	"context"
	"flag"
	"net"
	"os"

	fopegrpc "github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightoperation/adapters/grpc"
	foperm "github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightoperation/adapters/rabbitmq"
	fopeapp "github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightoperation/app"
	fplgrpc "github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightplan/adapters/grpc"
	fplrm "github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightplan/adapters/rabbitmq"
	fplapp "github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightplan/app"
	frepgrpc "github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightreport/adapters/grpc"
	freprm "github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightreport/adapters/rabbitmq"
	frepapp "github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightreport/app"

	cgrpc "github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/adapters/grpc"
	cpg "github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/adapters/postgresql"
	crm "github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/adapters/rabbitmq"

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
	s := grpc.NewServer(grpc.UnaryInterceptor(cgrpc.LogBodyInterceptor()))

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
