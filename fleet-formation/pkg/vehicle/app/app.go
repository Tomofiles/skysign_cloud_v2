package app

import "github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/vehicle/service"

// Application .
type Application struct {
	Services Services
}

// Services .
type Services struct {
	ManageVehicle service.ManageVehicleService
}
