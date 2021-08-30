package communication

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Telemetryを一つ新しく作成し、初期状態を検証する。
func TestCreateNewTelemetry(t *testing.T) {
	a := assert.New(t)

	telemetry := NewTelemetry()

	a.Equal(telemetry.latitude, 0.0)
	a.Equal(telemetry.longitude, 0.0)
	a.Equal(telemetry.altitude, 0.0)
	a.Equal(telemetry.relativeAltitude, 0.0)
	a.Equal(telemetry.speed, 0.0)
	a.Equal(telemetry.armed, Disarmed)
	a.Equal(telemetry.flightMode, "NONE")
	a.Equal(telemetry.x, 0.0)
	a.Equal(telemetry.y, 0.0)
	a.Equal(telemetry.z, 0.0)
	a.Equal(telemetry.w, 0.0)
}

// Snapshotを元にTelemetryを作成し、初期状態を検証する。
func TestCreateNewTelemetryFromSnapshot(t *testing.T) {
	a := assert.New(t)

	snapshot := TelemetrySnapshot{
		Latitude:         1.0,
		Longitude:        2.0,
		Altitude:         3.0,
		RelativeAltitude: 4.0,
		Speed:            5.0,
		Armed:            Armed,
		FlightMode:       "NONE",
		X:                6.0,
		Y:                7.0,
		Z:                8.0,
		W:                9.0,
	}
	telemetry := NewTelemetryBySnapshot(snapshot)

	a.Equal(telemetry.latitude, 1.0)
	a.Equal(telemetry.longitude, 2.0)
	a.Equal(telemetry.altitude, 3.0)
	a.Equal(telemetry.relativeAltitude, 4.0)
	a.Equal(telemetry.speed, 5.0)
	a.Equal(telemetry.armed, Armed)
	a.Equal(telemetry.flightMode, "NONE")
	a.Equal(telemetry.x, 6.0)
	a.Equal(telemetry.y, 7.0)
	a.Equal(telemetry.z, 8.0)
	a.Equal(telemetry.w, 9.0)
}

// TelemetryからSnapshotを取得し、内部状態を検証する。
func TestGetSnapshotFromTelemetry(t *testing.T) {
	a := assert.New(t)

	telemetry := &Telemetry{
		latitude:         1.0,
		longitude:        2.0,
		altitude:         3.0,
		relativeAltitude: 4.0,
		speed:            5.0,
		armed:            Armed,
		flightMode:       "NONE",
		x:                6.0,
		y:                7.0,
		z:                8.0,
		w:                9.0,
	}

	snapshot := telemetry.GetSnapshot()

	a.Equal(snapshot.Latitude, 1.0)
	a.Equal(snapshot.Longitude, 2.0)
	a.Equal(snapshot.Altitude, 3.0)
	a.Equal(snapshot.RelativeAltitude, 4.0)
	a.Equal(snapshot.Speed, 5.0)
	a.Equal(snapshot.Armed, Armed)
	a.Equal(snapshot.FlightMode, "NONE")
	a.Equal(snapshot.X, 6.0)
	a.Equal(snapshot.Y, 7.0)
	a.Equal(snapshot.Z, 8.0)
	a.Equal(snapshot.W, 9.0)
}

// TelemetryのIsDisarmedを検証する。
func TestIsDisarmedOfTelemetry(t *testing.T) {
	a := assert.New(t)

	armedTelemetry := &Telemetry{
		armed: Armed,
	}
	disarmedTelemetry := &Telemetry{
		armed: Disarmed,
	}

	a.False(armedTelemetry.IsDisarmed())
	a.True(disarmedTelemetry.IsDisarmed())
}
