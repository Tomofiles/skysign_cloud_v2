package net.tomofiles.skysign.communication.usecase.dpo;

import net.tomofiles.skysign.communication.domain.communication.CommandId;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;

public interface PullCommandRequestDpo {
    public CommunicationId getCommId();
    public CommandId getCommandId();
}