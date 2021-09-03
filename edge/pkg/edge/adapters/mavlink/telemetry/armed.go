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

// AdapterArmed .
func AdapterArmed(ctx context.Context, url string) (<-chan *edge.Armed, error) {
	gr, err := grpc.NewGrpcClientConnectionWithBlock(url)
	if err != nil {
		log.Println("grpc client connection error:", err)
		return nil, err
	}

	telemetry := mavsdk_rpc_telemetry.NewTelemetryServiceClient(gr)

	armedReceiver, err := AdapterArmedInternal(ctx, glog.NewSupport(), telemetry)

	armedStream := AdapterArmedSubscriber(armedReceiver, glog.NewSupport())

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
) <-chan *edge.Armed {
	armedStream := make(chan *edge.Armed)

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
				armed := &edge.Armed{
					Armed: response.GetIsArmed(),
				}
				armedStream <- armed
			}
		}()
	}()

	return armedStream
}
