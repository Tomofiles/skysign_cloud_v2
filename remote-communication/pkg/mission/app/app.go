package app

import "github.com/Tomofiles/skysign_cloud_v2/remote-communication/pkg/mission/service"

// Application .
type Application struct {
	Services Services
}

// Services .
type Services struct {
	ManageMission service.ManageMissionService
	EdgeMission   service.EdgeMissionService
}
