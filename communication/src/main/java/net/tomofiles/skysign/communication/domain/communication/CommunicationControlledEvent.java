package net.tomofiles.skysign.communication.domain.communication;

import java.time.LocalDateTime;

import lombok.AllArgsConstructor;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.ToString;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleId;
import net.tomofiles.skysign.communication.event.Event;

@AllArgsConstructor
@Getter
@EqualsAndHashCode
@ToString
public class CommunicationControlledEvent implements Event {
    private final CommunicationId communicationId;
    private final VehicleId vehicleId;
    private final MissionId missionId;
    private final LocalDateTime updateDateTime;
}