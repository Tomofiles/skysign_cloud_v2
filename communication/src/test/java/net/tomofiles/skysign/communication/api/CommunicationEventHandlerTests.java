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
import net.tomofiles.skysign.communication.service.ManageCommunicationService;

import static net.tomofiles.skysign.communication.domain.communication.CommunicationObjectMother.newNormalCommunication;
import static net.tomofiles.skysign.communication.api.EventObjectMother.newNormalCommunicationIdGaveEvent;
import static net.tomofiles.skysign.communication.api.EventObjectMother.newNormalCommunicationIdRemovedEvent;

public class CommunicationEventHandlerTests {
    
    private static final CommunicationId DEFAULT_COMMUNICATION_ID = new CommunicationId(UUID.randomUUID().toString());
    private static final String QUEUE_NAME_GAVE_EVENT = "queue_name_gave_event";
    private static final String QUEUE_NAME_REMOVED_EVENT = "queue_name_removed_event";

    @Mock
    private CommunicationRepository repository;

    @InjectMocks
    private ManageCommunicationService service;

    private CommunicationEventHandler eventHandler;

    @BeforeEach
    public void beforeEach() {
        initMocks(this);

        this.eventHandler = new CommunicationEventHandler(this.service);
        this.eventHandler.setQUEUE_NAME_GAVE_EVENT(QUEUE_NAME_GAVE_EVENT);
        this.eventHandler.setQUEUE_NAME_REMOVED_EVENT(QUEUE_NAME_REMOVED_EVENT);
    }

    /**
     * Vehicleが作成されたときにCommunicationIDが付与されたイベントを
     * 受信した場合の処理を確認する。<br>
     * 新しくCommunicationのレコードが作成されたことを検証する。
     */
    @Test
    public void fireCommunicationIdGaveEvent() throws Exception {
        this.eventHandler.processCommunicationIdGaveEvent(
            newNormalCommunicationIdGaveEvent(DEFAULT_COMMUNICATION_ID));

        ArgumentCaptor<Communication> commCaptor = ArgumentCaptor.forClass(Communication.class);
        verify(this.repository, times(1)).save(commCaptor.capture());

        assertThat(commCaptor.getValue().getId()).isEqualTo(DEFAULT_COMMUNICATION_ID);
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
                        null)); // テストに使用しないためNull

        this.eventHandler.processCommunicationIdRemovedEvent(
            newNormalCommunicationIdRemovedEvent(DEFAULT_COMMUNICATION_ID));

        verify(this.repository, times(1)).remove(DEFAULT_COMMUNICATION_ID);
    }
}