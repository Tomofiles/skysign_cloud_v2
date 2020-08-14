package net.tomofiles.skysign.communication.api.dpo;

import java.util.Arrays;

import lombok.RequiredArgsConstructor;
import net.tomofiles.skysign.communication.domain.communication.CommandType;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.service.dpo.PushCommandRequestDpo;
import proto.skysign.PushCommandRequest;

@RequiredArgsConstructor
public class PushCommandRequestDpoGrpc implements PushCommandRequestDpo {

    private final PushCommandRequest request;

    @Override
    public CommunicationId getCommId() {
        return new CommunicationId(this.request.getId());
    }

    @Override
    public CommandType getCommandType() {
        return Arrays.asList(CommandType.values()).stream()
                .filter(t -> t.name().equals(request.getType().name()))
                .findAny()
                .orElse(null);
    }
}