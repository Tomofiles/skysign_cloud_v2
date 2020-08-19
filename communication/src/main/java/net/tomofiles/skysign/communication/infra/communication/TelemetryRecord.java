package net.tomofiles.skysign.communication.infra.communication;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class TelemetryRecord {
    private String commId;
    private double latitude;
    private double longitude;
    private double altitude;
    private double relativeAltitude;
    private double speed;
    private boolean armed;
    private String flightMode;
    private double oriX;
    private double oriY;
    private double oriZ;
    private double oriW;
}