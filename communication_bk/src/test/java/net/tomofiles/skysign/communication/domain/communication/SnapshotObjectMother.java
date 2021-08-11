package net.tomofiles.skysign.communication.domain.communication;

public class SnapshotObjectMother {
    
    /**
     * テスト用Telemetryスナップショットを生成する。
     */
    public static TelemetrySnapshot newNormalTelemetrySnapshot() {
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

        return new TelemetrySnapshot(
                latitude,
                longitude,
                altitude,
                relativeAltitude,
                speed,
                armed,
                flightMode,
                orientationX,
                orientationY,
                orientationZ,
                orientationW
        );
    }
}