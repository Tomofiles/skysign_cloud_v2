package net.tomofiles.skysign.communication.api.dpo;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

import net.tomofiles.skysign.communication.domain.communication.Communication;
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
                .map(communication -> {
                        return proto.skysign.common.Communication.newBuilder()
                                .setId(communication.getId().getId())
                                .setVehicleId(communication.getVehicleId().getId())
                                .setMissionId(communication.getMissionId().getId())
                                .build();
                })
                .collect(Collectors.toList());
        return proto.skysign.ListCommunicationsResponses.newBuilder()
                .addAllCommunications(r)
                .build();
    }

}