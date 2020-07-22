package net.tomofiles.skysign.communication.domain.communication;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.Mock;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertNull;
import static org.mockito.Mockito.when;
import static org.mockito.MockitoAnnotations.initMocks;

import net.tomofiles.skysign.communication.domain.common.Version;

public class UserCommunicationTests {
    
    @Mock
    private CommunicationRepository repository;

    @BeforeEach
    public void beforeEach() {
        initMocks(this);
    }

    /**
     * Userが、新しいCommunicationエンティティを作成する。<br>
     * Communicationエンティティの初期状態を検証する。
     */
    @Test
    public void createNewCommunicationTest() {
        CommunicationId id = CommunicationId.newId();

        Communication communication = CommunicationFactory.newInstance(id);

        assertEquals(communication.getId(), id);
        assertNull(communication.getMissionId());
        assertEquals(communication.getCommands().size(), 0);
        assertEquals(communication.getTelemetry(), Telemetry.newInstance());
        assertEquals(communication.getVersion(), new Version(1));
    }

    /**
     * Userが、既存のCommunicationエンティティにCommandを追加する。<br>
     * Commandが追加され、IDとTimeが付与されていることを検証する。
     */
    @Test
    public void pushCommandToCommunicationTest() {
        CommunicationId id = CommunicationId.newId();

        when(repository.getById(id)).thenReturn(CommunicationFactory.newInstance(id));

        Communication communication = repository.getById(id);

        communication.pushCommand(CommandType.ARM);

        assertEquals(communication.getCommandIds().size(), 1);

        assertEquals(communication.getId(), id);
        assertNull(communication.getMissionId());
        assertEquals(communication.getCommands().size(), 1);
        assertNotNull(communication.getCommands().get(0).getId());
        assertEquals(communication.getCommands().get(0).getType(), CommandType.ARM);
        assertNotNull(communication.getCommands().get(0).getTime());
        assertEquals(communication.getTelemetry(), Telemetry.newInstance());
        assertEquals(communication.getVersion(), new Version(1));
    }

    /**
     * Userが、既存のCommunicationエンティティからTelemetryを取得する。<br>
     * Telemetryのスナップショットが生成され、返却されることを検証する。
     */
    @Test
    public void pullTelemetryFromCommunicationTest() {
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
        before.setTelemetry(Telemetry.newInstance()
                .setPosition(latitude, longitude, altitude, relativeAltitude, speed)
                .setArmed(armed)
                .setFlightMode(flightMode)
                .setOrientation(orientationX, orientationY, orientationZ, orientationW)
        );

        when(repository.getById(id)).thenReturn(before);

        Communication communication = repository.getById(id);

        TelemetrySnapshot telemetry = communication.pullTelemetry();

        assertEquals(telemetry.getLatitude(), latitude);
        assertEquals(telemetry.getLongitude(), longitude);
        assertEquals(telemetry.getAltitude(), altitude);
        assertEquals(telemetry.getRelativeAltitude(), relativeAltitude);
        assertEquals(telemetry.getSpeed(), speed);
        assertEquals(telemetry.isArmed(), armed);
        assertEquals(telemetry.getFlightMode(), flightMode);
        assertEquals(telemetry.getX(), orientationX);
        assertEquals(telemetry.getY(), orientationY);
        assertEquals(telemetry.getZ(), orientationZ);
        assertEquals(telemetry.getW(), orientationW);
    }
}