package net.tomofiles.skysign.communication.domain.communication;

import java.util.function.Consumer;

public class CommandSendService {
    public static void send(
        CommunicationRepository communicationRepository,
        Consumer<Communication> setCommunication,
        Consumer<CommandId> setCommandId,
        CommunicationId communicationId,
        CommandType type
    ) {
        Communication communication = communicationRepository.getById(communicationId);

        if (communication == null) {
            return;
        }

        CommandId commandId = communication.pushCommand(type);

        communicationRepository.save(communication);

        setCommunication.accept(communication);
        setCommandId.accept(commandId);
    }
}