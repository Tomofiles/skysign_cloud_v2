package mavlink

import (
	"context"
	"io"
	"log"

	"edge/pkg/edge"
	"edge/pkg/edge/adapters/glog"
	"edge/pkg/edge/adapters/grpc"
	"edge/pkg/edge/common"
	mavsdk_rpc_telemetry "edge/pkg/protos/telemetry"
)

// AdapterVelocity .
func AdapterVelocity(ctx context.Context, url string) (<-chan *edge.Velocity, error) {
	gr, err := grpc.NewGrpcClientConnectionWithBlock(url)
	if err != nil {
		log.Println("grpc client connection error:", err)
		return nil, err
	}

	telemetry := mavsdk_rpc_telemetry.NewTelemetryServiceClient(gr)

	velocityReceiver, err := AdapterVelocityInternal(ctx, glog.NewSupport(), telemetry)

	velocityStream := AdapterVelocitySubscriber(velocityReceiver, glog.NewSupport())

	return velocityStream, nil
}

// AdapterVelocityInternal .
func AdapterVelocityInternal(
	ctx context.Context,
	support common.Support,
	telemetry mavsdk_rpc_telemetry.TelemetryServiceClient,
) (velocityReceiver mavsdk_rpc_telemetry.TelemetryService_SubscribeGroundSpeedNedClient, err error) {
	defer func() {
		if err != nil {
			support.NotifyError("velocity telemetry error: %v", err)
		}
	}()

	velocityRequest := mavsdk_rpc_telemetry.SubscribeGroundSpeedNedRequest{}
	velocityReceiver, err = telemetry.SubscribeGroundSpeedNed(ctx, &velocityRequest)

	return
}

// AdapterVelocitySubscriber .
func AdapterVelocitySubscriber(
	velocityReceiver mavsdk_rpc_telemetry.TelemetryService_SubscribeGroundSpeedNedClient,
	support common.Support,
) <-chan *edge.Velocity {
	velocityStream := make(chan *edge.Velocity)

	go func() {
		defer func() {
			if err := velocityReceiver.CloseSend(); err != nil {
				support.NotifyError("velocity telemetry error: %v", err)
			}
		}()
		defer close(velocityStream)
		func() {
			var err error
			defer func() {
				if err != nil {
					support.NotifyError("velocity receive error: %v", err)
					return
				}
				support.NotifyInfo("velocity receive finish")
			}()
			for {
				response, ret := velocityReceiver.Recv()
				if ret == io.EOF {
					return
				}
				if ret != nil {
					err = ret
					return
				}
				velocity := &edge.Velocity{
					North: float64(response.GetGroundSpeedNed().GetVelocityNorthMS()),
					East:  float64(response.GetGroundSpeedNed().GetVelocityEastMS()),
					Down:  float64(response.GetGroundSpeedNed().GetVelocityDownMS()),
				}
				velocityStream <- velocity
			}
		}()
	}()

	return velocityStream
}
