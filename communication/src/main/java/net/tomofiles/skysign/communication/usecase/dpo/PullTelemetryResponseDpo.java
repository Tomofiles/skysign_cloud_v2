package net.tomofiles.skysign.communication.usecase.dpo;

import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.domain.communication.TelemetrySnapshot;

public interface PullTelemetryResponseDpo {
    public void setTelemetry(CommunicationId communicationId, TelemetrySnapshot telemetrySnapshot);
}