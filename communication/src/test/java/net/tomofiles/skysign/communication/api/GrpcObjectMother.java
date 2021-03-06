package net.tomofiles.skysign.communication.api;

import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import proto.skysign.PullTelemetryResponse;
import proto.skysign.common.Telemetry;

public class GrpcObjectMother {

    /**
     * テスト用Telemetryオブジェクトを生成する。
     */
    public static PullTelemetryResponse newNormalPullTelemetryResponseGrpc(CommunicationId communicationId) {
        return PullTelemetryResponse.newBuilder()
                .setId(communicationId.getId())
                .setTelemetry(newNormalTelemetryGrpc())
                .build();
    }

    /**
     * テスト用Telemetryオブジェクトを生成する。
     */
    public static Telemetry newNormalTelemetryGrpc() {
        double latitude = 0.0d;
        double longitude = 1.0d;
        double altitude = 2.0d;
        double relativeAltitude = 3.0d;
        double speed = 4.0d;
        boolean armed = true;
        String flightMode = "INFLIGHT";
        double orientationX = 5.0d;
        double orientationY = 6.0d;
        double orientationZ = 7.0d;
        double orientationW = 8.0d;

        return Telemetry.newBuilder()
                .setLatitude(latitude)
                .setLongitude(longitude)
                .setAltitude(altitude)
                .setRelativeAltitude(relativeAltitude)
                .setSpeed(speed)
                .setArmed(armed)
                .setFlightMode(flightMode)
                .setOrientationX(orientationX)
                .setOrientationY(orientationY)
                .setOrientationZ(orientationZ)
                .setOrientationW(orientationW)
                .build();
    }
}