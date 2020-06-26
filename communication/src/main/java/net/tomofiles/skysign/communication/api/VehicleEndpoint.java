package net.tomofiles.skysign.communication.api;

import org.lognet.springboot.grpc.GRpcService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;

import io.grpc.stub.StreamObserver;
import net.tomofiles.skysign.communication.usecase.ManageVehicleService;
import proto.skysign.CreateVehicleRequest;
import proto.skysign.Vehicle;
import proto.skysign.VehicleServiceGrpc.VehicleServiceImplBase;

@GRpcService
@Controller
public class VehicleEndpoint extends VehicleServiceImplBase {
    
    @Autowired
    private ManageVehicleService service;

    @Override
    public void createVehicle(CreateVehicleRequest request, StreamObserver<Vehicle> responseObserver) {
        String id = this.service.createVehicle(request.getName(), request.getCommId());

        Vehicle r = Vehicle.newBuilder().setId(id).setName(request.getName()).setCommId(request.getCommId()).build();
        responseObserver.onNext(r); 
        responseObserver.onCompleted();
    }
}