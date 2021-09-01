package main

import (
	"context"
	"flag"
	"net"
	"time"

	rrm "remote-communication/pkg/communication/adapters/rabbitmq"
	rapp "remote-communication/pkg/communication/app"
	rports "remote-communication/pkg/communication/ports"
	mapp "remote-communication/pkg/mission/app"
	mports "remote-communication/pkg/mission/ports"
	"remote-communication/pkg/skysign_proto"

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

	db, err := cpg.NewPostgresqlConnection("remote-communication")
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
	missionApp := mapp.NewApplication(ctx, txm, psm)

	revt := rports.NewEventHandler(communicationApp)
	mevt := mports.NewEventHandler(missionApp)

	rports.SubscribeEventHandler(ctx, psm, revt)
	mports.SubscribeEventHandler(ctx, psm, mevt)

	rrm.SubscribePublishHandler(psm)

	rsvc := rports.NewGrpcServer(communicationApp)
	msvc := mports.NewGrpcServer(missionApp)

	skysign_proto.RegisterCommunicationUserServiceServer(s, &rsvc)
	skysign_proto.RegisterCommunicationEdgeServiceServer(s, &rsvc)
	skysign_proto.RegisterUploadMissionEdgeServiceServer(s, &msvc)

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
