package net.tomofiles.skysign.vehicle.api.event.event;

import lombok.AllArgsConstructor;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.ToString;

@AllArgsConstructor
@Getter
@EqualsAndHashCode
@ToString
public class VehicleCopiedWhenFlightplanCopiedEvent {
    private final String originalId;
    private final String newId;
}