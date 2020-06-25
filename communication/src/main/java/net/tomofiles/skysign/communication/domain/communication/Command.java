package net.tomofiles.skysign.communication.domain.communication;

import lombok.AccessLevel;
import lombok.EqualsAndHashCode;
import lombok.RequiredArgsConstructor;

@RequiredArgsConstructor(access = AccessLevel.PACKAGE)
@EqualsAndHashCode(of = {"id"})
class Command {
    private final CommandId id;
    private final CommandType type;
}