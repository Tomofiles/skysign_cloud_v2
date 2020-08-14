package net.tomofiles.skysign.communication.service.dpo;

import net.tomofiles.skysign.communication.domain.communication.CommunicationId;

public interface PullTelemetryRequestDpo {
    public CommunicationId getCommId();
}