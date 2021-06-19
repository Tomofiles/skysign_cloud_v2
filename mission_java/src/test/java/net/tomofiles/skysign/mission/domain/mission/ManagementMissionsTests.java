package net.tomofiles.skysign.mission.domain.mission;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.Mock;

import static com.google.common.truth.Truth.assertThat;
import static org.junit.Assert.assertThrows;
import static org.junit.jupiter.api.Assertions.assertAll;
import static org.mockito.MockitoAnnotations.initMocks;

import java.util.UUID;
import java.util.function.Supplier;

import static net.tomofiles.skysign.mission.domain.mission.MissionObjectMother.newSingleNavigation;
import static net.tomofiles.skysign.mission.domain.mission.MissionObjectMother.newSeveralNavigation;

public class ManagementMissionsTests {

    private static final MissionId DEFAULT_MISSION_ID1 = new MissionId(UUID.randomUUID().toString());
    private static final MissionId DEFAULT_MISSION_ID2 = new MissionId(UUID.randomUUID().toString());
    private static final Version DEFAULT_VERSION1 = new Version(UUID.randomUUID().toString());
    private static final Version DEFAULT_VERSION2 = new Version(UUID.randomUUID().toString());
    private static final Supplier<Generator> DEFAULT_GENERATOR = () -> {
        return new Generator(){
            private int count = 0;

            @Override
            public MissionId newMissionId() {
                return DEFAULT_MISSION_ID1;
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
            () -> assertThat(mission.getId()).isEqualTo(DEFAULT_MISSION_ID1),
            () -> assertThat(mission.getMissionName()).isNull(),
            () -> assertThat(mission.getNavigation()).isNull(),
            () -> assertThat(mission.isCarbonCopy()).isFalse(),
            () -> assertThat(mission.getVersion()).isEqualTo(DEFAULT_VERSION1),
            () -> assertThat(mission.getNewVersion()).isEqualTo(DEFAULT_VERSION1)
        );
    }

    /**
     * Userが、新しいMissionエンティティに対してMission Nameを付与する。
     */
    @Test
    public void changeMissionsNameTest() {
        Mission mission = MissionFactory.newInstance(DEFAULT_GENERATOR.get());

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

    /**
     * 既存のMissionエンティティのカーボンコピーを作成する。<br>
     * コピーされたMissionエンティティの内部状態を検証する。
     */
    @Test
    public void carbonCopyMissionTest() {
        Generator generator = DEFAULT_GENERATOR.get();
        Mission original = MissionFactory.newInstance(generator);

        String missionName = "now mission";
        Navigation navigation = newSeveralNavigation();

        original.setMissionName(missionName);
        original.setNavigation(navigation);

        Mission mission = MissionFactory.copy(original, DEFAULT_MISSION_ID2, generator);

        assertAll(
            () -> assertThat(mission.getId()).isEqualTo(DEFAULT_MISSION_ID2),
            () -> assertThat(mission.getMissionName()).isEqualTo(missionName),
            () -> assertThat(mission.getNavigation()).isEqualTo(newSeveralNavigation()),
            () -> assertThat(mission.isCarbonCopy()).isTrue(),
            () -> assertThat(mission.getVersion()).isEqualTo(DEFAULT_VERSION1),
            () -> assertThat(mission.getNewVersion()).isEqualTo(DEFAULT_VERSION1)
        );
    }

    /**
     * Navigationがない既存のMissionエンティティのカーボンコピーを作成する。<br>
     * コピーされたMissionエンティティの内部状態を検証する。
     */
    @Test
    public void carbonCopyNoneNavigationMissionTest() {
        Generator generator = DEFAULT_GENERATOR.get();
        Mission original = MissionFactory.newInstance(generator);

        String missionName = "now mission";

        original.setMissionName(missionName);

        Mission mission = MissionFactory.copy(original, DEFAULT_MISSION_ID2, generator);

        assertAll(
            () -> assertThat(mission.getId()).isEqualTo(DEFAULT_MISSION_ID2),
            () -> assertThat(mission.getMissionName()).isEqualTo(missionName),
            () -> assertThat(mission.getNavigation()).isNull(),
            () -> assertThat(mission.isCarbonCopy()).isTrue(),
            () -> assertThat(mission.getVersion()).isEqualTo(DEFAULT_VERSION1),
            () -> assertThat(mission.getNewVersion()).isEqualTo(DEFAULT_VERSION1)
        );
    }

    /**
     * カーボンコピーされたMissionエンティティに対してMission Nameを付与する。
     * 更新時に例外がスローされることを検証する。
     */
    @Test
    public void cannotChangeErrorWhenChangeCarbonCopiedMissionsNameTest() {
        Generator generator = DEFAULT_GENERATOR.get();
        Mission original = MissionFactory.newInstance(generator);

        String missionName = "now mission";
        Navigation navigation = newSeveralNavigation();

        original.setMissionName(missionName);
        original.setNavigation(navigation);

        Mission mission = MissionFactory.copy(original, DEFAULT_MISSION_ID2, generator);

        String newMissionName = "new mission";

        CannotChangeMissionException e
                = assertThrows(
                        CannotChangeMissionException.class,
                        () -> mission.nameMission(newMissionName));

        assertThat(e).hasMessageThat().contains("cannot change carbon copied mission");
    }

    /**
     * カーボンコピーされたMissionエンティティに対してNavigationを付与する。
     * 更新時に例外がスローされることを検証する。
     */
    @Test
    public void cannotChangeErrorWhenChangeCarbonCopiedMissionsNavigationTest() {
        Generator generator = DEFAULT_GENERATOR.get();
        Mission original = MissionFactory.newInstance(generator);

        String missionName = "now mission";

        original.setMissionName(missionName);

        Mission mission = MissionFactory.copy(original, DEFAULT_MISSION_ID2, generator);

        CannotChangeMissionException e
                = assertThrows(
                        CannotChangeMissionException.class,
                        () -> mission.replaceNavigationWith(newSeveralNavigation()));

        assertThat(e).hasMessageThat().contains("cannot change carbon copied mission");
    }

}