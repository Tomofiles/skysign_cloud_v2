package net.tomofiles.skysign.communication.service.dpo;

import net.tomofiles.skysign.communication.domain.communication.CommandId;
import net.tomofiles.skysign.communication.domain.communication.Communication;

public interface PushUploadMissionResponseDpo {
    public void setCommunication(Communication communication);
    public void setCommandId(CommandId commandId);
}