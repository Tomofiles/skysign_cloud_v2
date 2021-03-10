package app

import "flightreport/pkg/flightreport/service"

// Application .
type Application struct {
	Services Services
}

// Services .
type Services struct {
	ManageFlightoperation service.ManageFlightoperationService
}
