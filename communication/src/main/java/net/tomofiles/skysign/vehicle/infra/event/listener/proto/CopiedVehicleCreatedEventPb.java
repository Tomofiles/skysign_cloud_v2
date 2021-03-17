package net.tomofiles.skysign.vehicle.infra.event.listener.proto;

import org.springframework.amqp.core.Message;
import org.springframework.amqp.core.MessageProperties;

import net.tomofiles.skysign.vehicle.domain.vehicle.CopiedVehicleCreatedEvent;

public class CopiedVehicleCreatedEventPb {

    private final proto.skysign.event.CopiedVehicleCreatedEvent event;

    public CopiedVehicleCreatedEventPb(CopiedVehicleCreatedEvent event) {
        this.event = proto.skysign.event.CopiedVehicleCreatedEvent.newBuilder()
            .setVehicleId(event.getVehicleId().getId())
            .setCommunicationId(event.getCommunicationId().getId())
            .setFlightplanId(event.getFlightplanId().getId())
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