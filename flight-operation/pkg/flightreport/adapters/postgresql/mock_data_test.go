package postgresql

import (
	frep "flight-operation/pkg/flightreport/domain/flightreport"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DefaultID = frep.ID("flightreport-id")
const DefaultName = "flightreport-name"
const DefaultDescription = "flightreport-description"
const DefaultFleetID = frep.FleetID("fleet-id")

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
	id          string
	name        string
	description string
	fleetID     string
}

func (f *flightreportComponentMock) GetID() string {
	return f.id
}

func (f *flightreportComponentMock) GetName() string {
	return f.name
}

func (f *flightreportComponentMock) GetDescription() string {
	return f.description
}

func (f *flightreportComponentMock) GetFleetID() string {
	return f.fleetID
}
