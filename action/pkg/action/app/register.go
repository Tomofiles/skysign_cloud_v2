package app

import (
	"action/pkg/action/adapters/postgresql"
	"action/pkg/action/domain/txmanager"
	"action/pkg/action/service"
	"context"
)

// NewApplication .
func NewApplication(
	ctx context.Context,
	txm txmanager.TransactionManager,
) Application {
	return newApplication(ctx, txm)
}

func newApplication(
	ctx context.Context,
	txm txmanager.TransactionManager,
) Application {
	repo := postgresql.NewActionRepository()
	return Application{
		Services: Services{
			ManageAction:  service.NewManageActionService(repo, txm),
			OperateAction: service.NewOperateActionService(repo, txm),
		},
	}
}
