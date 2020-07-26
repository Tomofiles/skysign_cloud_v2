package net.tomofiles.skysign.communication.infra.vehicle;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.InjectMocks;
import org.mockito.Mock;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNull;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;
import static org.mockito.MockitoAnnotations.initMocks;

import java.util.ArrayList;
import java.util.List;

import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.domain.vehicle.Vehicle;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleFactory;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleId;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleRepository;
import net.tomofiles.skysign.communication.domain.vehicle.Version;
import net.tomofiles.skysign.communication.infra.common.DeleteCondition;

public class VehicleRepositoryTests {
    
    @Mock
    private VehicleMapper vehicleMapper;

    @InjectMocks
    private VehicleRepository repository = new VehicleRepositoryImpl();

    @BeforeEach
    public void beforeEach() {
        initMocks(this);
    }

    /**
     * リポジトリーからVehicleエンティティを一つ取得する。
     */
    @Test
    public void getVehicleByIdTest() {
        VehicleId id = VehicleId.newId();
        String oldVehicleName = "old vehicle";
        CommunicationId oldCommId = CommunicationId.newId();
        Version version = Version.newVersion();

        VehicleRecord record = new VehicleRecord();
        record.setId(id.getId());
        record.setName(oldVehicleName);
        record.setCommId(oldCommId.getId());
        record.setVersion(version.getVersion());

        when(vehicleMapper.find(id.getId())).thenReturn(record);

        Vehicle vehicle = repository.getById(id);

        assertEquals(vehicle.getId(), id);
        assertEquals(vehicle.getVehicleName(), oldVehicleName);
        assertEquals(vehicle.getCommId(), oldCommId);
        assertEquals(vehicle.getVersion(), version);
        assertEquals(vehicle.getNewVersion(), version);
    }

    /**
     * リポジトリーからVehicleエンティティを一つ取得する。<br>
     * 対象のエンティティが存在しない場合、NULLが返却されることを検証する。
     */
    @Test
    public void getNoVehicleByIdTest() {
        VehicleId id = VehicleId.newId();

        Vehicle vehicle = repository.getById(id);

        assertNull(vehicle);
    }

    /**
     * リポジトリーからVehicleエンティティをすべて取得する。
     */
    @Test
    public void getAllVehiclesTest() {
        VehicleId id = VehicleId.newId();
        String oldVehicleName = "old vehicle";
        CommunicationId oldCommId = CommunicationId.newId();
        Version version = Version.newVersion();

        VehicleRecord record = new VehicleRecord();
        record.setId(id.getId());
        record.setName(oldVehicleName);
        record.setCommId(oldCommId.getId());
        record.setVersion(version.getVersion());

        List<VehicleRecord> records = new ArrayList<>();
        records.add(record);
        records.add(record);
        records.add(record);

        when(vehicleMapper.findAll()).thenReturn(records);

        List<Vehicle> vehicles = repository.getAll();

        assertEquals(vehicles.size(), 3);

        assertEquals(vehicles.get(0).getId(), id);
        assertEquals(vehicles.get(0).getVehicleName(), oldVehicleName);
        assertEquals(vehicles.get(0).getCommId(), oldCommId);
        assertEquals(vehicles.get(0).getVersion(), version);
        assertEquals(vehicles.get(0).getNewVersion(), version);
    }

    /**
     * リポジトリーからVehicleエンティティをすべて取得する。<br>
     * エンティティが存在しない場合、空リストが返却されることを検証する。
     */
    @Test
    public void getAllNoVehiclesTest() {
        List<Vehicle> vehicles = repository.getAll();

        assertEquals(vehicles.size(), 0);
    }

    /**
     * リポジトリーにVehicleエンティティを一つ保存する。<br>
     * 既存のエンティティが無いため、新規登録されることを検証する。
     */
    @Test
    public void saveNewVehicleTest() {
        VehicleId id = VehicleId.newId();
        String oldVehicleName = "old vehicle";
        CommunicationId oldCommId = CommunicationId.newId();

        Vehicle vehicle = VehicleFactory.newInstance(id);
        Version version = vehicle.getVersion();

        vehicle.nameVehicle(oldVehicleName);
        vehicle.giveCommId(oldCommId);

        Version newVersion = vehicle.getNewVersion();

        repository.save(vehicle);

        VehicleRecord record = new VehicleRecord();
        record.setId(id.getId());
        record.setName(oldVehicleName);
        record.setCommId(oldCommId.getId());
        record.setVersion(version.getVersion());
        record.setNewVersion(newVersion.getVersion());

        verify(vehicleMapper, times(1)).create(record);
        verify(vehicleMapper, times(0)).update(any());
    }

    /**
     * リポジトリーにVehicleエンティティを一つ保存する。<br>
     * 既存のエンティティが存在するため、更新されることを検証する。
     */
    @Test
    public void savePreExistVehicleTest() {
        VehicleId id = VehicleId.newId();
        String oldVehicleName = "old vehicle";
        CommunicationId oldCommId = CommunicationId.newId();
        Version version = Version.newVersion();

        VehicleRecord before = new VehicleRecord();
        before.setId(id.getId());
        before.setName(oldVehicleName);
        before.setCommId(oldCommId.getId());
        before.setVersion(version.getVersion());

        when(vehicleMapper.find(id.getId())).thenReturn(before);

        Vehicle vehicle = repository.getById(id);

        String newVehicleName = "new vehicle";
        CommunicationId newCommId = CommunicationId.newId();

        vehicle.nameVehicle(newVehicleName);
        vehicle.giveCommId(newCommId);

        Version newVersion = vehicle.getNewVersion();

        repository.save(vehicle);

        VehicleRecord after = new VehicleRecord();
        after.setId(id.getId());
        after.setName(newVehicleName);
        after.setCommId(newCommId.getId());
        after.setVersion(version.getVersion());
        after.setNewVersion(newVersion.getVersion());

        verify(vehicleMapper, times(0)).create(any());
        verify(vehicleMapper, times(1)).update(after);
    }

    /**
     * リポジトリーからVehicleエンティティを一つ削除する。
     */
    @Test
    public void removeVehicleTest() {
        VehicleId id = VehicleId.newId();
        Version version = Version.newVersion();

        repository.remove(id, version);

        DeleteCondition condition = new DeleteCondition();
        condition.setId(id.getId());
        condition.setVersion(version.getVersion());

        verify(vehicleMapper, times(1)).delete(condition);
    }
}