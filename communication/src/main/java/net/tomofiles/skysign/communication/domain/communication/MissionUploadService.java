package net.tomofiles.skysign.communication.domain.communication;

import java.util.function.Consumer;

public class MissionUploadService {
    public static void send(
        CommunicationRepository communicationRepository,
        Consumer<Communication> setCommunication,
        Consumer<CommandId> setCommandId,
        CommunicationId communicationId,
        MissionId missionId
    ) {
        Communication communication = communicationRepository.getById(communicationId);

        if (communication == null) {
            return;
        }

        CommandId commandId = communication.pushUploadMission(missionId);

        communicationRepository.save(communication);

        setCommunication.accept(communication);
        setCommandId.accept(commandId);
    }
}