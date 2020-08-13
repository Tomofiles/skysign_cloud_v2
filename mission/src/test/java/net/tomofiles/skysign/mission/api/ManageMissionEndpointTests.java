package net.tomofiles.skysign.mission.api;

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
import net.tomofiles.skysign.mission.domain.mission.Generator;
import net.tomofiles.skysign.mission.domain.mission.Mission;
import net.tomofiles.skysign.mission.domain.mission.MissionId;
import net.tomofiles.skysign.mission.domain.mission.MissionRepository;
import net.tomofiles.skysign.mission.domain.mission.Version;
import net.tomofiles.skysign.mission.service.ManageMissionService;
import proto.skysign.DeleteMissionRequest;
import proto.skysign.Empty;
import proto.skysign.GetMissionRequest;
import proto.skysign.ListMissionsRequest;
import proto.skysign.ListMissionsResponses;

import static net.tomofiles.skysign.mission.api.GrpcObjectMother.newSingleItemMissionGrpc;
import static net.tomofiles.skysign.mission.api.GrpcObjectMother.newSingleItemMissionNoIDGrpc;
import static net.tomofiles.skysign.mission.domain.mission.MissionObjectMother.newSingleNavigationMission;

public class ManageMissionEndpointTests {
    
    private static final MissionId DEFAULT_MISSION_ID = new MissionId(UUID.randomUUID().toString());
    private static final Version DEFAULT_VERSION = new Version(UUID.randomUUID().toString());
    private static final Supplier<Generator> DEFAULT_GENERATOR = () -> {
        return new Generator(){
            @Override
            public MissionId newMissionId() {
                return DEFAULT_MISSION_ID;
            }

            @Override
            public Version newVersion() {
                return DEFAULT_VERSION;
            }
        };
    };
    
    @Mock
    private MissionRepository repository;

    @Mock
    private Generator generator;

    @InjectMocks
    private ManageMissionService service;

    private ManageMissionEndpoint endpoint;

    @BeforeEach
    public void beforeEach() {
        initMocks(this);

        endpoint = new ManageMissionEndpoint(service);
    }

    /**
     * ユーザーは、全件取得APIを実行し、すべてのMissionをリスト形式で取得できる。
     */
    @Test
    public void getAllApi() {
        when(repository.getAll()).thenReturn(Arrays.asList(new Mission[] {
            newSingleNavigationMission(DEFAULT_MISSION_ID, DEFAULT_VERSION, DEFAULT_GENERATOR.get()),
            newSingleNavigationMission(DEFAULT_MISSION_ID, DEFAULT_VERSION, DEFAULT_GENERATOR.get()),
            newSingleNavigationMission(DEFAULT_MISSION_ID, DEFAULT_VERSION, DEFAULT_GENERATOR.get())
        }));

        ListMissionsRequest request = ListMissionsRequest.newBuilder()
                .build();
        StreamRecorder<ListMissionsResponses> responseObserver = StreamRecorder.create();
        endpoint.listMissions(request, responseObserver);

        assertThat(responseObserver.getError()).isNull();
        List<ListMissionsResponses> results = responseObserver.getValues();
        assertThat(results).hasSize(1);
        ListMissionsResponses response = results.get(0);
        assertThat(response).isEqualTo(ListMissionsResponses.newBuilder()
                .addMissions(newSingleItemMissionGrpc(DEFAULT_MISSION_ID))
                .addMissions(newSingleItemMissionGrpc(DEFAULT_MISSION_ID))
                .addMissions(newSingleItemMissionGrpc(DEFAULT_MISSION_ID))
                .build());
    }

    /**
     * ユーザーは、全件取得APIを実行し、DBエラーのよりINTERNALエラーを検出できる。
     */
    @Test
    public void getAllApiInternalError() {
        when(repository.getAll()).thenThrow(new IllegalStateException());

        ListMissionsRequest request = ListMissionsRequest.newBuilder()
                .build();
        StreamRecorder<ListMissionsResponses> responseObserver = StreamRecorder.create();
        endpoint.listMissions(request, responseObserver);

        assertThat(responseObserver.getError()).isNotNull();
        assertThat(responseObserver.getError()).isInstanceOf(StatusRuntimeException.class);
        assertThat(((StatusRuntimeException)responseObserver.getError()).getStatus().getCode())
                .isEqualTo(Status.INTERNAL.getCode());
    }

    /**
     * ユーザーは、１件取得APIを実行し、対象のMissionを取得できる。
     */
    @Test
    public void getOneApi() {
        when(repository.getById(DEFAULT_MISSION_ID))
                .thenReturn(newSingleNavigationMission(
                        DEFAULT_MISSION_ID,
                        DEFAULT_VERSION,
                        DEFAULT_GENERATOR.get()));

        GetMissionRequest request = GetMissionRequest.newBuilder()
                .setId(DEFAULT_MISSION_ID.getId())
                .build();
        StreamRecorder<proto.skysign.Mission> responseObserver = StreamRecorder.create();
        endpoint.getMission(request, responseObserver);

        assertThat(responseObserver.getError()).isNull();
        List<proto.skysign.Mission> results = responseObserver.getValues();
        assertThat(results).hasSize(1);
        proto.skysign.Mission response = results.get(0);
        assertThat(response).isEqualTo(newSingleItemMissionGrpc(DEFAULT_MISSION_ID));
    }

    /**
     * ユーザーは、１件取得APIを実行し、未存在のID指定によりNOT_FOUNDエラーを検出できる。
     */
    @Test
    public void getOneApiNotFoundError() {
        GetMissionRequest request = GetMissionRequest.newBuilder()
                .setId(DEFAULT_MISSION_ID.getId())
                .build();
        StreamRecorder<proto.skysign.Mission> responseObserver = StreamRecorder.create();
        endpoint.getMission(request, responseObserver);

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
        when(repository.getById(DEFAULT_MISSION_ID)).thenThrow(new IllegalStateException());

        GetMissionRequest request = GetMissionRequest.newBuilder()
                .setId(DEFAULT_MISSION_ID.getId())
                .build();
        StreamRecorder<proto.skysign.Mission> responseObserver = StreamRecorder.create();
        endpoint.getMission(request, responseObserver);

        assertThat(responseObserver.getError()).isNotNull();
        assertThat(responseObserver.getError()).isInstanceOf(StatusRuntimeException.class);
        assertThat(((StatusRuntimeException)responseObserver.getError()).getStatus().getCode())
                .isEqualTo(Status.INTERNAL.getCode());
    }

    /**
     * ユーザーは、登録APIを実行し、新しいMissionを登録できる。
     */
    @Test
    public void createApi() {
        when(generator.newMissionId()).thenReturn(DEFAULT_MISSION_ID);
        when(generator.newVersion()).thenReturn(DEFAULT_VERSION);

        proto.skysign.Mission request = newSingleItemMissionNoIDGrpc();
        StreamRecorder<proto.skysign.Mission> responseObserver = StreamRecorder.create();
        endpoint.createMission(request, responseObserver);

        verify(repository, times(1)).save(newSingleNavigationMission(
                DEFAULT_MISSION_ID,
                DEFAULT_VERSION,
                DEFAULT_GENERATOR.get()
        ));

        assertThat(responseObserver.getError()).isNull();
        List<proto.skysign.Mission> results = responseObserver.getValues();
        assertThat(results).hasSize(1);
        proto.skysign.Mission response = results.get(0);
        assertThat(response).isEqualTo(newSingleItemMissionGrpc(DEFAULT_MISSION_ID));
    }

    /**
     * ユーザーは、登録APIを実行し、DBエラーのよりINTERNALエラーを検出できる。
     */
    @Test
    public void createApiInternalError() {
        doThrow(new IllegalStateException()).when(repository).save(any());

        proto.skysign.Mission request = newSingleItemMissionNoIDGrpc();
        StreamRecorder<proto.skysign.Mission> responseObserver = StreamRecorder.create();
        endpoint.createMission(request, responseObserver);

        assertThat(responseObserver.getError()).isNotNull();
        assertThat(responseObserver.getError()).isInstanceOf(StatusRuntimeException.class);
        assertThat(((StatusRuntimeException)responseObserver.getError()).getStatus().getCode())
                .isEqualTo(Status.INTERNAL.getCode());
    }

    /**
     * ユーザーは、更新APIを実行し、既存のMissionを更新できる。
     */
    @Test
    public void updateApi() {
        when(repository.getById(DEFAULT_MISSION_ID))
                .thenReturn(newSingleNavigationMission(
                        DEFAULT_MISSION_ID,
                        DEFAULT_VERSION,
                        DEFAULT_GENERATOR.get()));

        proto.skysign.Mission request = newSingleItemMissionGrpc(DEFAULT_MISSION_ID);
        StreamRecorder<proto.skysign.Mission> responseObserver = StreamRecorder.create();
        endpoint.updateMission(request, responseObserver);

        verify(repository, times(1)).save(newSingleNavigationMission(
                DEFAULT_MISSION_ID,
                DEFAULT_VERSION,
                DEFAULT_GENERATOR.get()));

        assertThat(responseObserver.getError()).isNull();
        List<proto.skysign.Mission> results = responseObserver.getValues();
        assertThat(results).hasSize(1);
        proto.skysign.Mission response = results.get(0);
        assertThat(response).isEqualTo(newSingleItemMissionGrpc(DEFAULT_MISSION_ID));
    }

    /**
     * ユーザーは、更新APIを実行し、未存在のID指定によりNOT_FOUNDエラーを検出できる。
     */
    @Test
    public void updateApiNotFoundError() {
        proto.skysign.Mission request = newSingleItemMissionGrpc(DEFAULT_MISSION_ID);
        StreamRecorder<proto.skysign.Mission> responseObserver = StreamRecorder.create();
        endpoint.updateMission(request, responseObserver);

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
        when(repository.getById(DEFAULT_MISSION_ID))
                .thenReturn(newSingleNavigationMission(
                        DEFAULT_MISSION_ID,
                        DEFAULT_VERSION,
                        DEFAULT_GENERATOR.get()));
        doThrow(new IllegalStateException()).when(repository).save(any());

        proto.skysign.Mission request = newSingleItemMissionGrpc(DEFAULT_MISSION_ID);
        StreamRecorder<proto.skysign.Mission> responseObserver = StreamRecorder.create();
        endpoint.updateMission(request, responseObserver);

        assertThat(responseObserver.getError()).isNotNull();
        assertThat(responseObserver.getError()).isInstanceOf(StatusRuntimeException.class);
        assertThat(((StatusRuntimeException)responseObserver.getError()).getStatus().getCode())
                .isEqualTo(Status.INTERNAL.getCode());
    }

    /**
     * ユーザーは、削除APIを実行し、既存のMissionを削除できる。
     */
    @Test
    public void deleteApi() {
        when(repository.getById(DEFAULT_MISSION_ID))
                .thenReturn(newSingleNavigationMission(
                        DEFAULT_MISSION_ID,
                        DEFAULT_VERSION,
                        DEFAULT_GENERATOR.get()));

        DeleteMissionRequest request = DeleteMissionRequest.newBuilder()
                .setId(DEFAULT_MISSION_ID.getId())
                .build();
        StreamRecorder<Empty> responseObserver = StreamRecorder.create();
        endpoint.deleteMission(request, responseObserver);

        verify(repository, times(1)).remove(DEFAULT_MISSION_ID, DEFAULT_VERSION);

        assertThat(responseObserver.getError()).isNull();
        List<Empty> results = responseObserver.getValues();
        assertThat(results).hasSize(1);
    }

    /**
     * ユーザーは、削除APIを実行し、未存在のID指定によりNOT_FOUNDエラーを検出できる。
     */
    @Test
    public void deleteApiNotFoundError() {
        DeleteMissionRequest request = DeleteMissionRequest.newBuilder()
                .setId(DEFAULT_MISSION_ID.getId())
                .build();
        StreamRecorder<Empty> responseObserver = StreamRecorder.create();
        endpoint.deleteMission(request, responseObserver);

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
        when(repository.getById(DEFAULT_MISSION_ID))
                .thenReturn(newSingleNavigationMission(
                        DEFAULT_MISSION_ID,
                        DEFAULT_VERSION,
                        DEFAULT_GENERATOR.get()));
        doThrow(new IllegalStateException()).when(repository).remove(any(), any());

        DeleteMissionRequest request = DeleteMissionRequest.newBuilder()
                .setId(DEFAULT_MISSION_ID.getId())
                .build();
        StreamRecorder<Empty> responseObserver = StreamRecorder.create();
        endpoint.deleteMission(request, responseObserver);

        assertThat(responseObserver.getError()).isNotNull();
        assertThat(responseObserver.getError()).isInstanceOf(StatusRuntimeException.class);
        assertThat(((StatusRuntimeException)responseObserver.getError()).getStatus().getCode())
                .isEqualTo(Status.INTERNAL.getCode());
    }
}