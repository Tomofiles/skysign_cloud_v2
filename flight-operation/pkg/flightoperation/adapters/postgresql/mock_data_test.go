package postgresql

import (
	fope "flight-operation/pkg/flightoperation/domain/flightoperation"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DefaultID = fope.ID("flightoperation-id")
const DefaultName = "flightoperation-name"
const DefaultDescription = "flightoperation-description"
const DefaultFleetID = fope.FleetID("fleet-id")
const DefaultIsCompleted = fope.Completed
const DefaultVersion = fope.Version("version")

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

// Flightoperation構成オブジェクトモック
type flightoperationComponentMock struct {
	id          string
	name        string
	description string
	fleetID     string
	isCompleted bool
	version     string
}

func (f *flightoperationComponentMock) GetID() string {
	return f.id
}

// GetName .
func (f *flightoperationComponentMock) GetName() string {
	return f.name
}

// GetDescription .
func (f *flightoperationComponentMock) GetDescription() string {
	return f.description
}

// GetFleetID .
func (f *flightoperationComponentMock) GetFleetID() string {
	return f.fleetID
}

func (f *flightoperationComponentMock) GetIsCompleted() bool {
	return f.isCompleted
}

func (f *flightoperationComponentMock) GetVersion() string {
	return f.version
}
