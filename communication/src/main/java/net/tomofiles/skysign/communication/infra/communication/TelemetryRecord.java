package net.tomofiles.skysign.communication.infra.communication;

import lombok.Data;

@Data
public class TelemetryRecord {
    private String commId;
    private double latitude;
    private double longitude;
    private double altitude;
    private double speed;
    private boolean armed;
    private String flightMode;
    private double oriX;
    private double oriY;
    private double oriZ;
    private double oriW;
}