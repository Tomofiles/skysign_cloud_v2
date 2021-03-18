package net.tomofiles.skysign.vehicle.domain.vehicle;

import static com.google.common.truth.Truth.assertThat;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.Mock;

import static org.junit.Assert.assertThrows;
import static org.junit.jupiter.api.Assertions.assertAll;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;
import static org.mockito.MockitoAnnotations.initMocks;

import java.util.UUID;
import java.util.function.Supplier;

import net.tomofiles.skysign.vehicle.event.Publisher;

import static net.tomofiles.skysign.vehicle.domain.vehicle.VehicleObjectMother.newNormalVehicle;

public class ManagementVehiclesTests {
    
    private static final VehicleId DEFAULT_VEHICLE_ID1 = new VehicleId(UUID.randomUUID().toString());
    private static final VehicleId DEFAULT_VEHICLE_ID2 = new VehicleId(UUID.randomUUID().toString());
    private static final CommunicationId DEFAULT_COMMUNICATION_ID = new CommunicationId("comm id");
    private static final Version DEFAULT_VERSION1 = new Version(UUID.randomUUID().toString());
    private static final Version DEFAULT_VERSION2 = new Version(UUID.randomUUID().toString());
    private static final Supplier<Generator> DEFAULT_GENERATOR_1CALL = () -> {
        return new Generator(){
            @Override
            public VehicleId newVehicleId() {
                return DEFAULT_VEHICLE_ID1;
            }

            @Override
            public Version newVersion() {
                return DEFAULT_VERSION2;
            }
        };
    };
    private static final Supplier<Generator> DEFAULT_GENERATOR_2CALL = () -> {
        return new Generator(){
            private int count = 0;

            @Override
            public VehicleId newVehicleId() {
                return DEFAULT_VEHICLE_ID1;
            }

            @Override
            public Version newVersion() {
                if (count == 0) {
                    count++;
                    return DEFAULT_VERSION1;
                } else {
                    return DEFAULT_VERSION2;
                }
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
     * Userが、新しいVehicleエンティティを作成する。<br>
     * Vehicleエンティティの初期状態を検証する。
     */
    @Test
    public void createNewVehicleTest() {
        Vehicle vehicle = VehicleFactory.newInstance(DEFAULT_GENERATOR_2CALL.get());

        assertAll(
            () -> assertThat(vehicle.getId()).isEqualTo(DEFAULT_VEHICLE_ID1),
            () -> assertThat(vehicle.getVehicleName()).isNull(),
            () -> assertThat(vehicle.getCommId()).isNull(),
            () -> assertThat(vehicle.isCarbonCopy()).isFalse(),
            () -> assertThat(vehicle.getVersion()).isEqualTo(DEFAULT_VERSION1),
            () -> assertThat(vehicle.getNewVersion()).isEqualTo(DEFAULT_VERSION1)
        );
    }

    /**
     * Userが、新しいVehicleエンティティに対してVehicle Nameを付与する。
     */
    @Test
    public void changeVehiclesNameTest() {
        Vehicle vehicle = VehicleFactory.newInstance(DEFAULT_GENERATOR_2CALL.get());

        String newVehicleName = "new vehicle";
        vehicle.nameVehicle(newVehicleName);

        assertAll(
            () -> assertThat(vehicle.getVehicleName()).isEqualTo(newVehicleName),
            () -> assertThat(vehicle.getVersion()).isEqualTo(DEFAULT_VERSION1),
            () -> assertThat(vehicle.getNewVersion()).isEqualTo(DEFAULT_VERSION2)
        );
    }

    /**
     * Userが、新しいVehicleエンティティを作成して、CommunicationIDを付与する。<br>
     * CommunicationIDを付与したら、イベントを生成して発行する。<br>
     * その際、新しいCommunicationIDのみ、購読者に通知されることを検証する。
     */
    @Test
    public void changeNewVehiclesCommIdAndPublishEventTest() {
        Vehicle vehicle = VehicleFactory.newInstance(DEFAULT_GENERATOR_2CALL.get());

        vehicle.setPublisher(this.publisher);

        CommunicationId newCommId = new CommunicationId("new comm id");
        vehicle.giveCommId(newCommId);

        CommunicationIdGaveEvent event
                = new CommunicationIdGaveEvent(
                    newCommId,
                    DEFAULT_VEHICLE_ID1,
                    DEFAULT_VERSION2
                );

        assertAll(
            () -> assertThat(vehicle.getCommId()).isEqualTo(newCommId),
            () -> assertThat(vehicle.getVersion()).isEqualTo(DEFAULT_VERSION1),
            () -> assertThat(vehicle.getNewVersion()).isEqualTo(DEFAULT_VERSION2),
            () -> verify(this.publisher, times(1)).publish(event)
        );
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
        when(this.repository.getById(DEFAULT_VEHICLE_ID1))
                .thenReturn(newNormalVehicle(DEFAULT_VEHICLE_ID1, DEFAULT_VERSION1, DEFAULT_GENERATOR_1CALL.get()));

        Vehicle vehicle = this.repository.getById(DEFAULT_VEHICLE_ID1);

        vehicle.setPublisher(this.publisher);

        CommunicationId newCommId = new CommunicationId("new comm id");
        vehicle.giveCommId(newCommId);

        CommunicationIdRemovedEvent removedEvent
                = new CommunicationIdRemovedEvent(
                    DEFAULT_COMMUNICATION_ID,
                    DEFAULT_VERSION2
                );

        CommunicationIdGaveEvent gaveEvent
                = new CommunicationIdGaveEvent(
                    newCommId,
                    DEFAULT_VEHICLE_ID1,
                    DEFAULT_VERSION2
                );

        assertAll(
            () -> assertThat(vehicle.getCommId()).isEqualTo(newCommId),
            () -> assertThat(vehicle.getVersion()).isEqualTo(DEFAULT_VERSION1),
            () -> assertThat(vehicle.getNewVersion()).isEqualTo(DEFAULT_VERSION2),
            () -> verify(this.publisher, times(1)).publish(removedEvent),
            () -> verify(this.publisher, times(1)).publish(gaveEvent)
        );
    }

    /**
     * Userが、既存のVehicleエンティティに対してCommunicationIDを更新する。<br>
     * VehicleエンティティからCommunicationIDを削除することで、<br>
     * イベントを生成して発行する。<br>
     * その際、古いCommunicationIDが、購読者に通知されることを検証する。
     */
    @Test
    public void removePreExistVehiclesCommIdAndPublishEventTest() {
        when(this.repository.getById(DEFAULT_VEHICLE_ID1))
                .thenReturn(newNormalVehicle(DEFAULT_VEHICLE_ID1, DEFAULT_VERSION1, DEFAULT_GENERATOR_1CALL.get()));

        Vehicle vehicle = this.repository.getById(DEFAULT_VEHICLE_ID1);

        vehicle.setPublisher(this.publisher);

        vehicle.removeCommId();

        CommunicationIdRemovedEvent event
                = new CommunicationIdRemovedEvent(
                    DEFAULT_COMMUNICATION_ID,
                    DEFAULT_VERSION2
                );

        assertAll(
            () -> assertThat(vehicle.getCommId()).isNull(),
            () -> assertThat(vehicle.getVersion()).isEqualTo(DEFAULT_VERSION1),
            () -> assertThat(vehicle.getNewVersion()).isEqualTo(DEFAULT_VERSION2),
            () -> verify(this.publisher, times(1)).publish(event)
        );
    }

    /**
     * 既存のVehicleエンティティのカーボンコピーを作成する。<br>
     * コピーされたVehicleエンティティの内部状態を検証する。
     */
    @Test
    public void carbonCopyVehicleTest() {
        Generator generator = DEFAULT_GENERATOR_1CALL.get();
        Vehicle original = VehicleFactory.newInstance(generator);

        String vehicleName = "now vehicle";
        CommunicationId commId = new CommunicationId("now comm id");

        original.setVehicleName(vehicleName);
        original.setCommId(commId);

        Vehicle vehicle = VehicleFactory.copy(original, DEFAULT_VEHICLE_ID2, generator);

        assertAll(
            () -> assertThat(vehicle.getId()).isEqualTo(DEFAULT_VEHICLE_ID2),
            () -> assertThat(vehicle.getVehicleName()).isEqualTo(vehicleName),
            () -> assertThat(vehicle.getCommId()).isEqualTo(commId),
            () -> assertThat(vehicle.isCarbonCopy()).isTrue(),
            () -> assertThat(vehicle.getVersion()).isEqualTo(DEFAULT_VERSION2),
            () -> assertThat(vehicle.getNewVersion()).isEqualTo(DEFAULT_VERSION2)
        );
    }

    /**
     * カーボンコピーされたVehicleエンティティに対してVehicle Nameを付与する。
     * 更新時に例外がスローされることを検証する。
     */
    @Test
    public void cannotChangeErrorWhenChangeCarbonCopiedVehiclesNameTest() {
        Generator generator = DEFAULT_GENERATOR_1CALL.get();
        Vehicle original = VehicleFactory.newInstance(generator);

        String vehicleName = "now vehicle";
        CommunicationId commId = new CommunicationId("now comm id");

        original.setVehicleName(vehicleName);
        original.setCommId(commId);

        Vehicle vehicle = VehicleFactory.copy(original, DEFAULT_VEHICLE_ID2, generator);

        String newVehicleName = "new vehicle";

        CannotChangeVehicleException e
                = assertThrows(
                        CannotChangeVehicleException.class,
                        () -> vehicle.nameVehicle(newVehicleName));

        assertThat(e).hasMessageThat().contains("cannot change carbon copied vehicle");
    }

    /**
     * カーボンコピーされたVehicleエンティティに対してCommunicationIDを付与する。
     * 更新時に例外がスローされることを検証する。
     */
    @Test
    public void cannotChangeErrorWhenChangeCarbonCopiedVehiclesCommIdTest() {
        Generator generator = DEFAULT_GENERATOR_1CALL.get();
        Vehicle original = VehicleFactory.newInstance(generator);

        String vehicleName = "now vehicle";
        CommunicationId commId = new CommunicationId("now comm id");

        original.setVehicleName(vehicleName);
        original.setCommId(commId);

        Vehicle vehicle = VehicleFactory.copy(original, DEFAULT_VEHICLE_ID2, generator);

        CommunicationId newCommId = new CommunicationId("new comm id");

        CannotChangeVehicleException e
                = assertThrows(
                        CannotChangeVehicleException.class,
                        () -> vehicle.giveCommId(newCommId));

        assertThat(e).hasMessageThat().contains("cannot change carbon copied vehicle");
    }

    /**
     * カーボンコピーされたVehicleエンティティのCommunicationIDを削除する。
     * 更新時に例外がスローされることを検証する。
     */
    @Test
    public void cannotChangeErrorWhenRemoveCarbonCopiedVehiclesCommIdTest() {
        Generator generator = DEFAULT_GENERATOR_1CALL.get();
        Vehicle original = VehicleFactory.newInstance(generator);

        String vehicleName = "now vehicle";
        CommunicationId commId = new CommunicationId("now comm id");

        original.setVehicleName(vehicleName);
        original.setCommId(commId);

        Vehicle vehicle = VehicleFactory.copy(original, DEFAULT_VEHICLE_ID2, generator);

        CannotChangeVehicleException e
                = assertThrows(
                        CannotChangeVehicleException.class,
                        () -> vehicle.removeCommId());

        assertThat(e).hasMessageThat().contains("cannot change carbon copied vehicle");
    }

}