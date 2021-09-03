package mavlink

import (
	"context"
	"io"
	"log"

	"google.golang.org/grpc"

	"edge/pkg/edge"
	"edge/pkg/edge/adapters/glog"
	"edge/pkg/edge/common"
	mavsdk_rpc_telemetry "edge/pkg/protos/telemetry"
)

// AdapterQuaternion .
func AdapterQuaternion(ctx context.Context, url string) (<-chan *edge.Quaternion, error) {
	gr, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println("grpc client connection error:", err)
		return nil, err
	}

	telemetry := mavsdk_rpc_telemetry.NewTelemetryServiceClient(gr)

	quaternionReceiver, err := AdapterQuaternionInternal(ctx, glog.NewSupport(), telemetry)

	quaternionStream := AdapterQuaternionSubscriber(quaternionReceiver, glog.NewSupport())

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
