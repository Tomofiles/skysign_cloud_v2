package net.tomofiles.skysign.communication.domain.communication.component;

import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class TelemetryComponentDto {
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