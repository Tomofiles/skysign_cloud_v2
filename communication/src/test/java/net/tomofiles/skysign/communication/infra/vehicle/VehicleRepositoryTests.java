package net.tomofiles.skysign.communication.infra.vehicle;

import static com.google.common.truth.Truth.assertThat;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.InjectMocks;
import org.mockito.Mock;

import static org.junit.jupiter.api.Assertions.assertAll;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;
import static org.mockito.MockitoAnnotations.initMocks;

import java.util.Arrays;
import java.util.List;
import java.util.UUID;
import java.util.function.Supplier;

import net.tomofiles.skysign.communication.domain.vehicle.Generator;
import net.tomofiles.skysign.communication.domain.vehicle.Vehicle;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleId;
import net.tomofiles.skysign.communication.domain.vehicle.Version;
import net.tomofiles.skysign.communication.infra.common.DeleteCondition;

import static net.tomofiles.skysign.communication.domain.vehicle.VehicleObjectMother.newNormalVehicle;
import static net.tomofiles.skysign.communication.infra.vehicle.RecordObjectMother.newNormalVehicleRecord;

public class VehicleRepositoryTests {
    
    private static final VehicleId DEFAULT_VEHICLE_ID = new VehicleId(UUID.randomUUID().toString());
    private static final Version DEFAULT_VERSION = new Version(UUID.randomUUID().toString());
    private static final Supplier<Generator> DEFAULT_GENERATOR = () -> {
        return new Generator(){
            @Override
            public VehicleId newVehicleId() {
                return DEFAULT_VEHICLE_ID;
            }

            @Override
            public Version newVersion() {
                return DEFAULT_VERSION;
            }
        };
    };

    @Mock
    private VehicleMapper vehicleMapper;

    @InjectMocks
    private VehicleRepositoryImpl repository;

    @BeforeEach
    public void beforeEach() {
        initMocks(this);
    }

    /**
     * リポジトリーからVehicleエンティティを一つ取得する。
     */
    @Test
    public void getVehicleByIdTest() {
        when(vehicleMapper.find(DEFAULT_VEHICLE_ID.getId()))
                .thenReturn(newNormalVehicleRecord(DEFAULT_VEHICLE_ID, DEFAULT_VERSION, DEFAULT_GENERATOR.get()));

        Vehicle vehicle = repository.getById(DEFAULT_VEHICLE_ID);

        Vehicle expectVehicle = newNormalVehicle(DEFAULT_VEHICLE_ID, DEFAULT_VERSION, DEFAULT_GENERATOR.get());

        assertAll(
            () -> assertThat(vehicle.getId()).isEqualTo(expectVehicle.getId()),
            () -> assertThat(vehicle.getVehicleName()).isEqualTo(expectVehicle.getVehicleName()),
            () -> assertThat(vehicle.getCommId()).isEqualTo(expectVehicle.getCommId()),
            () -> assertThat(vehicle.getVersion()).isEqualTo(expectVehicle.getVersion()),
            () -> assertThat(vehicle.getNewVersion()).isEqualTo(expectVehicle.getNewVersion())
        );
    }

    /**
     * リポジトリーからVehicleエンティティを一つ取得する。<br>
     * 対象のエンティティが存在しない場合、NULLが返却されることを検証する。
     */
    @Test
    public void getNoVehicleByIdTest() {
        Vehicle vehicle = repository.getById(DEFAULT_VEHICLE_ID);

        assertThat(vehicle).isNull();
    }

    /**
     * リポジトリーからVehicleエンティティをすべて取得する。
     */
    @Test
    public void getAllVehiclesTest() {
        when(vehicleMapper.findAll()).thenReturn(Arrays.asList(new VehicleRecord[] {
            newNormalVehicleRecord(DEFAULT_VEHICLE_ID, DEFAULT_VERSION, DEFAULT_GENERATOR.get()),
            newNormalVehicleRecord(DEFAULT_VEHICLE_ID, DEFAULT_VERSION, DEFAULT_GENERATOR.get()),
            newNormalVehicleRecord(DEFAULT_VEHICLE_ID, DEFAULT_VERSION, DEFAULT_GENERATOR.get())
        }));

        List<Vehicle> vehicles = repository.getAll();

        Vehicle expectVehicle = newNormalVehicle(DEFAULT_VEHICLE_ID, DEFAULT_VERSION, DEFAULT_GENERATOR.get());

        assertAll(
            () -> assertThat(vehicles).hasSize(3),
            () -> assertThat(vehicles.get(0).getId()).isEqualTo(expectVehicle.getId()),
            () -> assertThat(vehicles.get(0).getVehicleName()).isEqualTo(expectVehicle.getVehicleName()),
            () -> assertThat(vehicles.get(0).getCommId()).isEqualTo(expectVehicle.getCommId()),
            () -> assertThat(vehicles.get(0).getVersion()).isEqualTo(expectVehicle.getVersion()),
            () -> assertThat(vehicles.get(0).getNewVersion()).isEqualTo(expectVehicle.getNewVersion())
        );
    }

    /**
     * リポジトリーからVehicleエンティティをすべて取得する。<br>
     * エンティティが存在しない場合、空リストが返却されることを検証する。
     */
    @Test
    public void getAllNoVehiclesTest() {
        List<Vehicle> vehicles = repository.getAll();

        assertThat(vehicles).hasSize(0);
    }

    /**
     * リポジトリーにVehicleエンティティを一つ保存する。<br>
     * 既存のエンティティが無いため、新規登録されることを検証する。
     */
    @Test
    public void saveNewVehicleTest() {
        repository.save(newNormalVehicle(DEFAULT_VEHICLE_ID, DEFAULT_VERSION, DEFAULT_GENERATOR.get()));

        verify(vehicleMapper, times(1))
                .create(newNormalVehicleRecord(DEFAULT_VEHICLE_ID, DEFAULT_VERSION, DEFAULT_GENERATOR.get()));
    }

    /**
     * リポジトリーにVehicleエンティティを一つ保存する。<br>
     * 既存のエンティティが存在するため、更新されることを検証する。
     */
    @Test
    public void savePreExistVehicleTest() {
        when(vehicleMapper.find(DEFAULT_VEHICLE_ID.getId()))
                .thenReturn(newNormalVehicleRecord(DEFAULT_VEHICLE_ID, DEFAULT_VERSION, DEFAULT_GENERATOR.get()));

        Vehicle vehicle = repository.getById(DEFAULT_VEHICLE_ID);

        repository.save(vehicle);

        verify(vehicleMapper, times(1))
                .update(newNormalVehicleRecord(DEFAULT_VEHICLE_ID, DEFAULT_VERSION, DEFAULT_GENERATOR.get()));
    }

    /**
     * リポジトリーからVehicleエンティティを一つ削除する。
     */
    @Test
    public void removeVehicleTest() {
        repository.remove(DEFAULT_VEHICLE_ID, DEFAULT_VERSION);

        DeleteCondition condition = new DeleteCondition();
        condition.setId(DEFAULT_VEHICLE_ID.getId());
        condition.setVersion(DEFAULT_VERSION.getVersion());

        verify(vehicleMapper, times(1)).delete(condition);
    }
}