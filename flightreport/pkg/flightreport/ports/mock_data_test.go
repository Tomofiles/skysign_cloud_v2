package ports

import (
	frep "flightreport/pkg/flightreport/domain/flightreport"
	"flightreport/pkg/flightreport/service"

	"github.com/stretchr/testify/mock"
)

const DefaultFlightreportID = frep.ID("flightreport-id")
const DefaultFlightoperationID = frep.FlightoperationID("flightoperation-id")

type manageFlightreportServiceMock struct {
	mock.Mock
}

func (s *manageFlightreportServiceMock) GetFlightreport(
	requestDpo service.GetFlightreportRequestDpo,
	responseDpo service.GetFlightreportResponseDpo,
) error {
	ret := s.Called()
	if flightreport := ret.Get(0); flightreport != nil {
		f := flightreport.(flightreportMock)
		responseDpo(f.id, f.flightoperationID)
	}
	return ret.Error(1)
}

func (s *manageFlightreportServiceMock) ListFlightreports(
	responseEachDpo service.ListFlightreportsResponseDpo,
) error {
	ret := s.Called()
	if flightreports := ret.Get(0); flightreports != nil {
		for _, f := range flightreports.([]flightreportMock) {
			responseEachDpo(f.id, f.flightoperationID)
		}
	}
	return ret.Error(1)
}

func (s *manageFlightreportServiceMock) CreateFlightreport(
	requestDpo service.CreateFlightreportRequestDpo,
) error {
	ret := s.Called()
	return ret.Error(0)
}

type flightreportMock struct {
	id, flightoperationID string
}
