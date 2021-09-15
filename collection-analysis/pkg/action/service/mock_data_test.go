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
	LatitudeDegree:    1.0,
	LongitudeDegree:   2.0,
	AltitudeM:         3.0,
	RelativeAltitudeM: 4.0,
	SpeedMS:           5.0,
	Armed:             true,
	FlightMode:        "state",
	OrientationX:      6.0,
	OrientationY:      7.0,
	OrientationZ:      8.0,
	OrientationW:      9.0,
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
	PointNumber       int
	LatitudeDegree    float64
	LongitudeDegree   float64
	AltitudeM         float64
	RelativeAltitudeM float64
	SpeedMS           float64
	Armed             bool
	FlightMode        string
	OrientationX      float64
	OrientationY      float64
	OrientationZ      float64
	OrientationW      float64
}

func (c *trajectoryPointComponentMock) GetPointNumber() int {
	return c.PointNumber
}
func (c *trajectoryPointComponentMock) GetLatitudeDegree() float64 {
	return c.LatitudeDegree
}
func (c *trajectoryPointComponentMock) GetLongitudeDegree() float64 {
	return c.LongitudeDegree
}
func (c *trajectoryPointComponentMock) GetAltitudeM() float64 {
	return c.AltitudeM
}
func (c *trajectoryPointComponentMock) GetRelativeAltitudeM() float64 {
	return c.RelativeAltitudeM
}
func (c *trajectoryPointComponentMock) GetSpeedMS() float64 {
	return c.SpeedMS
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

type telemetryCommandMock struct {
	CommunicationID   string
	LatitudeDegree    float64
	LongitudeDegree   float64
	AltitudeM         float64
	RelativeAltitudeM float64
	SpeedMS           float64
	Armed             bool
	FlightMode        string
	OrientationX      float64
	OrientationY      float64
	OrientationZ      float64
	OrientationW      float64
}

func (r *telemetryCommandMock) GetCommunicationID() string {
	return r.CommunicationID
}
func (r *telemetryCommandMock) GetLatitudeDegree() float64 {
	return r.LatitudeDegree
}
func (r *telemetryCommandMock) GetLongitudeDegree() float64 {
	return r.LongitudeDegree
}
func (r *telemetryCommandMock) GetAltitudeM() float64 {
	return r.AltitudeM
}
func (r *telemetryCommandMock) GetRelativeAltitudeM() float64 {
	return r.RelativeAltitudeM
}
func (r *telemetryCommandMock) GetSpeedMS() float64 {
	return r.SpeedMS
}
func (r *telemetryCommandMock) GetArmed() bool {
	return r.Armed
}
func (r *telemetryCommandMock) GetFlightMode() string {
	return r.FlightMode
}
func (r *telemetryCommandMock) GetOrientationX() float64 {
	return r.OrientationX
}
func (r *telemetryCommandMock) GetOrientationY() float64 {
	return r.OrientationY
}
func (r *telemetryCommandMock) GetOrientationZ() float64 {
	return r.OrientationZ
}
func (r *telemetryCommandMock) GetOrientationW() float64 {
	return r.OrientationW
}
