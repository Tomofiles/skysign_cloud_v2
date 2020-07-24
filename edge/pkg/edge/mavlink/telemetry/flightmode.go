package mavlink

import (
	"context"
	"io"
	"log"

	"google.golang.org/grpc"

	"edge/pkg/edge"
	mavsdk_rpc_telemetry "edge/pkg/protos/telemetry"
)

// AdapterFlightMode .
func AdapterFlightMode(ctx context.Context, url string) (<-chan *edge.FlightMode, error) {
	gr, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Println("grpc client connection error:", err)
		return nil, err
	}

	telemetry := mavsdk_rpc_telemetry.NewTelemetryServiceClient(gr)

	flightModeRequest := mavsdk_rpc_telemetry.SubscribeFlightModeRequest{}
	flightModeReceiver, err := telemetry.SubscribeFlightMode(ctx, &flightModeRequest)
	if err != nil {
		log.Println("flightMode request error:", err)
		return nil, err
	}

	flightModeStream := make(chan *edge.FlightMode)
	go func() {
		defer gr.Close()
		defer close(flightModeStream)
		for {
			response, err := flightModeReceiver.Recv()
			if err == io.EOF {
				log.Println("flightMode response io eof error:", err)
				return
			}
			if err != nil {
				log.Println("flightMode response other error:", err)
				return
			}
			flightMode := &edge.FlightMode{
				FlightMode: response.GetFlightMode().String(),
			}
			flightModeStream <- flightMode
		}
	}()

	return flightModeStream, nil
}
