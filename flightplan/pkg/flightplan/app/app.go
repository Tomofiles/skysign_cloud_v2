package app

import (
	"flightplan/pkg/flightplan/service"
)

// Application .
type Application struct {
	Services Services
}

// Services .
type Services struct {
	ManageFlightplan service.ManageFlightplanService
	ChangeFlightplan service.ChangeFlightplanService
	ManageFleet      service.ManageFleetService
	AssignFleet      service.AssignFleetService
}
