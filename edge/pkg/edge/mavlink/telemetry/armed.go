package mavlink

import (
	"context"
	"io"
	"log"

	"google.golang.org/grpc"

	"edge/pkg/edge"
	mavsdk_rpc_telemetry "edge/pkg/protos/telemetry"
)

// AdapterArmed .
func AdapterArmed(ctx context.Context, url string) (<-chan *edge.Armed, error) {
	gr, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Println("grpc client connection error:", err)
		return nil, err
	}

	telemetry := mavsdk_rpc_telemetry.NewTelemetryServiceClient(gr)

	armedRequest := mavsdk_rpc_telemetry.SubscribeArmedRequest{}
	armedReceiver, err := telemetry.SubscribeArmed(ctx, &armedRequest)
	if err != nil {
		log.Println("armed request error:", err)
		return nil, err
	}

	armedStream := make(chan *edge.Armed)
	go func() {
		defer gr.Close()
		defer close(armedStream)
		for {
			response, err := armedReceiver.Recv()
			if err == io.EOF {
				log.Println("armed response io eof error:", err)
				return
			}
			if err != nil {
				log.Println("armed response other error:", err)
				return
			}
			armed := &edge.Armed{
				Armed: response.GetIsArmed(),
			}
			armedStream <- armed
		}
	}()

	return armedStream, nil
}
