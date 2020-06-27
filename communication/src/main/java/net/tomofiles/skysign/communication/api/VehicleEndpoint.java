package net.tomofiles.skysign.communication.api;

import java.util.List;
import java.util.stream.Collectors;

import org.lognet.springboot.grpc.GRpcService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;

import io.grpc.stub.StreamObserver;
import net.tomofiles.skysign.communication.usecase.ManageVehicleService;
import net.tomofiles.skysign.communication.usecase.dto.VehicleDto;
import proto.skysign.CreateVehicleRequest;
import proto.skysign.DeleteVehicleRequest;
import proto.skysign.Empty;
import proto.skysign.GetVehicleRequest;
import proto.skysign.ListVehiclesRequest;
import proto.skysign.ListVehiclesResponses;
import proto.skysign.UpdateVehicleRequest;
import proto.skysign.Vehicle;
import proto.skysign.VehicleServiceGrpc.VehicleServiceImplBase;

@GRpcService
@Controller
public class VehicleEndpoint extends VehicleServiceImplBase {
    
    @Autowired
    private ManageVehicleService service;

    @Override
    public void listVehicles(ListVehiclesRequest request, StreamObserver<ListVehiclesResponses> responseObserver) {
        List<VehicleDto> vehicles = this.service.getAllVehicle();

        List<Vehicle> r = vehicles.stream()
                .map(vehicle -> {
                    return Vehicle.newBuilder()
                            .setId(vehicle.getId())
                            .setName(vehicle.getName())
                            .setCommId(vehicle.getCommId())
                            .build();
                })
                .collect(Collectors.toList());
        responseObserver.onNext(ListVehiclesResponses.newBuilder().addAllVehicles(r).build()); 
        responseObserver.onCompleted();
    }

    @Override
    public void getVehicle(GetVehicleRequest request, StreamObserver<Vehicle> responseObserver) {
        VehicleDto vehicle = this.service.getVehicle(request.getId());

        Vehicle r = Vehicle.newBuilder()
                .setId(vehicle.getId())
                .setName(vehicle.getName())
                .setCommId(vehicle.getCommId())
                .build();
        responseObserver.onNext(r); 
        responseObserver.onCompleted();
    }

    @Override
    public void createVehicle(CreateVehicleRequest request, StreamObserver<Vehicle> responseObserver) {
        String id = this.service.createVehicle(request.getName(), request.getCommId());

        Vehicle r = Vehicle.newBuilder().setId(id).setName(request.getName()).setCommId(request.getCommId()).build();
        responseObserver.onNext(r); 
        responseObserver.onCompleted();
    }

    @Override
    public void updateVehicle(UpdateVehicleRequest request, StreamObserver<Vehicle> responseObserver) {
        this.service.updateVehicle(request.getId(), request.getName(), request.getCommId());

        Vehicle r = Vehicle.newBuilder().setId(request.getId()).setName(request.getName()).setCommId(request.getCommId()).build();
        responseObserver.onNext(r); 
        responseObserver.onCompleted();
    }

    @Override
    public void deleteVehicle(DeleteVehicleRequest request, StreamObserver<Empty> responseObserver) {
        this.service.deleteVehicle(request.getId());

        responseObserver.onNext(Empty.newBuilder().build()); 
        responseObserver.onCompleted();
    }
}