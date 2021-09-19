package app

import "github.com/Tomofiles/skysign_cloud_v2/remote-communication/pkg/communication/service"

// Application .
type Application struct {
	Services Services
}

// Services .
type Services struct {
	ManageCommunication service.ManageCommunicationService
	UserCommunication   service.UserCommunicationService
	EdgeCommunication   service.EdgeCommunicationService
}
