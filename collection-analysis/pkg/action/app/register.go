package app

import (
	"context"

	"github.com/Tomofiles/skysign_cloud_v2/collection-analysis/pkg/action/adapters/postgresql"
	"github.com/Tomofiles/skysign_cloud_v2/collection-analysis/pkg/action/service"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/txmanager"
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
