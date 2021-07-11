package net.tomofiles.skysign.communication.api.proto;

import java.util.stream.Collectors;

import com.google.protobuf.InvalidProtocolBufferException;

import net.tomofiles.skysign.communication.api.event.event.CopiedMissionCreatedEvent;
import net.tomofiles.skysign.communication.api.event.event.Waypoint;


public class CopiedMissionCreatedEventPb {

    private final proto.skysign.event.CopiedMissionCreatedEvent event;

    public CopiedMissionCreatedEventPb(byte[] message) throws InvalidProtocolBufferException {
        this.event = proto.skysign.event.CopiedMissionCreatedEvent.parseFrom(message);
    }

    @Override
    public String toString() {
        return this.event.toString().replaceAll("\\r\\n|\\r|\\n", " ");
    }

    public CopiedMissionCreatedEvent getEvent() {
        return new CopiedMissionCreatedEvent(
            this.event.getMission().getNavigation().getUploadId(),
            this.event.getMission().getNavigation().getWaypointsList().stream()
                .map(waypoint -> {
                    return new Waypoint(
                        waypoint.getLatitude(),
                        waypoint.getLongitude(),
                        waypoint.getRelativeHeight(),
                        waypoint.getSpeed()
                    );
                })
                .collect(Collectors.toList())
            );
    }
}