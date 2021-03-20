package rabbitmq

import (
	"flightplan/pkg/flightplan/domain/fleet"
	"flightplan/pkg/skysign_proto"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const missionCopiedWhenFlightplanCopiedEventExchangeName = "fleet.mission_copied_when_flightplan_copied_event"

// PublishMissionCopiedWhenFlightplanCopiedEvent .
func PublishMissionCopiedWhenFlightplanCopiedEvent(
	ch Channel,
	event fleet.MissionCopiedWhenFlightplanCopiedEvent,
) error {
	if err := ch.FanoutExchangeDeclare(
		missionCopiedWhenFlightplanCopiedEventExchangeName,
	); err != nil {
		return err
	}

	eventPb := skysign_proto.MissionCopiedWhenFlightplanCopiedEvent{
		FlightplanId:      event.GetFlightplanID(),
		OriginalMissionId: event.GetOriginalID(),
		NewMissionId:      event.GetNewID(),
	}
	eventBin, err := proto.Marshal(&eventPb)
	if err != nil {
		return err
	}

	if err := ch.Publish(
		missionCopiedWhenFlightplanCopiedEventExchangeName,
		eventBin,
	); err != nil {
		return err
	}

	glog.Infof("PUBLISH , Event: %s, Message: %s", missionCopiedWhenFlightplanCopiedEventExchangeName, eventPb.String())
	return nil
}
