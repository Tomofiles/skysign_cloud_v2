package net.tomofiles.skysign.vehicle.infra.vehicle;

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

import net.tomofiles.skysign.vehicle.domain.vehicle.Generator;
import net.tomofiles.skysign.vehicle.domain.vehicle.Vehicle;
import net.tomofiles.skysign.vehicle.domain.vehicle.VehicleId;
import net.tomofiles.skysign.vehicle.domain.vehicle.Version;
import net.tomofiles.skysign.vehicle.infra.common.DeleteCondition;

import static net.tomofiles.skysign.vehicle.domain.vehicle.VehicleObjectMother.newCarbonCopiedVehicle;
import static net.tomofiles.skysign.vehicle.infra.vehicle.RecordObjectMother.newCarbonCopiedVehicleRecord;

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
        when(this.vehicleMapper.find(DEFAULT_VEHICLE_ID.getId()))
                .thenReturn(newCarbonCopiedVehicleRecord(DEFAULT_VEHICLE_ID, DEFAULT_VERSION, DEFAULT_GENERATOR.get()));

        Vehicle vehicle = this.repository.getById(DEFAULT_VEHICLE_ID);

        Vehicle expectVehicle = newCarbonCopiedVehicle(DEFAULT_VEHICLE_ID, DEFAULT_VERSION, DEFAULT_GENERATOR.get());

        assertAll(
            () -> assertThat(vehicle.getId()).isEqualTo(expectVehicle.getId()),
            () -> assertThat(vehicle.getVehicleName()).isEqualTo(expectVehicle.getVehicleName()),
            () -> assertThat(vehicle.getCommunicationId()).isEqualTo(expectVehicle.getCommunicationId()),
            () -> assertThat(vehicle.isCarbonCopy()).isEqualTo(expectVehicle.isCarbonCopy()),
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
        Vehicle vehicle = this.repository.getById(DEFAULT_VEHICLE_ID);

        assertThat(vehicle).isNull();
    }

    /**
     * リポジトリーからVehicleエンティティをすべて取得する。
     */
    @Test
    public void getAllVehiclesTest() {
        when(this.vehicleMapper.findAll()).thenReturn(Arrays.asList(new VehicleRecord[] {
            newCarbonCopiedVehicleRecord(DEFAULT_VEHICLE_ID, DEFAULT_VERSION, DEFAULT_GENERATOR.get()),
            newCarbonCopiedVehicleRecord(DEFAULT_VEHICLE_ID, DEFAULT_VERSION, DEFAULT_GENERATOR.get()),
            newCarbonCopiedVehicleRecord(DEFAULT_VEHICLE_ID, DEFAULT_VERSION, DEFAULT_GENERATOR.get())
        }));

        List<Vehicle> vehicles = this.repository.getAll();

        Vehicle expectVehicle = newCarbonCopiedVehicle(DEFAULT_VEHICLE_ID, DEFAULT_VERSION, DEFAULT_GENERATOR.get());

        assertAll(
            () -> assertThat(vehicles).hasSize(3),
            () -> assertThat(vehicles.get(0).getId()).isEqualTo(expectVehicle.getId()),
            () -> assertThat(vehicles.get(0).getVehicleName()).isEqualTo(expectVehicle.getVehicleName()),
            () -> assertThat(vehicles.get(0).getCommunicationId()).isEqualTo(expectVehicle.getCommunicationId()),
            () -> assertThat(vehicles.get(0).isCarbonCopy()).isEqualTo(expectVehicle.isCarbonCopy()),
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
        List<Vehicle> vehicles = this.repository.getAll();

        assertThat(vehicles).hasSize(0);
    }

    /**
     * リポジトリーからカーボンコピーでないオリジナルのVehicleエンティティをすべて取得する。
     */
    @Test
    public void getAllOriginalVehiclesTest() {
        when(this.vehicleMapper.findAllOriginal()).thenReturn(Arrays.asList(new VehicleRecord[] {
            newCarbonCopiedVehicleRecord(DEFAULT_VEHICLE_ID, DEFAULT_VERSION, DEFAULT_GENERATOR.get()),
            newCarbonCopiedVehicleRecord(DEFAULT_VEHICLE_ID, DEFAULT_VERSION, DEFAULT_GENERATOR.get()),
            newCarbonCopiedVehicleRecord(DEFAULT_VEHICLE_ID, DEFAULT_VERSION, DEFAULT_GENERATOR.get())
        }));

        List<Vehicle> vehicles = this.repository.getAllOriginal();

        Vehicle expectVehicle = newCarbonCopiedVehicle(DEFAULT_VEHICLE_ID, DEFAULT_VERSION, DEFAULT_GENERATOR.get());

        assertAll(
            () -> assertThat(vehicles).hasSize(3),
            () -> assertThat(vehicles.get(0).getId()).isEqualTo(expectVehicle.getId()),
            () -> assertThat(vehicles.get(0).getVehicleName()).isEqualTo(expectVehicle.getVehicleName()),
            () -> assertThat(vehicles.get(0).getCommunicationId()).isEqualTo(expectVehicle.getCommunicationId()),
            () -> assertThat(vehicles.get(0).isCarbonCopy()).isEqualTo(expectVehicle.isCarbonCopy()),
            () -> assertThat(vehicles.get(0).getVersion()).isEqualTo(expectVehicle.getVersion()),
            () -> assertThat(vehicles.get(0).getNewVersion()).isEqualTo(expectVehicle.getNewVersion())
        );
    }

    /**
     * リポジトリーからカーボンコピーでないオリジナルのVehicleエンティティをすべて取得する。<br>
     * エンティティが存在しない場合、空リストが返却されることを検証する。
     */
    @Test
    public void getAllOriginalNoVehiclesTest() {
        List<Vehicle> vehicles = this.repository.getAllOriginal();

        assertThat(vehicles).hasSize(0);
    }

    /**
     * リポジトリーにVehicleエンティティを一つ保存する。<br>
     * 既存のエンティティが無いため、新規登録されることを検証する。
     */
    @Test
    public void saveNewVehicleTest() {
        this.repository.save(newCarbonCopiedVehicle(DEFAULT_VEHICLE_ID, DEFAULT_VERSION, DEFAULT_GENERATOR.get()));

        verify(this.vehicleMapper, times(1))
                .create(newCarbonCopiedVehicleRecord(DEFAULT_VEHICLE_ID, DEFAULT_VERSION, DEFAULT_GENERATOR.get()));
    }

    /**
     * リポジトリーにVehicleエンティティを一つ保存する。<br>
     * 既存のエンティティが存在するため、更新されることを検証する。
     */
    @Test
    public void savePreExistVehicleTest() {
        when(this.vehicleMapper.find(DEFAULT_VEHICLE_ID.getId()))
                .thenReturn(newCarbonCopiedVehicleRecord(DEFAULT_VEHICLE_ID, DEFAULT_VERSION, DEFAULT_GENERATOR.get()));

        Vehicle vehicle = this.repository.getById(DEFAULT_VEHICLE_ID);

        repository.save(vehicle);

        verify(this.vehicleMapper, times(1))
                .update(newCarbonCopiedVehicleRecord(DEFAULT_VEHICLE_ID, DEFAULT_VERSION, DEFAULT_GENERATOR.get()));
    }

    /**
     * リポジトリーからVehicleエンティティを一つ削除する。
     */
    @Test
    public void removeVehicleTest() {
        this.repository.remove(DEFAULT_VEHICLE_ID, DEFAULT_VERSION);

        DeleteCondition condition = new DeleteCondition();
        condition.setId(DEFAULT_VEHICLE_ID.getId());
        condition.setVersion(DEFAULT_VERSION.getVersion());

        verify(this.vehicleMapper, times(1)).delete(condition);
    }
}