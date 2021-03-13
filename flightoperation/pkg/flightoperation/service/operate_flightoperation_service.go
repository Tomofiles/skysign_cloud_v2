package service

import (
	"flightoperation/pkg/flightoperation/domain/event"
	fope "flightoperation/pkg/flightoperation/domain/flightoperation"
	"flightoperation/pkg/flightoperation/domain/txmanager"
)

// OperateFlightoperationService .
type OperateFlightoperationService interface {
	CompleteFlightoperation(requestDpo CompleteFlightoperationRequestDpo) error
}

// CompleteFlightoperationRequestDpo .
type CompleteFlightoperationRequestDpo interface {
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
	requestDpo CompleteFlightoperationRequestDpo,
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
				requestDpo,
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
	requestDpo CompleteFlightoperationRequestDpo,
) error {
	flightoperation, err := s.repo.GetByID(tx, fope.ID(requestDpo.GetID()))
	if err != nil {
		return err
	}

	flightoperation.SetPublisher(pub)
	if err := flightoperation.Complete(); err != nil {
		return err
	}

	if err := s.repo.Save(tx, flightoperation); err != nil {
		return err
	}

	return nil
}
