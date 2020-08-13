package net.tomofiles.skysign.mission.api.dpo;

import lombok.RequiredArgsConstructor;
import net.tomofiles.skysign.mission.domain.mission.MissionId;
import net.tomofiles.skysign.mission.service.dpo.DeleteMissionRequestDpo;
import proto.skysign.DeleteMissionRequest;

@RequiredArgsConstructor
public class DeleteMissionRequestDpoGrpc implements DeleteMissionRequestDpo {

    private final DeleteMissionRequest request;

    @Override
    public MissionId getMissionId() {
        return new MissionId(request.getId());
    }
}