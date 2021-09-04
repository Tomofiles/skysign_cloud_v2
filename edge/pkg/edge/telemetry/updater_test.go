package telemetry

import (
	"context"
	"edge/pkg/edge"
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

	tlm := NewTelemetry()

	connectionStateStream := make(chan *edge.ConnectionState)
	positionStream := make(chan *edge.Position)
	quaternionStream := make(chan *edge.Quaternion)
	velocityStream := make(chan *edge.Velocity)
	armedStream := make(chan *edge.Armed)
	flightModeStream := make(chan *edge.FlightMode)

	updateExit := Updater(
		ctx.Done(),
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

	expectTelemetry := &telemetry{}

	a.Equal(expectTelemetry, tlm)
	a.Equal([]string{"telemetry updater done"}, supportMock.messages)
}

// TestTelemetryUpdaterConnectionState .
func TestTelemetryUpdaterConnectionState(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	tlm := NewTelemetry()

	connectionStateStream := make(chan *edge.ConnectionState)
	positionStream := make(chan *edge.Position)
	quaternionStream := make(chan *edge.Quaternion)
	velocityStream := make(chan *edge.Velocity)
	armedStream := make(chan *edge.Armed)
	flightModeStream := make(chan *edge.FlightMode)

	updateExit := Updater(
		ctx.Done(),
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

	expectTelemetry := &telemetry{
		id: DefaultEdgeVehicleID,
	}

	a.Equal(expectTelemetry, tlm)
	a.Equal([]string{"connectionStateStream close"}, supportMock.messages)
}

// TestTelemetryUpdaterPosition .
func TestTelemetryUpdaterPosition(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	tlm := NewTelemetry()

	connectionStateStream := make(chan *edge.ConnectionState)
	positionStream := make(chan *edge.Position)
	quaternionStream := make(chan *edge.Quaternion)
	velocityStream := make(chan *edge.Velocity)
	armedStream := make(chan *edge.Armed)
	flightModeStream := make(chan *edge.FlightMode)

	updateExit := Updater(
		ctx.Done(),
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

	expectTelemetry := &telemetry{
		latitude:         1.0,
		longitude:        2.0,
		altitude:         3.0,
		relativeAltitude: 4.0,
	}

	a.Equal(expectTelemetry, tlm)
	a.Equal([]string{"positionStream close"}, supportMock.messages)
}

// TestTelemetryUpdaterQuaternion .
func TestTelemetryUpdaterQuaternion(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	tlm := NewTelemetry()

	connectionStateStream := make(chan *edge.ConnectionState)
	positionStream := make(chan *edge.Position)
	quaternionStream := make(chan *edge.Quaternion)
	velocityStream := make(chan *edge.Velocity)
	armedStream := make(chan *edge.Armed)
	flightModeStream := make(chan *edge.FlightMode)

	updateExit := Updater(
		ctx.Done(),
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

	expectTelemetry := &telemetry{
		orientationX: 1.0,
		orientationY: 2.0,
		orientationZ: 3.0,
		orientationW: 4.0,
	}

	a.Equal(expectTelemetry, tlm)
	a.Equal([]string{"quaternionStream close"}, supportMock.messages)
}

// TestTelemetryUpdaterVelocity .
func TestTelemetryUpdaterVelocity(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	tlm := NewTelemetry()

	connectionStateStream := make(chan *edge.ConnectionState)
	positionStream := make(chan *edge.Position)
	quaternionStream := make(chan *edge.Quaternion)
	velocityStream := make(chan *edge.Velocity)
	armedStream := make(chan *edge.Armed)
	flightModeStream := make(chan *edge.FlightMode)

	updateExit := Updater(
		ctx.Done(),
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

	// GroundSpeed = √n^2+e^2）
	expectTelemetry := &telemetry{
		speed: math.Sqrt(1.0*1.0 + 2.0*2.0),
	}

	a.Equal(expectTelemetry, tlm)
	a.Equal([]string{"velocityStream close"}, supportMock.messages)
}

// TestTelemetryUpdaterArmed .
func TestTelemetryUpdaterArmed(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	tlm := NewTelemetry()

	connectionStateStream := make(chan *edge.ConnectionState)
	positionStream := make(chan *edge.Position)
	quaternionStream := make(chan *edge.Quaternion)
	velocityStream := make(chan *edge.Velocity)
	armedStream := make(chan *edge.Armed)
	flightModeStream := make(chan *edge.FlightMode)

	updateExit := Updater(
		ctx.Done(),
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

	expectTelemetry := &telemetry{
		armed: true,
	}

	a.Equal(expectTelemetry, tlm)
	a.Equal([]string{"armedStream close"}, supportMock.messages)
}

// TestTelemetryUpdaterFlightMode .
func TestTelemetryUpdaterFlightMode(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	tlm := NewTelemetry()

	connectionStateStream := make(chan *edge.ConnectionState)
	positionStream := make(chan *edge.Position)
	quaternionStream := make(chan *edge.Quaternion)
	velocityStream := make(chan *edge.Velocity)
	armedStream := make(chan *edge.Armed)
	flightModeStream := make(chan *edge.FlightMode)

	updateExit := Updater(
		ctx.Done(),
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

	expectTelemetry := &telemetry{
		flightMode: "XXX",
	}

	a.Equal(expectTelemetry, tlm)
	a.Equal([]string{"flightModeStream close"}, supportMock.messages)
}

// TestTelemetryGet .
func TestTelemetryGet(t *testing.T) {
	a := assert.New(t)

	tlm := NewTelemetry().(*telemetry)
	tlm.id = DefaultEdgeVehicleID
	tlm.latitude = 1.0
	tlm.longitude = 2.0
	tlm.altitude = 3.0
	tlm.relativeAltitude = 4.0
	tlm.speed = 5.0
	tlm.armed = true
	tlm.flightMode = "XXX"
	tlm.orientationX = 6.0
	tlm.orientationY = 7.0
	tlm.orientationZ = 8.0
	tlm.orientationW = 9.0

	snapshot := tlm.Get()

	expectTelemetry := &edge.Telemetry{
		ID: DefaultEdgeVehicleID,
		State: &edge.State{
			Latitude:         1.0,
			Longitude:        2.0,
			Altitude:         3.0,
			RelativeAltitude: 4.0,
			Speed:            5.0,
			Armed:            true,
			FlightMode:       "XXX",
			OrientationX:     6.0,
			OrientationY:     7.0,
			OrientationZ:     8.0,
			OrientationW:     9.0,
		},
	}

	a.Equal(expectTelemetry, snapshot)
}
