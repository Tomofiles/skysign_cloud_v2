package grpc

import (
	f "flight-operation/pkg/flightplan/domain/flightplan"
	"flight-operation/pkg/flightplan/service"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

const DefaultFlightplanID = "flightplan-id"
const DefaultFlightplanName = "flightplan-name"
const DefaultFlightplanDescription = "flightplan-description"
const DefaultFlightplanFleetID = "fleet-id"
const DefaultFleetNumberOfVehicles = 3

type manageFlightplanServiceMock struct {
	mock.Mock
	OriginalID string
	NewID      string
}

func (s *manageFlightplanServiceMock) GetFlightplan(
	command service.GetFlightplanCommand,
	retrievedModel service.RetrievedModel,
) error {
	ret := s.Called()
	if model := ret.Get(0); model != nil {
		f := model.(service.FlightplanPresentationModel)
		retrievedModel(f)
	}
	return ret.Error(1)
}

func (s *manageFlightplanServiceMock) ListFlightplans(
	retrievedModel service.RetrievedModel,
) error {
	ret := s.Called()
	if models := ret.Get(0); models != nil {
		for _, f := range models.([]service.FlightplanPresentationModel) {
			retrievedModel(f)
		}
	}
	return ret.Error(1)
}

func (s *manageFlightplanServiceMock) CreateFlightplan(
	command service.CreateFlightplanCommand,
	createdID service.CreatedID,
	fleetID service.FleetID,
) error {
	ret := s.Called()
	if model := ret.Get(0); model != nil {
		f := model.(service.FlightplanPresentationModel)
		createdID(f.GetFlightplan().GetID())
		fleetID(f.GetFlightplan().GetFleetID())
	}
	return ret.Error(1)
}

func (s *manageFlightplanServiceMock) UpdateFlightplan(
	command service.UpdateFlightplanCommand,
	fleetID service.FleetID,
) error {
	ret := s.Called()
	if model := ret.Get(0); model != nil {
		f := model.(service.FlightplanPresentationModel)
		fleetID(f.GetFlightplan().GetFleetID())
	}
	return ret.Error(1)
}

func (s *manageFlightplanServiceMock) DeleteFlightplan(
	command service.DeleteFlightplanCommand,
) error {
	ret := s.Called()
	return ret.Error(0)
}

type changeFlightplanServiceMock struct {
	mock.Mock
}

func (s *changeFlightplanServiceMock) ChangeNumberOfVehicles(
	command service.ChangeNumberOfVehiclesCommand,
) error {
	ret := s.Called()
	return ret.Error(0)
}

type executeFlightplanServiceMock struct {
	mock.Mock
}

func (s *executeFlightplanServiceMock) ExecuteFlightplan(
	command service.ExecuteFlightplanCommand,
) error {
	ret := s.Called()
	return ret.Error(0)
}

type flightplanModelMock struct {
	flightplan *f.Flightplan
}

func (f *flightplanModelMock) GetFlightplan() service.Flightplan {
	return &flightplanMock{
		flightplan: f.flightplan,
	}
}

type flightplanMock struct {
	flightplan *f.Flightplan
}

func (f *flightplanMock) GetID() string {
	return string(f.flightplan.GetID())
}

func (f *flightplanMock) GetName() string {
	return f.flightplan.GetName()
}

func (f *flightplanMock) GetDescription() string {
	return f.flightplan.GetDescription()
}

func (f *flightplanMock) GetFleetID() string {
	return string(f.flightplan.GetFleetID())
}

// Flightplan構成オブジェクトモック
type flightplanComponentMock struct {
	ID          string
	Name        string
	Description string
	FleetID     string
	Version     string
}

func (f *flightplanComponentMock) GetID() string {
	return f.ID
}

func (f *flightplanComponentMock) GetName() string {
	return f.Name
}

func (f *flightplanComponentMock) GetDescription() string {
	return f.Description
}

func (f *flightplanComponentMock) GetFleetID() string {
	return f.FleetID
}

func (f *flightplanComponentMock) GetVersion() string {
	return f.Version
}

type serviceRegistrarMock struct {
	descs []*grpc.ServiceDesc
	impls []interface{}
}

func (s *serviceRegistrarMock) RegisterService(desc *grpc.ServiceDesc, impl interface{}) {
	s.descs = append(s.descs, desc)
	s.impls = append(s.impls, impl)
}
