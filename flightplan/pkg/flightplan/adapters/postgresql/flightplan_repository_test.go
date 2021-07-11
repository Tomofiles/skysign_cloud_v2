package postgresql

import (
	"flightplan/pkg/flightplan/adapters/uuid"
	fpl "flightplan/pkg/flightplan/domain/flightplan"
	"regexp"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestFlightplanRepositoryGetSingleWhenGetAll(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightplans"`)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "description", "fleet_id", "version"}).
				AddRow(DefaultFlightplanID, DefaultFlightplanName, DefaultFlightplanDescription, DefaultFlightplanFleetID, DefaultFlightplanVersion),
		)

	gen := uuid.NewFlightplanUUID()
	repository := NewFlightplanRepository(gen)

	flightplans, err := repository.GetAll(db)

	expectFpls := []*fpl.Flightplan{
		fpl.AssembleFrom(
			gen,
			&flightplanComponentMock{
				id:          string(DefaultFlightplanID),
				name:        DefaultFlightplanName,
				description: DefaultFlightplanDescription,
				fleetID:     string(DefaultFlightplanFleetID),
				version:     string(DefaultFlightplanVersion),
			},
		),
	}

	a.Nil(err)
	a.Len(flightplans, 1)
	a.Equal(flightplans, expectFpls)
}

func TestFlightplanRepositoryGetMultipleWhenGetAll(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	const (
		DefaultFlightplanID1          = DefaultFlightplanID + "-1"
		DefaultFlightplanName1        = DefaultFlightplanName + "-1"
		DefaultFlightplanDescription1 = DefaultFlightplanDescription + "-1"
		DefaultFlightplanFleetID1     = DefaultFlightplanFleetID + "-1"
		DefaultFlightplanVersion1     = DefaultFlightplanVersion + "-1"
		DefaultFlightplanID2          = DefaultFlightplanID + "-2"
		DefaultFlightplanName2        = DefaultFlightplanName + "-2"
		DefaultFlightplanDescription2 = DefaultFlightplanDescription + "-2"
		DefaultFlightplanFleetID2     = DefaultFlightplanFleetID + "-2"
		DefaultFlightplanVersion2     = DefaultFlightplanVersion + "-2"
		DefaultFlightplanID3          = DefaultFlightplanID + "-3"
		DefaultFlightplanName3        = DefaultFlightplanName + "-3"
		DefaultFlightplanDescription3 = DefaultFlightplanDescription + "-3"
		DefaultFlightplanFleetID3     = DefaultFlightplanFleetID + "-3"
		DefaultFlightplanVersion3     = DefaultFlightplanVersion + "-3"
	)

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightplans"`)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "description", "fleet_id", "version"}).
				AddRow(DefaultFlightplanID1, DefaultFlightplanName1, DefaultFlightplanDescription1, DefaultFlightplanFleetID1, DefaultFlightplanVersion1).
				AddRow(DefaultFlightplanID2, DefaultFlightplanName2, DefaultFlightplanDescription2, DefaultFlightplanFleetID2, DefaultFlightplanVersion2).
				AddRow(DefaultFlightplanID3, DefaultFlightplanName3, DefaultFlightplanDescription3, DefaultFlightplanFleetID3, DefaultFlightplanVersion3),
		)

	gen := uuid.NewFlightplanUUID()
	repository := NewFlightplanRepository(gen)

	flightplans, err := repository.GetAll(db)

	expectFpls := []*fpl.Flightplan{
		fpl.AssembleFrom(
			gen,
			&flightplanComponentMock{
				id:          string(DefaultFlightplanID1),
				name:        DefaultFlightplanName1,
				description: DefaultFlightplanDescription1,
				fleetID:     string(DefaultFlightplanFleetID1),
				version:     string(DefaultFlightplanVersion1),
			},
		),
		fpl.AssembleFrom(
			gen,
			&flightplanComponentMock{
				id:          string(DefaultFlightplanID2),
				name:        DefaultFlightplanName2,
				description: DefaultFlightplanDescription2,
				fleetID:     string(DefaultFlightplanFleetID2),
				version:     string(DefaultFlightplanVersion2),
			},
		),
		fpl.AssembleFrom(
			gen,
			&flightplanComponentMock{
				id:          string(DefaultFlightplanID3),
				name:        DefaultFlightplanName3,
				description: DefaultFlightplanDescription3,
				fleetID:     string(DefaultFlightplanFleetID3),
				version:     string(DefaultFlightplanVersion3),
			},
		),
	}

	a.Nil(err)
	a.Len(flightplans, 3)
	a.Equal(flightplans, expectFpls)
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
			regexp.QuoteMeta(`SELECT * FROM "flightplans"`)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "description", "fleet_id", "version"}),
		)

	gen := uuid.NewFlightplanUUID()
	repository := NewFlightplanRepository(gen)

	flightplans, err := repository.GetAll(db)

	var expectFpls []*fpl.Flightplan

	a.Nil(err)
	a.Len(flightplans, 0)
	a.Equal(flightplans, expectFpls)
}

func TestFlightplanRepositoryGetByID(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightplans" WHERE id = $1`)).
		WithArgs(DefaultFlightplanID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "description", "fleet_id", "version"}).
				AddRow(DefaultFlightplanID, DefaultFlightplanName, DefaultFlightplanDescription, DefaultFlightplanFleetID, DefaultFlightplanVersion),
		)

	gen := uuid.NewFlightplanUUID()
	repository := NewFlightplanRepository(gen)

	flightplan, err := repository.GetByID(db, DefaultFlightplanID)

	expectFpl := fpl.AssembleFrom(
		gen,
		&flightplanComponentMock{
			id:          string(DefaultFlightplanID),
			name:        DefaultFlightplanName,
			description: DefaultFlightplanDescription,
			fleetID:     string(DefaultFlightplanFleetID),
			version:     string(DefaultFlightplanVersion),
		},
	)

	a.Nil(err)
	a.Equal(flightplan, expectFpl)
}

func TestFlightplanRepositoryNotFoundWhenGetByID(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightplans" WHERE id = $1`)).
		WithArgs(DefaultFlightplanID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "description", "fleet_id", "version"}),
		)

	gen := uuid.NewFlightplanUUID()
	repository := NewFlightplanRepository(gen)

	flightplan, err := repository.GetByID(db, DefaultFlightplanID)

	a.Nil(flightplan)
	a.Equal(err, fpl.ErrNotFound)
}

func TestFlightplanRepositoryCreateSave(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightplans" WHERE id = $1`)).
		WithArgs(DefaultFlightplanID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "description", "fleet_id", "version"}),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "flightplans" ("id","name","description","fleet_id","version") VALUES ($1,$2,$3,$4,$5)`)).
		WithArgs(DefaultFlightplanID, DefaultFlightplanName, DefaultFlightplanDescription, DefaultFlightplanFleetID, DefaultFlightplanVersion).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	gen := uuid.NewFlightplanUUID()
	repository := NewFlightplanRepository(gen)

	flightplan := fpl.AssembleFrom(
		gen,
		&flightplanComponentMock{
			id:          string(DefaultFlightplanID),
			name:        DefaultFlightplanName,
			description: DefaultFlightplanDescription,
			fleetID:     string(DefaultFlightplanFleetID),
			version:     string(DefaultFlightplanVersion),
		},
	)

	err = repository.Save(db, flightplan)

	a.Nil(err)
}

func TestFlightplanRepositoryUpdateSave(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	const (
		AfterName              = DefaultFlightplanName + "-after"
		AfterDescription       = DefaultFlightplanDescription + "-after"
		AfterFlightplanFleetID = DefaultFlightplanFleetID + "-after"
		AfterVersion           = DefaultFlightplanVersion + "-after"
	)

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightplans" WHERE id = $1`)).
		WithArgs(DefaultFlightplanID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "description", "fleet_id", "version"}).
				AddRow(DefaultFlightplanID, DefaultFlightplanName, DefaultFlightplanDescription, DefaultFlightplanFleetID, DefaultFlightplanVersion),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`UPDATE "flightplans" SET "name"=$1,"description"=$2,"fleet_id"=$3,"version"=$4 WHERE "id" = $5`)).
		WithArgs(AfterName, AfterDescription, AfterFlightplanFleetID, AfterVersion, DefaultFlightplanID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	gen := uuid.NewFlightplanUUID()
	repository := NewFlightplanRepository(gen)

	flightplan := fpl.AssembleFrom(
		gen,
		&flightplanComponentMock{
			id:          string(DefaultFlightplanID),
			name:        AfterName,
			description: AfterDescription,
			fleetID:     string(AfterFlightplanFleetID),
			version:     string(AfterVersion),
		},
	)

	err = repository.Save(db, flightplan)

	a.Nil(err)
}

func TestFlightplanRepositoryDelete(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightplans" WHERE id = $1`)).
		WithArgs(DefaultFlightplanID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "description", "fleet_id", "version"}).
				AddRow(DefaultFlightplanID, DefaultFlightplanName, DefaultFlightplanDescription, DefaultFlightplanFleetID, DefaultFlightplanVersion),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`DELETE FROM "flightplans" WHERE "flightplans"."id" = $1`)).
		WithArgs(DefaultFlightplanID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	gen := uuid.NewFlightplanUUID()
	repository := NewFlightplanRepository(gen)

	err = repository.Delete(db, DefaultFlightplanID)

	a.Nil(err)
}

func TestFlightplanRepositoryNotFoundWhenDelete(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightplans" WHERE id = $1`)).
		WithArgs(DefaultFlightplanID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "description", "fleet_id", "version"}),
		)

	gen := uuid.NewFlightplanUUID()
	repository := NewFlightplanRepository(gen)

	err = repository.Delete(db, DefaultFlightplanID)

	a.Equal(err, fpl.ErrNotFound)
}
