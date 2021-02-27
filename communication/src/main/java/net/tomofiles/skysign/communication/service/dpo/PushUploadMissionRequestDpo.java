package net.tomofiles.skysign.communication.service.dpo;

import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.domain.communication.MissionId;

public interface PushUploadMissionRequestDpo {
    public CommunicationId getCommId();
    public MissionId getMissionId();
}