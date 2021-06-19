package net.tomofiles.skysign.mission.service.dpo;

import net.tomofiles.skysign.mission.domain.mission.MissionId;

public interface CarbonCopyMissionRequestDpo {
    public MissionId getOriginalId();
    public MissionId getNewId();
}