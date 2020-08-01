package net.tomofiles.skysign.mission.infra.mission;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.InjectMocks;
import org.mockito.Mock;

import net.tomofiles.skysign.mission.domain.mission.GeodesicCoordinates;
import net.tomofiles.skysign.mission.domain.mission.Height;
import net.tomofiles.skysign.mission.domain.mission.Mission;
import net.tomofiles.skysign.mission.domain.mission.MissionFactory;
import net.tomofiles.skysign.mission.domain.mission.MissionId;
import net.tomofiles.skysign.mission.domain.mission.MissionRepository;
import net.tomofiles.skysign.mission.domain.mission.Navigation;
import net.tomofiles.skysign.mission.domain.mission.Speed;
import net.tomofiles.skysign.mission.domain.mission.Version;
import net.tomofiles.skysign.mission.infra.common.DeleteCondition;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNull;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;
import static org.mockito.MockitoAnnotations.initMocks;

import java.util.Arrays;
import java.util.List;

public class MissionRepositoryTests {
    
    @Mock
    private MissionMapper missionMapper;

    @Mock
    private WaypointMapper waypointMapper;

    @InjectMocks
    private MissionRepository repository = new MissionRepositoryImpl();

    @BeforeEach
    public void beforeEach() {
        initMocks(this);
    }

    /**
     * リポジトリーからMissionエンティティを一つ取得する。<br>
     * Waypointの順序がバラバラでも、リポジトリーがソートして返却することを検証する
     */
    @Test
    public void getMissionByIdTest() {
        MissionId id = MissionId.newId();
        String missionName = "mission name";
        Height takeoffPointGroundHeight = Height.fromM(10.0);
        Version version = Version.newVersion();

        double latitude1 = 0.0d;
        double longitude1 = 1.0d;
        double heightWGS84M1 = 2.0d;
        double speedMS1 = 4.0d;
        double latitude2 = 10.0d;
        double longitude2 = 11.0d;
        double heightWGS84M2 = 12.0d;
        double speedMS2 = 14.0d;
        double latitude3 = 20.0d;
        double longitude3 = 21.0d;
        double heightWGS84M3 = 22.0d;
        double speedMS3 = 24.0d;

        MissionRecord missionRecord = new MissionRecord(
                id.getId(),
                missionName,
                takeoffPointGroundHeight.getHeightM(),
                version.getVersion(),
                version.getVersion());

        WaypointRecord waypointRecord1 = new WaypointRecord(
                id.getId(),
                1,
                latitude1,
                longitude1,
                heightWGS84M1,
                speedMS1);
        WaypointRecord waypointRecord2 = new WaypointRecord(
                id.getId(),
                2,
                latitude2,
                longitude2,
                heightWGS84M2,
                speedMS2);
        WaypointRecord waypointRecord3 = new WaypointRecord(
                id.getId(),
                3,
                latitude3,
                longitude3,
                heightWGS84M3,
                speedMS3);

        when(missionMapper.find(id.getId())).thenReturn(missionRecord);
        when(waypointMapper.find(id.getId())).thenReturn(Arrays.asList(new WaypointRecord[] {
            waypointRecord3, // 順序がバラバラ
            waypointRecord1, // 順序がバラバラ
            waypointRecord2  // 順序がバラバラ
        }));

        Mission mission = repository.getById(id);

        Navigation expectNavigation = new Navigation();
        expectNavigation.setTakeoffPointGroundHeight(takeoffPointGroundHeight);
        expectNavigation.pushNextWaypoint(
            new GeodesicCoordinates(latitude1, longitude1),
            Height.distanceFrom(Height.fromM(heightWGS84M1), takeoffPointGroundHeight),
            Speed.fromMS(speedMS1));
        expectNavigation.pushNextWaypoint(
            new GeodesicCoordinates(latitude2, longitude2),
            Height.distanceFrom(Height.fromM(heightWGS84M2), takeoffPointGroundHeight),
            Speed.fromMS(speedMS2));
        expectNavigation.pushNextWaypoint(
            new GeodesicCoordinates(latitude3, longitude3),
            Height.distanceFrom(Height.fromM(heightWGS84M3), takeoffPointGroundHeight),
            Speed.fromMS(speedMS3));

        assertEquals(mission.getId(), id);
        assertEquals(mission.getMissionName(), missionName);
        assertEquals(mission.getNavigation(), expectNavigation);
        assertEquals(mission.getVersion(), version);
        assertEquals(mission.getNewVersion(), version);
    }

    /**
     * リポジトリーからMissionエンティティを一つ取得する。<br>
     * 対象のエンティティが存在しない場合、NULLが返却されることを検証する。
     */
    @Test
    public void getNoMissionByIdTest() {
        MissionId id = MissionId.newId();

        Mission mission = repository.getById(id);

        assertNull(mission);
    }

    /**
     * リポジトリーからMissionエンティティをすべて取得する。
     */
    @Test
    public void getAllMissionsTest() {
        MissionId id = MissionId.newId();
        String missionName = "mission name";
        Height takeoffPointGroundHeight = Height.fromM(10.0);
        Version version = Version.newVersion();

        double latitude1 = 0.0d;
        double longitude1 = 1.0d;
        double heightWGS84M1 = 2.0d;
        double speedMS1 = 4.0d;

        MissionRecord missionRecord = new MissionRecord(
                id.getId(),
                missionName,
                takeoffPointGroundHeight.getHeightM(),
                version.getVersion(),
                version.getVersion());

        WaypointRecord waypointRecord = new WaypointRecord(
                id.getId(),
                1,
                latitude1,
                longitude1,
                heightWGS84M1,
                speedMS1);

        when(missionMapper.findAll()).thenReturn(Arrays.asList(new MissionRecord[] {
                missionRecord,
                missionRecord,
                missionRecord
        }));
        when(waypointMapper.find(id.getId())).thenReturn(Arrays.asList(new WaypointRecord[] {waypointRecord}));

        List<Mission> missions = repository.getAll();

        Navigation expectNavigation = new Navigation();
        expectNavigation.setTakeoffPointGroundHeight(takeoffPointGroundHeight);
        expectNavigation.pushNextWaypoint(
            new GeodesicCoordinates(latitude1, longitude1),
            Height.distanceFrom(Height.fromM(heightWGS84M1), takeoffPointGroundHeight),
            Speed.fromMS(speedMS1));

        assertEquals(missions.size(), 3);

        assertEquals(missions.get(0).getId(), id);
        assertEquals(missions.get(0).getMissionName(), missionName);
        assertEquals(missions.get(0).getNavigation(), expectNavigation);
        assertEquals(missions.get(0).getVersion(), version);
        assertEquals(missions.get(0).getNewVersion(), version);
    }

    /**
     * リポジトリーからMissionエンティティをすべて取得する。<br>
     * エンティティが存在しない場合、空リストが返却されることを検証する。
     */
    @Test
    public void getAllNoMissionsTest() {
        List<Mission> missions = repository.getAll();

        assertEquals(missions.size(), 0);
    }

    /**
     * リポジトリーにMissionエンティティを一つ保存する。<br>
     * 既存のエンティティが無いため、新規登録されることを検証する。
     */
    @Test
    public void saveNewMissionTest() {
        MissionId id = MissionId.newId();
        String missionName = "mission name";
        Height takeoffPointGroundHeight = Height.fromM(10.0);

        double latitude1 = 0.0d;
        double longitude1 = 1.0d;
        double heightWGS84M1 = 2.0d;
        double speedMS1 = 4.0d;

        Mission mission = MissionFactory.newInstance(id);
        Version version = mission.getVersion();

        Navigation navigation = new Navigation();
        navigation.setTakeoffPointGroundHeight(takeoffPointGroundHeight);
        navigation.pushNextWaypoint(
            new GeodesicCoordinates(latitude1, longitude1),
            Height.distanceFrom(Height.fromM(heightWGS84M1), takeoffPointGroundHeight),
            Speed.fromMS(speedMS1));

        mission.nameMission(missionName);
        mission.replaceNavigationWith(navigation);

        Version newVersion = mission.getNewVersion();

        repository.save(mission);
        
        MissionRecord missionRecord = new MissionRecord(
                id.getId(),
                missionName,
                takeoffPointGroundHeight.getHeightM(),
                version.getVersion(),
                newVersion.getVersion());

        WaypointRecord waypointRecord = new WaypointRecord(
                id.getId(),
                1,
                latitude1,
                longitude1,
                heightWGS84M1,
                speedMS1);

        verify(missionMapper, times(1)).create(missionRecord);
        verify(waypointMapper, times(1)).create(waypointRecord);
    }

    /**
     * リポジトリーにMissionエンティティを一つ保存する。<br>
     * 既存のエンティティが存在するため、更新されることを検証する。
     */
    @Test
    public void savePreExistMissionTest() {
        MissionId id = MissionId.newId();
        String oldMissionName = "old mission";
        String newMissionName = "new mission";
        Height takeoffPointGroundHeight1 = Height.fromM(10.0);
        Height takeoffPointGroundHeight2 = Height.fromM(10.0);
        Version version = Version.newVersion();

        double latitude1 = 0.0d;
        double longitude1 = 1.0d;
        double heightWGS84M1 = 2.0d;
        double speedMS1 = 4.0d;
        double latitude2 = 10.0d;
        double longitude2 = 11.0d;
        double heightWGS84M2 = 12.0d;
        double speedMS2 = 14.0d;

        MissionRecord missionRecordBefore = new MissionRecord(
                id.getId(),
                oldMissionName,
                takeoffPointGroundHeight1.getHeightM(),
                version.getVersion(),
                version.getVersion());

        WaypointRecord waypointRecordBefore = new WaypointRecord(
                id.getId(),
                1,
                latitude1,
                longitude1,
                heightWGS84M1,
                speedMS1);

        when(missionMapper.find(id.getId())).thenReturn(missionRecordBefore);
        when(waypointMapper.find(id.getId())).thenReturn(Arrays.asList(new WaypointRecord[] {waypointRecordBefore}));

        Mission mission = repository.getById(id);

        Navigation navigation = new Navigation();
        navigation.setTakeoffPointGroundHeight(takeoffPointGroundHeight2);
        navigation.pushNextWaypoint(
            new GeodesicCoordinates(latitude2, longitude2),
            Height.distanceFrom(Height.fromM(heightWGS84M2), takeoffPointGroundHeight2),
            Speed.fromMS(speedMS2));

        mission.nameMission(newMissionName);
        mission.replaceNavigationWith(navigation);

        Version newVersion = mission.getNewVersion();

        repository.save(mission);
        
        MissionRecord missionRecordAfter = new MissionRecord(
                id.getId(),
                newMissionName,
                takeoffPointGroundHeight2.getHeightM(),
                version.getVersion(),
                newVersion.getVersion());

        WaypointRecord waypointRecordAfter = new WaypointRecord(
                id.getId(),
                1,
                latitude2,
                longitude2,
                heightWGS84M2,
                speedMS2);

        verify(missionMapper, times(1)).update(missionRecordAfter);
        verify(waypointMapper, times(1)).delete(id.getId());
        verify(waypointMapper, times(1)).create(waypointRecordAfter);
    }

    /**
     * リポジトリーからMissionエンティティを一つ削除する。
     */
    @Test
    public void removeMissionTest() {
        MissionId id = MissionId.newId();
        Version version = Version.newVersion();

        repository.remove(id, version);

        DeleteCondition condition = new DeleteCondition();
        condition.setId(id.getId());
        condition.setVersion(version.getVersion());

        verify(missionMapper, times(1)).delete(condition);
        verify(waypointMapper, times(1)).delete(id.getId());
    }
}