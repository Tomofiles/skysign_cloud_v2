package service

import (
	act "collection-analysis/pkg/action/domain/action"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/txmanager"

	"github.com/stretchr/testify/mock"
)

const DefaultActionID = act.ID("action-id")
const DefaultActionCommunicationID = act.CommunicationID("communication-id")
const DefaultActionFleetID = act.FleetID("fleet-id")

var DefaultTelemetrySnapshot = act.TelemetrySnapshot{
	Latitude:         1.0,
	Longitude:        2.0,
	Altitude:         3.0,
	RelativeAltitude: 4.0,
	Speed:            5.0,
	Armed:            true,
	FlightMode:       "state",
	OrientationX:     6.0,
	OrientationY:     7.0,
	OrientationZ:     8.0,
	OrientationW:     9.0,
}

type actionRepositoryMock struct {
	mock.Mock
	action *act.Action
}

func (r *actionRepositoryMock) GetByID(
	tx txmanager.Tx,
	id act.ID,
) (*act.Action, error) {
	ret := r.Called(id)
	var a *act.Action
	if ret.Get(0) == nil {
		a = nil
	} else {
		a = ret.Get(0).(*act.Action)
	}
	return a, ret.Error(1)
}

func (r *actionRepositoryMock) GetAllActiveByFleetID(
	tx txmanager.Tx,
	fleet act.FleetID,
) ([]*act.Action, error) {
	ret := r.Called(fleet)
	var a []*act.Action
	if ret.Get(0) == nil {
		a = nil
	} else {
		a = ret.Get(0).([]*act.Action)
	}
	return a, ret.Error(1)
}

func (r *actionRepositoryMock) GetActiveByCommunicationID(
	tx txmanager.Tx,
	communicationID act.CommunicationID,
) (*act.Action, error) {
	ret := r.Called(communicationID)
	var a *act.Action
	if ret.Get(0) == nil {
		a = nil
	} else {
		a = ret.Get(0).(*act.Action)
	}
	return a, ret.Error(1)
}

func (r *actionRepositoryMock) Save(
	tx txmanager.Tx,
	action *act.Action,
) error {
	ret := r.Called(action)
	r.action = action
	return ret.Error(0)
}

type txManagerMock struct {
	isOpe, isEH error
}

func (txm *txManagerMock) Do(operation func(txmanager.Tx) error) error {
	txm.isOpe = operation(nil)
	return nil
}

func (txm *txManagerMock) DoAndEndHook(operation func(txmanager.Tx) error, endHook func() error) error {
	txm.isOpe = operation(nil)
	txm.isEH = endHook()
	return nil
}

type actionComponentMock struct {
	ID               string
	CommunicationID  string
	FleetID          string
	IsCompleted      bool
	TrajectoryPoints []act.TrajectoryPointComponent
}

func (c *actionComponentMock) GetID() string {
	return c.ID
}
func (c *actionComponentMock) GetCommunicationID() string {
	return c.CommunicationID
}
func (c *actionComponentMock) GetFleetID() string {
	return c.FleetID
}
func (c *actionComponentMock) GetIsCompleted() bool {
	return c.IsCompleted
}
func (c *actionComponentMock) GetTrajectory() []act.TrajectoryPointComponent {
	return c.TrajectoryPoints
}

type trajectoryPointComponentMock struct {
	PointNumber      int
	Latitude         float64
	Longitude        float64
	Altitude         float64
	RelativeAltitude float64
	Speed            float64
	Armed            bool
	FlightMode       string
	OrientationX     float64
	OrientationY     float64
	OrientationZ     float64
	OrientationW     float64
}

func (c *trajectoryPointComponentMock) GetPointNumber() int {
	return c.PointNumber
}
func (c *trajectoryPointComponentMock) GetLatitude() float64 {
	return c.Latitude
}
func (c *trajectoryPointComponentMock) GetLongitude() float64 {
	return c.Longitude
}
func (c *trajectoryPointComponentMock) GetAltitude() float64 {
	return c.Altitude
}
func (c *trajectoryPointComponentMock) GetRelativeAltitude() float64 {
	return c.RelativeAltitude
}
func (c *trajectoryPointComponentMock) GetSpeed() float64 {
	return c.Speed
}
func (c *trajectoryPointComponentMock) GetArmed() bool {
	return c.Armed
}
func (c *trajectoryPointComponentMock) GetFlightMode() string {
	return c.FlightMode
}
func (c *trajectoryPointComponentMock) GetOrientationX() float64 {
	return c.OrientationX
}
func (c *trajectoryPointComponentMock) GetOrientationY() float64 {
	return c.OrientationY
}
func (c *trajectoryPointComponentMock) GetOrientationZ() float64 {
	return c.OrientationZ
}
func (c *trajectoryPointComponentMock) GetOrientationW() float64 {
	return c.OrientationW
}

type createCommandMock struct {
	VehicleId, CommunicationId, FleetId string
}

func (r *createCommandMock) GetID() string {
	return r.VehicleId
}

func (r *createCommandMock) GetCommunicationID() string {
	return r.CommunicationId
}

func (r *createCommandMock) GetFleetID() string {
	return r.FleetId
}

type actionIDCommandMock struct {
	ID string
}

func (r *actionIDCommandMock) GetID() string {
	return r.ID
}

type fleetIDCommandMock struct {
	FleetID string
}

func (r *fleetIDCommandMock) GetFleetID() string {
	return r.FleetID
}

type telemetryRequestMock struct {
	CommunicationID  string
	Latitude         float64
	Longitude        float64
	Altitude         float64
	RelativeAltitude float64
	Speed            float64
	Armed            bool
	FlightMode       string
	OrientationX     float64
	OrientationY     float64
	OrientationZ     float64
	OrientationW     float64
}

func (r *telemetryRequestMock) GetCommunicationID() string {
	return r.CommunicationID
}
func (r *telemetryRequestMock) GetLatitude() float64 {
	return r.Latitude
}
func (r *telemetryRequestMock) GetLongitude() float64 {
	return r.Longitude
}
func (r *telemetryRequestMock) GetAltitude() float64 {
	return r.Altitude
}
func (r *telemetryRequestMock) GetRelativeAltitude() float64 {
	return r.RelativeAltitude
}
func (r *telemetryRequestMock) GetSpeed() float64 {
	return r.Speed
}
func (r *telemetryRequestMock) GetArmed() bool {
	return r.Armed
}
func (r *telemetryRequestMock) GetFlightMode() string {
	return r.FlightMode
}
func (r *telemetryRequestMock) GetOrientationX() float64 {
	return r.OrientationX
}
func (r *telemetryRequestMock) GetOrientationY() float64 {
	return r.OrientationY
}
func (r *telemetryRequestMock) GetOrientationZ() float64 {
	return r.OrientationZ
}
func (r *telemetryRequestMock) GetOrientationW() float64 {
	return r.OrientationW
}
