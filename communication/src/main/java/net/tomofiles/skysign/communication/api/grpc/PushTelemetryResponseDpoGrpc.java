package net.tomofiles.skysign.communication.api.grpc;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

import lombok.RequiredArgsConstructor;
import net.tomofiles.skysign.communication.domain.communication.CommandId;
import net.tomofiles.skysign.communication.domain.communication.Communication;
import net.tomofiles.skysign.communication.service.dpo.PushTelemetryResponseDpo;
import proto.skysign.PushTelemetryRequest;

@RequiredArgsConstructor
public class PushTelemetryResponseDpoGrpc implements PushTelemetryResponseDpo {

    private final PushTelemetryRequest request;
    private Communication communication = null;
    private List<CommandId> commandIds = new ArrayList<>();

    @Override
    public void setCommandIds(List<CommandId> commandIds) {
        this.commandIds = commandIds;
    }

	@Override
	public void setCommunication(Communication communication) {
		this.communication = communication;
	}

    public boolean notExistCommunication() {
        return communication == null;
    }

    public proto.skysign.PushTelemetryResponse getGrpcResponse() {
        return proto.skysign.PushTelemetryResponse.newBuilder()
                .setId(request.getId())
                .addAllCommandIds(
                    this.commandIds.stream()
                        .map(CommandId::getId)
                        .collect(Collectors.toList())
                )
                .build();
    }
}