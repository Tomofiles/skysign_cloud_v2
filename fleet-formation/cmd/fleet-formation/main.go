package main

import (
	"context"
	"flag"
	"net"
	"os"

	fgrpc "github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/fleet/adapters/grpc"
	frm "github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/fleet/adapters/rabbitmq"
	fapp "github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/fleet/app"
	mgrpc "github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/mission/adapters/grpc"
	mrm "github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/mission/adapters/rabbitmq"
	mapp "github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/mission/app"
	vgrpc "github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/vehicle/adapters/grpc"
	vrm "github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/vehicle/adapters/rabbitmq"
	vapp "github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/vehicle/app"

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
	port = flag.String("port", "5001", "fleet-formation port")
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

	db, err := cpg.NewPostgresqlConnection("fleet-formation")
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

	fleetApp := fapp.NewApplication(ctx, txm, psm)
	vehilceApp := vapp.NewApplication(ctx, txm, psm)
	missionApp := mapp.NewApplication(ctx, txm, psm)

	frm.SubscribeEventHandler(ctx, psm, fleetApp)
	vrm.SubscribeEventHandler(ctx, psm, vehilceApp)
	mrm.SubscribeEventHandler(ctx, psm, missionApp)

	frm.SubscribeEventPublisher(psm)
	vrm.SubscribeEventPublisher(psm)
	mrm.SubscribeEventPublisher(psm)

	fgrpc.SubscribeGrpcServer(s, fleetApp)
	vgrpc.SubscribeGrpcServer(s, vehilceApp)
	mgrpc.SubscribeGrpcServer(s, missionApp)

	glog.Info("start fleet-formation server")
	return s.Serve(listen)
}

func main() {
	if err := run(); err != nil {
		glog.Error(err)
		os.Exit(1)
	}
}
