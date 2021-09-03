package mavlink

import (
	"context"
	"io"
	"log"

	"google.golang.org/grpc"

	"edge/pkg/edge"
	"edge/pkg/edge/common"
	mavsdk_rpc_telemetry "edge/pkg/protos/telemetry"
)

// AdapterFlightMode .
func AdapterFlightMode(ctx context.Context, url string) (<-chan *edge.FlightMode, error) {
	gr, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println("grpc client connection error:", err)
		return nil, err
	}

	telemetry := mavsdk_rpc_telemetry.NewTelemetryServiceClient(gr)

	flightModeReceiver, err := AdapterFlightModeInternal(ctx, common.NewSupport(), telemetry)

	flightModeStream := AdapterFlightModeSubscriber(flightModeReceiver, common.NewSupport())

	return flightModeStream, nil
}

// AdapterFlightModeInternal .
func AdapterFlightModeInternal(
	ctx context.Context,
	support common.Support,
	telemetry mavsdk_rpc_telemetry.TelemetryServiceClient,
) (flightModeReceiver mavsdk_rpc_telemetry.TelemetryService_SubscribeFlightModeClient, err error) {
	defer func() {
		if err != nil {
			support.NotifyError("flightMode telemetry error: %v", err)
		}
	}()

	flightModeRequest := mavsdk_rpc_telemetry.SubscribeFlightModeRequest{}
	flightModeReceiver, err = telemetry.SubscribeFlightMode(ctx, &flightModeRequest)

	return
}

// AdapterFlightModeSubscriber .
func AdapterFlightModeSubscriber(
	flightModeReceiver mavsdk_rpc_telemetry.TelemetryService_SubscribeFlightModeClient,
	support common.Support,
) <-chan *edge.FlightMode {
	flightModeStream := make(chan *edge.FlightMode)

	go func() {
		defer func() {
			if err := flightModeReceiver.CloseSend(); err != nil {
				support.NotifyError("flightMode telemetry error: %v", err)
			}
		}()
		defer close(flightModeStream)
		func() {
			var err error
			defer func() {
				if err != nil {
					support.NotifyError("flightMode receive error: %v", err)
					return
				}
				support.NotifyInfo("flightMode receive finish")
			}()
			for {
				response, ret := flightModeReceiver.Recv()
				if ret == io.EOF {
					return
				}
				if ret != nil {
					err = ret
					return
				}
				flightMode := &edge.FlightMode{
					FlightMode: response.GetFlightMode().String(),
				}
				flightModeStream <- flightMode
			}
		}()
	}()

	return flightModeStream
}
