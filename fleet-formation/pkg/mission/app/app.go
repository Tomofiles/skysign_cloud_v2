package app

import "github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/mission/service"

// Application .
type Application struct {
	Services Services
}

// Services .
type Services struct {
	ManageMission service.ManageMissionService
}
