package net.tomofiles.skysign.communication.usecase;

import java.util.NoSuchElementException;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import org.springframework.transaction.annotation.Transactional;

import net.tomofiles.skysign.communication.domain.communication.CommandId;
import net.tomofiles.skysign.communication.domain.communication.Communication;
import net.tomofiles.skysign.communication.domain.communication.CommunicationFactory;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.domain.communication.CommunicationRepository;

@Component
public class CommunicateVehicleService {

    @Autowired
    private CommunicationRepository communicationRepository;

    @Transactional
    public void pushTelemetry(String commId) {
        CommunicationId id = new CommunicationId(commId);
        Communication communication = this.communicationRepository.getById(id);

        if (communication == null) {
            throw new NoSuchElementException("communication-idに合致するCommunicationが存在しません。");
        }

        communication.pushTelemetry(0, 0);
        
        this.communicationRepository.save(communication);
    }

    @Transactional
    public void pullCommand(String commId, String commandId) {
        CommunicationId id = new CommunicationId(commId);
        CommandId cid = new CommandId(commandId);

        Communication communication = this.communicationRepository.getById(id);

        if (communication == null) {
            throw new NoSuchElementException("communication-idに合致するCommunicationが存在しません。");
        }

        communication.pullCommandById(cid);
        
        this.communicationRepository.save(communication);
    }
}