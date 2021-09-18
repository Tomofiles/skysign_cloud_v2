package grpc

import (
	"fleet-formation/pkg/vehicle/app"
	"testing"

	proto "github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSingleVehiclesListVehicles(t *testing.T) {
	a := assert.New(t)

	service := manageVehicleServiceMock{}

	vehicles := []vehicleMock{
		{
			ID:              DefaultVehicleID,
			Name:            DefaultVehicleName,
			CommunicationID: DefaultCommunicationID,
		},
	}
	service.On("ListVehicles", mock.Anything).Return(vehicles, nil)

	app := app.Application{
		Services: app.Services{
			ManageVehicle: &service,
		},
	}

	grpc := NewManageVehicleServiceServer(app)

	request := &proto.Empty{}
	response, err := grpc.ListVehicles(
		nil,
		request,
	)

	expectResponse := &proto.ListVehiclesResponses{
		Vehicles: []*proto.Vehicle{
			{
				Id:              DefaultVehicleID,
				Name:            DefaultVehicleName,
				CommunicationId: DefaultCommunicationID,
			},
		},
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestMultipleVehiclesListVehicles(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultVehicleID1       = DefaultVehicleID + "-1"
		DefaultVehicleName1     = DefaultVehicleName + "-1"
		DefaultCommunicationID1 = DefaultCommunicationID + "-1"
		DefaultVehicleID2       = DefaultVehicleID + "-2"
		DefaultVehicleName2     = DefaultVehicleName + "-2"
		DefaultCommunicationID2 = DefaultCommunicationID + "-2"
		DefaultVehicleID3       = DefaultVehicleID + "-3"
		DefaultVehicleName3     = DefaultVehicleName + "-3"
		DefaultCommunicationID3 = DefaultCommunicationID + "-3"
	)

	service := manageVehicleServiceMock{}

	vehicles := []vehicleMock{
		{
			ID:              DefaultVehicleID1,
			Name:            DefaultVehicleName1,
			CommunicationID: DefaultCommunicationID1,
		},
		{
			ID:              DefaultVehicleID2,
			Name:            DefaultVehicleName2,
			CommunicationID: DefaultCommunicationID2,
		},
		{
			ID:              DefaultVehicleID3,
			Name:            DefaultVehicleName3,
			CommunicationID: DefaultCommunicationID3,
		},
	}
	service.On("ListVehicles", mock.Anything).Return(vehicles, nil)

	app := app.Application{
		Services: app.Services{
			ManageVehicle: &service,
		},
	}

	grpc := NewManageVehicleServiceServer(app)

	request := &proto.Empty{}
	response, err := grpc.ListVehicles(
		nil,
		request,
	)

	expectResponse := &proto.ListVehiclesResponses{
		Vehicles: []*proto.Vehicle{
			{
				Id:              DefaultVehicleID1,
				Name:            DefaultVehicleName1,
				CommunicationId: DefaultCommunicationID1,
			},
			{
				Id:              DefaultVehicleID2,
				Name:            DefaultVehicleName2,
				CommunicationId: DefaultCommunicationID2,
			},
			{
				Id:              DefaultVehicleID3,
				Name:            DefaultVehicleName3,
				CommunicationId: DefaultCommunicationID3,
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

	request := &proto.Empty{}
	response, err := grpc.ListVehicles(
		nil,
		request,
	)

	expectResponse := &proto.ListVehiclesResponses{}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestGetVehicle(t *testing.T) {
	a := assert.New(t)

	service := manageVehicleServiceMock{}

	vehicle := vehicleMock{
		ID:              DefaultVehicleID,
		Name:            DefaultVehicleName,
		CommunicationID: DefaultCommunicationID,
	}
	service.On("GetVehicle", mock.Anything, mock.Anything).Return(vehicle, nil)

	app := app.Application{
		Services: app.Services{
			ManageVehicle: &service,
		},
	}

	grpc := NewManageVehicleServiceServer(app)

	request := &proto.GetVehicleRequest{
		Id: DefaultVehicleID,
	}
	response, err := grpc.GetVehicle(
		nil,
		request,
	)

	expectResponse := &proto.Vehicle{
		Id:              DefaultVehicleID,
		Name:            DefaultVehicleName,
		CommunicationId: DefaultCommunicationID,
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestCreateVehicle(t *testing.T) {
	a := assert.New(t)

	service := manageVehicleServiceMock{}

	vehicle := vehicleMock{
		ID:              DefaultVehicleID,
		Name:            DefaultVehicleName,
		CommunicationID: DefaultCommunicationID,
	}
	service.On("CreateVehicle", mock.Anything, mock.Anything).Return(vehicle, nil)

	app := app.Application{
		Services: app.Services{
			ManageVehicle: &service,
		},
	}

	grpc := NewManageVehicleServiceServer(app)

	request := &proto.Vehicle{
		Name:            DefaultVehicleName,
		CommunicationId: DefaultCommunicationID,
	}
	response, err := grpc.CreateVehicle(
		nil,
		request,
	)

	expectResponse := &proto.Vehicle{
		Id:              DefaultVehicleID,
		Name:            DefaultVehicleName,
		CommunicationId: DefaultCommunicationID,
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestUpdateVehicle(t *testing.T) {
	a := assert.New(t)

	service := manageVehicleServiceMock{}

	service.On("UpdateVehicle", mock.Anything, mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			ManageVehicle: &service,
		},
	}

	grpc := NewManageVehicleServiceServer(app)

	request := &proto.Vehicle{
		Id:              DefaultVehicleID,
		Name:            DefaultVehicleName,
		CommunicationId: DefaultCommunicationID,
	}
	response, err := grpc.UpdateVehicle(
		nil,
		request,
	)

	expectResponse := &proto.Vehicle{
		Id:              DefaultVehicleID,
		Name:            DefaultVehicleName,
		CommunicationId: DefaultCommunicationID,
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestDeleteVehicle(t *testing.T) {
	a := assert.New(t)

	service := manageVehicleServiceMock{}

	service.On("DeleteVehicle", mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			ManageVehicle: &service,
		},
	}

	grpc := NewManageVehicleServiceServer(app)

	request := &proto.DeleteVehicleRequest{
		Id: DefaultVehicleID,
	}
	response, err := grpc.DeleteVehicle(
		nil,
		request,
	)

	expectResponse := &proto.Empty{}

	a.Nil(err)
	a.Equal(response, expectResponse)
}
