package mavlink

import (
	"context"
	"io"
	"strconv"

	"edge/pkg/edge"
	"edge/pkg/edge/domain/common"

	mavsdk_rpc_core "edge/pkg/protos/core"

	"google.golang.org/grpc"
)

// AdapterConnectionState .
func AdapterConnectionState(ctx context.Context, gr *grpc.ClientConn, support common.Support) (<-chan *edge.ConnectionState, error) {
	core := mavsdk_rpc_core.NewCoreServiceClient(gr)

	connectionStateReceiver, err := AdapterConnectionStateInternal(ctx, support, core)
	if err != nil {
		return nil, err
	}

	connectionStateStream := AdapterConnectionStateSubscriber(connectionStateReceiver, support)

	return connectionStateStream, nil
}

// AdapterConnectionStateInternal .
func AdapterConnectionStateInternal(
	ctx context.Context,
	support common.Support,
	core mavsdk_rpc_core.CoreServiceClient,
) (connectionStateReceiver mavsdk_rpc_core.CoreService_SubscribeConnectionStateClient, err error) {
	defer func() {
		if err != nil {
			support.NotifyError("connectionState core error: %v", err)
		}
	}()

	connectionStateRequest := mavsdk_rpc_core.SubscribeConnectionStateRequest{}
	connectionStateReceiver, err = core.SubscribeConnectionState(ctx, &connectionStateRequest)

	return
}

// AdapterConnectionStateSubscriber .
func AdapterConnectionStateSubscriber(
	connectionStateReceiver mavsdk_rpc_core.CoreService_SubscribeConnectionStateClient,
	support common.Support,
) <-chan *edge.ConnectionState {
	connectionStateStream := make(chan *edge.ConnectionState)

	go func() {
		defer func() {
			if err := connectionStateReceiver.CloseSend(); err != nil {
				support.NotifyError("connectionState core error: %v", err)
			}
		}()
		defer close(connectionStateStream)
		func() {
			var err error
			defer func() {
				if err != nil {
					support.NotifyError("connectionState receive error: %v", err)
					return
				}
				support.NotifyInfo("connectionState receive finish")
			}()
			for {
				response, ret := connectionStateReceiver.Recv()
				if ret == io.EOF {
					return
				}
				if ret != nil {
					err = ret
					return
				}
				connectionState := &edge.ConnectionState{
					VehicleID: strconv.FormatUint(response.ConnectionState.GetUuid(), 10),
				}
				connectionStateStream <- connectionState
			}
		}()
	}()

	return connectionStateStream
}
