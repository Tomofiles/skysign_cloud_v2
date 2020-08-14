package net.tomofiles.skysign.communication.api.dpo;

import lombok.RequiredArgsConstructor;
import net.tomofiles.skysign.communication.domain.communication.CommandId;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.usecase.dpo.PullCommandRequestDpo;
import proto.skysign.PullCommandRequest;

@RequiredArgsConstructor
public class PullCommandRequestDpoGrpc implements PullCommandRequestDpo {

    private final PullCommandRequest request;

    @Override
    public CommunicationId getCommId() {
        return new CommunicationId(this.request.getId());
    }

    @Override
    public CommandId getCommandId() {
        return new CommandId(this.request.getCommandId());
    }
}