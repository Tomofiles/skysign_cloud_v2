package net.tomofiles.skysign.communication.service.dpo;

import java.util.List;

import net.tomofiles.skysign.communication.domain.communication.CommandId;
import net.tomofiles.skysign.communication.domain.communication.Communication;

public interface PushTelemetryResponseDpo {
    public void setCommunication(Communication communication);
    public void setCommandIds(List<CommandId> commandIds);
}