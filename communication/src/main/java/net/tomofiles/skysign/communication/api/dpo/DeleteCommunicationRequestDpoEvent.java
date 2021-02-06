package net.tomofiles.skysign.communication.api.dpo;

import lombok.RequiredArgsConstructor;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.domain.vehicle.CommunicationIdRemovedEvent;
import net.tomofiles.skysign.communication.service.dpo.DeleteCommunicationRequestDpo;

@RequiredArgsConstructor
public class DeleteCommunicationRequestDpoEvent implements DeleteCommunicationRequestDpo {
    
    private final CommunicationIdRemovedEvent event;

    @Override
    public CommunicationId getCommId() {
        return this.event.getCommunicationId();
    }
}