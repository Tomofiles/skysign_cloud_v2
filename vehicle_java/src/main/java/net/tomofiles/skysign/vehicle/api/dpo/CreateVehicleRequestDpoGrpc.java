package net.tomofiles.skysign.vehicle.api.dpo;

import lombok.RequiredArgsConstructor;
import net.tomofiles.skysign.vehicle.domain.vehicle.CommunicationId;
import net.tomofiles.skysign.vehicle.service.dpo.CreateVehicleRequestDpo;
import proto.skysign.common.Vehicle;

@RequiredArgsConstructor
public class CreateVehicleRequestDpoGrpc implements CreateVehicleRequestDpo {

    private final Vehicle request;

    @Override
    public String getVehicleName() {
        return this.request.getName();
    }

    @Override
    public CommunicationId getCommunicationId() {
        return new CommunicationId(this.request.getCommunicationId());
    }
}