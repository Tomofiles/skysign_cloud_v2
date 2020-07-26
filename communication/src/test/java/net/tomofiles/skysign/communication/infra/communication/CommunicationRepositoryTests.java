package net.tomofiles.skysign.communication.infra.communication;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.InjectMocks;
import org.mockito.Mock;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNull;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;
import static org.mockito.MockitoAnnotations.initMocks;

import java.time.LocalDateTime;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

import net.tomofiles.skysign.communication.domain.communication.CommandId;
import net.tomofiles.skysign.communication.domain.communication.CommandType;
import net.tomofiles.skysign.communication.domain.communication.Communication;
import net.tomofiles.skysign.communication.domain.communication.CommunicationFactory;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.domain.communication.CommunicationRepository;
import net.tomofiles.skysign.communication.domain.communication.MissionId;
import net.tomofiles.skysign.communication.domain.communication.component.CommandComponentDto;
import net.tomofiles.skysign.communication.domain.communication.component.CommunicationComponentDto;
import net.tomofiles.skysign.communication.domain.communication.component.TelemetryComponentDto;


public class CommunicationRepositoryTests {
    
    @Mock
    private CommunicationMapper communicationMapper;

    @Mock
    private TelemetryMapper telemetryMapper;

    @Mock
    private CommandMapper commandMapper;

    @InjectMocks
    private CommunicationRepository repository = new CommunicationRepositoryImpl();

    @BeforeEach
    public void beforeEach() {
        initMocks(this);
    }

    /**
     * リポジトリーからCommunicationエンティティを一つ取得する。
     */
    @Test
    public void getCommunicationByIdTest() {
        CommunicationId id = CommunicationId.newId();
        MissionId missionId = new MissionId("mission id");
        double latitude = 0.0d;
        double longitude = 1.0d;
        double altitude = 2.0d;
        double relativeAltitude = 3.0d;
        double speed = 4.0d;
        boolean armed = true;
        String flightMode = "INFLIGHT";
        double orientationX = 5.0d;
        double orientationY = 6.0d;
        double orientationZ = 7.0d;
        double orientationW = 8.0d;
        CommandId commandId1 = CommandId.newId();
        CommandType type1 = CommandType.ARM;
        LocalDateTime time1 = LocalDateTime.of(2020, 07, 22, 10, 30, 25);
        CommandId commandId2 = CommandId.newId();
        CommandType type2 = CommandType.DISARM;
        LocalDateTime time2 = LocalDateTime.of(2020, 07, 22, 10, 30, 30);
        CommandId commandId3 = CommandId.newId();
        CommandType type3 = CommandType.UPLOAD;
        LocalDateTime time3 = LocalDateTime.of(2020, 07, 22, 10, 30, 45);


        CommunicationRecord record = new CommunicationRecord();
        record.setId(id.getId());
        record.setMissionId(missionId.getId());

        when(communicationMapper.find(id.getId())).thenReturn(record);

        TelemetryRecord telemRecord = new TelemetryRecord();
        telemRecord.setCommId(id.getId());
        telemRecord.setLatitude(latitude);
        telemRecord.setLongitude(longitude);
        telemRecord.setAltitude(altitude);
        telemRecord.setRelativeAltitude(relativeAltitude);
        telemRecord.setSpeed(speed);
        telemRecord.setArmed(armed);
        telemRecord.setFlightMode(flightMode);
        telemRecord.setOriX(orientationX);
        telemRecord.setOriY(orientationY);
        telemRecord.setOriZ(orientationZ);
        telemRecord.setOriW(orientationW);

        when(telemetryMapper.find(id.getId())).thenReturn(telemRecord);

        CommandRecord commRecord1 = new CommandRecord();
        commRecord1.setId(commandId1.getId());
        commRecord1.setCommId(id.getId());
        commRecord1.setType(type1.name());
        commRecord1.setTime(time1);
        CommandRecord commRecord2 = new CommandRecord();
        commRecord2.setId(commandId2.getId());
        commRecord2.setCommId(id.getId());
        commRecord2.setType(type2.name());
        commRecord2.setTime(time2);
        CommandRecord commRecord3 = new CommandRecord();
        commRecord3.setId(commandId3.getId());
        commRecord3.setCommId(id.getId());
        commRecord3.setType(type3.name());
        commRecord3.setTime(time3);

        List<CommandRecord> commRecords = new ArrayList<>();
        commRecords.add(commRecord1);
        commRecords.add(commRecord2);
        commRecords.add(commRecord3);

        when(commandMapper.findByCommId(id.getId())).thenReturn(commRecords);

        Communication communication = repository.getById(id);
        
        CommunicationComponentDto dto = CommunicationFactory.takeApart(communication);

        assertEquals(dto.getId(), id.getId());
        assertEquals(dto.getMissionId(), missionId.getId());
        assertEquals(dto.getCommands().get(0).getId(), commandId1.getId());
        assertEquals(dto.getCommands().get(0).getType(), type1.name());
        assertEquals(dto.getCommands().get(0).getTime(), time1);
        assertEquals(dto.getCommands().get(1).getId(), commandId2.getId());
        assertEquals(dto.getCommands().get(1).getType(), type2.name());
        assertEquals(dto.getCommands().get(1).getTime(), time2);
        assertEquals(dto.getCommands().get(2).getId(), commandId3.getId());
        assertEquals(dto.getCommands().get(2).getType(), type3.name());
        assertEquals(dto.getCommands().get(2).getTime(), time3);
        assertEquals(dto.getTelemetry().getLatitude(), latitude);
        assertEquals(dto.getTelemetry().getLongitude(), longitude);
        assertEquals(dto.getTelemetry().getAltitude(), altitude);
        assertEquals(dto.getTelemetry().getRelativeAltitude(), relativeAltitude);
        assertEquals(dto.getTelemetry().getSpeed(), speed);
        assertEquals(dto.getTelemetry().isArmed(), armed);
        assertEquals(dto.getTelemetry().getFlightMode(), flightMode);
        assertEquals(dto.getTelemetry().getOriX(), orientationX);
        assertEquals(dto.getTelemetry().getOriY(), orientationY);
        assertEquals(dto.getTelemetry().getOriZ(), orientationZ);
        assertEquals(dto.getTelemetry().getOriW(), orientationW);
    }

    /**
     * リポジトリーからCommunicationエンティティを一つ取得する。<br>
     * 対象のエンティティが存在しない場合、NULLが返却されることを検証する。
     */
    @Test
    public void getNoCommunicationByIdTest() {
        CommunicationId id = CommunicationId.newId();

        Communication communication = repository.getById(id);

        assertNull(communication);
    }

    /**
     * リポジトリーにCommunicationエンティティを一つ保存する。<br>
     * 既存のエンティティが無いため、新規登録されることを検証する。
     */
    @Test
    public void saveNewCommunicationTest() {
        CommunicationId id = CommunicationId.newId();
        MissionId missionId = new MissionId("new mission id");

        double latitude = 0.0d;
        double longitude = 1.0d;
        double altitude = 2.0d;
        double relativeAltitude = 3.0d;
        double speed = 4.0d;
        boolean armed = true;
        String flightMode = "INFLIGHT";
        double orientationX = 5.0d;
        double orientationY = 6.0d;
        double orientationZ = 7.0d;
        double orientationW = 8.0d;

        CommandId commandId1 = CommandId.newId();
        CommandType type1 = CommandType.ARM;
        LocalDateTime time1 = LocalDateTime.of(2020, 07, 22, 10, 30, 25);
        CommandId commandId2 = CommandId.newId();
        CommandType type2 = CommandType.DISARM;
        LocalDateTime time2 = LocalDateTime.of(2020, 07, 22, 10, 30, 30);

        Communication communication = CommunicationFactory.assembleFrom(
                new CommunicationComponentDto(
                        id.getId(),
                        missionId.getId(),
                        new TelemetryComponentDto(
                                latitude,
                                longitude,
                                altitude,
                                relativeAltitude,
                                speed,
                                armed,
                                flightMode,
                                orientationX,
                                orientationY,
                                orientationZ,
                                orientationW),
                        Arrays.asList(new CommandComponentDto[] {
                                new CommandComponentDto(
                                        commandId1.getId(),
                                        type1.name(),
                                        time1)
                        })
                )
        );

        CommunicationRecord before = new CommunicationRecord();
        before.setId(id.getId());

        when(communicationMapper.find(id.getId())).thenReturn(before);

        TelemetryRecord telemBefore = new TelemetryRecord();
        telemBefore.setCommId(id.getId());

        when(telemetryMapper.find(id.getId())).thenReturn(telemBefore);

        List<CommandRecord> commBefores = new ArrayList<>();
        CommandRecord commBefore2 = new CommandRecord();
        commBefore2.setId(commandId2.getId());
        commBefore2.setCommId(id.getId());
        commBefore2.setType(type2.name());
        commBefore2.setTime(time2);

        commBefores.add(commBefore2);

        when(commandMapper.findByCommId(id.getId())).thenReturn(commBefores);

        repository.save(communication);

        CommunicationRecord after = new CommunicationRecord();
        after.setId(id.getId());
        after.setMissionId(missionId.getId());

        TelemetryRecord telemAfter = new TelemetryRecord();
        telemAfter.setCommId(id.getId());
        telemAfter.setLatitude(latitude);
        telemAfter.setLongitude(longitude);
        telemAfter.setAltitude(altitude);
        telemAfter.setRelativeAltitude(relativeAltitude);
        telemAfter.setSpeed(speed);
        telemAfter.setArmed(armed);
        telemAfter.setFlightMode(flightMode);
        telemAfter.setOriX(orientationX);
        telemAfter.setOriY(orientationY);
        telemAfter.setOriZ(orientationZ);
        telemAfter.setOriW(orientationW);

        CommandRecord commRecord1 = new CommandRecord();
        commRecord1.setId(commandId1.getId());
        commRecord1.setCommId(id.getId());
        commRecord1.setType(type1.name());
        commRecord1.setTime(time1);

        verify(communicationMapper, times(1)).update(after);
        verify(telemetryMapper, times(1)).update(telemAfter);
        verify(commandMapper, times(1)).create(commRecord1);
        verify(commandMapper, times(1)).delete(commandId2.getId());
    }

    /**
     * リポジトリーにCommunicationエンティティを一つ保存する。<br>
     * 既存のエンティティが存在するため、更新されることを検証する。
     */
    @Test
    public void savePreExistCommunicationTest() {
        CommunicationId id = CommunicationId.newId();
        MissionId missionId = new MissionId("new mission id");

        double latitude = 0.0d;
        double longitude = 1.0d;
        double altitude = 2.0d;
        double relativeAltitude = 3.0d;
        double speed = 4.0d;
        boolean armed = true;
        String flightMode = "INFLIGHT";
        double orientationX = 5.0d;
        double orientationY = 6.0d;
        double orientationZ = 7.0d;
        double orientationW = 8.0d;

        CommandId commandId1 = CommandId.newId();
        CommandType type1 = CommandType.ARM;
        LocalDateTime time1 = LocalDateTime.of(2020, 07, 22, 10, 30, 25);

        Communication communication = CommunicationFactory.assembleFrom(
                new CommunicationComponentDto(
                        id.getId(),
                        missionId.getId(),
                        new TelemetryComponentDto(
                                latitude,
                                longitude,
                                altitude,
                                relativeAltitude,
                                speed,
                                armed,
                                flightMode,
                                orientationX,
                                orientationY,
                                orientationZ,
                                orientationW),
                        Arrays.asList(new CommandComponentDto[] {
                                new CommandComponentDto(
                                        commandId1.getId(),
                                        type1.name(),
                                        time1)
                        })
                )
        );

        repository.save(communication);

        CommunicationRecord record = new CommunicationRecord();
        record.setId(id.getId());
        record.setMissionId(missionId.getId());

        TelemetryRecord telemRecord = new TelemetryRecord();
        telemRecord.setCommId(id.getId());
        telemRecord.setLatitude(latitude);
        telemRecord.setLongitude(longitude);
        telemRecord.setAltitude(altitude);
        telemRecord.setRelativeAltitude(relativeAltitude);
        telemRecord.setSpeed(speed);
        telemRecord.setArmed(armed);
        telemRecord.setFlightMode(flightMode);
        telemRecord.setOriX(orientationX);
        telemRecord.setOriY(orientationY);
        telemRecord.setOriZ(orientationZ);
        telemRecord.setOriW(orientationW);

        CommandRecord commRecord1 = new CommandRecord();
        commRecord1.setId(commandId1.getId());
        commRecord1.setCommId(id.getId());
        commRecord1.setType(type1.name());
        commRecord1.setTime(time1);

        verify(communicationMapper, times(1)).create(record);
        verify(telemetryMapper, times(1)).create(telemRecord);
        verify(commandMapper, times(1)).create(commRecord1);
    }

    /**
     * リポジトリーからCommunicationエンティティを一つ削除する。
     */
    @Test
    public void removeCommunicationTest() {
        CommunicationId id = CommunicationId.newId();

        repository.remove(id);

        verify(communicationMapper, times(1)).delete(id.getId());
    }
}