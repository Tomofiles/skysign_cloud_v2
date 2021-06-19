package net.tomofiles.skysign.mission.service.dpo;

import net.tomofiles.skysign.mission.domain.mission.MissionId;
import net.tomofiles.skysign.mission.domain.mission.Navigation;

public interface UpdateMissionRequestDpo {
    public MissionId getMissionId();
    public String getMissionName();
    public Navigation getNavigation();
}