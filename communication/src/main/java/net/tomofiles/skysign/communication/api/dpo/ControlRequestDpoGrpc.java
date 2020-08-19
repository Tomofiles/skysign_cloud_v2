package net.tomofiles.skysign.communication.api.dpo;

import lombok.RequiredArgsConstructor;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.service.dpo.ControlRequestDpo;
import proto.skysign.ControlRequest;

@RequiredArgsConstructor
public class ControlRequestDpoGrpc implements ControlRequestDpo {

    private final ControlRequest request;

    @Override
    public CommunicationId getCommId() {
        return new CommunicationId(this.request.getId());
    }
}