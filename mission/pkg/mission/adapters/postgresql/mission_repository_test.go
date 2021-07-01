package postgresql

import (
	"mission/pkg/mission/adapters/uuid"
	m "mission/pkg/mission/domain/mission"
	"regexp"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestMissionRepositoryGetSingleWhenGetAll(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "missions"`)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "is_carbon_copy", "version"}).
				AddRow(DefaultMissionID, DefaultMissionName, m.CarbonCopy, DefaultMissionVersion),
		)
	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "navigations" WHERE mission_id = $1`)).
		WithArgs(DefaultMissionID).
		WillReturnRows(
			sqlmock.NewRows([]string{"mission_id", "takeoff_point_ground_height_wgs84_ellipsoid_m"}).
				AddRow(DefaultMissionID, DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM),
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

	gen := uuid.NewMissionUUID()
	repository := NewMissionRepository(gen)

	missions, err := repository.GetAll(db)

	waypointComps := []waypointComponentMock{
		{
			PointOrder:      1,
			LatitudeDegree:  11.0,
			LongitudeDegree: 21.0,
			RelativeHeightM: 31.0,
			SpeedMS:         41.0,
		},
		{
			PointOrder:      2,
			LatitudeDegree:  12.0,
			LongitudeDegree: 22.0,
			RelativeHeightM: 32.0,
			SpeedMS:         42.0,
		},
		{
			PointOrder:      3,
			LatitudeDegree:  13.0,
			LongitudeDegree: 23.0,
			RelativeHeightM: 33.0,
			SpeedMS:         43.0,
		},
	}
	navigationComp := navigationComponentMock{
		TakeoffPointGroundHeightWGS84EllipsoidM: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
		Waypoints:                               waypointComps,
	}
	missionComp := missionComponentMock{
		ID:           string(DefaultMissionID),
		Name:         DefaultMissionName,
		Navigation:   navigationComp,
		IsCarbonCopy: m.CarbonCopy,
		Version:      string(DefaultMissionVersion),
	}
	expectMs := []*m.Mission{
		m.AssembleFrom(
			gen,
			&missionComp,
		),
	}

	a.Nil(err)
	a.Len(missions, 1)
	a.Equal(missions, expectMs)
}

func TestMissionRepositoryGetMultipleWhenGetAll(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	const (
		DefaultMissionID1      = DefaultMissionID + "-1"
		DefaultMissionName1    = DefaultMissionName + "-1"
		DefaultMissionVersion1 = DefaultMissionVersion + "-1"
		DefaultMissionID2      = DefaultMissionID + "-2"
		DefaultMissionName2    = DefaultMissionName + "-2"
		DefaultMissionVersion2 = DefaultMissionVersion + "-2"
		DefaultMissionID3      = DefaultMissionID + "-3"
		DefaultMissionName3    = DefaultMissionName + "-3"
		DefaultMissionVersion3 = DefaultMissionVersion + "-3"
	)

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "missions"`)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "is_carbon_copy", "version"}).
				AddRow(DefaultMissionID1, DefaultMissionName1, m.CarbonCopy, DefaultMissionVersion1).
				AddRow(DefaultMissionID2, DefaultMissionName2, m.CarbonCopy, DefaultMissionVersion2).
				AddRow(DefaultMissionID3, DefaultMissionName3, m.CarbonCopy, DefaultMissionVersion3),
		)
	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "navigations" WHERE mission_id = $1`)).
		WithArgs(DefaultMissionID1).
		WillReturnRows(
			sqlmock.NewRows([]string{"mission_id", "takeoff_point_ground_height_wgs84_ellipsoid_m"}).
				AddRow(DefaultMissionID1, DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM),
		)
	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "waypoints" WHERE mission_id = $1`)).
		WithArgs(DefaultMissionID1).
		WillReturnRows(
			sqlmock.NewRows([]string{"mission_id", "point_order", "latitude_degree", "longitude_degree", "relative_height_m", "speed_ms"}).
				AddRow(DefaultMissionID1, 1, 11.0, 21.0, 31.0, 41.0).
				AddRow(DefaultMissionID1, 2, 12.0, 22.0, 32.0, 42.0).
				AddRow(DefaultMissionID1, 3, 13.0, 23.0, 33.0, 43.0),
		)
	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "navigations" WHERE mission_id = $1`)).
		WithArgs(DefaultMissionID2).
		WillReturnRows(
			sqlmock.NewRows([]string{"mission_id", "takeoff_point_ground_height_wgs84_ellipsoid_m"}).
				AddRow(DefaultMissionID2, DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM),
		)
	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "waypoints" WHERE mission_id = $1`)).
		WithArgs(DefaultMissionID2).
		WillReturnRows(
			sqlmock.NewRows([]string{"mission_id", "point_order", "latitude_degree", "longitude_degree", "relative_height_m", "speed_ms"}).
				AddRow(DefaultMissionID2, 1, 11.0, 21.0, 31.0, 41.0).
				AddRow(DefaultMissionID2, 2, 12.0, 22.0, 32.0, 42.0).
				AddRow(DefaultMissionID2, 3, 13.0, 23.0, 33.0, 43.0),
		)
	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "navigations" WHERE mission_id = $1`)).
		WithArgs(DefaultMissionID3).
		WillReturnRows(
			sqlmock.NewRows([]string{"mission_id", "takeoff_point_ground_height_wgs84_ellipsoid_m"}).
				AddRow(DefaultMissionID3, DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM),
		)
	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "waypoints" WHERE mission_id = $1`)).
		WithArgs(DefaultMissionID3).
		WillReturnRows(
			sqlmock.NewRows([]string{"mission_id", "point_order", "latitude_degree", "longitude_degree", "relative_height_m", "speed_ms"}).
				AddRow(DefaultMissionID3, 1, 11.0, 21.0, 31.0, 41.0).
				AddRow(DefaultMissionID3, 2, 12.0, 22.0, 32.0, 42.0).
				AddRow(DefaultMissionID3, 3, 13.0, 23.0, 33.0, 43.0),
		)

	gen := uuid.NewMissionUUID()
	repository := NewMissionRepository(gen)

	missions, err := repository.GetAll(db)

	waypointComps := []waypointComponentMock{
		{
			PointOrder:      1,
			LatitudeDegree:  11.0,
			LongitudeDegree: 21.0,
			RelativeHeightM: 31.0,
			SpeedMS:         41.0,
		},
		{
			PointOrder:      2,
			LatitudeDegree:  12.0,
			LongitudeDegree: 22.0,
			RelativeHeightM: 32.0,
			SpeedMS:         42.0,
		},
		{
			PointOrder:      3,
			LatitudeDegree:  13.0,
			LongitudeDegree: 23.0,
			RelativeHeightM: 33.0,
			SpeedMS:         43.0,
		},
	}
	navigationComp := navigationComponentMock{
		TakeoffPointGroundHeightWGS84EllipsoidM: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
		Waypoints:                               waypointComps,
	}
	missionComp1 := missionComponentMock{
		ID:           string(DefaultMissionID1),
		Name:         DefaultMissionName1,
		Navigation:   navigationComp,
		IsCarbonCopy: m.CarbonCopy,
		Version:      string(DefaultMissionVersion1),
	}
	missionComp2 := missionComponentMock{
		ID:           string(DefaultMissionID2),
		Name:         DefaultMissionName2,
		Navigation:   navigationComp,
		IsCarbonCopy: m.CarbonCopy,
		Version:      string(DefaultMissionVersion2),
	}
	missionComp3 := missionComponentMock{
		ID:           string(DefaultMissionID3),
		Name:         DefaultMissionName3,
		Navigation:   navigationComp,
		IsCarbonCopy: m.CarbonCopy,
		Version:      string(DefaultMissionVersion3),
	}
	expectMs := []*m.Mission{
		m.AssembleFrom(
			gen,
			&missionComp1,
		),
		m.AssembleFrom(
			gen,
			&missionComp2,
		),
		m.AssembleFrom(
			gen,
			&missionComp3,
		),
	}

	a.Nil(err)
	a.Len(missions, 3)
	a.Equal(missions, expectMs)
}

func TestMissionRepositoryGetNoneWhenGetAll(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "missions"`)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "is_carbon_copy", "version"}),
		)

	gen := uuid.NewMissionUUID()
	repository := NewMissionRepository(gen)

	missions, err := repository.GetAll(db)

	a.Nil(err)
	a.Len(missions, 0)
}

func TestMissionRepositoryGetSingleWhenGetAllOrigin(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "missions" WHERE is_carbon_copy = false`)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "is_carbon_copy", "version"}).
				AddRow(DefaultMissionID, DefaultMissionName, m.CarbonCopy, DefaultMissionVersion),
		)
	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "navigations" WHERE mission_id = $1`)).
		WithArgs(DefaultMissionID).
		WillReturnRows(
			sqlmock.NewRows([]string{"mission_id", "takeoff_point_ground_height_wgs84_ellipsoid_m"}).
				AddRow(DefaultMissionID, DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM),
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

	gen := uuid.NewMissionUUID()
	repository := NewMissionRepository(gen)

	missions, err := repository.GetAllOrigin(db)

	waypointComps := []waypointComponentMock{
		{
			PointOrder:      1,
			LatitudeDegree:  11.0,
			LongitudeDegree: 21.0,
			RelativeHeightM: 31.0,
			SpeedMS:         41.0,
		},
		{
			PointOrder:      2,
			LatitudeDegree:  12.0,
			LongitudeDegree: 22.0,
			RelativeHeightM: 32.0,
			SpeedMS:         42.0,
		},
		{
			PointOrder:      3,
			LatitudeDegree:  13.0,
			LongitudeDegree: 23.0,
			RelativeHeightM: 33.0,
			SpeedMS:         43.0,
		},
	}
	navigationComp := navigationComponentMock{
		TakeoffPointGroundHeightWGS84EllipsoidM: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
		Waypoints:                               waypointComps,
	}
	missionComp := missionComponentMock{
		ID:           string(DefaultMissionID),
		Name:         DefaultMissionName,
		Navigation:   navigationComp,
		IsCarbonCopy: m.CarbonCopy,
		Version:      string(DefaultMissionVersion),
	}
	expectMs := []*m.Mission{
		m.AssembleFrom(
			gen,
			&missionComp,
		),
	}

	a.Nil(err)
	a.Len(missions, 1)
	a.Equal(missions, expectMs)
}

func TestMissionRepositoryGetMultipleWhenGetAllOrigin(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	const (
		DefaultMissionID1      = DefaultMissionID + "-1"
		DefaultMissionName1    = DefaultMissionName + "-1"
		DefaultMissionVersion1 = DefaultMissionVersion + "-1"
		DefaultMissionID2      = DefaultMissionID + "-2"
		DefaultMissionName2    = DefaultMissionName + "-2"
		DefaultMissionVersion2 = DefaultMissionVersion + "-2"
		DefaultMissionID3      = DefaultMissionID + "-3"
		DefaultMissionName3    = DefaultMissionName + "-3"
		DefaultMissionVersion3 = DefaultMissionVersion + "-3"
	)

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "missions" WHERE is_carbon_copy = false`)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "is_carbon_copy", "version"}).
				AddRow(DefaultMissionID1, DefaultMissionName1, m.CarbonCopy, DefaultMissionVersion1).
				AddRow(DefaultMissionID2, DefaultMissionName2, m.CarbonCopy, DefaultMissionVersion2).
				AddRow(DefaultMissionID3, DefaultMissionName3, m.CarbonCopy, DefaultMissionVersion3),
		)
	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "navigations" WHERE mission_id = $1`)).
		WithArgs(DefaultMissionID1).
		WillReturnRows(
			sqlmock.NewRows([]string{"mission_id", "takeoff_point_ground_height_wgs84_ellipsoid_m"}).
				AddRow(DefaultMissionID1, DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM),
		)
	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "waypoints" WHERE mission_id = $1`)).
		WithArgs(DefaultMissionID1).
		WillReturnRows(
			sqlmock.NewRows([]string{"mission_id", "point_order", "latitude_degree", "longitude_degree", "relative_height_m", "speed_ms"}).
				AddRow(DefaultMissionID1, 1, 11.0, 21.0, 31.0, 41.0).
				AddRow(DefaultMissionID1, 2, 12.0, 22.0, 32.0, 42.0).
				AddRow(DefaultMissionID1, 3, 13.0, 23.0, 33.0, 43.0),
		)
	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "navigations" WHERE mission_id = $1`)).
		WithArgs(DefaultMissionID2).
		WillReturnRows(
			sqlmock.NewRows([]string{"mission_id", "takeoff_point_ground_height_wgs84_ellipsoid_m"}).
				AddRow(DefaultMissionID2, DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM),
		)
	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "waypoints" WHERE mission_id = $1`)).
		WithArgs(DefaultMissionID2).
		WillReturnRows(
			sqlmock.NewRows([]string{"mission_id", "point_order", "latitude_degree", "longitude_degree", "relative_height_m", "speed_ms"}).
				AddRow(DefaultMissionID2, 1, 11.0, 21.0, 31.0, 41.0).
				AddRow(DefaultMissionID2, 2, 12.0, 22.0, 32.0, 42.0).
				AddRow(DefaultMissionID2, 3, 13.0, 23.0, 33.0, 43.0),
		)
	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "navigations" WHERE mission_id = $1`)).
		WithArgs(DefaultMissionID3).
		WillReturnRows(
			sqlmock.NewRows([]string{"mission_id", "takeoff_point_ground_height_wgs84_ellipsoid_m"}).
				AddRow(DefaultMissionID3, DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM),
		)
	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "waypoints" WHERE mission_id = $1`)).
		WithArgs(DefaultMissionID3).
		WillReturnRows(
			sqlmock.NewRows([]string{"mission_id", "point_order", "latitude_degree", "longitude_degree", "relative_height_m", "speed_ms"}).
				AddRow(DefaultMissionID3, 1, 11.0, 21.0, 31.0, 41.0).
				AddRow(DefaultMissionID3, 2, 12.0, 22.0, 32.0, 42.0).
				AddRow(DefaultMissionID3, 3, 13.0, 23.0, 33.0, 43.0),
		)

	gen := uuid.NewMissionUUID()
	repository := NewMissionRepository(gen)

	missions, err := repository.GetAllOrigin(db)

	waypointComps := []waypointComponentMock{
		{
			PointOrder:      1,
			LatitudeDegree:  11.0,
			LongitudeDegree: 21.0,
			RelativeHeightM: 31.0,
			SpeedMS:         41.0,
		},
		{
			PointOrder:      2,
			LatitudeDegree:  12.0,
			LongitudeDegree: 22.0,
			RelativeHeightM: 32.0,
			SpeedMS:         42.0,
		},
		{
			PointOrder:      3,
			LatitudeDegree:  13.0,
			LongitudeDegree: 23.0,
			RelativeHeightM: 33.0,
			SpeedMS:         43.0,
		},
	}
	navigationComp := navigationComponentMock{
		TakeoffPointGroundHeightWGS84EllipsoidM: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
		Waypoints:                               waypointComps,
	}
	missionComp1 := missionComponentMock{
		ID:           string(DefaultMissionID1),
		Name:         DefaultMissionName1,
		Navigation:   navigationComp,
		IsCarbonCopy: m.CarbonCopy,
		Version:      string(DefaultMissionVersion1),
	}
	missionComp2 := missionComponentMock{
		ID:           string(DefaultMissionID2),
		Name:         DefaultMissionName2,
		Navigation:   navigationComp,
		IsCarbonCopy: m.CarbonCopy,
		Version:      string(DefaultMissionVersion2),
	}
	missionComp3 := missionComponentMock{
		ID:           string(DefaultMissionID3),
		Name:         DefaultMissionName3,
		Navigation:   navigationComp,
		IsCarbonCopy: m.CarbonCopy,
		Version:      string(DefaultMissionVersion3),
	}
	expectMs := []*m.Mission{
		m.AssembleFrom(
			gen,
			&missionComp1,
		),
		m.AssembleFrom(
			gen,
			&missionComp2,
		),
		m.AssembleFrom(
			gen,
			&missionComp3,
		),
	}

	a.Nil(err)
	a.Len(missions, 3)
	a.Equal(missions, expectMs)
}

func TestMissionRepositoryGetNoneWhenGetAllOrigin(t *testing.T) {
	a := assert.New(t)

	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "missions" WHERE is_carbon_copy = false`)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "is_carbon_copy", "version"}),
		)

	gen := uuid.NewMissionUUID()
	repository := NewMissionRepository(gen)

	missions, err := repository.GetAllOrigin(db)

	a.Nil(err)
	a.Len(missions, 0)
}

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
			sqlmock.NewRows([]string{"id", "name", "is_carbon_copy", "version"}).
				AddRow(DefaultMissionID, DefaultMissionName, m.CarbonCopy, DefaultMissionVersion),
		)
	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "navigations" WHERE mission_id = $1`)).
		WithArgs(DefaultMissionID).
		WillReturnRows(
			sqlmock.NewRows([]string{"mission_id", "takeoff_point_ground_height_wgs84_ellipsoid_m"}).
				AddRow(DefaultMissionID, DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM),
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

	gen := uuid.NewMissionUUID()
	repository := NewMissionRepository(gen)

	mission, err := repository.GetByID(db, DefaultMissionID)

	waypointComps := []waypointComponentMock{
		{
			PointOrder:      1,
			LatitudeDegree:  11.0,
			LongitudeDegree: 21.0,
			RelativeHeightM: 31.0,
			SpeedMS:         41.0,
		},
		{
			PointOrder:      2,
			LatitudeDegree:  12.0,
			LongitudeDegree: 22.0,
			RelativeHeightM: 32.0,
			SpeedMS:         42.0,
		},
		{
			PointOrder:      3,
			LatitudeDegree:  13.0,
			LongitudeDegree: 23.0,
			RelativeHeightM: 33.0,
			SpeedMS:         43.0,
		},
	}
	navigationComp := navigationComponentMock{
		TakeoffPointGroundHeightWGS84EllipsoidM: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
		Waypoints:                               waypointComps,
	}
	missionComp := missionComponentMock{
		ID:           string(DefaultMissionID),
		Name:         DefaultMissionName,
		Navigation:   navigationComp,
		IsCarbonCopy: m.CarbonCopy,
		Version:      string(DefaultMissionVersion),
	}
	expectM := m.AssembleFrom(
		gen,
		&missionComp,
	)

	a.Nil(err)
	a.Equal(mission, expectM)
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
			sqlmock.NewRows([]string{"id", "name", "is_carbon_copy", "version"}),
		)

	gen := uuid.NewMissionUUID()
	repository := NewMissionRepository(gen)

	mission, err := repository.GetByID(db, DefaultMissionID)

	a.Nil(mission)
	a.Equal(err, m.ErrNotFound)
}

func TestMissionRepositorySingleWaypointCreateSave(t *testing.T) {
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
			sqlmock.NewRows([]string{"id", "name", "is_carbon_copy", "version"}),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "missions" ("id","name","is_carbon_copy","version") VALUES ($1,$2,$3,$4)`)).
		WithArgs(DefaultMissionID, DefaultMissionName, m.CarbonCopy, DefaultMissionVersion).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "navigations" ("mission_id","takeoff_point_ground_height_wgs84_ellipsoid_m") VALUES ($1,$2)`)).
		WithArgs(DefaultMissionID, DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM).
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

	gen := uuid.NewMissionUUID()
	repository := NewMissionRepository(gen)

	waypointComps := []waypointComponentMock{
		{
			PointOrder:      1,
			LatitudeDegree:  11.0,
			LongitudeDegree: 21.0,
			RelativeHeightM: 31.0,
			SpeedMS:         41.0,
		},
	}
	navigationComp := navigationComponentMock{
		TakeoffPointGroundHeightWGS84EllipsoidM: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
		Waypoints:                               waypointComps,
	}
	missionComp := missionComponentMock{
		ID:           string(DefaultMissionID),
		Name:         DefaultMissionName,
		Navigation:   navigationComp,
		IsCarbonCopy: m.CarbonCopy,
		Version:      string(DefaultMissionVersion),
	}
	mission := m.AssembleFrom(
		gen,
		&missionComp,
	)

	err = repository.Save(db, mission)

	a.Nil(err)
}

func TestMissionRepositoryNoWaypointCreateSave(t *testing.T) {
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
			sqlmock.NewRows([]string{"id", "name", "is_carbon_copy", "version"}),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "missions" ("id","name","is_carbon_copy","version") VALUES ($1,$2,$3,$4)`)).
		WithArgs(DefaultMissionID, DefaultMissionName, m.CarbonCopy, DefaultMissionVersion).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "navigations" ("mission_id","takeoff_point_ground_height_wgs84_ellipsoid_m") VALUES ($1,$2)`)).
		WithArgs(DefaultMissionID, DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	gen := uuid.NewMissionUUID()
	repository := NewMissionRepository(gen)

	waypointComps := []waypointComponentMock{}
	navigationComp := navigationComponentMock{
		TakeoffPointGroundHeightWGS84EllipsoidM: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
		Waypoints:                               waypointComps,
	}
	missionComp := missionComponentMock{
		ID:           string(DefaultMissionID),
		Name:         DefaultMissionName,
		Navigation:   navigationComp,
		IsCarbonCopy: m.CarbonCopy,
		Version:      string(DefaultMissionVersion),
	}
	mission := m.AssembleFrom(
		gen,
		&missionComp,
	)

	err = repository.Save(db, mission)

	a.Nil(err)
}

func TestMissionRepositoryMultipleWaypointsCreateSave(t *testing.T) {
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
			sqlmock.NewRows([]string{"id", "name", "is_carbon_copy", "version"}),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "missions" ("id","name","is_carbon_copy","version") VALUES ($1,$2,$3,$4)`)).
		WithArgs(DefaultMissionID, DefaultMissionName, m.CarbonCopy, DefaultMissionVersion).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "navigations" ("mission_id","takeoff_point_ground_height_wgs84_ellipsoid_m") VALUES ($1,$2)`)).
		WithArgs(DefaultMissionID, DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM).
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

	gen := uuid.NewMissionUUID()
	repository := NewMissionRepository(gen)

	waypointComps := []waypointComponentMock{
		{
			PointOrder:      1,
			LatitudeDegree:  11.0,
			LongitudeDegree: 21.0,
			RelativeHeightM: 31.0,
			SpeedMS:         41.0,
		},
		{
			PointOrder:      2,
			LatitudeDegree:  12.0,
			LongitudeDegree: 22.0,
			RelativeHeightM: 32.0,
			SpeedMS:         42.0,
		},
		{
			PointOrder:      3,
			LatitudeDegree:  13.0,
			LongitudeDegree: 23.0,
			RelativeHeightM: 33.0,
			SpeedMS:         43.0,
		},
	}
	navigationComp := navigationComponentMock{
		TakeoffPointGroundHeightWGS84EllipsoidM: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
		Waypoints:                               waypointComps,
	}
	missionComp := missionComponentMock{
		ID:           string(DefaultMissionID),
		Name:         DefaultMissionName,
		Navigation:   navigationComp,
		IsCarbonCopy: m.CarbonCopy,
		Version:      string(DefaultMissionVersion),
	}
	mission := m.AssembleFrom(
		gen,
		&missionComp,
	)

	err = repository.Save(db, mission)

	a.Nil(err)
}

func TestMissionRepositorySingleWaypointUpdateSave(t *testing.T) {
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
			sqlmock.NewRows([]string{"id", "name", "is_carbon_copy", "version"}).
				AddRow(DefaultMissionID, DefaultMissionName, m.CarbonCopy, DefaultMissionVersion),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`UPDATE "missions" SET "name"=$1,"is_carbon_copy"=$2,"version"=$3 WHERE "id" = $4`)).
		WithArgs(DefaultMissionName, m.CarbonCopy, DefaultMissionVersion, DefaultMissionID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`DELETE FROM "navigations" WHERE mission_id = $1`)).
		WithArgs(DefaultMissionID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "navigations" ("mission_id","takeoff_point_ground_height_wgs84_ellipsoid_m") VALUES ($1,$2)`)).
		WithArgs(DefaultMissionID, DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM).
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

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "waypoints" ("mission_id","point_order","latitude_degree","longitude_degree","relative_height_m","speed_ms") VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(DefaultMissionID, 1, 11.0, 21.0, 31.0, 41.0).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	gen := uuid.NewMissionUUID()
	repository := NewMissionRepository(gen)

	waypointComps := []waypointComponentMock{
		{
			PointOrder:      1,
			LatitudeDegree:  11.0,
			LongitudeDegree: 21.0,
			RelativeHeightM: 31.0,
			SpeedMS:         41.0,
		},
	}
	navigationComp := navigationComponentMock{
		TakeoffPointGroundHeightWGS84EllipsoidM: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
		Waypoints:                               waypointComps,
	}
	missionComp := missionComponentMock{
		ID:           string(DefaultMissionID),
		Name:         DefaultMissionName,
		Navigation:   navigationComp,
		IsCarbonCopy: m.CarbonCopy,
		Version:      string(DefaultMissionVersion),
	}
	mission := m.AssembleFrom(
		gen,
		&missionComp,
	)

	err = repository.Save(db, mission)

	a.Nil(err)
}

func TestMissionRepositoryNoWaypointUpdateSave(t *testing.T) {
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
			sqlmock.NewRows([]string{"id", "name", "is_carbon_copy", "version"}).
				AddRow(DefaultMissionID, DefaultMissionName, m.CarbonCopy, DefaultMissionVersion),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`UPDATE "missions" SET "name"=$1,"is_carbon_copy"=$2,"version"=$3 WHERE "id" = $4`)).
		WithArgs(DefaultMissionName, m.CarbonCopy, DefaultMissionVersion, DefaultMissionID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`DELETE FROM "navigations" WHERE mission_id = $1`)).
		WithArgs(DefaultMissionID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "navigations" ("mission_id","takeoff_point_ground_height_wgs84_ellipsoid_m") VALUES ($1,$2)`)).
		WithArgs(DefaultMissionID, DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM).
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

	gen := uuid.NewMissionUUID()
	repository := NewMissionRepository(gen)

	waypointComps := []waypointComponentMock{}
	navigationComp := navigationComponentMock{
		TakeoffPointGroundHeightWGS84EllipsoidM: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
		Waypoints:                               waypointComps,
	}
	missionComp := missionComponentMock{
		ID:           string(DefaultMissionID),
		Name:         DefaultMissionName,
		Navigation:   navigationComp,
		IsCarbonCopy: m.CarbonCopy,
		Version:      string(DefaultMissionVersion),
	}
	mission := m.AssembleFrom(
		gen,
		&missionComp,
	)

	err = repository.Save(db, mission)

	a.Nil(err)
}

func TestMissionRepositoryMultipleWaypointsUpdateSave(t *testing.T) {
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
			sqlmock.NewRows([]string{"id", "name", "is_carbon_copy", "version"}).
				AddRow(DefaultMissionID, DefaultMissionName, m.CarbonCopy, DefaultMissionVersion),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`UPDATE "missions" SET "name"=$1,"is_carbon_copy"=$2,"version"=$3 WHERE "id" = $4`)).
		WithArgs(DefaultMissionName, m.CarbonCopy, DefaultMissionVersion, DefaultMissionID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`DELETE FROM "navigations" WHERE mission_id = $1`)).
		WithArgs(DefaultMissionID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "navigations" ("mission_id","takeoff_point_ground_height_wgs84_ellipsoid_m") VALUES ($1,$2)`)).
		WithArgs(DefaultMissionID, DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM).
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

	gen := uuid.NewMissionUUID()
	repository := NewMissionRepository(gen)

	waypointComps := []waypointComponentMock{
		{
			PointOrder:      1,
			LatitudeDegree:  11.0,
			LongitudeDegree: 21.0,
			RelativeHeightM: 31.0,
			SpeedMS:         41.0,
		},
		{
			PointOrder:      2,
			LatitudeDegree:  12.0,
			LongitudeDegree: 22.0,
			RelativeHeightM: 32.0,
			SpeedMS:         42.0,
		},
		{
			PointOrder:      3,
			LatitudeDegree:  13.0,
			LongitudeDegree: 23.0,
			RelativeHeightM: 33.0,
			SpeedMS:         43.0,
		},
	}
	navigationComp := navigationComponentMock{
		TakeoffPointGroundHeightWGS84EllipsoidM: DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM,
		Waypoints:                               waypointComps,
	}
	missionComp := missionComponentMock{
		ID:           string(DefaultMissionID),
		Name:         DefaultMissionName,
		Navigation:   navigationComp,
		IsCarbonCopy: m.CarbonCopy,
		Version:      string(DefaultMissionVersion),
	}
	mission := m.AssembleFrom(
		gen,
		&missionComp,
	)

	err = repository.Save(db, mission)

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
			sqlmock.NewRows([]string{"id", "name", "is_carbon_copy", "version"}).
				AddRow(DefaultMissionID, DefaultMissionName, m.CarbonCopy, DefaultMissionVersion),
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
			regexp.QuoteMeta(`DELETE FROM "navigations" WHERE mission_id = $1`)).
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

	gen := uuid.NewMissionUUID()
	repository := NewMissionRepository(gen)

	err = repository.Delete(db, DefaultMissionID)

	a.Nil(err)
}
