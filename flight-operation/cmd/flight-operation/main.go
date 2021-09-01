package main

import (
	"context"
	"flag"
	"net"
	"time"

	foperm "flight-operation/pkg/flightoperation/adapters/rabbitmq"
	fopeapp "flight-operation/pkg/flightoperation/app"
	fopeports "flight-operation/pkg/flightoperation/ports"
	fplrm "flight-operation/pkg/flightplan/adapters/rabbitmq"
	fplapp "flight-operation/pkg/flightplan/app"
	fplports "flight-operation/pkg/flightplan/ports"
	frepapp "flight-operation/pkg/flightreport/app"
	frepports "flight-operation/pkg/flightreport/ports"
	proto "flight-operation/pkg/skysign_proto"

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

	fopeevt := fopeports.NewEventHandler(fopeApp)
	frepevt := frepports.NewEventHandler(frepApp)

	fopeports.SubscribeEventHandler(ctx, psm, fopeevt)
	frepports.SubscribeEventHandler(ctx, psm, frepevt)

	fplrm.SubscribePublishHandler(psm)
	foperm.SubscribePublishHandler(psm)

	fplsvc := fplports.NewGrpcServer(fplApp)
	fopesvc := fopeports.NewGrpcServer(fopeApp)
	frepsvc := frepports.NewGrpcServer(frepApp)

	proto.RegisterManageFlightplanServiceServer(s, &fplsvc)
	proto.RegisterChangeFlightplanServiceServer(s, &fplsvc)
	proto.RegisterExecuteFlightplanServiceServer(s, &fplsvc)
	proto.RegisterOperateFlightServiceServer(s, &fopesvc)
	proto.RegisterReportFlightServiceServer(s, &frepsvc)

	glog.Info("start flight-operation server")
	return s.Serve(listen)
}

func main() {
	port = flag.String("port", "5001", "flight-operation port")
	flag.Parse()
	defer glog.Flush()

	for {
		if err := run(); err != nil {
			glog.Error(err)
			time.Sleep(10 * time.Second)
		}
	}
}
