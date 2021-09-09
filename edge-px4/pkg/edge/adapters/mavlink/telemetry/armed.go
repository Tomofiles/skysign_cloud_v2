package mavlink

import (
	"context"
	"io"

	"edge-px4/pkg/edge/domain/common"
	"edge-px4/pkg/edge/domain/model"
	mavsdk_rpc_telemetry "edge-px4/pkg/protos/telemetry"

	"google.golang.org/grpc"
)

// AdapterArmed .
func AdapterArmed(ctx context.Context, gr *grpc.ClientConn, support common.Support) (<-chan *model.Armed, error) {
	telemetry := mavsdk_rpc_telemetry.NewTelemetryServiceClient(gr)

	armedReceiver, err := AdapterArmedInternal(ctx, support, telemetry)
	if err != nil {
		return nil, err
	}

	armedStream := AdapterArmedSubscriber(armedReceiver, support)

	return armedStream, nil
}

// AdapterArmedInternal .
func AdapterArmedInternal(
	ctx context.Context,
	support common.Support,
	telemetry mavsdk_rpc_telemetry.TelemetryServiceClient,
) (armedReceiver mavsdk_rpc_telemetry.TelemetryService_SubscribeArmedClient, err error) {
	defer func() {
		if err != nil {
			support.NotifyError("armed telemetry error: %v", err)
		}
	}()

	armedRequest := mavsdk_rpc_telemetry.SubscribeArmedRequest{}
	armedReceiver, err = telemetry.SubscribeArmed(ctx, &armedRequest)

	return
}

// AdapterArmedSubscriber .
func AdapterArmedSubscriber(
	armedReceiver mavsdk_rpc_telemetry.TelemetryService_SubscribeArmedClient,
	support common.Support,
) <-chan *model.Armed {
	armedStream := make(chan *model.Armed)

	go func() {
		defer func() {
			if err := armedReceiver.CloseSend(); err != nil {
				support.NotifyError("armed telemetry error: %v", err)
			}
		}()
		defer close(armedStream)
		func() {
			var err error
			defer func() {
				if err != nil {
					support.NotifyError("armed receive error: %v", err)
					return
				}
				support.NotifyInfo("armed receive finish")
			}()
			for {
				response, ret := armedReceiver.Recv()
				if ret == io.EOF {
					return
				}
				if ret != nil {
					err = ret
					return
				}
				armed := &model.Armed{
					Armed: response.GetIsArmed(),
				}
				armedStream <- armed
			}
		}()
	}()

	return armedStream
}
