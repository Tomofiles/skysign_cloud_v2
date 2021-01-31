package net.tomofiles.skysign.communication.api.proto;

import com.google.protobuf.InvalidProtocolBufferException;

import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.domain.vehicle.CommunicationIdChangedEvent;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleId;
import net.tomofiles.skysign.communication.domain.vehicle.Version;

public class CommunicationIdChangedEventPb {

    private final proto.skysign.event.CommunicationIdChangedEvent event;

    public CommunicationIdChangedEventPb(byte[] message) throws InvalidProtocolBufferException {
        this.event = proto.skysign.event.CommunicationIdChangedEvent.parseFrom(message);
    }

    @Override
    public String toString() {
        return this.event.toString().replaceAll("\\r\\n|\\r|\\n", " ");
    }

    public CommunicationIdChangedEvent getEvent() {
        return new CommunicationIdChangedEvent(
            new CommunicationId(this.event.getBeforeId()),
            new CommunicationId(this.event.getAfterId()),
            new VehicleId(this.event.getVehicleId()),
            new Version(this.event.getVersion()));
    }
}