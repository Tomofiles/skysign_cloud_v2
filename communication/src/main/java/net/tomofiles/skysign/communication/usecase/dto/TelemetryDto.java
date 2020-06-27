package net.tomofiles.skysign.communication.usecase.dto;

import lombok.Data;

@Data
public class TelemetryDto {
    private double latitude;
    private double longitude;
    private double altitude;
    private double speed;
    private boolean armed;
    private String flightMode;
    private double[] orientation;
}