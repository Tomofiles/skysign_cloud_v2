package net.tomofiles.skysign.vehicle.infra.vehicle;

import net.tomofiles.skysign.vehicle.domain.vehicle.CommunicationId;
import net.tomofiles.skysign.vehicle.domain.vehicle.Generator;
import net.tomofiles.skysign.vehicle.domain.vehicle.VehicleId;
import net.tomofiles.skysign.vehicle.domain.vehicle.Version;

public class RecordObjectMother {
    
    /**
     * 通常のVehicleレコードを生成する。
     */
    public static VehicleRecord newNormalVehicleRecord(VehicleId vehicleId, Version version, Generator generator) {
        return new VehicleRecord(
                vehicleId.getId(),
                "vehicle name",
                new CommunicationId("comm id").getId(),
                false,
                version.getVersion(),
                version.getVersion());
    }

    /**
     * カーボンコピーされたVehicleレコードを生成する。
     */
    public static VehicleRecord newCarbonCopiedVehicleRecord(VehicleId vehicleId, Version version, Generator generator) {
        return new VehicleRecord(
                vehicleId.getId(),
                "vehicle name",
                new CommunicationId("comm id").getId(),
                true,
                version.getVersion(),
                version.getVersion());
    }

}