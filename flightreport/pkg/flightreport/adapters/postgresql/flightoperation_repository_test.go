package postgresql

import (
	"flightreport/pkg/flightreport/adapters/uuid"
	fope "flightreport/pkg/flightreport/domain/flightoperation"
	"regexp"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestFlightoperationRepositoryGetSingleWhenGetAll(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightoperations"`)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "flightplan_id"}).
				AddRow(DefaultFlightoperationID, DefaultFlightplanID),
		)

	gen := uuid.NewFlightoperationUUID()
	repository := NewFlightoperationRepository(gen)

	flightoperations, err := repository.GetAll(db)

	expectFopes := []*fope.Flightoperation{
		fope.AssembleFrom(
			gen,
			&flightoperationComponentMock{
				id:           string(DefaultFlightoperationID),
				flightplanID: string(DefaultFlightplanID),
			},
		),
	}

	a.Nil(err)
	a.Len(flightoperations, 1)
	a.Equal(flightoperations, expectFopes)
}

func TestFlightoperationRepositoryGetMultipleWhenGetAll(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	const (
		DefaultFlightoperationID1 = DefaultFlightoperationID + "-1"
		DefaultFlightoperationID2 = DefaultFlightoperationID + "-2"
		DefaultFlightoperationID3 = DefaultFlightoperationID + "-3"
		DefaultFlightplanID1      = DefaultFlightplanID + "-1"
		DefaultFlightplanID2      = DefaultFlightplanID + "-2"
		DefaultFlightplanID3      = DefaultFlightplanID + "-3"
	)

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightoperations"`)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "flightplan_id"}).
				AddRow(DefaultFlightoperationID1, DefaultFlightplanID1).
				AddRow(DefaultFlightoperationID2, DefaultFlightplanID2).
				AddRow(DefaultFlightoperationID3, DefaultFlightplanID3),
		)

	gen := uuid.NewFlightoperationUUID()
	repository := NewFlightoperationRepository(gen)

	flightoperations, err := repository.GetAll(db)

	expectFopes := []*fope.Flightoperation{
		fope.AssembleFrom(
			gen,
			&flightoperationComponentMock{
				id:           string(DefaultFlightoperationID1),
				flightplanID: string(DefaultFlightplanID1),
			},
		),
		fope.AssembleFrom(
			gen,
			&flightoperationComponentMock{
				id:           string(DefaultFlightoperationID2),
				flightplanID: string(DefaultFlightplanID2),
			},
		),
		fope.AssembleFrom(
			gen,
			&flightoperationComponentMock{
				id:           string(DefaultFlightoperationID3),
				flightplanID: string(DefaultFlightplanID3),
			},
		),
	}

	a.Nil(err)
	a.Len(flightoperations, 3)
	a.Equal(flightoperations, expectFopes)
}

func TestFlightplanRepositoryGetNoneWhenGetAll(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightoperations"`)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "flightplan_id"}),
		)

	gen := uuid.NewFlightoperationUUID()
	repository := NewFlightoperationRepository(gen)

	flightoperations, err := repository.GetAll(db)

	var expectFopes []*fope.Flightoperation

	a.Nil(err)
	a.Len(flightoperations, 0)
	a.Equal(flightoperations, expectFopes)
}

func TestFlightoperationRepositoryGetByID(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightoperations" WHERE id = $1`)).
		WithArgs(DefaultFlightoperationID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "flightplan_id"}).
				AddRow(DefaultFlightoperationID, DefaultFlightplanID),
		)

	gen := uuid.NewFlightoperationUUID()
	repository := NewFlightoperationRepository(gen)

	flightoperation, err := repository.GetByID(db, DefaultFlightoperationID)

	expectFope := fope.AssembleFrom(
		gen,
		&flightoperationComponentMock{
			id:           string(DefaultFlightoperationID),
			flightplanID: string(DefaultFlightplanID),
		},
	)

	a.Nil(err)
	a.Equal(flightoperation, expectFope)
}

func TestFlightoperationRepositoryNotFoundWhenGetByID(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightoperations" WHERE id = $1`)).
		WithArgs(DefaultFlightoperationID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "flightplan_id"}),
		)

	gen := uuid.NewFlightoperationUUID()
	repository := NewFlightoperationRepository(gen)

	flightoperation, err := repository.GetByID(db, DefaultFlightoperationID)

	a.Nil(flightoperation)
	a.Equal(err, fope.ErrNotFound)
}

func TestFlightoperationRepositorySave(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightoperations" WHERE id = $1`)).
		WithArgs(DefaultFlightoperationID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "flightplan_id"}),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "flightoperations" ("id","flightplan_id") VALUES ($1,$2)`)).
		WithArgs(DefaultFlightoperationID, DefaultFlightplanID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	gen := uuid.NewFlightoperationUUID()
	repository := NewFlightoperationRepository(gen)

	flightoperation := fope.AssembleFrom(
		gen,
		&flightoperationComponentMock{
			id:           string(DefaultFlightoperationID),
			flightplanID: string(DefaultFlightplanID),
		},
	)

	err = repository.Save(db, flightoperation)

	a.Nil(err)
}

func TestFlightoperationRepositoryUpdateSkipWhenAlreadyExistWhenSave(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightoperations" WHERE id = $1`)).
		WithArgs(DefaultFlightoperationID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "flightplan_id"}).
				AddRow(DefaultFlightoperationID, DefaultFlightplanID),
		)

	gen := uuid.NewFlightoperationUUID()
	repository := NewFlightoperationRepository(gen)

	flightoperation := fope.AssembleFrom(
		gen,
		&flightoperationComponentMock{
			id:           string(DefaultFlightoperationID),
			flightplanID: string(DefaultFlightplanID),
		},
	)

	err = repository.Save(db, flightoperation)

	a.Nil(err)
}
