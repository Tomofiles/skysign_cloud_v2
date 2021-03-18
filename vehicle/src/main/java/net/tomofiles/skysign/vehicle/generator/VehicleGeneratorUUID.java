package net.tomofiles.skysign.vehicle.generator;

import java.util.UUID;

import org.springframework.stereotype.Component;

import net.tomofiles.skysign.vehicle.domain.vehicle.Generator;
import net.tomofiles.skysign.vehicle.domain.vehicle.VehicleId;
import net.tomofiles.skysign.vehicle.domain.vehicle.Version;

@Component
public class VehicleGeneratorUUID implements Generator {
    @Override
    public VehicleId newVehicleId() {
        return new VehicleId(UUID.randomUUID().toString());
    }

    @Override
    public Version newVersion() {
        return new Version(UUID.randomUUID().toString());
    }
}