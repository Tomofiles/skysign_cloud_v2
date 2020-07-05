package net.tomofiles.skysign.communication.domain.communication;

import lombok.AccessLevel;
import lombok.AllArgsConstructor;
import lombok.Getter;

@AllArgsConstructor(access = AccessLevel.PRIVATE)
@Getter
class Telemetry {
    private Position position;
    private double speed;
    private boolean armed;
    private String flightMode;
    private Orientation orientation;

    public static Telemetry newInstance() {
        return new Telemetry(
                new Position(
                        0.0,
                        0.0,
                        0.0,
                        0.0),
                0.0,
                false,
                "NONE",
                new Orientation(
                        0.0,
                        0.0,
                        0.0,
                        0.0)
        );
    }

    public Telemetry setPosition(
            double latitude,
            double longitude,
            double altitude,
            double relativeAltitude,
            double speed) {
        return new Telemetry(
                new Position(
                        latitude,
                        longitude,
                        altitude,
                        relativeAltitude),
                speed,
                this.armed,
                this.flightMode,
                this.orientation
        );
    }

    public Telemetry setArmed(boolean armed) {
        return new Telemetry(
                this.position,
                this.speed,
                armed,
                this.flightMode,
                this.orientation
        );
    }

    public Telemetry setFlightMode(String flightMode) {
        return new Telemetry(
                this.position,
                this.speed,
                this.armed,
                flightMode,
                this.orientation
        );
    }

    public Telemetry setOrientation(
            double orientationX,
            double orientationY,
            double orientationZ,
            double orientationW) {
        return new Telemetry(
                this.position,
                this.speed,
                this.armed,
                this.flightMode,
                new Orientation(
                        orientationX,
                        orientationY,
                        orientationZ,
                        orientationW)
        );
    }
}

@AllArgsConstructor
@Getter
class Position {
    private final double latitude;
    private final double longitude;
    private final double altitude;
    private final double relativeAltitude;
}

@AllArgsConstructor
@Getter
class Orientation {
    private final double x;
    private final double y;
    private final double z;
    private final double w;
}