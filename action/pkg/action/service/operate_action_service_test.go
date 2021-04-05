package service

import (
	act "action/pkg/action/domain/action"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCompleteActionTransaction(t *testing.T) {
	a := assert.New(t)

	trajectoryPointComps := []act.TrajectoryPointComponent{}
	actionComp := actionComponentMock{
		ID:               string(DefaultActionID),
		CommunicationID:  string(DefaultActionCommunicationID),
		FlightplanID:     string(DefaultActionFlightplanID),
		IsCompleted:      act.Active,
		TrajectoryPoints: trajectoryPointComps,
	}
	action := act.AssembleFrom(
		&actionComp,
	)
	actions := []*act.Action{action}

	repo := &actionRepositoryMock{}
	repo.On("GetAllActiveByFlightplanID", DefaultActionFlightplanID).Return(actions, nil)
	repo.On("Save", mock.Anything).Return(nil)
	txm := &txManagerMock{}

	service := &operateActionService{
		repo: repo,
		txm:  txm,
	}

	req := &flightplanIDRequestMock{
		FlightplanID: string(DefaultActionFlightplanID),
	}
	ret := service.CompleteAction(req)

	a.Nil(ret)
	a.Nil(txm.isOpe)
	a.Nil(txm.isEH)
}

func TestCompleteActionOperation(t *testing.T) {
	a := assert.New(t)

	trajectoryPointComps := []act.TrajectoryPointComponent{}
	actionComp := actionComponentMock{
		ID:               string(DefaultActionID),
		CommunicationID:  string(DefaultActionCommunicationID),
		FlightplanID:     string(DefaultActionFlightplanID),
		IsCompleted:      act.Active,
		TrajectoryPoints: trajectoryPointComps,
	}
	action := act.AssembleFrom(
		&actionComp,
	)
	actions := []*act.Action{action}

	repo := &actionRepositoryMock{}
	repo.On("GetAllActiveByFlightplanID", DefaultActionFlightplanID).Return(actions, nil)
	repo.On("Save", mock.Anything).Return(nil)

	service := &operateActionService{
		repo: repo,
		txm:  nil,
	}

	req := &flightplanIDRequestMock{
		FlightplanID: string(DefaultActionFlightplanID),
	}
	ret := service.completeActionOperation(
		nil,
		req,
	)

	expectTrajectoryPointComps := []act.TrajectoryPointComponent{}
	expectActionComp := actionComponentMock{
		ID:               string(DefaultActionID),
		CommunicationID:  string(DefaultActionCommunicationID),
		FlightplanID:     string(DefaultActionFlightplanID),
		IsCompleted:      act.Completed,
		TrajectoryPoints: expectTrajectoryPointComps,
	}
	expectAct := act.AssembleFrom(
		&expectActionComp,
	)

	a.Nil(ret)
	a.Equal(repo.action, expectAct)
}

func TestNoActionWhenCompleteActionOperation(t *testing.T) {
	a := assert.New(t)

	actions := []*act.Action{}

	repo := &actionRepositoryMock{}
	repo.On("GetAllActiveByFlightplanID", DefaultActionFlightplanID).Return(actions, nil)

	service := &operateActionService{
		repo: repo,
		txm:  nil,
	}

	req := &flightplanIDRequestMock{
		FlightplanID: string(DefaultActionFlightplanID),
	}
	ret := service.completeActionOperation(
		nil,
		req,
	)

	a.Nil(ret)
	a.Nil(repo.action)
}

func TestPushTelemetryTransaction(t *testing.T) {
	a := assert.New(t)

	trajectoryPointComps := []act.TrajectoryPointComponent{}
	actionComp := actionComponentMock{
		ID:               string(DefaultActionID),
		CommunicationID:  string(DefaultActionCommunicationID),
		FlightplanID:     string(DefaultActionFlightplanID),
		IsCompleted:      act.Active,
		TrajectoryPoints: trajectoryPointComps,
	}
	action := act.AssembleFrom(
		&actionComp,
	)

	repo := &actionRepositoryMock{}
	repo.On("GetActiveByCommunicationID", DefaultActionCommunicationID).Return(action, nil)
	repo.On("Save", mock.Anything).Return(nil)
	txm := &txManagerMock{}

	service := &operateActionService{
		repo: repo,
		txm:  txm,
	}

	req := &telemetryRequestMock{
		CommunicationID:  string(DefaultActionCommunicationID),
		Latitude:         1.0,
		Longitude:        2.0,
		Altitude:         3.0,
		RelativeAltitude: 4.0,
		Speed:            5.0,
		Armed:            true,
		FlightMode:       "state",
		OrientationX:     6.0,
		OrientationY:     7.0,
		OrientationZ:     8.0,
		OrientationW:     9.0,
	}
	ret := service.PushTelemetry(req)

	a.Nil(ret)
	a.Nil(txm.isOpe)
	a.Nil(txm.isEH)
}

func TestPushTelemetryOperation(t *testing.T) {
	a := assert.New(t)

	trajectoryPointComps := []act.TrajectoryPointComponent{}
	actionComp := actionComponentMock{
		ID:               string(DefaultActionID),
		CommunicationID:  string(DefaultActionCommunicationID),
		FlightplanID:     string(DefaultActionFlightplanID),
		IsCompleted:      act.Active,
		TrajectoryPoints: trajectoryPointComps,
	}
	action := act.AssembleFrom(
		&actionComp,
	)

	repo := &actionRepositoryMock{}
	repo.On("GetActiveByCommunicationID", DefaultActionCommunicationID).Return(action, nil)
	repo.On("Save", mock.Anything).Return(nil)

	service := &operateActionService{
		repo: repo,
		txm:  nil,
	}

	req := &telemetryRequestMock{
		CommunicationID:  string(DefaultActionCommunicationID),
		Latitude:         1.0,
		Longitude:        2.0,
		Altitude:         3.0,
		RelativeAltitude: 4.0,
		Speed:            5.0,
		Armed:            true,
		FlightMode:       "state",
		OrientationX:     6.0,
		OrientationY:     7.0,
		OrientationZ:     8.0,
		OrientationW:     9.0,
	}
	ret := service.pushTelemetryOperation(
		nil,
		req,
	)

	expectTrajectoryPointComps := []act.TrajectoryPointComponent{
		&trajectoryPointComponentMock{
			1,
			1.0,
			2.0,
			3.0,
			4.0,
			5.0,
			true,
			"state",
			6.0,
			7.0,
			8.0,
			9.0,
		},
	}
	expectActionComp := actionComponentMock{
		ID:               string(DefaultActionID),
		CommunicationID:  string(DefaultActionCommunicationID),
		FlightplanID:     string(DefaultActionFlightplanID),
		IsCompleted:      act.Active,
		TrajectoryPoints: expectTrajectoryPointComps,
	}
	expectAct := act.AssembleFrom(
		&expectActionComp,
	)

	a.Nil(ret)
	a.Equal(repo.action, expectAct)
}
