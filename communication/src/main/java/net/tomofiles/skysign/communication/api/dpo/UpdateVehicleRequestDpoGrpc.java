package net.tomofiles.skysign.communication.api.dpo;

import lombok.RequiredArgsConstructor;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleId;
import net.tomofiles.skysign.communication.usecase.dpo.UpdateVehicleRequestDpo;
import proto.skysign.UpdateVehicleRequest;

@RequiredArgsConstructor
public class UpdateVehicleRequestDpoGrpc implements UpdateVehicleRequestDpo {

    private final UpdateVehicleRequest request;

    @Override
    public VehicleId getVehicleId() {
        return new VehicleId(request.getId());
    }

    @Override
    public String getVehicleName() {
        return this.request.getName();
    }

    @Override
    public CommunicationId getCommId() {
        return new CommunicationId(this.request.getCommId());
    }
}