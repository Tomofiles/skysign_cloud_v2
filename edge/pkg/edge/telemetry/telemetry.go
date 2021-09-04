package telemetry

import (
	"edge/pkg/edge"
	"edge/pkg/edge/common"
	"math"
	"sync"
)

// Telemetry struct
type Telemetry interface {
	Updater(
		done <-chan struct{},
		support common.Support,
		connectionStateStream <-chan *edge.ConnectionState,
		positionStream <-chan *edge.Position,
		quaternionStream <-chan *edge.Quaternion,
		velocityStream <-chan *edge.Velocity,
		armedStream <-chan *edge.Armed,
		flightModeStream <-chan *edge.FlightMode) <-chan interface{}
	Get() *edge.Telemetry
}

type telemetry struct {
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

var rwm sync.RWMutex

// NewTelemetry .
func NewTelemetry() Telemetry {
	return &telemetry{}
}

func (t *telemetry) Updater(
	done <-chan struct{},
	support common.Support,
	connectionStateStream <-chan *edge.ConnectionState,
	positionStream <-chan *edge.Position,
	quaternionStream <-chan *edge.Quaternion,
	velocityStream <-chan *edge.Velocity,
	armedStream <-chan *edge.Armed,
	flightModeStream <-chan *edge.FlightMode,
) <-chan interface{} {
	updateExit := make(chan interface{})

	go func() {
		defer close(updateExit)
		for {
			select {
			case <-done:
				support.NotifyInfo("telemetry updater done")
				return
			case connectionState, ok := <-connectionStateStream:
				if !ok {
					support.NotifyInfo("connectionStateStream close")
					return
				}
				rwm.Lock()
				t.id = connectionState.VehicleID
				rwm.Unlock()
			case position, ok := <-positionStream:
				if !ok {
					support.NotifyInfo("positionStream close")
					return
				}
				rwm.Lock()
				t.latitude = position.Latitude
				t.longitude = position.Longitude
				t.altitude = position.Altitude
				t.relativeAltitude = position.RelativeAltitude
				rwm.Unlock()
			case quaternion, ok := <-quaternionStream:
				if !ok {
					support.NotifyInfo("quaternionStream close")
					return
				}
				rwm.Lock()
				t.orientationX = quaternion.X
				t.orientationY = quaternion.Y
				t.orientationZ = quaternion.Z
				t.orientationW = quaternion.W
				rwm.Unlock()
			case velocity, ok := <-velocityStream:
				if !ok {
					support.NotifyInfo("velocityStream close")
					return
				}
				rwm.Lock()
				// NEDフレームから速度の合成（GroundSpeed = √n^2+e^2）
				t.speed = math.Sqrt(velocity.North*velocity.North + velocity.East*velocity.East)
				rwm.Unlock()
			case armed, ok := <-armedStream:
				if !ok {
					support.NotifyInfo("armedStream close")
					return
				}
				rwm.Lock()
				t.armed = armed.Armed
				rwm.Unlock()
			case flightMode, ok := <-flightModeStream:
				if !ok {
					support.NotifyInfo("flightModeStream close")
					return
				}
				rwm.Lock()
				t.flightMode = flightMode.FlightMode
				rwm.Unlock()
			}
		}
	}()

	return updateExit
}

func (t *telemetry) Get() *edge.Telemetry {
	rwmLocker := rwm.RLocker()
	rwmLocker.Lock()

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

	rwmLocker.Unlock()

	return telemetry
}
