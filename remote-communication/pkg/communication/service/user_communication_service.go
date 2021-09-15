package service

import (
	c "remote-communication/pkg/communication/domain/communication"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/event"
	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/txmanager"
)

// UserCommunicationService .
type UserCommunicationService interface {
	PushCommand(command PushCommandCommand, pushedCommandID PushedCommandID) error
	PushUploadMission(command PushUploadMissionCommand, pushedCommandID PushedCommandID) error
	PullTelemetry(command PullTelemetryCommand, pulledTelemetry PulledTelemetry) error
}

// PushCommandCommand .
type PushCommandCommand interface {
	GetID() string
	GetType() string
}

// PushUploadMissionCommand .
type PushUploadMissionCommand interface {
	GetID() string
	GetMissionID() string
}

// PullTelemetryCommand .
type PullTelemetryCommand interface {
	GetID() string
}

// UserTelemetry .
type UserTelemetry interface {
	GetLatitudeDegree() float64
	GetLongitudeDegree() float64
	GetAltitudeM() float64
	GetRelativeAltitudeM() float64
	GetSpeedMS() float64
	GetArmed() bool
	GetFlightMode() string
	GetX() float64
	GetY() float64
	GetZ() float64
	GetW() float64
}

// PushedCommandID .
type PushedCommandID = func(commandID string)

// PulledTelemetry .
type PulledTelemetry = func(telemetry UserTelemetry)

// NewUserCommunicationService .
func NewUserCommunicationService(
	gen c.Generator,
	repo c.Repository,
	txm txmanager.TransactionManager,
	psm event.PubSubManager,
) UserCommunicationService {
	return &userCommunicationService{
		gen:  gen,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}
}

type userCommunicationService struct {
	gen  c.Generator
	repo c.Repository
	txm  txmanager.TransactionManager
	psm  event.PubSubManager
}

func (s *userCommunicationService) PushCommand(
	command PushCommandCommand,
	pushedCommandID PushedCommandID,
) error {
	pub, chClose, err := s.psm.GetPublisher()
	if err != nil {
		return err
	}
	defer chClose()

	return s.txm.DoAndEndHook(
		func(tx txmanager.Tx) error {
			return s.pushCommandOperation(
				tx,
				pub,
				command,
				pushedCommandID,
			)
		},
		func() error {
			return pub.Flush()
		},
	)
}

func (s *userCommunicationService) pushCommandOperation(
	tx txmanager.Tx,
	pub event.Publisher,
	command PushCommandCommand,
	pushedCommandID PushedCommandID,
) error {
	commandID, err := c.PushCommandService(
		tx,
		s.gen,
		s.repo,
		pub,
		c.ID(command.GetID()),
		c.CommandType(command.GetType()),
	)
	if err != nil {
		return err
	}

	pushedCommandID(string(commandID))
	return nil
}

func (s *userCommunicationService) PushUploadMission(
	command PushUploadMissionCommand,
	pushedCommandID PushedCommandID,
) error {
	pub, chClose, err := s.psm.GetPublisher()
	if err != nil {
		return err
	}
	defer chClose()

	return s.txm.DoAndEndHook(
		func(tx txmanager.Tx) error {
			return s.pushUploadMissionOperation(
				tx,
				pub,
				command,
				pushedCommandID,
			)
		},
		func() error {
			return pub.Flush()
		},
	)
}

func (s *userCommunicationService) pushUploadMissionOperation(
	tx txmanager.Tx,
	pub event.Publisher,
	command PushUploadMissionCommand,
	pushedCommandID PushedCommandID,
) error {
	commandID, err := c.PushUploadMissionService(
		tx,
		s.gen,
		s.repo,
		pub,
		c.ID(command.GetID()),
		c.MissionID(command.GetMissionID()),
	)
	if err != nil {
		return err
	}

	pushedCommandID(string(commandID))
	return nil
}

func (s *userCommunicationService) PullTelemetry(
	command PullTelemetryCommand,
	pulledTelemetry PulledTelemetry,
) error {
	pub, chClose, err := s.psm.GetPublisher()
	if err != nil {
		return err
	}
	defer chClose()

	return s.txm.DoAndEndHook(
		func(tx txmanager.Tx) error {
			return s.pullTelemetryOperation(
				tx,
				pub,
				command,
				pulledTelemetry,
			)
		},
		func() error {
			return pub.Flush()
		},
	)
}

func (s *userCommunicationService) pullTelemetryOperation(
	tx txmanager.Tx,
	pub event.Publisher,
	command PullTelemetryCommand,
	pulledTelemetry PulledTelemetry,
) error {
	snapshot, err := c.PullTelemetryService(
		tx,
		s.gen,
		s.repo,
		pub,
		c.ID(command.GetID()),
	)
	if err != nil {
		return err
	}

	telemetry := &telemetry{
		latitudeDegree:    snapshot.LatitudeDegree,
		longitudeDegree:   snapshot.LongitudeDegree,
		altitudeM:         snapshot.AltitudeM,
		relativeAltitudeM: snapshot.RelativeAltitudeM,
		speedMS:           snapshot.SpeedMS,
		armed:             snapshot.Armed,
		flightMode:        snapshot.FlightMode,
		x:                 snapshot.X,
		y:                 snapshot.Y,
		z:                 snapshot.Z,
		w:                 snapshot.W,
	}

	pulledTelemetry(telemetry)
	return nil
}

type telemetry struct {
	latitudeDegree    float64
	longitudeDegree   float64
	altitudeM         float64
	relativeAltitudeM float64
	speedMS           float64
	armed             bool
	flightMode        string
	x                 float64
	y                 float64
	z                 float64
	w                 float64
}

func (t *telemetry) GetLatitudeDegree() float64 {
	return t.latitudeDegree
}

func (t *telemetry) GetLongitudeDegree() float64 {
	return t.longitudeDegree
}

func (t *telemetry) GetAltitudeM() float64 {
	return t.altitudeM
}

func (t *telemetry) GetRelativeAltitudeM() float64 {
	return t.relativeAltitudeM
}

func (t *telemetry) GetSpeedMS() float64 {
	return t.speedMS
}

func (t *telemetry) GetArmed() bool {
	return t.armed
}

func (t *telemetry) GetFlightMode() string {
	return t.flightMode
}

func (t *telemetry) GetX() float64 {
	return t.x
}

func (t *telemetry) GetY() float64 {
	return t.y
}

func (t *telemetry) GetZ() float64 {
	return t.z
}

func (t *telemetry) GetW() float64 {
	return t.w
}
