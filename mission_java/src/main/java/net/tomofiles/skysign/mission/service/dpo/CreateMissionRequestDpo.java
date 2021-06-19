package net.tomofiles.skysign.mission.service.dpo;

import net.tomofiles.skysign.mission.domain.mission.Navigation;

public interface CreateMissionRequestDpo {
    public String getMissionName();
    public Navigation getNavigation();
}