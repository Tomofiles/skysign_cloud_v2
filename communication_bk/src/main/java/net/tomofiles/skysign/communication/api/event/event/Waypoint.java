package net.tomofiles.skysign.communication.api.event.event;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class Waypoint {
    private double latitude;
    private double longitude;
    private double relativeHeight;
    private double speed;
}