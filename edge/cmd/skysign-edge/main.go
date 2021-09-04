package main

import (
	"context"
	"edge/pkg/edge/adapters/glog"
	"edge/pkg/edge/adapters/grpc"
	"edge/pkg/edge/builder"
	"edge/pkg/edge/telemetry"
	"log"
	"os"
	"os/signal"
	"time"
)

var (
	mavsdk = "localhost:50051"
	cloud  = "http://localhost:8080"
)

func main() {
	mavsdkAddressEnv := os.Getenv("MAVSDK_ADDRESS")
	cloudAddressEnv := os.Getenv("CLOUD_ADDRESS")

	if mavsdkAddressEnv != "" {
		mavsdk = mavsdkAddressEnv
	}
	if cloudAddressEnv != "" {
		cloud = cloudAddressEnv
	}

	support := glog.NewSupport()

	go func() {
		for {
			t := time.NewTimer(1 * time.Second)
			select {
			case <-t.C:
				ctx := context.Background()
				ctx, cancel := context.WithCancel(ctx)

				gr, err := grpc.NewGrpcClientConnectionWithBlock(mavsdk)
				if err != nil {
					log.Println("grpc client connection error:", err)
					continue
				}

				telemetryStream, err := builder.MavlinkTelemetry(ctx, gr, support)
				if err != nil {
					log.Println("mavlink telemetry error:", err)
					cancel()
					continue
				}

				telemetry := telemetry.NewTelemetry()
				updateExit := telemetry.Updater(
					ctx.Done(),
					support,
					telemetryStream.ConnectionStateStream,
					telemetryStream.PositionStream,
					telemetryStream.QuaternionStream,
					telemetryStream.VelocityStream,
					telemetryStream.ArmedStream,
					telemetryStream.FlightModeStream,
				)

				commandStream := builder.Cloudlink(ctx, cloud, telemetry)

				err = builder.MavlinkCommand(
					ctx,
					gr,
					support,
					commandStream.CommandStream,
					commandStream.MissionStream,
				)
				if err != nil {
					log.Println("mavlink command error:", err)
					cancel()
					continue
				}

				// 障害時動作確認用
				// go func() {
				// 	t := time.NewTimer(5 * time.Second)
				// 	select {
				// 	case <-t.C:
				// 		cancel()
				// 	}
				// }()

				<-updateExit
				log.Println("update exit.")
				cancel()
			}
		}
	}()

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)

	<-stop

	time.Sleep(1 * time.Second)

	defer log.Printf("Skysign Edge end.")
}
