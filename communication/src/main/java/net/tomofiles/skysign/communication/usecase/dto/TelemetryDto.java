package net.tomofiles.skysign.communication.usecase.dto;

import lombok.Data;

@Data
public class TelemetryDto {
    private String name;
    private double latitude;
    private double longitude;
    private double altitude;
    private double relativeAltitude;
    private double speed;
    private boolean armed;
    private String flightMode;
    private double orientationX;
    private double orientationY;
    private double orientationZ;
    private double orientationW;
}