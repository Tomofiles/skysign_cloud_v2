package service

import (
	"testing"
	v "vehicle/pkg/vehicle/domain/vehicle"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetVehicleTransaction(t *testing.T) {
	a := assert.New(t)

	vehicle := v.AssembleFrom(
		nil,
		&vehicleComponentMock{
			ID:              string(DefaultVehicleID),
			Name:            DefaultVehicleName,
			CommunicationID: string(DefaultVehicleCommunicationID),
			Version:         string(DefaultVehicleVersion),
		},
	)

	repo := &vehicleRepositoryMock{}
	txm := &txManagerMock{}

	repo.On("GetByID", DefaultVehicleID).Return(vehicle, nil)

	service := &manageVehicleService{
		gen:  nil,
		repo: repo,
		txm:  txm,
		psm:  nil,
	}

	req := &vehicleIDRequestMock{
		ID: string(DefaultVehicleID),
	}
	var resCall bool
	ret := service.GetVehicle(
		req,
		func(id, name, communicationID string) {
			resCall = true
		},
	)

	a.Nil(ret)
	a.True(resCall)
	a.Nil(txm.isOpe)
}

func TestGetVehicleOperation(t *testing.T) {
	a := assert.New(t)

	vehicle := v.AssembleFrom(
		nil,
		&vehicleComponentMock{
			ID:              string(DefaultVehicleID),
			Name:            DefaultVehicleName,
			CommunicationID: string(DefaultVehicleCommunicationID),
			Version:         string(DefaultVehicleVersion),
		},
	)

	repo := &vehicleRepositoryMock{}
	repo.On("GetByID", DefaultVehicleID).Return(vehicle, nil)

	service := &manageVehicleService{
		gen:  nil,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	req := &vehicleIDRequestMock{
		ID: string(DefaultVehicleID),
	}
	var resID, resName, resCommunicationID string
	ret := service.getVehicleOperation(
		nil,
		req,
		func(id, name, communicationID string) {
			resID = id
			resName = name
			resCommunicationID = communicationID
		},
	)

	a.Nil(ret)
	a.Equal(resID, string(DefaultVehicleID))
	a.Equal(resName, DefaultVehicleName)
	a.Equal(resCommunicationID, string(DefaultVehicleCommunicationID))
}

func TestListVehiclesTransaction(t *testing.T) {
	a := assert.New(t)

	vehicles := []*v.Vehicle{
		v.AssembleFrom(
			nil,
			&vehicleComponentMock{
				ID:              string(DefaultVehicleID),
				Name:            DefaultVehicleName,
				CommunicationID: string(DefaultVehicleCommunicationID),
				Version:         string(DefaultVehicleVersion),
			},
		),
	}

	repo := &vehicleRepositoryMock{}
	repo.On("GetAllOrigin").Return(vehicles, nil)
	txm := &txManagerMock{}

	service := &manageVehicleService{
		gen:  nil,
		repo: repo,
		txm:  txm,
		psm:  nil,
	}

	var resCall bool
	ret := service.ListVehicles(
		func(id, name, communicationID string) {
			resCall = true
		},
	)

	a.Nil(ret)
	a.True(resCall)
	a.Nil(txm.isOpe)
}

func TestListVehiclesOperation(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultVehicleID1              = string(DefaultVehicleID) + "-1"
		DefaultVehicleName1            = DefaultVehicleName + "-1"
		DefaultVehicleCommunicationID1 = string(DefaultVehicleCommunicationID) + "-1"
		DefaultVehicleVersion1         = string(DefaultVehicleVersion) + "-1"
		DefaultVehicleID2              = string(DefaultVehicleID) + "-2"
		DefaultVehicleName2            = DefaultVehicleName + "-2"
		DefaultVehicleCommunicationID2 = string(DefaultVehicleCommunicationID) + "-2"
		DefaultVehicleVersion2         = string(DefaultVehicleVersion) + "-2"
		DefaultVehicleID3              = string(DefaultVehicleID) + "-3"
		DefaultVehicleName3            = DefaultVehicleName + "-3"
		DefaultVehicleCommunicationID3 = string(DefaultVehicleCommunicationID) + "-3"
		DefaultVehicleVersion3         = string(DefaultVehicleVersion) + "-3"
	)

	vehicles := []*v.Vehicle{
		v.AssembleFrom(
			nil,
			&vehicleComponentMock{
				ID:              DefaultVehicleID1,
				Name:            DefaultVehicleName1,
				CommunicationID: DefaultVehicleCommunicationID1,
				Version:         DefaultVehicleVersion1,
			},
		),
		v.AssembleFrom(
			nil,
			&vehicleComponentMock{
				ID:              DefaultVehicleID2,
				Name:            DefaultVehicleName2,
				CommunicationID: DefaultVehicleCommunicationID2,
				Version:         DefaultVehicleVersion2,
			},
		),
		v.AssembleFrom(
			nil,
			&vehicleComponentMock{
				ID:              DefaultVehicleID3,
				Name:            DefaultVehicleName3,
				CommunicationID: DefaultVehicleCommunicationID3,
				Version:         DefaultVehicleVersion3,
			},
		),
	}

	repo := &vehicleRepositoryMock{}
	repo.On("GetAllOrigin").Return(vehicles, nil)

	service := &manageVehicleService{
		gen:  nil,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	var resID, resName, resCommunicationID []string
	ret := service.listVehiclesOperation(
		nil,
		func(id, name, communicationID string) {
			resID = append(resID, id)
			resName = append(resName, name)
			resCommunicationID = append(resCommunicationID, communicationID)
		},
	)

	a.Nil(ret)
	a.Equal(resID, []string{DefaultVehicleID1, DefaultVehicleID2, DefaultVehicleID3})
	a.Equal(resName, []string{DefaultVehicleName1, DefaultVehicleName2, DefaultVehicleName3})
	a.Equal(resCommunicationID, []string{DefaultVehicleCommunicationID1, DefaultVehicleCommunicationID2, DefaultVehicleCommunicationID3})
}

func TestCreateVehicleTransaction(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultVehicleVersion1 = DefaultVehicleVersion + "-1"
		DefaultVehicleVersion2 = DefaultVehicleVersion + "-2"
		DefaultVehicleVersion3 = DefaultVehicleVersion + "-3"
	)

	gen := &generatorMock{
		id:       DefaultVehicleID,
		versions: []v.Version{DefaultVehicleVersion1, DefaultVehicleVersion2, DefaultVehicleVersion3},
	}
	repo := &vehicleRepositoryMock{}
	txm := &txManagerMock{}
	pub := &publisherMock{}
	psm := &pubSubManagerMock{}

	var isClose bool
	close := func() error {
		isClose = true
		return nil
	}

	psm.On("GetPublisher").Return(pub, close, nil)
	repo.On("Save", mock.Anything).Return(nil)

	service := &manageVehicleService{
		gen:  gen,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}

	req := &vehicleRequestMock{
		Name:            DefaultVehicleName,
		CommunicationID: string(DefaultVehicleCommunicationID),
	}
	var resCall bool
	ret := service.CreateVehicle(
		req,
		func(id, name, communicationID string) {
			resCall = true
		},
	)

	a.Nil(ret)
	a.True(resCall)
	a.Len(pub.events, 1)
	a.True(isClose)
	a.True(pub.isFlush)
	a.Nil(txm.isOpe)
	a.Nil(txm.isEH)
}

func TestCreateVehicleOperation(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultVehicleVersion1 = DefaultVehicleVersion + "-1"
		DefaultVehicleVersion2 = DefaultVehicleVersion + "-2"
		DefaultVehicleVersion3 = DefaultVehicleVersion + "-3"
	)

	gen := &generatorMock{
		id:       DefaultVehicleID,
		versions: []v.Version{DefaultVehicleVersion1, DefaultVehicleVersion2, DefaultVehicleVersion3},
	}
	repo := &vehicleRepositoryMock{}
	repo.On("Save", mock.Anything).Return(nil)
	pub := &publisherMock{}

	service := &manageVehicleService{
		gen:  gen,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	req := &vehicleRequestMock{
		Name:            DefaultVehicleName,
		CommunicationID: string(DefaultVehicleCommunicationID),
	}
	var resID, resName, resCommunicationID string
	ret := service.createVehicleOperation(
		nil,
		pub,
		req,
		func(id, name, communicationID string) {
			resID = id
			resName = name
			resCommunicationID = communicationID
		},
	)

	expectEvent := v.CommunicationIdGaveEvent{CommunicationID: DefaultVehicleCommunicationID}

	a.Nil(ret)
	a.Equal(resID, string(DefaultVehicleID))
	a.Equal(resName, DefaultVehicleName)
	a.Equal(resCommunicationID, string(DefaultVehicleCommunicationID))
	a.Len(pub.events, 1)
	a.Equal(pub.events[0], expectEvent)
}

func TestUpdateVehicleTransaction(t *testing.T) {
	a := assert.New(t)

	var (
		AfterVehicleName            = DefaultVehicleName + "-after"
		AfterVehicleCommunicationID = DefaultVehicleCommunicationID + "-after"
		DefaultVehicleVersion1      = DefaultVehicleVersion + "-1"
		DefaultVehicleVersion2      = DefaultVehicleVersion + "-2"
	)

	gen := &generatorMock{
		id:       DefaultVehicleID,
		versions: []v.Version{DefaultVehicleVersion1, DefaultVehicleVersion2},
	}

	vehicle := v.AssembleFrom(
		gen,
		&vehicleComponentMock{
			ID:              string(DefaultVehicleID),
			Name:            DefaultVehicleName,
			CommunicationID: string(DefaultVehicleCommunicationID),
			Version:         string(DefaultVehicleVersion),
		},
	)

	repo := &vehicleRepositoryMock{}
	txm := &txManagerMock{}
	pub := &publisherMock{}
	psm := &pubSubManagerMock{}

	var isClose bool
	close := func() error {
		isClose = true
		return nil
	}

	psm.On("GetPublisher").Return(pub, close, nil)
	repo.On("GetByID", DefaultVehicleID).Return(vehicle, nil)
	repo.On("Save", mock.Anything).Return(nil)

	service := &manageVehicleService{
		gen:  gen,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}

	req := &vehicleRequestMock{
		ID:              string(DefaultVehicleID),
		Name:            AfterVehicleName,
		CommunicationID: string(AfterVehicleCommunicationID),
	}
	var resCall bool
	ret := service.UpdateVehicle(
		req,
		func(id, name, communicationID string) {
			resCall = true
		},
	)

	a.Nil(ret)
	a.True(resCall)
	a.Len(pub.events, 2)
	a.True(isClose)
	a.True(pub.isFlush)
	a.Nil(txm.isOpe)
	a.Nil(txm.isEH)
}

func TestUpdateVehicleOperation(t *testing.T) {
	a := assert.New(t)

	var (
		AfterVehicleName            = DefaultVehicleName + "-after"
		AfterVehicleCommunicationID = DefaultVehicleCommunicationID + "-after"
		DefaultVehicleVersion1      = DefaultVehicleVersion + "-1"
		DefaultVehicleVersion2      = DefaultVehicleVersion + "-2"
	)

	gen := &generatorMock{
		id:       DefaultVehicleID,
		versions: []v.Version{DefaultVehicleVersion1, DefaultVehicleVersion2},
	}

	vehicle := v.AssembleFrom(
		gen,
		&vehicleComponentMock{
			ID:              string(DefaultVehicleID),
			Name:            DefaultVehicleName,
			CommunicationID: string(DefaultVehicleCommunicationID),
			Version:         string(DefaultVehicleVersion),
		},
	)

	repo := &vehicleRepositoryMock{}
	repo.On("GetByID", DefaultVehicleID).Return(vehicle, nil)
	repo.On("Save", mock.Anything).Return(nil)
	pub := &publisherMock{}

	service := &manageVehicleService{
		gen:  nil,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	req := &vehicleRequestMock{
		ID:              string(DefaultVehicleID),
		Name:            AfterVehicleName,
		CommunicationID: string(AfterVehicleCommunicationID),
	}
	var resID, resName, resCommunicationID string
	ret := service.updateVehicleOperation(
		nil,
		pub,
		req,
		func(id, name, communicationID string) {
			resID = id
			resName = name
			resCommunicationID = communicationID
		},
	)

	expectEvent1 := v.CommunicationIdGaveEvent{CommunicationID: AfterVehicleCommunicationID}
	expectEvent2 := v.CommunicationIdRemovedEvent{CommunicationID: DefaultVehicleCommunicationID}

	a.Nil(ret)
	a.Equal(resID, string(DefaultVehicleID))
	a.Equal(resName, AfterVehicleName)
	a.Equal(resCommunicationID, string(AfterVehicleCommunicationID))
	a.Len(pub.events, 2)
	a.Equal(pub.events, []interface{}{expectEvent2, expectEvent1})
}

func TestDeleteVehicleTransaction(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMock{
		versions: []v.Version{DefaultVehicleVersion},
	}

	vehicle := v.AssembleFrom(
		gen,
		&vehicleComponentMock{
			ID:              string(DefaultVehicleID),
			Name:            DefaultVehicleName,
			CommunicationID: string(DefaultVehicleCommunicationID),
			Version:         string(DefaultVehicleVersion),
		},
	)

	repo := &vehicleRepositoryMock{}
	txm := &txManagerMock{}
	pub := &publisherMock{}
	psm := &pubSubManagerMock{}

	var isClose bool
	close := func() error {
		isClose = true
		return nil
	}

	psm.On("GetPublisher").Return(pub, close, nil)
	repo.On("GetByID", DefaultVehicleID).Return(vehicle, nil)
	repo.On("Delete", mock.Anything).Return(nil)

	service := &manageVehicleService{
		gen:  gen,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}

	req := &vehicleIDRequestMock{
		ID: string(DefaultVehicleID),
	}
	ret := service.DeleteVehicle(req)

	a.Nil(ret)
	a.Len(pub.events, 1)
	a.True(isClose)
	a.True(pub.isFlush)
	a.Nil(txm.isOpe)
	a.Nil(txm.isEH)
}

func TestDeleteVehicleOperation(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMock{
		versions: []v.Version{DefaultVehicleVersion},
	}

	vehicle := v.AssembleFrom(
		gen,
		&vehicleComponentMock{
			ID:              string(DefaultVehicleID),
			Name:            DefaultVehicleName,
			CommunicationID: string(DefaultVehicleCommunicationID),
			Version:         string(DefaultVehicleVersion),
		},
	)

	repo := &vehicleRepositoryMock{}
	repo.On("GetByID", DefaultVehicleID).Return(vehicle, nil)
	repo.On("Delete", mock.Anything).Return(nil)
	pub := &publisherMock{}

	service := &manageVehicleService{
		gen:  gen,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	req := &vehicleIDRequestMock{
		ID: string(DefaultVehicleID),
	}
	ret := service.deleteVehicleOperation(
		nil,
		pub,
		req,
	)

	expectEvent := v.CommunicationIdRemovedEvent{
		CommunicationID: DefaultVehicleCommunicationID,
	}

	a.Nil(ret)
	a.Len(pub.events, 1)
	a.Equal(pub.events[0], expectEvent)
}

func TestCarbonCopyVehicleTransaction(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultOriginalID = DefaultVehicleID + "-original"
		DefaultNewID      = DefaultVehicleID + "-new"
	)

	gen := &generatorMock{}

	vehicle := v.AssembleFrom(
		gen,
		&vehicleComponentMock{
			ID:              string(DefaultVehicleID),
			Name:            DefaultVehicleName,
			CommunicationID: string(DefaultVehicleCommunicationID),
			Version:         string(DefaultVehicleVersion),
		},
	)

	repo := &vehicleRepositoryMock{}
	txm := &txManagerMock{}
	pub := &publisherMock{}
	psm := &pubSubManagerMock{}

	var isClose bool
	close := func() error {
		isClose = true
		return nil
	}

	psm.On("GetPublisher").Return(pub, close, nil)
	repo.On("GetByID", DefaultNewID).Return(nil, v.ErrNotFound)
	repo.On("GetByID", DefaultOriginalID).Return(vehicle, nil)
	repo.On("Save", mock.Anything).Return(nil)

	service := &manageVehicleService{
		gen:  gen,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}

	req := &carbonCopyRequestMock{
		OriginalID: string(DefaultOriginalID),
		NewID:      string(DefaultNewID),
	}
	ret := service.CarbonCopyVehicle(req)

	a.Nil(ret)
	a.Len(pub.events, 1)
	a.True(isClose)
	a.True(pub.isFlush)
	a.Nil(txm.isOpe)
	a.Nil(txm.isEH)
}

func TestCarbonCopyVehicleOperation(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultOriginalID = DefaultVehicleID + "-original"
		DefaultNewID      = DefaultVehicleID + "-new"
	)

	gen := &generatorMock{}

	vehicle := v.AssembleFrom(
		gen,
		&vehicleComponentMock{
			ID:              string(DefaultVehicleID),
			Name:            DefaultVehicleName,
			CommunicationID: string(DefaultVehicleCommunicationID),
			Version:         string(DefaultVehicleVersion),
		},
	)

	repo := &vehicleRepositoryMock{}
	repo.On("GetByID", DefaultNewID).Return(nil, v.ErrNotFound)
	repo.On("GetByID", DefaultOriginalID).Return(vehicle, nil)
	repo.On("Save", mock.Anything).Return(nil)
	pub := &publisherMock{}

	service := &manageVehicleService{
		gen:  gen,
		repo: repo,
		txm:  nil,
		psm:  nil,
	}

	req := &carbonCopyRequestMock{
		OriginalID:   string(DefaultOriginalID),
		NewID:        string(DefaultNewID),
		FlightplanID: string(DefaultFlightplanID),
	}
	ret := service.carbonCopyVehicleOperation(
		nil,
		pub,
		req,
	)

	expectEvent := v.CopiedVehicleCreatedEvent{
		ID:              DefaultNewID,
		CommunicationID: DefaultVehicleCommunicationID,
		FlightplanID:    DefaultFlightplanID,
	}

	a.Nil(ret)
	a.Len(pub.events, 1)
	a.Equal(pub.events[0], expectEvent)
}
