package net.tomofiles.skysign.communication.api.dpo;

import lombok.RequiredArgsConstructor;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleId;
import net.tomofiles.skysign.communication.usecase.dpo.GetVehicleRequestDpo;
import proto.skysign.GetVehicleRequest;

@RequiredArgsConstructor
public class GetVehicleRequestDpoGrpc implements GetVehicleRequestDpo {

    private final GetVehicleRequest request;

    @Override
    public VehicleId getVehicleId() {
        return new VehicleId(request.getId());
    }
}