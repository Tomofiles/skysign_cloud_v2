package net.tomofiles.skysign.communication.infra.communication;

import java.time.LocalDateTime;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.EqualsAndHashCode;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
@EqualsAndHashCode(of = {"id"})
public class CommandRecord {
    private String id;
    private String communicationId;
    private String type;
    private LocalDateTime time;
}