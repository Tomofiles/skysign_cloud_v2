package main

import (
	"context"
	"flag"
	"net"
	"os"

	rgrpc "github.com/Tomofiles/skysign_cloud_v2/remote-communication/pkg/communication/adapters/grpc"
	rrm "github.com/Tomofiles/skysign_cloud_v2/remote-communication/pkg/communication/adapters/rabbitmq"
	rapp "github.com/Tomofiles/skysign_cloud_v2/remote-communication/pkg/communication/app"
	mgrpc "github.com/Tomofiles/skysign_cloud_v2/remote-communication/pkg/mission/adapters/grpc"
	mrm "github.com/Tomofiles/skysign_cloud_v2/remote-communication/pkg/mission/adapters/rabbitmq"
	mapp "github.com/Tomofiles/skysign_cloud_v2/remote-communication/pkg/mission/app"

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
	port = flag.String("port", "5001", "remote-communication port")
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

	rrm.SubscribeEventHandler(ctx, psm, communicationApp)
	mrm.SubscribeEventHandler(ctx, psm, missionApp)

	rrm.SubscribeEventPublisher(psm)

	rgrpc.SubscribeGrpcServer(s, communicationApp)
	mgrpc.SubscribeGrpcServer(s, missionApp)

	glog.Info("start remote-communication server")
	return s.Serve(listen)
}

func main() {
	if err := run(); err != nil {
		glog.Error(err)
		os.Exit(1)
	}
}
