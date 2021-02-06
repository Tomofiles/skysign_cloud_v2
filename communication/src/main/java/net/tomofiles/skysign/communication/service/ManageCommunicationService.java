package net.tomofiles.skysign.communication.service;

import org.springframework.stereotype.Component;
import org.springframework.transaction.annotation.Transactional;

import lombok.AllArgsConstructor;
import net.tomofiles.skysign.communication.domain.communication.Communication;
import net.tomofiles.skysign.communication.domain.communication.CommunicationFactory;
import net.tomofiles.skysign.communication.domain.communication.CommunicationRepository;
import net.tomofiles.skysign.communication.domain.communication.Generator;
import net.tomofiles.skysign.communication.service.dpo.CreateCommunicationRequestDpo;
import net.tomofiles.skysign.communication.service.dpo.DeleteCommunicationRequestDpo;
import net.tomofiles.skysign.communication.service.dpo.ManageCommunicationResponseDpo;

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
    public void deleteCommunication(DeleteCommunicationRequestDpo requestDpo, ManageCommunicationResponseDpo responseDpo) {
        Communication communication = this.communicationRepository.getById(requestDpo.getCommId());

        if (communication == null) {
            return;
        }

        this.communicationRepository.remove(communication.getId());

        responseDpo.setCommunication(communication);
    }
}