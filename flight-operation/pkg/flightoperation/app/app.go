package app

import "flight-operation/pkg/flightoperation/service"

// Application .
type Application struct {
	Services Services
}

// Services .
type Services struct {
	ManageFlightoperation  service.ManageFlightoperationService
	OperateFlightoperation service.OperateFlightoperationService
}
