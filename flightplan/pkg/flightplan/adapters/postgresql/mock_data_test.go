package postgresql

import (
	"flightplan/pkg/flightplan/domain/flightplan"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DefaultID = flightplan.ID("flightplan-id")
const DefaultVersion = flightplan.Version("version")
const DefaultName = "flightplan-name"
const DefaultDescription = "flightplan-description"

func GetNewDbMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, mock, err
	}

	gormDB, err := gorm.Open(
		postgres.New(
			postgres.Config{
				Conn: db,
			}), &gorm.Config{})

	if err != nil {
		return gormDB, mock, err
	}

	return gormDB, mock, err
}

// Flightplan構成オブジェクトモック
type flightplanComponentMock struct {
	id          string
	name        string
	description string
	version     string
}

func (f *flightplanComponentMock) GetID() string {
	return f.id
}

func (f *flightplanComponentMock) GetName() string {
	return f.name
}

func (f *flightplanComponentMock) GetDescription() string {
	return f.description
}

func (f *flightplanComponentMock) GetVersion() string {
	return f.version
}
