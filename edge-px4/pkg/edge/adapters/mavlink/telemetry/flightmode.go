package mavlink

import (
	"context"
	"io"

	"edge-px4/pkg/edge/domain/common"
	"edge-px4/pkg/edge/domain/model"
	mavsdk_rpc_telemetry "edge-px4/pkg/protos/telemetry"

	"google.golang.org/grpc"
)

// AdapterFlightMode .
func AdapterFlightMode(ctx context.Context, gr *grpc.ClientConn, support common.Support) (<-chan *model.FlightMode, error) {
	telemetry := mavsdk_rpc_telemetry.NewTelemetryServiceClient(gr)

	flightModeReceiver, err := AdapterFlightModeInternal(ctx, support, telemetry)
	if err != nil {
		return nil, err
	}

	flightModeStream := AdapterFlightModeSubscriber(flightModeReceiver, support)

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
) <-chan *model.FlightMode {
	flightModeStream := make(chan *model.FlightMode)

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
				flightMode := &model.FlightMode{
					FlightMode: response.GetFlightMode().String(),
				}
				flightModeStream <- flightMode
			}
		}()
	}()

	return flightModeStream
}
