package net.tomofiles.skysign.vehicle.domain.vehicle;

import lombok.AllArgsConstructor;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.ToString;
import net.tomofiles.skysign.vehicle.event.Event;

@AllArgsConstructor
@Getter
@EqualsAndHashCode
@ToString
public class CommunicationIdGaveEvent implements Event {
    private final CommunicationId communicationId;
    private final VehicleId vehicleId;
    private final Version version;
}