package grpc

import (
	"fleet-formation/pkg/mission/service"

	proto "github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
)

// MissionProtoTransformerFromModel .
func MissionProtoTransformerFromModel(
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
				Latitude:         w.GetLatitudeDegree(),
				Longitude:        w.GetLongitudeDegree(),
				RelativeAltitude: w.GetRelativeAltitudeM(),
				Speed:            w.GetSpeedMS(),
			},
		)
	}
	mission.Navigation = &proto.Navigation{
		TakeoffPointGroundAltitude: model.GetMission().GetNavigation().GetTakeoffPointGroundAltitudeM(),
		Waypoints:                  waypoints,
		UploadId:                   model.GetMission().GetNavigation().GetUploadID(),
	}
	return mission
}
