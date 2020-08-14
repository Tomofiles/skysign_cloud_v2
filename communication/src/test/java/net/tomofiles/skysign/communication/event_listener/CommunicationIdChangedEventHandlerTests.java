package net.tomofiles.skysign.communication.event_listener;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.ArgumentCaptor;
import org.mockito.InjectMocks;
import org.mockito.Mock;

import static com.google.common.truth.Truth.assertThat;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;
import static org.mockito.MockitoAnnotations.initMocks;

import java.util.UUID;

import net.tomofiles.skysign.communication.domain.communication.Communication;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.domain.communication.CommunicationRepository;
import net.tomofiles.skysign.communication.domain.vehicle.CommunicationIdChangedEvent;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleId;
import net.tomofiles.skysign.communication.domain.vehicle.Version;
import net.tomofiles.skysign.communication.usecase.ManageCommunicationService;

import static net.tomofiles.skysign.communication.domain.communication.CommunicationObjectMother.newNormalCommunication;

public class CommunicationIdChangedEventHandlerTests {
    
    private static final CommunicationId DEFAULT_COMMUNICATION_ID_BEFORE = new CommunicationId(UUID.randomUUID().toString());
    private static final CommunicationId DEFAULT_COMMUNICATION_ID_AFTER = new CommunicationId(UUID.randomUUID().toString());
    private static final VehicleId DEFAULT_VEHICLE_ID = new VehicleId(UUID.randomUUID().toString());
    private static final Version DEFAULT_VERSION = new Version(UUID.randomUUID().toString());

    @Mock
    private CommunicationRepository repository;

    @InjectMocks
    private ManageCommunicationService service;

    private CommunicationIdChangedEventHandler eventHandler;

    @BeforeEach
    public void beforeEach() {
        initMocks(this);

        eventHandler = new CommunicationIdChangedEventHandler(service);
    }

    @Test
    public void firstFireCommunicationIdChangedEvent() {
        CommunicationIdChangedEvent event = new CommunicationIdChangedEvent(
                null,
                DEFAULT_COMMUNICATION_ID_AFTER,
                DEFAULT_VEHICLE_ID,
                DEFAULT_VERSION
        );

        this.eventHandler.processCommunicationIdChangedEvent(event);

        ArgumentCaptor<Communication> commCaptor = ArgumentCaptor.forClass(Communication.class);
        verify(repository, times(1)).save(commCaptor.capture());
        verify(repository, times(0)).remove(any());

        assertThat(commCaptor.getValue().getId()).isEqualTo(DEFAULT_COMMUNICATION_ID_AFTER);
        assertThat(commCaptor.getValue().getVehicleId()).isEqualTo(DEFAULT_VEHICLE_ID);
    }

    @Test
    public void secondFireCommunicationIdChangedEvent() {
        when(repository.getById(DEFAULT_COMMUNICATION_ID_BEFORE))
                .thenReturn(newNormalCommunication(
                        DEFAULT_COMMUNICATION_ID_BEFORE,
                        DEFAULT_VEHICLE_ID,
                        null, // テストに使用しないためNull
                        null)); // テストに使用しないためNull

        CommunicationIdChangedEvent event = new CommunicationIdChangedEvent(
                DEFAULT_COMMUNICATION_ID_BEFORE,
                DEFAULT_COMMUNICATION_ID_AFTER,
                DEFAULT_VEHICLE_ID,
                DEFAULT_VERSION
        );

        this.eventHandler.processCommunicationIdChangedEvent(event);

        ArgumentCaptor<Communication> commCaptor = ArgumentCaptor.forClass(Communication.class);
        verify(repository, times(1)).save(commCaptor.capture());
        verify(repository, times(1)).remove(DEFAULT_COMMUNICATION_ID_BEFORE);

        assertThat(commCaptor.getValue().getId()).isEqualTo(DEFAULT_COMMUNICATION_ID_AFTER);
        assertThat(commCaptor.getValue().getVehicleId()).isEqualTo(DEFAULT_VEHICLE_ID);
    }
}