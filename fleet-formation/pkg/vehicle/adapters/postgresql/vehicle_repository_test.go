package postgresql

import (
	"regexp"
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/vehicle/adapters/uuid"
	v "github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/vehicle/domain/vehicle"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/adapters/postgresql"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestVehicleRepositoryGetSingleWhenGetAll(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "vehicles"`)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "communication_id", "is_carbon_copy", "version"}).
				AddRow(DefaultVehicleID, DefaultVehicleName, DefaultVehicleCommunicationID, v.CarbonCopy, DefaultVehicleVersion),
		)

	gen := uuid.NewVehicleUUID()
	repository := NewVehicleRepository(gen)

	vehicles, err := repository.GetAll(db)

	expectVehicles := []*v.Vehicle{
		v.AssembleFrom(
			gen,
			&vehicleComponentMock{
				ID:              string(DefaultVehicleID),
				Name:            DefaultVehicleName,
				CommunicationID: string(DefaultVehicleCommunicationID),
				IsCarbonCopy:    v.CarbonCopy,
				Version:         string(DefaultVehicleVersion),
			},
		),
	}

	a.Nil(err)
	a.Len(vehicles, 1)
	a.Equal(vehicles, expectVehicles)
}

func TestVehicleRepositoryGetMultipleWhenGetAll(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	const (
		DefaultVehicleID1              = DefaultVehicleID + "-1"
		DefaultVehicleName1            = DefaultVehicleName + "-1"
		DefaultVehicleCommunicationID1 = DefaultVehicleCommunicationID + "-1"
		DefaultVehicleVersion1         = DefaultVehicleVersion + "-1"
		DefaultVehicleID2              = DefaultVehicleID + "-2"
		DefaultVehicleName2            = DefaultVehicleName + "-2"
		DefaultVehicleCommunicationID2 = DefaultVehicleCommunicationID + "-2"
		DefaultVehicleVersion2         = DefaultVehicleVersion + "-2"
		DefaultVehicleID3              = DefaultVehicleID + "-3"
		DefaultVehicleName3            = DefaultVehicleName + "-3"
		DefaultVehicleCommunicationID3 = DefaultVehicleCommunicationID + "-3"
		DefaultVehicleVersion3         = DefaultVehicleVersion + "-3"
	)

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "vehicles"`)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "communication_id", "is_carbon_copy", "version"}).
				AddRow(DefaultVehicleID1, DefaultVehicleName1, DefaultVehicleCommunicationID1, v.CarbonCopy, DefaultVehicleVersion1).
				AddRow(DefaultVehicleID2, DefaultVehicleName2, DefaultVehicleCommunicationID2, v.CarbonCopy, DefaultVehicleVersion2).
				AddRow(DefaultVehicleID3, DefaultVehicleName3, DefaultVehicleCommunicationID3, v.CarbonCopy, DefaultVehicleVersion3),
		)

	gen := uuid.NewVehicleUUID()
	repository := NewVehicleRepository(gen)

	vehicles, err := repository.GetAll(db)

	expectVehicles := []*v.Vehicle{
		v.AssembleFrom(
			gen,
			&vehicleComponentMock{
				ID:              string(DefaultVehicleID1),
				Name:            DefaultVehicleName1,
				CommunicationID: string(DefaultVehicleCommunicationID1),
				IsCarbonCopy:    v.CarbonCopy,
				Version:         string(DefaultVehicleVersion1),
			},
		),
		v.AssembleFrom(
			gen,
			&vehicleComponentMock{
				ID:              string(DefaultVehicleID2),
				Name:            DefaultVehicleName2,
				CommunicationID: string(DefaultVehicleCommunicationID2),
				IsCarbonCopy:    v.CarbonCopy,
				Version:         string(DefaultVehicleVersion2),
			},
		),
		v.AssembleFrom(
			gen,
			&vehicleComponentMock{
				ID:              string(DefaultVehicleID3),
				Name:            DefaultVehicleName3,
				CommunicationID: string(DefaultVehicleCommunicationID3),
				IsCarbonCopy:    v.CarbonCopy,
				Version:         string(DefaultVehicleVersion3),
			},
		),
	}

	a.Nil(err)
	a.Len(vehicles, 3)
	a.Equal(vehicles, expectVehicles)
}

func TestVehicleRepositoryGetNoneWhenGetAll(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "vehicles"`)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "communication_id", "is_carbon_copy", "version"}),
		)

	gen := uuid.NewVehicleUUID()
	repository := NewVehicleRepository(gen)

	vehicles, err := repository.GetAll(db)

	var expectVehicles []*v.Vehicle

	a.Nil(err)
	a.Len(vehicles, 0)
	a.Equal(vehicles, expectVehicles)
}

func TestVehicleRepositoryGetSingleWhenGetAllOrigin(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "vehicles" WHERE is_carbon_copy = false`)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "communication_id", "is_carbon_copy", "version"}).
				AddRow(DefaultVehicleID, DefaultVehicleName, DefaultVehicleCommunicationID, v.CarbonCopy, DefaultVehicleVersion),
		)

	gen := uuid.NewVehicleUUID()
	repository := NewVehicleRepository(gen)

	vehicles, err := repository.GetAllOrigin(db)

	expectVehicles := []*v.Vehicle{
		v.AssembleFrom(
			gen,
			&vehicleComponentMock{
				ID:              string(DefaultVehicleID),
				Name:            DefaultVehicleName,
				CommunicationID: string(DefaultVehicleCommunicationID),
				IsCarbonCopy:    v.CarbonCopy,
				Version:         string(DefaultVehicleVersion),
			},
		),
	}

	a.Nil(err)
	a.Len(vehicles, 1)
	a.Equal(vehicles, expectVehicles)
}

func TestVehicleRepositoryGetMultipleWhenGetAllOrigin(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	const (
		DefaultVehicleID1              = DefaultVehicleID + "-1"
		DefaultVehicleName1            = DefaultVehicleName + "-1"
		DefaultVehicleCommunicationID1 = DefaultVehicleCommunicationID + "-1"
		DefaultVehicleVersion1         = DefaultVehicleVersion + "-1"
		DefaultVehicleID2              = DefaultVehicleID + "-2"
		DefaultVehicleName2            = DefaultVehicleName + "-2"
		DefaultVehicleCommunicationID2 = DefaultVehicleCommunicationID + "-2"
		DefaultVehicleVersion2         = DefaultVehicleVersion + "-2"
		DefaultVehicleID3              = DefaultVehicleID + "-3"
		DefaultVehicleName3            = DefaultVehicleName + "-3"
		DefaultVehicleCommunicationID3 = DefaultVehicleCommunicationID + "-3"
		DefaultVehicleVersion3         = DefaultVehicleVersion + "-3"
	)

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "vehicles" WHERE is_carbon_copy = false`)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "communication_id", "is_carbon_copy", "version"}).
				AddRow(DefaultVehicleID1, DefaultVehicleName1, DefaultVehicleCommunicationID1, v.CarbonCopy, DefaultVehicleVersion1).
				AddRow(DefaultVehicleID2, DefaultVehicleName2, DefaultVehicleCommunicationID2, v.CarbonCopy, DefaultVehicleVersion2).
				AddRow(DefaultVehicleID3, DefaultVehicleName3, DefaultVehicleCommunicationID3, v.CarbonCopy, DefaultVehicleVersion3),
		)

	gen := uuid.NewVehicleUUID()
	repository := NewVehicleRepository(gen)

	vehicles, err := repository.GetAllOrigin(db)

	expectVehicles := []*v.Vehicle{
		v.AssembleFrom(
			gen,
			&vehicleComponentMock{
				ID:              string(DefaultVehicleID1),
				Name:            DefaultVehicleName1,
				CommunicationID: string(DefaultVehicleCommunicationID1),
				IsCarbonCopy:    v.CarbonCopy,
				Version:         string(DefaultVehicleVersion1),
			},
		),
		v.AssembleFrom(
			gen,
			&vehicleComponentMock{
				ID:              string(DefaultVehicleID2),
				Name:            DefaultVehicleName2,
				CommunicationID: string(DefaultVehicleCommunicationID2),
				IsCarbonCopy:    v.CarbonCopy,
				Version:         string(DefaultVehicleVersion2),
			},
		),
		v.AssembleFrom(
			gen,
			&vehicleComponentMock{
				ID:              string(DefaultVehicleID3),
				Name:            DefaultVehicleName3,
				CommunicationID: string(DefaultVehicleCommunicationID3),
				IsCarbonCopy:    v.CarbonCopy,
				Version:         string(DefaultVehicleVersion3),
			},
		),
	}

	a.Nil(err)
	a.Len(vehicles, 3)
	a.Equal(vehicles, expectVehicles)
}

func TestVehicleRepositoryGetNoneWhenGetAllOrigin(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "vehicles" WHERE is_carbon_copy = false`)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "communication_id", "is_carbon_copy", "version"}),
		)

	gen := uuid.NewVehicleUUID()
	repository := NewVehicleRepository(gen)

	vehicles, err := repository.GetAllOrigin(db)

	var expectVehicles []*v.Vehicle

	a.Nil(err)
	a.Len(vehicles, 0)
	a.Equal(vehicles, expectVehicles)
}

func TestVehicleRepositoryGetByID(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "vehicles" WHERE id = $1`)).
		WithArgs(DefaultVehicleID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "communication_id", "is_carbon_copy", "version"}).
				AddRow(DefaultVehicleID, DefaultVehicleName, DefaultVehicleCommunicationID, v.CarbonCopy, DefaultVehicleVersion),
		)

	gen := uuid.NewVehicleUUID()
	repository := NewVehicleRepository(gen)

	vehicle, err := repository.GetByID(db, DefaultVehicleID)

	expectVehicle := v.AssembleFrom(
		gen,
		&vehicleComponentMock{
			ID:              string(DefaultVehicleID),
			Name:            DefaultVehicleName,
			CommunicationID: string(DefaultVehicleCommunicationID),
			IsCarbonCopy:    v.CarbonCopy,
			Version:         string(DefaultVehicleVersion),
		},
	)

	a.Nil(err)
	a.Equal(vehicle, expectVehicle)
}

func TestVehicleRepositoryNotFoundWhenGetByID(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "vehicles" WHERE id = $1`)).
		WithArgs(DefaultVehicleID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "communication_id", "is_carbon_copy", "version"}),
		)

	gen := uuid.NewVehicleUUID()
	repository := NewVehicleRepository(gen)

	vehicle, err := repository.GetByID(db, DefaultVehicleID)

	a.Nil(vehicle)
	a.Equal(err, v.ErrNotFound)
}

func TestVehicleRepositoryCreateSave(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "vehicles" WHERE id = $1`)).
		WithArgs(DefaultVehicleID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "communication_id", "is_carbon_copy", "version"}),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "vehicles" ("id","name","communication_id","is_carbon_copy","version") VALUES ($1,$2,$3,$4,$5)`)).
		WithArgs(DefaultVehicleID, DefaultVehicleName, DefaultVehicleCommunicationID, v.CarbonCopy, DefaultVehicleVersion).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	gen := uuid.NewVehicleUUID()
	repository := NewVehicleRepository(gen)

	vehicle := v.AssembleFrom(
		gen,
		&vehicleComponentMock{
			ID:              string(DefaultVehicleID),
			Name:            DefaultVehicleName,
			CommunicationID: string(DefaultVehicleCommunicationID),
			IsCarbonCopy:    v.CarbonCopy,
			Version:         string(DefaultVehicleVersion),
		},
	)

	err = repository.Save(db, vehicle)

	a.Nil(err)
}

func TestVehicleRepositoryUpdateSave(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	const (
		AfterName            = DefaultVehicleName + "-after"
		AfterCommunicationID = DefaultVehicleCommunicationID + "-after"
		AfterVersion         = DefaultVehicleVersion + "-after"
	)

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "vehicles" WHERE id = $1`)).
		WithArgs(DefaultVehicleID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "communication_id", "is_carbon_copy", "version"}).
				AddRow(DefaultVehicleID, DefaultVehicleName, DefaultVehicleCommunicationID, v.CarbonCopy, DefaultVehicleVersion),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`UPDATE "vehicles" SET "name"=$1,"communication_id"=$2,"is_carbon_copy"=$3,"version"=$4 WHERE "id" = $5`)).
		WithArgs(AfterName, AfterCommunicationID, v.CarbonCopy, AfterVersion, DefaultVehicleID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	gen := uuid.NewVehicleUUID()
	repository := NewVehicleRepository(gen)

	vehicle := v.AssembleFrom(
		gen,
		&vehicleComponentMock{
			ID:              string(DefaultVehicleID),
			Name:            AfterName,
			CommunicationID: string(AfterCommunicationID),
			IsCarbonCopy:    v.CarbonCopy,
			Version:         string(AfterVersion),
		},
	)

	err = repository.Save(db, vehicle)

	a.Nil(err)
}

func TestVehicleRepositoryDelete(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "vehicles" WHERE id = $1`)).
		WithArgs(DefaultVehicleID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "communication_id", "is_carbon_copy", "version"}).
				AddRow(DefaultVehicleID, DefaultVehicleName, DefaultVehicleCommunicationID, v.CarbonCopy, DefaultVehicleVersion),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`DELETE FROM "vehicles" WHERE "vehicles"."id" = $1`)).
		WithArgs(DefaultVehicleID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	gen := uuid.NewVehicleUUID()
	repository := NewVehicleRepository(gen)

	err = repository.Delete(db, DefaultVehicleID)

	a.Nil(err)
}
