package net.tomofiles.skysign.vehicle.api.dpo;

import lombok.RequiredArgsConstructor;
import net.tomofiles.skysign.vehicle.domain.vehicle.VehicleId;
import net.tomofiles.skysign.vehicle.service.dpo.GetVehicleRequestDpo;
import proto.skysign.GetVehicleRequest;

@RequiredArgsConstructor
public class GetVehicleRequestDpoGrpc implements GetVehicleRequestDpo {

    private final GetVehicleRequest request;

    @Override
    public VehicleId getId() {
        return new VehicleId(request.getId());
    }
}