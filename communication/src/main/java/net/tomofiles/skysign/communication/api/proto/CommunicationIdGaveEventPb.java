package net.tomofiles.skysign.communication.api.proto;

import com.google.protobuf.InvalidProtocolBufferException;

import net.tomofiles.skysign.communication.api.event.event.CommunicationIdGaveEvent;


public class CommunicationIdGaveEventPb {

    private final proto.skysign.event.CommunicationIdGaveEvent event;

    public CommunicationIdGaveEventPb(byte[] message) throws InvalidProtocolBufferException {
        this.event = proto.skysign.event.CommunicationIdGaveEvent.parseFrom(message);
    }

    @Override
    public String toString() {
        return this.event.toString().replaceAll("\\r\\n|\\r|\\n", " ");
    }

    public CommunicationIdGaveEvent getEvent() {
        return new CommunicationIdGaveEvent(
            this.event.getCommunicationId(),
            this.event.getVehicleId(),
            this.event.getVersion());
    }
}