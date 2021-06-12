package net.tomofiles.skysign.vehicle.infra.vehicle;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class VehicleRecord {
    private String id;
    private String name;
    private String communicationId;
    private boolean isCarbonCopy;
    private String version;
    private String newVersion;
}