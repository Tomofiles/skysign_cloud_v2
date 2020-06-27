package net.tomofiles.skysign.communication.domain.communication.component;

import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class CommandComponentDto {
    private String id;
    private String type;
}