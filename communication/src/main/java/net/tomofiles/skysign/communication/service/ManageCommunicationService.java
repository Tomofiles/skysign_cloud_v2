package net.tomofiles.skysign.communication.service;

import org.springframework.stereotype.Component;
import org.springframework.transaction.annotation.Transactional;

import lombok.AllArgsConstructor;
import net.tomofiles.skysign.communication.domain.communication.Communication;
import net.tomofiles.skysign.communication.domain.communication.CommunicationFactory;
import net.tomofiles.skysign.communication.domain.communication.CommunicationRepository;
import net.tomofiles.skysign.communication.domain.communication.Generator;
import net.tomofiles.skysign.communication.service.dpo.CreateCommunicationRequestDpo;
import net.tomofiles.skysign.communication.service.dpo.ManageCommunicationResponseDpo;
import net.tomofiles.skysign.communication.service.dpo.RecreateCommunicationRequestDpo;

@Component
@AllArgsConstructor
public class ManageCommunicationService {

    private final CommunicationRepository communicationRepository;
    private final Generator generator;

    @Transactional
    public void createCommunication(CreateCommunicationRequestDpo requestDpo, ManageCommunicationResponseDpo responseDpo) {
        Communication communication = CommunicationFactory.newInstance(
                requestDpo.getCommId(),
                requestDpo.getVehicleId(),
                this.generator);

        this.communicationRepository.save(communication);

        responseDpo.setCommunication(communication);
    }

    @Transactional
    public void recreateCommunication(RecreateCommunicationRequestDpo requestDpo, ManageCommunicationResponseDpo responseDpo) {
        Communication oldComm = this.communicationRepository.getById(requestDpo.getBeforeCommId());

        if (oldComm != null) {
            this.communicationRepository.remove(oldComm.getId());
        }
        
        Communication newComm = CommunicationFactory.newInstance(
                requestDpo.getAfterCommId(),
                requestDpo.getVehicleId(),
                generator);

        this.communicationRepository.save(newComm);

        responseDpo.setCommunication(newComm);
    }
}