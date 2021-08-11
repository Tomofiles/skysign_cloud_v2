package net.tomofiles.skysign.communication.service.dpo;

import net.tomofiles.skysign.communication.domain.communication.Communication;
import net.tomofiles.skysign.communication.domain.communication.MissionId;

public interface PullUploadMissionResponseDpo {
    public void setCommunication(Communication communication);
    public void setMissionId(MissionId missionId);
}