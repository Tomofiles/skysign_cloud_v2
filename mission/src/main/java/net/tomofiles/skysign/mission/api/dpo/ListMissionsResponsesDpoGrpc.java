package net.tomofiles.skysign.mission.api.dpo;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

import net.tomofiles.skysign.mission.domain.mission.Mission;
import net.tomofiles.skysign.mission.service.dpo.ListMissionsResponsesDpo;

public class ListMissionsResponsesDpoGrpc implements ListMissionsResponsesDpo {

    private List<Mission> missions;

    public ListMissionsResponsesDpoGrpc() {
        this.missions = new ArrayList<>();
    }

    @Override
    public void setMissions(List<Mission> missions) {
        this.missions = missions;
    }

    public proto.skysign.ListMissionsResponses getGrpcResponse() {
        List<proto.skysign.Mission> r = this.missions.stream().map(mission -> {
            return proto.skysign.Mission.newBuilder()
                    .setId(mission.getId().getId())
                    .setName(mission.getMissionName())
                    .setTakeoffPointGroundHeight(mission.getNavigation().getTakeoffPointGroundHeight().getHeightM())
                    .addAllItems(mission.getNavigation().getWaypoints().stream().map(waypoint -> {
                        return proto.skysign.MissionItem.newBuilder()
                                .setLatitude(waypoint.getLatitude())
                                .setLongitude(waypoint.getLongitude())
                                .setRelativeHeight(waypoint.getRelativeHeightM())
                                .setSpeed(waypoint.getSpeedMS())
                                .build();
                        }).collect(Collectors.toList()))
                    .build();
        }).collect(Collectors.toList());

        return proto.skysign.ListMissionsResponses.newBuilder()
                .addAllMissions(r)
                .build();
    }
}