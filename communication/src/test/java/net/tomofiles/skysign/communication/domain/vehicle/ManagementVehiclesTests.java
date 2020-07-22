package net.tomofiles.skysign.communication.domain.vehicle;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.Mock;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNull;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;
import static org.mockito.MockitoAnnotations.initMocks;

import net.tomofiles.skysign.communication.domain.common.Version;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.event.Publisher;

public class ManagementVehiclesTests {
    
    @Mock
    private Publisher publisher;

    @Mock
    private VehicleRepository repository;

    @BeforeEach
    public void beforeEach() {
        initMocks(this);
    }

    /**
     * Userが、新しいVehicleエンティティを作成する。<br>
     * Vehicleエンティティの初期状態を検証する。
     */
    @Test
    public void createNewVehicleTest() {
        VehicleId id = VehicleId.newId();

        Vehicle vehicle = VehicleFactory.newInstance(id);

        assertEquals(vehicle.getId(), id);
        assertNull(vehicle.getVehicleName());
        assertNull(vehicle.getCommId());
        assertEquals(vehicle.getVersion(), new Version(1));
    }

    /**
     * Userが、既存のVehicleエンティティに対してVehicle Nameを更新する。<br>
     * Vehicle Name以外の変化が無いことを検証する。
     */
    @Test
    public void changeVehiclesNameTest() {
        VehicleId id = VehicleId.newId();

        String oldVehicleName = "old vehicle";
        CommunicationId oldCommId = CommunicationId.newId();
        Version version = new Version(1);

        Vehicle before = new Vehicle(id);
        before.setVehicleName(oldVehicleName);
        before.setCommId(oldCommId);
        before.setVersion(version);

        when(repository.getById(id)).thenReturn(before);

        Vehicle vehicle = repository.getById(id);

        String newVehicleName = "new vehicle";
        vehicle.nameVehicle(newVehicleName);

        assertEquals(vehicle.getId(), id);
        assertEquals(vehicle.getVehicleName(), newVehicleName);
        assertEquals(vehicle.getCommId(), oldCommId);
        assertEquals(vehicle.getVersion(), version);
    }

    /**
     * Userが、新しいVehicleエンティティを作成して、CommunicationIDを付与する。<br>
     * CommunicationIDを付与したら、イベントを生成して発行する。<br>
     * その際、新しいCommunicationIDのみ、購読者に通知されることを検証する。
     */
    @Test
    public void changeNewVehiclesCommIdAndPublishEventTest() {
        VehicleId id = VehicleId.newId();

        CommunicationId newCommId = CommunicationId.newId();

        Vehicle vehicle = VehicleFactory.newInstance(id);

        vehicle.setPublisher(publisher);

        vehicle.giveCommId(newCommId);

        CommunicationIdChangedEvent event
                = new CommunicationIdChangedEvent(
                    null,
                    newCommId
                );

        assertEquals(vehicle.getId(), id);
        assertNull(vehicle.getVehicleName());
        assertEquals(vehicle.getCommId(), newCommId);
        assertEquals(vehicle.getVersion(), new Version(1));

        verify(publisher, times(1)).publish(event);
    }

    /**
     * Userが、既存のVehicleエンティティに対してCommunicationIDを更新する。<br>
     * VehicleエンティティにCommunicationIDを付与することで、<br>
     * イベントを生成して発行する。<br>
     * その際、古いCommunicationIDと新しいCommunicationIDの両方が、<br>
     * 購読者に通知されることを検証する。
     */
    @Test
    public void changePreExistVehiclesCommIdAndPublishEventTest() {
        VehicleId id = VehicleId.newId();

        String oldVehicleName = "old vehicle";
        CommunicationId oldCommId = CommunicationId.newId();
        Version version = new Version(1);

        Vehicle before = new Vehicle(id);
        before.setVehicleName(oldVehicleName);
        before.setCommId(oldCommId);
        before.setVersion(version);

        when(repository.getById(id)).thenReturn(before);

        Vehicle vehicle = repository.getById(id);

        vehicle.setPublisher(publisher);

        CommunicationId newCommId = CommunicationId.newId();
        vehicle.giveCommId(newCommId);

        CommunicationIdChangedEvent event
                = new CommunicationIdChangedEvent(
                    oldCommId,
                    newCommId
                );

        assertEquals(vehicle.getId(), id);
        assertEquals(vehicle.getVehicleName(), oldVehicleName);
        assertEquals(vehicle.getCommId(), newCommId);
        assertEquals(vehicle.getVersion(), version);

        verify(publisher, times(1)).publish(event);
    }
}