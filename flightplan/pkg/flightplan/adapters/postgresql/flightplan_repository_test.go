package postgresql

import (
	"flightplan/pkg/flightplan/adapters/uuid"
	fpl "flightplan/pkg/flightplan/domain/flightplan"
	"regexp"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetSingleWhenGetAll(t *testing.T) {
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
			sqlmock.NewRows([]string{"id", "name", "description", "version"}).
				AddRow(DefaultID, DefaultName, DefaultDescription, DefaultVersion),
		)

	gen := uuid.NewFlightplanUUID()
	repository := NewFlightplanRepository(gen)

	flightplans, err := repository.GetAll(db)

	expectFpls := []*fpl.Flightplan{
		fpl.AssembleFrom(
			gen,
			&flightplanComponentMock{
				id:          string(DefaultID),
				name:        DefaultName,
				description: DefaultDescription,
				version:     string(DefaultVersion),
			},
		),
	}

	a.Nil(err)
	a.Len(flightplans, 1)
	a.Equal(flightplans, expectFpls)
}

func TestGetMultipleWhenGetAll(t *testing.T) {
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
		DefaultVersion1     = DefaultVersion + "-1"
		DefaultID2          = DefaultID + "-2"
		DefaultName2        = DefaultName + "-2"
		DefaultDescription2 = DefaultDescription + "-2"
		DefaultVersion2     = DefaultVersion + "-2"
		DefaultID3          = DefaultID + "-3"
		DefaultName3        = DefaultName + "-3"
		DefaultDescription3 = DefaultDescription + "-3"
		DefaultVersion3     = DefaultVersion + "-3"
	)

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightplans"`)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "description", "version"}).
				AddRow(DefaultID1, DefaultName1, DefaultDescription1, DefaultVersion1).
				AddRow(DefaultID2, DefaultName2, DefaultDescription2, DefaultVersion2).
				AddRow(DefaultID3, DefaultName3, DefaultDescription3, DefaultVersion3),
		)

	gen := uuid.NewFlightplanUUID()
	repository := NewFlightplanRepository(gen)

	flightplans, err := repository.GetAll(db)

	expectFpls := []*fpl.Flightplan{
		fpl.AssembleFrom(
			gen,
			&flightplanComponentMock{
				id:          string(DefaultID1),
				name:        DefaultName1,
				description: DefaultDescription1,
				version:     string(DefaultVersion1),
			},
		),
		fpl.AssembleFrom(
			gen,
			&flightplanComponentMock{
				id:          string(DefaultID2),
				name:        DefaultName2,
				description: DefaultDescription2,
				version:     string(DefaultVersion2),
			},
		),
		fpl.AssembleFrom(
			gen,
			&flightplanComponentMock{
				id:          string(DefaultID3),
				name:        DefaultName3,
				description: DefaultDescription3,
				version:     string(DefaultVersion3),
			},
		),
	}

	a.Nil(err)
	a.Len(flightplans, 3)
	a.Equal(flightplans, expectFpls)
}

func TestGetNoneWhenGetAll(t *testing.T) {
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
			sqlmock.NewRows([]string{"id", "name", "description", "version"}),
		)

	gen := uuid.NewFlightplanUUID()
	repository := NewFlightplanRepository(gen)

	flightplans, err := repository.GetAll(db)

	var expectFpls []*fpl.Flightplan

	a.Nil(err)
	a.Len(flightplans, 0)
	a.Equal(flightplans, expectFpls)
}

func TestGetErrorWhenGetAll(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightplans"`)).
		WillReturnError(
			fpl.ErrGet,
		)

	gen := uuid.NewFlightplanUUID()
	repository := NewFlightplanRepository(gen)

	flightplans, err := repository.GetAll(db)

	a.Nil(flightplans)
	a.Equal(err, fpl.ErrGet)
}

func TestGetByID(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightplans" WHERE id = $1`)).
		WithArgs(DefaultID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "description", "version"}).
				AddRow(DefaultID, DefaultName, DefaultDescription, DefaultVersion),
		)

	gen := uuid.NewFlightplanUUID()
	repository := NewFlightplanRepository(gen)

	flightplan, err := repository.GetByID(db, DefaultID)

	expectFpl := fpl.AssembleFrom(
		gen,
		&flightplanComponentMock{
			id:          string(DefaultID),
			name:        DefaultName,
			description: DefaultDescription,
			version:     string(DefaultVersion),
		},
	)

	a.Nil(err)
	a.Equal(flightplan, expectFpl)
}

func TestGetErrorWhenGetByID(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightplans" WHERE id = $1`)).
		WithArgs(DefaultID).
		WillReturnError(
			fpl.ErrGet,
		)

	gen := uuid.NewFlightplanUUID()
	repository := NewFlightplanRepository(gen)

	flightplan, err := repository.GetByID(db, DefaultID)

	a.Nil(flightplan)
	a.Equal(err, fpl.ErrGet)
}

func TestNotFoundErrorWhenGetByID(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightplans" WHERE id = $1`)).
		WithArgs(DefaultID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "description", "version"}),
		)

	gen := uuid.NewFlightplanUUID()
	repository := NewFlightplanRepository(gen)

	flightplan, err := repository.GetByID(db, DefaultID)

	a.Nil(flightplan)
	a.Nil(err)
}

func TestCreateSave(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightplans" WHERE id = $1`)).
		WithArgs(DefaultID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "description", "version"}),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "flightplans" ("id","name","description","version") VALUES ($1,$2,$3,$4)`)).
		WithArgs(DefaultID, DefaultName, DefaultDescription, DefaultVersion).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	gen := uuid.NewFlightplanUUID()
	repository := NewFlightplanRepository(gen)

	flightplan := fpl.AssembleFrom(
		gen,
		&flightplanComponentMock{
			id:          string(DefaultID),
			name:        DefaultName,
			description: DefaultDescription,
			version:     string(DefaultVersion),
		},
	)

	err = repository.Save(db, flightplan)

	a.Nil(err)
}

func TestCreateErrorWhenSave(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightplans" WHERE id = $1`)).
		WithArgs(DefaultID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "description", "version"}),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "flightplans" ("id","name","description","version") VALUES ($1,$2,$3,$4)`)).
		WithArgs(DefaultID, DefaultName, DefaultDescription, DefaultVersion).
		WillReturnError(
			fpl.ErrSave,
		)

	gen := uuid.NewFlightplanUUID()
	repository := NewFlightplanRepository(gen)

	flightplan := fpl.AssembleFrom(
		gen,
		&flightplanComponentMock{
			id:          string(DefaultID),
			name:        DefaultName,
			description: DefaultDescription,
			version:     string(DefaultVersion),
		},
	)

	err = repository.Save(db, flightplan)

	a.Equal(err, fpl.ErrSave)
}

func TestUpdateSave(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	const (
		AfterName        = DefaultName + "-after"
		AfterDescription = DefaultDescription + "-after"
		AfterVersion     = DefaultVersion + "-after"
	)

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightplans" WHERE id = $1`)).
		WithArgs(DefaultID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "description", "version"}).
				AddRow(DefaultID, DefaultName, DefaultDescription, DefaultVersion),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`UPDATE "flightplans" SET "name"=$1,"description"=$2,"version"=$3 WHERE "id" = $4`)).
		WithArgs(AfterName, AfterDescription, AfterVersion, DefaultID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	gen := uuid.NewFlightplanUUID()
	repository := NewFlightplanRepository(gen)

	flightplan := fpl.AssembleFrom(
		gen,
		&flightplanComponentMock{
			id:          string(DefaultID),
			name:        AfterName,
			description: AfterDescription,
			version:     string(AfterVersion),
		},
	)

	err = repository.Save(db, flightplan)

	a.Nil(err)
}

func TestUpdateErrorWhenSave(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	const (
		AfterName        = DefaultName + "-after"
		AfterDescription = DefaultDescription + "-after"
		AfterVersion     = DefaultVersion + "-after"
	)

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightplans" WHERE id = $1`)).
		WithArgs(DefaultID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "description", "version"}).
				AddRow(DefaultID, DefaultName, DefaultDescription, DefaultVersion),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`UPDATE "flightplans" SET "name"=$1,"description"=$2,"version"=$3 WHERE "id" = $4`)).
		WithArgs(AfterName, AfterDescription, AfterVersion, DefaultID).
		WillReturnError(
			fpl.ErrSave,
		)

	gen := uuid.NewFlightplanUUID()
	repository := NewFlightplanRepository(gen)

	flightplan := fpl.AssembleFrom(
		gen,
		&flightplanComponentMock{
			id:          string(DefaultID),
			name:        AfterName,
			description: AfterDescription,
			version:     string(AfterVersion),
		},
	)

	err = repository.Save(db, flightplan)

	a.Equal(err, fpl.ErrSave)
}

func TestGetErrorWhenSave(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightplans" WHERE id = $1`)).
		WithArgs(DefaultID).
		WillReturnError(
			fpl.ErrGet,
		)

	gen := uuid.NewFlightplanUUID()
	repository := NewFlightplanRepository(gen)

	flightplan := fpl.AssembleFrom(
		gen,
		&flightplanComponentMock{
			id:          string(DefaultID),
			name:        DefaultName,
			description: DefaultDescription,
			version:     string(DefaultVersion),
		},
	)

	err = repository.Save(db, flightplan)

	a.Equal(err, fpl.ErrGet)
}

func TestDelete(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightplans" WHERE id = $1`)).
		WithArgs(DefaultID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "description", "version"}).
				AddRow(DefaultID, DefaultName, DefaultDescription, DefaultVersion),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`DELETE FROM "flightplans" WHERE "flightplans"."id" = $1`)).
		WithArgs(DefaultID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	gen := uuid.NewFlightplanUUID()
	repository := NewFlightplanRepository(gen)

	err = repository.Delete(db, DefaultID)

	a.Nil(err)
}

func TestGetErrorWhenDelete(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightplans" WHERE id = $1`)).
		WithArgs(DefaultID).
		WillReturnError(
			fpl.ErrGet,
		)

	gen := uuid.NewFlightplanUUID()
	repository := NewFlightplanRepository(gen)

	err = repository.Delete(db, DefaultID)

	a.Equal(err, fpl.ErrGet)
}

func TestNotFoundWhenDelete(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightplans" WHERE id = $1`)).
		WithArgs(DefaultID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "description", "version"}),
		)

	gen := uuid.NewFlightplanUUID()
	repository := NewFlightplanRepository(gen)

	err = repository.Delete(db, DefaultID)

	a.Nil(err)
}

func TestDeleteErrorWhenDelete(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightplans" WHERE id = $1`)).
		WithArgs(DefaultID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "description", "version"}).
				AddRow(DefaultID, DefaultName, DefaultDescription, DefaultVersion),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`DELETE FROM "flightplans" WHERE "flightplans"."id" = $1`)).
		WithArgs(DefaultID).
		WillReturnError(
			fpl.ErrDelete,
		)

	gen := uuid.NewFlightplanUUID()
	repository := NewFlightplanRepository(gen)

	err = repository.Delete(db, DefaultID)

	a.Equal(err, fpl.ErrDelete)
}
