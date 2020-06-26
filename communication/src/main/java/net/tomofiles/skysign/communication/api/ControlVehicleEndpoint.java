package net.tomofiles.skysign.communication.api;

import org.lognet.springboot.grpc.GRpcService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;

import io.grpc.stub.StreamObserver;
import net.tomofiles.skysign.communication.usecase.ControlVehicleService;
import proto.skysign.StandByRequest;
import proto.skysign.StandByResponse;
import proto.skysign.CommunicationUserServiceGrpc.CommunicationUserServiceImplBase;

@GRpcService
@Controller
public class ControlVehicleEndpoint extends CommunicationUserServiceImplBase {
    
    @Autowired
    private ControlVehicleService service;

    @Override
    public void standBy(StandByRequest request, StreamObserver<StandByResponse> responseObserver) {
        this.service.standBy(request.getId(), request.getMissionId());

        StandByResponse r = StandByResponse.newBuilder().setId(request.getId()).setMissionId(request.getMissionId()).build();
        responseObserver.onNext(r); 
        responseObserver.onCompleted();
    }
}