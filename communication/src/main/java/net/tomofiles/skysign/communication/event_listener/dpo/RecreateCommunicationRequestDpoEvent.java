package net.tomofiles.skysign.communication.event_listener.dpo;

import lombok.RequiredArgsConstructor;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.domain.vehicle.CommunicationIdChangedEvent;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleId;
import net.tomofiles.skysign.communication.usecase.dpo.RecreateCommunicationRequestDpo;

@RequiredArgsConstructor
public class RecreateCommunicationRequestDpoEvent implements RecreateCommunicationRequestDpo {

    private final CommunicationIdChangedEvent event;

    @Override
    public CommunicationId getBeforeCommId() {
        return this.event.getBeforeId();
    }

    @Override
    public CommunicationId getAfterCommId() {
        return this.event.getAfterId();
    }

    @Override
    public VehicleId getVehicleId() {
        return this.event.getVehicleId();
    }
}