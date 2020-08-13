package net.tomofiles.skysign.mission.api.dpo;

import lombok.RequiredArgsConstructor;
import net.tomofiles.skysign.mission.domain.mission.MissionId;
import net.tomofiles.skysign.mission.service.dpo.GetMissionRequestDpo;
import proto.skysign.GetMissionRequest;

@RequiredArgsConstructor
public class GetMissionRequestDpoGrpc implements GetMissionRequestDpo {

    private final GetMissionRequest request;

    @Override
    public MissionId getMissionId() {
        return new MissionId(request.getId());
    }
}