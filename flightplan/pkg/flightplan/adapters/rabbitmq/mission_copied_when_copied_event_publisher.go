package rabbitmq

import (
	"flightplan/pkg/flightplan/domain/fleet"
	"flightplan/pkg/skysign_proto"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const missionCopiedWhenCopiedEventExchangeName = "flightplan.mission_copied_when_copied_event"

// PublishMissionCopiedWhenCopiedEvent .
func PublishMissionCopiedWhenCopiedEvent(
	ch Channel,
	event fleet.MissionCopiedWhenCopiedEvent,
) error {
	if err := ch.FanoutExchangeDeclare(
		missionCopiedWhenCopiedEventExchangeName,
	); err != nil {
		return err
	}

	eventPb := skysign_proto.MissionCopiedWhenCopiedEvent{
		OriginalMissionId: event.GetOriginalID(),
		NewMissionId:      event.GetNewID(),
	}
	eventBin, err := proto.Marshal(&eventPb)
	if err != nil {
		return err
	}

	if err := ch.Publish(
		missionCopiedWhenCopiedEventExchangeName,
		eventBin,
	); err != nil {
		return err
	}

	glog.Infof("PUBLISH , Event: %s, Message: %s", missionCopiedWhenCopiedEventExchangeName, eventPb.String())
	return nil
}
