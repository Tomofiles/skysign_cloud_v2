package app

import "fleet-formation/pkg/mission/service"

// Application .
type Application struct {
	Services Services
}

// Services .
type Services struct {
	ManageMission service.ManageMissionService
}
