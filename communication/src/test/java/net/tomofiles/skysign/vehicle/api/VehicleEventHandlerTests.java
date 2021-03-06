package net.tomofiles.skysign.vehicle.api;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.ArgumentCaptor;
import org.mockito.InjectMocks;
import org.mockito.Mock;

import static com.google.common.truth.Truth.assertThat;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;
import static org.mockito.MockitoAnnotations.initMocks;

import java.util.UUID;
import java.util.function.Supplier;

import net.tomofiles.skysign.vehicle.domain.vehicle.Generator;
import net.tomofiles.skysign.vehicle.domain.vehicle.Vehicle;
import net.tomofiles.skysign.vehicle.domain.vehicle.Version;
import net.tomofiles.skysign.vehicle.domain.vehicle.VehicleId;
import net.tomofiles.skysign.vehicle.domain.vehicle.VehicleRepository;
import net.tomofiles.skysign.vehicle.service.ManageVehicleService;

import static net.tomofiles.skysign.vehicle.api.EventObjectMother.newNormalVehicleCopiedWhenCopiedEvent;
import static net.tomofiles.skysign.vehicle.domain.vehicle.VehicleObjectMother.newNormalVehicle;

public class VehicleEventHandlerTests {
    
    private static final VehicleId ORIGINAL_VEHICLE_ID = new VehicleId(UUID.randomUUID().toString());
    private static final VehicleId NEW_VEHICLE_ID = new VehicleId(UUID.randomUUID().toString());
    private static final Version DEFAULT_VERSION = new Version(UUID.randomUUID().toString());
    private static final String EXCHANGE_NAME_COPIED_EVENT = "exchange_name_copied_event";
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
    private VehicleRepository repository;

    @InjectMocks
    private ManageVehicleService service;

    private VehicleEventHandler eventHandler;

    @BeforeEach
    public void beforeEach() {
        initMocks(this);

        this.eventHandler = new VehicleEventHandler(this.service);
        this.eventHandler.setEXCHANGE_NAME_COPIED_EVENT(EXCHANGE_NAME_COPIED_EVENT);
    }

    /**
     * Flightplanがコピーされたときに新たなVehicleIDが発行されたイベントを
     * 受信した場合の処理を確認する。<br>
     * 対象のVehicleのカーボンコピーが作成されたことを検証する。
     */
    @Test
    public void fireVehicleCopiedWhenFlightplanCopiedEvent() throws Exception {
        when(this.repository.getById(ORIGINAL_VEHICLE_ID))
                .thenReturn(newNormalVehicle(
                        ORIGINAL_VEHICLE_ID,
                        DEFAULT_VERSION,
                        DEFAULT_GENERATOR.get()));

        this.eventHandler.processVehicleCopiedWhenFlightplanCopiedEvent(
            newNormalVehicleCopiedWhenCopiedEvent(
                ORIGINAL_VEHICLE_ID,
                NEW_VEHICLE_ID
            ));

        ArgumentCaptor<Vehicle> vehicleCaptor = ArgumentCaptor.forClass(Vehicle.class);
        verify(this.repository, times(1)).save(vehicleCaptor.capture());

        assertThat(vehicleCaptor.getValue().getId()).isEqualTo(NEW_VEHICLE_ID);
    }
}