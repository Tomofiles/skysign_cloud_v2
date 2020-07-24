package mavlink

import (
	"context"
	"io"
	"log"

	"google.golang.org/grpc"

	"edge/pkg/edge"
	mavsdk_rpc_telemetry "edge/pkg/protos/telemetry"
)

// AdapterVelocity .
func AdapterVelocity(ctx context.Context, url string) (<-chan *edge.Velocity, error) {
	gr, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Println("grpc client connection error:", err)
		return nil, err
	}

	telemetry := mavsdk_rpc_telemetry.NewTelemetryServiceClient(gr)

	velocityRequest := mavsdk_rpc_telemetry.SubscribeGroundSpeedNedRequest{}
	velocityReceiver, err := telemetry.SubscribeGroundSpeedNed(ctx, &velocityRequest)
	if err != nil {
		log.Println("velocity request error:", err)
		return nil, err
	}

	velocityStream := make(chan *edge.Velocity)
	go func() {
		defer gr.Close()
		defer close(velocityStream)
		for {
			response, err := velocityReceiver.Recv()
			if err == io.EOF {
				log.Println("velocity response io eof error:", err)
				return
			}
			if err != nil {
				log.Println("velocity response other error:", err)
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

	return velocityStream, nil
}
