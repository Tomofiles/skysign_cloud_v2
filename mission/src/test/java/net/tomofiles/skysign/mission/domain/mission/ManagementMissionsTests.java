package net.tomofiles.skysign.mission.domain.mission;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.Mock;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertNull;
import static org.mockito.Mockito.when;
import static org.mockito.MockitoAnnotations.initMocks;

public class ManagementMissionsTests {
    
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
        MissionId id = MissionId.newId();

        Mission mission = MissionFactory.newInstance(id);

        assertEquals(mission.getId(), id);
        assertNull(mission.getMissionName());
        assertNull(mission.getNavigation());
        assertNotNull(mission.getVersion());
        assertNotNull(mission.getNewVersion());
        assertEquals(mission.getVersion(), mission.getNewVersion());
    }

    /**
     * Userが、既存のMissionエンティティに対してMission Nameを更新する。<br>
     * Mission Name以外の変化が無いことを検証する。
     */
    @Test
    public void changeMissionsNameTest() {
        MissionId id = MissionId.newId();

        String oldMissionName = "old mission";
        Version version = Version.newVersion();

        Mission before = new Mission(id);
        before.setMissionName(oldMissionName);
        before.setVersion(version);
        before.setNewVersion(version);

        when(repository.getById(id)).thenReturn(before);

        Mission mission = repository.getById(id);

        String newMissionName = "new mission";
        mission.nameMission(newMissionName);

        assertEquals(mission.getId(), id);
        assertEquals(mission.getMissionName(), newMissionName);
        assertNull(mission.getNavigation());
        assertEquals(mission.getVersion(), version);
        assertNotEquals(mission.getVersion(), mission.getNewVersion());
    }

    /**
     * Userが、Missionエンティティに新たなNavigationを追加する。<br>
     * Navigation自体の詳細なテストは、MissionNavigationTestsにまとめている。
     */
    @Test
    public void addNavigationToMissionTest() {
        MissionId id = MissionId.newId();

        Mission mission = MissionFactory.newInstance(id);

        Navigation navigation = new Navigation();
        navigation.setTakeoffPointGroundHeight(Height.fromM(0.0));
        navigation.pushNextWaypoint(
            new GeodesicCoordinates(1.0, 2.0),
            Height.fromM(3.0),
            Speed.fromMS(4.0));

        mission.replaceNavigationWith(navigation);

        Navigation expectNavigation = new Navigation();
        expectNavigation.setTakeoffPointGroundHeight(Height.fromM(0.0));
        expectNavigation.pushNextWaypoint(
            new GeodesicCoordinates(1.0, 2.0),
            Height.fromM(3.0),
            Speed.fromMS(4.0));

        assertEquals(mission.getId(), id);
        assertNull(mission.getMissionName());
        assertEquals(mission.getNavigation(), expectNavigation);
        assertNotEquals(mission.getVersion(), mission.getNewVersion());
    }
}