package app

import "github.com/Tomofiles/skysign_cloud_v2/collection-analysis/pkg/action/service"

// Application .
type Application struct {
	Services Services
}

// Services .
type Services struct {
	ManageAction  service.ManageActionService
	OperateAction service.OperateActionService
}
