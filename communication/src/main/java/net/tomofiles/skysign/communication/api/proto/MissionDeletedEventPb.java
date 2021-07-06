package net.tomofiles.skysign.communication.api.proto;

import com.google.protobuf.InvalidProtocolBufferException;

import net.tomofiles.skysign.communication.api.event.event.MissionDeletedEvent;


public class MissionDeletedEventPb {

    private final proto.skysign.event.MissionDeletedEvent event;

    public MissionDeletedEventPb(byte[] message) throws InvalidProtocolBufferException {
        this.event = proto.skysign.event.MissionDeletedEvent.parseFrom(message);
    }

    @Override
    public String toString() {
        return this.event.toString().replaceAll("\\r\\n|\\r|\\n", " ");
    }

    public MissionDeletedEvent getEvent() {
        return new MissionDeletedEvent(
            this.event.getUploadMissionId());
    }
}