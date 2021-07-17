package postgresql

import (
	v "vehicle/pkg/vehicle/domain/vehicle"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DefaultVehicleID = v.ID("vehicle-id")
const DefaultVehicleVersion = v.Version("version")
const DefaultVehicleName = "vehicle-name"
const DefaultVehicleCommunicationID = v.CommunicationID("communication-id")
const DefaultFleetID = v.FleetID("fleet-id")

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

type vehicleComponentMock struct {
	ID              string
	Name            string
	CommunicationID string
	IsCarbonCopy    bool
	Version         string
}

func (v *vehicleComponentMock) GetID() string {
	return v.ID
}

func (v *vehicleComponentMock) GetName() string {
	return v.Name
}

func (v *vehicleComponentMock) GetCommunicationID() string {
	return v.CommunicationID
}

func (v *vehicleComponentMock) GetIsCarbonCopy() bool {
	return v.IsCarbonCopy
}

func (v *vehicleComponentMock) GetVersion() string {
	return v.Version
}
