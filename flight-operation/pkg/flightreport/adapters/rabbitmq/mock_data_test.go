package rabbitmq

import (
	"context"

	frep "github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightreport/domain/flightreport"
	"github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightreport/service"

	crm "github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/adapters/rabbitmq"

	"github.com/stretchr/testify/mock"
)

const DefaultFlightreportID = frep.ID("flightreport-id")
const DefaultFlightreportName = "flightreport-name"
const DefaultFlightreportDescription = "flightreport-description"
const DefaultFlightreportFleetID = frep.FleetID("fleet-id")

type manageFlightreportServiceMock struct {
	mock.Mock
	name, description, fleetID string
}

func (s *manageFlightreportServiceMock) GetFlightreport(
	command service.GetFlightreportCommand,
	retrievedModel service.RetrievedModel,
) error {
	ret := s.Called()
	if model := ret.Get(0); model != nil {
		f := model.(service.FlightreportPresentationModel)
		retrievedModel(f)
	}
	return ret.Error(1)
}

func (s *manageFlightreportServiceMock) ListFlightreports(
	retrievedModel service.RetrievedModel,
) error {
	ret := s.Called()
	if models := ret.Get(0); models != nil {
		for _, f := range models.([]service.FlightreportPresentationModel) {
			retrievedModel(f)
		}
	}
	return ret.Error(1)
}

func (s *manageFlightreportServiceMock) CreateFlightreport(
	command service.CreateFlightreportCommand,
) error {
	s.name = command.GetFlightreport().GetName()
	s.description = command.GetFlightreport().GetDescription()
	s.fleetID = command.GetFlightreport().GetFleetID()
	ret := s.Called()
	return ret.Error(0)
}

type flightreportComponentMock struct {
	ID          string
	Name        string
	Description string
	FleetID     string
}

func (f *flightreportComponentMock) GetID() string {
	return f.ID
}

func (f *flightreportComponentMock) GetName() string {
	return f.Name
}

func (f *flightreportComponentMock) GetDescription() string {
	return f.Description
}

func (f *flightreportComponentMock) GetFleetID() string {
	return f.FleetID
}

type flightreportModelMock struct {
	flightreport *frep.Flightreport
}

func (f *flightreportModelMock) GetFlightreport() service.Flightreport {
	return &flightreportMock{
		flightreport: f.flightreport,
	}
}

type flightreportMock struct {
	flightreport *frep.Flightreport
}

func (f *flightreportMock) GetID() string {
	return string(f.flightreport.GetID())
}

func (f *flightreportMock) GetName() string {
	return f.flightreport.GetName()
}

func (f *flightreportMock) GetDescription() string {
	return f.flightreport.GetDescription()
}

func (f *flightreportMock) GetFleetID() string {
	return string(f.flightreport.GetFleetID())
}

type pubSubManagerMock struct {
	consumers       []consumer
	publishHandlers []func(ch crm.Channel, e interface{})
}

func (h *pubSubManagerMock) SetConsumer(ctx context.Context, exchangeName, queueName string, handler func([]byte)) error {
	h.consumers = append(
		h.consumers,
		consumer{
			exchangeName: exchangeName,
			queueName:    queueName,
			handler:      handler,
		})
	return nil
}

func (h *pubSubManagerMock) SetPublishHandler(handler func(ch crm.Channel, e interface{})) error {
	h.publishHandlers = append(h.publishHandlers, handler)
	return nil
}

type consumer struct {
	exchangeName, queueName string
	handler                 func([]byte)
}
