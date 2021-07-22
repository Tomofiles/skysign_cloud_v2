package app

import "collection-analysis/pkg/action/service"

// Application .
type Application struct {
	Services Services
}

// Services .
type Services struct {
	ManageAction  service.ManageActionService
	OperateAction service.OperateActionService
}
