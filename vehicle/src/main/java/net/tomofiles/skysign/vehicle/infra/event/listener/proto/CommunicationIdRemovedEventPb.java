package net.tomofiles.skysign.vehicle.infra.event.listener.proto;

import org.springframework.amqp.core.Message;
import org.springframework.amqp.core.MessageProperties;

import net.tomofiles.skysign.vehicle.domain.vehicle.CommunicationIdRemovedEvent;

public class CommunicationIdRemovedEventPb {

    private final proto.skysign.event.CommunicationIdRemovedEvent event;

    public CommunicationIdRemovedEventPb(CommunicationIdRemovedEvent event) {
        this.event = proto.skysign.event.CommunicationIdRemovedEvent.newBuilder()
            .setCommunicationId(event.getCommunicationId().getId())
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