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
import net.tomofiles.skysign.communication.service.CommunicationUserService;
import proto.skysign.common.CommandType;
import proto.skysign.PullTelemetryRequest;
import proto.skysign.PullTelemetryResponse;
import proto.skysign.PushCommandRequest;
import proto.skysign.PushCommandResponse;
import proto.skysign.PushUploadMissionRequest;
import proto.skysign.PushUploadMissionResponse;

import static net.tomofiles.skysign.communication.api.GrpcObjectMother.newNormalPullTelemetryResponseGrpc;
import static net.tomofiles.skysign.communication.domain.communication.CommunicationObjectMother.newNormalCommunication;
import static net.tomofiles.skysign.communication.domain.communication.ComponentDtoObjectMother.newSingleCommandComponentDto;
import static net.tomofiles.skysign.communication.domain.communication.ComponentDtoObjectMother.newSingleUploadMissionComponentDto;

public class CommunicationUserEndpointTests {
    
    private static final CommunicationId DEFAULT_COMMUNICATION_ID = new CommunicationId(UUID.randomUUID().toString());
    private static final CommandId DEFAULT_COMMAND_ID = new CommandId(UUID.randomUUID().toString());
    private static final String DEFAULT_COMMAND_TYPE = "ARM";
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

        this.endpoint = new CommunicationUserEndpoint(this.service);
    }

    /**
     * ユーザーは、飛行コマンド送信APIを実行し、対象のCommunicationにコマンドを送信できる。
     */
    @Test
    public void pushCommandApi() {
        when(this.repository.getById(DEFAULT_COMMUNICATION_ID))
                .thenReturn(newNormalCommunication(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_GENERATOR.get()));

        PushCommandRequest request = PushCommandRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .setType(CommandType.valueOf(DEFAULT_COMMAND_TYPE))
                .build();
        StreamRecorder<PushCommandResponse> responseObserver = StreamRecorder.create();
        this.endpoint.pushCommand(request, responseObserver);

        ArgumentCaptor<Communication> commCaptor = ArgumentCaptor.forClass(Communication.class);
        verify(this.repository, times(1)).save(commCaptor.capture());

        CommunicationComponentDto dto = CommunicationFactory.takeApart(commCaptor.getValue());
        assertThat(dto.getCommands()).hasSize(1);
        assertThat(dto.getCommands().get(0))
                .isEqualTo(newSingleCommandComponentDto(DEFAULT_GENERATOR.get(),
                        net.tomofiles.skysign.communication.domain.communication.CommandType.valueOf(DEFAULT_COMMAND_TYPE)
                ));

        assertThat(responseObserver.getError()).isNull();
        List<PushCommandResponse> results = responseObserver.getValues();
        assertThat(results).hasSize(1);
        PushCommandResponse response = results.get(0);
        assertThat(response).isEqualTo(PushCommandResponse.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .setType(CommandType.valueOf(DEFAULT_COMMAND_TYPE))
                .setCommandId(DEFAULT_COMMAND_ID.getId())
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
        this.endpoint.pushCommand(request, responseObserver);

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
        when(this.repository.getById(DEFAULT_COMMUNICATION_ID)).thenThrow(new IllegalStateException());

        PushCommandRequest request = PushCommandRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .setType(CommandType.valueOf(DEFAULT_COMMAND_TYPE))
                .build();
        StreamRecorder<PushCommandResponse> responseObserver = StreamRecorder.create();
        this.endpoint.pushCommand(request, responseObserver);

        assertThat(responseObserver.getError()).isNotNull();
        assertThat(responseObserver.getError()).isInstanceOf(StatusRuntimeException.class);
        assertThat(((StatusRuntimeException)responseObserver.getError()).getStatus().getCode())
                .isEqualTo(Status.INTERNAL.getCode());
    }

    /**
     * ユーザーは、ミッションアップロード送信APIを実行し、対象のCommunicationにコマンドを送信できる。
     */
    @Test
    public void pushUploadMissionApi() {
        when(this.repository.getById(DEFAULT_COMMUNICATION_ID))
                .thenReturn(newNormalCommunication(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_GENERATOR.get()));

        PushUploadMissionRequest request = PushUploadMissionRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .setMissionId(DEFAULT_MISSION_ID.getId())
                .build();
        StreamRecorder<PushUploadMissionResponse> responseObserver = StreamRecorder.create();
        this.endpoint.pushUploadMission(request, responseObserver);

        ArgumentCaptor<Communication> commCaptor = ArgumentCaptor.forClass(Communication.class);
        verify(this.repository, times(1)).save(commCaptor.capture());

        CommunicationComponentDto dto = CommunicationFactory.takeApart(commCaptor.getValue());
        assertThat(dto.getCommands()).hasSize(1);
        assertThat(dto.getCommands().get(0))
                .isEqualTo(newSingleCommandComponentDto(DEFAULT_GENERATOR.get(),
                        net.tomofiles.skysign.communication.domain.communication.CommandType.UPLOAD));
        assertThat(dto.getUploadMissions()).hasSize(1);
        assertThat(dto.getUploadMissions().get(0))
                .isEqualTo(newSingleUploadMissionComponentDto(DEFAULT_GENERATOR.get(), DEFAULT_MISSION_ID));

        assertThat(responseObserver.getError()).isNull();
        List<PushUploadMissionResponse> results = responseObserver.getValues();
        assertThat(results).hasSize(1);
        PushUploadMissionResponse response = results.get(0);
        assertThat(response).isEqualTo(PushUploadMissionResponse.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .setMissionId(DEFAULT_MISSION_ID.getId())
                .setCommandId(DEFAULT_COMMAND_ID.getId())
                .build());
    }

    /**
     * ユーザーは、ミッションアップロード送信APIを実行し、未存在のID指定によりNOT_FOUNDエラーを検出できる。
     */
    @Test
    public void pushUploadMissionApiNotFoundError() {
        PushUploadMissionRequest request = PushUploadMissionRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .setMissionId(DEFAULT_MISSION_ID.getId())
                .build();
        StreamRecorder<PushUploadMissionResponse> responseObserver = StreamRecorder.create();
        this.endpoint.pushUploadMission(request, responseObserver);

        assertThat(responseObserver.getError()).isNotNull();
        assertThat(responseObserver.getError()).isInstanceOf(StatusRuntimeException.class);
        assertThat(((StatusRuntimeException)responseObserver.getError()).getStatus().getCode())
                .isEqualTo(Status.NOT_FOUND.getCode());
    }

    /**
     * ユーザーは、ミッションアップロード送信APIを実行し、DBエラーのよりINTERNALエラーを検出できる。
     */
    @Test
    public void pushUploadMissionApiInternalError() {
        when(this.repository.getById(DEFAULT_COMMUNICATION_ID)).thenThrow(new IllegalStateException());

        PushUploadMissionRequest request = PushUploadMissionRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .setMissionId(DEFAULT_MISSION_ID.getId())
                .build();
        StreamRecorder<PushUploadMissionResponse> responseObserver = StreamRecorder.create();
        this.endpoint.pushUploadMission(request, responseObserver);

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
        when(this.repository.getById(DEFAULT_COMMUNICATION_ID))
                .thenReturn(newNormalCommunication(
                        DEFAULT_COMMUNICATION_ID,
                        DEFAULT_GENERATOR.get()));

        PullTelemetryRequest request = PullTelemetryRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .build();
        StreamRecorder<PullTelemetryResponse> responseObserver = StreamRecorder.create();
        this.endpoint.pullTelemetry(request, responseObserver);

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
        this.endpoint.pullTelemetry(request, responseObserver);

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
        when(this.repository.getById(DEFAULT_COMMUNICATION_ID)).thenThrow(new IllegalStateException());

        PullTelemetryRequest request = PullTelemetryRequest.newBuilder()
                .setId(DEFAULT_COMMUNICATION_ID.getId())
                .build();
        StreamRecorder<PullTelemetryResponse> responseObserver = StreamRecorder.create();
        this.endpoint.pullTelemetry(request, responseObserver);

        assertThat(responseObserver.getError()).isNotNull();
        assertThat(responseObserver.getError()).isInstanceOf(StatusRuntimeException.class);
        assertThat(((StatusRuntimeException)responseObserver.getError()).getStatus().getCode())
                .isEqualTo(Status.INTERNAL.getCode());
    }
}