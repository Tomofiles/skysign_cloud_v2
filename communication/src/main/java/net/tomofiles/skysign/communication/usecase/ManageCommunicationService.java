package net.tomofiles.skysign.communication.usecase;

import org.springframework.stereotype.Component;
import org.springframework.transaction.annotation.Transactional;

import lombok.AllArgsConstructor;
import net.tomofiles.skysign.communication.domain.communication.Communication;
import net.tomofiles.skysign.communication.domain.communication.CommunicationFactory;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.domain.communication.CommunicationRepository;

@Component
@AllArgsConstructor
public class ManageCommunicationService {

    private final CommunicationRepository communicationRepository;

    @Transactional
    public void recreateCommunication(String beforeId, String afterId) {
        Communication communication;

        if (beforeId != null) {
            communication = this.communicationRepository.getById(new CommunicationId(beforeId));

            if (communication != null) {
                this.communicationRepository.remove(new CommunicationId(beforeId));
            }    
        }
        
        communication = CommunicationFactory.newInstance(new CommunicationId(afterId));

        this.communicationRepository.save(communication);
    }
}