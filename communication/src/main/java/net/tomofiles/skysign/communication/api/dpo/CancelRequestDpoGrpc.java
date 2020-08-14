package net.tomofiles.skysign.communication.api.dpo;

import lombok.RequiredArgsConstructor;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.usecase.dpo.CancelRequestDpo;
import proto.skysign.CancelRequest;

@RequiredArgsConstructor
public class CancelRequestDpoGrpc implements CancelRequestDpo {

    private final CancelRequest request;

    @Override
    public CommunicationId getCommId() {
        return new CommunicationId(this.request.getId());
    }
}