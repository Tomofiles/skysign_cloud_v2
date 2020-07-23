package net.tomofiles.skysign.communication.domain.communication.component;

import java.util.List;

import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class CommunicationComponentDto {
    private String id;
    private String missionId;
    private TelemetryComponentDto telemetry;
    private List<CommandComponentDto> commands;
}