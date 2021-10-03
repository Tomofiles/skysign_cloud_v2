package grpc

import (
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/vehicle/app"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSingleVehiclesListVehicles(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultVehicleID              = NewVehicleID()
		DefaultVehicleCommunicationID = NewVehicleCommunicationID()
	)

	service := manageVehicleServiceMock{}

	vehicles := []vehicleMock{
		{
			ID:              DefaultVehicleID,
			Name:            DefaultVehicleName,
			CommunicationID: DefaultVehicleCommunicationID,
		},
	}
	service.On("ListVehicles", mock.Anything).Return(vehicles, nil)

	app := app.Application{
		Services: app.Services{
			ManageVehicle: &service,
		},
	}

	grpc := NewManageVehicleServiceServer(app)

	request := &skysign_proto.Empty{}
	response, err := grpc.ListVehicles(
		nil,
		request,
	)

	expectResponse := &skysign_proto.ListVehiclesResponses{
		Vehicles: []*skysign_proto.Vehicle{
			{
				Id:              DefaultVehicleID,
				Name:            DefaultVehicleName,
				CommunicationId: DefaultVehicleCommunicationID,
			},
		},
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestMultipleVehiclesListVehicles(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultVehicleID1              = NewVehicleID()
		DefaultVehicleName1            = DefaultVehicleName + "-1"
		DefaultVehicleCommunicationID1 = NewVehicleCommunicationID()
		DefaultVehicleID2              = NewVehicleID()
		DefaultVehicleName2            = DefaultVehicleName + "-2"
		DefaultVehicleCommunicationID2 = NewVehicleCommunicationID()
		DefaultVehicleID3              = NewVehicleID()
		DefaultVehicleName3            = DefaultVehicleName + "-3"
		DefaultVehicleCommunicationID3 = NewVehicleCommunicationID()
	)

	service := manageVehicleServiceMock{}

	vehicles := []vehicleMock{
		{
			ID:              DefaultVehicleID1,
			Name:            DefaultVehicleName1,
			CommunicationID: DefaultVehicleCommunicationID1,
		},
		{
			ID:              DefaultVehicleID2,
			Name:            DefaultVehicleName2,
			CommunicationID: DefaultVehicleCommunicationID2,
		},
		{
			ID:              DefaultVehicleID3,
			Name:            DefaultVehicleName3,
			CommunicationID: DefaultVehicleCommunicationID3,
		},
	}
	service.On("ListVehicles", mock.Anything).Return(vehicles, nil)

	app := app.Application{
		Services: app.Services{
			ManageVehicle: &service,
		},
	}

	grpc := NewManageVehicleServiceServer(app)

	request := &skysign_proto.Empty{}
	response, err := grpc.ListVehicles(
		nil,
		request,
	)

	expectResponse := &skysign_proto.ListVehiclesResponses{
		Vehicles: []*skysign_proto.Vehicle{
			{
				Id:              DefaultVehicleID1,
				Name:            DefaultVehicleName1,
				CommunicationId: DefaultVehicleCommunicationID1,
			},
			{
				Id:              DefaultVehicleID2,
				Name:            DefaultVehicleName2,
				CommunicationId: DefaultVehicleCommunicationID2,
			},
			{
				Id:              DefaultVehicleID3,
				Name:            DefaultVehicleName3,
				CommunicationId: DefaultVehicleCommunicationID3,
			},
		},
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestNoneVehiclesListVehicles(t *testing.T) {
	a := assert.New(t)

	service := manageVehicleServiceMock{}

	vehicles := []vehicleMock{}
	service.On("ListVehicles", mock.Anything).Return(vehicles, nil)

	app := app.Application{
		Services: app.Services{
			ManageVehicle: &service,
		},
	}

	grpc := NewManageVehicleServiceServer(app)

	request := &skysign_proto.Empty{}
	response, err := grpc.ListVehicles(
		nil,
		request,
	)

	expectResponse := &skysign_proto.ListVehiclesResponses{}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestGetVehicle(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultVehicleID              = NewVehicleID()
		DefaultVehicleCommunicationID = NewVehicleCommunicationID()
	)

	service := manageVehicleServiceMock{}

	vehicle := vehicleMock{
		ID:              DefaultVehicleID,
		Name:            DefaultVehicleName,
		CommunicationID: DefaultVehicleCommunicationID,
	}
	service.On("GetVehicle", mock.Anything, mock.Anything).Return(vehicle, nil)

	app := app.Application{
		Services: app.Services{
			ManageVehicle: &service,
		},
	}

	grpc := NewManageVehicleServiceServer(app)

	request := &skysign_proto.GetVehicleRequest{
		Id: DefaultVehicleID,
	}
	response, err := grpc.GetVehicle(
		nil,
		request,
	)

	expectResponse := &skysign_proto.Vehicle{
		Id:              DefaultVehicleID,
		Name:            DefaultVehicleName,
		CommunicationId: DefaultVehicleCommunicationID,
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestCreateVehicle(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultVehicleID              = NewVehicleID()
		DefaultVehicleCommunicationID = NewVehicleCommunicationID()
	)

	service := manageVehicleServiceMock{}

	vehicle := vehicleMock{
		ID:              DefaultVehicleID,
		Name:            DefaultVehicleName,
		CommunicationID: DefaultVehicleCommunicationID,
	}
	service.On("CreateVehicle", mock.Anything, mock.Anything).Return(vehicle, nil)

	app := app.Application{
		Services: app.Services{
			ManageVehicle: &service,
		},
	}

	grpc := NewManageVehicleServiceServer(app)

	request := &skysign_proto.Vehicle{
		Name:            DefaultVehicleName,
		CommunicationId: DefaultVehicleCommunicationID,
	}
	response, err := grpc.CreateVehicle(
		nil,
		request,
	)

	expectResponse := &skysign_proto.Vehicle{
		Id:              DefaultVehicleID,
		Name:            DefaultVehicleName,
		CommunicationId: DefaultVehicleCommunicationID,
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestUpdateVehicle(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultVehicleID              = NewVehicleID()
		DefaultVehicleCommunicationID = NewVehicleCommunicationID()
	)

	service := manageVehicleServiceMock{}

	service.On("UpdateVehicle", mock.Anything, mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			ManageVehicle: &service,
		},
	}

	grpc := NewManageVehicleServiceServer(app)

	request := &skysign_proto.Vehicle{
		Id:              DefaultVehicleID,
		Name:            DefaultVehicleName,
		CommunicationId: DefaultVehicleCommunicationID,
	}
	response, err := grpc.UpdateVehicle(
		nil,
		request,
	)

	expectResponse := &skysign_proto.Vehicle{
		Id:              DefaultVehicleID,
		Name:            DefaultVehicleName,
		CommunicationId: DefaultVehicleCommunicationID,
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestDeleteVehicle(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultVehicleID = NewVehicleID()
	)

	service := manageVehicleServiceMock{}

	service.On("DeleteVehicle", mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			ManageVehicle: &service,
		},
	}

	grpc := NewManageVehicleServiceServer(app)

	request := &skysign_proto.DeleteVehicleRequest{
		Id: DefaultVehicleID,
	}
	response, err := grpc.DeleteVehicle(
		nil,
		request,
	)

	expectResponse := &skysign_proto.Empty{}

	a.Nil(err)
	a.Equal(response, expectResponse)
}
