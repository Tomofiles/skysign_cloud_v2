package net.tomofiles.skysign.communication.service.dpo;

import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.domain.communication.TelemetrySnapshot;

public interface PushTelemetryRequestDpo {
    public CommunicationId getCommId();
    public TelemetrySnapshot getTelemetry();
}