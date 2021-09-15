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
	rwm               sync.RWMutex
	id                string
	latitudeDegree    float64
	longitudeDegree   float64
	altitudeM         float64
	relativeAltitudeM float64
	speedMS           float64
	armed             bool
	flightMode        string
	orientationX      float64
	orientationY      float64
	orientationZ      float64
	orientationW      float64
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

	t.latitudeDegree = position.LatitudeDegree
	t.longitudeDegree = position.LongitudeDegree
	t.altitudeM = position.AltitudeM
	t.relativeAltitudeM = position.RelativeAltitudeM
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
	t.speedMS = math.Sqrt(velocity.NorthMS*velocity.NorthMS + velocity.EastMS*velocity.EastMS)
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
			LatitudeDegree:    t.latitudeDegree,
			LongitudeDegree:   t.longitudeDegree,
			AltitudeM:         t.altitudeM,
			RelativeAltitudeM: t.relativeAltitudeM,
			SpeedMS:           t.speedMS,
			Armed:             t.armed,
			FlightMode:        t.flightMode,
			OrientationX:      t.orientationX,
			OrientationY:      t.orientationY,
			OrientationZ:      t.orientationZ,
			OrientationW:      t.orientationW,
		},
	}

	return telemetry, nil
}
