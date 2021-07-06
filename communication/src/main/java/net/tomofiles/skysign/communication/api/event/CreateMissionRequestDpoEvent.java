package net.tomofiles.skysign.communication.api.event;

import java.util.List;
import java.util.stream.Collectors;

import lombok.RequiredArgsConstructor;
import net.tomofiles.skysign.communication.api.event.event.MissionCreatedEvent;
import net.tomofiles.skysign.communication.domain.mission.MissionId;
import net.tomofiles.skysign.communication.service.dpo.CreateMissionRequestDpo;
import net.tomofiles.skysign.communication.service.dpo.WaypointDpo;

@RequiredArgsConstructor
public class CreateMissionRequestDpoEvent implements CreateMissionRequestDpo {

    private final MissionCreatedEvent event;

    @Override
    public MissionId getMissionId() {
        return new MissionId(this.event.getMissionId());
    }

    @Override
    public List<WaypointDpo> getWaypoints() {
        return this.event.getWaypoints().stream()
            .map(w -> {
                return new WaypointDpo(
                    w.getLatitude(), 
                    w.getLongitude(), 
                    w.getRelativeHeight(), 
                    w.getSpeed());
            })
            .collect(Collectors.toList());
    }
}