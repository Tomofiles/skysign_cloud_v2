package net.tomofiles.skysign.communication.api.grpc;

import lombok.RequiredArgsConstructor;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.domain.communication.TelemetrySnapshot;
import net.tomofiles.skysign.communication.service.dpo.PushTelemetryRequestDpo;
import proto.skysign.PushTelemetryRequest;

@RequiredArgsConstructor
public class PushTelemetryRequestDpoGrpc implements PushTelemetryRequestDpo {

    private final PushTelemetryRequest request;

    @Override
    public CommunicationId getCommId() {
        return new CommunicationId(this.request.getId());
    }

    @Override
    public TelemetrySnapshot getTelemetry() {
        return new TelemetrySnapshot(
            request.getTelemetry().getLatitude(),
            request.getTelemetry().getLongitude(),
            request.getTelemetry().getAltitude(),
            request.getTelemetry().getRelativeAltitude(),
            request.getTelemetry().getSpeed(),
            request.getTelemetry().getArmed(),
            request.getTelemetry().getFlightMode(),
            request.getTelemetry().getOrientationX(),
            request.getTelemetry().getOrientationY(),
            request.getTelemetry().getOrientationZ(),
            request.getTelemetry().getOrientationW()
        );
    }
}