package net.tomofiles.skysign.communication.infra.vehicle;

import lombok.Data;

@Data
public class VehicleRecord {
    private String id;
    private String name;
    private String commId;
    private String version;
    private String newVersion;
}