package net.tomofiles.skysign.communication.api;

import static com.google.common.truth.Truth.assertThat;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;
import static org.mockito.MockitoAnnotations.initMocks;

import java.time.LocalDateTime;
import java.util.Arrays;
import java.util.List;
import java.util.UUID;
import java.util.function.Supplier;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.ArgumentCaptor;
import org.mockito.InjectMocks;
import org.mockito.Mock;

import io.grpc.Status;
import io.grpc.StatusRuntimeException;
import io.grpc.internal.testing.StreamRecorder;
import net.tomofiles.skysign.communication.domain.communication.CommandId;
import net.tomofiles.skysign.communication.domain.communication.Communication;
import net.tomofiles.skysign.communication.domain.communication.CommunicationFactory;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.domain.communication.CommunicationRepository;
import net.tomofiles.skysign.communication.domain.communication.Generator;
import net.tomofiles.skysign.communication.domain.communication.MissionId;
import net.tomofiles.skysign.communication.domain.communication.component.CommunicationComponentDto;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleId;
import net.tomofiles.skysign.communication.service.CommunicationUserService;
import proto.skysign.CancelRequest;
import proto.skysign.CancelResponse;
import proto.skysign.common.CommandType;
import proto.skysign.common.Empty;
import proto.skysign.ListCommunicationsResponses;
import proto.skysign.PullTelemetryRequest;
import proto.skysign.PullTelemetryResponse;
import proto.skysign.PushCommandRequest;
import proto.skysign.PushCommandResponse;
import proto.skysign.StagingRequest;
import proto.skysign.StagingResponse;

import static net.tomofiles.skysign.communication.api.GrpcObjectMother.newNormalCommunicationGrpc;
import static net.tomofiles.skysign.communication.api.GrpcObjectMother.newNormalPullTelemetryResponseGrpc;
import static net.tomofiles.skysign.communication.domain.communication.CommunicationObjectMother.newNormalCommunication;
import static net.tomofiles.skysign.communication.domain.communication.ComponentDtoObjectMother.newSingleCommandComponentDto;

public class CommunicationUserEndpointTests {
    
    private static final CommunicationId DEFAULT_COMMUNICATION_ID = new CommunicationId(UUID.randomUUID().toString());
    private static final CommandId DEFAULT_COMMAND_ID = new CommandId(UUID.randomUUID().toString());
    private static final String DEFAULT_COMMAND_TYPE = "ARM";
    private static final VehicleId DEFAULT_VEHICLE_ID = new VehicleId(UUID.randomUUID().toString());
    private static final MissionId DEFAULT_MISSION_ID = new MissionId(UUID.randomUUID().toString());
    private static final LocalDateTime DEFAULT_COMMAND_TIME = LocalDateTime.of(2020, 1, 1, 0, 0, 0);
    private static final Supplier<Generator> DEFAULT_GENERATOR = () -> {
        return new Generator(){
            @Override
            public CommandId newCommandId() {
                return DEFAULT_COMMAND_ID;
            }
            @Override
            public LocalDateTime newTime() {
                return DEFAULT_COMMAND_TIME;
            }
        };
    };

    @Mock
    private CommunicationRepository repository;

    @InjectMocks
    private CommunicationUserService service;

    private CommunicationUserEndpoint endpoint;

    @BeforeEach
    public void beforeEach() {
        initMocks(this);

        endpoint = new CommunicationUserEndpoint(service);
    }

    /**
     * ユーザーは、全件取得APIを実行し、すべてのCommunicationをリスト形式で取得できる。
     */
    @Test
    public void listCommunicationsApi() {
        when(repository.getAll())
                .thenReturn(Arrays.asList(new Communication[] {
                        newNormalCommunication(
                                DEFAULT_COMMUNICATION_ID,
                                DEFAULT_VEHICLE_ID,
                                DEFAULT_MISSION_ID,
                                DEFAULT_GENERATOR.get()),
                        newNormalCommunication(
                                DEFAULT_COMMUNICATION_ID,
                                DEFAULT_VEHICLE_ID,
                                DEFAULT_MISSION_ID,
                                DEFAULT_GENERATOR.get()),
                        newNormalCommunication(
                                DEFAULT_COMMUNICATION_ID,
                                DEFAULT_VEHICLE_ID,
                                DEFAULT_MISSION_ID,
                                DEFAULT_GENERATOR.get())
                }));

        Empty request = Empty.newBuilder().build();
        StreamRecorder<ListCommunicationsResponses> responseObserver = StreamRecorder.create();
        endpoint.listCommunications(request, responseObserver);

        assertThat(responseObserver.getError()).isNull();
        List<ListCommunicationsResponses> results = responseObserver.getValues();
        assertThat(results).hasSize(1);
        ListCommunicationsResponses response = results.get(0);
        assertThat(response).isEqualTo(ListCommunicationsResponses.newBuilder()
                .addCommunications(newNormalCommunicationGrpc(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_VEHICLE_ID,
                        DEFAULT_MISSION_ID))
                .addCommunications(newNormalCommunicationGrpc(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_VEHICLE_ID,
                        DEFAULT_MISSION_ID))
                .addCommunications(newNormalCommunicationGrpc(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_VEHICLE_ID,
                        DEFAULT_MISSION_ID))
                .build());
    }

    /**
     * ユーザーは、全件取得APIを実行し、未存在により空のリストを取得できる。
     */
    @Test
    public void listCommunicationsApiNotFoundError() {
        Empty request = Empty.newBuilder().build();
        StreamRecorder<ListCommunicationsResponses> responseObserver = StreamRecorder.create();
        endpoint.listCommunications(request, responseObserver);

        assertThat(responseObserver.getError()).isNull();
        List<ListCommunicationsResponses> results = responseObserver.getValues();
        assertThat(results).hasSize(1);
        ListCommunicationsResponses response = results.get(0);
        assertThat(response).isEqualTo(ListCommunicationsResponses.newBuilder().build());
    }

    /**
     * ユーザーは、全件取得APIを実行し、DBエラーのよりINTERNALエラーを検出できる。
     */
    @Test
    public void listCommunicationsApiInternalError() {
        when(repository.getAll()).thenThrow(new IllegalStateException());

        Empty request = Empty.newBuilder().build();
        StreamRecorder<ListCommunicationsResponses> responseObserver = StreamRecorder.create();
        endpoint.listCommunications(request, responseObserver);

        assertThat(responseObserver.getError()).isNotNull();
        assertThat(responseObserver.getError()).isInstanceOf(StatusRuntimeException.class);
        assertThat(((StatusRuntimeException)responseObserver.getError()).getStatus().getCode())
                .isEqualTo(Status.INTERNAL.getCode());
    }

    /**
     * ユーザーは、飛行コマンド送信APIを実行し、対象のCommunicationにコマンドを送信できる。
     */
    @Test
    public void pushCommandApi() {
        when(repository.getById(DEFAULT_COMMUNICATION_ID))
                .thenReturn(newNormalCommunication(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_VEHICLE_ID,
                        DEFAULT_MISSION_ID,
                        DEFAULT_GENERATOR.get()));

        PushCommandRequest request = PushCommandRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .setType(CommandType.valueOf(DEFAULT_COMMAND_TYPE))
                .build();
        StreamRecorder<PushCommandResponse> responseObserver = StreamRecorder.create();
        endpoint.pushCommand(request, responseObserver);

        ArgumentCaptor<Communication> commCaptor = ArgumentCaptor.forClass(Communication.class);
        verify(repository, times(1)).save(commCaptor.capture());

        CommunicationComponentDto dto = CommunicationFactory.takeApart(commCaptor.getValue());
        assertThat(dto.getCommands()).hasSize(1);
        assertThat(dto.getCommands().get(0))
                .isEqualTo(newSingleCommandComponentDto(DEFAULT_GENERATOR.get()));

        assertThat(responseObserver.getError()).isNull();
        List<PushCommandResponse> results = responseObserver.getValues();
        assertThat(results).hasSize(1);
        PushCommandResponse response = results.get(0);
        assertThat(response).isEqualTo(PushCommandResponse.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .setType(CommandType.valueOf(DEFAULT_COMMAND_TYPE))
                .build());
    }

    /**
     * ユーザーは、飛行コマンド送信APIを実行し、未存在のID指定によりNOT_FOUNDエラーを検出できる。
     */
    @Test
    public void pushCommandApiNotFoundError() {
        PushCommandRequest request = PushCommandRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .setType(CommandType.valueOf(DEFAULT_COMMAND_TYPE))
                .build();
        StreamRecorder<PushCommandResponse> responseObserver = StreamRecorder.create();
        endpoint.pushCommand(request, responseObserver);

        assertThat(responseObserver.getError()).isNotNull();
        assertThat(responseObserver.getError()).isInstanceOf(StatusRuntimeException.class);
        assertThat(((StatusRuntimeException)responseObserver.getError()).getStatus().getCode())
                .isEqualTo(Status.NOT_FOUND.getCode());
    }

    /**
     * ユーザーは、飛行コマンド送信APIを実行し、DBエラーのよりINTERNALエラーを検出できる。
     */
    @Test
    public void pushCommandApiInternalError() {
        when(repository.getById(DEFAULT_COMMUNICATION_ID)).thenThrow(new IllegalStateException());

        PushCommandRequest request = PushCommandRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .setType(CommandType.valueOf(DEFAULT_COMMAND_TYPE))
                .build();
        StreamRecorder<PushCommandResponse> responseObserver = StreamRecorder.create();
        endpoint.pushCommand(request, responseObserver);

        assertThat(responseObserver.getError()).isNotNull();
        assertThat(responseObserver.getError()).isInstanceOf(StatusRuntimeException.class);
        assertThat(((StatusRuntimeException)responseObserver.getError()).getStatus().getCode())
                .isEqualTo(Status.INTERNAL.getCode());
    }

    /**
     * ユーザーは、テレメトリー受信APIを実行し、対象のCommunicationからテレメトリーを受信できる。
     */
    @Test
    public void pullTelemetryApi() {
        when(repository.getById(DEFAULT_COMMUNICATION_ID))
                .thenReturn(newNormalCommunication(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_VEHICLE_ID,
                        DEFAULT_MISSION_ID,
                        DEFAULT_GENERATOR.get()));

        PullTelemetryRequest request = PullTelemetryRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .build();
        StreamRecorder<PullTelemetryResponse> responseObserver = StreamRecorder.create();
        endpoint.pullTelemetry(request, responseObserver);

        assertThat(responseObserver.getError()).isNull();
        List<PullTelemetryResponse> results = responseObserver.getValues();
        assertThat(results).hasSize(1);
        PullTelemetryResponse response = results.get(0);
        assertThat(response).isEqualTo(newNormalPullTelemetryResponseGrpc(DEFAULT_COMMUNICATION_ID));
    }

    /**
     * ユーザーは、テレメトリー受信APIを実行し、未存在のID指定によりNOT_FOUNDエラーを検出できる。
     */
    @Test
    public void pullTelemetryApiNotFoundError() {
        PullTelemetryRequest request = PullTelemetryRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .build();
        StreamRecorder<PullTelemetryResponse> responseObserver = StreamRecorder.create();
        endpoint.pullTelemetry(request, responseObserver);

        assertThat(responseObserver.getError()).isNotNull();
        assertThat(responseObserver.getError()).isInstanceOf(StatusRuntimeException.class);
        assertThat(((StatusRuntimeException)responseObserver.getError()).getStatus().getCode())
                .isEqualTo(Status.NOT_FOUND.getCode());
    }

    /**
     * ユーザーは、テレメトリー受信APIを実行し、DBエラーのよりINTERNALエラーを検出できる。
     */
    @Test
    public void pullTelemetryApiInternalError() {
        when(repository.getById(DEFAULT_COMMUNICATION_ID)).thenThrow(new IllegalStateException());

        PullTelemetryRequest request = PullTelemetryRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .build();
        StreamRecorder<PullTelemetryResponse> responseObserver = StreamRecorder.create();
        endpoint.pullTelemetry(request, responseObserver);

        assertThat(responseObserver.getError()).isNotNull();
        assertThat(responseObserver.getError()).isInstanceOf(StatusRuntimeException.class);
        assertThat(((StatusRuntimeException)responseObserver.getError()).getStatus().getCode())
                .isEqualTo(Status.INTERNAL.getCode());
    }

    /**
     * ユーザーは、ステージングAPIを実行し、対象のCommunicationにMissionIdを紐付ける。
     */
    @Test
    public void stagingApi() {
        MissionId newMissionId = new MissionId("new mission id");
        when(repository.getById(DEFAULT_COMMUNICATION_ID))
                .thenReturn(newNormalCommunication(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_VEHICLE_ID,
                        DEFAULT_MISSION_ID,
                        DEFAULT_GENERATOR.get()));

        StagingRequest request = StagingRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .setMissionId(newMissionId.getId())
                .build();
        StreamRecorder<StagingResponse> responseObserver = StreamRecorder.create();
        endpoint.staging(request, responseObserver);

        ArgumentCaptor<Communication> commCaptor = ArgumentCaptor.forClass(Communication.class);
        verify(repository, times(1)).save(commCaptor.capture());

        assertThat(commCaptor.getValue().getMissionId())
                .isEqualTo(newMissionId);

        assertThat(responseObserver.getError()).isNull();
        List<StagingResponse> results = responseObserver.getValues();
        assertThat(results).hasSize(1);
        StagingResponse response = results.get(0);
        assertThat(response).isEqualTo(StagingResponse.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .setMissionId(newMissionId.getId())
                .build());
    }

    /**
     * ユーザーは、ステージングAPIを実行し、未存在のID指定によりNOT_FOUNDエラーを検出できる。
     */
    @Test
    public void stagingApiNotFoundError() {
        StagingRequest request = StagingRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .setMissionId(DEFAULT_MISSION_ID.getId())
                .build();
        StreamRecorder<StagingResponse> responseObserver = StreamRecorder.create();
        endpoint.staging(request, responseObserver);

        assertThat(responseObserver.getError()).isNotNull();
        assertThat(responseObserver.getError()).isInstanceOf(StatusRuntimeException.class);
        assertThat(((StatusRuntimeException)responseObserver.getError()).getStatus().getCode())
                .isEqualTo(Status.NOT_FOUND.getCode());
    }

    /**
     * ユーザーは、ステージングAPIを実行し、DBエラーのよりINTERNALエラーを検出できる。
     */
    @Test
    public void stagingApiInternalError() {
        when(repository.getById(DEFAULT_COMMUNICATION_ID)).thenThrow(new IllegalStateException());

        StagingRequest request = StagingRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .setMissionId(DEFAULT_MISSION_ID.getId())
                .build();
        StreamRecorder<StagingResponse> responseObserver = StreamRecorder.create();
        endpoint.staging(request, responseObserver);

        assertThat(responseObserver.getError()).isNotNull();
        assertThat(responseObserver.getError()).isInstanceOf(StatusRuntimeException.class);
        assertThat(((StatusRuntimeException)responseObserver.getError()).getStatus().getCode())
                .isEqualTo(Status.INTERNAL.getCode());
    }

    /**
     * ユーザーは、キャンセルAPIを実行し、対象のCommunicationのMissionIdの紐付けを解除する。
     */
    @Test
    public void cancelApi() {
        when(repository.getById(DEFAULT_COMMUNICATION_ID))
                .thenReturn(newNormalCommunication(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_VEHICLE_ID,
                        DEFAULT_MISSION_ID,
                        DEFAULT_GENERATOR.get()));

        CancelRequest request = CancelRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .build();
        StreamRecorder<CancelResponse> responseObserver = StreamRecorder.create();
        endpoint.cancel(request, responseObserver);

        ArgumentCaptor<Communication> commCaptor = ArgumentCaptor.forClass(Communication.class);
        verify(repository, times(1)).save(commCaptor.capture());

        assertThat(commCaptor.getValue().getMissionId())
                .isNull();;

        assertThat(responseObserver.getError()).isNull();
        List<CancelResponse> results = responseObserver.getValues();
        assertThat(results).hasSize(1);
        CancelResponse response = results.get(0);
        assertThat(response).isEqualTo(CancelResponse.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .build());
    }

    /**
     * ユーザーは、キャンセルAPIを実行し、未存在のID指定によりNOT_FOUNDエラーを検出できる。
     */
    @Test
    public void cancelApiNotFoundError() {
        CancelRequest request = CancelRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .build();
        StreamRecorder<CancelResponse> responseObserver = StreamRecorder.create();
        endpoint.cancel(request, responseObserver);

        assertThat(responseObserver.getError()).isNotNull();
        assertThat(responseObserver.getError()).isInstanceOf(StatusRuntimeException.class);
        assertThat(((StatusRuntimeException)responseObserver.getError()).getStatus().getCode())
                .isEqualTo(Status.NOT_FOUND.getCode());
    }

    /**
     * ユーザーは、キャンセルAPIを実行し、DBエラーのよりINTERNALエラーを検出できる。
     */
    @Test
    public void cancelApiInternalError() {
        when(repository.getById(DEFAULT_COMMUNICATION_ID)).thenThrow(new IllegalStateException());

        CancelRequest request = CancelRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .build();
        StreamRecorder<CancelResponse> responseObserver = StreamRecorder.create();
        endpoint.cancel(request, responseObserver);

        assertThat(responseObserver.getError()).isNotNull();
        assertThat(responseObserver.getError()).isInstanceOf(StatusRuntimeException.class);
        assertThat(((StatusRuntimeException)responseObserver.getError()).getStatus().getCode())
                .isEqualTo(Status.INTERNAL.getCode());
    }
}