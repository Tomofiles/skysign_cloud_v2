package net.tomofiles.skysign.communication.api;

import java.util.NoSuchElementException;

import org.lognet.springboot.grpc.GRpcService;
import org.springframework.stereotype.Controller;

import io.grpc.Status;
import io.grpc.stub.StreamObserver;
import lombok.AllArgsConstructor;
import net.tomofiles.skysign.communication.api.dpo.CreateVehicleRequestDpoGrpc;
import net.tomofiles.skysign.communication.api.dpo.CreateVehicleResponseDpoGrpc;
import net.tomofiles.skysign.communication.api.dpo.DeleteVehicleRequestDpoGrpc;
import net.tomofiles.skysign.communication.api.dpo.DeleteVehicleResponseDpoGrpc;
import net.tomofiles.skysign.communication.api.dpo.GetVehicleRequestDpoGrpc;
import net.tomofiles.skysign.communication.api.dpo.GetVehicleResponseDpoGrpc;
import net.tomofiles.skysign.communication.api.dpo.ListVehiclesResponsesDpoGrpc;
import net.tomofiles.skysign.communication.api.dpo.UpdateVehicleRequestDpoGrpc;
import net.tomofiles.skysign.communication.api.dpo.UpdateVehicleResponseDpoGrpc;
import net.tomofiles.skysign.communication.usecase.ManageVehicleService;
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
    public void createVehicle(CreateVehicleRequest request, StreamObserver<Vehicle> responseObserver) {
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
    public void updateVehicle(UpdateVehicleRequest request, StreamObserver<Vehicle> responseObserver) {
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