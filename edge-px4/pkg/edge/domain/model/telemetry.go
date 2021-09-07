package model

import (
	"errors"
	"math"
	"sync"
)

var (
	ErrNotPrepared = errors.New("no telemetry prepared")
)

// Telemetry .
type Telemetry interface {
	SetConnectionState(connectionState *ConnectionState)
	SetPosition(position *Position)
	SetQuaternion(quaternion *Quaternion)
	SetVelocity(velocity *Velocity)
	SetArmed(armed *Armed)
	SetFlightMode(flightMode *FlightMode)
	Get() (*PushTelemetry, error)
}

type telemetry struct {
	rwm              sync.RWMutex
	id               string
	latitude         float64
	longitude        float64
	altitude         float64
	relativeAltitude float64
	speed            float64
	armed            bool
	flightMode       string
	orientationX     float64
	orientationY     float64
	orientationZ     float64
	orientationW     float64
}

// NewTelemetry .
func NewTelemetry() Telemetry {
	return &telemetry{}
}

func (t *telemetry) SetConnectionState(connectionState *ConnectionState) {
	t.rwm.Lock()
	defer t.rwm.Unlock()

	t.id = connectionState.VehicleID
}

func (t *telemetry) SetPosition(position *Position) {
	t.rwm.Lock()
	defer t.rwm.Unlock()

	t.latitude = position.Latitude
	t.longitude = position.Longitude
	t.altitude = position.Altitude
	t.relativeAltitude = position.RelativeAltitude
}

func (t *telemetry) SetQuaternion(quaternion *Quaternion) {
	t.rwm.Lock()
	defer t.rwm.Unlock()

	t.orientationX = quaternion.X
	t.orientationY = quaternion.Y
	t.orientationZ = quaternion.Z
	t.orientationW = quaternion.W
}

func (t *telemetry) SetVelocity(velocity *Velocity) {
	t.rwm.Lock()
	defer t.rwm.Unlock()

	// NEDフレームから速度の合成（GroundSpeed = √n^2+e^2）
	t.speed = math.Sqrt(velocity.North*velocity.North + velocity.East*velocity.East)
}

func (t *telemetry) SetArmed(armed *Armed) {
	t.rwm.Lock()
	defer t.rwm.Unlock()

	t.armed = armed.Armed
}

func (t *telemetry) SetFlightMode(flightMode *FlightMode) {
	t.rwm.Lock()
	defer t.rwm.Unlock()

	t.flightMode = flightMode.FlightMode
}

func (t *telemetry) Get() (*PushTelemetry, error) {
	rwmLocker := t.rwm.RLocker()
	rwmLocker.Lock()
	defer rwmLocker.Unlock()

	if t.flightMode == "" {
		return nil, ErrNotPrepared
	}

	telemetry := &PushTelemetry{
		ID: t.id,
		State: &State{
			Latitude:         t.latitude,
			Longitude:        t.longitude,
			Altitude:         t.altitude,
			RelativeAltitude: t.relativeAltitude,
			Speed:            t.speed,
			Armed:            t.armed,
			FlightMode:       t.flightMode,
			OrientationX:     t.orientationX,
			OrientationY:     t.orientationY,
			OrientationZ:     t.orientationZ,
			OrientationW:     t.orientationW,
		},
	}

	return telemetry, nil
}
