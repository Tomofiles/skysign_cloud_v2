package net.tomofiles.skysign.communication.domain.communication;

import lombok.AllArgsConstructor;
import lombok.Getter;

@Getter
@AllArgsConstructor
public enum CommandType {
    UPLOAD(1, "upload"),
    NONE(999, "none");

    private final int code;
    private final String name;
}