package ports

import (
	fope "flightreport/pkg/flightreport/domain/flightoperation"
	"flightreport/pkg/flightreport/service"

	"github.com/stretchr/testify/mock"
)

const DefaultFlightoperationID = fope.ID("flightoperation-id")
const DefaultFlightplanID = fope.FlightplanID("flightplan-id")

type manageFlightoperationServiceMock struct {
	mock.Mock
}

func (s *manageFlightoperationServiceMock) GetFlightoperation(
	requestDpo service.GetFlightoperationRequestDpo,
	responseDpo service.GetFlightoperationResponseDpo,
) error {
	ret := s.Called()
	if flightoperation := ret.Get(0); flightoperation != nil {
		f := flightoperation.(flightoperationMock)
		responseDpo(f.id, f.flightplanID)
	}
	return ret.Error(1)
}

func (s *manageFlightoperationServiceMock) ListFlightoperations(
	responseEachDpo service.ListFlightoperationsResponseDpo,
) error {
	ret := s.Called()
	if flightoperations := ret.Get(0); flightoperations != nil {
		for _, f := range flightoperations.([]flightoperationMock) {
			responseEachDpo(f.id, f.flightplanID)
		}
	}
	return ret.Error(1)
}

func (s *manageFlightoperationServiceMock) CreateFlightoperation(
	requestDpo service.CreateFlightoperationRequestDpo,
) error {
	ret := s.Called()
	return ret.Error(0)
}

type flightoperationMock struct {
	id, flightplanID string
}
