package net.tomofiles.skysign.communication.usecase.dto;

import java.util.Arrays;

import lombok.AllArgsConstructor;
import lombok.Getter;
import net.tomofiles.skysign.communication.domain.communication.CommandType;

@Getter
@AllArgsConstructor
public enum ControlCommandType {
    UPLOAD(CommandType.UPLOAD),
    ARM(CommandType.ARM),
    DISARM(CommandType.DISARM),
    NONE(CommandType.NONE);

    private CommandType type;

    public static ControlCommandType valueOf(CommandType type) {
        return Arrays.asList(ControlCommandType.values()).stream()
                .filter(t -> t.type == type)
                .findAny()
                .orElse(null);
    }
}