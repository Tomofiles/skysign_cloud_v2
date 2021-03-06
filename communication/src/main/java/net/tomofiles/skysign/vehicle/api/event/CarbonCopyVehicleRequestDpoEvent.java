package net.tomofiles.skysign.vehicle.api.event;

import lombok.RequiredArgsConstructor;
import net.tomofiles.skysign.vehicle.api.event.event.VehicleCopiedWhenFlightplanCopiedEvent;
import net.tomofiles.skysign.vehicle.domain.vehicle.VehicleId;
import net.tomofiles.skysign.vehicle.service.dpo.CarbonCopyVehicleRequestDpo;

@RequiredArgsConstructor
public class CarbonCopyVehicleRequestDpoEvent implements CarbonCopyVehicleRequestDpo {

    private final VehicleCopiedWhenFlightplanCopiedEvent event;

    @Override
    public VehicleId getOriginalId() {
        return new VehicleId(event.getOriginalId());
    }

    @Override
    public VehicleId getNewId() {
        return new VehicleId(event.getNewId());
    }

}