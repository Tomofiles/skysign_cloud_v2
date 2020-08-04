package net.tomofiles.skysign.mission.domain.mission;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.Mock;

import static com.google.common.truth.Truth.assertThat;
import static org.junit.jupiter.api.Assertions.assertAll;
import static org.mockito.Mockito.when;
import static org.mockito.MockitoAnnotations.initMocks;

import java.util.UUID;
import java.util.function.Supplier;

import static net.tomofiles.skysign.mission.domain.mission.MissionObjectMother.newSingleNavigationMission;
import static net.tomofiles.skysign.mission.domain.mission.MissionObjectMother.newSingleNavigation;

public class ManagementMissionsTests {

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

    @Mock
    private MissionRepository repository;

    @BeforeEach
    public void beforeEach() {
        initMocks(this);
    }

    /**
     * Userが、新しいMissionエンティティを作成する。<br>
     * Missionエンティティの初期状態を検証する。
     */
    @Test
    public void createNewMissionTest() {
        Mission mission = MissionFactory.newInstance(DEFAULT_GENERATOR.get());

        assertAll(
            () -> assertThat(mission.getId()).isEqualTo(DEFAULT_MISSION_ID),
            () -> assertThat(mission.getMissionName()).isNull(),
            () -> assertThat(mission.getNavigation()).isNull(),
            () -> assertThat(mission.getVersion()).isEqualTo(DEFAULT_VERSION1),
            () -> assertThat(mission.getNewVersion()).isEqualTo(DEFAULT_VERSION1)
        );
    }

    /**
     * Userが、既存のMissionエンティティに対してMission Nameを更新する。
     */
    @Test
    public void changeMissionsNameTest() {
        when(repository.getById(DEFAULT_MISSION_ID)).thenReturn(newSingleNavigationMission(DEFAULT_GENERATOR.get()));

        Mission mission = repository.getById(DEFAULT_MISSION_ID);

        String newMissionName = "new mission";
        mission.nameMission(newMissionName);

        assertAll(
            () -> assertThat(mission.getMissionName()).isEqualTo(newMissionName),
            () -> assertThat(mission.getVersion()).isEqualTo(DEFAULT_VERSION1),
            () -> assertThat(mission.getNewVersion()).isEqualTo(DEFAULT_VERSION2)
        );
    }

    /**
     * Userが、Missionエンティティに新たなNavigationを追加する。<br>
     * Navigation自体の詳細なテストは、MissionNavigationTestsにまとめている。
     */
    @Test
    public void addNavigationToMissionTest() {
        Mission mission = MissionFactory.newInstance(DEFAULT_GENERATOR.get());

        mission.replaceNavigationWith(newSingleNavigation());

        assertAll(
            () -> assertThat(mission.getNavigation()).isEqualTo(newSingleNavigation()),
            () -> assertThat(mission.getVersion()).isEqualTo(DEFAULT_VERSION1),
            () -> assertThat(mission.getNewVersion()).isEqualTo(DEFAULT_VERSION2)
        );
    }
}