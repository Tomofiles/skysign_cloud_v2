package net.tomofiles.skysign.communication.service;

import org.springframework.stereotype.Component;
import org.springframework.transaction.annotation.Transactional;

import lombok.AllArgsConstructor;
import net.tomofiles.skysign.communication.domain.communication.CommandId;
import net.tomofiles.skysign.communication.domain.communication.Communication;
import net.tomofiles.skysign.communication.domain.communication.CommunicationRepository;
import net.tomofiles.skysign.communication.domain.communication.TelemetrySnapshot;
import net.tomofiles.skysign.communication.service.dpo.PullTelemetryRequestDpo;
import net.tomofiles.skysign.communication.service.dpo.PullTelemetryResponseDpo;
import net.tomofiles.skysign.communication.service.dpo.PushCommandRequestDpo;
import net.tomofiles.skysign.communication.service.dpo.PushCommandResponseDpo;
import net.tomofiles.skysign.communication.service.dpo.PushUploadMissionRequestDpo;
import net.tomofiles.skysign.communication.service.dpo.PushUploadMissionResponseDpo;

@Component
@AllArgsConstructor
public class CommunicationUserService {

    private final CommunicationRepository communicationRepository;

    @Transactional
    public void pushCommand(PushCommandRequestDpo requestDpo, PushCommandResponseDpo responseDpo) {
        Communication communication = this.communicationRepository.getById(requestDpo.getCommId());

        if (communication == null) {
            return;
        }

        CommandId commandId = communication.pushCommand(requestDpo.getCommandType());

        this.communicationRepository.save(communication);

        responseDpo.setCommunication(communication);
        responseDpo.setCommandId(commandId);
    }

    @Transactional
    public void pushUploadMission(PushUploadMissionRequestDpo requestDpo, PushUploadMissionResponseDpo responseDpo) {
        Communication communication = this.communicationRepository.getById(requestDpo.getCommId());

        if (communication == null) {
            return;
        }

        CommandId commandId = communication.pushUploadMission(requestDpo.getMissionId());

        this.communicationRepository.save(communication);

        responseDpo.setCommunication(communication);
        responseDpo.setCommandId(commandId);
    }

    @Transactional
    public void pullTelemetry(PullTelemetryRequestDpo requestDpo, PullTelemetryResponseDpo responseDpo) {
        Communication communication = this.communicationRepository.getById(requestDpo.getCommId());

        if (communication == null) {
            return;
        }

        TelemetrySnapshot telemetry = communication.pullTelemetry();

        responseDpo.setTelemetry(communication.getId(), telemetry);
    }
}