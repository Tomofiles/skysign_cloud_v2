package net.tomofiles.skysign.vehicle.api.proto;

import com.google.protobuf.InvalidProtocolBufferException;

import net.tomofiles.skysign.vehicle.api.event.event.VehicleCopiedWhenFlightplanCopiedEvent;


public class VehicleCopiedWhenFlightplanCopiedEventPb {

    private final proto.skysign.event.VehicleCopiedWhenCopiedEvent event;

    public VehicleCopiedWhenFlightplanCopiedEventPb(byte[] message) throws InvalidProtocolBufferException {
        this.event = proto.skysign.event.VehicleCopiedWhenCopiedEvent.parseFrom(message);
    }

    @Override
    public String toString() {
        return this.event.toString().replaceAll("\\r\\n|\\r|\\n", " ");
    }

    public VehicleCopiedWhenFlightplanCopiedEvent getEvent() {
        return new VehicleCopiedWhenFlightplanCopiedEvent(
            this.event.getOriginalVehicleId(),
            this.event.getNewVehicleId());
    }
}