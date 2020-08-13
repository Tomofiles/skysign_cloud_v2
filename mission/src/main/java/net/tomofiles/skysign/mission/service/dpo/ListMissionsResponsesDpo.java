package net.tomofiles.skysign.mission.service.dpo;

import java.util.List;

import net.tomofiles.skysign.mission.domain.mission.Mission;

public interface ListMissionsResponsesDpo {
    public void setMissions(List<Mission> missions);
}