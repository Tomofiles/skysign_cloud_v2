package rabbitmq

import (
	"github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/fleet/domain/fleet"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	crm "github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/adapters/rabbitmq"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const missionCopiedEventExchangeName = "fleet.mission_copied_event"

// PublishMissionCopiedEvent .
func PublishMissionCopiedEvent(
	ch crm.Channel,
	event fleet.MissionCopiedEvent,
) error {
	if err := ch.FanoutExchangeDeclare(
		missionCopiedEventExchangeName,
	); err != nil {
		return err
	}

	eventPb := skysign_proto.MissionCopiedEvent{
		FleetId:           event.GetFleetID(),
		OriginalMissionId: event.GetOriginalID(),
		NewMissionId:      event.GetNewID(),
	}
	eventBin, err := proto.Marshal(&eventPb)
	if err != nil {
		return err
	}

	if err := ch.Publish(
		missionCopiedEventExchangeName,
		eventBin,
	); err != nil {
		return err
	}

	glog.Infof("PUBLISH , Event: %s, Message: %s", missionCopiedEventExchangeName, eventPb.String())
	return nil
}
