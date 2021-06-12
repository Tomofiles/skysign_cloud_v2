package net.tomofiles.skysign.vehicle.api.dpo;

import lombok.RequiredArgsConstructor;
import net.tomofiles.skysign.vehicle.domain.vehicle.CommunicationId;
import net.tomofiles.skysign.vehicle.domain.vehicle.VehicleId;
import net.tomofiles.skysign.vehicle.service.dpo.UpdateVehicleRequestDpo;
import proto.skysign.common.Vehicle;

@RequiredArgsConstructor
public class UpdateVehicleRequestDpoGrpc implements UpdateVehicleRequestDpo {

    private final Vehicle request;

    @Override
    public VehicleId getId() {
        return new VehicleId(request.getId());
    }

    @Override
    public String getVehicleName() {
        return this.request.getName();
    }

    @Override
    public CommunicationId getCommunicationId() {
        return new CommunicationId(this.request.getCommunicationId());
    }
}