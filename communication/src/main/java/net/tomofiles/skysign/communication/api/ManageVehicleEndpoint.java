package net.tomofiles.skysign.communication.api;

import java.util.List;
import java.util.NoSuchElementException;
import java.util.stream.Collectors;

import org.lognet.springboot.grpc.GRpcService;
import org.springframework.stereotype.Controller;

import io.grpc.Status;
import io.grpc.stub.StreamObserver;
import lombok.AllArgsConstructor;
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
import proto.skysign.ManageVehicleServiceGrpc.ManageVehicleServiceImplBase;

@GRpcService
@Controller
@AllArgsConstructor
public class ManageVehicleEndpoint extends ManageVehicleServiceImplBase {

    private final ManageVehicleService service;

    @Override
    public void listVehicles(ListVehiclesRequest request, StreamObserver<ListVehiclesResponses> responseObserver) {
        List<VehicleDto> vehicles;
        
        try {
            vehicles = this.service.getAllVehicle();
        } catch (Exception e) {
            responseObserver.onError(Status
                    .INTERNAL
                    .withCause(e)
                    .asRuntimeException());
            return;
        }

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
        VehicleDto vehicle;
        
        try {
            vehicle = this.service.getVehicle(request.getId());
        } catch (NoSuchElementException e) {
            responseObserver.onError(Status
                    .NOT_FOUND
                    .withCause(e)
                    .withDescription(e.getMessage())
                    .asRuntimeException());
            return;
        } catch (Exception e) {
            responseObserver.onError(Status
                    .INTERNAL
                    .withCause(e)
                    .asRuntimeException());
            return;
        }

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
        String id;
        
        try {
            id = this.service.createVehicle(request.getName(), request.getCommId());
        } catch (Exception e) {
            responseObserver.onError(Status
                    .INTERNAL
                    .withCause(e)
                    .asRuntimeException());
            return;
        }

        Vehicle r = Vehicle.newBuilder().setId(id).setName(request.getName()).setCommId(request.getCommId()).build();
        responseObserver.onNext(r); 
        responseObserver.onCompleted();
    }

    @Override
    public void updateVehicle(UpdateVehicleRequest request, StreamObserver<Vehicle> responseObserver) {
        try {
            this.service.updateVehicle(request.getId(), request.getName(), request.getCommId());
        } catch (NoSuchElementException e) {
            responseObserver.onError(Status
                    .NOT_FOUND
                    .withCause(e)
                    .withDescription(e.getMessage())
                    .asRuntimeException());
            return;
        } catch (Exception e) {
            responseObserver.onError(Status
                    .INTERNAL
                    .withCause(e)
                    .asRuntimeException());
            return;
        }

        Vehicle r = Vehicle.newBuilder().setId(request.getId()).setName(request.getName()).setCommId(request.getCommId()).build();
        responseObserver.onNext(r); 
        responseObserver.onCompleted();
    }

    @Override
    public void deleteVehicle(DeleteVehicleRequest request, StreamObserver<Empty> responseObserver) {
        try {
            this.service.deleteVehicle(request.getId());
        } catch (NoSuchElementException e) {
            responseObserver.onError(Status
                    .NOT_FOUND
                    .withCause(e)
                    .withDescription(e.getMessage())
                    .asRuntimeException());
            return;
        } catch (Exception e) {
            responseObserver.onError(Status
                    .INTERNAL
                    .withCause(e)
                    .asRuntimeException());
            return;
        }

        responseObserver.onNext(Empty.newBuilder().build()); 
        responseObserver.onCompleted();
    }
}