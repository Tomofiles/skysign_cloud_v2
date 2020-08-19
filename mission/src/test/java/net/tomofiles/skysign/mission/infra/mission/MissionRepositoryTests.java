package net.tomofiles.skysign.mission.infra.mission;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.InjectMocks;
import org.mockito.Mock;

import static com.google.common.truth.Truth.assertThat;

import net.tomofiles.skysign.mission.domain.mission.Generator;
import net.tomofiles.skysign.mission.domain.mission.Mission;
import net.tomofiles.skysign.mission.domain.mission.MissionId;
import net.tomofiles.skysign.mission.domain.mission.Version;
import net.tomofiles.skysign.mission.infra.common.DeleteCondition;

import static org.junit.jupiter.api.Assertions.assertAll;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;
import static org.mockito.MockitoAnnotations.initMocks;

import java.util.Arrays;
import java.util.List;
import java.util.UUID;
import java.util.function.Supplier;

import static net.tomofiles.skysign.mission.domain.mission.MissionObjectMother.newSingleNavigationMission;
import static net.tomofiles.skysign.mission.domain.mission.MissionObjectMother.newSeveralNavigationMission;
import static net.tomofiles.skysign.mission.infra.mission.RecordObjectMother.newNormalMissionRecord;
import static net.tomofiles.skysign.mission.infra.mission.RecordObjectMother.newSingleWaypointRecord;
import static net.tomofiles.skysign.mission.infra.mission.RecordObjectMother.newSeveralWaypointRecords;
import static net.tomofiles.skysign.mission.infra.mission.RecordObjectMother.newSeveralInRondomOrderWaypointRecords;

public class MissionRepositoryTests {

    private static final MissionId DEFAULT_MISSION_ID = new MissionId(UUID.randomUUID().toString());
    private static final Version DEFAULT_VERSION = new Version(UUID.randomUUID().toString());
    private static final Supplier<Generator> DEFAULT_GENERATOR = () -> {
        return new Generator(){
            @Override
            public MissionId newMissionId() {
                return DEFAULT_MISSION_ID;
            }

            @Override
            public Version newVersion() {
                return DEFAULT_VERSION;
            }
        };
    };
    
    @Mock
    private MissionMapper missionMapper;

    @Mock
    private WaypointMapper waypointMapper;

    @InjectMocks
    private MissionRepositoryImpl repository;

    @BeforeEach
    public void beforeEach() {
        initMocks(this);
    }

    /**
     * リポジトリーからMissionエンティティを一つ取得する。<br>
     * Waypointの順序がバラバラでも、リポジトリーがソートして返却することを検証する。
     */
    @Test
    public void getMissionByIdTest() {
        when(missionMapper.find(DEFAULT_MISSION_ID.getId()))
                .thenReturn(newNormalMissionRecord(DEFAULT_MISSION_ID, DEFAULT_VERSION));

        when(waypointMapper.find(DEFAULT_MISSION_ID.getId()))
                .thenReturn(newSeveralInRondomOrderWaypointRecords(DEFAULT_MISSION_ID));

        Mission mission = repository.getById(DEFAULT_MISSION_ID);

        Mission expectedMission = newSeveralNavigationMission(
                DEFAULT_MISSION_ID,
                DEFAULT_VERSION,
                DEFAULT_GENERATOR.get());

        assertAll(
            () -> assertThat(mission.getId()).isEqualTo(expectedMission.getId()),
            () -> assertThat(mission.getMissionName()).isEqualTo(expectedMission.getMissionName()),
            () -> assertThat(mission.getNavigation()).isEqualTo(expectedMission.getNavigation()),
            () -> assertThat(mission.getVersion()).isEqualTo(expectedMission.getVersion()),
            () -> assertThat(mission.getNewVersion()).isEqualTo(expectedMission.getNewVersion())
        );
    }

    /**
     * リポジトリーからMissionエンティティを一つ取得する。<br>
     * 対象のエンティティが存在しない場合、NULLが返却されることを検証する。
     */
    @Test
    public void getNoMissionByIdTest() {
        Mission mission = repository.getById(DEFAULT_MISSION_ID);

        assertThat(mission).isNull();
    }

    /**
     * リポジトリーからMissionエンティティをすべて取得する。
     */
    @Test
    public void getAllMissionsTest() {
        when(missionMapper.findAll())
                .thenReturn(Arrays.asList(new MissionRecord[] {
                        newNormalMissionRecord(DEFAULT_MISSION_ID, DEFAULT_VERSION),
                        newNormalMissionRecord(DEFAULT_MISSION_ID, DEFAULT_VERSION),
                        newNormalMissionRecord(DEFAULT_MISSION_ID, DEFAULT_VERSION)
                }));
        when(waypointMapper.find(DEFAULT_MISSION_ID.getId()))
                .thenReturn(
                        newSeveralWaypointRecords(DEFAULT_MISSION_ID)
                );

        List<Mission> missions = repository.getAll();

        Mission expectedMission = newSeveralNavigationMission(
                DEFAULT_MISSION_ID,
                DEFAULT_VERSION,
                DEFAULT_GENERATOR.get());

        assertAll(
            () -> assertThat(missions).hasSize(3),
            () -> assertThat(missions.get(0).getId()).isEqualTo(expectedMission.getId()),
            () -> assertThat(missions.get(0).getMissionName()).isEqualTo(expectedMission.getMissionName()),
            () -> assertThat(missions.get(0).getNavigation()).isEqualTo(expectedMission.getNavigation()),
            () -> assertThat(missions.get(0).getVersion()).isEqualTo(expectedMission.getVersion()),
            () -> assertThat(missions.get(0).getNewVersion()).isEqualTo(expectedMission.getNewVersion())
        );
    }

    /**
     * リポジトリーからMissionエンティティをすべて取得する。<br>
     * エンティティが存在しない場合、空リストが返却されることを検証する。
     */
    @Test
    public void getAllNoMissionsTest() {
        List<Mission> missions = repository.getAll();

        assertThat(missions).hasSize(0);
    }

    /**
     * リポジトリーにMissionエンティティを一つ保存する。<br>
     * 既存のエンティティが無いため、新規登録されることを検証する。
     */
    @Test
    public void saveNewMissionTest() {
        repository.save(newSingleNavigationMission(
                DEFAULT_MISSION_ID,
                DEFAULT_VERSION,
                DEFAULT_GENERATOR.get()));

        verify(missionMapper, times(1)).create(newNormalMissionRecord(DEFAULT_MISSION_ID, DEFAULT_VERSION));
        verify(waypointMapper, times(1)).create(newSingleWaypointRecord(DEFAULT_MISSION_ID));
    }

    /**
     * リポジトリーにMissionエンティティを一つ保存する。<br>
     * 既存のエンティティが存在するため、更新されることを検証する。
     */
    @Test
    public void savePreExistMissionTest() {
        when(missionMapper.find(DEFAULT_MISSION_ID.getId()))
                .thenReturn(newNormalMissionRecord(DEFAULT_MISSION_ID, DEFAULT_VERSION));
        when(waypointMapper.find(DEFAULT_MISSION_ID.getId()))
                .thenReturn(Arrays.asList(new WaypointRecord[] {
                        newSingleWaypointRecord(DEFAULT_MISSION_ID)
                }));

        Mission mission = repository.getById(DEFAULT_MISSION_ID);

        repository.save(mission);
        
        verify(missionMapper, times(1)).update(newNormalMissionRecord(DEFAULT_MISSION_ID, DEFAULT_VERSION));
        verify(waypointMapper, times(1)).delete(DEFAULT_MISSION_ID.getId());
        verify(waypointMapper, times(1)).create(newSingleWaypointRecord(DEFAULT_MISSION_ID));
    }

    /**
     * リポジトリーからMissionエンティティを一つ削除する。
     */
    @Test
    public void removeMissionTest() {
        repository.remove(DEFAULT_MISSION_ID, DEFAULT_VERSION);

        DeleteCondition condition = new DeleteCondition();
        condition.setId(DEFAULT_MISSION_ID.getId());
        condition.setVersion(DEFAULT_VERSION.getVersion());

        verify(missionMapper, times(1)).delete(condition);
        verify(waypointMapper, times(1)).delete(DEFAULT_MISSION_ID.getId());
    }
}