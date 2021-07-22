package service

import (
	act "collection-analysis/pkg/action/domain/action"
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

	command := &createCommandMock{
		VehicleId:       string(DefaultActionID),
		CommunicationId: string(DefaultActionCommunicationID),
		FleetId:         string(DefaultActionFleetID),
	}
	ret := service.CreateAction(command)

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

	command := &createCommandMock{
		VehicleId:       string(DefaultActionID),
		CommunicationId: string(DefaultActionCommunicationID),
		FleetId:         string(DefaultActionFleetID),
	}
	ret := service.createActionOperation(nil, command)

	trajectoryPointComps := []act.TrajectoryPointComponent{}
	actionComp := actionComponentMock{
		ID:               string(DefaultActionID),
		CommunicationID:  string(DefaultActionCommunicationID),
		FleetID:          string(DefaultActionFleetID),
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
		FleetID:          string(DefaultActionFleetID),
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

	command := &actionIDCommandMock{
		ID: string(DefaultActionID),
	}
	var resCall bool
	ret := service.GetTrajectory(
		command,
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
		FleetID:          string(DefaultActionFleetID),
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

	command := &actionIDCommandMock{
		ID: string(DefaultActionID),
	}
	var snapshot act.TelemetrySnapshot
	ret := service.getTrajectoryOperation(
		nil,
		command,
		func(s act.TelemetrySnapshot) {
			snapshot = s
		},
	)

	a.Nil(ret)
	a.Equal(snapshot, DefaultTelemetrySnapshot)
}
