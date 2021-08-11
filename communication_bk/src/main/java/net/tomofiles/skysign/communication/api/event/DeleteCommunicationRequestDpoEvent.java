package net.tomofiles.skysign.communication.api.event;

import lombok.RequiredArgsConstructor;
import net.tomofiles.skysign.communication.api.event.event.CommunicationIdRemovedEvent;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.service.dpo.DeleteCommunicationRequestDpo;

@RequiredArgsConstructor
public class DeleteCommunicationRequestDpoEvent implements DeleteCommunicationRequestDpo {
    
    private final CommunicationIdRemovedEvent event;

    @Override
    public CommunicationId getCommId() {
        return new CommunicationId(this.event.getCommunicationId());
    }
}