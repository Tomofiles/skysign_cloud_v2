package postgresql

import (
	"fleet-formation/pkg/fleet/domain/fleet"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DefaultFleetID = fleet.ID("fleet-id")
const DefaultFleetAssignmentID = fleet.AssignmentID("assignment-id")
const DefaultFleetEventID = fleet.EventID("event-id")
const DefaultFleetVehicleID = fleet.VehicleID("vehicle-id")
const DefaultFleetMissionID = fleet.MissionID("mission-id")
const DefaultFleetVersion = fleet.Version("version")

func GetNewDbMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gormDB, err := gorm.Open(
		postgres.New(
			postgres.Config{
				Conn: db,
			}), &gorm.Config{})

	if err != nil {
		return nil, nil, err
	}

	return gormDB, mock, nil
}

// Fleet構成オブジェクトモック
type fleetComponentMock struct {
	id           string
	isCarbonCopy bool
	assignments  []assignmentComponentMock
	events       []eventComponentMock
	version      string
}

func (f *fleetComponentMock) GetID() string {
	return f.id
}

func (f *fleetComponentMock) GetIsCarbonCopy() bool {
	return f.isCarbonCopy
}

func (f *fleetComponentMock) GetVersion() string {
	return f.version
}

func (f *fleetComponentMock) GetAssignments() []fleet.AssignmentComponent {
	var assignments []fleet.AssignmentComponent
	for _, a := range f.assignments {
		assignments = append(
			assignments,
			&assignmentComponentMock{
				id:        a.id,
				fleetID:   a.fleetID,
				vehicleID: a.vehicleID,
			},
		)
	}
	return assignments
}

func (f *fleetComponentMock) GetEvents() []fleet.EventComponent {
	var events []fleet.EventComponent
	for _, e := range f.events {
		events = append(
			events,
			&eventComponentMock{
				id:           e.id,
				fleetID:      e.fleetID,
				assignmentID: e.assignmentID,
				missionID:    e.missionID,
			},
		)
	}
	return events
}

// Assignment構成オブジェクトモック
type assignmentComponentMock struct {
	id        string
	fleetID   string
	vehicleID string
}

func (a *assignmentComponentMock) GetID() string {
	return a.id
}

func (a *assignmentComponentMock) GetFleetID() string {
	return a.fleetID
}

func (a *assignmentComponentMock) GetVehicleID() string {
	return a.vehicleID
}

// Event構成オブジェクトモック
type eventComponentMock struct {
	id           string
	fleetID      string
	assignmentID string
	missionID    string
}

func (e *eventComponentMock) GetID() string {
	return e.id
}

func (e *eventComponentMock) GetFleetID() string {
	return e.fleetID
}

func (e *eventComponentMock) GetAssignmentID() string {
	return e.assignmentID
}

func (e *eventComponentMock) GetMissionID() string {
	return e.missionID
}
