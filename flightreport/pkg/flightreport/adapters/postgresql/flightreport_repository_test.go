package postgresql

import (
	"flightreport/pkg/flightreport/adapters/uuid"
	frep "flightreport/pkg/flightreport/domain/flightreport"
	"regexp"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestFlightreportRepositoryGetSingleWhenGetAll(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightreports"`)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "description", "fleet_id"}).
				AddRow(DefaultID, DefaultName, DefaultDescription, DefaultFleetID),
		)

	gen := uuid.NewFlightreportUUID()
	repository := NewFlightreportRepository(gen)

	flightreports, err := repository.GetAll(db)

	expectFopes := []*frep.Flightreport{
		frep.AssembleFrom(
			gen,
			&flightreportComponentMock{
				id:          string(DefaultID),
				name:        DefaultName,
				description: DefaultDescription,
				fleetID:     string(DefaultFleetID),
			},
		),
	}

	a.Nil(err)
	a.Len(flightreports, 1)
	a.Equal(flightreports, expectFopes)
}

func TestFlightreportRepositoryGetMultipleWhenGetAll(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	const (
		DefaultID1          = DefaultID + "-1"
		DefaultName1        = DefaultName + "-1"
		DefaultDescription1 = DefaultDescription + "-1"
		DefaultFleetID1     = DefaultFleetID + "-1"
		DefaultID2          = DefaultID + "-2"
		DefaultName2        = DefaultName + "-2"
		DefaultDescription2 = DefaultDescription + "-2"
		DefaultFleetID2     = DefaultFleetID + "-2"
		DefaultID3          = DefaultID + "-3"
		DefaultName3        = DefaultName + "-3"
		DefaultDescription3 = DefaultDescription + "-3"
		DefaultFleetID3     = DefaultFleetID + "-3"
	)

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightreports"`)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "description", "fleet_id"}).
				AddRow(DefaultID1, DefaultName1, DefaultDescription1, DefaultFleetID1).
				AddRow(DefaultID2, DefaultName2, DefaultDescription2, DefaultFleetID2).
				AddRow(DefaultID3, DefaultName3, DefaultDescription3, DefaultFleetID3),
		)

	gen := uuid.NewFlightreportUUID()
	repository := NewFlightreportRepository(gen)

	flightreports, err := repository.GetAll(db)

	expectFopes := []*frep.Flightreport{
		frep.AssembleFrom(
			gen,
			&flightreportComponentMock{
				id:          string(DefaultID1),
				name:        DefaultName1,
				description: DefaultDescription1,
				fleetID:     string(DefaultFleetID1),
			},
		),
		frep.AssembleFrom(
			gen,
			&flightreportComponentMock{
				id:          string(DefaultID2),
				name:        DefaultName2,
				description: DefaultDescription2,
				fleetID:     string(DefaultFleetID2),
			},
		),
		frep.AssembleFrom(
			gen,
			&flightreportComponentMock{
				id:          string(DefaultID3),
				name:        DefaultName3,
				description: DefaultDescription3,
				fleetID:     string(DefaultFleetID3),
			},
		),
	}

	a.Nil(err)
	a.Len(flightreports, 3)
	a.Equal(flightreports, expectFopes)
}

func TestFlightreportRepositoryGetNoneWhenGetAll(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightreports"`)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "description", "fleet_id"}),
		)

	gen := uuid.NewFlightreportUUID()
	repository := NewFlightreportRepository(gen)

	flightreports, err := repository.GetAll(db)

	var expectFopes []*frep.Flightreport

	a.Nil(err)
	a.Len(flightreports, 0)
	a.Equal(flightreports, expectFopes)
}

func TestFlightreportRepositoryGetByID(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightreports" WHERE id = $1`)).
		WithArgs(DefaultID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "description", "fleet_id"}).
				AddRow(DefaultID, DefaultName, DefaultDescription, DefaultFleetID),
		)

	gen := uuid.NewFlightreportUUID()
	repository := NewFlightreportRepository(gen)

	flightreport, err := repository.GetByID(db, DefaultID)

	expectFope := frep.AssembleFrom(
		gen,
		&flightreportComponentMock{
			id:          string(DefaultID),
			name:        DefaultName,
			description: DefaultDescription,
			fleetID:     string(DefaultFleetID),
		},
	)

	a.Nil(err)
	a.Equal(flightreport, expectFope)
}

func TestFlightreportRepositoryNotFoundWhenGetByID(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightreports" WHERE id = $1`)).
		WithArgs(DefaultID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "description", "fleet_id"}),
		)

	gen := uuid.NewFlightreportUUID()
	repository := NewFlightreportRepository(gen)

	flightreport, err := repository.GetByID(db, DefaultID)

	a.Nil(flightreport)
	a.Equal(err, frep.ErrNotFound)
}

func TestFlightreportRepositorySave(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightreports" WHERE id = $1`)).
		WithArgs(DefaultID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "description", "fleet_id"}),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "flightreports" ("id","name","description","fleet_id") VALUES ($1,$2,$3,$4)`)).
		WithArgs(DefaultID, DefaultName, DefaultDescription, DefaultFleetID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	gen := uuid.NewFlightreportUUID()
	repository := NewFlightreportRepository(gen)

	flightreport := frep.AssembleFrom(
		gen,
		&flightreportComponentMock{
			id:          string(DefaultID),
			name:        DefaultName,
			description: DefaultDescription,
			fleetID:     string(DefaultFleetID),
		},
	)

	err = repository.Save(db, flightreport)

	a.Nil(err)
}

func TestFlightreportRepositoryUpdateSkipWhenAlreadyExistWhenSave(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightreports" WHERE id = $1`)).
		WithArgs(DefaultID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "description", "fleet_id"}).
				AddRow(DefaultID, DefaultName, DefaultDescription, DefaultFleetID),
		)

	gen := uuid.NewFlightreportUUID()
	repository := NewFlightreportRepository(gen)

	flightreport := frep.AssembleFrom(
		gen,
		&flightreportComponentMock{
			id:          string(DefaultID),
			name:        DefaultName,
			description: DefaultDescription,
			fleetID:     string(DefaultFleetID),
		},
	)

	err = repository.Save(db, flightreport)

	a.Nil(err)
}
