package service

import (
	"flightplan/pkg/flightplan/domain/event"
	fpl "flightplan/pkg/flightplan/domain/flightplan"
	"flightplan/pkg/flightplan/domain/txmanager"
)

// ManageFlightplanService .
type ManageFlightplanService struct {
	gen  fpl.Generator
	repo fpl.Repository
	txm  txmanager.TransactionManager
	psm  event.PubSubManager
}

// NewManageFlightplanService .
func NewManageFlightplanService(
	gen fpl.Generator,
	repo fpl.Repository,
	txm txmanager.TransactionManager,
	psm event.PubSubManager,
) ManageFlightplanService {
	return ManageFlightplanService{
		gen:  gen,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}
}

// GetFlightplan .
func (s *ManageFlightplanService) GetFlightplan(
	requestDpo GetFlightplanRequestDpo,
	responseDpo GetFlightplanResponseDpo,
) error {
	return s.txm.Do(func(tx txmanager.Tx) error {
		return s.getFlightplanOperation(
			tx,
			requestDpo,
			responseDpo,
		)
	})
}

func (s *ManageFlightplanService) getFlightplanOperation(
	tx txmanager.Tx,
	requestDpo GetFlightplanRequestDpo,
	responseDpo GetFlightplanResponseDpo,
) error {
	flightplan, err := s.repo.GetByID(tx, fpl.ID(requestDpo.GetID()))
	if err != nil {
		return err
	}
	if flightplan == nil {
		return fpl.ErrNotFound
	}

	responseDpo(
		string(flightplan.GetID()),
		flightplan.GetName(),
		flightplan.GetDescription(),
	)
	return nil
}

// ListFlightplans .
func (s *ManageFlightplanService) ListFlightplans(
	responseEachDpo ListFlightplansResponseDpo,
) error {
	return s.txm.Do(func(tx txmanager.Tx) error {
		return s.listFlightplansOperation(
			tx,
			responseEachDpo,
		)
	})
}

func (s *ManageFlightplanService) listFlightplansOperation(
	tx txmanager.Tx,
	responseEachDpo ListFlightplansResponseDpo,
) error {
	flightplans, err := s.repo.GetAll(tx)
	if err != nil {
		return err
	}

	for _, f := range flightplans {
		responseEachDpo(
			string(f.GetID()),
			f.GetName(),
			f.GetDescription(),
		)
	}
	return nil
}

// CreateFlightplan .
func (s *ManageFlightplanService) CreateFlightplan(
	requestDpo CreateFlightplanRequestDpo,
	responseDpo CreateFlightplanResponseDpo,
) error {
	pub, chClose, err := s.psm.GetPublisher()
	if err != nil {
		return err
	}
	defer chClose()

	return s.txm.DoAndEndHook(
		func(tx txmanager.Tx) error {
			return s.createFlightplanOperation(
				tx,
				pub,
				requestDpo,
				responseDpo,
			)
		},
		func() error {
			return pub.Flush()
		},
	)
}

func (s *ManageFlightplanService) createFlightplanOperation(
	tx txmanager.Tx,
	pub event.Publisher,
	requestDpo CreateFlightplanRequestDpo,
	responseDpo CreateFlightplanResponseDpo,
) error {
	id, ret := fpl.CreateNewFlightplan(
		tx,
		s.gen,
		s.repo,
		pub,
		requestDpo.GetName(),
		requestDpo.GetDescription())
	if ret != nil {
		return ret
	}

	responseDpo(id, requestDpo.GetName(), requestDpo.GetDescription())
	return nil
}

// UpdateFlightplan .
func (s *ManageFlightplanService) UpdateFlightplan(
	requestDpo UpdateFlightplanRequestDpo,
	responseDpo UpdateFlightplanResponseDpo,
) error {
	pub, chClose, err := s.psm.GetPublisher()
	if err != nil {
		return err
	}
	defer chClose()

	return s.txm.DoAndEndHook(
		func(tx txmanager.Tx) error {
			return s.updateFlightplanOperation(
				tx,
				pub,
				requestDpo,
				responseDpo,
			)
		},
		func() error {
			return pub.Flush()
		},
	)
}

func (s *ManageFlightplanService) updateFlightplanOperation(
	tx txmanager.Tx,
	pub event.Publisher,
	requestDpo UpdateFlightplanRequestDpo,
	responseDpo UpdateFlightplanResponseDpo,
) error {
	flightplan, err := s.repo.GetByID(tx, fpl.ID(requestDpo.GetID()))
	if err != nil {
		return err
	}
	if flightplan == nil {
		return fpl.ErrNotFound
	}

	flightplan.NameFlightplan(requestDpo.GetName())
	flightplan.ChangeDescription(requestDpo.GetDescription())

	if ret := s.repo.Save(tx, flightplan); ret != nil {
		return ret
	}

	responseDpo(
		string(flightplan.GetID()),
		flightplan.GetName(),
		flightplan.GetDescription(),
	)
	return nil
}

// DeleteFlightplan .
func (s *ManageFlightplanService) DeleteFlightplan(
	requestDpo DeleteFlightplanRequestDpo,
) error {
	pub, chClose, err := s.psm.GetPublisher()
	if err != nil {
		return err
	}
	defer chClose()

	return s.txm.DoAndEndHook(
		func(tx txmanager.Tx) error {
			return s.deleteFlightplanOperation(
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

func (s *ManageFlightplanService) deleteFlightplanOperation(
	tx txmanager.Tx,
	pub event.Publisher,
	requestDpo DeleteFlightplanRequestDpo,
) error {
	if ret := fpl.DeleteFlightplan(
		tx,
		s.repo,
		pub,
		fpl.ID(requestDpo.GetID()),
	); ret != nil {
		return ret
	}

	return nil
}

// CreateFlightplanRequestDpo .
type CreateFlightplanRequestDpo interface {
	GetName() string
	GetDescription() string
}

// CreateFlightplanResponseDpo .
type CreateFlightplanResponseDpo = func(id, name, description string)

// UpdateFlightplanRequestDpo .
type UpdateFlightplanRequestDpo interface {
	GetID() string
	GetName() string
	GetDescription() string
}

// UpdateFlightplanResponseDpo .
type UpdateFlightplanResponseDpo = func(id, name, description string)

// GetFlightplanRequestDpo .
type GetFlightplanRequestDpo interface {
	GetID() string
}

// GetFlightplanResponseDpo .
type GetFlightplanResponseDpo = func(id, name, description string)

// ListFlightplansResponseDpo .
type ListFlightplansResponseDpo = func(id, name, description string)

// DeleteFlightplanRequestDpo .
type DeleteFlightplanRequestDpo interface {
	GetID() string
}
