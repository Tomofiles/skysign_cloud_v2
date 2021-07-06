package rabbitmq

import (
	"mission/pkg/mission/domain/mission"
	"mission/pkg/skysign_proto"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const missionDeletedEventExchangeName = "mission.mission_deleted_event"

// PublishMissionDeletedEvent .
func PublishMissionDeletedEvent(
	ch Channel,
	event mission.DeletedEvent,
) error {
	if err := ch.FanoutExchangeDeclare(
		missionDeletedEventExchangeName,
	); err != nil {
		return err
	}

	eventPb := skysign_proto.MissionDeletedEvent{
		MissionId:       event.GetMissionID(),
		UploadMissionId: event.GetUploadID(),
	}
	eventBin, err := proto.Marshal(&eventPb)
	if err != nil {
		return err
	}

	if err := ch.Publish(
		missionDeletedEventExchangeName,
		eventBin,
	); err != nil {
		return err
	}

	glog.Infof("PUBLISH , Event: %s, Message: %s", missionDeletedEventExchangeName, eventPb.String())
	return nil
}
