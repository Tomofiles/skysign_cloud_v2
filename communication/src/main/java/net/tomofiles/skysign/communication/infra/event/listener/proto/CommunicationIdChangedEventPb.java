package net.tomofiles.skysign.communication.infra.event.listener.proto;

import org.springframework.amqp.core.Message;
import org.springframework.amqp.core.MessageProperties;

import net.tomofiles.skysign.communication.domain.vehicle.CommunicationIdChangedEvent;

public class CommunicationIdChangedEventPb {

    private final proto.skysign.event.CommunicationIdChangedEvent event;

    public CommunicationIdChangedEventPb(CommunicationIdChangedEvent event) {
        if (event.isFirst()) {
            this.event = proto.skysign.event.CommunicationIdChangedEvent.newBuilder()
                .setAfterId(event.getAfterId().getId())
                .setVehicleId(event.getVehicleId().getId())
                .setVersion(event.getVersion().getVersion())
                .build();
        } else {
            this.event = proto.skysign.event.CommunicationIdChangedEvent.newBuilder()
                .setBeforeId(event.getBeforeId().getId())
                .setAfterId(event.getAfterId().getId())
                .setVehicleId(event.getVehicleId().getId())
                .setVersion(event.getVersion().getVersion())
                .build();
        }
    }

    @Override
    public String toString() {
        return this.event.toString().replaceAll("\\r\\n|\\r|\\n", " ");
    }

    public Message getMessage() {
        return new Message(this.event.toByteArray(), new MessageProperties());
    }
}