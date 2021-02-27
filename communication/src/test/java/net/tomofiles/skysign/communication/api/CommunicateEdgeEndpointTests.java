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
import net.tomofiles.skysign.communication.service.CommunicateEdgeService;
import proto.skysign.common.CommandType;
import proto.skysign.PullCommandRequest;
import proto.skysign.PullCommandResponse;
import proto.skysign.PullUploadMissionRequest;
import proto.skysign.PullUploadMissionResponse;
import proto.skysign.PushTelemetryRequest;
import proto.skysign.PushTelemetryResponse;

import static net.tomofiles.skysign.communication.api.GrpcObjectMother.newNormalTelemetryGrpc;
import static net.tomofiles.skysign.communication.domain.communication.CommunicationObjectMother.newSingleCommandCommunication;
import static net.tomofiles.skysign.communication.domain.communication.ComponentDtoObjectMother.newNormalTelemetryComponentDto;

public class CommunicateEdgeEndpointTests {
    
    private static final CommunicationId DEFAULT_COMMUNICATION_ID = new CommunicationId(UUID.randomUUID().toString());
    private static final CommandId DEFAULT_COMMAND_ID = new CommandId(UUID.randomUUID().toString());
    private static final String DEFAULT_COMMAND_TYPE = "ARM";
    private static final MissionId DEFAULT_MISSION_ID = new MissionId("MISSION_ID_SAMPLE_1");
    private static final boolean DEFAULT_CONTROLLED = true;
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

        this.endpoint = new CommunicateEdgeEndpoint(this.service);
    }

    /**
     * エッジは、テレメトリー送信APIを実行し、対象のCommunicationにテレメトリーを送信できる。<br>
     * Communicationが保持するコマンドは0件のため、返却オブジェクトのコマンドリストは0件であること。
     */
    @Test
    public void pushTelemetryWithNoCommandApi() {
        when(this.repository.getById(DEFAULT_COMMUNICATION_ID))
                .thenReturn(CommunicationFactory.newInstance(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_GENERATOR.get()));

        PushTelemetryRequest request = PushTelemetryRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .setTelemetry(newNormalTelemetryGrpc())
                .build();
        StreamRecorder<PushTelemetryResponse> responseObserver = StreamRecorder.create();
        this.endpoint.pushTelemetry(request, responseObserver);

        ArgumentCaptor<Communication> commCaptor = ArgumentCaptor.forClass(Communication.class);
        verify(this.repository, times(1)).save(commCaptor.capture());

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
        when(this.repository.getById(DEFAULT_COMMUNICATION_ID))
                .thenReturn(newSingleCommandCommunication(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_CONTROLLED,
                        DEFAULT_GENERATOR.get(),
                        DEFAULT_GENERATOR.get(),
                        DEFAULT_GENERATOR.get()));

        PushTelemetryRequest request = PushTelemetryRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .setTelemetry(newNormalTelemetryGrpc())
                .build();
        StreamRecorder<PushTelemetryResponse> responseObserver = StreamRecorder.create();
        this.endpoint.pushTelemetry(request, responseObserver);

        ArgumentCaptor<Communication> commCaptor = ArgumentCaptor.forClass(Communication.class);
        verify(this.repository, times(1)).save(commCaptor.capture());

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
        this.endpoint.pushTelemetry(request, responseObserver);

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
        when(this.repository.getById(DEFAULT_COMMUNICATION_ID)).thenThrow(new IllegalStateException());

        PushTelemetryRequest request = PushTelemetryRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .setTelemetry(newNormalTelemetryGrpc())
                .build();
        StreamRecorder<PushTelemetryResponse> responseObserver = StreamRecorder.create();
        this.endpoint.pushTelemetry(request, responseObserver);

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
        when(this.repository.getById(DEFAULT_COMMUNICATION_ID))
                .thenReturn(newSingleCommandCommunication(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_CONTROLLED,
                        DEFAULT_GENERATOR.get(),
                        DEFAULT_GENERATOR.get(),
                        DEFAULT_GENERATOR.get()));

        PullCommandRequest request = PullCommandRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .setCommandId(DEFAULT_COMMAND_ID.getId())
                .build();
        StreamRecorder<PullCommandResponse> responseObserver = StreamRecorder.create();
        this.endpoint.pullCommand(request, responseObserver);

        ArgumentCaptor<Communication> commCaptor = ArgumentCaptor.forClass(Communication.class);
        verify(this.repository, times(1)).save(commCaptor.capture());

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
        this.endpoint.pullCommand(request, responseObserver);

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
        when(this.repository.getById(DEFAULT_COMMUNICATION_ID))
                .thenReturn(CommunicationFactory.newInstance(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_GENERATOR.get()));

        PullCommandRequest request = PullCommandRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .setCommandId(DEFAULT_COMMAND_ID.getId())
                .build();
        StreamRecorder<PullCommandResponse> responseObserver = StreamRecorder.create();
        this.endpoint.pullCommand(request, responseObserver);

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
        when(this.repository.getById(DEFAULT_COMMUNICATION_ID)).thenThrow(new IllegalStateException());

        PullCommandRequest request = PullCommandRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .setCommandId(DEFAULT_COMMAND_ID.getId())
                .build();
        StreamRecorder<PullCommandResponse> responseObserver = StreamRecorder.create();
        this.endpoint.pullCommand(request, responseObserver);

        assertThat(responseObserver.getError()).isNotNull();
        assertThat(responseObserver.getError()).isInstanceOf(StatusRuntimeException.class);
        assertThat(((StatusRuntimeException)responseObserver.getError()).getStatus().getCode())
                .isEqualTo(Status.INTERNAL.getCode());
    }

    /**
     * エッジは、アップロードミッション受信APIを実行し、対象のCommunicationからアップロードミッションを受信できる。<br>
     * Communicationからアップロードミッションを取得すると、エンティティ内から削除されることを検証する。
     */
    @Test
    public void pullUploadMissionApi() {
        when(this.repository.getById(DEFAULT_COMMUNICATION_ID))
                .thenReturn(newSingleCommandCommunication(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_CONTROLLED,
                        DEFAULT_GENERATOR.get(),
                        DEFAULT_GENERATOR.get(),
                        DEFAULT_GENERATOR.get()));

        PullUploadMissionRequest request = PullUploadMissionRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .setCommandId(DEFAULT_COMMAND_ID.getId())
                .build();
        StreamRecorder<PullUploadMissionResponse> responseObserver = StreamRecorder.create();
        this.endpoint.pullUploadMission(request, responseObserver);

        ArgumentCaptor<Communication> commCaptor = ArgumentCaptor.forClass(Communication.class);
        verify(this.repository, times(1)).save(commCaptor.capture());

        CommunicationComponentDto dto = CommunicationFactory.takeApart(commCaptor.getValue());
        assertThat(dto.getUploadMissions()).hasSize(0);

        assertThat(responseObserver.getError()).isNull();
        List<PullUploadMissionResponse> results = responseObserver.getValues();
        assertThat(results).hasSize(1);
        PullUploadMissionResponse response = results.get(0);
        assertThat(response).isEqualTo(PullUploadMissionResponse.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .setCommandId(DEFAULT_COMMAND_ID.getId())
                .setMissionId(DEFAULT_MISSION_ID.getId())
                .build());
    }

    /**
     * エッジは、アップロードミッション受信APIを実行し、未存在のID指定によりNOT_FOUNDエラーを検出できる。<br>
     * Communicationエンティティが存在しないケース。
     */
    @Test
    public void pullUploadMissionApiNotFoundCommunicationError() {
        PullUploadMissionRequest request = PullUploadMissionRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .setCommandId(DEFAULT_COMMAND_ID.getId())
                .build();
        StreamRecorder<PullUploadMissionResponse> responseObserver = StreamRecorder.create();
        this.endpoint.pullUploadMission(request, responseObserver);

        assertThat(responseObserver.getError()).isNotNull();
        assertThat(responseObserver.getError()).isInstanceOf(StatusRuntimeException.class);
        assertThat(((StatusRuntimeException)responseObserver.getError()).getStatus().getCode())
                .isEqualTo(Status.NOT_FOUND.getCode());
    }

    /**
     * エッジは、アップロードミッション受信APIを実行し、未存在のID指定によりNOT_FOUNDエラーを検出できる。<br>
     * Communicationエンティティ内のコマンドが存在しないケース。
     */
    @Test
    public void pullUploadMissionApiNotFoundCommandInCommunicationError() {
        when(this.repository.getById(DEFAULT_COMMUNICATION_ID))
                .thenReturn(CommunicationFactory.newInstance(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_GENERATOR.get()));

        PullUploadMissionRequest request = PullUploadMissionRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .setCommandId(DEFAULT_COMMAND_ID.getId())
                .build();
        StreamRecorder<PullUploadMissionResponse> responseObserver = StreamRecorder.create();
        this.endpoint.pullUploadMission(request, responseObserver);

        assertThat(responseObserver.getError()).isNotNull();
        assertThat(responseObserver.getError()).isInstanceOf(StatusRuntimeException.class);
        assertThat(((StatusRuntimeException)responseObserver.getError()).getStatus().getCode())
                .isEqualTo(Status.NOT_FOUND.getCode());
    }

    /**
     * エッジは、アップロードミッション受信APIを実行し、DBエラーのよりINTERNALエラーを検出できる。
     */
    @Test
    public void pullUploadMissionApiInternalError() {
        when(this.repository.getById(DEFAULT_COMMUNICATION_ID)).thenThrow(new IllegalStateException());

        PullUploadMissionRequest request = PullUploadMissionRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .setCommandId(DEFAULT_COMMAND_ID.getId())
                .build();
        StreamRecorder<PullUploadMissionResponse> responseObserver = StreamRecorder.create();
        this.endpoint.pullUploadMission(request, responseObserver);

        assertThat(responseObserver.getError()).isNotNull();
        assertThat(responseObserver.getError()).isInstanceOf(StatusRuntimeException.class);
        assertThat(((StatusRuntimeException)responseObserver.getError()).getStatus().getCode())
                .isEqualTo(Status.INTERNAL.getCode());
    }
}