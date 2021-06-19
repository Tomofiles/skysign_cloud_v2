package app

import "vehicle/pkg/vehicle/service"

// Application .
type Application struct {
	Services Services
}

// Services .
type Services struct {
	ManageVehicle service.ManageVehicleService
}
