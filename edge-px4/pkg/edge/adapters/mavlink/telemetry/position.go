package mavlink

import (
	"context"
	"io"

	"edge-px4/pkg/edge"
	"edge-px4/pkg/edge/domain/common"
	mavsdk_rpc_telemetry "edge-px4/pkg/protos/telemetry"

	"google.golang.org/grpc"
)

// AdapterPosition .
func AdapterPosition(ctx context.Context, gr *grpc.ClientConn, support common.Support) (<-chan *edge.Position, error) {
	telemetry := mavsdk_rpc_telemetry.NewTelemetryServiceClient(gr)

	positionReceiver, err := AdapterPositionInternal(ctx, support, telemetry)
	if err != nil {
		return nil, err
	}

	positionStream := AdapterPositionSubscriber(positionReceiver, support)

	return positionStream, nil
}

// AdapterPositionInternal .
func AdapterPositionInternal(
	ctx context.Context,
	support common.Support,
	telemetry mavsdk_rpc_telemetry.TelemetryServiceClient,
) (positionReceiver mavsdk_rpc_telemetry.TelemetryService_SubscribePositionClient, err error) {
	defer func() {
		if err != nil {
			support.NotifyError("position telemetry error: %v", err)
		}
	}()

	positionRequest := mavsdk_rpc_telemetry.SubscribePositionRequest{}
	positionReceiver, err = telemetry.SubscribePosition(ctx, &positionRequest)

	return
}

// AdapterPositionSubscriber .
func AdapterPositionSubscriber(
	positionReceiver mavsdk_rpc_telemetry.TelemetryService_SubscribePositionClient,
	support common.Support,
) <-chan *edge.Position {
	positionStream := make(chan *edge.Position)

	go func() {
		defer func() {
			if err := positionReceiver.CloseSend(); err != nil {
				support.NotifyError("position telemetry error: %v", err)
			}
		}()
		defer close(positionStream)
		func() {
			var err error
			defer func() {
				if err != nil {
					support.NotifyError("position receive error: %v", err)
					return
				}
				support.NotifyInfo("position receive finish")
			}()
			for {
				response, ret := positionReceiver.Recv()
				if ret == io.EOF {
					return
				}
				if ret != nil {
					err = ret
					return
				}
				position := &edge.Position{
					Latitude:         response.GetPosition().LatitudeDeg,
					Longitude:        response.GetPosition().LongitudeDeg,
					Altitude:         float64(response.GetPosition().AbsoluteAltitudeM),
					RelativeAltitude: float64(response.GetPosition().RelativeAltitudeM),
				}
				positionStream <- position
			}
		}()
	}()

	return positionStream
}
