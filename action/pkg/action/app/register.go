package app

import (
	"action/pkg/action/domain/event"
	"action/pkg/action/domain/txmanager"
	"context"
)

// NewApplication .
func NewApplication(
	ctx context.Context,
	txm txmanager.TransactionManager,
	psm event.PubSubManager,
) Application {
	return newApplication(ctx, txm, psm)
}

func newApplication(
	ctx context.Context,
	txm txmanager.TransactionManager,
	psm event.PubSubManager,
) Application {
	return Application{
		Services: Services{},
	}
}
