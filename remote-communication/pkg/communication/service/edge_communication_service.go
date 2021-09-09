package service

import (
	c "remote-communication/pkg/communication/domain/communication"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/event"
	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/txmanager"
)

// EdgeCommunicationService .
type EdgeCommunicationService interface {
	PullCommand(command PullCommandCommand, pulledCommandType PulledCommandType) error
	PullUploadMission(command PullUploadMissionCommand, pulledMissionID PulledMissionID) error
	PushTelemetry(command PushTelemetryCommand, pulledCommandIDs PulledCommandIDs) error
}

// PullCommandCommand .
type PullCommandCommand interface {
	GetID() string
	GetCommandID() string
}

// PullUploadMissionCommand .
type PullUploadMissionCommand interface {
	GetID() string
	GetCommandID() string
}

// PushTelemetryCommand .
type PushTelemetryCommand interface {
	GetID() string
	GetTelemetry() EdgeTelemetry
}

// EdgeTelemetry .
type EdgeTelemetry interface {
	GetLatitude() float64
	GetLongitude() float64
	GetAltitude() float64
	GetRelativeAltitude() float64
	GetSpeed() float64
	GetArmed() bool
	GetFlightMode() string
	GetX() float64
	GetY() float64
	GetZ() float64
	GetW() float64
}

// PulledCommandType .
type PulledCommandType = func(cType string)

// PulledMissionID .
type PulledMissionID = func(missionID string)

// PulledCommandIDs .
type PulledCommandIDs = func(commandIDs []string)

// NewEdgeCommunicationService .
func NewEdgeCommunicationService(
	gen c.Generator,
	repo c.Repository,
	txm txmanager.TransactionManager,
	psm event.PubSubManager,
) EdgeCommunicationService {
	return &edgeCommunicationService{
		gen:  gen,
		repo: repo,
		txm:  txm,
		psm:  psm,
	}
}

type edgeCommunicationService struct {
	gen  c.Generator
	repo c.Repository
	txm  txmanager.TransactionManager
	psm  event.PubSubManager
}

func (s *edgeCommunicationService) PullCommand(
	command PullCommandCommand,
	pulledCommandType PulledCommandType,
) error {
	pub, chClose, err := s.psm.GetPublisher()
	if err != nil {
		return err
	}
	defer chClose()

	return s.txm.DoAndEndHook(
		func(tx txmanager.Tx) error {
			return s.pullCommandOperation(
				tx,
				pub,
				command,
				pulledCommandType,
			)
		},
		func() error {
			return pub.Flush()
		},
	)
}

func (s *edgeCommunicationService) pullCommandOperation(
	tx txmanager.Tx,
	pub event.Publisher,
	command PullCommandCommand,
	pulledCommandType PulledCommandType,
) error {
	cType, err := c.PullCommandService(
		tx,
		s.gen,
		s.repo,
		pub,
		c.ID(command.GetID()),
		c.CommandID(command.GetCommandID()),
	)
	if err != nil {
		return err
	}

	pulledCommandType(string(cType))
	return nil
}

func (s *edgeCommunicationService) PullUploadMission(
	command PullUploadMissionCommand,
	pulledMissionID PulledMissionID,
) error {
	pub, chClose, err := s.psm.GetPublisher()
	if err != nil {
		return err
	}
	defer chClose()

	return s.txm.DoAndEndHook(
		func(tx txmanager.Tx) error {
			return s.pullUploadMissionOperation(
				tx,
				pub,
				command,
				pulledMissionID,
			)
		},
		func() error {
			return pub.Flush()
		},
	)
}

func (s *edgeCommunicationService) pullUploadMissionOperation(
	tx txmanager.Tx,
	pub event.Publisher,
	command PullUploadMissionCommand,
	pulledMissionID PulledMissionID,
) error {
	missionID, err := c.PullUploadMissionService(
		tx,
		s.gen,
		s.repo,
		pub,
		c.ID(command.GetID()),
		c.CommandID(command.GetCommandID()),
	)
	if err != nil {
		return err
	}

	pulledMissionID(string(missionID))
	return nil
}

func (s *edgeCommunicationService) PushTelemetry(
	command PushTelemetryCommand,
	pulledCommandIDs PulledCommandIDs,
) error {
	pub, chClose, err := s.psm.GetPublisher()
	if err != nil {
		return err
	}
	defer chClose()

	return s.txm.DoAndEndHook(
		func(tx txmanager.Tx) error {
			return s.pushTelemetryOperation(
				tx,
				pub,
				command,
				pulledCommandIDs,
			)
		},
		func() error {
			return pub.Flush()
		},
	)
}

func (s *edgeCommunicationService) pushTelemetryOperation(
	tx txmanager.Tx,
	pub event.Publisher,
	command PushTelemetryCommand,
	pulledCommandIDs PulledCommandIDs,
) error {
	snapshot := c.TelemetrySnapshot{
		Latitude:         command.GetTelemetry().GetLatitude(),
		Longitude:        command.GetTelemetry().GetLongitude(),
		Altitude:         command.GetTelemetry().GetAltitude(),
		RelativeAltitude: command.GetTelemetry().GetRelativeAltitude(),
		Speed:            command.GetTelemetry().GetSpeed(),
		Armed:            command.GetTelemetry().GetArmed(),
		FlightMode:       command.GetTelemetry().GetFlightMode(),
		X:                command.GetTelemetry().GetX(),
		Y:                command.GetTelemetry().GetY(),
		Z:                command.GetTelemetry().GetZ(),
		W:                command.GetTelemetry().GetW(),
	}

	commandIDs, err := c.PushTelemetryService(
		tx,
		s.gen,
		s.repo,
		pub,
		c.ID(command.GetID()),
		snapshot,
	)
	if err != nil {
		return err
	}

	var strCommandIDs []string
	for _, commandID := range commandIDs {
		strCommandIDs = append(strCommandIDs, string(commandID))
	}

	pulledCommandIDs(strCommandIDs)
	return nil
}
