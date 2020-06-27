package net.tomofiles.skysign.communication.domain.communication;

import lombok.AccessLevel;
import lombok.AllArgsConstructor;
import lombok.Getter;

@AllArgsConstructor(access = AccessLevel.PACKAGE)
@Getter
public class TelemetrySnapshot {
    private double latitude;
    private double longitude;
    private double altitude;
    private double speed;
    private boolean armed;
    private String flightMode;
    private double x;
    private double y;
    private double z;
    private double w;
}