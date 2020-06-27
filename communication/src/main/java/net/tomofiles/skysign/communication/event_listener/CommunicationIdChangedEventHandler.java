package net.tomofiles.skysign.communication.event_listener;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.scheduling.annotation.Async;
import org.springframework.stereotype.Component;
import org.springframework.transaction.event.TransactionPhase;
import org.springframework.transaction.event.TransactionalEventListener;

import net.tomofiles.skysign.communication.domain.vehicle.CommunicationIdChangedEvent;
import net.tomofiles.skysign.communication.usecase.ManageCommunicationService;

@Component
public class CommunicationIdChangedEventHandler {
    
    @Autowired
    private ManageCommunicationService manageCommunicationService;

    @TransactionalEventListener(phase = TransactionPhase.AFTER_COMMIT)
    @Async
    public void processCommunicationIdChangedEvent(CommunicationIdChangedEvent event) {
        this.manageCommunicationService.recreateCommunication(
                event.getBeforeId() == null ? null : event.getBeforeId().getId(),
                event.getAfterId() == null ? null : event.getAfterId().getId());
    }
}