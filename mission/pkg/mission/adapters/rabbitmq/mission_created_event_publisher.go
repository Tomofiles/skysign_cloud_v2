package rabbitmq

import (
	"mission/pkg/mission/domain/mission"
	"mission/pkg/skysign_proto"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const missionCreatedEventExchangeName = "mission.mission_created_event"

// PublishMissionCreatedEvent .
func PublishMissionCreatedEvent(
	ch Channel,
	event mission.CreatedEvent,
) error {
	if err := ch.FanoutExchangeDeclare(
		missionCreatedEventExchangeName,
	); err != nil {
		return err
	}

	var navigation *skysign_proto.Navigation
	if event.GetMission().GetNavigation() != nil {
		var waypoints []*skysign_proto.Waypoint
		event.GetMission().GetNavigation().ProvideWaypointsInterest(
			func(pointOrder int, latitudeDegree, longitudeDegree, relativeHeightM, speedMS float64) {
				waypoints = append(
					waypoints,
					&skysign_proto.Waypoint{
						Latitude:       latitudeDegree,
						Longitude:      longitudeDegree,
						RelativeHeight: relativeHeightM,
						Speed:          speedMS,
					},
				)
			},
		)
		navigation = &skysign_proto.Navigation{
			TakeoffPointGroundHeight: event.GetMission().GetNavigation().GetTakeoffPointGroundHeightWGS84EllipsoidM(),
			Waypoints:                waypoints,
			UploadId:                 string(event.GetMission().GetNavigation().GetUploadID()),
		}
	}
	eventPb := skysign_proto.MissionCreatedEvent{
		MissionId: event.GetMissionID(),
		Mission: &skysign_proto.Mission{
			Id:         string(event.GetMission().GetID()),
			Name:       event.GetMission().GetName(),
			Navigation: navigation,
		},
	}
	eventBin, err := proto.Marshal(&eventPb)
	if err != nil {
		return err
	}

	if err := ch.Publish(
		missionCreatedEventExchangeName,
		eventBin,
	); err != nil {
		return err
	}

	glog.Infof("PUBLISH , Event: %s, Message: %s", missionCreatedEventExchangeName, eventPb.String())
	return nil
}
