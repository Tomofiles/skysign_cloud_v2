package telemetry

import (
	"context"
	"edge-px4/pkg/edge"
	"edge-px4/pkg/edge/domain/telemetry"
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

	tlm := telemetry.NewTelemetry()

	connectionStateStream := make(chan *edge.ConnectionState)
	positionStream := make(chan *edge.Position)
	quaternionStream := make(chan *edge.Quaternion)
	velocityStream := make(chan *edge.Velocity)
	armedStream := make(chan *edge.Armed)
	flightModeStream := make(chan *edge.FlightMode)

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

	expectTelemetry := telemetry.NewTelemetry()

	a.Equal(expectTelemetry, tlm)
	a.Equal([]string{"telemetry updater done"}, supportMock.messages)
}

// TestTelemetryUpdaterConnectionState .
func TestTelemetryUpdaterConnectionState(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	tlm := telemetry.NewTelemetry()

	connectionStateStream := make(chan *edge.ConnectionState)
	positionStream := make(chan *edge.Position)
	quaternionStream := make(chan *edge.Quaternion)
	velocityStream := make(chan *edge.Velocity)
	armedStream := make(chan *edge.Armed)
	flightModeStream := make(chan *edge.FlightMode)

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

	response1 := &edge.ConnectionState{
		VehicleID: DefaultEdgeVehicleID,
	}
	connectionStateStream <- response1
	close(connectionStateStream)

	<-updateExit

	expectTelemetry := telemetry.NewTelemetry()
	expectTelemetry.SetConnectionState(&edge.ConnectionState{VehicleID: DefaultEdgeVehicleID})

	a.Equal(expectTelemetry, tlm)
	a.Equal([]string{"connectionStateStream close"}, supportMock.messages)
}

// TestTelemetryUpdaterPosition .
func TestTelemetryUpdaterPosition(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	tlm := telemetry.NewTelemetry()

	connectionStateStream := make(chan *edge.ConnectionState)
	positionStream := make(chan *edge.Position)
	quaternionStream := make(chan *edge.Quaternion)
	velocityStream := make(chan *edge.Velocity)
	armedStream := make(chan *edge.Armed)
	flightModeStream := make(chan *edge.FlightMode)

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

	response1 := &edge.Position{
		Latitude:         1.0,
		Longitude:        2.0,
		Altitude:         3.0,
		RelativeAltitude: 4.0,
	}
	positionStream <- response1
	close(positionStream)

	<-updateExit

	expectTelemetry := telemetry.NewTelemetry()
	expectTelemetry.SetPosition(&edge.Position{
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

	tlm := telemetry.NewTelemetry()

	connectionStateStream := make(chan *edge.ConnectionState)
	positionStream := make(chan *edge.Position)
	quaternionStream := make(chan *edge.Quaternion)
	velocityStream := make(chan *edge.Velocity)
	armedStream := make(chan *edge.Armed)
	flightModeStream := make(chan *edge.FlightMode)

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

	response1 := &edge.Quaternion{
		X: 1.0,
		Y: 2.0,
		Z: 3.0,
		W: 4.0,
	}
	quaternionStream <- response1
	close(quaternionStream)

	<-updateExit

	expectTelemetry := telemetry.NewTelemetry()
	expectTelemetry.SetQuaternion(&edge.Quaternion{
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

	tlm := telemetry.NewTelemetry()
	tlm.SetFlightMode(&edge.FlightMode{FlightMode: "XXX"})

	connectionStateStream := make(chan *edge.ConnectionState)
	positionStream := make(chan *edge.Position)
	quaternionStream := make(chan *edge.Quaternion)
	velocityStream := make(chan *edge.Velocity)
	armedStream := make(chan *edge.Armed)
	flightModeStream := make(chan *edge.FlightMode)

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

	response1 := &edge.Velocity{
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

	tlm := telemetry.NewTelemetry()

	connectionStateStream := make(chan *edge.ConnectionState)
	positionStream := make(chan *edge.Position)
	quaternionStream := make(chan *edge.Quaternion)
	velocityStream := make(chan *edge.Velocity)
	armedStream := make(chan *edge.Armed)
	flightModeStream := make(chan *edge.FlightMode)

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

	response1 := &edge.Armed{
		Armed: true,
	}
	armedStream <- response1
	close(armedStream)

	<-updateExit

	expectTelemetry := telemetry.NewTelemetry()
	expectTelemetry.SetArmed(&edge.Armed{
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

	tlm := telemetry.NewTelemetry()

	connectionStateStream := make(chan *edge.ConnectionState)
	positionStream := make(chan *edge.Position)
	quaternionStream := make(chan *edge.Quaternion)
	velocityStream := make(chan *edge.Velocity)
	armedStream := make(chan *edge.Armed)
	flightModeStream := make(chan *edge.FlightMode)

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

	response1 := &edge.FlightMode{
		FlightMode: "XXX",
	}
	flightModeStream <- response1
	close(flightModeStream)

	<-updateExit

	expectTelemetry := telemetry.NewTelemetry()
	expectTelemetry.SetFlightMode(&edge.FlightMode{
		FlightMode: "XXX",
	})

	a.Equal(expectTelemetry, tlm)
	a.Equal([]string{"flightModeStream close"}, supportMock.messages)
}

// TestTelemetryGet .
func TestTelemetryGet(t *testing.T) {
	a := assert.New(t)

	tlm := telemetry.NewTelemetry()
	tlm.SetConnectionState(&edge.ConnectionState{VehicleID: DefaultEdgeVehicleID})
	tlm.SetPosition(&edge.Position{
		Latitude:         1.0,
		Longitude:        2.0,
		Altitude:         3.0,
		RelativeAltitude: 4.0,
	})
	tlm.SetQuaternion(&edge.Quaternion{
		X: 6.0,
		Y: 7.0,
		Z: 8.0,
		W: 9.0,
	})
	tlm.SetVelocity(&edge.Velocity{
		North: 1.0,
		East:  2.0,
		Down:  3.0,
	})
	tlm.SetArmed(&edge.Armed{
		Armed: true,
	})
	tlm.SetFlightMode(&edge.FlightMode{
		FlightMode: "XXX",
	})

	snapshot, err := tlm.Get()

	expectTelemetry := &edge.Telemetry{
		ID: DefaultEdgeVehicleID,
		State: &edge.State{
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
