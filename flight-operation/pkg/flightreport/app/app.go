package app

import "flight-operation/pkg/flightreport/service"

// Application .
type Application struct {
	Services Services
}

// Services .
type Services struct {
	ManageFlightreport service.ManageFlightreportService
}
