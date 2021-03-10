package postgresql

import (
	fope "flightreport/pkg/flightreport/domain/flightoperation"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DefaultFlightoperationID = fope.ID("flightoperation-id")
const DefaultFlightplanID = fope.FlightplanID("flightplan-id")

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
}

func (f *flightoperationComponentMock) GetID() string {
	return f.id
}

func (f *flightoperationComponentMock) GetFlightplanID() string {
	return f.flightplanID
}
