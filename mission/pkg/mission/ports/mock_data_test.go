package ports

// import (
// 	"vehicle/pkg/vehicle/service"

// 	"github.com/stretchr/testify/mock"
// )

// const DefaultVehicleID = "vehicle-id"
// const DefaultVehicleName = "vehicle-name"
// const DefaultCommunicationID = "communication-id"
// const DefaultFlightplanID = "flightplan-id"

// type manageVehicleServiceMock struct {
// 	mock.Mock
// 	OriginalID   string
// 	NewID        string
// 	FlightplanID string
// }

// func (s *manageVehicleServiceMock) GetVehicle(
// 	requestDpo service.GetVehicleRequestDpo,
// 	responseDpo service.GetVehicleResponseDpo,
// ) error {
// 	ret := s.Called()
// 	if vehicle := ret.Get(0); vehicle != nil {
// 		f := vehicle.(vehicleMock)
// 		responseDpo(f.id, f.name, f.communicationID)
// 	}
// 	return ret.Error(1)
// }

// func (s *manageVehicleServiceMock) ListVehicles(
// 	responseEachDpo service.ListVehiclesResponseDpo,
// ) error {
// 	ret := s.Called()
// 	if vehicles := ret.Get(0); vehicles != nil {
// 		for _, f := range vehicles.([]vehicleMock) {
// 			responseEachDpo(f.id, f.name, f.communicationID)
// 		}
// 	}
// 	return ret.Error(1)
// }

// func (s *manageVehicleServiceMock) CreateVehicle(
// 	requestDpo service.CreateVehicleRequestDpo,
// 	responseDpo service.CreateVehicleResponseDpo,
// ) error {
// 	ret := s.Called()
// 	if vehicle := ret.Get(0); vehicle != nil {
// 		f := vehicle.(vehicleMock)
// 		responseDpo(f.id, f.name, f.communicationID)
// 	}
// 	return ret.Error(1)
// }

// func (s *manageVehicleServiceMock) UpdateVehicle(
// 	requestDpo service.UpdateVehicleRequestDpo,
// 	responseDpo service.UpdateVehicleResponseDpo,
// ) error {
// 	ret := s.Called()
// 	if vehicle := ret.Get(0); vehicle != nil {
// 		f := vehicle.(vehicleMock)
// 		responseDpo(f.id, f.name, f.communicationID)
// 	}
// 	return ret.Error(1)
// }

// func (s *manageVehicleServiceMock) DeleteVehicle(
// 	requestDpo service.DeleteVehicleRequestDpo,
// ) error {
// 	ret := s.Called()
// 	return ret.Error(0)
// }

// func (s *manageVehicleServiceMock) CarbonCopyVehicle(
// 	requestDpo service.CarbonCopyVehicleRequestDpo,
// ) error {
// 	ret := s.Called()
// 	s.OriginalID = requestDpo.GetOriginalID()
// 	s.NewID = requestDpo.GetNewID()
// 	s.FlightplanID = requestDpo.GetFlightplanID()
// 	return ret.Error(0)
// }

// type vehicleMock struct {
// 	id, name, communicationID string
// }
