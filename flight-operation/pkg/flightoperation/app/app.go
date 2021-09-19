package app

import "github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightoperation/service"

// Application .
type Application struct {
	Services Services
}

// Services .
type Services struct {
	ManageFlightoperation  service.ManageFlightoperationService
	OperateFlightoperation service.OperateFlightoperationService
}
