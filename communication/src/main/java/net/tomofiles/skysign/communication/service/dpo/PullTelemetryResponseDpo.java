package net.tomofiles.skysign.communication.service.dpo;

import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.domain.communication.TelemetrySnapshot;

public interface PullTelemetryResponseDpo {
    public void setTelemetry(CommunicationId communicationId, TelemetrySnapshot telemetrySnapshot);
}