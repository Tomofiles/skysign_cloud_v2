package grpc

import (
	"github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/vehicle/service"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

const DefaultVehicleID = "vehicle-id"
const DefaultVehicleName = "vehicle-name"
const DefaultCommunicationID = "communication-id"
const DefaultFleetID = "fleet-id"

type manageVehicleServiceMock struct {
	mock.Mock
	OriginalID string
	NewID      string
	FleetID    string
}

func (s *manageVehicleServiceMock) GetVehicle(
	command service.GetVehicleCommand,
	model service.RetrievedModel,
) error {
	ret := s.Called()
	if vehicle := ret.Get(0); vehicle != nil {
		f := vehicle.(vehicleMock)
		model(
			&vehicleModelMock{
				vehicle: &f,
			},
		)
	}
	return ret.Error(1)
}

func (s *manageVehicleServiceMock) ListVehicles(
	model service.RetrievedModel,
) error {
	ret := s.Called()
	if vehicles := ret.Get(0); vehicles != nil {
		for _, f := range vehicles.([]vehicleMock) {
			model(
				&vehicleModelMock{
					vehicle: &f,
				},
			)
		}
	}
	return ret.Error(1)
}

func (s *manageVehicleServiceMock) CreateVehicle(
	command service.CreateVehicleCommand,
	createdID service.CreatedID,
) error {
	ret := s.Called()
	if vehicle := ret.Get(0); vehicle != nil {
		f := vehicle.(vehicleMock)
		createdID(f.ID)
	}
	return ret.Error(1)
}

func (s *manageVehicleServiceMock) UpdateVehicle(
	command service.UpdateVehicleCommand,
) error {
	ret := s.Called()
	return ret.Error(0)
}

func (s *manageVehicleServiceMock) DeleteVehicle(
	command service.DeleteVehicleCommand,
) error {
	ret := s.Called()
	return ret.Error(0)
}

func (s *manageVehicleServiceMock) CarbonCopyVehicle(
	command service.CarbonCopyVehicleCommand,
) error {
	ret := s.Called()
	s.OriginalID = command.GetOriginalID()
	s.NewID = command.GetNewID()
	s.FleetID = command.GetFleetID()
	return ret.Error(0)
}

type vehicleModelMock struct {
	vehicle *vehicleMock
}

func (f *vehicleModelMock) GetVehicle() service.Vehicle {
	return &vehicleMock{
		ID:              f.vehicle.ID,
		Name:            f.vehicle.Name,
		CommunicationID: f.vehicle.CommunicationID,
	}
}

type vehicleMock struct {
	ID              string
	Name            string
	CommunicationID string
}

func (f *vehicleMock) GetID() string {
	return f.ID
}

func (f *vehicleMock) GetName() string {
	return f.Name
}

func (f *vehicleMock) GetCommunicationID() string {
	return f.CommunicationID
}

type serviceRegistrarMock struct {
	descs []*grpc.ServiceDesc
	impls []interface{}
}

func (s *serviceRegistrarMock) RegisterService(desc *grpc.ServiceDesc, impl interface{}) {
	s.descs = append(s.descs, desc)
	s.impls = append(s.impls, impl)
}
