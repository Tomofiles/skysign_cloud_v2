package postgresql

import (
	frep "flightreport/pkg/flightreport/domain/flightreport"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DefaultFlightreportID = frep.ID("flightreport-id")
const DefaultFlightoperationID = frep.FlightoperationID("flightoperation-id")

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

// Flightreport構成オブジェクトモック
type flightreportComponentMock struct {
	id                string
	flightoperationID string
}

func (f *flightreportComponentMock) GetID() string {
	return f.id
}

func (f *flightreportComponentMock) GetFlightoperationID() string {
	return f.flightoperationID
}
