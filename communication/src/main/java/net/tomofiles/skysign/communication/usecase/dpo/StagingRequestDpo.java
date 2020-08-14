package net.tomofiles.skysign.communication.usecase.dpo;

import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.domain.communication.MissionId;

public interface StagingRequestDpo {
    public CommunicationId getCommId();
    public MissionId getMissionId();
}