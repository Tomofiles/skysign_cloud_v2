package grpc

import (
	fope "github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightoperation/domain/flightoperation"
	"github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightoperation/service"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

const DefaultID = fope.ID("flightoperation-id")
const DefaultName = "flightoperation-name"
const DefaultDescription = "flightoperation-description"
const DefaultFleetID = fope.FleetID("fleet-id")
const DefaultIsCompleted = fope.Completed
const DefaultVersion = fope.Version("version")

type manageFlightoperationServiceMock struct {
	mock.Mock
	name, description, fleetID string
}

func (s *manageFlightoperationServiceMock) GetFlightoperation(
	command service.GetFlightoperationCommand,
	model service.RetrievedModel,
) error {
	ret := s.Called()
	if flightoperation := ret.Get(0); flightoperation != nil {
		f := flightoperation.(flightoperationMock)
		model(
			&flightoperationModelMock{
				flightoperation: &f,
			},
		)
	}
	return ret.Error(1)
}

func (s *manageFlightoperationServiceMock) ListFlightoperations(
	model service.RetrievedModel,
) error {
	ret := s.Called()
	if flightoperations := ret.Get(0); flightoperations != nil {
		for _, f := range flightoperations.([]flightoperationMock) {
			model(
				&flightoperationModelMock{
					flightoperation: &f,
				},
			)
		}
	}
	return ret.Error(1)
}

func (s *manageFlightoperationServiceMock) CreateFlightoperation(
	command service.CreateFlightoperationCommand,
) error {
	s.name = command.GetFlightoperation().GetName()
	s.description = command.GetFlightoperation().GetDescription()
	s.fleetID = command.GetFlightoperation().GetFleetID()
	ret := s.Called()
	return ret.Error(0)
}

type operateFlightoperationServiceMock struct {
	mock.Mock
}

func (s *operateFlightoperationServiceMock) CompleteFlightoperation(
	command service.CompleteFlightoperationCommand,
) error {
	ret := s.Called()
	return ret.Error(0)
}

type flightoperationModelMock struct {
	flightoperation *flightoperationMock
}

func (f *flightoperationModelMock) GetFlightoperation() service.Flightoperation {
	return f.flightoperation
}

type flightoperationMock struct {
	id          string
	name        string
	description string
	fleetID     string
}

func (f *flightoperationMock) GetID() string {
	return f.id
}

func (f *flightoperationMock) GetName() string {
	return f.name
}

func (f *flightoperationMock) GetDescription() string {
	return f.description
}

func (f *flightoperationMock) GetFleetID() string {
	return f.fleetID
}

type serviceRegistrarMock struct {
	descs []*grpc.ServiceDesc
	impls []interface{}
}

func (s *serviceRegistrarMock) RegisterService(desc *grpc.ServiceDesc, impl interface{}) {
	s.descs = append(s.descs, desc)
	s.impls = append(s.impls, impl)
}
