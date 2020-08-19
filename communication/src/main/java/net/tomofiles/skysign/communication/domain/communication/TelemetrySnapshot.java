package net.tomofiles.skysign.communication.domain.communication;

import lombok.AllArgsConstructor;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.ToString;

@AllArgsConstructor
@Getter
@EqualsAndHashCode
@ToString
public class TelemetrySnapshot {
    private double latitude;
    private double longitude;
    private double altitude;
    private double relativeAltitude;
    private double speed;
    private boolean armed;
    private String flightMode;
    private double x;
    private double y;
    private double z;
    private double w;
}