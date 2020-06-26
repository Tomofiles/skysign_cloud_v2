package net.tomofiles.skysign.communication.usecase.dto;

import lombok.AllArgsConstructor;
import lombok.Getter;
import net.tomofiles.skysign.communication.domain.communication.CommandType;

@Getter
@AllArgsConstructor
public enum ControlCommandType {
    UPLOAD(CommandType.UPLOAD),
    NONE(CommandType.NONE);

    private CommandType type;
}