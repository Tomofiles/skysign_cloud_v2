package net.tomofiles.skysign.vehicle.domain.vehicle;

import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.RequiredArgsConstructor;
import lombok.ToString;

@Getter
@RequiredArgsConstructor
@EqualsAndHashCode
@ToString
public class FlightplanId {
    private final String id;
}