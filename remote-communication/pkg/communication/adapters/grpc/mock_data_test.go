package grpc

import (
	"remote-communication/pkg/communication/service"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

const DefaultCommunicationID = "communication-id"
const DefaultCommunicationCommandID = "command-id"
const DefaultCommunicationMissionID = "mission-id"

type userCommunicationServiceMock struct {
	mock.Mock
	ID, Type, MissionID string
}

func (s *userCommunicationServiceMock) PushCommand(
	command service.PushCommandCommand,
	pushedCommandID service.PushedCommandID,
) error {
	ret := s.Called()
	s.ID = command.GetID()
	s.Type = command.GetType()
	var f string
	if ret.Get(0) != nil {
		f = ret.Get(0).(string)
	}
	pushedCommandID(f)
	return ret.Error(1)
}

func (s *userCommunicationServiceMock) PushUploadMission(
	command service.PushUploadMissionCommand,
	pushedCommandID service.PushedCommandID,
) error {
	ret := s.Called()
	s.ID = command.GetID()
	s.MissionID = command.GetMissionID()
	var f string
	if ret.Get(0) != nil {
		f = ret.Get(0).(string)
	}
	pushedCommandID(f)
	return ret.Error(1)
}

func (s *userCommunicationServiceMock) PullTelemetry(
	command service.PullTelemetryCommand,
	pulledTelemetry service.PulledTelemetry,
) error {
	ret := s.Called()
	s.ID = command.GetID()
	var f service.UserTelemetry
	if ret.Get(0) != nil {
		f = ret.Get(0).(service.UserTelemetry)
	}
	pulledTelemetry(f)
	return ret.Error(1)
}

type edgeCommunicationServiceMock struct {
	mock.Mock
	ID, CommandID string
	Telemetry     service.EdgeTelemetry
}

func (s *edgeCommunicationServiceMock) PullCommand(
	command service.PullCommandCommand,
	pulledCommandType service.PulledCommandType,
) error {
	ret := s.Called()
	s.ID = command.GetID()
	s.CommandID = command.GetCommandID()
	var f string
	if ret.Get(0) != nil {
		f = ret.Get(0).(string)
	}
	pulledCommandType(f)
	return ret.Error(1)
}

func (s *edgeCommunicationServiceMock) PullUploadMission(
	command service.PullUploadMissionCommand,
	pulledMissionID service.PulledMissionID,
) error {
	ret := s.Called()
	s.ID = command.GetID()
	s.CommandID = command.GetCommandID()
	var f string
	if ret.Get(0) != nil {
		f = ret.Get(0).(string)
	}
	pulledMissionID(f)
	return ret.Error(1)
}

func (s *edgeCommunicationServiceMock) PushTelemetry(
	command service.PushTelemetryCommand,
	pulledCommandIDs service.PulledCommandIDs,
) error {
	ret := s.Called()
	s.ID = command.GetID()
	s.Telemetry = command.GetTelemetry()
	var f []string
	if ret.Get(0) != nil {
		f = ret.Get(0).([]string)
	}
	pulledCommandIDs(f)
	return ret.Error(1)
}

type serviceRegistrarMock struct {
	descs []*grpc.ServiceDesc
	impls []interface{}
}

func (s *serviceRegistrarMock) RegisterService(desc *grpc.ServiceDesc, impl interface{}) {
	s.descs = append(s.descs, desc)
	s.impls = append(s.impls, impl)
}
