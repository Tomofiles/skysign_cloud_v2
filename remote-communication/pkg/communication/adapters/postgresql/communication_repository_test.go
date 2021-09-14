package postgresql

import (
	"regexp"
	"remote-communication/pkg/communication/adapters/uuid"
	c "remote-communication/pkg/communication/domain/communication"
	"testing"
	"time"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/adapters/postgresql"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCommunicationRepositoryGetByID(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "communications" WHERE id = $1`)).
		WithArgs(DefaultCommunicationID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).
				AddRow(DefaultCommunicationID),
		)
	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "telemetries" WHERE communication_id = $1`)).
		WithArgs(DefaultCommunicationID).
		WillReturnRows(
			sqlmock.NewRows([]string{"communication_id", "latitude_degree", "longitude_degree", "altitude_m", "relative_altitude_m", "speed_ms", "armed", "flight_mode", "orientation_x", "orientation_y", "orientation_z", "orientation_w"}).
				AddRow(DefaultCommunicationID, 1.0, 2.0, 3.0, 4.0, 5.0, true, "state", 6.0, 7.0, 8.0, 9.0),
		)
	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "commands" WHERE communication_id = $1`)).
		WithArgs(DefaultCommunicationID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "communication_id", "type", "time"}).
				AddRow(DefaultCommunicationCommandID, DefaultCommunicationID, c.CommandTypeARM, DefaultCommunicationTime),
		)
	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "upload_missions" WHERE communication_id = $1`)).
		WithArgs(DefaultCommunicationID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "communication_id", "mission_id"}).
				AddRow(DefaultCommunicationCommandID, DefaultCommunicationID, DefaultCommunicationMissionID),
		)

	gen := uuid.NewCommunicationUUID()
	repository := NewCommunicationRepository(gen)

	communication, err := repository.GetByID(db, DefaultCommunicationID)

	commandComps := []*commandComponentMock{
		{
			id:    string(DefaultCommunicationCommandID),
			cType: string(c.CommandTypeARM),
			time:  DefaultCommunicationTime,
		},
	}
	uploadMissionComps := []*uploadMissionComponentMock{
		{
			commandID: string(DefaultCommunicationCommandID),
			missionID: string(DefaultCommunicationMissionID),
		},
	}
	telemetryComp := telemetryComponentMock{
		latitude:         1.0,
		longitude:        2.0,
		altitude:         3.0,
		relativeAltitude: 4.0,
		speed:            5.0,
		armed:            true,
		flightMode:       "state",
		x:                6.0,
		y:                7.0,
		z:                8.0,
		w:                9.0,
	}
	communicationComp := communicationComponentMock{
		id:             string(DefaultCommunicationID),
		telemetry:      &telemetryComp,
		commands:       commandComps,
		uploadMissions: uploadMissionComps,
	}
	expectComm := c.AssembleFrom(
		gen,
		&communicationComp,
	)

	a.Nil(err)
	a.Equal(communication, expectComm)
}

func TestCommunicationRepositoryNotFoundWhenGetByID(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "communications" WHERE id = $1`)).
		WithArgs(DefaultCommunicationID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}),
		)

	gen := uuid.NewCommunicationUUID()
	repository := NewCommunicationRepository(gen)

	communication, err := repository.GetByID(db, DefaultCommunicationID)

	a.Equal(err, c.ErrNotFound)
	a.Nil(communication)
}

func TestCommunicationRepositorySingleDataCreateSave(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "communications" WHERE id = $1`)).
		WithArgs(DefaultCommunicationID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "communications" ("id") VALUES ($1)`)).
		WithArgs(DefaultCommunicationID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "telemetries" ("communication_id","latitude_degree","longitude_degree","altitude_m","relative_altitude_m","speed_ms","armed","flight_mode","orientation_x","orientation_y","orientation_z","orientation_w") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)`)).
		WithArgs(DefaultCommunicationID, 1.0, 2.0, 3.0, 4.0, 5.0, true, "state", 6.0, 7.0, 8.0, 9.0).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "commands" ("id","communication_id","type","time") VALUES ($1,$2,$3,$4)`)).
		WithArgs(DefaultCommunicationCommandID, DefaultCommunicationID, c.CommandTypeARM, DefaultCommunicationTime).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "upload_missions" ("id","communication_id","mission_id") VALUES ($1,$2,$3)`)).
		WithArgs(DefaultCommunicationCommandID, DefaultCommunicationID, DefaultCommunicationMissionID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	gen := uuid.NewCommunicationUUID()
	repository := NewCommunicationRepository(gen)

	commandComps := []*commandComponentMock{
		{
			id:    string(DefaultCommunicationCommandID),
			cType: string(c.CommandTypeARM),
			time:  DefaultCommunicationTime,
		},
	}
	uploadMissionComps := []*uploadMissionComponentMock{
		{
			commandID: string(DefaultCommunicationCommandID),
			missionID: string(DefaultCommunicationMissionID),
		},
	}
	telemetryComp := telemetryComponentMock{
		latitude:         1.0,
		longitude:        2.0,
		altitude:         3.0,
		relativeAltitude: 4.0,
		speed:            5.0,
		armed:            true,
		flightMode:       "state",
		x:                6.0,
		y:                7.0,
		z:                8.0,
		w:                9.0,
	}
	communicationComp := communicationComponentMock{
		id:             string(DefaultCommunicationID),
		telemetry:      &telemetryComp,
		commands:       commandComps,
		uploadMissions: uploadMissionComps,
	}
	communication := c.AssembleFrom(
		gen,
		&communicationComp,
	)

	err = repository.Save(db, communication)

	a.Nil(err)
}

func TestCommunicationRepositoryNoneDataCreateSave(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "communications" WHERE id = $1`)).
		WithArgs(DefaultCommunicationID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "communications" ("id") VALUES ($1)`)).
		WithArgs(DefaultCommunicationID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "telemetries" ("communication_id","latitude_degree","longitude_degree","altitude_m","relative_altitude_m","speed_ms","armed","flight_mode","orientation_x","orientation_y","orientation_z","orientation_w") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)`)).
		WithArgs(DefaultCommunicationID, 1.0, 2.0, 3.0, 4.0, 5.0, true, "state", 6.0, 7.0, 8.0, 9.0).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	gen := uuid.NewCommunicationUUID()
	repository := NewCommunicationRepository(gen)

	commandComps := []*commandComponentMock{}
	uploadMissionComps := []*uploadMissionComponentMock{}
	telemetryComp := telemetryComponentMock{
		latitude:         1.0,
		longitude:        2.0,
		altitude:         3.0,
		relativeAltitude: 4.0,
		speed:            5.0,
		armed:            true,
		flightMode:       "state",
		x:                6.0,
		y:                7.0,
		z:                8.0,
		w:                9.0,
	}
	communicationComp := communicationComponentMock{
		id:             string(DefaultCommunicationID),
		telemetry:      &telemetryComp,
		commands:       commandComps,
		uploadMissions: uploadMissionComps,
	}
	communication := c.AssembleFrom(
		gen,
		&communicationComp,
	)

	err = repository.Save(db, communication)

	a.Nil(err)
}

func TestCommunicationRepositoryMultipleDataCreateSave(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	data := []string{"1", "2", "3"}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "communications" WHERE id = $1`)).
		WithArgs(DefaultCommunicationID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "communications" ("id") VALUES ($1)`)).
		WithArgs(DefaultCommunicationID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "telemetries" ("communication_id","latitude_degree","longitude_degree","altitude_m","relative_altitude_m","speed_ms","armed","flight_mode","orientation_x","orientation_y","orientation_z","orientation_w") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)`)).
		WithArgs(DefaultCommunicationID, 1.0, 2.0, 3.0, 4.0, 5.0, true, "state", 6.0, 7.0, 8.0, 9.0).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	commandsArgs := []interface{}{}
	uploadMissionsArgs := []interface{}{}
	for _, i := range data {
		commandsArgs = append(
			commandsArgs,
			string(DefaultCommunicationCommandID)+i,
			string(DefaultCommunicationID),
			string(c.CommandTypeARM)+i,
			DefaultCommunicationTime.Add(time.Second),
		)
		uploadMissionsArgs = append(
			uploadMissionsArgs,
			string(DefaultCommunicationCommandID)+i,
			string(DefaultCommunicationID),
			string(DefaultCommunicationMissionID)+i,
		)
	}

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "commands" ("id","communication_id","type","time") VALUES ($1,$2,$3,$4),($5,$6,$7,$8),($9,$10,$11,$12)`)).
		WithArgs(
			commandsArgs[0],
			commandsArgs[1],
			commandsArgs[2],
			commandsArgs[3],
			commandsArgs[4],
			commandsArgs[5],
			commandsArgs[6],
			commandsArgs[7],
			commandsArgs[8],
			commandsArgs[9],
			commandsArgs[10],
			commandsArgs[11],
		).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "upload_missions" ("id","communication_id","mission_id") VALUES ($1,$2,$3),($4,$5,$6),($7,$8,$9)`)).
		WithArgs(
			uploadMissionsArgs[0],
			uploadMissionsArgs[1],
			uploadMissionsArgs[2],
			uploadMissionsArgs[3],
			uploadMissionsArgs[4],
			uploadMissionsArgs[5],
			uploadMissionsArgs[6],
			uploadMissionsArgs[7],
			uploadMissionsArgs[8],
		).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	gen := uuid.NewCommunicationUUID()
	repository := NewCommunicationRepository(gen)

	commandComps := []*commandComponentMock{}
	uploadMissionComps := []*uploadMissionComponentMock{}
	for _, i := range data {
		commandComps = append(
			commandComps,
			&commandComponentMock{
				id:    string(DefaultCommunicationCommandID) + i,
				cType: string(c.CommandTypeARM) + i,
				time:  DefaultCommunicationTime.Add(time.Second),
			},
		)
		uploadMissionComps = append(
			uploadMissionComps,
			&uploadMissionComponentMock{
				commandID: string(DefaultCommunicationCommandID) + i,
				missionID: string(DefaultCommunicationMissionID) + i,
			},
		)
	}
	telemetryComp := telemetryComponentMock{
		latitude:         1.0,
		longitude:        2.0,
		altitude:         3.0,
		relativeAltitude: 4.0,
		speed:            5.0,
		armed:            true,
		flightMode:       "state",
		x:                6.0,
		y:                7.0,
		z:                8.0,
		w:                9.0,
	}
	communicationComp := communicationComponentMock{
		id:             string(DefaultCommunicationID),
		telemetry:      &telemetryComp,
		commands:       commandComps,
		uploadMissions: uploadMissionComps,
	}
	communication := c.AssembleFrom(
		gen,
		&communicationComp,
	)

	err = repository.Save(db, communication)

	a.Nil(err)
}

func TestCommunicationRepositorySingleDataUpdateSave(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "communications" WHERE id = $1`)).
		WithArgs(DefaultCommunicationID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).
				AddRow(DefaultCommunicationID),
		)

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "communications" WHERE "id" = $1 ORDER BY "communications"."id" LIMIT 1`)).
		WithArgs(DefaultCommunicationID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).
				AddRow(DefaultCommunicationID),
		)

	// mock.
	// 	ExpectExec(
	// 		regexp.QuoteMeta(`UPDATE "communications" SET "id"=$1 WHERE "id" = $2`)).
	// 	WithArgs(DefaultCommunicationID, DefaultCommunicationID).
	// 	WillReturnResult(
	// 		sqlmock.NewResult(1, 1),
	// 	)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`UPDATE "telemetries" SET "latitude_degree"=$1,"longitude_degree"=$2,"altitude_m"=$3,"relative_altitude_m"=$4,"speed_ms"=$5,"armed"=$6,"flight_mode"=$7,"orientation_x"=$8,"orientation_y"=$9,"orientation_z"=$10,"orientation_w"=$11 WHERE "communication_id" = $12`)).
		WithArgs(1.0, 2.0, 3.0, 4.0, 5.0, true, "state", 6.0, 7.0, 8.0, 9.0, DefaultCommunicationID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`DELETE FROM "commands" WHERE communication_id = $1`)).
		WithArgs(DefaultCommunicationID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "commands" ("id","communication_id","type","time") VALUES ($1,$2,$3,$4)`)).
		WithArgs(DefaultCommunicationCommandID, DefaultCommunicationID, c.CommandTypeARM, DefaultCommunicationTime).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`DELETE FROM "upload_missions" WHERE communication_id = $1`)).
		WithArgs(DefaultCommunicationID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "upload_missions" ("id","communication_id","mission_id") VALUES ($1,$2,$3)`)).
		WithArgs(DefaultCommunicationCommandID, DefaultCommunicationID, DefaultCommunicationMissionID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	gen := uuid.NewCommunicationUUID()
	repository := NewCommunicationRepository(gen)

	commandComps := []*commandComponentMock{
		{
			id:    string(DefaultCommunicationCommandID),
			cType: string(c.CommandTypeARM),
			time:  DefaultCommunicationTime,
		},
	}
	uploadMissionComps := []*uploadMissionComponentMock{
		{
			commandID: string(DefaultCommunicationCommandID),
			missionID: string(DefaultCommunicationMissionID),
		},
	}
	telemetryComp := telemetryComponentMock{
		latitude:         1.0,
		longitude:        2.0,
		altitude:         3.0,
		relativeAltitude: 4.0,
		speed:            5.0,
		armed:            true,
		flightMode:       "state",
		x:                6.0,
		y:                7.0,
		z:                8.0,
		w:                9.0,
	}
	communicationComp := communicationComponentMock{
		id:             string(DefaultCommunicationID),
		telemetry:      &telemetryComp,
		commands:       commandComps,
		uploadMissions: uploadMissionComps,
	}
	communication := c.AssembleFrom(
		gen,
		&communicationComp,
	)

	err = repository.Save(db, communication)

	a.Nil(err)
}

func TestCommunicationRepositoryNoneDataUpdateSave(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "communications" WHERE id = $1`)).
		WithArgs(DefaultCommunicationID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).
				AddRow(DefaultCommunicationID),
		)

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "communications" WHERE "id" = $1 ORDER BY "communications"."id" LIMIT 1`)).
		WithArgs(DefaultCommunicationID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).
				AddRow(DefaultCommunicationID),
		)

	// mock.
	// 	ExpectExec(
	// 		regexp.QuoteMeta(`UPDATE "communications" SET "id"=$1 WHERE "id" = $2`)).
	// 	WithArgs(DefaultCommunicationID, DefaultCommunicationID).
	// 	WillReturnResult(
	// 		sqlmock.NewResult(1, 1),
	// 	)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`UPDATE "telemetries" SET "latitude_degree"=$1,"longitude_degree"=$2,"altitude_m"=$3,"relative_altitude_m"=$4,"speed_ms"=$5,"armed"=$6,"flight_mode"=$7,"orientation_x"=$8,"orientation_y"=$9,"orientation_z"=$10,"orientation_w"=$11 WHERE "communication_id" = $12`)).
		WithArgs(1.0, 2.0, 3.0, 4.0, 5.0, true, "state", 6.0, 7.0, 8.0, 9.0, DefaultCommunicationID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`DELETE FROM "commands" WHERE communication_id = $1`)).
		WithArgs(DefaultCommunicationID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`DELETE FROM "upload_missions" WHERE communication_id = $1`)).
		WithArgs(DefaultCommunicationID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	gen := uuid.NewCommunicationUUID()
	repository := NewCommunicationRepository(gen)

	commandComps := []*commandComponentMock{}
	uploadMissionComps := []*uploadMissionComponentMock{}
	telemetryComp := telemetryComponentMock{
		latitude:         1.0,
		longitude:        2.0,
		altitude:         3.0,
		relativeAltitude: 4.0,
		speed:            5.0,
		armed:            true,
		flightMode:       "state",
		x:                6.0,
		y:                7.0,
		z:                8.0,
		w:                9.0,
	}
	communicationComp := communicationComponentMock{
		id:             string(DefaultCommunicationID),
		telemetry:      &telemetryComp,
		commands:       commandComps,
		uploadMissions: uploadMissionComps,
	}
	communication := c.AssembleFrom(
		gen,
		&communicationComp,
	)

	err = repository.Save(db, communication)

	a.Nil(err)
}

func TestCommunicationRepositoryMultipleDataUpdateSave(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	data := []string{"1", "2", "3"}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "communications" WHERE id = $1`)).
		WithArgs(DefaultCommunicationID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).
				AddRow(DefaultCommunicationID),
		)

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "communications" WHERE "id" = $1 ORDER BY "communications"."id" LIMIT 1`)).
		WithArgs(DefaultCommunicationID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).
				AddRow(DefaultCommunicationID),
		)

	// mock.
	// 	ExpectExec(
	// 		regexp.QuoteMeta(`UPDATE "communications" SET "id"=$1 WHERE "id" = $2`)).
	// 	WithArgs(DefaultCommunicationID, DefaultCommunicationID).
	// 	WillReturnResult(
	// 		sqlmock.NewResult(1, 1),
	// 	)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`UPDATE "telemetries" SET "latitude_degree"=$1,"longitude_degree"=$2,"altitude_m"=$3,"relative_altitude_m"=$4,"speed_ms"=$5,"armed"=$6,"flight_mode"=$7,"orientation_x"=$8,"orientation_y"=$9,"orientation_z"=$10,"orientation_w"=$11 WHERE "communication_id" = $12`)).
		WithArgs(1.0, 2.0, 3.0, 4.0, 5.0, true, "state", 6.0, 7.0, 8.0, 9.0, DefaultCommunicationID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	commandsArgs := []interface{}{}
	uploadMissionsArgs := []interface{}{}
	for _, i := range data {
		commandsArgs = append(
			commandsArgs,
			string(DefaultCommunicationCommandID)+i,
			string(DefaultCommunicationID),
			string(c.CommandTypeARM)+i,
			DefaultCommunicationTime.Add(time.Second),
		)
		uploadMissionsArgs = append(
			uploadMissionsArgs,
			string(DefaultCommunicationCommandID)+i,
			string(DefaultCommunicationID),
			string(DefaultCommunicationMissionID)+i,
		)
	}

	mock.
		ExpectExec(
			regexp.QuoteMeta(`DELETE FROM "commands" WHERE communication_id = $1`)).
		WithArgs(DefaultCommunicationID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "commands" ("id","communication_id","type","time") VALUES ($1,$2,$3,$4),($5,$6,$7,$8),($9,$10,$11,$12)`)).
		WithArgs(
			commandsArgs[0],
			commandsArgs[1],
			commandsArgs[2],
			commandsArgs[3],
			commandsArgs[4],
			commandsArgs[5],
			commandsArgs[6],
			commandsArgs[7],
			commandsArgs[8],
			commandsArgs[9],
			commandsArgs[10],
			commandsArgs[11],
		).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`DELETE FROM "upload_missions" WHERE communication_id = $1`)).
		WithArgs(DefaultCommunicationID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "upload_missions" ("id","communication_id","mission_id") VALUES ($1,$2,$3),($4,$5,$6),($7,$8,$9)`)).
		WithArgs(
			uploadMissionsArgs[0],
			uploadMissionsArgs[1],
			uploadMissionsArgs[2],
			uploadMissionsArgs[3],
			uploadMissionsArgs[4],
			uploadMissionsArgs[5],
			uploadMissionsArgs[6],
			uploadMissionsArgs[7],
			uploadMissionsArgs[8],
		).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	gen := uuid.NewCommunicationUUID()
	repository := NewCommunicationRepository(gen)

	commandComps := []*commandComponentMock{}
	uploadMissionComps := []*uploadMissionComponentMock{}
	for _, i := range data {
		commandComps = append(
			commandComps,
			&commandComponentMock{
				id:    string(DefaultCommunicationCommandID) + i,
				cType: string(c.CommandTypeARM) + i,
				time:  DefaultCommunicationTime.Add(time.Second),
			},
		)
		uploadMissionComps = append(
			uploadMissionComps,
			&uploadMissionComponentMock{
				commandID: string(DefaultCommunicationCommandID) + i,
				missionID: string(DefaultCommunicationMissionID) + i,
			},
		)
	}
	telemetryComp := telemetryComponentMock{
		latitude:         1.0,
		longitude:        2.0,
		altitude:         3.0,
		relativeAltitude: 4.0,
		speed:            5.0,
		armed:            true,
		flightMode:       "state",
		x:                6.0,
		y:                7.0,
		z:                8.0,
		w:                9.0,
	}
	communicationComp := communicationComponentMock{
		id:             string(DefaultCommunicationID),
		telemetry:      &telemetryComp,
		commands:       commandComps,
		uploadMissions: uploadMissionComps,
	}
	communication := c.AssembleFrom(
		gen,
		&communicationComp,
	)

	err = repository.Save(db, communication)

	a.Nil(err)
}

func TestCommunicationRepositoryDelete(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "communications" WHERE id = $1`)).
		WithArgs(DefaultCommunicationID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).
				AddRow(DefaultCommunicationID),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`DELETE FROM "communications" WHERE "communications"."id" = $1`)).
		WithArgs(DefaultCommunicationID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`DELETE FROM "telemetries" WHERE communication_id = $1`)).
		WithArgs(DefaultCommunicationID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`DELETE FROM "commands" WHERE communication_id = $1`)).
		WithArgs(DefaultCommunicationID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`DELETE FROM "upload_missions" WHERE communication_id = $1`)).
		WithArgs(DefaultCommunicationID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	gen := uuid.NewCommunicationUUID()
	repository := NewCommunicationRepository(gen)

	err = repository.Delete(db, DefaultCommunicationID)

	a.Nil(err)
}

func TestCommunicationRepositoryNotFoundWhenDelete(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "communications" WHERE id = $1`)).
		WithArgs(DefaultCommunicationID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}),
		)

	gen := uuid.NewCommunicationUUID()
	repository := NewCommunicationRepository(gen)

	err = repository.Delete(db, DefaultCommunicationID)

	a.Equal(err, c.ErrNotFound)
}
