package net.tomofiles.skysign.communication.usecase;

import java.util.List;
import org.springframework.stereotype.Component;
import org.springframework.transaction.annotation.Transactional;

import lombok.AllArgsConstructor;
import net.tomofiles.skysign.communication.domain.communication.CommandId;
import net.tomofiles.skysign.communication.domain.communication.CommandType;
import net.tomofiles.skysign.communication.domain.communication.Communication;
import net.tomofiles.skysign.communication.domain.communication.CommunicationRepository;
import net.tomofiles.skysign.communication.usecase.dpo.GetCommunicationRequestDpo;
import net.tomofiles.skysign.communication.usecase.dpo.GetCommunicationResponseDpo;
import net.tomofiles.skysign.communication.usecase.dpo.PullCommandRequestDpo;
import net.tomofiles.skysign.communication.usecase.dpo.PullCommandResponseDpo;
import net.tomofiles.skysign.communication.usecase.dpo.PushTelemetryRequestDpo;
import net.tomofiles.skysign.communication.usecase.dpo.PushTelemetryResponseDpo;

@Component
@AllArgsConstructor
public class CommunicateEdgeService {

    private final CommunicationRepository communicationRepository;

    @Transactional
    public void pushTelemetry(PushTelemetryRequestDpo requestDpo, PushTelemetryResponseDpo responseDpo) {
        Communication communication = this.communicationRepository.getById(requestDpo.getCommId());

        if (communication == null) {
            return;
        }

        responseDpo.setCommunication(communication);

        communication.pushTelemetry(requestDpo.getTelemetry());
        
        List<CommandId> commandIds = communication.getCommandIds();

        this.communicationRepository.save(communication);

        responseDpo.setCommandIds(commandIds);
    }

    @Transactional
    public void pullCommand(PullCommandRequestDpo requestDpo, PullCommandResponseDpo responseDpo) {
        Communication communication = this.communicationRepository.getById(requestDpo.getCommId());

        if (communication == null) {
            return;
        }

        responseDpo.setCommunication(communication);

        CommandType commandType = communication.pullCommandById(requestDpo.getCommandId());
        
        if (commandType == null) {
            return;
        }

        this.communicationRepository.save(communication);

        responseDpo.setCommandType(commandType);
    }

    @Transactional
    public void getCommunication(GetCommunicationRequestDpo requestDpo, GetCommunicationResponseDpo responseDpo) {
        Communication communication = this.communicationRepository.getById(requestDpo.getCommId());

        responseDpo.setCommunication(communication);
    }
}