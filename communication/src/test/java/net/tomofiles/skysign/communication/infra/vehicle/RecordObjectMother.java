package net.tomofiles.skysign.communication.infra.vehicle;

import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.domain.vehicle.Generator;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleId;
import net.tomofiles.skysign.communication.domain.vehicle.Version;

public class RecordObjectMother {
    
    /**
     * 通常のVehicleレコードを生成する。
     */
    public static VehicleRecord newNormalVehicleRecord(VehicleId vehicleId, Version version, Generator generator) {
        return new VehicleRecord(
                vehicleId.getId(),
                "vehicle name",
                new CommunicationId("comm id").getId(),
                version.getVersion(),
                version.getVersion());
    }

}