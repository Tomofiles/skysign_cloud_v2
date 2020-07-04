package telemetry

import (
	"edge/pkg/edge"
	"log"
	"sync"
)

// Telemetry struct
type Telemetry interface {
	Updater(
		done <-chan struct{},
		connStateStream <-chan *edge.ConnectionState,
		positionStream <-chan *edge.Position,
		quaternionStream <-chan *edge.Quaternion,
		velosityStream <-chan *edge.Velocity,
		armedStream <-chan *edge.Armed,
		flightModeStream <-chan *edge.FlightMode) <-chan interface{}
	Get() *edge.Telemetry
}

type telemetry struct {
	id               string
	latitude         float64
	longitude        float64
	absoluteAltitude float64
	relativeAltitude float64
	speedNorth       float64
	speedEast        float64
	speedDown        float64
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
	connStateStream <-chan *edge.ConnectionState,
	positionStream <-chan *edge.Position,
	quaternionStream <-chan *edge.Quaternion,
	velosityStream <-chan *edge.Velocity,
	armedStream <-chan *edge.Armed,
	flightModeStream <-chan *edge.FlightMode) <-chan interface{} {
	updateExit := make(chan interface{})

	go func() {
		defer close(updateExit)
		for {
			select {
			case <-done:
				log.Println("telemetry updater done.")
				return
			case connState, ok := <-connStateStream:
				if !ok {
					log.Println("connStateStream close.")
					return
				}
				rwm.Lock()
				t.id = connState.VehicleID
				rwm.Unlock()
			case position, ok := <-positionStream:
				if !ok {
					log.Println("positionStream close.")
					return
				}
				rwm.Lock()
				t.latitude = position.Latitude
				t.longitude = position.Longitude
				t.absoluteAltitude = position.AbsoluteAltitude
				t.relativeAltitude = position.RelativeAltitude
				rwm.Unlock()
			case quaternion, ok := <-quaternionStream:
				if !ok {
					log.Println("quaternionStream close.")
					return
				}
				rwm.Lock()
				t.orientationX = quaternion.X
				t.orientationY = quaternion.Y
				t.orientationZ = quaternion.Z
				t.orientationW = quaternion.W
				rwm.Unlock()
			case velosity, ok := <-velosityStream:
				if !ok {
					log.Println("velosityStream close.")
					return
				}
				rwm.Lock()
				t.speedNorth = velosity.North
				t.speedEast = velosity.East
				t.speedDown = velosity.Down
				rwm.Unlock()
			case armed, ok := <-armedStream:
				if !ok {
					log.Println("armedStream close.")
					return
				}
				rwm.Lock()
				t.armed = armed.Armed
				rwm.Unlock()
			case flightMode, ok := <-flightModeStream:
				if !ok {
					log.Println("flightModeStream close.")
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
		ID:               t.id,
		Latitude:         t.latitude,
		Longitude:        t.longitude,
		AbsoluteAltitude: t.absoluteAltitude,
		RelativeAltitude: t.relativeAltitude,
		SpeedNorth:       t.speedNorth,
		SpeedEast:        t.speedEast,
		SpeedDown:        t.speedDown,
		Armed:            t.armed,
		FlightMode:       t.flightMode,
		OrientationX:     t.orientationX,
		OrientationY:     t.orientationY,
		OrientationZ:     t.orientationZ,
		OrientationW:     t.orientationW,
	}

	rwmLocker.Unlock()

	return telemetry
}
