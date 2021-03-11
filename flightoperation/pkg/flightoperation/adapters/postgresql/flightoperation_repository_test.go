package postgresql

import (
	"flightoperation/pkg/flightoperation/adapters/uuid"
	fope "flightoperation/pkg/flightoperation/domain/flightoperation"
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
			sqlmock.NewRows([]string{"id", "flightplan_id", "is_completed", "version"}).
				AddRow(DefaultFlightoperationID, DefaultFlightplanID, DefaultIsCompleted, DefaultVersion),
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
				isCompleted:  DefaultIsCompleted,
				version:      string(DefaultVersion),
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
		DefaultVersion1           = DefaultVersion + "-1"
		DefaultVersion2           = DefaultVersion + "-2"
		DefaultVersion3           = DefaultVersion + "-3"
	)

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightoperations"`)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "flightplan_id", "is_completed", "version"}).
				AddRow(DefaultFlightoperationID1, DefaultFlightplanID1, DefaultIsCompleted, DefaultVersion1).
				AddRow(DefaultFlightoperationID2, DefaultFlightplanID2, DefaultIsCompleted, DefaultVersion2).
				AddRow(DefaultFlightoperationID3, DefaultFlightplanID3, DefaultIsCompleted, DefaultVersion3),
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
				isCompleted:  DefaultIsCompleted,
				version:      string(DefaultVersion1),
			},
		),
		fope.AssembleFrom(
			gen,
			&flightoperationComponentMock{
				id:           string(DefaultFlightoperationID2),
				flightplanID: string(DefaultFlightplanID2),
				isCompleted:  DefaultIsCompleted,
				version:      string(DefaultVersion2),
			},
		),
		fope.AssembleFrom(
			gen,
			&flightoperationComponentMock{
				id:           string(DefaultFlightoperationID3),
				flightplanID: string(DefaultFlightplanID3),
				isCompleted:  DefaultIsCompleted,
				version:      string(DefaultVersion3),
			},
		),
	}

	a.Nil(err)
	a.Len(flightoperations, 3)
	a.Equal(flightoperations, expectFopes)
}

func TestFlightoperationRepositoryGetNoneWhenGetAll(t *testing.T) {
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
			sqlmock.NewRows([]string{"id", "flightplan_id", "is_completed", "version"}),
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
			sqlmock.NewRows([]string{"id", "flightplan_id", "is_completed", "version"}).
				AddRow(DefaultFlightoperationID, DefaultFlightplanID, DefaultIsCompleted, DefaultVersion),
		)

	gen := uuid.NewFlightoperationUUID()
	repository := NewFlightoperationRepository(gen)

	flightoperation, err := repository.GetByID(db, DefaultFlightoperationID)

	expectFope := fope.AssembleFrom(
		gen,
		&flightoperationComponentMock{
			id:           string(DefaultFlightoperationID),
			flightplanID: string(DefaultFlightplanID),
			isCompleted:  DefaultIsCompleted,
			version:      string(DefaultVersion),
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
			sqlmock.NewRows([]string{"id", "flightplan_id", "is_completed", "version"}),
		)

	gen := uuid.NewFlightoperationUUID()
	repository := NewFlightoperationRepository(gen)

	flightoperation, err := repository.GetByID(db, DefaultFlightoperationID)

	a.Nil(flightoperation)
	a.Equal(err, fope.ErrNotFound)
}

func TestFlightoperationRepositoryCreateSave(t *testing.T) {
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
			sqlmock.NewRows([]string{"id", "flightplan_id", "is_completed", "version"}),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "flightoperations" ("id","flightplan_id","is_completed","version") VALUES ($1,$2,$3,$4)`)).
		WithArgs(DefaultFlightoperationID, DefaultFlightplanID, DefaultIsCompleted, DefaultVersion).
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
			isCompleted:  DefaultIsCompleted,
			version:      string(DefaultVersion),
		},
	)

	err = repository.Save(db, flightoperation)

	a.Nil(err)
}

func TestFlightoperationRepositoryUpdateSave(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	const (
		AfterFlightplanID = DefaultFlightplanID + "-after"
		AfterIsCompleted  = !DefaultIsCompleted
		AfterVersion      = DefaultVersion + "-after"
	)

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightoperations" WHERE id = $1`)).
		WithArgs(DefaultFlightoperationID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "flightplan_id", "is_completed", "version"}).
				AddRow(DefaultFlightoperationID, DefaultFlightplanID, DefaultIsCompleted, DefaultVersion),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`UPDATE "flightoperations" SET "flightplan_id"=$1,"is_completed"=$2,"version"=$3 WHERE "id" = $4`)).
		WithArgs(AfterFlightplanID, AfterIsCompleted, AfterVersion, DefaultFlightoperationID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	gen := uuid.NewFlightoperationUUID()
	repository := NewFlightoperationRepository(gen)

	flightoperation := fope.AssembleFrom(
		gen,
		&flightoperationComponentMock{
			id:           string(DefaultFlightoperationID),
			flightplanID: string(AfterFlightplanID),
			isCompleted:  AfterIsCompleted,
			version:      string(AfterVersion),
		},
	)

	err = repository.Save(db, flightoperation)

	a.Nil(err)
}
