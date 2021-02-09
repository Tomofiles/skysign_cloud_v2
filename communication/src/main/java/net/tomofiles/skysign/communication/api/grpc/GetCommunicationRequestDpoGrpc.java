package net.tomofiles.skysign.communication.api.grpc;

import lombok.RequiredArgsConstructor;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.service.dpo.GetCommunicationRequestDpo;
import proto.skysign.GetCommunicationRequest;

@RequiredArgsConstructor
public class GetCommunicationRequestDpoGrpc implements GetCommunicationRequestDpo {

    private final GetCommunicationRequest request;

    @Override
    public CommunicationId getCommId() {
        return new CommunicationId(this.request.getId());
    }
}