package communication

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Telemetryを一つ新しく作成し、初期状態を検証する。
func TestCreateNewTelemetry(t *testing.T) {
	a := assert.New(t)

	telemetry := NewTelemetry()

	a.Equal(telemetry.latitudeDegree, 0.0)
	a.Equal(telemetry.longitudeDegree, 0.0)
	a.Equal(telemetry.altitudeM, 0.0)
	a.Equal(telemetry.relativeAltitudeM, 0.0)
	a.Equal(telemetry.speedMS, 0.0)
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
		LatitudeDegree:    1.0,
		LongitudeDegree:   2.0,
		AltitudeM:         3.0,
		RelativeAltitudeM: 4.0,
		SpeedMS:           5.0,
		Armed:             Armed,
		FlightMode:        "NONE",
		X:                 6.0,
		Y:                 7.0,
		Z:                 8.0,
		W:                 9.0,
	}
	telemetry := NewTelemetryBySnapshot(snapshot)

	a.Equal(telemetry.latitudeDegree, 1.0)
	a.Equal(telemetry.longitudeDegree, 2.0)
	a.Equal(telemetry.altitudeM, 3.0)
	a.Equal(telemetry.relativeAltitudeM, 4.0)
	a.Equal(telemetry.speedMS, 5.0)
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
		latitudeDegree:    1.0,
		longitudeDegree:   2.0,
		altitudeM:         3.0,
		relativeAltitudeM: 4.0,
		speedMS:           5.0,
		armed:             Armed,
		flightMode:        "NONE",
		x:                 6.0,
		y:                 7.0,
		z:                 8.0,
		w:                 9.0,
	}

	snapshot := telemetry.GetSnapshot()

	a.Equal(snapshot.LatitudeDegree, 1.0)
	a.Equal(snapshot.LongitudeDegree, 2.0)
	a.Equal(snapshot.AltitudeM, 3.0)
	a.Equal(snapshot.RelativeAltitudeM, 4.0)
	a.Equal(snapshot.SpeedMS, 5.0)
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
