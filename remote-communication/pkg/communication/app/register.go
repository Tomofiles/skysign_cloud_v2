package app

import (
	"context"
	"remote-communication/pkg/common/domain/event"
	"remote-communication/pkg/common/domain/txmanager"
	"remote-communication/pkg/communication/adapters/postgresql"
	"remote-communication/pkg/communication/adapters/uuid"
	"remote-communication/pkg/communication/service"
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
	gen := uuid.NewCommunicationUUID()
	repo := postgresql.NewCommunicationRepository(gen)
	return Application{
		Services: Services{
			ManageCommunication: service.NewManageCommunicationService(gen, repo, txm, psm),
			UserCommunication:   service.NewUserCommunicationService(gen, repo, txm, psm),
			EdgeCommunication:   service.NewEdgeCommunicationService(gen, repo, txm, psm),
		},
	}
}
