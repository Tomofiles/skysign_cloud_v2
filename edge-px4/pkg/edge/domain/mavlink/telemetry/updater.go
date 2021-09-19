package telemetry

import (
	"context"

	"github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/edge/domain/common"
	"github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/edge/domain/model"
)

// ConnectionStateUpdater .
func ConnectionStateUpdater(
	ctx context.Context,
	support common.Support,
	telemetry model.Telemetry,
	connectionStateStream <-chan *model.ConnectionState,
) <-chan struct{} {
	updaterExit := make(chan struct{})

	go func() {
		defer close(updaterExit)
		for {
			select {
			case <-ctx.Done():
				support.NotifyInfo("telemetry CONNECTION_STATE done")
				return
			case connectionState, ok := <-connectionStateStream:
				if !ok {
					support.NotifyInfo("telemetry CONNECTION_STATE close")
					return
				}
				telemetry.SetConnectionState(connectionState)
			}
		}
	}()

	return updaterExit
}

// PositionUpdater .
func PositionUpdater(
	ctx context.Context,
	support common.Support,
	telemetry model.Telemetry,
	positionStream <-chan *model.Position,
) <-chan struct{} {
	updaterExit := make(chan struct{})

	go func() {
		defer close(updaterExit)
		for {
			select {
			case <-ctx.Done():
				support.NotifyInfo("telemetry POSITION done")
				return
			case position, ok := <-positionStream:
				if !ok {
					support.NotifyInfo("telemetry POSITION close")
					return
				}
				telemetry.SetPosition(position)
			}
		}
	}()

	return updaterExit
}

// QuaternionUpdater .
func QuaternionUpdater(
	ctx context.Context,
	support common.Support,
	telemetry model.Telemetry,
	quaternionStream <-chan *model.Quaternion,
) <-chan struct{} {
	updaterExit := make(chan struct{})

	go func() {
		defer close(updaterExit)
		for {
			select {
			case <-ctx.Done():
				support.NotifyInfo("telemetry QUATERNION done")
				return
			case quaternion, ok := <-quaternionStream:
				if !ok {
					support.NotifyInfo("telemetry QUATERNION close")
					return
				}
				telemetry.SetQuaternion(quaternion)
			}
		}
	}()

	return updaterExit
}

// VelocityUpdater .
func VelocityUpdater(
	ctx context.Context,
	support common.Support,
	telemetry model.Telemetry,
	velocityStream <-chan *model.Velocity,
) <-chan struct{} {
	updaterExit := make(chan struct{})

	go func() {
		defer close(updaterExit)
		for {
			select {
			case <-ctx.Done():
				support.NotifyInfo("telemetry VELOCITY done")
				return
			case velocity, ok := <-velocityStream:
				if !ok {
					support.NotifyInfo("telemetry VELOCITY close")
					return
				}
				telemetry.SetVelocity(velocity)
			}
		}
	}()

	return updaterExit
}

// ArmedUpdater .
func ArmedUpdater(
	ctx context.Context,
	support common.Support,
	telemetry model.Telemetry,
	armedStream <-chan *model.Armed,
) <-chan struct{} {
	updaterExit := make(chan struct{})

	go func() {
		defer close(updaterExit)
		for {
			select {
			case <-ctx.Done():
				support.NotifyInfo("telemetry ARMED done")
				return
			case armed, ok := <-armedStream:
				if !ok {
					support.NotifyInfo("telemetry ARMED close")
					return
				}
				telemetry.SetArmed(armed)
			}
		}
	}()

	return updaterExit
}

// FlightModeUpdater .
func FlightModeUpdater(
	ctx context.Context,
	support common.Support,
	telemetry model.Telemetry,
	flightModeStream <-chan *model.FlightMode,
) <-chan struct{} {
	updaterExit := make(chan struct{})

	go func() {
		defer close(updaterExit)
		for {
			select {
			case <-ctx.Done():
				support.NotifyInfo("telemetry FLIGHT_MODE done")
				return
			case flightMode, ok := <-flightModeStream:
				if !ok {
					support.NotifyInfo("telemetry FLIGHT_MODE close")
					return
				}
				telemetry.SetFlightMode(flightMode)
			}
		}
	}()

	return updaterExit
}
