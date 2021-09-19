package app

import "github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightreport/service"

// Application .
type Application struct {
	Services Services
}

// Services .
type Services struct {
	ManageFlightreport service.ManageFlightreportService
}
