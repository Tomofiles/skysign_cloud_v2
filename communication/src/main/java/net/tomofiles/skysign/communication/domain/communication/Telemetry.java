package net.tomofiles.skysign.communication.domain.communication;

import lombok.AllArgsConstructor;
import lombok.Getter;

@AllArgsConstructor
@Getter
class Telemetry {
    private final double latitude;
    private final double longitude;
}