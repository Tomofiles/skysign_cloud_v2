package net.tomofiles.skysign.communication.domain.communication;

import java.time.LocalDateTime;

import lombok.AccessLevel;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.RequiredArgsConstructor;

@Getter
@RequiredArgsConstructor(access = AccessLevel.PACKAGE)
@EqualsAndHashCode(of = {"id"})
class Command {
    private final CommandId id;
    private final CommandType type;
    private final LocalDateTime time;

    public static Command empty(CommandId id) {
        return new Command(id, CommandType.NONE, null);
    }
}