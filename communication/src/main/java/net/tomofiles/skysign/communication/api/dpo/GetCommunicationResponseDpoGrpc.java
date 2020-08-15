package net.tomofiles.skysign.communication.api.dpo;

import net.tomofiles.skysign.communication.domain.communication.Communication;
import net.tomofiles.skysign.communication.service.dpo.GetCommunicationResponseDpo;

public class GetCommunicationResponseDpoGrpc implements GetCommunicationResponseDpo {

    private Communication communication = null;

	@Override
	public void setCommunication(Communication communication) {
		this.communication = communication;
	}

    public boolean notExistCommunication() {
        return communication == null;
    }

    public proto.skysign.common.Communication getGrpcResponse() {
        return proto.skysign.common.Communication.newBuilder()
                .setId(communication.getId().getId())
                .setVehicleId(communication.getVehicleId().getId())
                .setMissionId(communication.getMissionId().getId())
                .build();
    }
}