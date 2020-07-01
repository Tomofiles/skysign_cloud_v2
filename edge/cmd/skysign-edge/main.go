package main

import (
	"edge/pkg/edge"
	"log"
	"os"
	"os/signal"
	"time"

	"google.golang.org/grpc"
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

	gr, err := grpc.Dial(mavsdk, grpc.WithInsecure())
	if err != nil {
		log.Fatal("grpc client connection error:", err)
	}
	defer gr.Close()

	mavlink := edge.NewMavlink(gr)
	mavlink.Listen()

	go mavlink.SendTelemetry(cloud)

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)

	<-stop

	time.Sleep(1 * time.Second)

	defer log.Printf("Skysign Edge end.")
}
