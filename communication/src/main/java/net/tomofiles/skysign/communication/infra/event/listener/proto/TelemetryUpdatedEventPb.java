package net.tomofiles.skysign.communication.infra.event.listener.proto;

import org.springframework.amqp.core.Message;
import org.springframework.amqp.core.MessageProperties;

import net.tomofiles.skysign.communication.domain.communication.TelemetryUpdatedEvent;

public class TelemetryUpdatedEventPb {

    private final proto.skysign.event.TelemetryUpdatedEvent event;

    public TelemetryUpdatedEventPb(TelemetryUpdatedEvent event) {
        this.event = proto.skysign.event.TelemetryUpdatedEvent.newBuilder()
            .setCommunicationId(event.getCommunicationId().getId())
            .setTelemetry(
                proto.skysign.common.Telemetry.newBuilder()
                    .setLatitude(event.getTelemetry().getLatitude())
                    .setLongitude(event.getTelemetry().getLongitude())
                    .setAltitude(event.getTelemetry().getAltitude())
                    .setRelativeAltitude(event.getTelemetry().getRelativeAltitude())
                    .setSpeed(event.getTelemetry().getSpeed())
                    .setArmed(event.getTelemetry().isArmed())
                    .setFlightMode(event.getTelemetry().getFlightMode())
                    .setOrientationX(event.getTelemetry().getX())
                    .setOrientationY(event.getTelemetry().getY())
                    .setOrientationZ(event.getTelemetry().getZ())
                    .setOrientationW(event.getTelemetry().getW())
                    .build()
            )
            .build();
    }

    @Override
    public String toString() {
        return this.event.toString().replaceAll("\\r\\n|\\r|\\n", " ");
    }

    public Message getMessage() {
        return new Message(this.event.toByteArray(), new MessageProperties());
    }
}