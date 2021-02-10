package net.tomofiles.skysign.communication.service;

import java.util.List;

import org.springframework.stereotype.Component;
import org.springframework.transaction.annotation.Transactional;

import lombok.AllArgsConstructor;
import net.tomofiles.skysign.communication.domain.communication.CommandId;
import net.tomofiles.skysign.communication.domain.communication.Communication;
import net.tomofiles.skysign.communication.domain.communication.CommunicationRepository;
import net.tomofiles.skysign.communication.domain.communication.TelemetrySnapshot;
import net.tomofiles.skysign.communication.service.dpo.ControlRequestDpo;
import net.tomofiles.skysign.communication.service.dpo.ControlResponseDpo;
import net.tomofiles.skysign.communication.service.dpo.ListCommunicationsResponsesDpo;
import net.tomofiles.skysign.communication.service.dpo.PullTelemetryRequestDpo;
import net.tomofiles.skysign.communication.service.dpo.PullTelemetryResponseDpo;
import net.tomofiles.skysign.communication.service.dpo.PushCommandRequestDpo;
import net.tomofiles.skysign.communication.service.dpo.PushCommandResponseDpo;
import net.tomofiles.skysign.communication.service.dpo.PushUploadMissionRequestDpo;
import net.tomofiles.skysign.communication.service.dpo.PushUploadMissionResponseDpo;
import net.tomofiles.skysign.communication.service.dpo.UncontrolRequestDpo;
import net.tomofiles.skysign.communication.service.dpo.UncontrolResponseDpo;

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
    public void control(ControlRequestDpo requestDpo, ControlResponseDpo responseDpo) {
        Communication communication = this.communicationRepository.getById(requestDpo.getCommId());

        if (communication == null) {
            return;
        }

        communication.control();

        this.communicationRepository.save(communication);

        responseDpo.setCommunication(communication);
    }

    @Transactional
    public void uncontrol(UncontrolRequestDpo requestDpo, UncontrolResponseDpo responseDpo) {
        Communication communication = this.communicationRepository.getById(requestDpo.getCommId());

        if (communication == null) {
            return;
        }

        communication.uncontrol();

        this.communicationRepository.save(communication);

        responseDpo.setCommunication(communication);
    }

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