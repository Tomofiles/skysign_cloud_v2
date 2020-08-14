package net.tomofiles.skysign.communication.api.dpo;

import lombok.RequiredArgsConstructor;
import net.tomofiles.skysign.communication.domain.communication.CommandType;
import net.tomofiles.skysign.communication.domain.communication.Communication;
import net.tomofiles.skysign.communication.usecase.dpo.PullCommandResponseDpo;
import proto.skysign.PullCommandRequest;

@RequiredArgsConstructor
public class PullCommandResponseDpoGrpc implements PullCommandResponseDpo {

    private final PullCommandRequest request;
    private Communication communication = null;
    private CommandType commandType = null;

    @Override
    public void setCommunication(Communication communication) {
        this.communication = communication;
    }

    @Override
    public void setCommandType(CommandType commandType) {
        this.commandType = commandType;
    }

    public boolean notExistCommunication() {
        return communication == null;
    }

    public boolean notExistCommand() {
        return commandType == null;
    }

    public proto.skysign.PullCommandResponse getGrpcResponse() {
        return proto.skysign.PullCommandResponse.newBuilder()
                .setId(request.getId())
                .setCommandId(request.getCommandId())
                .setType(proto.skysign.CommandType.valueOf(commandType.name()))
                .build();
    }
}