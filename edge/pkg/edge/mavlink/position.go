package mavlink

import (
	"context"
	"io"
	"log"

	"google.golang.org/grpc"

	"edge/pkg/edge"
	mavsdk_rpc_telemetry "edge/pkg/protos/telemetry"
)

// AdapterPosition .
func AdapterPosition(ctx context.Context, url string) (<-chan *edge.Position, error) {
	gr, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Println("grpc client connection error:", err)
		return nil, err
	}

	telemetry := mavsdk_rpc_telemetry.NewTelemetryServiceClient(gr)

	positionRequest := mavsdk_rpc_telemetry.SubscribePositionRequest{}
	positionReceiver, err := telemetry.SubscribePosition(ctx, &positionRequest)
	if err != nil {
		log.Println("position request error:", err)
		return nil, err
	}

	positionStream := make(chan *edge.Position)
	go func() {
		defer gr.Close()
		defer close(positionStream)
		for {
			response, err := positionReceiver.Recv()
			if err == io.EOF {
				log.Println("position response io eof error:", err)
				return
			}
			if err != nil {
				log.Println("position response other error:", err)
				return
			}
			position := &edge.Position{
				Latitude:         response.GetPosition().LatitudeDeg,
				Longitude:        response.GetPosition().LongitudeDeg,
				AbsoluteAltitude: float64(response.GetPosition().AbsoluteAltitudeM),
				RelativeAltitude: float64(response.GetPosition().RelativeAltitudeM),
			}
			positionStream <- position
		}
	}()

	return positionStream, nil
}
