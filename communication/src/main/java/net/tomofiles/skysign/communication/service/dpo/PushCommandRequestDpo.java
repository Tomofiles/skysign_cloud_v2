package net.tomofiles.skysign.communication.service.dpo;

import net.tomofiles.skysign.communication.domain.communication.CommandType;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;

public interface PushCommandRequestDpo {
    public CommunicationId getCommId();
    public CommandType getCommandType();
}