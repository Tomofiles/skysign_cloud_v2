package net.tomofiles.skysign.vehicle.infra.event.listener.proto;

import org.springframework.amqp.core.Message;
import org.springframework.amqp.core.MessageProperties;

import net.tomofiles.skysign.vehicle.domain.vehicle.CommunicationIdGaveEvent;

public class CommunicationIdGaveEventPb {

    private final proto.skysign.event.CommunicationIdGaveEvent event;

    public CommunicationIdGaveEventPb(CommunicationIdGaveEvent event) {
        this.event = proto.skysign.event.CommunicationIdGaveEvent.newBuilder()
            .setCommunicationId(event.getCommunicationId().getId())
            .setVehicleId(event.getVehicleId().getId())
            .setVersion(event.getVersion().getVersion())
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