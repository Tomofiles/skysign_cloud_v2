package postgresql

import (
	"flight-operation/pkg/flightoperation/adapters/uuid"
	fope "flight-operation/pkg/flightoperation/domain/flightoperation"
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
			sqlmock.NewRows([]string{"id", "name", "description", "fleet_id", "is_completed", "version"}).
				AddRow(DefaultID, DefaultName, DefaultDescription, DefaultFleetID, DefaultIsCompleted, DefaultVersion),
		)

	gen := uuid.NewFlightoperationUUID()
	repository := NewFlightoperationRepository(gen)

	flightoperations, err := repository.GetAll(db)

	expectFopes := []*fope.Flightoperation{
		fope.AssembleFrom(
			gen,
			&flightoperationComponentMock{
				id:          string(DefaultID),
				name:        DefaultName,
				description: DefaultDescription,
				fleetID:     string(DefaultFleetID),
				isCompleted: DefaultIsCompleted,
				version:     string(DefaultVersion),
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
		DefaultID1          = DefaultID + "-1"
		DefaultID2          = DefaultID + "-2"
		DefaultID3          = DefaultID + "-3"
		DefaultName1        = DefaultName + "-1"
		DefaultName2        = DefaultName + "-2"
		DefaultName3        = DefaultName + "-3"
		DefaultDescription1 = DefaultDescription + "-1"
		DefaultDescription2 = DefaultDescription + "-2"
		DefaultDescription3 = DefaultDescription + "-3"
		DefaultFleetID1     = DefaultFleetID + "-1"
		DefaultFleetID2     = DefaultFleetID + "-2"
		DefaultFleetID3     = DefaultFleetID + "-3"
		DefaultVersion1     = DefaultVersion + "-1"
		DefaultVersion2     = DefaultVersion + "-2"
		DefaultVersion3     = DefaultVersion + "-3"
	)

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightoperations"`)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "description", "fleet_id", "is_completed", "version"}).
				AddRow(DefaultID1, DefaultName1, DefaultDescription1, DefaultFleetID1, DefaultIsCompleted, DefaultVersion1).
				AddRow(DefaultID2, DefaultName2, DefaultDescription2, DefaultFleetID2, DefaultIsCompleted, DefaultVersion2).
				AddRow(DefaultID3, DefaultName3, DefaultDescription3, DefaultFleetID3, DefaultIsCompleted, DefaultVersion3),
		)

	gen := uuid.NewFlightoperationUUID()
	repository := NewFlightoperationRepository(gen)

	flightoperations, err := repository.GetAll(db)

	expectFopes := []*fope.Flightoperation{
		fope.AssembleFrom(
			gen,
			&flightoperationComponentMock{
				id:          string(DefaultID1),
				name:        DefaultName1,
				description: DefaultDescription1,
				fleetID:     string(DefaultFleetID1),
				isCompleted: DefaultIsCompleted,
				version:     string(DefaultVersion1),
			},
		),
		fope.AssembleFrom(
			gen,
			&flightoperationComponentMock{
				id:          string(DefaultID2),
				name:        DefaultName2,
				description: DefaultDescription2,
				fleetID:     string(DefaultFleetID2),
				isCompleted: DefaultIsCompleted,
				version:     string(DefaultVersion2),
			},
		),
		fope.AssembleFrom(
			gen,
			&flightoperationComponentMock{
				id:          string(DefaultID3),
				name:        DefaultName3,
				description: DefaultDescription3,
				fleetID:     string(DefaultFleetID3),
				isCompleted: DefaultIsCompleted,
				version:     string(DefaultVersion3),
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
			sqlmock.NewRows([]string{"id", "name", "description", "fleet_id", "is_completed", "version"}),
		)

	gen := uuid.NewFlightoperationUUID()
	repository := NewFlightoperationRepository(gen)

	flightoperations, err := repository.GetAll(db)

	var expectFopes []*fope.Flightoperation

	a.Nil(err)
	a.Len(flightoperations, 0)
	a.Equal(flightoperations, expectFopes)
}

func TestFlightoperationRepositoryGetSingleWhenGetAllOperating(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightoperations" WHERE is_completed = false`)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "description", "fleet_id", "is_completed", "version"}).
				AddRow(DefaultID, DefaultName, DefaultDescription, DefaultFleetID, DefaultIsCompleted, DefaultVersion),
		)

	gen := uuid.NewFlightoperationUUID()
	repository := NewFlightoperationRepository(gen)

	flightoperations, err := repository.GetAllOperating(db)

	expectFopes := []*fope.Flightoperation{
		fope.AssembleFrom(
			gen,
			&flightoperationComponentMock{
				id:          string(DefaultID),
				name:        DefaultName,
				description: DefaultDescription,
				fleetID:     string(DefaultFleetID),
				isCompleted: DefaultIsCompleted,
				version:     string(DefaultVersion),
			},
		),
	}

	a.Nil(err)
	a.Len(flightoperations, 1)
	a.Equal(flightoperations, expectFopes)
}

func TestFlightoperationRepositoryGetMultipleWhenGetAllOperating(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	const (
		DefaultID1          = DefaultID + "-1"
		DefaultID2          = DefaultID + "-2"
		DefaultID3          = DefaultID + "-3"
		DefaultName1        = DefaultName + "-1"
		DefaultName2        = DefaultName + "-2"
		DefaultName3        = DefaultName + "-3"
		DefaultDescription1 = DefaultDescription + "-1"
		DefaultDescription2 = DefaultDescription + "-2"
		DefaultDescription3 = DefaultDescription + "-3"
		DefaultFleetID1     = DefaultFleetID + "-1"
		DefaultFleetID2     = DefaultFleetID + "-2"
		DefaultFleetID3     = DefaultFleetID + "-3"
		DefaultVersion1     = DefaultVersion + "-1"
		DefaultVersion2     = DefaultVersion + "-2"
		DefaultVersion3     = DefaultVersion + "-3"
	)

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightoperations" WHERE is_completed = false`)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "description", "fleet_id", "is_completed", "version"}).
				AddRow(DefaultID1, DefaultName1, DefaultDescription1, DefaultFleetID1, DefaultIsCompleted, DefaultVersion1).
				AddRow(DefaultID2, DefaultName2, DefaultDescription2, DefaultFleetID2, DefaultIsCompleted, DefaultVersion2).
				AddRow(DefaultID3, DefaultName3, DefaultDescription3, DefaultFleetID3, DefaultIsCompleted, DefaultVersion3),
		)

	gen := uuid.NewFlightoperationUUID()
	repository := NewFlightoperationRepository(gen)

	flightoperations, err := repository.GetAllOperating(db)

	expectFopes := []*fope.Flightoperation{
		fope.AssembleFrom(
			gen,
			&flightoperationComponentMock{
				id:          string(DefaultID1),
				name:        DefaultName1,
				description: DefaultDescription1,
				fleetID:     string(DefaultFleetID1),
				isCompleted: DefaultIsCompleted,
				version:     string(DefaultVersion1),
			},
		),
		fope.AssembleFrom(
			gen,
			&flightoperationComponentMock{
				id:          string(DefaultID2),
				name:        DefaultName2,
				description: DefaultDescription2,
				fleetID:     string(DefaultFleetID2),
				isCompleted: DefaultIsCompleted,
				version:     string(DefaultVersion2),
			},
		),
		fope.AssembleFrom(
			gen,
			&flightoperationComponentMock{
				id:          string(DefaultID3),
				name:        DefaultName3,
				description: DefaultDescription3,
				fleetID:     string(DefaultFleetID3),
				isCompleted: DefaultIsCompleted,
				version:     string(DefaultVersion3),
			},
		),
	}

	a.Nil(err)
	a.Len(flightoperations, 3)
	a.Equal(flightoperations, expectFopes)
}

func TestFlightoperationRepositoryGetNoneWhenGetAllOperating(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightoperations" WHERE is_completed = false`)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "description", "fleet_id", "is_completed", "version"}),
		)

	gen := uuid.NewFlightoperationUUID()
	repository := NewFlightoperationRepository(gen)

	flightoperations, err := repository.GetAllOperating(db)

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
		WithArgs(DefaultID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "description", "fleet_id", "is_completed", "version"}).
				AddRow(DefaultID, DefaultName, DefaultDescription, DefaultFleetID, DefaultIsCompleted, DefaultVersion),
		)

	gen := uuid.NewFlightoperationUUID()
	repository := NewFlightoperationRepository(gen)

	flightoperation, err := repository.GetByID(db, DefaultID)

	expectFope := fope.AssembleFrom(
		gen,
		&flightoperationComponentMock{
			id:          string(DefaultID),
			name:        DefaultName,
			description: DefaultDescription,
			fleetID:     string(DefaultFleetID),
			isCompleted: DefaultIsCompleted,
			version:     string(DefaultVersion),
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
		WithArgs(DefaultID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "description", "fleet_id", "is_completed", "version"}),
		)

	gen := uuid.NewFlightoperationUUID()
	repository := NewFlightoperationRepository(gen)

	flightoperation, err := repository.GetByID(db, DefaultID)

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
		WithArgs(DefaultID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "description", "fleet_id", "is_completed", "version"}),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "flightoperations" ("id","name","description","fleet_id","is_completed","version") VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(DefaultID, DefaultName, DefaultDescription, DefaultFleetID, DefaultIsCompleted, DefaultVersion).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	gen := uuid.NewFlightoperationUUID()
	repository := NewFlightoperationRepository(gen)

	flightoperation := fope.AssembleFrom(
		gen,
		&flightoperationComponentMock{
			id:          string(DefaultID),
			name:        DefaultName,
			description: DefaultDescription,
			fleetID:     string(DefaultFleetID),
			isCompleted: DefaultIsCompleted,
			version:     string(DefaultVersion),
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
		AfterName        = DefaultName + "-after"
		AfterDescription = DefaultDescription + "-after"
		AfterFleetID     = DefaultFleetID + "-after"
		AfterIsCompleted = !DefaultIsCompleted
		AfterVersion     = DefaultVersion + "-after"
	)

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "flightoperations" WHERE id = $1`)).
		WithArgs(DefaultID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "description", "fleet_id", "is_completed", "version"}).
				AddRow(DefaultID, DefaultName, DefaultDescription, DefaultFleetID, DefaultIsCompleted, DefaultVersion),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`UPDATE "flightoperations" SET "name"=$1,"description"=$2,"fleet_id"=$3,"is_completed"=$4,"version"=$5 WHERE "id" = $6`)).
		WithArgs(AfterName, AfterDescription, AfterFleetID, AfterIsCompleted, AfterVersion, DefaultID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	gen := uuid.NewFlightoperationUUID()
	repository := NewFlightoperationRepository(gen)

	flightoperation := fope.AssembleFrom(
		gen,
		&flightoperationComponentMock{
			id:          string(DefaultID),
			name:        AfterName,
			description: AfterDescription,
			fleetID:     string(AfterFleetID),
			isCompleted: AfterIsCompleted,
			version:     string(AfterVersion),
		},
	)

	err = repository.Save(db, flightoperation)

	a.Nil(err)
}
