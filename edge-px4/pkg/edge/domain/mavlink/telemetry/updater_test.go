package telemetry

import (
	"context"
	"edge-px4/pkg/edge/domain/model"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestTelemetryUpdaterConnectionStateContextDone .
func TestTelemetryUpdaterConnectionStateContextDone(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	supportMock := &supportMock{}

	tlm := model.NewTelemetry()

	connectionStateStream := make(chan *model.ConnectionState)

	updaterExit := ConnectionStateUpdater(
		ctx,
		supportMock,
		tlm,
		connectionStateStream,
	)

	cancel()

	<-updaterExit

	expectTelemetry := model.NewTelemetry()

	a.Equal(expectTelemetry, tlm)
	a.Equal([]string{"telemetry CONNECTION_STATE done"}, supportMock.messages)
}

// TestTelemetryUpdaterConnectionState .
func TestTelemetryUpdaterConnectionState(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	tlm := model.NewTelemetry()

	connectionStateStream := make(chan *model.ConnectionState)

	updaterExit := ConnectionStateUpdater(
		ctx,
		supportMock,
		tlm,
		connectionStateStream,
	)

	response1 := &model.ConnectionState{
		VehicleID: DefaultEdgeVehicleID,
	}
	connectionStateStream <- response1
	close(connectionStateStream)

	<-updaterExit

	expectTelemetry := model.NewTelemetry()
	expectTelemetry.SetConnectionState(&model.ConnectionState{VehicleID: DefaultEdgeVehicleID})

	a.Equal(expectTelemetry, tlm)
	a.Equal([]string{"telemetry CONNECTION_STATE close"}, supportMock.messages)
}

// TestTelemetryUpdaterPositionContextDone .
func TestTelemetryUpdaterPositionContextDone(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	supportMock := &supportMock{}

	tlm := model.NewTelemetry()

	positionStream := make(chan *model.Position)

	updaterExit := PositionUpdater(
		ctx,
		supportMock,
		tlm,
		positionStream,
	)

	cancel()

	<-updaterExit

	expectTelemetry := model.NewTelemetry()

	a.Equal(expectTelemetry, tlm)
	a.Equal([]string{"telemetry POSITION done"}, supportMock.messages)
}

// TestTelemetryUpdaterPosition .
func TestTelemetryUpdaterPosition(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	tlm := model.NewTelemetry()

	positionStream := make(chan *model.Position)

	updaterExit := PositionUpdater(
		ctx,
		supportMock,
		tlm,
		positionStream,
	)

	response1 := &model.Position{
		LatitudeDegree:    1.0,
		LongitudeDegree:   2.0,
		AltitudeM:         3.0,
		RelativeAltitudeM: 4.0,
	}
	positionStream <- response1
	close(positionStream)

	<-updaterExit

	expectTelemetry := model.NewTelemetry()
	expectTelemetry.SetPosition(&model.Position{
		LatitudeDegree:    1.0,
		LongitudeDegree:   2.0,
		AltitudeM:         3.0,
		RelativeAltitudeM: 4.0,
	})

	a.Equal(expectTelemetry, tlm)
	a.Equal([]string{"telemetry POSITION close"}, supportMock.messages)
}

// TestTelemetryUpdaterQuaternionContextDone .
func TestTelemetryUpdaterQuaternionContextDone(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	supportMock := &supportMock{}

	tlm := model.NewTelemetry()

	quaternionStream := make(chan *model.Quaternion)

	updaterExit := QuaternionUpdater(
		ctx,
		supportMock,
		tlm,
		quaternionStream,
	)

	cancel()

	<-updaterExit

	expectTelemetry := model.NewTelemetry()

	a.Equal(expectTelemetry, tlm)
	a.Equal([]string{"telemetry QUATERNION done"}, supportMock.messages)
}

// TestTelemetryUpdaterQuaternion .
func TestTelemetryUpdaterQuaternion(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	tlm := model.NewTelemetry()

	quaternionStream := make(chan *model.Quaternion)

	updaterExit := QuaternionUpdater(
		ctx,
		supportMock,
		tlm,
		quaternionStream,
	)

	response1 := &model.Quaternion{
		X: 1.0,
		Y: 2.0,
		Z: 3.0,
		W: 4.0,
	}
	quaternionStream <- response1
	close(quaternionStream)

	<-updaterExit

	expectTelemetry := model.NewTelemetry()
	expectTelemetry.SetQuaternion(&model.Quaternion{
		X: 1.0,
		Y: 2.0,
		Z: 3.0,
		W: 4.0,
	})

	a.Equal(expectTelemetry, tlm)
	a.Equal([]string{"telemetry QUATERNION close"}, supportMock.messages)
}

// TestTelemetryUpdaterVelocityContextDone .
func TestTelemetryUpdaterVelocityContextDone(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	supportMock := &supportMock{}

	tlm := model.NewTelemetry()

	velocityStream := make(chan *model.Velocity)

	updaterExit := VelocityUpdater(
		ctx,
		supportMock,
		tlm,
		velocityStream,
	)

	cancel()

	<-updaterExit

	expectTelemetry := model.NewTelemetry()

	a.Equal(expectTelemetry, tlm)
	a.Equal([]string{"telemetry VELOCITY done"}, supportMock.messages)
}

// TestTelemetryUpdaterVelocity .
func TestTelemetryUpdaterVelocity(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	tlm := model.NewTelemetry()
	tlm.SetFlightMode(&model.FlightMode{FlightMode: "XXX"})

	velocityStream := make(chan *model.Velocity)

	updaterExit := VelocityUpdater(
		ctx,
		supportMock,
		tlm,
		velocityStream,
	)

	response1 := &model.Velocity{
		NorthMS: 1.0,
		EastMS:  2.0,
		DownMS:  3.0,
	}
	velocityStream <- response1
	close(velocityStream)

	<-updaterExit

	snapshot, err := tlm.Get()

	a.Nil(err)
	a.Equal(math.Sqrt(1.0*1.0+2.0*2.0), snapshot.State.SpeedMS) // GroundSpeed = √n^2+e^2）
	a.Equal([]string{"telemetry VELOCITY close"}, supportMock.messages)
}

// TestTelemetryUpdaterArmedContextDone .
func TestTelemetryUpdaterArmedContextDone(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	supportMock := &supportMock{}

	tlm := model.NewTelemetry()

	armedStream := make(chan *model.Armed)

	updaterExit := ArmedUpdater(
		ctx,
		supportMock,
		tlm,
		armedStream,
	)

	cancel()

	<-updaterExit

	expectTelemetry := model.NewTelemetry()

	a.Equal(expectTelemetry, tlm)
	a.Equal([]string{"telemetry ARMED done"}, supportMock.messages)
}

// TestTelemetryUpdaterArmed .
func TestTelemetryUpdaterArmed(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	tlm := model.NewTelemetry()

	armedStream := make(chan *model.Armed)

	updaterExit := ArmedUpdater(
		ctx,
		supportMock,
		tlm,
		armedStream,
	)

	response1 := &model.Armed{
		Armed: true,
	}
	armedStream <- response1
	close(armedStream)

	<-updaterExit

	expectTelemetry := model.NewTelemetry()
	expectTelemetry.SetArmed(&model.Armed{
		Armed: true,
	})

	a.Equal(expectTelemetry, tlm)
	a.Equal([]string{"telemetry ARMED close"}, supportMock.messages)
}

// TestTelemetryUpdaterFlightModeContextDone .
func TestTelemetryUpdaterFlightModeContextDone(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	supportMock := &supportMock{}

	tlm := model.NewTelemetry()

	flightModeStream := make(chan *model.FlightMode)

	updaterExit := FlightModeUpdater(
		ctx,
		supportMock,
		tlm,
		flightModeStream,
	)

	cancel()

	<-updaterExit

	expectTelemetry := model.NewTelemetry()

	a.Equal(expectTelemetry, tlm)
	a.Equal([]string{"telemetry FLIGHT_MODE done"}, supportMock.messages)
}

// TestTelemetryUpdaterFlightMode .
func TestTelemetryUpdaterFlightMode(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	tlm := model.NewTelemetry()

	flightModeStream := make(chan *model.FlightMode)

	updaterExit := FlightModeUpdater(
		ctx,
		supportMock,
		tlm,
		flightModeStream,
	)

	response1 := &model.FlightMode{
		FlightMode: "XXX",
	}
	flightModeStream <- response1
	close(flightModeStream)

	<-updaterExit

	expectTelemetry := model.NewTelemetry()
	expectTelemetry.SetFlightMode(&model.FlightMode{
		FlightMode: "XXX",
	})

	a.Equal(expectTelemetry, tlm)
	a.Equal([]string{"telemetry FLIGHT_MODE close"}, supportMock.messages)
}

// TestTelemetryGet .
func TestTelemetryGet(t *testing.T) {
	a := assert.New(t)

	tlm := model.NewTelemetry()
	tlm.SetConnectionState(&model.ConnectionState{VehicleID: DefaultEdgeVehicleID})
	tlm.SetPosition(&model.Position{
		LatitudeDegree:    1.0,
		LongitudeDegree:   2.0,
		AltitudeM:         3.0,
		RelativeAltitudeM: 4.0,
	})
	tlm.SetQuaternion(&model.Quaternion{
		X: 6.0,
		Y: 7.0,
		Z: 8.0,
		W: 9.0,
	})
	tlm.SetVelocity(&model.Velocity{
		NorthMS: 1.0,
		EastMS:  2.0,
		DownMS:  3.0,
	})
	tlm.SetArmed(&model.Armed{
		Armed: true,
	})
	tlm.SetFlightMode(&model.FlightMode{
		FlightMode: "XXX",
	})

	snapshot, err := tlm.Get()

	expectTelemetry := &model.PushTelemetry{
		ID: DefaultEdgeVehicleID,
		State: &model.State{
			LatitudeDegree:    1.0,
			LongitudeDegree:   2.0,
			AltitudeM:         3.0,
			RelativeAltitudeM: 4.0,
			SpeedMS:           math.Sqrt(1.0*1.0 + 2.0*2.0),
			Armed:             true,
			FlightMode:        "XXX",
			OrientationX:      6.0,
			OrientationY:      7.0,
			OrientationZ:      8.0,
			OrientationW:      9.0,
		},
	}

	a.Equal(expectTelemetry, snapshot)
	a.Nil(err)
}
