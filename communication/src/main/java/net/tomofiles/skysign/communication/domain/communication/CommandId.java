package net.tomofiles.skysign.communication.domain.communication;

import java.util.UUID;

import lombok.AccessLevel;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.RequiredArgsConstructor;

@Getter
@RequiredArgsConstructor(access = AccessLevel.PUBLIC)
@EqualsAndHashCode(of = {"id"})
public class CommandId {
    private final String id;

    public static CommandId newId() {
        return new CommandId(UUID.randomUUID().toString());
    }
}