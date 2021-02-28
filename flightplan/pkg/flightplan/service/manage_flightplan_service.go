package service

import (
	"flightplan/pkg/flightplan/domain/event"
	fpl "flightplan/pkg/flightplan/domain/flightplan"
	"flightplan/pkg/flightplan/domain/txmanager"
)

// ManageFlightplanService .
type ManageFlightplanService interface {
	GetFlightplan(requestDpo GetFlightplanRequestDpo, responseDpo GetFlightplanResponseDpo) error
	ListFlightplans(responseEachDpo ListFlightplansResponseDpo) error
	CreateFlightplan(requestDpo CreateFlightplanRequestDpo, responseDpo CreateFlightplanResponseDpo) error
	UpdateFlightplan(requestDpo UpdateFlightplanRequestDpo, responseDpo UpdateFlightplanResponseDpo) error
	DeleteFlightplan(requestDpo DeleteFlightplanRequestDpo) error
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

// CarbonCopyFlightplanRequestDpo .
type CarbonCopyFlightplanRequestDpo interface {
	GetOriginalID() string
	GetNewID() string
}

// NewManageFlightplanService .
func NewManageFlightplanService(
	gen fpl.Generator,
	repo fpl.Repository,
	txm txmanager.TransactionManager,
	psm event.PubSubManager,
) ManageFlightplanService {
	return &manageFlightplanService{
		gen:  gen,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}
}

type manageFlightplanService struct {
	gen  fpl.Generator
	repo fpl.Repository
	txm  txmanager.TransactionManager
	psm  event.PubSubManager
}

func (s *manageFlightplanService) GetFlightplan(
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

func (s *manageFlightplanService) getFlightplanOperation(
	tx txmanager.Tx,
	requestDpo GetFlightplanRequestDpo,
	responseDpo GetFlightplanResponseDpo,
) error {
	flightplan, err := s.repo.GetByID(tx, fpl.ID(requestDpo.GetID()))
	if err != nil {
		return err
	}

	responseDpo(
		string(flightplan.GetID()),
		flightplan.GetName(),
		flightplan.GetDescription(),
	)
	return nil
}

func (s *manageFlightplanService) ListFlightplans(
	responseEachDpo ListFlightplansResponseDpo,
) error {
	return s.txm.Do(func(tx txmanager.Tx) error {
		return s.listFlightplansOperation(
			tx,
			responseEachDpo,
		)
	})
}

func (s *manageFlightplanService) listFlightplansOperation(
	tx txmanager.Tx,
	responseEachDpo ListFlightplansResponseDpo,
) error {
	flightplans, err := s.repo.GetAllOrigin(tx)
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

func (s *manageFlightplanService) CreateFlightplan(
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

func (s *manageFlightplanService) createFlightplanOperation(
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

func (s *manageFlightplanService) UpdateFlightplan(
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

func (s *manageFlightplanService) updateFlightplanOperation(
	tx txmanager.Tx,
	pub event.Publisher,
	requestDpo UpdateFlightplanRequestDpo,
	responseDpo UpdateFlightplanResponseDpo,
) error {
	flightplan, err := s.repo.GetByID(tx, fpl.ID(requestDpo.GetID()))
	if err != nil {
		return err
	}

	if err := flightplan.NameFlightplan(requestDpo.GetName()); err != nil {
		return err
	}
	if err := flightplan.ChangeDescription(requestDpo.GetDescription()); err != nil {
		return err
	}

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

func (s *manageFlightplanService) DeleteFlightplan(
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

func (s *manageFlightplanService) deleteFlightplanOperation(
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

func (s *manageFlightplanService) CarbonCopyFlightplan(
	requestDpo CarbonCopyFlightplanRequestDpo,
) error {
	pub, chClose, err := s.psm.GetPublisher()
	if err != nil {
		return err
	}
	defer chClose()

	return s.txm.DoAndEndHook(
		func(tx txmanager.Tx) error {
			return s.carbonCopyFlightplanOperation(
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

func (s *manageFlightplanService) carbonCopyFlightplanOperation(
	tx txmanager.Tx,
	pub event.Publisher,
	requestDpo CarbonCopyFlightplanRequestDpo,
) error {
	if ret := fpl.CarbonCopyFlightplan(
		tx,
		s.gen,
		s.repo,
		pub,
		fpl.ID(requestDpo.GetOriginalID()),
		fpl.ID(requestDpo.GetNewID()),
	); ret != nil {
		return ret
	}

	return nil
}
