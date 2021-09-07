package mavlink

import (
	"context"
	"io"

	"edge-px4/pkg/edge"
	"edge-px4/pkg/edge/domain/common"
	mavsdk_rpc_telemetry "edge-px4/pkg/protos/telemetry"

	"google.golang.org/grpc"
)

// AdapterQuaternion .
func AdapterQuaternion(ctx context.Context, gr *grpc.ClientConn, support common.Support) (<-chan *edge.Quaternion, error) {
	telemetry := mavsdk_rpc_telemetry.NewTelemetryServiceClient(gr)

	quaternionReceiver, err := AdapterQuaternionInternal(ctx, support, telemetry)
	if err != nil {
		return nil, err
	}

	quaternionStream := AdapterQuaternionSubscriber(quaternionReceiver, support)

	return quaternionStream, nil
}

// AdapterQuaternionInternal .
func AdapterQuaternionInternal(
	ctx context.Context,
	support common.Support,
	telemetry mavsdk_rpc_telemetry.TelemetryServiceClient,
) (quaternionReceiver mavsdk_rpc_telemetry.TelemetryService_SubscribeAttitudeQuaternionClient, err error) {
	defer func() {
		if err != nil {
			support.NotifyError("quaternion telemetry error: %v", err)
		}
	}()

	quaternionRequest := mavsdk_rpc_telemetry.SubscribeAttitudeQuaternionRequest{}
	quaternionReceiver, err = telemetry.SubscribeAttitudeQuaternion(ctx, &quaternionRequest)
	return
}

// AdapterQuaternionSubscriber .
func AdapterQuaternionSubscriber(
	quaternionReceiver mavsdk_rpc_telemetry.TelemetryService_SubscribeAttitudeQuaternionClient,
	support common.Support,
) <-chan *edge.Quaternion {
	quaternionStream := make(chan *edge.Quaternion)

	go func() {
		defer func() {
			if err := quaternionReceiver.CloseSend(); err != nil {
				support.NotifyError("quaternion telemetry error: %v", err)
			}
		}()
		defer close(quaternionStream)
		func() {
			var err error
			defer func() {
				if err != nil {
					support.NotifyError("quaternion receive error: %v", err)
					return
				}
				support.NotifyInfo("quaternion receive finish")
			}()
			for {
				response, ret := quaternionReceiver.Recv()
				if ret == io.EOF {
					return
				}
				if ret != nil {
					err = ret
					return
				}
				quaternion := &edge.Quaternion{
					X: float64(response.GetAttitudeQuaternion().GetX()),
					Y: float64(response.GetAttitudeQuaternion().GetY()),
					Z: float64(response.GetAttitudeQuaternion().GetZ()),
					W: float64(response.GetAttitudeQuaternion().GetW()),
				}
				quaternionStream <- quaternion
			}
		}()
	}()

	return quaternionStream
}
