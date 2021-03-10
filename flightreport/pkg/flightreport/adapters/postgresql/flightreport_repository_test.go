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
			sqlmock.NewRows([]string{"id", "flightoperation_id"}).
				AddRow(DefaultFlightreportID, DefaultFlightoperationID),
		)

	gen := uuid.NewFlightreportUUID()
	repository := NewFlightreportRepository(gen)

	flightreports, err := repository.GetAll(db)

	expectFopes := []*frep.Flightreport{
		frep.AssembleFrom(
			gen,
			&flightreportComponentMock{
				id:                string(DefaultFlightreportID),
				flightoperationID: string(DefaultFlightoperationID),
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
		DefaultFlightreportID1    = DefaultFlightreportID + "-1"
		DefaultFlightreportID2    = DefaultFlightreportID + "-2"
		DefaultFlightreportID3    = DefaultFlightreportID + "-3"
		DefaultFlightoperationID1 = DefaultFlightoperationID + "-1"
		DefaultFlightoperationID2 = DefaultFlightoperationID + "-2"
		DefaultFlightoperationID3 = DefaultFlightoperationID + "-3"
	)

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightreports"`)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "flightoperation_id"}).
				AddRow(DefaultFlightreportID1, DefaultFlightoperationID1).
				AddRow(DefaultFlightreportID2, DefaultFlightoperationID2).
				AddRow(DefaultFlightreportID3, DefaultFlightoperationID3),
		)

	gen := uuid.NewFlightreportUUID()
	repository := NewFlightreportRepository(gen)

	flightreports, err := repository.GetAll(db)

	expectFopes := []*frep.Flightreport{
		frep.AssembleFrom(
			gen,
			&flightreportComponentMock{
				id:                string(DefaultFlightreportID1),
				flightoperationID: string(DefaultFlightoperationID1),
			},
		),
		frep.AssembleFrom(
			gen,
			&flightreportComponentMock{
				id:                string(DefaultFlightreportID2),
				flightoperationID: string(DefaultFlightoperationID2),
			},
		),
		frep.AssembleFrom(
			gen,
			&flightreportComponentMock{
				id:                string(DefaultFlightreportID3),
				flightoperationID: string(DefaultFlightoperationID3),
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
			sqlmock.NewRows([]string{"id", "flightoperation_id"}),
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
		WithArgs(DefaultFlightreportID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "flightoperation_id"}).
				AddRow(DefaultFlightreportID, DefaultFlightoperationID),
		)

	gen := uuid.NewFlightreportUUID()
	repository := NewFlightreportRepository(gen)

	flightreport, err := repository.GetByID(db, DefaultFlightreportID)

	expectFope := frep.AssembleFrom(
		gen,
		&flightreportComponentMock{
			id:                string(DefaultFlightreportID),
			flightoperationID: string(DefaultFlightoperationID),
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
		WithArgs(DefaultFlightreportID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "flightoperation_id"}),
		)

	gen := uuid.NewFlightreportUUID()
	repository := NewFlightreportRepository(gen)

	flightreport, err := repository.GetByID(db, DefaultFlightreportID)

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
		WithArgs(DefaultFlightreportID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "flightoperation_id"}),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "flightreports" ("id","flightoperation_id") VALUES ($1,$2)`)).
		WithArgs(DefaultFlightreportID, DefaultFlightoperationID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	gen := uuid.NewFlightreportUUID()
	repository := NewFlightreportRepository(gen)

	flightreport := frep.AssembleFrom(
		gen,
		&flightreportComponentMock{
			id:                string(DefaultFlightreportID),
			flightoperationID: string(DefaultFlightoperationID),
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
		WithArgs(DefaultFlightreportID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "flightoperation_id"}).
				AddRow(DefaultFlightreportID, DefaultFlightoperationID),
		)

	gen := uuid.NewFlightreportUUID()
	repository := NewFlightreportRepository(gen)

	flightreport := frep.AssembleFrom(
		gen,
		&flightreportComponentMock{
			id:                string(DefaultFlightreportID),
			flightoperationID: string(DefaultFlightoperationID),
		},
	)

	err = repository.Save(db, flightreport)

	a.Nil(err)
}
