package app

import (
	"collection-analysis/pkg/action/adapters/postgresql"
	"collection-analysis/pkg/action/service"
	"collection-analysis/pkg/common/domain/txmanager"
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
