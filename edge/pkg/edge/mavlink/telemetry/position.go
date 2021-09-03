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

// AdapterPosition .
func AdapterPosition(ctx context.Context, url string) (<-chan *edge.Position, error) {
	gr, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println("grpc client connection error:", err)
		return nil, err
	}

	telemetry := mavsdk_rpc_telemetry.NewTelemetryServiceClient(gr)

	positionReceiver, err := AdapterPositionInternal(ctx, common.NewSupport(), telemetry)

	positionStream := AdapterPositionSubscriber(positionReceiver, common.NewSupport())

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
