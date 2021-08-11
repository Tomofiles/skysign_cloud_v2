package net.tomofiles.skysign.communication.domain.communication;

import lombok.AccessLevel;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.RequiredArgsConstructor;
import lombok.ToString;

@Getter
@RequiredArgsConstructor(access = AccessLevel.PACKAGE)
@EqualsAndHashCode(of = {"id"})
@ToString
class UploadMission {
    private final CommandId id;
    private final MissionId missionId;

    public static UploadMission empty(CommandId id) {
        return new UploadMission(id, null);
    }
}