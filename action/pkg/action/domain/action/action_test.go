package action

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Telemetryを更新する。
// Armed状態のTelemetryを受信した場合、Trajectoryが伸長されることを検証する。
func TestPushTelemetry_ReceiveArmedTelemetry(t *testing.T) {
	a := assert.New(t)

	action := &Action{
		id:              DefaultID,
		communicationID: DefaultCommunicationID,
		flightplanID:    DefaultFlightplanID,
		isCompleted:     Active,
		trajectory: Trajectory{
			numberOfPoints:   0,
			trajectoryPoints: []trajectoryPoint{},
		},
	}

	ret := action.PushTelemetry(
		TelemetrySnapshot{
			Armed: true,
		},
	)

	a.Equal(action.trajectory.numberOfPoints, 1)
	a.Len(action.trajectory.trajectoryPoints, 1)
	a.True(action.trajectory.trajectoryPoints[0].Armed)
	a.Nil(ret)
}

// Telemetryを更新する。
// ActionがComplete状態の場合、エラーが発生して、
// 何も更新されないことを検証する。
func TestCannotChangeErrorWhenPushTelemetry_CompletedAction(t *testing.T) {
	a := assert.New(t)

	action := &Action{
		id:              DefaultID,
		communicationID: DefaultCommunicationID,
		flightplanID:    DefaultFlightplanID,
		isCompleted:     Completed,
		trajectory: Trajectory{
			numberOfPoints:   0,
			trajectoryPoints: []trajectoryPoint{},
		},
	}

	ret := action.PushTelemetry(
		TelemetrySnapshot{
			Armed: true,
		},
	)

	a.Equal(action.trajectory.numberOfPoints, 0)
	a.Len(action.trajectory.trajectoryPoints, 0)
	a.Equal(ret, ErrCannotChange)
}

// Telemetryを更新する。
// Disarmed状態のTelemetryを受信した場合、何も更新されないことを検証する。
func TestNoUpdateWhenPushTelemetry_ReceiveDisarmedTelemetry(t *testing.T) {
	a := assert.New(t)

	action := &Action{
		id:              DefaultID,
		communicationID: DefaultCommunicationID,
		flightplanID:    DefaultFlightplanID,
		isCompleted:     Active,
		trajectory: Trajectory{
			numberOfPoints:   0,
			trajectoryPoints: []trajectoryPoint{},
		},
	}

	ret := action.PushTelemetry(
		TelemetrySnapshot{
			Armed: false,
		},
	)

	a.Equal(action.trajectory.numberOfPoints, 0)
	a.Len(action.trajectory.trajectoryPoints, 0)
	a.Nil(ret)
}

// Complete状態に更新する。
// ActionがActive状態の場合、正常に更新されることを検証する。
func TestComplete(t *testing.T) {
	a := assert.New(t)

	action := &Action{
		id:              DefaultID,
		communicationID: DefaultCommunicationID,
		flightplanID:    DefaultFlightplanID,
		isCompleted:     Active,
		trajectory: Trajectory{
			numberOfPoints:   0,
			trajectoryPoints: []trajectoryPoint{},
		},
	}

	ret := action.Complete()

	a.True(action.isCompleted)
	a.Nil(ret)
}

// Complete状態に更新する。
// ActionがComplete状態の場合、エラーが発生して、
// 何も更新されないことを検証する。
func TestCannotChangeErrorWhenComplete(t *testing.T) {
	a := assert.New(t)

	action := &Action{
		id:              DefaultID,
		communicationID: DefaultCommunicationID,
		flightplanID:    DefaultFlightplanID,
		isCompleted:     Completed,
		trajectory: Trajectory{
			numberOfPoints:   0,
			trajectoryPoints: []trajectoryPoint{},
		},
	}

	ret := action.Complete()

	a.Equal(ret, ErrCannotChange)
}

// TrajectoryのSnapshotを取得する。
// TrajectoryPointが0件の場合、Snapshotのダブルディスパッチが
// コールされないことを検証する。
func TestProvideTrajectoryInterest_NoTrajectory(t *testing.T) {
	a := assert.New(t)

	action := &Action{
		id:              DefaultID,
		communicationID: DefaultCommunicationID,
		flightplanID:    DefaultFlightplanID,
		isCompleted:     Active,
		trajectory: Trajectory{
			numberOfPoints:   0,
			trajectoryPoints: []trajectoryPoint{},
		},
	}

	isCall := false
	action.ProvideTrajectoryInterest(
		func(snapshot TelemetrySnapshot) {
			isCall = true
		},
	)

	a.False(isCall)
}

// TrajectoryのSnapshotを取得する。
// TrajectoryPointが複数件の場合、Snapshotのダブルディスパッチが
// 当該回数コールされることを検証する。
func TestProvideTrajectoryInterest_MultipleTrajectory(t *testing.T) {
	a := assert.New(t)

	action := &Action{
		id:              DefaultID,
		communicationID: DefaultCommunicationID,
		flightplanID:    DefaultFlightplanID,
		isCompleted:     Active,
		trajectory: Trajectory{
			numberOfPoints: 0,
			trajectoryPoints: []trajectoryPoint{
				{Armed: true},
				{Armed: false},
				{Armed: true},
				{Armed: true},
			},
		},
	}

	armeds := []bool{}
	action.ProvideTrajectoryInterest(
		func(snapshot TelemetrySnapshot) {
			armeds = append(armeds, snapshot.Armed)
		},
	)

	a.Equal(armeds, []bool{true, false, true, true})
}
