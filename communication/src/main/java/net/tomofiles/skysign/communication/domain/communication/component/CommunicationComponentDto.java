package net.tomofiles.skysign.communication.domain.communication.component;

import java.util.List;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.EqualsAndHashCode;
import lombok.ToString;

@Data
@AllArgsConstructor
@EqualsAndHashCode
@ToString
public class CommunicationComponentDto {
    private String id;
    private String missionId;
    private TelemetryComponentDto telemetry;
    private List<CommandComponentDto> commands;
}