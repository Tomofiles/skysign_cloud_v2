package net.tomofiles.skysign.mission.api.dpo;

import java.util.stream.Collectors;

import net.tomofiles.skysign.mission.domain.mission.Mission;
import net.tomofiles.skysign.mission.service.dpo.UpdateMissionResponseDpo;

public class UpdateMissionResponseDpoGrpc implements UpdateMissionResponseDpo {

    private Mission mission = null;

    @Override
    public void setMission(Mission mission) {
        this.mission = mission;
    }

    public boolean isEmpty() {
        return mission == null;
    }

    public proto.skysign.common.Mission getGrpcResponse() {
        return proto.skysign.common.Mission.newBuilder()
                .setId(mission.getId().getId())
                .setName(mission.getMissionName())
                .setTakeoffPointGroundHeight(mission.getNavigation().getTakeoffPointGroundHeight().getHeightM())
                .addAllItems(mission.getNavigation().getWaypoints().stream().map(waypoint -> {
                    return proto.skysign.common.MissionItem.newBuilder()
                            .setLatitude(waypoint.getLatitude())
                            .setLongitude(waypoint.getLongitude())
                            .setRelativeHeight(waypoint.getRelativeHeightM())
                            .setSpeed(waypoint.getSpeedMS())
                            .build();
                    }).collect(Collectors.toList()))
                .build();
    }
}