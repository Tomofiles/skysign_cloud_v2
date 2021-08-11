package net.tomofiles.skysign.communication.api.grpc;

import java.util.stream.Collectors;

import net.tomofiles.skysign.communication.domain.mission.Mission;
import net.tomofiles.skysign.communication.service.dpo.GetUploadMissionResponseDpo;

public class GetUploadMissionResponseDpoGrpc implements GetUploadMissionResponseDpo {

    private Mission mission = null;

    @Override
    public void setMission(Mission mission) {
        this.mission = mission;
    }

    public boolean isEmpty() {
        return this.mission == null;
    }

    public proto.skysign.UploadMission getGrpcResponse() {
        return proto.skysign.UploadMission.newBuilder()
                .setId(mission.getId().getId())
                .addAllWaypoints(mission.getWaypoints().stream().map(waypoint -> {
                    return proto.skysign.common.Waypoint.newBuilder()
                        .setLatitude(waypoint.getLatitude())
                        .setLongitude(waypoint.getLongitude())
                        .setRelativeHeight(waypoint.getRelativeHeightM())
                        .setSpeed(waypoint.getSpeedMS())
                        .build();
                }).collect(Collectors.toList()))
                .build();
    }
}