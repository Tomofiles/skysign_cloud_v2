package net.tomofiles.skysign.communication.domain.communication;

import lombok.Getter;

@Getter
public enum CommandType {
    ARM,
    DISARM,
    UPLOAD,
    START,
    PAUSE,
    TAKEOFF,
    LAND,
    RETURN,
    NONE;
}