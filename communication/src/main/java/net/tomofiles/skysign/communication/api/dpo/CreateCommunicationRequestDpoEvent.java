package net.tomofiles.skysign.communication.api.dpo;

import lombok.RequiredArgsConstructor;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.domain.vehicle.CommunicationIdChangedEvent;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleId;
import net.tomofiles.skysign.communication.service.dpo.CreateCommunicationRequestDpo;

@RequiredArgsConstructor
public class CreateCommunicationRequestDpoEvent implements CreateCommunicationRequestDpo {
    
    private final CommunicationIdChangedEvent event;

    @Override
    public CommunicationId getCommId() {
        return this.event.getAfterId();
    }

    @Override
    public VehicleId getVehicleId() {
        return this.event.getVehicleId();
    }
}