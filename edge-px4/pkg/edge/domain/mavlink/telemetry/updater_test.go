package telemetry

import (
	"context"
	"edge-px4/pkg/edge/domain/model"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestTelemetryUpdaterContextDone .
func TestTelemetryUpdaterContextDone(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	supportMock := &supportMock{}

	tlm := model.NewTelemetry()

	connectionStateStream := make(chan *model.ConnectionState)
	positionStream := make(chan *model.Position)
	quaternionStream := make(chan *model.Quaternion)
	velocityStream := make(chan *model.Velocity)
	armedStream := make(chan *model.Armed)
	flightModeStream := make(chan *model.FlightMode)

	updateExit := Updater(
		ctx,
		supportMock,
		tlm,
		connectionStateStream,
		positionStream,
		quaternionStream,
		velocityStream,
		armedStream,
		flightModeStream,
	)

	cancel()

	<-updateExit

	expectTelemetry := model.NewTelemetry()

	a.Equal(expectTelemetry, tlm)
	a.Equal([]string{"telemetry updater done"}, supportMock.messages)
}

// TestTelemetryUpdaterConnectionState .
func TestTelemetryUpdaterConnectionState(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	tlm := model.NewTelemetry()

	connectionStateStream := make(chan *model.ConnectionState)
	positionStream := make(chan *model.Position)
	quaternionStream := make(chan *model.Quaternion)
	velocityStream := make(chan *model.Velocity)
	armedStream := make(chan *model.Armed)
	flightModeStream := make(chan *model.FlightMode)

	updateExit := Updater(
		ctx,
		supportMock,
		tlm,
		connectionStateStream,
		positionStream,
		quaternionStream,
		velocityStream,
		armedStream,
		flightModeStream,
	)

	response1 := &model.ConnectionState{
		VehicleID: DefaultEdgeVehicleID,
	}
	connectionStateStream <- response1
	close(connectionStateStream)

	<-updateExit

	expectTelemetry := model.NewTelemetry()
	expectTelemetry.SetConnectionState(&model.ConnectionState{VehicleID: DefaultEdgeVehicleID})

	a.Equal(expectTelemetry, tlm)
	a.Equal([]string{"connectionStateStream close"}, supportMock.messages)
}

// TestTelemetryUpdaterPosition .
func TestTelemetryUpdaterPosition(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	tlm := model.NewTelemetry()

	connectionStateStream := make(chan *model.ConnectionState)
	positionStream := make(chan *model.Position)
	quaternionStream := make(chan *model.Quaternion)
	velocityStream := make(chan *model.Velocity)
	armedStream := make(chan *model.Armed)
	flightModeStream := make(chan *model.FlightMode)

	updateExit := Updater(
		ctx,
		supportMock,
		tlm,
		connectionStateStream,
		positionStream,
		quaternionStream,
		velocityStream,
		armedStream,
		flightModeStream,
	)

	response1 := &model.Position{
		Latitude:         1.0,
		Longitude:        2.0,
		Altitude:         3.0,
		RelativeAltitude: 4.0,
	}
	positionStream <- response1
	close(positionStream)

	<-updateExit

	expectTelemetry := model.NewTelemetry()
	expectTelemetry.SetPosition(&model.Position{
		Latitude:         1.0,
		Longitude:        2.0,
		Altitude:         3.0,
		RelativeAltitude: 4.0,
	})

	a.Equal(expectTelemetry, tlm)
	a.Equal([]string{"positionStream close"}, supportMock.messages)
}

// TestTelemetryUpdaterQuaternion .
func TestTelemetryUpdaterQuaternion(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	tlm := model.NewTelemetry()

	connectionStateStream := make(chan *model.ConnectionState)
	positionStream := make(chan *model.Position)
	quaternionStream := make(chan *model.Quaternion)
	velocityStream := make(chan *model.Velocity)
	armedStream := make(chan *model.Armed)
	flightModeStream := make(chan *model.FlightMode)

	updateExit := Updater(
		ctx,
		supportMock,
		tlm,
		connectionStateStream,
		positionStream,
		quaternionStream,
		velocityStream,
		armedStream,
		flightModeStream,
	)

	response1 := &model.Quaternion{
		X: 1.0,
		Y: 2.0,
		Z: 3.0,
		W: 4.0,
	}
	quaternionStream <- response1
	close(quaternionStream)

	<-updateExit

	expectTelemetry := model.NewTelemetry()
	expectTelemetry.SetQuaternion(&model.Quaternion{
		X: 1.0,
		Y: 2.0,
		Z: 3.0,
		W: 4.0,
	})

	a.Equal(expectTelemetry, tlm)
	a.Equal([]string{"quaternionStream close"}, supportMock.messages)
}

// TestTelemetryUpdaterVelocity .
func TestTelemetryUpdaterVelocity(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	tlm := model.NewTelemetry()
	tlm.SetFlightMode(&model.FlightMode{FlightMode: "XXX"})

	connectionStateStream := make(chan *model.ConnectionState)
	positionStream := make(chan *model.Position)
	quaternionStream := make(chan *model.Quaternion)
	velocityStream := make(chan *model.Velocity)
	armedStream := make(chan *model.Armed)
	flightModeStream := make(chan *model.FlightMode)

	updateExit := Updater(
		ctx,
		supportMock,
		tlm,
		connectionStateStream,
		positionStream,
		quaternionStream,
		velocityStream,
		armedStream,
		flightModeStream,
	)

	response1 := &model.Velocity{
		North: 1.0,
		East:  2.0,
		Down:  3.0,
	}
	velocityStream <- response1
	close(velocityStream)

	<-updateExit

	snapshot, err := tlm.Get()

	a.Nil(err)
	a.Equal(math.Sqrt(1.0*1.0+2.0*2.0), snapshot.State.Speed) // GroundSpeed = √n^2+e^2）
	a.Equal([]string{"velocityStream close"}, supportMock.messages)
}

// TestTelemetryUpdaterArmed .
func TestTelemetryUpdaterArmed(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	tlm := model.NewTelemetry()

	connectionStateStream := make(chan *model.ConnectionState)
	positionStream := make(chan *model.Position)
	quaternionStream := make(chan *model.Quaternion)
	velocityStream := make(chan *model.Velocity)
	armedStream := make(chan *model.Armed)
	flightModeStream := make(chan *model.FlightMode)

	updateExit := Updater(
		ctx,
		supportMock,
		tlm,
		connectionStateStream,
		positionStream,
		quaternionStream,
		velocityStream,
		armedStream,
		flightModeStream,
	)

	response1 := &model.Armed{
		Armed: true,
	}
	armedStream <- response1
	close(armedStream)

	<-updateExit

	expectTelemetry := model.NewTelemetry()
	expectTelemetry.SetArmed(&model.Armed{
		Armed: true,
	})

	a.Equal(expectTelemetry, tlm)
	a.Equal([]string{"armedStream close"}, supportMock.messages)
}

// TestTelemetryUpdaterFlightMode .
func TestTelemetryUpdaterFlightMode(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	tlm := model.NewTelemetry()

	connectionStateStream := make(chan *model.ConnectionState)
	positionStream := make(chan *model.Position)
	quaternionStream := make(chan *model.Quaternion)
	velocityStream := make(chan *model.Velocity)
	armedStream := make(chan *model.Armed)
	flightModeStream := make(chan *model.FlightMode)

	updateExit := Updater(
		ctx,
		supportMock,
		tlm,
		connectionStateStream,
		positionStream,
		quaternionStream,
		velocityStream,
		armedStream,
		flightModeStream,
	)

	response1 := &model.FlightMode{
		FlightMode: "XXX",
	}
	flightModeStream <- response1
	close(flightModeStream)

	<-updateExit

	expectTelemetry := model.NewTelemetry()
	expectTelemetry.SetFlightMode(&model.FlightMode{
		FlightMode: "XXX",
	})

	a.Equal(expectTelemetry, tlm)
	a.Equal([]string{"flightModeStream close"}, supportMock.messages)
}

// TestTelemetryGet .
func TestTelemetryGet(t *testing.T) {
	a := assert.New(t)

	tlm := model.NewTelemetry()
	tlm.SetConnectionState(&model.ConnectionState{VehicleID: DefaultEdgeVehicleID})
	tlm.SetPosition(&model.Position{
		Latitude:         1.0,
		Longitude:        2.0,
		Altitude:         3.0,
		RelativeAltitude: 4.0,
	})
	tlm.SetQuaternion(&model.Quaternion{
		X: 6.0,
		Y: 7.0,
		Z: 8.0,
		W: 9.0,
	})
	tlm.SetVelocity(&model.Velocity{
		North: 1.0,
		East:  2.0,
		Down:  3.0,
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
			Latitude:         1.0,
			Longitude:        2.0,
			Altitude:         3.0,
			RelativeAltitude: 4.0,
			Speed:            math.Sqrt(1.0*1.0 + 2.0*2.0),
			Armed:            true,
			FlightMode:       "XXX",
			OrientationX:     6.0,
			OrientationY:     7.0,
			OrientationZ:     8.0,
			OrientationW:     9.0,
		},
	}

	a.Equal(expectTelemetry, snapshot)
	a.Nil(err)
}
