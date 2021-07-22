package app

import (
	"fleet-formation/pkg/fleet/service"
)

// Application .
type Application struct {
	Services Services
}

// Services .
type Services struct {
	ManageFleet service.ManageFleetService
	AssignFleet service.AssignFleetService
}
