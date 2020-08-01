package net.tomofiles.skysign.mission.domain.mission;

import lombok.AccessLevel;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.RequiredArgsConstructor;
import lombok.Setter;

@RequiredArgsConstructor(access = AccessLevel.PACKAGE)
@EqualsAndHashCode(of = {"id"})
public class Mission {
    @Getter
    private final MissionId id;
    
    @Getter
    @Setter(value = AccessLevel.PACKAGE)
    private String missionName = null;

    @Getter
    @Setter(value = AccessLevel.PACKAGE)
    private Navigation navigation = null;

    @Getter
    @Setter(value = AccessLevel.PACKAGE)
    private Version version;

    @Getter
    @Setter(value = AccessLevel.PACKAGE)
    private Version newVersion;

    public void nameMission(String name) {
        this.missionName = name;
        this.newVersion = Version.newVersion();
    }

    public void replaceNavigationWith(Navigation navigation) {
        this.navigation = navigation;
        this.newVersion = Version.newVersion();
    }
}