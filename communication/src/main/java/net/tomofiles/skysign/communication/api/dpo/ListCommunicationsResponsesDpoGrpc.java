package net.tomofiles.skysign.communication.api.dpo;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

import net.tomofiles.skysign.communication.domain.communication.Communication;
import net.tomofiles.skysign.communication.domain.communication.CommunicationFactory;
import net.tomofiles.skysign.communication.service.dpo.ListCommunicationsResponsesDpo;

public class ListCommunicationsResponsesDpoGrpc implements ListCommunicationsResponsesDpo {

    private List<Communication> communications;

    public ListCommunicationsResponsesDpoGrpc() {
        this.communications = new ArrayList<>();
    }

    @Override
    public void setCommunications(List<Communication> communications) {
        this.communications = communications;
    }
    public proto.skysign.ListCommunicationsResponses getGrpcResponse() {
        List<proto.skysign.common.Communication> r = this.communications.stream()
                .map(CommunicationFactory::takeApart)
                .map(communication -> {
                    proto.skysign.common.Communication nonMissionIdComm =  proto.skysign.common.Communication.newBuilder()
                            .setId(communication.getId())
                            .setVehicleId(communication.getVehicleId())
                            .setIsControlled(communication.isControlled())
                            .build();
                    if (communication.getMissionId() == null) {
                        return nonMissionIdComm;
                    } else {
                        return proto.skysign.common.Communication.newBuilder(nonMissionIdComm)
                                .setMissionId(communication.getMissionId())
                                .build();
                    }
                })
                .collect(Collectors.toList());
        return proto.skysign.ListCommunicationsResponses.newBuilder()
                .addAllCommunications(r)
                .build();
    }

}