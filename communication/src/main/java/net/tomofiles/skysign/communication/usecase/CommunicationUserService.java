package net.tomofiles.skysign.communication.usecase;

import java.util.List;

import org.springframework.stereotype.Component;
import org.springframework.transaction.annotation.Transactional;

import lombok.AllArgsConstructor;
import net.tomofiles.skysign.communication.domain.communication.Communication;
import net.tomofiles.skysign.communication.domain.communication.CommunicationRepository;
import net.tomofiles.skysign.communication.domain.communication.TelemetrySnapshot;
import net.tomofiles.skysign.communication.usecase.dpo.CancelRequestDpo;
import net.tomofiles.skysign.communication.usecase.dpo.CancelResponseDpo;
import net.tomofiles.skysign.communication.usecase.dpo.ListCommunicationsResponsesDpo;
import net.tomofiles.skysign.communication.usecase.dpo.PullTelemetryRequestDpo;
import net.tomofiles.skysign.communication.usecase.dpo.PullTelemetryResponseDpo;
import net.tomofiles.skysign.communication.usecase.dpo.PushCommandRequestDpo;
import net.tomofiles.skysign.communication.usecase.dpo.PushCommandResponseDpo;
import net.tomofiles.skysign.communication.usecase.dpo.StagingRequestDpo;
import net.tomofiles.skysign.communication.usecase.dpo.StagingResponseDpo;

@Component
@AllArgsConstructor
public class CommunicationUserService {

    private final CommunicationRepository communicationRepository;

    @Transactional
    public void listCommunications(ListCommunicationsResponsesDpo responsesDpo) {
        List<Communication> communications = this.communicationRepository.getAll();

        responsesDpo.setCommunications(communications);
    }

    @Transactional
    public void staging(StagingRequestDpo requestDpo, StagingResponseDpo responseDpo) {
        Communication communication = this.communicationRepository.getById(requestDpo.getCommId());

        if (communication == null) {
            return;
        }

        communication.staging(requestDpo.getMissionId());

        this.communicationRepository.save(communication);

        responseDpo.setCommunication(communication);
    }

    @Transactional
    public void cancel(CancelRequestDpo requestDpo, CancelResponseDpo responseDpo) {
        Communication communication = this.communicationRepository.getById(requestDpo.getCommId());

        if (communication == null) {
            return;
        }

        communication.cancel();

        this.communicationRepository.save(communication);

        responseDpo.setCommunication(communication);
    }

    @Transactional
    public void pushCommand(PushCommandRequestDpo requestDpo, PushCommandResponseDpo responseDpo) {
        Communication communication = this.communicationRepository.getById(requestDpo.getCommId());

        if (communication == null) {
            return;
        }

        communication.pushCommand(requestDpo.getCommandType());

        this.communicationRepository.save(communication);

        responseDpo.setCommunication(communication);
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