package net.tomofiles.skysign.communication.generator;

import java.time.LocalDateTime;
import java.util.UUID;

import org.springframework.stereotype.Component;

import net.tomofiles.skysign.communication.domain.communication.CommandId;
import net.tomofiles.skysign.communication.domain.communication.Generator;

@Component
public class CommunicationGeneratorUUID implements Generator {
    @Override
    public CommandId newCommandId() {
        return new CommandId(UUID.randomUUID().toString());
    }

    @Override
    public LocalDateTime newTime() {
        return LocalDateTime.now();
    }
}