package postgresql

import (
	act "collection-analysis/pkg/action/domain/action"
	"regexp"
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/adapters/postgresql"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestActionRepositoryGetByID(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "actions" WHERE id = $1`)).
		WithArgs(DefaultActionID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "communication_id", "fleet_id", "is_completed"}).
				AddRow(DefaultActionID, DefaultActionCommunicationID, DefaultActionFleetID, act.Completed),
		)
	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "trajectory_points" WHERE action_id = $1`)).
		WithArgs(DefaultActionID).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"action_id",
				"point_number",
				"latitude",
				"longitude",
				"altitude",
				"relative_altitude",
				"speed",
				"armed",
				"flight_mode",
				"orientation_x",
				"orientation_y",
				"orientation_z",
				"orientation_w"}).
				AddRow(
					DefaultActionID,
					1,
					1.0,
					2.0,
					3.0,
					4.0,
					5.0,
					true,
					"state",
					6.0,
					7.0,
					8.0,
					9.0),
		)

	repository := NewActionRepository()

	action, err := repository.GetByID(db, DefaultActionID)

	trajectoryPointComps := []act.TrajectoryPointComponent{
		&trajectoryPointComponentMock{
			1,
			1.0,
			2.0,
			3.0,
			4.0,
			5.0,
			true,
			"state",
			6.0,
			7.0,
			8.0,
			9.0,
		},
	}
	actionComp := actionComponentMock{
		ID:               string(DefaultActionID),
		CommunicationID:  string(DefaultActionCommunicationID),
		FleetID:          string(DefaultActionFleetID),
		IsCompleted:      act.Completed,
		TrajectoryPoints: trajectoryPointComps,
	}
	expectAct := act.AssembleFrom(
		&actionComp,
	)

	a.Nil(err)
	a.Equal(action, expectAct)
}

func TestActionRepositoryNotFoundWhenGetByID(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "actions" WHERE id = $1`)).
		WithArgs(DefaultActionID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "communication_id", "fleet_id", "is_completed"}),
		)

	repository := NewActionRepository()

	action, err := repository.GetByID(db, DefaultActionID)

	a.Nil(action)
	a.Equal(err, act.ErrNotFound)
}

func TestActionRepositoryGetSingleWhenGetAllActiveByFleetID(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "actions" WHERE fleet_id = $1 AND is_completed = false`)).
		WithArgs(DefaultActionFleetID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "communication_id", "fleet_id", "is_completed"}).
				AddRow(DefaultActionID, DefaultActionCommunicationID, DefaultActionFleetID, act.Completed),
		)
	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "trajectory_points" WHERE action_id = $1`)).
		WithArgs(DefaultActionID).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"action_id",
				"point_number",
				"latitude",
				"longitude",
				"altitude",
				"relative_altitude",
				"speed",
				"armed",
				"flight_mode",
				"orientation_x",
				"orientation_y",
				"orientation_z",
				"orientation_w"}).
				AddRow(
					DefaultActionID,
					1,
					1.0,
					2.0,
					3.0,
					4.0,
					5.0,
					true,
					"state",
					6.0,
					7.0,
					8.0,
					9.0),
		)

	repository := NewActionRepository()

	actions, err := repository.GetAllActiveByFleetID(db, DefaultActionFleetID)

	trajectoryPointComps := []act.TrajectoryPointComponent{
		&trajectoryPointComponentMock{
			1,
			1.0,
			2.0,
			3.0,
			4.0,
			5.0,
			true,
			"state",
			6.0,
			7.0,
			8.0,
			9.0,
		},
	}
	actionComp := actionComponentMock{
		ID:               string(DefaultActionID),
		CommunicationID:  string(DefaultActionCommunicationID),
		FleetID:          string(DefaultActionFleetID),
		IsCompleted:      act.Completed,
		TrajectoryPoints: trajectoryPointComps,
	}
	expectAct := act.AssembleFrom(
		&actionComp,
	)
	expectActs := []*act.Action{expectAct}

	a.Nil(err)
	a.Equal(actions, expectActs)
}

func TestActionRepositoryGetMultipleWhenGetAllActiveByFleetID(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultActionID1 = DefaultActionID + "-1"
		DefaultActionID2 = DefaultActionID + "-2"
		DefaultActionID3 = DefaultActionID + "-3"
	)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "actions" WHERE fleet_id = $1 AND is_completed = false`)).
		WithArgs(DefaultActionFleetID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "communication_id", "fleet_id", "is_completed"}).
				AddRow(DefaultActionID1, DefaultActionCommunicationID, DefaultActionFleetID, act.Completed).
				AddRow(DefaultActionID2, DefaultActionCommunicationID, DefaultActionFleetID, act.Completed).
				AddRow(DefaultActionID3, DefaultActionCommunicationID, DefaultActionFleetID, act.Completed),
		)
	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "trajectory_points" WHERE action_id = $1`)).
		WithArgs(DefaultActionID1).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"action_id",
				"point_number",
				"latitude",
				"longitude",
				"altitude",
				"relative_altitude",
				"speed",
				"armed",
				"flight_mode",
				"orientation_x",
				"orientation_y",
				"orientation_z",
				"orientation_w"}).
				AddRow(
					DefaultActionID1,
					1,
					1.0,
					2.0,
					3.0,
					4.0,
					5.0,
					true,
					"state",
					6.0,
					7.0,
					8.0,
					9.0),
		)
	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "trajectory_points" WHERE action_id = $1`)).
		WithArgs(DefaultActionID2).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"action_id",
				"point_number",
				"latitude",
				"longitude",
				"altitude",
				"relative_altitude",
				"speed",
				"armed",
				"flight_mode",
				"orientation_x",
				"orientation_y",
				"orientation_z",
				"orientation_w"}).
				AddRow(
					DefaultActionID2,
					1,
					1.0,
					2.0,
					3.0,
					4.0,
					5.0,
					true,
					"state",
					6.0,
					7.0,
					8.0,
					9.0),
		)
	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "trajectory_points" WHERE action_id = $1`)).
		WithArgs(DefaultActionID3).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"action_id",
				"point_number",
				"latitude",
				"longitude",
				"altitude",
				"relative_altitude",
				"speed",
				"armed",
				"flight_mode",
				"orientation_x",
				"orientation_y",
				"orientation_z",
				"orientation_w"}).
				AddRow(
					DefaultActionID3,
					1,
					1.0,
					2.0,
					3.0,
					4.0,
					5.0,
					true,
					"state",
					6.0,
					7.0,
					8.0,
					9.0),
		)

	repository := NewActionRepository()

	actions, err := repository.GetAllActiveByFleetID(db, DefaultActionFleetID)

	expectActs := []*act.Action{
		act.AssembleFrom(
			&actionComponentMock{
				ID:              string(DefaultActionID1),
				CommunicationID: string(DefaultActionCommunicationID),
				FleetID:         string(DefaultActionFleetID),
				IsCompleted:     act.Completed,
				TrajectoryPoints: []act.TrajectoryPointComponent{
					&trajectoryPointComponentMock{
						1,
						1.0,
						2.0,
						3.0,
						4.0,
						5.0,
						true,
						"state",
						6.0,
						7.0,
						8.0,
						9.0,
					},
				},
			},
		),
		act.AssembleFrom(
			&actionComponentMock{
				ID:              string(DefaultActionID2),
				CommunicationID: string(DefaultActionCommunicationID),
				FleetID:         string(DefaultActionFleetID),
				IsCompleted:     act.Completed,
				TrajectoryPoints: []act.TrajectoryPointComponent{
					&trajectoryPointComponentMock{
						1,
						1.0,
						2.0,
						3.0,
						4.0,
						5.0,
						true,
						"state",
						6.0,
						7.0,
						8.0,
						9.0,
					},
				},
			},
		),
		act.AssembleFrom(
			&actionComponentMock{
				ID:              string(DefaultActionID3),
				CommunicationID: string(DefaultActionCommunicationID),
				FleetID:         string(DefaultActionFleetID),
				IsCompleted:     act.Completed,
				TrajectoryPoints: []act.TrajectoryPointComponent{
					&trajectoryPointComponentMock{
						1,
						1.0,
						2.0,
						3.0,
						4.0,
						5.0,
						true,
						"state",
						6.0,
						7.0,
						8.0,
						9.0,
					},
				},
			},
		),
	}

	a.Nil(err)
	a.Equal(actions, expectActs)
}

func TestActionRepositoryGetNoneWhenGetAllActiveByFleetID(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "actions" WHERE fleet_id = $1 AND is_completed = false`)).
		WithArgs(DefaultActionFleetID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "communication_id", "fleet_id", "is_completed"}),
		)

	repository := NewActionRepository()

	actions, err := repository.GetAllActiveByFleetID(db, DefaultActionFleetID)

	var expectActs []*act.Action

	a.Nil(err)
	a.Equal(actions, expectActs)
}

func TestActionRepositoryGetActiveByCommunicationID(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "actions" WHERE communication_id = $1 AND is_completed = false`)).
		WithArgs(DefaultActionCommunicationID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "communication_id", "fleet_id", "is_completed"}).
				AddRow(DefaultActionID, DefaultActionCommunicationID, DefaultActionFleetID, act.Completed),
		)
	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "trajectory_points" WHERE action_id = $1`)).
		WithArgs(DefaultActionID).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"action_id",
				"point_number",
				"latitude",
				"longitude",
				"altitude",
				"relative_altitude",
				"speed",
				"armed",
				"flight_mode",
				"orientation_x",
				"orientation_y",
				"orientation_z",
				"orientation_w"}).
				AddRow(
					DefaultActionID,
					1,
					1.0,
					2.0,
					3.0,
					4.0,
					5.0,
					true,
					"state",
					6.0,
					7.0,
					8.0,
					9.0),
		)

	repository := NewActionRepository()

	action, err := repository.GetActiveByCommunicationID(db, DefaultActionCommunicationID)

	trajectoryPointComps := []act.TrajectoryPointComponent{
		&trajectoryPointComponentMock{
			1,
			1.0,
			2.0,
			3.0,
			4.0,
			5.0,
			true,
			"state",
			6.0,
			7.0,
			8.0,
			9.0,
		},
	}
	actionComp := actionComponentMock{
		ID:               string(DefaultActionID),
		CommunicationID:  string(DefaultActionCommunicationID),
		FleetID:          string(DefaultActionFleetID),
		IsCompleted:      act.Completed,
		TrajectoryPoints: trajectoryPointComps,
	}
	expectAct := act.AssembleFrom(
		&actionComp,
	)

	a.Nil(err)
	a.Equal(action, expectAct)
}

func TestActionRepositoryNotFoundWhenGetActiveByCommunicationID(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "actions" WHERE communication_id = $1 AND is_completed = false`)).
		WithArgs(DefaultActionCommunicationID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "communication_id", "fleet_id", "is_completed"}),
		)

	repository := NewActionRepository()

	action, err := repository.GetActiveByCommunicationID(db, DefaultActionCommunicationID)

	a.Nil(action)
	a.Equal(err, act.ErrNotFound)
}

func TestActionRepositorySingleTrajectoryPointDataCreateSave(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "actions" WHERE id = $1`)).
		WithArgs(DefaultActionID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "communication_id", "fleet_id", "is_completed"}),
		)

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT count(1) FROM "trajectory_points" WHERE action_id = $1`)).
		WithArgs(DefaultActionID).
		WillReturnRows(
			sqlmock.NewRows([]string{"count(1)"}).
				AddRow(0),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "actions" ("id","communication_id","fleet_id","is_completed") VALUES ($1,$2,$3,$4)`)).
		WithArgs(DefaultActionID, DefaultActionCommunicationID, DefaultActionFleetID, act.Completed).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "trajectory_points" ("action_id","point_number","latitude","longitude","altitude","relative_altitude","speed","armed","flight_mode","orientation_x","orientation_y","orientation_z","orientation_w") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)`)).
		WithArgs(
			DefaultActionID,
			1,
			1.0,
			2.0,
			3.0,
			4.0,
			5.0,
			true,
			"state",
			6.0,
			7.0,
			8.0,
			9.0).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	repository := NewActionRepository()

	trajectoryPointComps := []act.TrajectoryPointComponent{
		&trajectoryPointComponentMock{
			1,
			1.0,
			2.0,
			3.0,
			4.0,
			5.0,
			true,
			"state",
			6.0,
			7.0,
			8.0,
			9.0,
		},
	}
	actionComp := actionComponentMock{
		ID:               string(DefaultActionID),
		CommunicationID:  string(DefaultActionCommunicationID),
		FleetID:          string(DefaultActionFleetID),
		IsCompleted:      act.Completed,
		TrajectoryPoints: trajectoryPointComps,
	}
	action := act.AssembleFrom(
		&actionComp,
	)

	err = repository.Save(db, action)

	a.Nil(err)
}

func TestActionRepositoryNoneTrajectoryPointDataCreateSave(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "actions" WHERE id = $1`)).
		WithArgs(DefaultActionID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "communication_id", "fleet_id", "is_completed"}),
		)

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT count(1) FROM "trajectory_points" WHERE action_id = $1`)).
		WithArgs(DefaultActionID).
		WillReturnRows(
			sqlmock.NewRows([]string{"count(1)"}).
				AddRow(0),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "actions" ("id","communication_id","fleet_id","is_completed") VALUES ($1,$2,$3,$4)`)).
		WithArgs(DefaultActionID, DefaultActionCommunicationID, DefaultActionFleetID, act.Completed).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	repository := NewActionRepository()

	trajectoryPointComps := []act.TrajectoryPointComponent{}
	actionComp := actionComponentMock{
		ID:               string(DefaultActionID),
		CommunicationID:  string(DefaultActionCommunicationID),
		FleetID:          string(DefaultActionFleetID),
		IsCompleted:      act.Completed,
		TrajectoryPoints: trajectoryPointComps,
	}
	action := act.AssembleFrom(
		&actionComp,
	)

	err = repository.Save(db, action)

	a.Nil(err)
}

func TestActionRepositoryMultipleTrajectoryPointDataCreateSave(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "actions" WHERE id = $1`)).
		WithArgs(DefaultActionID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "communication_id", "fleet_id", "is_completed"}),
		)

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT count(1) FROM "trajectory_points" WHERE action_id = $1`)).
		WithArgs(DefaultActionID).
		WillReturnRows(
			sqlmock.NewRows([]string{"count(1)"}).
				AddRow(0),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "actions" ("id","communication_id","fleet_id","is_completed") VALUES ($1,$2,$3,$4)`)).
		WithArgs(DefaultActionID, DefaultActionCommunicationID, DefaultActionFleetID, act.Completed).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "trajectory_points" ("action_id","point_number","latitude","longitude","altitude","relative_altitude","speed","armed","flight_mode","orientation_x","orientation_y","orientation_z","orientation_w") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13),($14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26)`)).
		WithArgs(
			DefaultActionID,
			1,
			1.0,
			2.0,
			3.0,
			4.0,
			5.0,
			true,
			"state",
			6.0,
			7.0,
			8.0,
			9.0,
			DefaultActionID,
			2,
			1.0,
			2.0,
			3.0,
			4.0,
			5.0,
			true,
			"state",
			6.0,
			7.0,
			8.0,
			9.0).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	repository := NewActionRepository()

	trajectoryPointComps := []act.TrajectoryPointComponent{
		&trajectoryPointComponentMock{
			1,
			1.0,
			2.0,
			3.0,
			4.0,
			5.0,
			true,
			"state",
			6.0,
			7.0,
			8.0,
			9.0,
		},
		&trajectoryPointComponentMock{
			2,
			1.0,
			2.0,
			3.0,
			4.0,
			5.0,
			true,
			"state",
			6.0,
			7.0,
			8.0,
			9.0,
		},
	}
	actionComp := actionComponentMock{
		ID:               string(DefaultActionID),
		CommunicationID:  string(DefaultActionCommunicationID),
		FleetID:          string(DefaultActionFleetID),
		IsCompleted:      act.Completed,
		TrajectoryPoints: trajectoryPointComps,
	}
	action := act.AssembleFrom(
		&actionComp,
	)

	err = repository.Save(db, action)

	a.Nil(err)
}

func TestActionRepositorySingleTrajectoryPointDataUpdateSave_NonePreTrajectoryPoint(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "actions" WHERE id = $1`)).
		WithArgs(DefaultActionID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "communication_id", "fleet_id", "is_completed"}).
				AddRow(DefaultActionID, DefaultActionCommunicationID, DefaultActionFleetID, act.Completed),
		)

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT count(1) FROM "trajectory_points" WHERE action_id = $1`)).
		WithArgs(DefaultActionID).
		WillReturnRows(
			sqlmock.NewRows([]string{"count(1)"}).
				AddRow(0),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`UPDATE "actions" SET "communication_id"=$1,"fleet_id"=$2,"is_completed"=$3 WHERE "id" = $4`)).
		WithArgs(DefaultActionCommunicationID, DefaultActionFleetID, act.Completed, DefaultActionID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "trajectory_points" ("action_id","point_number","latitude","longitude","altitude","relative_altitude","speed","armed","flight_mode","orientation_x","orientation_y","orientation_z","orientation_w") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)`)).
		WithArgs(
			DefaultActionID,
			1,
			1.0,
			2.0,
			3.0,
			4.0,
			5.0,
			true,
			"state",
			6.0,
			7.0,
			8.0,
			9.0).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	repository := NewActionRepository()

	trajectoryPointComps := []act.TrajectoryPointComponent{
		&trajectoryPointComponentMock{
			1,
			1.0,
			2.0,
			3.0,
			4.0,
			5.0,
			true,
			"state",
			6.0,
			7.0,
			8.0,
			9.0,
		},
	}
	actionComp := actionComponentMock{
		ID:               string(DefaultActionID),
		CommunicationID:  string(DefaultActionCommunicationID),
		FleetID:          string(DefaultActionFleetID),
		IsCompleted:      act.Completed,
		TrajectoryPoints: trajectoryPointComps,
	}
	action := act.AssembleFrom(
		&actionComp,
	)

	err = repository.Save(db, action)

	a.Nil(err)
}

func TestActionRepositorySingleTrajectoryPointDataUpdateSave_ExistPreTrajectoryPoint(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "actions" WHERE id = $1`)).
		WithArgs(DefaultActionID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "communication_id", "fleet_id", "is_completed"}).
				AddRow(DefaultActionID, DefaultActionCommunicationID, DefaultActionFleetID, act.Completed),
		)

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT count(1) FROM "trajectory_points" WHERE action_id = $1`)).
		WithArgs(DefaultActionID).
		WillReturnRows(
			sqlmock.NewRows([]string{"count(1)"}).
				AddRow(1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`UPDATE "actions" SET "communication_id"=$1,"fleet_id"=$2,"is_completed"=$3 WHERE "id" = $4`)).
		WithArgs(DefaultActionCommunicationID, DefaultActionFleetID, act.Completed, DefaultActionID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	repository := NewActionRepository()

	trajectoryPointComps := []act.TrajectoryPointComponent{
		&trajectoryPointComponentMock{
			1,
			1.0,
			2.0,
			3.0,
			4.0,
			5.0,
			true,
			"state",
			6.0,
			7.0,
			8.0,
			9.0,
		},
	}
	actionComp := actionComponentMock{
		ID:               string(DefaultActionID),
		CommunicationID:  string(DefaultActionCommunicationID),
		FleetID:          string(DefaultActionFleetID),
		IsCompleted:      act.Completed,
		TrajectoryPoints: trajectoryPointComps,
	}
	action := act.AssembleFrom(
		&actionComp,
	)

	err = repository.Save(db, action)

	a.Nil(err)
}

func TestActionRepositoryMultipeTrajectoryPointDataUpdateSave_NonePreTrajectoryPoint(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "actions" WHERE id = $1`)).
		WithArgs(DefaultActionID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "communication_id", "fleet_id", "is_completed"}).
				AddRow(DefaultActionID, DefaultActionCommunicationID, DefaultActionFleetID, act.Completed),
		)

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT count(1) FROM "trajectory_points" WHERE action_id = $1`)).
		WithArgs(DefaultActionID).
		WillReturnRows(
			sqlmock.NewRows([]string{"count(1)"}).
				AddRow(0),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`UPDATE "actions" SET "communication_id"=$1,"fleet_id"=$2,"is_completed"=$3 WHERE "id" = $4`)).
		WithArgs(DefaultActionCommunicationID, DefaultActionFleetID, act.Completed, DefaultActionID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "trajectory_points" ("action_id","point_number","latitude","longitude","altitude","relative_altitude","speed","armed","flight_mode","orientation_x","orientation_y","orientation_z","orientation_w") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13),($14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26)`)).
		WithArgs(
			DefaultActionID,
			1,
			1.0,
			2.0,
			3.0,
			4.0,
			5.0,
			true,
			"state",
			6.0,
			7.0,
			8.0,
			9.0,
			DefaultActionID,
			2,
			1.0,
			2.0,
			3.0,
			4.0,
			5.0,
			true,
			"state",
			6.0,
			7.0,
			8.0,
			9.0).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	repository := NewActionRepository()

	trajectoryPointComps := []act.TrajectoryPointComponent{
		&trajectoryPointComponentMock{
			1,
			1.0,
			2.0,
			3.0,
			4.0,
			5.0,
			true,
			"state",
			6.0,
			7.0,
			8.0,
			9.0,
		},
		&trajectoryPointComponentMock{
			2,
			1.0,
			2.0,
			3.0,
			4.0,
			5.0,
			true,
			"state",
			6.0,
			7.0,
			8.0,
			9.0,
		},
	}
	actionComp := actionComponentMock{
		ID:               string(DefaultActionID),
		CommunicationID:  string(DefaultActionCommunicationID),
		FleetID:          string(DefaultActionFleetID),
		IsCompleted:      act.Completed,
		TrajectoryPoints: trajectoryPointComps,
	}
	action := act.AssembleFrom(
		&actionComp,
	)

	err = repository.Save(db, action)

	a.Nil(err)
}

func TestActionRepositoryMultipeTrajectoryPointDataUpdateSave_ExistPreTrajectoryPoint(t *testing.T) {
	a := assert.New(t)

	db, mock, err := postgresql.GetNewDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
		return
	}

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "actions" WHERE id = $1`)).
		WithArgs(DefaultActionID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "communication_id", "fleet_id", "is_completed"}).
				AddRow(DefaultActionID, DefaultActionCommunicationID, DefaultActionFleetID, act.Completed),
		)

	mock.
		ExpectQuery(
			regexp.QuoteMeta(`SELECT count(1) FROM "trajectory_points" WHERE action_id = $1`)).
		WithArgs(DefaultActionID).
		WillReturnRows(
			sqlmock.NewRows([]string{"count(1)"}).
				AddRow(1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`UPDATE "actions" SET "communication_id"=$1,"fleet_id"=$2,"is_completed"=$3 WHERE "id" = $4`)).
		WithArgs(DefaultActionCommunicationID, DefaultActionFleetID, act.Completed, DefaultActionID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.
		ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "trajectory_points" ("action_id","point_number","latitude","longitude","altitude","relative_altitude","speed","armed","flight_mode","orientation_x","orientation_y","orientation_z","orientation_w") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)`)).
		WithArgs(
			DefaultActionID,
			2,
			1.0,
			2.0,
			3.0,
			4.0,
			5.0,
			true,
			"state",
			6.0,
			7.0,
			8.0,
			9.0).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	repository := NewActionRepository()

	trajectoryPointComps := []act.TrajectoryPointComponent{
		&trajectoryPointComponentMock{
			1,
			1.0,
			2.0,
			3.0,
			4.0,
			5.0,
			true,
			"state",
			6.0,
			7.0,
			8.0,
			9.0,
		},
		&trajectoryPointComponentMock{
			2,
			1.0,
			2.0,
			3.0,
			4.0,
			5.0,
			true,
			"state",
			6.0,
			7.0,
			8.0,
			9.0,
		},
	}
	actionComp := actionComponentMock{
		ID:               string(DefaultActionID),
		CommunicationID:  string(DefaultActionCommunicationID),
		FleetID:          string(DefaultActionFleetID),
		IsCompleted:      act.Completed,
		TrajectoryPoints: trajectoryPointComps,
	}
	action := act.AssembleFrom(
		&actionComp,
	)

	err = repository.Save(db, action)

	a.Nil(err)
}
