package service

import (
	c "remote-communication/pkg/communication/domain/communication"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/event"
	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/txmanager"
)

// ManageCommunicationService .
type ManageCommunicationService interface {
	CreateCommunication(command CreateCommunicationCommand) error
	DeleteCommunication(command DeleteCommunicationCommand) error
}

// CreateCommunicationCommand .
type CreateCommunicationCommand interface {
	GetID() string
}

// DeleteCommunicationCommand .
type DeleteCommunicationCommand interface {
	GetID() string
}

// NewManageCommunicationService .
func NewManageCommunicationService(
	gen c.Generator,
	repo c.Repository,
	txm txmanager.TransactionManager,
	psm event.PubSubManager,
) ManageCommunicationService {
	return &manageCommunicationService{
		gen:  gen,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}
}

type manageCommunicationService struct {
	gen  c.Generator
	repo c.Repository
	txm  txmanager.TransactionManager
	psm  event.PubSubManager
}

func (s *manageCommunicationService) CreateCommunication(
	command CreateCommunicationCommand,
) error {
	pub, chClose, err := s.psm.GetPublisher()
	if err != nil {
		return err
	}
	defer chClose()

	return s.txm.DoAndEndHook(
		func(tx txmanager.Tx) error {
			return s.createCommunicationOperation(
				tx,
				pub,
				command,
			)
		},
		func() error {
			return pub.Flush()
		},
	)
}

func (s *manageCommunicationService) createCommunicationOperation(
	tx txmanager.Tx,
	pub event.Publisher,
	command CreateCommunicationCommand,
) error {
	communication := c.NewInstance(s.gen, c.ID(command.GetID()))

	if ret := s.repo.Save(tx, communication); ret != nil {
		return ret
	}

	return nil
}

func (s *manageCommunicationService) DeleteCommunication(
	command DeleteCommunicationCommand,
) error {
	pub, chClose, err := s.psm.GetPublisher()
	if err != nil {
		return err
	}
	defer chClose()

	return s.txm.DoAndEndHook(
		func(tx txmanager.Tx) error {
			return s.deleteCommunicationOperation(
				tx,
				pub,
				command,
			)
		},
		func() error {
			return pub.Flush()
		},
	)
}

func (s *manageCommunicationService) deleteCommunicationOperation(
	tx txmanager.Tx,
	pub event.Publisher,
	command DeleteCommunicationCommand,
) error {
	return s.repo.Delete(tx, c.ID(command.GetID()))
}
