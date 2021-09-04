package telemetry

import (
	"edge/pkg/edge"
	"edge/pkg/edge/common"
)

func Updater(
	done <-chan struct{},
	support common.Support,
	telemetry Telemetry,
	connectionStateStream <-chan *edge.ConnectionState,
	positionStream <-chan *edge.Position,
	quaternionStream <-chan *edge.Quaternion,
	velocityStream <-chan *edge.Velocity,
	armedStream <-chan *edge.Armed,
	flightModeStream <-chan *edge.FlightMode,
) <-chan struct{} {
	updateExit := make(chan struct{})

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
				telemetry.SetConnectionState(connectionState)
			case position, ok := <-positionStream:
				if !ok {
					support.NotifyInfo("positionStream close")
					return
				}
				telemetry.SetPosition(position)
			case quaternion, ok := <-quaternionStream:
				if !ok {
					support.NotifyInfo("quaternionStream close")
					return
				}
				telemetry.SetQuaternion(quaternion)
			case velocity, ok := <-velocityStream:
				if !ok {
					support.NotifyInfo("velocityStream close")
					return
				}
				telemetry.SetVelocity(velocity)
			case armed, ok := <-armedStream:
				if !ok {
					support.NotifyInfo("armedStream close")
					return
				}
				telemetry.SetArmed(armed)
			case flightMode, ok := <-flightModeStream:
				if !ok {
					support.NotifyInfo("flightModeStream close")
					return
				}
				telemetry.SetFlightMode(flightMode)
			}
		}
	}()

	return updateExit
}
