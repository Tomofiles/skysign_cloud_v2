package app

import (
	"flightplan/pkg/flightplan/event"
	"flightplan/pkg/flightplan/service"
)

// Application .
type Application struct {
	Pub      event.Publisher
	Services Services
}

// Services .
type Services struct {
	ManageFlightplan service.ManageFlightplanService
	ManageFleet      service.ManageFleetService
	AssignFleet      service.AssignFleetService
}
