package net.tomofiles.skysign.vehicle.api.dpo;

import lombok.RequiredArgsConstructor;
import net.tomofiles.skysign.vehicle.domain.vehicle.VehicleId;
import net.tomofiles.skysign.vehicle.service.dpo.DeleteVehicleRequestDpo;
import proto.skysign.DeleteVehicleRequest;

@RequiredArgsConstructor
public class DeleteVehicleRequestDpoGrpc implements DeleteVehicleRequestDpo {

    private final DeleteVehicleRequest request;

    @Override
    public VehicleId getVehicleId() {
        return new VehicleId(request.getId());
    }
}