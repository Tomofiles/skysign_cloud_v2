package mavlink

import (
	"context"
	"io"
	"log"

	"google.golang.org/grpc"

	"edge/pkg/edge"
	mavsdk_rpc_telemetry "edge/pkg/protos/telemetry"
)

// AdapterQuaternion .
func AdapterQuaternion(ctx context.Context, url string) (<-chan *edge.Quaternion, error) {
	gr, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Println("grpc client connection error:", err)
		return nil, err
	}

	telemetry := mavsdk_rpc_telemetry.NewTelemetryServiceClient(gr)

	quaternionRequest := mavsdk_rpc_telemetry.SubscribeAttitudeQuaternionRequest{}
	quaternionReceiver, err := telemetry.SubscribeAttitudeQuaternion(ctx, &quaternionRequest)
	if err != nil {
		log.Println("quaternion request error:", err)
		return nil, err
	}

	quaternionStream := make(chan *edge.Quaternion)
	go func() {
		defer gr.Close()
		defer close(quaternionStream)
		for {
			response, err := quaternionReceiver.Recv()
			if err == io.EOF {
				log.Println("quaternion response io eof error:", err)
				return
			}
			if err != nil {
				log.Println("quaternion response other error:", err)
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

	return quaternionStream, nil
}
