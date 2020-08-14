package net.tomofiles.skysign.communication.domain.communication;

import java.time.LocalDateTime;

public interface Generator {
    public CommandId newCommandId();
    public LocalDateTime newTime();
}