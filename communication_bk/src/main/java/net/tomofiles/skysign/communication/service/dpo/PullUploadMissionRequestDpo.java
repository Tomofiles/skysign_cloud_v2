package net.tomofiles.skysign.communication.service.dpo;

import net.tomofiles.skysign.communication.domain.communication.CommandId;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;

public interface PullUploadMissionRequestDpo {
    public CommunicationId getCommId();
    public CommandId getCommandId();
}