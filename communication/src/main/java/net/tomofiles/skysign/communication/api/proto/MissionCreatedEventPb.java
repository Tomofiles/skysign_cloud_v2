package net.tomofiles.skysign.communication.api.proto;

import java.util.stream.Collectors;

import com.google.protobuf.InvalidProtocolBufferException;

import net.tomofiles.skysign.communication.api.event.event.MissionCreatedEvent;
import net.tomofiles.skysign.communication.api.event.event.Waypoint;


public class MissionCreatedEventPb {

    private final proto.skysign.event.MissionCreatedEvent event;

    public MissionCreatedEventPb(byte[] message) throws InvalidProtocolBufferException {
        this.event = proto.skysign.event.MissionCreatedEvent.parseFrom(message);
    }

    @Override
    public String toString() {
        return this.event.toString().replaceAll("\\r\\n|\\r|\\n", " ");
    }

    public MissionCreatedEvent getEvent() {
        return new MissionCreatedEvent(
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