package net.tomofiles.skysign.communication.event_listener;

import org.springframework.scheduling.annotation.Async;
import org.springframework.stereotype.Component;
import org.springframework.transaction.event.TransactionPhase;
import org.springframework.transaction.event.TransactionalEventListener;

import lombok.RequiredArgsConstructor;
import net.tomofiles.skysign.communication.domain.vehicle.CommunicationIdChangedEvent;
import net.tomofiles.skysign.communication.event_listener.dpo.CreateCommunicationRequestDpoEvent;
import net.tomofiles.skysign.communication.event_listener.dpo.RecreateCommunicationRequestDpoEvent;
import net.tomofiles.skysign.communication.usecase.ManageCommunicationService;

@Component
@RequiredArgsConstructor
public class CommunicationIdChangedEventHandler {
    
    private final ManageCommunicationService manageCommunicationService;

    @TransactionalEventListener(phase = TransactionPhase.AFTER_COMMIT)
    @Async
    public void processCommunicationIdChangedEvent(CommunicationIdChangedEvent event) {
        if (event.isFirst()) {
            CreateCommunicationRequestDpoEvent requestDpo = new CreateCommunicationRequestDpoEvent(event);
            this.manageCommunicationService.createCommunication(requestDpo, communication -> {/** 何もしない */});
        } else {
            RecreateCommunicationRequestDpoEvent requestDpo = new RecreateCommunicationRequestDpoEvent(event);
            this.manageCommunicationService.recreateCommunication(requestDpo, communication -> {/** 何もしない */});
        }
    }
}