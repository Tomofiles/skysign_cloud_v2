package telemetry

import (
	"edge-px4/pkg/edge"
	"errors"
	"math"
	"sync"
)

var (
	ErrNotPrepared = errors.New("no telemetry prepared")
)

// Telemetry struct
type Telemetry interface {
	SetConnectionState(connectionState *edge.ConnectionState)
	SetPosition(position *edge.Position)
	SetQuaternion(quaternion *edge.Quaternion)
	SetVelocity(velocity *edge.Velocity)
	SetArmed(armed *edge.Armed)
	SetFlightMode(flightMode *edge.FlightMode)
	Get() (*edge.Telemetry, error)
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

func (t *telemetry) SetConnectionState(connectionState *edge.ConnectionState) {
	t.rwm.Lock()
	defer t.rwm.Unlock()

	t.id = connectionState.VehicleID
}

func (t *telemetry) SetPosition(position *edge.Position) {
	t.rwm.Lock()
	defer t.rwm.Unlock()

	t.latitude = position.Latitude
	t.longitude = position.Longitude
	t.altitude = position.Altitude
	t.relativeAltitude = position.RelativeAltitude
}

func (t *telemetry) SetQuaternion(quaternion *edge.Quaternion) {
	t.rwm.Lock()
	defer t.rwm.Unlock()

	t.orientationX = quaternion.X
	t.orientationY = quaternion.Y
	t.orientationZ = quaternion.Z
	t.orientationW = quaternion.W
}

func (t *telemetry) SetVelocity(velocity *edge.Velocity) {
	t.rwm.Lock()
	defer t.rwm.Unlock()

	// NEDフレームから速度の合成（GroundSpeed = √n^2+e^2）
	t.speed = math.Sqrt(velocity.North*velocity.North + velocity.East*velocity.East)
}

func (t *telemetry) SetArmed(armed *edge.Armed) {
	t.rwm.Lock()
	defer t.rwm.Unlock()

	t.armed = armed.Armed
}

func (t *telemetry) SetFlightMode(flightMode *edge.FlightMode) {
	t.rwm.Lock()
	defer t.rwm.Unlock()

	t.flightMode = flightMode.FlightMode
}

func (t *telemetry) Get() (*edge.Telemetry, error) {
	rwmLocker := t.rwm.RLocker()
	rwmLocker.Lock()
	defer rwmLocker.Unlock()

	if t.flightMode == "" {
		return nil, ErrNotPrepared
	}

	telemetry := &edge.Telemetry{
		ID: t.id,
		State: &edge.State{
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
