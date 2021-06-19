package net.tomofiles.skysign.mission.api.event;

import lombok.RequiredArgsConstructor;
import net.tomofiles.skysign.mission.api.event.event.MissionCopiedWhenFlightplanCopiedEvent;
import net.tomofiles.skysign.mission.domain.mission.MissionId;
import net.tomofiles.skysign.mission.service.dpo.CarbonCopyMissionRequestDpo;

@RequiredArgsConstructor
public class CarbonCopyMissionRequestDpoEvent implements CarbonCopyMissionRequestDpo {

    private final MissionCopiedWhenFlightplanCopiedEvent event;

    @Override
    public MissionId getOriginalId() {
        return new MissionId(event.getOriginalId());
    }

    @Override
    public MissionId getNewId() {
        return new MissionId(event.getNewId());
    }

}