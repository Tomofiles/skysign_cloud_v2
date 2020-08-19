package net.tomofiles.skysign.communication.api.dpo;

import lombok.RequiredArgsConstructor;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.service.dpo.PullTelemetryRequestDpo;
import proto.skysign.PullTelemetryRequest;

@RequiredArgsConstructor
public class PullTelemetryRequestDpoGrpc implements PullTelemetryRequestDpo {

    private final PullTelemetryRequest request;

    @Override
    public CommunicationId getCommId() {
        return new CommunicationId(this.request.getId());
    }
}