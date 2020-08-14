package net.tomofiles.skysign.communication.usecase.dpo;

import net.tomofiles.skysign.communication.domain.communication.CommandType;
import net.tomofiles.skysign.communication.domain.communication.Communication;

public interface PullCommandResponseDpo {
    public void setCommunication(Communication communication);
    public void setCommandType(CommandType commandType);
}