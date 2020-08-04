package net.tomofiles.skysign.mission.domain.mission;

import lombok.AccessLevel;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.Setter;
import lombok.ToString;

@EqualsAndHashCode(of = {"id"})
@ToString
public class Mission {
    @Getter
    private final MissionId id;

    private final Generator generator;
    
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

    Mission(Generator generator) {
        this.id = generator.newMissionId();
        Version version = generator.newVersion();

        this.version = version;
        this.newVersion = version;

        this.generator = generator;
    }

    public void nameMission(String name) {
        this.missionName = name;
        this.newVersion = this.generator.newVersion();
    }

    public void replaceNavigationWith(Navigation navigation) {
        this.navigation = navigation;
        this.newVersion = this.generator.newVersion();
    }
}