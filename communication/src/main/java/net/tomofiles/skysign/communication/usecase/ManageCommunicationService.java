package net.tomofiles.skysign.communication.usecase;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import org.springframework.transaction.annotation.Transactional;

import net.tomofiles.skysign.communication.domain.communication.Communication;
import net.tomofiles.skysign.communication.domain.communication.CommunicationFactory;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.domain.communication.CommunicationRepository;

@Component
public class ManageCommunicationService {

    @Autowired
    private CommunicationRepository communicationRepository;

    @Transactional
    public void recreateCommunication(CommunicationId beforeId, CommunicationId afterId) {
        Communication communication = this.communicationRepository.getById(beforeId);

        if (communication != null) {
            this.communicationRepository.remove(beforeId, communication.getVersion());
        }

        communication = CommunicationFactory.newInstance(afterId);

        this.communicationRepository.save(communication);
    }
}