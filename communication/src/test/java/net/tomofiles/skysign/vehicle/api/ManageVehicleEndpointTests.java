package net.tomofiles.skysign.vehicle.api;

import static com.google.common.truth.Truth.assertThat;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.doThrow;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;
import static org.mockito.MockitoAnnotations.initMocks;

import java.util.Arrays;
import java.util.List;
import java.util.UUID;
import java.util.function.Supplier;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.InjectMocks;
import org.mockito.Mock;

import io.grpc.Status;
import io.grpc.StatusRuntimeException;
import io.grpc.internal.testing.StreamRecorder;
import net.tomofiles.skysign.vehicle.domain.vehicle.Generator;
import net.tomofiles.skysign.vehicle.domain.vehicle.Vehicle;
import net.tomofiles.skysign.vehicle.domain.vehicle.VehicleId;
import net.tomofiles.skysign.vehicle.domain.vehicle.VehicleRepository;
import net.tomofiles.skysign.vehicle.domain.vehicle.Version;
import net.tomofiles.skysign.vehicle.event.Publisher;
import net.tomofiles.skysign.vehicle.service.ManageVehicleService;
import proto.skysign.DeleteVehicleRequest;
import proto.skysign.common.Empty;
import proto.skysign.GetVehicleRequest;
import proto.skysign.ListVehiclesResponses;

import static net.tomofiles.skysign.vehicle.api.GrpcObjectMother.newNormalVehicleGrpc;
import static net.tomofiles.skysign.vehicle.api.GrpcObjectMother.newNoIdVehicleGrpc;
import static net.tomofiles.skysign.vehicle.domain.vehicle.VehicleObjectMother.newNormalVehicle;

public class ManageVehicleEndpointTests {
    
    private static final VehicleId DEFAULT_VEHICLE_ID = new VehicleId(UUID.randomUUID().toString());
    private static final Version DEFAULT_VERSION1 = new Version(UUID.randomUUID().toString());
    private static final Version DEFAULT_VERSION2 = new Version(UUID.randomUUID().toString());
    private static final Supplier<Generator> DEFAULT_GENERATOR = () -> {
        return new Generator(){
            private int count = 0;

            @Override
            public VehicleId newVehicleId() {
                return DEFAULT_VEHICLE_ID;
            }

            @Override
            public Version newVersion() {
                if (count == 0) {
                    count++;
                    return DEFAULT_VERSION1;
                } else {
                    return DEFAULT_VERSION2;
                }
            }
        };
    };

    @Mock
    private VehicleRepository repository;

    @Mock
    private Generator generator;

    @Mock
    private Publisher publisher;

    @InjectMocks
    private ManageVehicleService service;

    private ManageVehicleEndpoint endpoint;

    @BeforeEach
    public void beforeEach() {
        initMocks(this);

        this.endpoint = new ManageVehicleEndpoint(this.service);
    }

    /**
     * ユーザーは、全件取得APIを実行し、すべてのVehicleをリスト形式で取得できる。
     */
    @Test
    public void getAllApi() {
        when(this.repository.getAll()).thenReturn(Arrays.asList(new Vehicle[] {
            newNormalVehicle(DEFAULT_VEHICLE_ID, DEFAULT_VERSION1, DEFAULT_GENERATOR.get()),
            newNormalVehicle(DEFAULT_VEHICLE_ID, DEFAULT_VERSION1, DEFAULT_GENERATOR.get()),
            newNormalVehicle(DEFAULT_VEHICLE_ID, DEFAULT_VERSION1, DEFAULT_GENERATOR.get())
        }));

        Empty request = Empty.newBuilder().build();
        StreamRecorder<ListVehiclesResponses> responseObserver = StreamRecorder.create();
        this.endpoint.listVehicles(request, responseObserver);

        assertThat(responseObserver.getError()).isNull();
        List<ListVehiclesResponses> results = responseObserver.getValues();
        assertThat(results).hasSize(1);
        ListVehiclesResponses response = results.get(0);
        assertThat(response).isEqualTo(ListVehiclesResponses.newBuilder()
                .addVehicles(newNormalVehicleGrpc(DEFAULT_VEHICLE_ID))
                .addVehicles(newNormalVehicleGrpc(DEFAULT_VEHICLE_ID))
                .addVehicles(newNormalVehicleGrpc(DEFAULT_VEHICLE_ID))
                .build());
    }

    /**
     * ユーザーは、全件取得APIを実行し、未存在により空のリストを取得できる。
     */
    @Test
    public void getAllApiNotFoundError() {
        Empty request = Empty.newBuilder().build();
        StreamRecorder<ListVehiclesResponses> responseObserver = StreamRecorder.create();
        this.endpoint.listVehicles(request, responseObserver);

        assertThat(responseObserver.getError()).isNull();
        List<ListVehiclesResponses> results = responseObserver.getValues();
        assertThat(results).hasSize(1);
        ListVehiclesResponses response = results.get(0);
        assertThat(response).isEqualTo(ListVehiclesResponses.newBuilder().build());
    }

    /**
     * ユーザーは、全件取得APIを実行し、DBエラーのよりINTERNALエラーを検出できる。
     */
    @Test
    public void getAllApiInternalError() {
        when(this.repository.getAll()).thenThrow(new IllegalStateException());

        Empty request = Empty.newBuilder().build();
        StreamRecorder<ListVehiclesResponses> responseObserver = StreamRecorder.create();
        this.endpoint.listVehicles(request, responseObserver);

        assertThat(responseObserver.getError()).isNotNull();
        assertThat(responseObserver.getError()).isInstanceOf(StatusRuntimeException.class);
        assertThat(((StatusRuntimeException)responseObserver.getError()).getStatus().getCode())
                .isEqualTo(Status.INTERNAL.getCode());
    }

    /**
     * ユーザーは、１件取得APIを実行し、対象のVehicleを取得できる。
     */
    @Test
    public void getOneApi() {
        when(this.repository.getById(DEFAULT_VEHICLE_ID))
                .thenReturn(newNormalVehicle(
                        DEFAULT_VEHICLE_ID,
                        DEFAULT_VERSION1,
                        DEFAULT_GENERATOR.get()));

        GetVehicleRequest request = GetVehicleRequest.newBuilder()
                .setId(DEFAULT_VEHICLE_ID.getId())
                .build();
        StreamRecorder<proto.skysign.common.Vehicle> responseObserver = StreamRecorder.create();
        this.endpoint.getVehicle(request, responseObserver);

        assertThat(responseObserver.getError()).isNull();
        List<proto.skysign.common.Vehicle> results = responseObserver.getValues();
        assertThat(results).hasSize(1);
        proto.skysign.common.Vehicle response = results.get(0);
        assertThat(response).isEqualTo(newNormalVehicleGrpc(DEFAULT_VEHICLE_ID));
    }

    /**
     * ユーザーは、１件取得APIを実行し、未存在のID指定によりNOT_FOUNDエラーを検出できる。
     */
    @Test
    public void getOneApiNotFoundError() {
        GetVehicleRequest request = GetVehicleRequest.newBuilder()
                .setId(DEFAULT_VEHICLE_ID.getId())
                .build();
        StreamRecorder<proto.skysign.common.Vehicle> responseObserver = StreamRecorder.create();
        this.endpoint.getVehicle(request, responseObserver);

        assertThat(responseObserver.getError()).isNotNull();
        assertThat(responseObserver.getError()).isInstanceOf(StatusRuntimeException.class);
        assertThat(((StatusRuntimeException)responseObserver.getError()).getStatus().getCode())
                .isEqualTo(Status.NOT_FOUND.getCode());
    }

    /**
     * ユーザーは、１件取得APIを実行し、DBエラーのよりINTERNALエラーを検出できる。
     */
    @Test
    public void getOneApiInternalError() {
        when(this.repository.getById(DEFAULT_VEHICLE_ID)).thenThrow(new IllegalStateException());

        GetVehicleRequest request = GetVehicleRequest.newBuilder()
                .setId(DEFAULT_VEHICLE_ID.getId())
                .build();
        StreamRecorder<proto.skysign.common.Vehicle> responseObserver = StreamRecorder.create();
        this.endpoint.getVehicle(request, responseObserver);

        assertThat(responseObserver.getError()).isNotNull();
        assertThat(responseObserver.getError()).isInstanceOf(StatusRuntimeException.class);
        assertThat(((StatusRuntimeException)responseObserver.getError()).getStatus().getCode())
                .isEqualTo(Status.INTERNAL.getCode());
    }

    /**
     * ユーザーは、登録APIを実行し、新しいVehicleを登録できる。
     */
    @Test
    public void createApi() {
        when(this.generator.newVehicleId()).thenReturn(DEFAULT_VEHICLE_ID);
        when(this.generator.newVersion()).thenReturn(DEFAULT_VERSION1);

        proto.skysign.common.Vehicle request = newNoIdVehicleGrpc();
        StreamRecorder<proto.skysign.common.Vehicle> responseObserver = StreamRecorder.create();
        this.endpoint.createVehicle(request, responseObserver);

        verify(this.repository, times(1)).save(newNormalVehicle(
                DEFAULT_VEHICLE_ID,
                DEFAULT_VERSION1,
                DEFAULT_GENERATOR.get()
        ));

        assertThat(responseObserver.getError()).isNull();
        List<proto.skysign.common.Vehicle> results = responseObserver.getValues();
        assertThat(results).hasSize(1);
        proto.skysign.common.Vehicle response = results.get(0);
        assertThat(response).isEqualTo(newNormalVehicleGrpc(DEFAULT_VEHICLE_ID));
    }

    /**
     * ユーザーは、登録APIを実行し、DBエラーのよりINTERNALエラーを検出できる。
     */
    @Test
    public void createApiInternalError() {
        doThrow(new IllegalStateException()).when(this.repository).save(any());

        proto.skysign.common.Vehicle request = newNoIdVehicleGrpc();
        StreamRecorder<proto.skysign.common.Vehicle> responseObserver = StreamRecorder.create();
        this.endpoint.createVehicle(request, responseObserver);

        assertThat(responseObserver.getError()).isNotNull();
        assertThat(responseObserver.getError()).isInstanceOf(StatusRuntimeException.class);
        assertThat(((StatusRuntimeException)responseObserver.getError()).getStatus().getCode())
                .isEqualTo(Status.INTERNAL.getCode());
    }

    /**
     * ユーザーは、更新APIを実行し、既存のVehicleを更新できる。
     */
    @Test
    public void updateApi() {
        when(this.repository.getById(DEFAULT_VEHICLE_ID))
                .thenReturn(newNormalVehicle(
                        DEFAULT_VEHICLE_ID,
                        DEFAULT_VERSION1,
                        DEFAULT_GENERATOR.get()));

        when(this.generator.newVehicleId()).thenReturn(DEFAULT_VEHICLE_ID);
        when(this.generator.newVersion()).thenReturn(DEFAULT_VERSION1);

        proto.skysign.common.Vehicle request = newNormalVehicleGrpc(DEFAULT_VEHICLE_ID);
        StreamRecorder<proto.skysign.common.Vehicle> responseObserver = StreamRecorder.create();
        this.endpoint.updateVehicle(request, responseObserver);

        verify(this.repository, times(1)).save(newNormalVehicle(
                DEFAULT_VEHICLE_ID,
                DEFAULT_VERSION1,
                DEFAULT_GENERATOR.get()
        ));

        assertThat(responseObserver.getError()).isNull();
        List<proto.skysign.common.Vehicle> results = responseObserver.getValues();
        assertThat(results).hasSize(1);
        proto.skysign.common.Vehicle response = results.get(0);
        assertThat(response).isEqualTo(newNormalVehicleGrpc(DEFAULT_VEHICLE_ID));
    }

    /**
     * ユーザーは、更新APIを実行し、未存在のID指定によりNOT_FOUNDエラーを検出できる。
     */
    @Test
    public void updateApiNotFoundError() {
        proto.skysign.common.Vehicle request = newNormalVehicleGrpc(DEFAULT_VEHICLE_ID);
        StreamRecorder<proto.skysign.common.Vehicle> responseObserver = StreamRecorder.create();
        this.endpoint.updateVehicle(request, responseObserver);

        assertThat(responseObserver.getError()).isNotNull();
        assertThat(responseObserver.getError()).isInstanceOf(StatusRuntimeException.class);
        assertThat(((StatusRuntimeException)responseObserver.getError()).getStatus().getCode())
                .isEqualTo(Status.NOT_FOUND.getCode());
    }

    /**
     * ユーザーは、更新APIを実行し、DBエラーのよりINTERNALエラーを検出できる。
     */
    @Test
    public void updateApiInternalError() {
        when(this.repository.getById(DEFAULT_VEHICLE_ID))
                .thenReturn(newNormalVehicle(
                        DEFAULT_VEHICLE_ID,
                        DEFAULT_VERSION1,
                        DEFAULT_GENERATOR.get()));
        doThrow(new IllegalStateException()).when(this.repository).save(any());

        proto.skysign.common.Vehicle request = newNormalVehicleGrpc(DEFAULT_VEHICLE_ID);
        StreamRecorder<proto.skysign.common.Vehicle> responseObserver = StreamRecorder.create();
        this.endpoint.updateVehicle(request, responseObserver);

        assertThat(responseObserver.getError()).isNotNull();
        assertThat(responseObserver.getError()).isInstanceOf(StatusRuntimeException.class);
        assertThat(((StatusRuntimeException)responseObserver.getError()).getStatus().getCode())
                .isEqualTo(Status.INTERNAL.getCode());
    }

    /**
     * ユーザーは、削除APIを実行し、既存のVehicleを削除できる。
     */
    @Test
    public void deleteApi() {
        when(this.repository.getById(DEFAULT_VEHICLE_ID))
                .thenReturn(newNormalVehicle(
                        DEFAULT_VEHICLE_ID,
                        DEFAULT_VERSION1,
                        DEFAULT_GENERATOR.get()));

        DeleteVehicleRequest request = DeleteVehicleRequest.newBuilder()
                .setId(DEFAULT_VEHICLE_ID.getId())
                .build();
        StreamRecorder<Empty> responseObserver = StreamRecorder.create();
        this.endpoint.deleteVehicle(request, responseObserver);

        verify(this.repository, times(1)).remove(DEFAULT_VEHICLE_ID, DEFAULT_VERSION1);

        assertThat(responseObserver.getError()).isNull();
        List<Empty> results = responseObserver.getValues();
        assertThat(results).hasSize(1);
    }

    /**
     * ユーザーは、削除APIを実行し、未存在のID指定によりNOT_FOUNDエラーを検出できる。
     */
    @Test
    public void deleteApiNotFoundError() {
        DeleteVehicleRequest request = DeleteVehicleRequest.newBuilder()
                .setId(DEFAULT_VEHICLE_ID.getId())
                .build();
        StreamRecorder<Empty> responseObserver = StreamRecorder.create();
        this.endpoint.deleteVehicle(request, responseObserver);

        assertThat(responseObserver.getError()).isNotNull();
        assertThat(responseObserver.getError()).isInstanceOf(StatusRuntimeException.class);
        assertThat(((StatusRuntimeException)responseObserver.getError()).getStatus().getCode())
                .isEqualTo(Status.NOT_FOUND.getCode());
    }

    /**
     * ユーザーは、削除APIを実行し、DBエラーのよりINTERNALエラーを検出できる。
     */
    @Test
    public void deleteApiInternalError() {
        when(this.repository.getById(DEFAULT_VEHICLE_ID))
                .thenReturn(newNormalVehicle(
                        DEFAULT_VEHICLE_ID,
                        DEFAULT_VERSION1,
                        DEFAULT_GENERATOR.get()));
        doThrow(new IllegalStateException()).when(this.repository).remove(any(), any());

        DeleteVehicleRequest request = DeleteVehicleRequest.newBuilder()
                .setId(DEFAULT_VEHICLE_ID.getId())
                .build();
        StreamRecorder<Empty> responseObserver = StreamRecorder.create();
        this.endpoint.deleteVehicle(request, responseObserver);

        assertThat(responseObserver.getError()).isNotNull();
        assertThat(responseObserver.getError()).isInstanceOf(StatusRuntimeException.class);
        assertThat(((StatusRuntimeException)responseObserver.getError()).getStatus().getCode())
                .isEqualTo(Status.INTERNAL.getCode());
    }
}