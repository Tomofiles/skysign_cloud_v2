package net.tomofiles.skysign.communication.api.event;

import lombok.RequiredArgsConstructor;
import net.tomofiles.skysign.communication.api.event.event.MissionDeletedEvent;
import net.tomofiles.skysign.communication.domain.mission.MissionId;
import net.tomofiles.skysign.communication.service.dpo.DeleteMissionRequestDpo;

@RequiredArgsConstructor
public class DeleteMissionRequestDpoEvent implements DeleteMissionRequestDpo {
    
    private final MissionDeletedEvent event;

    @Override
    public MissionId getMissionId() {
        return new MissionId(this.event.getMissionId());
    }
}