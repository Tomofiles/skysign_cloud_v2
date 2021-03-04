package main

import (
	"flag"

	"github.com/golang/glog"
)

var (
	port *string
)

// func run() error {
// 	ctx := context.Background()
// 	ctx, cancel := context.WithCancel(ctx)
// 	defer cancel()

// 	listen, err := net.Listen("tcp", ":"+*port)
// 	if err != nil {
// 		return err
// 	}
// 	defer listen.Close()
// 	s := grpc.NewServer(grpc.UnaryInterceptor(ports.LogBodyInterceptor()))

// db, err := postgresql.NewPostgresqlConnection()
// if err != nil {
// 	return err
// }
// txm := postgresql.NewGormTransactionManager(db)

// conn, err := rabbitmq.NewRabbitMQConnection()
// if err != nil {
// 	return err
// }
// defer conn.Close()
// psm := rabbitmq.NewPubSubManager(conn)

// application := app.NewApplication(ctx, txm, psm)

// svc := ports.NewGrpcServer(application)
// evt := ports.NewEventHandler(application)

// psm.SetConsumer(
// 	ctx,
// 	ports.FlightplanCreatedEventExchangeName,
// 	func(event []byte) {
// 		if err := evt.HandleCreatedEvent(ctx, event); err != nil {
// 			glog.Error(err)
// 		}
// 	},
// )
// psm.SetConsumer(
// 	ctx,
// 	ports.FlightplanDeletedEventExchangeName,
// 	func(event []byte) {
// 		if err := evt.HandleDeletedEvent(ctx, event); err != nil {
// 			glog.Error(err)
// 		}
// 	},
// )
// psm.SetConsumer(
// 	ctx,
// 	ports.FlightplanCopiedEventExchangeName,
// 	func(event []byte) {
// 		if err := evt.HandleCopiedEvent(ctx, event); err != nil {
// 			glog.Error(err)
// 		}
// 	},
// )

// application.Services.ManageFlightplan.CarbonCopyFlightplan(
// 	&carbonCopyRequestMock{
// 		o: "3f2a1599-5842-4ce6-aeac-a5456d124052",
// 		n: "b8961e11-5b55-49af-94ee-58a354ceafa5",
// 	},
// )

// proto.RegisterManageFlightplanServiceServer(s, &svc)
// proto.RegisterAssignAssetsToFlightplanServiceServer(s, &svc)

// 	glog.Info("start flightplan server")
// 	return s.Serve(listen)
// }

func main() {
	port = flag.String("port", "5001", "flightoperation port")
	flag.Parse()
	defer glog.Flush()

	// for {
	// 	if err := run(); err != nil {
	// 		glog.Error(err)
	// 		time.Sleep(10 * time.Second)
	// 	}
	// }
}

// type carbonCopyRequestMock struct {
// 	o, n string
// }

// func (r *carbonCopyRequestMock) GetOriginalID() string {
// 	return r.o
// }
// func (r *carbonCopyRequestMock) GetNewID() string {
// 	return r.n
// }
