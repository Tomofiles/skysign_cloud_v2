package postgresql

import (
	fope "flightoperation/pkg/flightoperation/domain/flightoperation"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DefaultFlightoperationID = fope.ID("flightoperation-id")
const DefaultFlightplanID = fope.FlightplanID("flightplan-id")
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
	id           string
	flightplanID string
	isCompleted  bool
	version      string
}

func (f *flightoperationComponentMock) GetID() string {
	return f.id
}

func (f *flightoperationComponentMock) GetFlightplanID() string {
	return f.flightplanID
}

func (f *flightoperationComponentMock) GetIsCompleted() bool {
	return f.isCompleted
}

func (f *flightoperationComponentMock) GetVersion() string {
	return f.version
}
