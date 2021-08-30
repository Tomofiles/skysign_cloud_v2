package app

import "remote-communication/pkg/mission/service"

// Application .
type Application struct {
	Services Services
}

// Services .
type Services struct {
	ManageMission service.ManageMissionService
	EdgeMission   service.EdgeMissionService
}
