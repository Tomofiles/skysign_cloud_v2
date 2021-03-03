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
    private final boolean isCarbonCopy;

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

    Mission(MissionId id, boolean isCarbonCopy, Version version, Generator generator) {
        this.id = id;
        this.isCarbonCopy = isCarbonCopy;
        this.version = version;
        this.newVersion = version;

        this.generator = generator;
    }

    static Mission newOriginal(MissionId id, Version version, Generator generator) {
        return new Mission(id, false, version, generator);
    } 

    static Mission newCarbonCopy(MissionId id, Version version, Generator generator) {
        return new Mission(id, true, version, generator);
    } 

    public void nameMission(String name) {
        if (this.isCarbonCopy) {
            throw new CannotChangeMissionException("cannot change carbon copied mission");
        }

        this.missionName = name;
        this.newVersion = this.generator.newVersion();
    }

    public void replaceNavigationWith(Navigation navigation) {
        if (this.isCarbonCopy) {
            throw new CannotChangeMissionException("cannot change carbon copied mission");
        }

        this.navigation = navigation;
        this.newVersion = this.generator.newVersion();
    }
}