package main

import (
	"context"
	"edge/pkg/edge/adapters/glog"
	"edge/pkg/edge/adapters/grpc"
	"edge/pkg/edge/builder"
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

				support := glog.NewSupport()

				updateExit, telemetry, err := builder.MavlinkTelemetry(ctx, gr, support)
				if err != nil {
					log.Println("mavlink telemetry error:", err)
					cancel()
					continue
				}

				commandStream, missionStream := builder.Cloudlink(ctx, cloud, telemetry)

				err = builder.MavlinkCommand(ctx, mavsdk, commandStream, missionStream)
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
