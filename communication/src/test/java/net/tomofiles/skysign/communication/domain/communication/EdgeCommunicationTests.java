package net.tomofiles.skysign.communication.domain.communication;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.Mock;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNull;
import static org.mockito.Mockito.when;
import static org.mockito.MockitoAnnotations.initMocks;

import java.time.LocalDateTime;
import java.util.List;

public class EdgeCommunicationTests {
    
    @Mock
    private CommunicationRepository repository;

    @BeforeEach
    public void beforeEach() {
        initMocks(this);
    }

    /**
     * Edgeが、既存のCommunicationエンティティのTelemetryを更新する。<br>
     * すべてのTelemetryのフィールドが更新されることを検証する。
     */
    @Test
    public void pushTelemetryToCommunicationTest() {
        CommunicationId id = CommunicationId.newId();
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

        Communication before = CommunicationFactory.newInstance(id);

        when(repository.getById(id)).thenReturn(before);

        Communication communication = repository.getById(id);

        communication.pushTelemetry(
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
                orientationW);

        Telemetry after = Telemetry.newInstance()
                .setPosition(latitude, longitude, altitude, relativeAltitude, speed)
                .setArmed(armed)
                .setFlightMode(flightMode)
                .setOrientation(orientationX, orientationY, orientationZ, orientationW);

        assertEquals(communication.getTelemetry(), after);
    }

    /**
     * Edgeが、既存のCommunicationエンティティからCommandIDリストを取得する。<br>
     * CommandIDはCommandをEdgeから古い順でCloudに取得しに来るため、<br>
     * CommandのTimeの昇順でソートされていることを検証する。
     */
    @Test
    public void pullCommandIDsFromCommunicationTest() {
        CommunicationId id = CommunicationId.newId();
        CommandId commandId1 = CommandId.newId();
        CommandId commandId2 = CommandId.newId();
        CommandId commandId3 = CommandId.newId();
        CommandId commandId4 = CommandId.newId();

        Communication before = CommunicationFactory.newInstance(id);
        before.getCommands().add(new Command(
                commandId1,
                CommandType.ARM,
                LocalDateTime.of(2020, 07, 22, 10, 30, 25))); // 3
        before.getCommands().add(new Command(
                commandId2,
                CommandType.ARM,
                LocalDateTime.of(2020, 07, 22, 10, 30, 12))); // 2
        before.getCommands().add(new Command(
                commandId3,
                CommandType.ARM,
                LocalDateTime.of(2020, 07, 22, 10, 30, 59))); // 4
        before.getCommands().add(new Command(
                commandId4,
                CommandType.ARM,
                LocalDateTime.of(2020, 07, 22, 10, 30, 2))); // 1

        when(repository.getById(id)).thenReturn(before);

        Communication communication = repository.getById(id);

        List<CommandId> commandIds = communication.getCommandIds();

        assertEquals(commandIds.get(0), commandId4);
        assertEquals(commandIds.get(1), commandId2);
        assertEquals(commandIds.get(2), commandId1);
        assertEquals(commandIds.get(3), commandId3);
    }

    /**
     * Edgeが、既存のCommunicationエンティティからCommandを取得する。<br>
     * CommandIDに合致するCommandが返却され、Communicationエンティティから<br>
     * Commandが削除されることを検証する。
     */
    @Test
    public void pullCommandFromCommunicationTest() {
        CommunicationId id = CommunicationId.newId();
        CommandId commandId = CommandId.newId();

        Communication before = CommunicationFactory.newInstance(id);
        before.getCommands().add(new Command(commandId, CommandType.ARM, LocalDateTime.now()));

        when(repository.getById(id)).thenReturn(before);

        Communication communication = repository.getById(id);

        CommandType type = communication.pullCommandById(commandId);

        assertEquals(type, CommandType.ARM);
        assertEquals(communication.getCommandIds().size(), 0);

        assertEquals(communication.getId(), id);
        assertNull(communication.getMissionId());
        assertEquals(communication.getCommands().size(), 0);
        assertEquals(communication.getTelemetry(), Telemetry.newInstance());
    }

    /**
     * Edgeが、既存のCommunicationエンティティからCommandを取得する。<br>
     * Commandはまだ発行されていないので、NULLが返却されることを検証する。
     */
    @Test
    public void pullNoCommandFromCommunicationTest() {
        CommunicationId id = CommunicationId.newId();

        when(repository.getById(id)).thenReturn(CommunicationFactory.newInstance(id));

        Communication communication = repository.getById(id);

        CommandType type = communication.pullCommandById(CommandId.newId());

        assertNull(type);
        assertEquals(communication.getCommandIds().size(), 0);

        assertEquals(communication.getId(), id);
        assertNull(communication.getMissionId());
        assertEquals(communication.getCommands().size(), 0);
        assertEquals(communication.getTelemetry(), Telemetry.newInstance());
    }
}