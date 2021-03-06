package net.tomofiles.skysign.mission.domain.mission;

import org.junit.jupiter.api.Test;

import static com.google.common.truth.Truth.assertThat;
import static org.junit.jupiter.api.Assertions.assertAll;

import java.util.UUID;
import java.util.function.Supplier;

import net.tomofiles.skysign.mission.domain.mission.component.MissionComponentDto;

import static net.tomofiles.skysign.mission.domain.mission.MissionObjectMother.newSeveralNavigationCarbonCopiedMission;
import static net.tomofiles.skysign.mission.domain.mission.ComponentDtoObjectMother.newSeveralNavigationCarbonCopiedMissionComponentDto;

public class ComponentMissionTests {
    
    private static final MissionId DEFAULT_MISSION_ID = new MissionId(UUID.randomUUID().toString());
    private static final Version DEFAULT_VERSION1 = new Version(UUID.randomUUID().toString());
    private static final Version DEFAULT_VERSION2 = new Version(UUID.randomUUID().toString());
    private static final Supplier<Generator> DEFAULT_GENERATOR = () -> {
        return new Generator(){
            private int count = 0;

            @Override
            public MissionId newMissionId() {
                return DEFAULT_MISSION_ID;
            }

            @Override
            public Version newVersion() {
                if (count == 0) {
                    count++;
                    return DEFAULT_VERSION1;
                } else {
                    return DEFAULT_VERSION2;
                }
            }
        };
    };

    /**
     * DTOからMissionエンティティを組み立てる。
     */
    @Test
    public void assembleIntoMissionTest() {
        Mission mission = MissionFactory.assembleFrom(
                newSeveralNavigationCarbonCopiedMissionComponentDto(
                        DEFAULT_MISSION_ID,
                        DEFAULT_VERSION1
                ),
                DEFAULT_GENERATOR.get());

        Mission expectedMission = newSeveralNavigationCarbonCopiedMission(
                DEFAULT_MISSION_ID,
                DEFAULT_VERSION1,
                DEFAULT_GENERATOR.get());

        assertAll(
            () -> assertThat(mission.getId()).isEqualTo(expectedMission.getId()),
            () -> assertThat(mission.getMissionName()).isEqualTo(expectedMission.getMissionName()),
            () -> assertThat(mission.getNavigation()).isEqualTo(expectedMission.getNavigation()),
            () -> assertThat(mission.isCarbonCopy()).isEqualTo(expectedMission.isCarbonCopy()),
            () -> assertThat(mission.getVersion()).isEqualTo(expectedMission.getVersion()),
            () -> assertThat(mission.getNewVersion()).isEqualTo(expectedMission.getNewVersion())
        );
    }

    /**
     * MissionエンティティからDTOに分解する。
     */
    @Test
    public void takeApartMissionTest() {
        MissionComponentDto dto = MissionFactory.takeApart(
                newSeveralNavigationCarbonCopiedMission(
                        DEFAULT_MISSION_ID,
                        DEFAULT_VERSION1,
                        DEFAULT_GENERATOR.get()));

        assertThat(dto).isEqualTo(newSeveralNavigationCarbonCopiedMissionComponentDto(DEFAULT_MISSION_ID, DEFAULT_VERSION1));
    }
}