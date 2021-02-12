package net.tomofiles.skysign.communication.api;

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

import net.tomofiles.skysign.communication.domain.communication.Communication;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.domain.communication.CommunicationRepository;
import net.tomofiles.skysign.communication.domain.vehicle.CommunicationIdGaveEvent;
import net.tomofiles.skysign.communication.domain.vehicle.CommunicationIdRemovedEvent;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleId;
import net.tomofiles.skysign.communication.domain.vehicle.Version;
import net.tomofiles.skysign.communication.infra.event.listener.proto.CommunicationIdGaveEventPb;
import net.tomofiles.skysign.communication.infra.event.listener.proto.CommunicationIdRemovedEventPb;
import net.tomofiles.skysign.communication.service.ManageCommunicationService;

import static net.tomofiles.skysign.communication.domain.communication.CommunicationObjectMother.newNormalCommunication;

public class CommunicationEventHandlerTests {
    
    private static final CommunicationId DEFAULT_COMMUNICATION_ID = new CommunicationId(UUID.randomUUID().toString());
    private static final VehicleId DEFAULT_VEHICLE_ID = new VehicleId(UUID.randomUUID().toString());
    private static final boolean DEFAULT_CONTROLLED = true;
    private static final Version DEFAULT_VERSION = new Version(UUID.randomUUID().toString());
    private static final String EXCHANGE_NAME_GAVE_EVENT = "exchange_name_gave_event";
    private static final String EXCHANGE_NAME_REMOVED_EVENT = "exchange_name_removed_event";

    @Mock
    private CommunicationRepository repository;

    @InjectMocks
    private ManageCommunicationService service;

    private CommunicationEventHandler eventHandler;

    @BeforeEach
    public void beforeEach() {
        initMocks(this);

        this.eventHandler = new CommunicationEventHandler(this.service);
        this.eventHandler.setEXCHANGE_NAME_GAVE_EVENT(EXCHANGE_NAME_GAVE_EVENT);
        this.eventHandler.setEXCHANGE_NAME_REMOVED_EVENT(EXCHANGE_NAME_REMOVED_EVENT);
    }

    /**
     * Vehicleが作成されたときにCommunicationIDが付与されたイベントを
     * 受信した場合の処理を確認する。<br>
     * 新しくCommunicationのレコードが作成されたことを検証する。
     */
    @Test
    public void fireCommunicationIdGaveEvent() throws Exception {
        CommunicationIdGaveEvent event = new CommunicationIdGaveEvent(
                DEFAULT_COMMUNICATION_ID,
                DEFAULT_VEHICLE_ID,
                DEFAULT_VERSION
        );
        CommunicationIdGaveEventPb eventPb = new CommunicationIdGaveEventPb(event);

        this.eventHandler.processCommunicationIdGaveEvent(eventPb.getMessage().getBody());

        ArgumentCaptor<Communication> commCaptor = ArgumentCaptor.forClass(Communication.class);
        verify(this.repository, times(1)).save(commCaptor.capture());

        assertThat(commCaptor.getValue().getId()).isEqualTo(DEFAULT_COMMUNICATION_ID);
        assertThat(commCaptor.getValue().getVehicleId()).isEqualTo(DEFAULT_VEHICLE_ID);
    }

    /**
     * Vehicleが更新されたときにCommunicationIDが削除されたイベントを
     * 受信した場合の処理を確認する。<br>
     * 古いCommunicationのレコードが削除されたことを検証する。
     */
    @Test
    public void fireCommunicationIdRemovedEvent() throws Exception {
        when(this.repository.getById(DEFAULT_COMMUNICATION_ID))
                .thenReturn(newNormalCommunication(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_VEHICLE_ID,
                        DEFAULT_CONTROLLED,
                        null, // テストに使用しないためNull
                        null)); // テストに使用しないためNull

        CommunicationIdRemovedEvent event = new CommunicationIdRemovedEvent(
                DEFAULT_COMMUNICATION_ID,
                DEFAULT_VERSION
        );
        CommunicationIdRemovedEventPb eventPb = new CommunicationIdRemovedEventPb(event);

        this.eventHandler.processCommunicationIdRemovedEvent(eventPb.getMessage().getBody());

        verify(this.repository, times(1)).remove(DEFAULT_COMMUNICATION_ID);
    }
}