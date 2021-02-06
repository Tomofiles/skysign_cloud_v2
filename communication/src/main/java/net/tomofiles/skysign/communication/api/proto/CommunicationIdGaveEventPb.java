package net.tomofiles.skysign.communication.api.proto;

import com.google.protobuf.InvalidProtocolBufferException;

import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.domain.vehicle.CommunicationIdGaveEvent;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleId;
import net.tomofiles.skysign.communication.domain.vehicle.Version;

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
            new CommunicationId(this.event.getCommunicationId()),
            new VehicleId(this.event.getVehicleId()),
            new Version(this.event.getVersion()));
    }
}