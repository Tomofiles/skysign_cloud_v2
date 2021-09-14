package main

import (
	"context"
	"flag"
	"net"
	"os"

	frm "fleet-formation/pkg/fleet/adapters/rabbitmq"
	fapp "fleet-formation/pkg/fleet/app"
	fports "fleet-formation/pkg/fleet/ports"
	mrm "fleet-formation/pkg/mission/adapters/rabbitmq"
	mapp "fleet-formation/pkg/mission/app"
	mports "fleet-formation/pkg/mission/ports"
	vrm "fleet-formation/pkg/vehicle/adapters/rabbitmq"
	vapp "fleet-formation/pkg/vehicle/app"
	vports "fleet-formation/pkg/vehicle/ports"

	proto "github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

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

	fevt := fports.NewEventHandler(fleetApp)
	vevt := vports.NewEventHandler(vehilceApp)
	mevt := mports.NewEventHandler(missionApp)

	fports.SubscribeEventHandler(ctx, psm, fevt)
	vports.SubscribeEventHandler(ctx, psm, vevt)
	mports.SubscribeEventHandler(ctx, psm, mevt)

	frm.SubscribePublishHandler(psm)
	vrm.SubscribePublishHandler(psm)
	mrm.SubscribePublishHandler(psm)

	fsvc := fports.NewGrpcServer(fleetApp)
	vsvc := vports.NewGrpcServer(vehilceApp)
	msvc := mports.NewGrpcServer(missionApp)

	proto.RegisterAssignAssetsToFleetServiceServer(s, &fsvc)
	proto.RegisterManageVehicleServiceServer(s, &vsvc)
	proto.RegisterManageMissionServiceServer(s, &msvc)

	glog.Info("start fleet-formation server")
	return s.Serve(listen)
}

func main() {
	if err := run(); err != nil {
		glog.Error(err)
		os.Exit(1)
	}
}
