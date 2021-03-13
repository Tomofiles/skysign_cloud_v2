package app

import "flightoperation/pkg/flightoperation/service"

// Application .
type Application struct {
	Services Services
}

// Services .
type Services struct {
	ManageFlightoperation  service.ManageFlightoperationService
	OperateFlightoperation service.OperateFlightoperationService
}
