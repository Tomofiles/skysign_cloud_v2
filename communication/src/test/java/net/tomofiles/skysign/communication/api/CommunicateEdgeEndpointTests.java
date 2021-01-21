package net.tomofiles.skysign.communication.api;

import static com.google.common.truth.Truth.assertThat;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;
import static org.mockito.MockitoAnnotations.initMocks;

import java.time.LocalDateTime;
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
import net.tomofiles.skysign.communication.service.CommunicateEdgeService;
import proto.skysign.common.CommandType;
import proto.skysign.GetCommunicationRequest;
import proto.skysign.PullCommandRequest;
import proto.skysign.PullCommandResponse;
import proto.skysign.PushTelemetryRequest;
import proto.skysign.PushTelemetryResponse;

import static net.tomofiles.skysign.communication.api.GrpcObjectMother.newNormalTelemetryGrpc;
import static net.tomofiles.skysign.communication.domain.communication.CommunicationObjectMother.newNormalCommunication;
import static net.tomofiles.skysign.communication.domain.communication.CommunicationObjectMother.newSingleCommandCommunication;
import static net.tomofiles.skysign.communication.domain.communication.ComponentDtoObjectMother.newNormalTelemetryComponentDto;

public class CommunicateEdgeEndpointTests {
    
    private static final CommunicationId DEFAULT_COMMUNICATION_ID = new CommunicationId(UUID.randomUUID().toString());
    private static final CommandId DEFAULT_COMMAND_ID = new CommandId(UUID.randomUUID().toString());
    private static final String DEFAULT_COMMAND_TYPE = "ARM";
    private static final VehicleId DEFAULT_VEHICLE_ID = new VehicleId(UUID.randomUUID().toString());
    private static final boolean DEFAULT_CONTROLLED = true;
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
    private CommunicateEdgeService service;

    private CommunicateEdgeEndpoint endpoint;

    @BeforeEach
    public void beforeEach() {
        initMocks(this);

        endpoint = new CommunicateEdgeEndpoint(service);
    }

    /**
     * エッジは、テレメトリー送信APIを実行し、対象のCommunicationにテレメトリーを送信できる。<br>
     * Communicationが保持するコマンドは0件のため、返却オブジェクトのコマンドリストは0件であること。
     */
    @Test
    public void pushTelemetryWithNoCommandApi() {
        when(repository.getById(DEFAULT_COMMUNICATION_ID))
                .thenReturn(CommunicationFactory.newInstance(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_VEHICLE_ID,
                        DEFAULT_GENERATOR.get()));

        PushTelemetryRequest request = PushTelemetryRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .setTelemetry(newNormalTelemetryGrpc())
                .build();
        StreamRecorder<PushTelemetryResponse> responseObserver = StreamRecorder.create();
        endpoint.pushTelemetry(request, responseObserver);

        ArgumentCaptor<Communication> commCaptor = ArgumentCaptor.forClass(Communication.class);
        verify(repository, times(1)).save(commCaptor.capture());

        CommunicationComponentDto dto = CommunicationFactory.takeApart(commCaptor.getValue());
        assertThat(dto.getTelemetry())
                .isEqualTo(newNormalTelemetryComponentDto());

        assertThat(responseObserver.getError()).isNull();
        List<PushTelemetryResponse> results = responseObserver.getValues();
        assertThat(results).hasSize(1);
        PushTelemetryResponse response = results.get(0);
        assertThat(response).isEqualTo(PushTelemetryResponse.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                // Commandが0件
                .build());
    }

    /**
     * エッジは、テレメトリー送信APIを実行し、対象のCommunicationにテレメトリーを送信できる。<br>
     * Communicationが保持するコマンドは1件のため、返却オブジェクトのコマンドリストは1件であること。
     */
    @Test
    public void pushTelemetryWithOneCommandApi() {
        when(repository.getById(DEFAULT_COMMUNICATION_ID))
                .thenReturn(newSingleCommandCommunication(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_VEHICLE_ID,
                        DEFAULT_CONTROLLED,
                        DEFAULT_MISSION_ID,
                        DEFAULT_GENERATOR.get()));

        PushTelemetryRequest request = PushTelemetryRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .setTelemetry(newNormalTelemetryGrpc())
                .build();
        StreamRecorder<PushTelemetryResponse> responseObserver = StreamRecorder.create();
        endpoint.pushTelemetry(request, responseObserver);

        ArgumentCaptor<Communication> commCaptor = ArgumentCaptor.forClass(Communication.class);
        verify(repository, times(1)).save(commCaptor.capture());

        CommunicationComponentDto dto = CommunicationFactory.takeApart(commCaptor.getValue());
        assertThat(dto.getTelemetry())
                .isEqualTo(newNormalTelemetryComponentDto());

        assertThat(responseObserver.getError()).isNull();
        List<PushTelemetryResponse> results = responseObserver.getValues();
        assertThat(results).hasSize(1);
        PushTelemetryResponse response = results.get(0);
        assertThat(response).isEqualTo(PushTelemetryResponse.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .addCommIds(DEFAULT_COMMAND_ID.getId())
                .build());
    }

    /**
     * エッジは、テレメトリー送信APIを実行し、未存在のID指定によりNOT_FOUNDエラーを検出できる。
     */
    @Test
    public void pushTelemetryApiNotFoundError() {
        PushTelemetryRequest request = PushTelemetryRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .setTelemetry(newNormalTelemetryGrpc())
                .build();
        StreamRecorder<PushTelemetryResponse> responseObserver = StreamRecorder.create();
        endpoint.pushTelemetry(request, responseObserver);

        assertThat(responseObserver.getError()).isNotNull();
        assertThat(responseObserver.getError()).isInstanceOf(StatusRuntimeException.class);
        assertThat(((StatusRuntimeException)responseObserver.getError()).getStatus().getCode())
                .isEqualTo(Status.NOT_FOUND.getCode());
    }

    /**
     * エッジは、テレメトリー送信APIを実行し、DBエラーのよりINTERNALエラーを検出できる。
     */
    @Test
    public void pushTelemetryApiInternalError() {
        when(repository.getById(DEFAULT_COMMUNICATION_ID)).thenThrow(new IllegalStateException());

        PushTelemetryRequest request = PushTelemetryRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .setTelemetry(newNormalTelemetryGrpc())
                .build();
        StreamRecorder<PushTelemetryResponse> responseObserver = StreamRecorder.create();
        endpoint.pushTelemetry(request, responseObserver);

        assertThat(responseObserver.getError()).isNotNull();
        assertThat(responseObserver.getError()).isInstanceOf(StatusRuntimeException.class);
        assertThat(((StatusRuntimeException)responseObserver.getError()).getStatus().getCode())
                .isEqualTo(Status.INTERNAL.getCode());
    }

    /**
     * エッジは、飛行コマンド受信APIを実行し、対象のCommunicationからコマンドを受信できる。<br>
     * Communicationからコマンドを取得すると、エンティティ内から削除されることを検証する。
     */
    @Test
    public void pullCommandApi() {
        when(repository.getById(DEFAULT_COMMUNICATION_ID))
                .thenReturn(newSingleCommandCommunication(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_VEHICLE_ID,
                        DEFAULT_CONTROLLED,
                        DEFAULT_MISSION_ID,
                        DEFAULT_GENERATOR.get()));

        PullCommandRequest request = PullCommandRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .setCommandId(DEFAULT_COMMAND_ID.getId())
                .build();
        StreamRecorder<PullCommandResponse> responseObserver = StreamRecorder.create();
        endpoint.pullCommand(request, responseObserver);

        ArgumentCaptor<Communication> commCaptor = ArgumentCaptor.forClass(Communication.class);
        verify(repository, times(1)).save(commCaptor.capture());

        CommunicationComponentDto dto = CommunicationFactory.takeApart(commCaptor.getValue());
        assertThat(dto.getCommands()).hasSize(0);

        assertThat(responseObserver.getError()).isNull();
        List<PullCommandResponse> results = responseObserver.getValues();
        assertThat(results).hasSize(1);
        PullCommandResponse response = results.get(0);
        assertThat(response).isEqualTo(PullCommandResponse.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .setCommandId(DEFAULT_COMMAND_ID.getId())
                .setType(CommandType.valueOf(DEFAULT_COMMAND_TYPE))
                .build());
    }

    /**
     * エッジは、飛行コマンド受信APIを実行し、未存在のID指定によりNOT_FOUNDエラーを検出できる。<br>
     * Communicationエンティティが存在しないケース。
     */
    @Test
    public void pullCommandApiNotFoundCommunicationError() {
        PullCommandRequest request = PullCommandRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .setCommandId(DEFAULT_COMMAND_ID.getId())
                .build();
        StreamRecorder<PullCommandResponse> responseObserver = StreamRecorder.create();
        endpoint.pullCommand(request, responseObserver);

        assertThat(responseObserver.getError()).isNotNull();
        assertThat(responseObserver.getError()).isInstanceOf(StatusRuntimeException.class);
        assertThat(((StatusRuntimeException)responseObserver.getError()).getStatus().getCode())
                .isEqualTo(Status.NOT_FOUND.getCode());
    }

    /**
     * エッジは、飛行コマンド受信APIを実行し、未存在のID指定によりNOT_FOUNDエラーを検出できる。<br>
     * Communicationエンティティ内のコマンドが存在しないケース。
     */
    @Test
    public void pullCommandApiNotFoundCommandInCommunicationError() {
        when(repository.getById(DEFAULT_COMMUNICATION_ID))
                .thenReturn(CommunicationFactory.newInstance(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_VEHICLE_ID,
                        DEFAULT_GENERATOR.get()));

        PullCommandRequest request = PullCommandRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .setCommandId(DEFAULT_COMMAND_ID.getId())
                .build();
        StreamRecorder<PullCommandResponse> responseObserver = StreamRecorder.create();
        endpoint.pullCommand(request, responseObserver);

        assertThat(responseObserver.getError()).isNotNull();
        assertThat(responseObserver.getError()).isInstanceOf(StatusRuntimeException.class);
        assertThat(((StatusRuntimeException)responseObserver.getError()).getStatus().getCode())
                .isEqualTo(Status.NOT_FOUND.getCode());
    }

    /**
     * エッジは、飛行コマンド受信APIを実行し、DBエラーのよりINTERNALエラーを検出できる。
     */
    @Test
    public void pullCommandApiInternalError() {
        when(repository.getById(DEFAULT_COMMUNICATION_ID)).thenThrow(new IllegalStateException());

        PullCommandRequest request = PullCommandRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .setCommandId(DEFAULT_COMMAND_ID.getId())
                .build();
        StreamRecorder<PullCommandResponse> responseObserver = StreamRecorder.create();
        endpoint.pullCommand(request, responseObserver);

        assertThat(responseObserver.getError()).isNotNull();
        assertThat(responseObserver.getError()).isInstanceOf(StatusRuntimeException.class);
        assertThat(((StatusRuntimeException)responseObserver.getError()).getStatus().getCode())
                .isEqualTo(Status.INTERNAL.getCode());
    }

    /**
     * エッジは、１件取得APIを実行し、対象のCommunicationを取得できる。
     */
    @Test
    public void getCommunicationApi() {
        when(repository.getById(DEFAULT_COMMUNICATION_ID))
                .thenReturn(newNormalCommunication(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_VEHICLE_ID,
                        DEFAULT_CONTROLLED,
                        DEFAULT_MISSION_ID,
                        DEFAULT_GENERATOR.get()));

        GetCommunicationRequest request = GetCommunicationRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .build();
        StreamRecorder<proto.skysign.common.Communication> responseObserver = StreamRecorder.create();
        endpoint.getCommunication(request, responseObserver);

        assertThat(responseObserver.getError()).isNull();
        List<proto.skysign.common.Communication> results = responseObserver.getValues();
        assertThat(results).hasSize(1);
        proto.skysign.common.Communication response = results.get(0);
        assertThat(response).isEqualTo(proto.skysign.common.Communication.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .setVehicleId(DEFAULT_VEHICLE_ID.getId())
                .setIsControlled(DEFAULT_CONTROLLED)
                .setMissionId(DEFAULT_MISSION_ID.getId())
                .build());
    }

    /**
     * エッジは、１件取得APIを実行し、対象のCommunicationを取得できる。<br>
     * Communicationはすべて未ステージングであり、MissionIdが空であること。
     */
    @Test
    public void getCommunicationNotStagingApi() {
        when(repository.getById(DEFAULT_COMMUNICATION_ID))
                .thenReturn(newNormalCommunication(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_VEHICLE_ID,
                        DEFAULT_CONTROLLED,
                        null,
                        DEFAULT_GENERATOR.get()));

        GetCommunicationRequest request = GetCommunicationRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .build();
        StreamRecorder<proto.skysign.common.Communication> responseObserver = StreamRecorder.create();
        endpoint.getCommunication(request, responseObserver);

        assertThat(responseObserver.getError()).isNull();
        List<proto.skysign.common.Communication> results = responseObserver.getValues();
        assertThat(results).hasSize(1);
        proto.skysign.common.Communication response = results.get(0);
        assertThat(response).isEqualTo(proto.skysign.common.Communication.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .setVehicleId(DEFAULT_VEHICLE_ID.getId())
                .setIsControlled(DEFAULT_CONTROLLED)
                // .setMissionId(DEFAULT_MISSION_ID.getId()) // MissionIdは空
                .build());
    }

    /**
     * エッジは、１件取得APIを実行し、未存在のID指定によりNOT_FOUNDエラーを検出できる。
     */
    @Test
    public void getCommunicationApiNotFoundError() {
        GetCommunicationRequest request = GetCommunicationRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .build();
        StreamRecorder<proto.skysign.common.Communication> responseObserver = StreamRecorder.create();
        endpoint.getCommunication(request, responseObserver);

        assertThat(responseObserver.getError()).isNotNull();
        assertThat(responseObserver.getError()).isInstanceOf(StatusRuntimeException.class);
        assertThat(((StatusRuntimeException)responseObserver.getError()).getStatus().getCode())
                .isEqualTo(Status.NOT_FOUND.getCode());
    }

    /**
     * エッジは、１件取得APIを実行し、DBエラーのよりINTERNALエラーを検出できる。
     */
    @Test
    public void getCommunicationApiInternalError() {
        when(repository.getById(DEFAULT_COMMUNICATION_ID)).thenThrow(new IllegalStateException());

        GetCommunicationRequest request = GetCommunicationRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .build();
        StreamRecorder<proto.skysign.common.Communication> responseObserver = StreamRecorder.create();
        endpoint.getCommunication(request, responseObserver);

        assertThat(responseObserver.getError()).isNotNull();
        assertThat(responseObserver.getError()).isInstanceOf(StatusRuntimeException.class);
        assertThat(((StatusRuntimeException)responseObserver.getError()).getStatus().getCode())
                .isEqualTo(Status.INTERNAL.getCode());
    }
}