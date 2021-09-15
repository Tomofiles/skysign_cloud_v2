package ports

import (
	"collection-analysis/pkg/action/app"
	act "collection-analysis/pkg/action/domain/action"
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSinglePointGetTrajectory(t *testing.T) {
	a := assert.New(t)

	service := manageActionServiceMock{}

	snapshots := []act.TelemetrySnapshot{
		{
			LatitudeDegree:    1.0,
			LongitudeDegree:   2.0,
			AltitudeM:         3.0,
			RelativeAltitudeM: 4.0,
			SpeedMS:           5.0,
			Armed:             true,
			FlightMode:        "state",
			OrientationX:      6.0,
			OrientationY:      7.0,
			OrientationZ:      8.0,
			OrientationW:      9.0,
		},
	}
	service.On("GetTrajectory", mock.Anything).Return(snapshots, nil)

	app := app.Application{
		Services: app.Services{
			ManageAction: &service,
		},
	}

	grpc := NewGrpcServer(app)

	request := &skysign_proto.GetTrajectoryRequest{
		VehicleId: string(DefaultActionID),
	}
	response, err := grpc.GetTrajectory(
		nil,
		request,
	)

	expectResponse := &skysign_proto.GetTrajectoryResponse{
		Telemetries: []*skysign_proto.Telemetry{
			{
				Latitude:         1.0,
				Longitude:        2.0,
				Altitude:         3.0,
				RelativeAltitude: 4.0,
				Speed:            5.0,
				Armed:            true,
				FlightMode:       "state",
				OrientationX:     6.0,
				OrientationY:     7.0,
				OrientationZ:     8.0,
				OrientationW:     9.0,
			},
		},
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestMultiplePointsGetTrajectory(t *testing.T) {
	a := assert.New(t)

	service := manageActionServiceMock{}

	snapshots := []act.TelemetrySnapshot{
		{
			LatitudeDegree:    1.0,
			LongitudeDegree:   2.0,
			AltitudeM:         3.0,
			RelativeAltitudeM: 4.0,
			SpeedMS:           5.0,
			Armed:             true,
			FlightMode:        "state1",
			OrientationX:      6.0,
			OrientationY:      7.0,
			OrientationZ:      8.0,
			OrientationW:      9.0,
		},
		{
			LatitudeDegree:    1.0,
			LongitudeDegree:   2.0,
			AltitudeM:         3.0,
			RelativeAltitudeM: 4.0,
			SpeedMS:           5.0,
			Armed:             true,
			FlightMode:        "state2",
			OrientationX:      6.0,
			OrientationY:      7.0,
			OrientationZ:      8.0,
			OrientationW:      9.0,
		},
		{
			LatitudeDegree:    1.0,
			LongitudeDegree:   2.0,
			AltitudeM:         3.0,
			RelativeAltitudeM: 4.0,
			SpeedMS:           5.0,
			Armed:             true,
			FlightMode:        "state3",
			OrientationX:      6.0,
			OrientationY:      7.0,
			OrientationZ:      8.0,
			OrientationW:      9.0,
		},
	}
	service.On("GetTrajectory", mock.Anything).Return(snapshots, nil)

	app := app.Application{
		Services: app.Services{
			ManageAction: &service,
		},
	}

	grpc := NewGrpcServer(app)

	request := &skysign_proto.GetTrajectoryRequest{
		VehicleId: string(DefaultActionID),
	}
	response, err := grpc.GetTrajectory(
		nil,
		request,
	)

	expectResponse := &skysign_proto.GetTrajectoryResponse{
		Telemetries: []*skysign_proto.Telemetry{
			{
				Latitude:         1.0,
				Longitude:        2.0,
				Altitude:         3.0,
				RelativeAltitude: 4.0,
				Speed:            5.0,
				Armed:            true,
				FlightMode:       "state1",
				OrientationX:     6.0,
				OrientationY:     7.0,
				OrientationZ:     8.0,
				OrientationW:     9.0,
			},
			{
				Latitude:         1.0,
				Longitude:        2.0,
				Altitude:         3.0,
				RelativeAltitude: 4.0,
				Speed:            5.0,
				Armed:            true,
				FlightMode:       "state2",
				OrientationX:     6.0,
				OrientationY:     7.0,
				OrientationZ:     8.0,
				OrientationW:     9.0,
			},
			{
				Latitude:         1.0,
				Longitude:        2.0,
				Altitude:         3.0,
				RelativeAltitude: 4.0,
				Speed:            5.0,
				Armed:            true,
				FlightMode:       "state3",
				OrientationX:     6.0,
				OrientationY:     7.0,
				OrientationZ:     8.0,
				OrientationW:     9.0,
			},
		},
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestNonePointsGetTrajectory(t *testing.T) {
	a := assert.New(t)

	service := manageActionServiceMock{}

	snapshots := []act.TelemetrySnapshot{}
	service.On("GetTrajectory", mock.Anything).Return(snapshots, nil)

	app := app.Application{
		Services: app.Services{
			ManageAction: &service,
		},
	}

	grpc := NewGrpcServer(app)

	request := &skysign_proto.GetTrajectoryRequest{
		VehicleId: string(DefaultActionID),
	}
	response, err := grpc.GetTrajectory(
		nil,
		request,
	)

	expectResponse := &skysign_proto.GetTrajectoryResponse{}

	a.Nil(err)
	a.Equal(response, expectResponse)
}
