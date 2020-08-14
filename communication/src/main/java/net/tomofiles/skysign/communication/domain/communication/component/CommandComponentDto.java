package net.tomofiles.skysign.communication.domain.communication.component;

import java.time.LocalDateTime;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.EqualsAndHashCode;
import lombok.ToString;

@Data
@AllArgsConstructor
@EqualsAndHashCode
@ToString
public class CommandComponentDto {
    private String id;
    private String type;
    private LocalDateTime time;
}