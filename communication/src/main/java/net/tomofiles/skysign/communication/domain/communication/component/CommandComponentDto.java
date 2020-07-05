package net.tomofiles.skysign.communication.domain.communication.component;

import java.time.LocalDateTime;

import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class CommandComponentDto {
    private String id;
    private String type;
    private LocalDateTime time;
}