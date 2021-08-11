package net.tomofiles.skysign.communication.service.dpo;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class WaypointDpo {
    private double latitude;
    private double longitude;
    private double relativeHeight;
    private double speed;
}