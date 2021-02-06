package net.tomofiles.skysign.communication.api.proto;

import com.google.protobuf.InvalidProtocolBufferException;

import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.domain.vehicle.CommunicationIdRemovedEvent;
import net.tomofiles.skysign.communication.domain.vehicle.Version;

public class CommunicationIdRemovedEventPb {

    private final proto.skysign.event.CommunicationIdRemovedEvent event;

    public CommunicationIdRemovedEventPb(byte[] message) throws InvalidProtocolBufferException {
        this.event = proto.skysign.event.CommunicationIdRemovedEvent.parseFrom(message);
    }

    @Override
    public String toString() {
        return this.event.toString().replaceAll("\\r\\n|\\r|\\n", " ");
    }

    public CommunicationIdRemovedEvent getEvent() {
        return new CommunicationIdRemovedEvent(
            new CommunicationId(this.event.getCommunicationId()),
            new Version(this.event.getVersion()));
    }
}