package ports

import (
	"mission/pkg/mission/service"
	proto "mission/pkg/skysign_proto"
)

func MissionProtoTransfomerFromModel(
	model service.MissionPresentationModel,
) *proto.Mission {
	mission := &proto.Mission{}
	mission.Id = model.GetMission().GetID()
	mission.Name = model.GetMission().GetName()
	waypoints := []*proto.Waypoint{}
	for _, w := range model.GetMission().GetNavigation().GetWaypoints() {
		waypoints = append(
			waypoints,
			&proto.Waypoint{
				Latitude:       w.GetLatitude(),
				Longitude:      w.GetLongitude(),
				RelativeHeight: w.GetRelativeHeight(),
				Speed:          w.GetSpeed(),
			},
		)
	}
	mission.Navigation = &proto.Navigation{
		TakeoffPointGroundHeight: model.GetMission().GetNavigation().GetTakeoffPointGroundHeight(),
		Waypoints:                waypoints,
	}
	return mission
}
