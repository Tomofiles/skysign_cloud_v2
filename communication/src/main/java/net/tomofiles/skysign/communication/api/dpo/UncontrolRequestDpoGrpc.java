package net.tomofiles.skysign.communication.api.dpo;

import lombok.RequiredArgsConstructor;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.service.dpo.UncontrolRequestDpo;
import proto.skysign.UncontrolRequest;

@RequiredArgsConstructor
public class UncontrolRequestDpoGrpc implements UncontrolRequestDpo {

    private final UncontrolRequest request;

    @Override
    public CommunicationId getCommId() {
        return new CommunicationId(this.request.getId());
    }
}