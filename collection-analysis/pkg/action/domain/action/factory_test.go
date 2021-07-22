package action

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 新しいActionを作成する。
// Actionの内部の初期状態と、TrajectoryPointが0件であることを検証する。
func TestCreateNewAction(t *testing.T) {
	a := assert.New(t)

	action := NewInstance(DefaultID, DefaultCommunicationID, DefaultFleetID)

	a.Equal(action.GetID(), DefaultID)
	a.Equal(action.GetCommunicationID(), DefaultCommunicationID)
	a.Equal(action.GetFleetID(), DefaultFleetID)
	a.Equal(action.isCompleted, Active)
	a.Equal(action.trajectory.numberOfPoints, 0)
	a.Len(action.trajectory.trajectoryPoints, 0)
}

// Actionを構成オブジェクトから組み立て直し、
// 内部状態を検証する。
func TestActionAssembleFromComponent(t *testing.T) {
	a := assert.New(t)

	action := AssembleFrom(
		&componentMock{
			ID:              string(DefaultID),
			CommunicationID: string(DefaultCommunicationID),
			FleetID:         string(DefaultFleetID),
			IsCompleted:     Completed,
			TrajectoryPoints: []TrajectoryPointComponent{
				&trajectoryPointComponentMock{
					trajectoryPoint: DefaultTrajectoryPoint,
				},
				&trajectoryPointComponentMock{
					trajectoryPoint: DefaultTrajectoryPoint,
				},
				&trajectoryPointComponentMock{
					trajectoryPoint: DefaultTrajectoryPoint,
				},
			},
		},
	)

	a.Equal(action.GetID(), DefaultID)
	a.Equal(action.GetCommunicationID(), DefaultCommunicationID)
	a.Equal(action.GetFleetID(), DefaultFleetID)
	a.Equal(action.isCompleted, Completed)
	a.Equal(action.trajectory.numberOfPoints, 3)
	a.Len(action.trajectory.trajectoryPoints, 3)
	a.Equal(action.trajectory.trajectoryPoints[0], DefaultTrajectoryPoint)
	a.Equal(action.trajectory.trajectoryPoints[1], DefaultTrajectoryPoint)
	a.Equal(action.trajectory.trajectoryPoints[2], DefaultTrajectoryPoint)
}

// Actionを構成オブジェクトに分解し、
// 内部状態を検証する。
func TestTakeApartAction(t *testing.T) {
	a := assert.New(t)

	action := &Action{
		id:              DefaultID,
		communicationID: DefaultCommunicationID,
		fleetID:         DefaultFleetID,
		isCompleted:     Completed,
		trajectory: Trajectory{
			numberOfPoints: 3,
			trajectoryPoints: []trajectoryPoint{
				DefaultTrajectoryPoint,
				DefaultTrajectoryPoint,
				DefaultTrajectoryPoint,
			},
		},
	}

	var retComp ActionComponent
	var retTPComps []TrajectoryPointComponent
	TakeApart(
		action,
		func(comp ActionComponent) {
			retComp = comp
		},
		func(comp TrajectoryPointComponent) {
			retTPComps = append(retTPComps, comp)
		},
	)

	a.Equal(retComp.GetID(), string(DefaultID))
	a.Equal(retComp.GetCommunicationID(), string(DefaultCommunicationID))
	a.Equal(retComp.GetFleetID(), string(DefaultFleetID))
	a.Equal(retComp.GetIsCompleted(), Completed)
	a.Len(retTPComps, 3)
	a.EqualValues(retTPComps[0], &trajectoryPointComponentMock{trajectoryPoint: DefaultTrajectoryPoint})
	a.EqualValues(retTPComps[1], &trajectoryPointComponentMock{trajectoryPoint: DefaultTrajectoryPoint})
	a.EqualValues(retTPComps[2], &trajectoryPointComponentMock{trajectoryPoint: DefaultTrajectoryPoint})
}
