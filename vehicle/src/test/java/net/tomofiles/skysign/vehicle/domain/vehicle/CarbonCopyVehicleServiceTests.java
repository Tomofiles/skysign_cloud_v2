package net.tomofiles.skysign.vehicle.domain.vehicle;

import java.util.UUID;
import java.util.function.Supplier;

import static com.google.common.truth.Truth.assertThat;
import static org.junit.jupiter.api.Assertions.assertAll;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;
import static org.mockito.MockitoAnnotations.initMocks;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.ArgumentCaptor;
import org.mockito.Mock;

import net.tomofiles.skysign.vehicle.event.Publisher;

import static net.tomofiles.skysign.vehicle.domain.vehicle.VehicleObjectMother.newNormalVehicle;
import static net.tomofiles.skysign.vehicle.domain.vehicle.VehicleObjectMother.newCarbonCopiedVehicle;

public class CarbonCopyVehicleServiceTests {
    
    private static final VehicleId ORIGINAL_VEHICLE_ID = new VehicleId(UUID.randomUUID().toString());
    private static final VehicleId NEW_VEHICLE_ID = new VehicleId(UUID.randomUUID().toString());
    private static final Version DEFAULT_VERSION = new Version(UUID.randomUUID().toString());
    private static final FlightplanId DEFAULT_FLIGHTPLAN_ID = new FlightplanId(UUID.randomUUID().toString());
    private static final CommunicationId DEFAULT_COMMUNICATION_ID = new CommunicationId("comm id");
    private static final Supplier<Generator> DEFAULT_GENERATOR = () -> {
        return new Generator(){
            @Override
            public VehicleId newVehicleId() {
                throw new IllegalStateException();
            }
            @Override
            public Version newVersion() {
                throw new IllegalStateException();
            }
        };
    };

    @Mock
    private Publisher publisher;

    @Mock
    private VehicleRepository repository;

    @BeforeEach
    public void beforeEach() {
        initMocks(this);
    }

    /**
     * Vehicleのカーボンコピーを作成するサービスをテストする。
     * 指定されたIDのVehicleを、指定されたIDでコピーする。
     * コピーされたVehicle作成されたことを表すイベントが発行されることを検証する。
     */
    @Test
    public void carbonCopyVehicleTest() {
        when(this.repository.getById(ORIGINAL_VEHICLE_ID))
                .thenReturn(newNormalVehicle(
                        ORIGINAL_VEHICLE_ID,
                        DEFAULT_VERSION,
                        DEFAULT_GENERATOR.get()));

        CarbonCopyVehicleService.copy(
                DEFAULT_GENERATOR.get(), 
                this.publisher,
                this.repository, 
                ORIGINAL_VEHICLE_ID, 
                NEW_VEHICLE_ID,
                DEFAULT_FLIGHTPLAN_ID);

        ArgumentCaptor<Vehicle> vehicleCaptor = ArgumentCaptor.forClass(Vehicle.class);
        verify(this.repository, times(1)).save(vehicleCaptor.capture());

        Vehicle expectVehicle = newCarbonCopiedVehicle(
                NEW_VEHICLE_ID,
                DEFAULT_VERSION,
                DEFAULT_GENERATOR.get());

        CopiedVehicleCreatedEvent event
                = new CopiedVehicleCreatedEvent(
                    NEW_VEHICLE_ID,
                    DEFAULT_COMMUNICATION_ID,
                    DEFAULT_FLIGHTPLAN_ID
                );

        assertAll(
            () -> assertThat(vehicleCaptor.getValue()).isEqualTo(expectVehicle),
            () -> verify(this.publisher, times(1)).publish(event)
        );
    }

    /**
     * Vehicleのカーボンコピーを作成するサービスをテストする。
     * コピー後のIDのVehicleがすでに存在する場合、コピーを行わず
     * 正常終了することを検証する。
     */
    @Test
    public void copySuccessWhenAlreadyExistsVehicleWhenCarbonCopyVehicleTest() {
        when(this.repository.getById(NEW_VEHICLE_ID))
                .thenReturn(newCarbonCopiedVehicle(
                        NEW_VEHICLE_ID,
                        DEFAULT_VERSION,
                        DEFAULT_GENERATOR.get()));

        CarbonCopyVehicleService.copy(
                DEFAULT_GENERATOR.get(), 
                this.publisher,
                this.repository, 
                ORIGINAL_VEHICLE_ID, 
                NEW_VEHICLE_ID,
                DEFAULT_FLIGHTPLAN_ID);

        assertAll(
            () -> verify(this.repository, times(0)).save(any()),
            () -> verify(this.publisher, times(0)).publish(any())
        );
    }

    /**
     * Vehicleのカーボンコピーを作成するサービスをテストする。
     * 指定されたIDのVehicleの取得がエラーとなった場合、
     * 正常終了することを検証する。
     */
    @Test
    public void getErrorWhenCarbonCopyVehicleTest() {
        CarbonCopyVehicleService.copy(
                DEFAULT_GENERATOR.get(), 
                this.publisher,
                this.repository, 
                ORIGINAL_VEHICLE_ID, 
                NEW_VEHICLE_ID,
                DEFAULT_FLIGHTPLAN_ID);

        assertAll(
            () -> verify(this.repository, times(0)).save(any()),
            () -> verify(this.publisher, times(0)).publish(any())
        );
    }

}