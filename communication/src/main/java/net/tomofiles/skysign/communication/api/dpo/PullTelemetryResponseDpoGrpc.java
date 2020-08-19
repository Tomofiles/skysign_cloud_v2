package net.tomofiles.skysign.communication.api.dpo;

import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.domain.communication.TelemetrySnapshot;
import net.tomofiles.skysign.communication.service.dpo.PullTelemetryResponseDpo;

public class PullTelemetryResponseDpoGrpc implements PullTelemetryResponseDpo {

    private CommunicationId communicationId = null;
    private TelemetrySnapshot telemetry = null;

    @Override
    public void setTelemetry(CommunicationId communicationId, TelemetrySnapshot telemetry) {
        this.communicationId = communicationId;
        this.telemetry = telemetry;
    }

    public boolean isEmpty() {
        return this.communicationId == null
                || this.telemetry == null;
    }

    public proto.skysign.PullTelemetryResponse getGrpcResponse() {
        return proto.skysign.PullTelemetryResponse.newBuilder()
                .setId(communicationId.getId())
                .setTelemetry(proto.skysign.common.Telemetry.newBuilder().setLatitude(telemetry.getLatitude())
                        .setLongitude(telemetry.getLongitude()).setAltitude(telemetry.getAltitude())
                        .setRelativeAltitude(telemetry.getRelativeAltitude()).setSpeed(telemetry.getSpeed())
                        .setArmed(telemetry.isArmed()).setFlightMode(telemetry.getFlightMode())
                        .setOrientationX(telemetry.getX()).setOrientationY(telemetry.getY())
                        .setOrientationZ(telemetry.getZ()).setOrientationW(telemetry.getW()))
                .build();
    }
}