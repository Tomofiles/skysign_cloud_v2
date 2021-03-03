package net.tomofiles.skysign.mission.api.proto;

import com.google.protobuf.InvalidProtocolBufferException;

import net.tomofiles.skysign.mission.api.event.event.MissionCopiedWhenFlightplanCopiedEvent;


public class MissionCopiedWhenFlightplanCopiedEventPb {

    private final proto.skysign.event.MissionCopiedWhenCopiedEvent event;

    public MissionCopiedWhenFlightplanCopiedEventPb(byte[] message) throws InvalidProtocolBufferException {
        this.event = proto.skysign.event.MissionCopiedWhenCopiedEvent.parseFrom(message);
    }

    @Override
    public String toString() {
        return this.event.toString().replaceAll("\\r\\n|\\r|\\n", " ");
    }

    public MissionCopiedWhenFlightplanCopiedEvent getEvent() {
        return new MissionCopiedWhenFlightplanCopiedEvent(
            this.event.getOriginalMissionId(),
            this.event.getNewMissionId());
    }
}