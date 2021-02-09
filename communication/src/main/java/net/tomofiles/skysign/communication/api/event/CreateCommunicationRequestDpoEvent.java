package net.tomofiles.skysign.communication.api.event;

import lombok.RequiredArgsConstructor;
import net.tomofiles.skysign.communication.api.event.event.CommunicationIdGaveEvent;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.domain.communication.VehicleId;
import net.tomofiles.skysign.communication.service.dpo.CreateCommunicationRequestDpo;

@RequiredArgsConstructor
public class CreateCommunicationRequestDpoEvent implements CreateCommunicationRequestDpo {
    
    private final CommunicationIdGaveEvent event;

    @Override
    public CommunicationId getCommId() {
        return new CommunicationId(this.event.getCommunicationId());
    }

    @Override
    public VehicleId getVehicleId() {
        return new VehicleId(this.event.getVehicleId());
    }
}