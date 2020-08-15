package net.tomofiles.skysign.communication.api.dpo;

import net.tomofiles.skysign.communication.domain.communication.Communication;
import net.tomofiles.skysign.communication.domain.communication.CommunicationFactory;
import net.tomofiles.skysign.communication.domain.communication.component.CommunicationComponentDto;
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
        CommunicationComponentDto dto = CommunicationFactory.takeApart(communication);
        proto.skysign.common.Communication nonMissionIdComm =  proto.skysign.common.Communication.newBuilder()
                .setId(dto.getId())
                .setVehicleId(dto.getVehicleId())
                .build();
        if (dto.getMissionId() == null) {
            return nonMissionIdComm;
        } else {
            return proto.skysign.common.Communication.newBuilder(nonMissionIdComm)
                    .setMissionId(dto.getMissionId())
                    .build();
        }
    }
}