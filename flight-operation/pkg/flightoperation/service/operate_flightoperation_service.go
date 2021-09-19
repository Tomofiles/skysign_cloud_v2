package service

import (
	fope "github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightoperation/domain/flightoperation"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/event"
	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/txmanager"
)

// OperateFlightoperationService .
type OperateFlightoperationService interface {
	CompleteFlightoperation(command CompleteFlightoperationCommand) error
}

// CompleteFlightoperationCommand .
type CompleteFlightoperationCommand interface {
	GetID() string
}

// NewOperateFlightoperationService .
func NewOperateFlightoperationService(
	gen fope.Generator,
	repo fope.Repository,
	txm txmanager.TransactionManager,
	psm event.PubSubManager,
) OperateFlightoperationService {
	return &operateFlightoperationService{
		gen:  gen,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}
}

type operateFlightoperationService struct {
	gen  fope.Generator
	repo fope.Repository
	txm  txmanager.TransactionManager
	psm  event.PubSubManager
}

func (s *operateFlightoperationService) CompleteFlightoperation(
	command CompleteFlightoperationCommand,
) error {
	pub, chClose, err := s.psm.GetPublisher()
	if err != nil {
		return err
	}
	defer chClose()

	return s.txm.DoAndEndHook(
		func(tx txmanager.Tx) error {
			return s.completeFlightoperationOperation(
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

func (s *operateFlightoperationService) completeFlightoperationOperation(
	tx txmanager.Tx,
	pub event.Publisher,
	command CompleteFlightoperationCommand,
) error {
	if ret := fope.CompleteFlightoperation(
		tx,
		s.repo,
		pub,
		fope.ID(command.GetID()),
	); ret != nil {
		return ret
	}

	return nil
}
