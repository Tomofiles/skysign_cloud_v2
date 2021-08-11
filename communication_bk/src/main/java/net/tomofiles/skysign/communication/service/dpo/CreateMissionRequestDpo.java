package net.tomofiles.skysign.communication.service.dpo;

import java.util.List;

import net.tomofiles.skysign.communication.domain.mission.MissionId;

public interface CreateMissionRequestDpo {
    public MissionId getMissionId();
    public List<WaypointDpo> getWaypoints();
}