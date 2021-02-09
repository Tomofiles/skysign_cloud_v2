package net.tomofiles.skysign.vehicle.api;

import java.util.NoSuchElementException;

import org.lognet.springboot.grpc.GRpcService;
import org.springframework.stereotype.Controller;

import io.grpc.Status;
import io.grpc.stub.StreamObserver;
import lombok.AllArgsConstructor;
import net.tomofiles.skysign.vehicle.api.dpo.CreateVehicleRequestDpoGrpc;
import net.tomofiles.skysign.vehicle.api.dpo.CreateVehicleResponseDpoGrpc;
import net.tomofiles.skysign.vehicle.api.dpo.DeleteVehicleRequestDpoGrpc;
import net.tomofiles.skysign.vehicle.api.dpo.DeleteVehicleResponseDpoGrpc;
import net.tomofiles.skysign.vehicle.api.dpo.GetVehicleRequestDpoGrpc;
import net.tomofiles.skysign.vehicle.api.dpo.GetVehicleResponseDpoGrpc;
import net.tomofiles.skysign.vehicle.api.dpo.ListVehiclesResponsesDpoGrpc;
import net.tomofiles.skysign.vehicle.api.dpo.UpdateVehicleRequestDpoGrpc;
import net.tomofiles.skysign.vehicle.api.dpo.UpdateVehicleResponseDpoGrpc;
import net.tomofiles.skysign.vehicle.service.ManageVehicleService;
import proto.skysign.DeleteVehicleRequest;
import proto.skysign.common.Empty;
import proto.skysign.GetVehicleRequest;
import proto.skysign.ListVehiclesResponses;
import proto.skysign.common.Vehicle;
import proto.skysign.ManageVehicleServiceGrpc.ManageVehicleServiceImplBase;

@GRpcService
@Controller
@AllArgsConstructor
public class ManageVehicleEndpoint extends ManageVehicleServiceImplBase {

    private final ManageVehicleService service;

    @Override
    public void listVehicles(Empty request, StreamObserver<ListVehiclesResponses> responseObserver) {
        ListVehiclesResponsesDpoGrpc responsesDpo = new ListVehiclesResponsesDpoGrpc();

        try {
            this.service.listVehicles(responsesDpo);
        } catch (Exception e) {
            responseObserver.onError(Status
                    .INTERNAL
                    .withCause(e)
                    .asRuntimeException());
            return;
        }

        responseObserver.onNext(responsesDpo.getGrpcResponse()); 
        responseObserver.onCompleted();
    }

    @Override
    public void getVehicle(GetVehicleRequest request, StreamObserver<Vehicle> responseObserver) {
        GetVehicleRequestDpoGrpc requestDpo = new GetVehicleRequestDpoGrpc(request);
        GetVehicleResponseDpoGrpc responsesDpo = new GetVehicleResponseDpoGrpc();

        try {
            this.service.getVehicle(requestDpo, responsesDpo);
        } catch (Exception e) {
            responseObserver.onError(Status
                    .INTERNAL
                    .withCause(e)
                    .asRuntimeException());
            return;
        }

        if (responsesDpo.isEmpty()) {
            responseObserver.onError(Status
                    .NOT_FOUND
                    .withCause(new NoSuchElementException())
                    .withDescription("vehicle-idに合致するVehicleが存在しません。")
                    .asRuntimeException());
            return;
        }

        responseObserver.onNext(responsesDpo.getGrpcResponse());
        responseObserver.onCompleted();
    }

    @Override
    public void createVehicle(Vehicle request, StreamObserver<Vehicle> responseObserver) {
        CreateVehicleRequestDpoGrpc requestDpo = new CreateVehicleRequestDpoGrpc(request);
        CreateVehicleResponseDpoGrpc responsesDpo = new CreateVehicleResponseDpoGrpc();

        try {
            this.service.createVehicle(requestDpo, responsesDpo);
        } catch (Exception e) {
            responseObserver.onError(Status
                    .INTERNAL
                    .withCause(e)
                    .asRuntimeException());
            return;
        }

        responseObserver.onNext(responsesDpo.getGrpcResponse());
        responseObserver.onCompleted();
    }

    @Override
    public void updateVehicle(Vehicle request, StreamObserver<Vehicle> responseObserver) {
        UpdateVehicleRequestDpoGrpc requestDpo = new UpdateVehicleRequestDpoGrpc(request);
        UpdateVehicleResponseDpoGrpc responsesDpo = new UpdateVehicleResponseDpoGrpc();

        try {
            this.service.updateVehicle(requestDpo, responsesDpo);
        } catch (Exception e) {
            responseObserver.onError(Status
                    .INTERNAL
                    .withCause(e)
                    .asRuntimeException());
            return;
        }

        if (responsesDpo.isEmpty()) {
            responseObserver.onError(Status
                    .NOT_FOUND
                    .withCause(new NoSuchElementException())
                    .withDescription("vehicle-idに合致するVehicleが存在しません。")
                    .asRuntimeException());
            return;
        }

        responseObserver.onNext(responsesDpo.getGrpcResponse());
        responseObserver.onCompleted();
    }

    @Override
    public void deleteVehicle(DeleteVehicleRequest request, StreamObserver<Empty> responseObserver) {
        DeleteVehicleRequestDpoGrpc requestDpo = new DeleteVehicleRequestDpoGrpc(request);
        DeleteVehicleResponseDpoGrpc responsesDpo = new DeleteVehicleResponseDpoGrpc();

        try {
            this.service.deleteVehicle(requestDpo, responsesDpo);
        } catch (Exception e) {
            responseObserver.onError(Status
                    .INTERNAL
                    .withCause(e)
                    .asRuntimeException());
            return;
        }

        if (responsesDpo.isEmpty()) {
            responseObserver.onError(Status
                    .NOT_FOUND
                    .withCause(new NoSuchElementException())
                    .withDescription("vehicle-idに合致するVehicleが存在しません。")
                    .asRuntimeException());
            return;
        }

        responseObserver.onNext(Empty.newBuilder().build()); 
        responseObserver.onCompleted();
    }
}