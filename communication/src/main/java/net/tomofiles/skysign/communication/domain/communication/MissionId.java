package net.tomofiles.skysign.communication.domain.communication;

import lombok.AccessLevel;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.RequiredArgsConstructor;

@Getter
@RequiredArgsConstructor(access = AccessLevel.PUBLIC)
@EqualsAndHashCode(of = {"id"})
public class MissionId {
    private final String id;
}