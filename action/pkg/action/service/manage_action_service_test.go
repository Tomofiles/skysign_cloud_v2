package service

import (
	act "action/pkg/action/domain/action"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateActionTransaction(t *testing.T) {
	a := assert.New(t)

	repo := &actionRepositoryMock{}
	repo.On("Save", mock.Anything).Return(nil)
	txm := &txManagerMock{}

	service := &manageActionService{
		repo: repo,
		txm:  txm,
	}

	req := &createRequestMock{
		VehicleId:       string(DefaultActionID),
		CommunicationId: string(DefaultActionCommunicationID),
		FlightplanId:    string(DefaultActionFlightplanID),
	}
	ret := service.CreateAction(req)

	a.Nil(ret)
	a.Nil(txm.isOpe)
	a.Nil(txm.isEH)
}

func TestCreateActionOperation(t *testing.T) {
	a := assert.New(t)

	repo := &actionRepositoryMock{}
	repo.On("Save", mock.Anything).Return(nil)

	service := &manageActionService{
		repo: repo,
		txm:  nil,
	}

	req := &createRequestMock{
		VehicleId:       string(DefaultActionID),
		CommunicationId: string(DefaultActionCommunicationID),
		FlightplanId:    string(DefaultActionFlightplanID),
	}
	ret := service.createActionOperation(nil, req)

	trajectoryPointComps := []act.TrajectoryPointComponent{}
	actionComp := actionComponentMock{
		ID:               string(DefaultActionID),
		CommunicationID:  string(DefaultActionCommunicationID),
		FlightplanID:     string(DefaultActionFlightplanID),
		IsCompleted:      act.Active,
		TrajectoryPoints: trajectoryPointComps,
	}
	expectAct := act.AssembleFrom(
		&actionComp,
	)

	a.Nil(ret)
	a.Equal(repo.action, expectAct)
}

func TestGetTrajectoryTransaction(t *testing.T) {
	a := assert.New(t)

	trajectoryPointComps := []act.TrajectoryPointComponent{
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
	repo.On("GetByID", DefaultActionID).Return(action, nil)
	txm := &txManagerMock{}

	service := &manageActionService{
		repo: repo,
		txm:  txm,
	}

	req := &idRequestMock{
		VehicleId: string(DefaultActionID),
	}
	var resCall bool
	ret := service.GetTrajectory(
		req,
		func(s act.TelemetrySnapshot) {
			resCall = true
		},
	)

	a.Nil(ret)
	a.True(resCall)
}

func TestGetTrajectoryOperation(t *testing.T) {
	a := assert.New(t)

	trajectoryPointComps := []act.TrajectoryPointComponent{
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
	repo.On("GetByID", DefaultActionID).Return(action, nil)

	service := &manageActionService{
		repo: repo,
		txm:  nil,
	}

	req := &idRequestMock{
		VehicleId: string(DefaultActionID),
	}
	var snapshot act.TelemetrySnapshot
	ret := service.getTrajectoryOperation(
		nil,
		req,
		func(s act.TelemetrySnapshot) {
			snapshot = s
		},
	)

	a.Nil(ret)
	a.Equal(snapshot, DefaultTelemetrySnapshot)
}
