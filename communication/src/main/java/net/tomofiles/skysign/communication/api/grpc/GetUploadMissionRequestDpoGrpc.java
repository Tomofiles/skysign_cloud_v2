package net.tomofiles.skysign.communication.api.grpc;

import lombok.RequiredArgsConstructor;
import net.tomofiles.skysign.communication.domain.mission.MissionId;
import net.tomofiles.skysign.communication.service.dpo.GetUploadMissionRequestDpo;
import proto.skysign.GetUploadMissionRequest;

@RequiredArgsConstructor
public class GetUploadMissionRequestDpoGrpc implements GetUploadMissionRequestDpo {
    
    private final GetUploadMissionRequest request;

    @Override
    public MissionId getMissionId() {
        return new MissionId(request.getId());
    }
}