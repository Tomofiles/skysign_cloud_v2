package mavlink

import (
	"context"
	"io"
	"log"
	"strconv"

	"google.golang.org/grpc"

	"edge/pkg/edge"

	mavsdk_rpc_core "edge/pkg/protos/core"
)

// AdapterConnectionState .
func AdapterConnectionState(ctx context.Context, url string) (<-chan *edge.ConnectionState, error) {
	gr, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Println("grpc client connection error:", err)
		return nil, err
	}

	core := mavsdk_rpc_core.NewCoreServiceClient(gr)

	connStateRequest := mavsdk_rpc_core.SubscribeConnectionStateRequest{}
	connStateReceiver, err := core.SubscribeConnectionState(ctx, &connStateRequest)
	if err != nil {
		log.Println("connState request error:", err)
		return nil, err
	}

	connStateStream := make(chan *edge.ConnectionState)
	go func() {
		defer gr.Close()
		defer close(connStateStream)
		for {
			response, err := connStateReceiver.Recv()
			if err == io.EOF {
				log.Println("connState response io eof error:", err)
				return
			}
			if err != nil {
				log.Println("connState response other error:", err)
				return
			}
			connState := &edge.ConnectionState{
				VehicleID: strconv.FormatUint(response.ConnectionState.GetUuid(), 10),
			}
			connStateStream <- connState
		}
	}()

	return connStateStream, nil
}
