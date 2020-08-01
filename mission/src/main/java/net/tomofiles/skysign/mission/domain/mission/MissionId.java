package net.tomofiles.skysign.mission.domain.mission;

import java.util.UUID;

import lombok.AccessLevel;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.RequiredArgsConstructor;

@Getter
@RequiredArgsConstructor(access = AccessLevel.PUBLIC)
@EqualsAndHashCode(of = {"id"})
public class MissionId {
    private final String id;

    public static MissionId newId() {
        return new MissionId(UUID.randomUUID().toString());
    }
    
}