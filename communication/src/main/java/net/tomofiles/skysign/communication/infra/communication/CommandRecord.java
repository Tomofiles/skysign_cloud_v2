package net.tomofiles.skysign.communication.infra.communication;

import lombok.Data;
import lombok.EqualsAndHashCode;

@Data
@EqualsAndHashCode(of = {"id"})
public class CommandRecord {
    private String id;
    private String commId;
    private String type;
}