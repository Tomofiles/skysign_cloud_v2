package postgresql

import (
	"fleet-formation/pkg/fleet/adapters/uuid"
	fl "fleet-formation/pkg/fleet/domain/fleet"
	"regexp"
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/adapters/postgresql"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestFleetRepositoryGetByID(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "fleets" WHERE id = $1`)).
		WithArgs(DefaultFleetID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "is_carbon_copy", "version"}).
				AddRow(DefaultFleetID, fl.CarbonCopy, DefaultFleetVersion),
		)
	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "assignments" WHERE fleet_id = $1`)).
		WithArgs(DefaultFleetID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "fleet_id", "vehicle_id"}).
				AddRow(DefaultFleetAssignmentID, DefaultFleetID, DefaultFleetVehicleID),
		)
	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "events" WHERE fleet_id = $1`)).
		WithArgs(DefaultFleetID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "fleet_id", "assignment_id", "mission_id"}).
				AddRow(DefaultFleetEventID, DefaultFleetID, DefaultFleetAssignmentID, DefaultFleetMissionID),
		)

	gen := uuid.NewFleetUUID()
	repository := NewFleetRepository(gen)

	fleet, err := repository.GetByID(db, DefaultFleetID)

	assignmentComps := []assignmentComponentMock{
		{
			id:        string(DefaultFleetAssignmentID),
			fleetID:   string(DefaultFleetID),
			vehicleID: string(DefaultFleetVehicleID),
		},
	}
	eventComps := []eventComponentMock{
		{
			id:           string(DefaultFleetEventID),
			fleetID:      string(DefaultFleetID),
			assignmentID: string(DefaultFleetAssignmentID),
			missionID:    string(DefaultFleetMissionID),
		},
	}
	fleetComp := fleetComponentMock{
		id:           string(DefaultFleetID),
		assignments:  assignmentComps,
		events:       eventComps,
		isCarbonCopy: fl.CarbonCopy,
		version:      string(DefaultFleetVersion),
	}
	expectFl := fl.AssembleFrom(
		gen,
		&fleetComp,
	)

	a.Nil(err)
	a.Equal(fleet, expectFl)
}

func TestFleetRepositoryNotFoundWhenGetByID(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "fleets" WHERE id = $1`)).
		WithArgs(DefaultFleetID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "is_carbon_copy", "version"}),
		)

	gen := uuid.NewFleetUUID()
	repository := NewFleetRepository(gen)

	fleet, err := repository.GetByID(db, DefaultFleetID)

	a.Equal(err, fl.ErrNotFound)
	a.Nil(fleet)
}

func TestFleetRepositorySingleDataCreateSave(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "fleets" WHERE id = $1`)).
		WithArgs(DefaultFleetID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "is_carbon_copy", "version"}),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "fleets" ("id","is_carbon_copy","version") VALUES ($1,$2,$3)`)).
		WithArgs(DefaultFleetID, fl.CarbonCopy, DefaultFleetVersion).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "assignments" ("id","fleet_id","vehicle_id") VALUES ($1,$2,$3)`)).
		WithArgs(DefaultFleetAssignmentID, DefaultFleetID, DefaultFleetVehicleID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "events" ("id","fleet_id","assignment_id","mission_id") VALUES ($1,$2,$3,$4)`)).
		WithArgs(DefaultFleetEventID, DefaultFleetID, DefaultFleetAssignmentID, DefaultFleetMissionID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	gen := uuid.NewFleetUUID()
	repository := NewFleetRepository(gen)

	assignmentComps := []assignmentComponentMock{
		{
			id:        string(DefaultFleetAssignmentID),
			fleetID:   string(DefaultFleetID),
			vehicleID: string(DefaultFleetVehicleID),
		},
	}
	eventComps := []eventComponentMock{
		{
			id:           string(DefaultFleetEventID),
			fleetID:      string(DefaultFleetID),
			assignmentID: string(DefaultFleetAssignmentID),
			missionID:    string(DefaultFleetMissionID),
		},
	}
	fleetComp := fleetComponentMock{
		id:           string(DefaultFleetID),
		assignments:  assignmentComps,
		events:       eventComps,
		isCarbonCopy: fl.CarbonCopy,
		version:      string(DefaultFleetVersion),
	}
	fleet := fl.AssembleFrom(
		gen,
		&fleetComp,
	)

	err = repository.Save(db, fleet)

	a.Nil(err)
}

func TestFleetRepositoryNoneDataCreateSave(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "fleets" WHERE id = $1`)).
		WithArgs(DefaultFleetID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "is_carbon_copy", "version"}),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "fleets" ("id","is_carbon_copy","version") VALUES ($1,$2,$3)`)).
		WithArgs(DefaultFleetID, fl.CarbonCopy, DefaultFleetVersion).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	gen := uuid.NewFleetUUID()
	repository := NewFleetRepository(gen)

	fleetComp := fleetComponentMock{
		id:           string(DefaultFleetID),
		isCarbonCopy: fl.CarbonCopy,
		version:      string(DefaultFleetVersion),
	}
	fleet := fl.AssembleFrom(
		gen,
		&fleetComp,
	)

	err = repository.Save(db, fleet)

	a.Nil(err)
}

func TestFleetRepositoryMultipleDataCreateSave(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	data := []string{"1", "2", "3"}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "fleets" WHERE id = $1`)).
		WithArgs(DefaultFleetID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "is_carbon_copy", "version"}),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "fleets" ("id","is_carbon_copy","version") VALUES ($1,$2,$3)`)).
		WithArgs(DefaultFleetID, fl.CarbonCopy, DefaultFleetVersion).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	assignmentsArgs := []string{}
	eventsArgs := []string{}
	for _, i := range data {
		assignmentsArgs = append(
			assignmentsArgs,
			string(DefaultFleetAssignmentID)+i,
			string(DefaultFleetID),
			string(DefaultFleetVehicleID)+i,
		)
		eventsArgs = append(
			eventsArgs,
			string(DefaultFleetEventID)+i,
			string(DefaultFleetID),
			string(DefaultFleetAssignmentID)+i,
			string(DefaultFleetMissionID)+i,
		)
	}

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "assignments" ("id","fleet_id","vehicle_id") VALUES ($1,$2,$3),($4,$5,$6),($7,$8,$9)`)).
		WithArgs(
			assignmentsArgs[0],
			assignmentsArgs[1],
			assignmentsArgs[2],
			assignmentsArgs[3],
			assignmentsArgs[4],
			assignmentsArgs[5],
			assignmentsArgs[6],
			assignmentsArgs[7],
			assignmentsArgs[8],
		).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "events" ("id","fleet_id","assignment_id","mission_id") VALUES ($1,$2,$3,$4),($5,$6,$7,$8),($9,$10,$11,$12)`)).
		WithArgs(
			eventsArgs[0],
			eventsArgs[1],
			eventsArgs[2],
			eventsArgs[3],
			eventsArgs[4],
			eventsArgs[5],
			eventsArgs[6],
			eventsArgs[7],
			eventsArgs[8],
			eventsArgs[9],
			eventsArgs[10],
			eventsArgs[11],
		).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	gen := uuid.NewFleetUUID()
	repository := NewFleetRepository(gen)

	assignmentComps := []assignmentComponentMock{}
	eventComps := []eventComponentMock{}
	for _, i := range data {
		assignmentComps = append(
			assignmentComps,
			assignmentComponentMock{
				id:        string(DefaultFleetAssignmentID) + i,
				fleetID:   string(DefaultFleetID),
				vehicleID: string(DefaultFleetVehicleID) + i,
			},
		)
		eventComps = append(
			eventComps,
			eventComponentMock{
				id:           string(DefaultFleetEventID) + i,
				fleetID:      string(DefaultFleetID),
				assignmentID: string(DefaultFleetAssignmentID) + i,
				missionID:    string(DefaultFleetMissionID) + i,
			},
		)
	}

	fleetComp := fleetComponentMock{
		id:           string(DefaultFleetID),
		assignments:  assignmentComps,
		events:       eventComps,
		isCarbonCopy: fl.CarbonCopy,
		version:      string(DefaultFleetVersion),
	}
	fleet := fl.AssembleFrom(
		gen,
		&fleetComp,
	)

	err = repository.Save(db, fleet)

	a.Nil(err)
}

func TestFleetRepositorySingleDataUpdateSave(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "fleets" WHERE id = $1`)).
		WithArgs(DefaultFleetID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "is_carbon_copy", "version"}).
				AddRow(DefaultFleetID, fl.CarbonCopy, DefaultFleetVersion),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`UPDATE "fleets" SET "is_carbon_copy"=$1,"version"=$2 WHERE "id" = $3`)).
		WithArgs(fl.CarbonCopy, DefaultFleetVersion, DefaultFleetID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`DELETE FROM "assignments" WHERE fleet_id = $1`)).
		WithArgs(DefaultFleetID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "assignments" ("id","fleet_id","vehicle_id") VALUES ($1,$2,$3)`)).
		WithArgs(DefaultFleetAssignmentID, DefaultFleetID, DefaultFleetVehicleID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`DELETE FROM "events" WHERE fleet_id = $1`)).
		WithArgs(DefaultFleetID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "events" ("id","fleet_id","assignment_id","mission_id") VALUES ($1,$2,$3,$4)`)).
		WithArgs(DefaultFleetEventID, DefaultFleetID, DefaultFleetAssignmentID, DefaultFleetMissionID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	gen := uuid.NewFleetUUID()
	repository := NewFleetRepository(gen)

	assignmentComps := []assignmentComponentMock{
		{
			id:        string(DefaultFleetAssignmentID),
			fleetID:   string(DefaultFleetID),
			vehicleID: string(DefaultFleetVehicleID),
		},
	}
	eventComps := []eventComponentMock{
		{
			id:           string(DefaultFleetEventID),
			fleetID:      string(DefaultFleetID),
			assignmentID: string(DefaultFleetAssignmentID),
			missionID:    string(DefaultFleetMissionID),
		},
	}
	fleetComp := fleetComponentMock{
		id:           string(DefaultFleetID),
		assignments:  assignmentComps,
		events:       eventComps,
		isCarbonCopy: fl.CarbonCopy,
		version:      string(DefaultFleetVersion),
	}
	fleet := fl.AssembleFrom(
		gen,
		&fleetComp,
	)

	err = repository.Save(db, fleet)

	a.Nil(err)
}

func TestFleetRepositoryNoneDataUpdateSave(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "fleets" WHERE id = $1`)).
		WithArgs(DefaultFleetID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "is_carbon_copy", "version"}).
				AddRow(DefaultFleetID, fl.CarbonCopy, DefaultFleetVersion),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`UPDATE "fleets" SET "is_carbon_copy"=$1,"version"=$2 WHERE "id" = $3`)).
		WithArgs(fl.CarbonCopy, DefaultFleetVersion, DefaultFleetID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`DELETE FROM "assignments" WHERE fleet_id = $1`)).
		WithArgs(DefaultFleetID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`DELETE FROM "events" WHERE fleet_id = $1`)).
		WithArgs(DefaultFleetID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	gen := uuid.NewFleetUUID()
	repository := NewFleetRepository(gen)

	fleetComp := fleetComponentMock{
		id:           string(DefaultFleetID),
		isCarbonCopy: fl.CarbonCopy,
		version:      string(DefaultFleetVersion),
	}
	fleet := fl.AssembleFrom(
		gen,
		&fleetComp,
	)

	err = repository.Save(db, fleet)

	a.Nil(err)
}

func TestFleetRepositoryMultipleDataUpdateSave(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	data := []string{"1", "2", "3"}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "fleets" WHERE id = $1`)).
		WithArgs(DefaultFleetID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "is_carbon_copy", "version"}).
				AddRow(DefaultFleetID, fl.CarbonCopy, DefaultFleetVersion),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`UPDATE "fleets" SET "is_carbon_copy"=$1,"version"=$2 WHERE "id" = $3`)).
		WithArgs(fl.CarbonCopy, DefaultFleetVersion, DefaultFleetID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	assignmentsArgs := []string{}
	eventsArgs := []string{}
	for _, i := range data {
		assignmentsArgs = append(
			assignmentsArgs,
			string(DefaultFleetAssignmentID)+i,
			string(DefaultFleetID),
			string(DefaultFleetVehicleID)+i,
		)
		eventsArgs = append(
			eventsArgs,
			string(DefaultFleetEventID)+i,
			string(DefaultFleetID),
			string(DefaultFleetAssignmentID)+i,
			string(DefaultFleetMissionID)+i,
		)
	}

	mock.
		ExpectExec(
			regexp.QuoteMeta(`DELETE FROM "assignments" WHERE fleet_id = $1`)).
		WithArgs(DefaultFleetID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "assignments" ("id","fleet_id","vehicle_id") VALUES ($1,$2,$3),($4,$5,$6),($7,$8,$9)`)).
		WithArgs(
			assignmentsArgs[0],
			assignmentsArgs[1],
			assignmentsArgs[2],
			assignmentsArgs[3],
			assignmentsArgs[4],
			assignmentsArgs[5],
			assignmentsArgs[6],
			assignmentsArgs[7],
			assignmentsArgs[8],
		).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`DELETE FROM "events" WHERE fleet_id = $1`)).
		WithArgs(DefaultFleetID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "events" ("id","fleet_id","assignment_id","mission_id") VALUES ($1,$2,$3,$4),($5,$6,$7,$8),($9,$10,$11,$12)`)).
		WithArgs(
			eventsArgs[0],
			eventsArgs[1],
			eventsArgs[2],
			eventsArgs[3],
			eventsArgs[4],
			eventsArgs[5],
			eventsArgs[6],
			eventsArgs[7],
			eventsArgs[8],
			eventsArgs[9],
			eventsArgs[10],
			eventsArgs[11],
		).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	gen := uuid.NewFleetUUID()
	repository := NewFleetRepository(gen)

	assignmentComps := []assignmentComponentMock{}
	eventComps := []eventComponentMock{}
	for _, i := range data {
		assignmentComps = append(
			assignmentComps,
			assignmentComponentMock{
				id:        string(DefaultFleetAssignmentID) + i,
				fleetID:   string(DefaultFleetID),
				vehicleID: string(DefaultFleetVehicleID) + i,
			},
		)
		eventComps = append(
			eventComps,
			eventComponentMock{
				id:           string(DefaultFleetEventID) + i,
				fleetID:      string(DefaultFleetID),
				assignmentID: string(DefaultFleetAssignmentID) + i,
				missionID:    string(DefaultFleetMissionID) + i,
			},
		)
	}

	fleetComp := fleetComponentMock{
		id:           string(DefaultFleetID),
		assignments:  assignmentComps,
		events:       eventComps,
		isCarbonCopy: fl.CarbonCopy,
		version:      string(DefaultFleetVersion),
	}
	fleet := fl.AssembleFrom(
		gen,
		&fleetComp,
	)

	err = repository.Save(db, fleet)

	a.Nil(err)
}

func TestFleetRepositoryDelete(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "fleets" WHERE id = $1`)).
		WithArgs(DefaultFleetID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "is_carbon_copy", "version"}).
				AddRow(DefaultFleetID, fl.CarbonCopy, DefaultFleetVersion),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`DELETE FROM "fleets" WHERE "fleets"."id" = $1`)).
		WithArgs(DefaultFleetID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`DELETE FROM "assignments" WHERE fleet_id = $1`)).
		WithArgs(DefaultFleetID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`DELETE FROM "events" WHERE fleet_id = $1`)).
		WithArgs(DefaultFleetID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	gen := uuid.NewFleetUUID()
	repository := NewFleetRepository(gen)

	err = repository.Delete(db, DefaultFleetID)

	a.Nil(err)
}

func TestFleetRepositoryNotFoundWhenDelete(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "fleets" WHERE id = $1`)).
		WithArgs(DefaultFleetID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "is_carbon_copy", "version"}),
		)

	gen := uuid.NewFleetUUID()
	repository := NewFleetRepository(gen)

	err = repository.Delete(db, DefaultFleetID)

	a.Equal(err, fl.ErrNotFound)
}
