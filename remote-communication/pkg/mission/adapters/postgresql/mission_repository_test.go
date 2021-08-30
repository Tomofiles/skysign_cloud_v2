package postgresql

import (
	"regexp"
	"remote-communication/pkg/mission/domain/mission"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestMissionRepositoryGetByID(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "missions" WHERE id = $1`)).
		WithArgs(DefaultMissionID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).
				AddRow(DefaultMissionID),
		)
	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "waypoints" WHERE mission_id = $1`)).
		WithArgs(DefaultMissionID).
		WillReturnRows(
			sqlmock.NewRows([]string{"mission_id", "point_order", "latitude_degree", "longitude_degree", "relative_height_m", "speed_ms"}).
				AddRow(DefaultMissionID, 1, 11.0, 21.0, 31.0, 41.0).
				AddRow(DefaultMissionID, 2, 12.0, 22.0, 32.0, 42.0).
				AddRow(DefaultMissionID, 3, 13.0, 23.0, 33.0, 43.0),
		)

	repository := NewMissionRepository()

	ms, err := repository.GetByID(db, DefaultMissionID)

	waypointComps := []*waypointComponentMock{
		{
			1, 11.0, 21.0, 31.0, 41.0,
		},
		{
			2, 12.0, 22.0, 32.0, 42.0,
		},
		{
			3, 13.0, 23.0, 33.0, 43.0,
		},
	}
	missionComp := &missionComponentMock{
		id:        string(DefaultMissionID),
		waypoints: waypointComps,
	}
	expectMission := mission.AssembleFrom(missionComp)

	a.Nil(err)
	a.Equal(ms, expectMission)
}

func TestMissionRepositoryNotFoundWhenGetByID(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "missions" WHERE id = $1`)).
		WithArgs(DefaultMissionID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}),
		)

	repository := NewMissionRepository()

	ms, err := repository.GetByID(db, DefaultMissionID)

	a.Equal(err, mission.ErrNotFound)
	a.Nil(ms)
}

func TestMissionRepositorySingleDataCreateSave(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "missions" WHERE id = $1`)).
		WithArgs(DefaultMissionID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "missions" ("id") VALUES ($1)`)).
		WithArgs(DefaultMissionID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "waypoints" ("mission_id","point_order","latitude_degree","longitude_degree","relative_height_m","speed_ms") VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(DefaultMissionID, 1, 11.0, 21.0, 31.0, 41.0).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	repository := NewMissionRepository()

	waypointComps := []*waypointComponentMock{
		{
			1, 11.0, 21.0, 31.0, 41.0,
		},
	}
	missionComp := &missionComponentMock{
		id:        string(DefaultMissionID),
		waypoints: waypointComps,
	}
	ms := mission.AssembleFrom(missionComp)

	err = repository.Save(db, ms)

	a.Nil(err)
}

func TestMissionRepositoryNoneDataCreateSave(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "missions" WHERE id = $1`)).
		WithArgs(DefaultMissionID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "missions" ("id") VALUES ($1)`)).
		WithArgs(DefaultMissionID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	repository := NewMissionRepository()

	waypointComps := []*waypointComponentMock{}
	missionComp := &missionComponentMock{
		id:        string(DefaultMissionID),
		waypoints: waypointComps,
	}
	ms := mission.AssembleFrom(missionComp)

	err = repository.Save(db, ms)

	a.Nil(err)
}

func TestMissionRepositoryMultipleDataCreateSave(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "missions" WHERE id = $1`)).
		WithArgs(DefaultMissionID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "missions" ("id") VALUES ($1)`)).
		WithArgs(DefaultMissionID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "waypoints" ("mission_id","point_order","latitude_degree","longitude_degree","relative_height_m","speed_ms") VALUES ($1,$2,$3,$4,$5,$6),($7,$8,$9,$10,$11,$12),($13,$14,$15,$16,$17,$18)`)).
		WithArgs(
			DefaultMissionID, 1, 11.0, 21.0, 31.0, 41.0,
			DefaultMissionID, 2, 12.0, 22.0, 32.0, 42.0,
			DefaultMissionID, 3, 13.0, 23.0, 33.0, 43.0,
		).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	repository := NewMissionRepository()

	waypointComps := []*waypointComponentMock{
		{
			1, 11.0, 21.0, 31.0, 41.0,
		},
		{
			2, 12.0, 22.0, 32.0, 42.0,
		},
		{
			3, 13.0, 23.0, 33.0, 43.0,
		},
	}
	missionComp := &missionComponentMock{
		id:        string(DefaultMissionID),
		waypoints: waypointComps,
	}
	ms := mission.AssembleFrom(missionComp)

	err = repository.Save(db, ms)

	a.Nil(err)
}

func TestMissionRepositorySingleDataUpdateSave(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "missions" WHERE id = $1`)).
		WithArgs(DefaultMissionID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).
				AddRow(DefaultMissionID),
		)

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "missions" WHERE "id" = $1 ORDER BY "missions"."id" LIMIT 1`)).
		WithArgs(DefaultMissionID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).
				AddRow(DefaultMissionID),
		)

	// mock.
	// 	ExpectExec(
	// 		regexp.QuoteMeta(`UPDATE "missions" SET "id"=$1 WHERE "id" = $2`)).
	// 	WithArgs(DefaultMissionID, DefaultMissionID).
	// 	WillReturnResult(
	// 		sqlmock.NewResult(1, 1),
	// 	)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`DELETE FROM "waypoints" WHERE mission_id = $1`)).
		WithArgs(DefaultMissionID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "waypoints" ("mission_id","point_order","latitude_degree","longitude_degree","relative_height_m","speed_ms") VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(DefaultMissionID, 1, 11.0, 21.0, 31.0, 41.0).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	repository := NewMissionRepository()

	waypointComps := []*waypointComponentMock{
		{
			1, 11.0, 21.0, 31.0, 41.0,
		},
	}
	missionComp := &missionComponentMock{
		id:        string(DefaultMissionID),
		waypoints: waypointComps,
	}
	ms := mission.AssembleFrom(missionComp)

	err = repository.Save(db, ms)

	a.Nil(err)
}

func TestMissionRepositoryNoneDataUpdateSave(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "missions" WHERE id = $1`)).
		WithArgs(DefaultMissionID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).
				AddRow(DefaultMissionID),
		)

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "missions" WHERE "id" = $1 ORDER BY "missions"."id" LIMIT 1`)).
		WithArgs(DefaultMissionID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).
				AddRow(DefaultMissionID),
		)

	// mock.
	// 	ExpectExec(
	// 		regexp.QuoteMeta(`UPDATE "missions" SET "id"=$1 WHERE "id" = $2`)).
	// 	WithArgs(DefaultMissionID, DefaultMissionID).
	// 	WillReturnResult(
	// 		sqlmock.NewResult(1, 1),
	// 	)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`DELETE FROM "waypoints" WHERE mission_id = $1`)).
		WithArgs(DefaultMissionID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	repository := NewMissionRepository()

	waypointComps := []*waypointComponentMock{}
	missionComp := &missionComponentMock{
		id:        string(DefaultMissionID),
		waypoints: waypointComps,
	}
	ms := mission.AssembleFrom(missionComp)

	err = repository.Save(db, ms)

	a.Nil(err)
}

func TestMissionRepositoryMultipleDataUpdateSave(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "missions" WHERE id = $1`)).
		WithArgs(DefaultMissionID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).
				AddRow(DefaultMissionID),
		)

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "missions" WHERE "id" = $1 ORDER BY "missions"."id" LIMIT 1`)).
		WithArgs(DefaultMissionID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).
				AddRow(DefaultMissionID),
		)

	// mock.
	// 	ExpectExec(
	// 		regexp.QuoteMeta(`UPDATE "missions" SET "id"=$1 WHERE "id" = $2`)).
	// 	WithArgs(DefaultMissionID, DefaultMissionID).
	// 	WillReturnResult(
	// 		sqlmock.NewResult(1, 1),
	// 	)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`DELETE FROM "waypoints" WHERE mission_id = $1`)).
		WithArgs(DefaultMissionID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "waypoints" ("mission_id","point_order","latitude_degree","longitude_degree","relative_height_m","speed_ms") VALUES ($1,$2,$3,$4,$5,$6),($7,$8,$9,$10,$11,$12),($13,$14,$15,$16,$17,$18)`)).
		WithArgs(
			DefaultMissionID, 1, 11.0, 21.0, 31.0, 41.0,
			DefaultMissionID, 2, 12.0, 22.0, 32.0, 42.0,
			DefaultMissionID, 3, 13.0, 23.0, 33.0, 43.0,
		).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	repository := NewMissionRepository()

	waypointComps := []*waypointComponentMock{
		{
			1, 11.0, 21.0, 31.0, 41.0,
		},
		{
			2, 12.0, 22.0, 32.0, 42.0,
		},
		{
			3, 13.0, 23.0, 33.0, 43.0,
		},
	}
	missionComp := &missionComponentMock{
		id:        string(DefaultMissionID),
		waypoints: waypointComps,
	}
	ms := mission.AssembleFrom(missionComp)

	err = repository.Save(db, ms)

	a.Nil(err)
}

func TestMissionRepositoryDelete(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "missions" WHERE id = $1`)).
		WithArgs(DefaultMissionID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).
				AddRow(DefaultMissionID),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`DELETE FROM "missions" WHERE "missions"."id" = $1`)).
		WithArgs(DefaultMissionID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`DELETE FROM "waypoints" WHERE mission_id = $1`)).
		WithArgs(DefaultMissionID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	repository := NewMissionRepository()

	err = repository.Delete(db, DefaultMissionID)

	a.Nil(err)
}

func TestMissionRepositoryNotFoundWhenDelete(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "missions" WHERE id = $1`)).
		WithArgs(DefaultMissionID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}),
		)

	repository := NewMissionRepository()

	err = repository.Delete(db, DefaultMissionID)

	a.Equal(err, mission.ErrNotFound)
}
