package main

import (
	"context"
	"edge/pkg/edge/builder"
	"log"
	"os"
	"os/signal"
	"time"
)

var (
	mavsdk = "localhost:50051"
	cloud  = "localhost:8889"
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

				updateExit, telemetry, err := builder.MavlinkTelemetry(ctx, mavsdk)
				if err != nil {
					log.Println("Mavlink telemetry error:", err)
					cancel()
					continue
				}

				builder.CloudlinkTelemetry(ctx, telemetry)

				go func() {
					t := time.NewTimer(5 * time.Second)
					select {
					case <-t.C:
						cancel()
					}
				}()

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
