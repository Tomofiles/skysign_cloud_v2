package app

import (
	"github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightplan/service"
)

// Application .
type Application struct {
	Services Services
}

// Services .
type Services struct {
	ManageFlightplan  service.ManageFlightplanService
	ChangeFlightplan  service.ChangeFlightplanService
	ExecuteFlightplan service.ExecuteFlightplanService
}
