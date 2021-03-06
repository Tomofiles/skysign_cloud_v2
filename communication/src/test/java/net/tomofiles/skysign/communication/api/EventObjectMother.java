package net.tomofiles.skysign.communication.api;

import net.tomofiles.skysign.communication.domain.communication.CommunicationId;

public class EventObjectMother {

    /**
     * テスト用CommunicationIdGaveEventのProtocolBuffersバイナリデータを生成する。
     */
    public static byte[] newNormalCommunicationIdGaveEvent(CommunicationId communicationId, String version) {
        return proto.skysign.event.CommunicationIdGaveEvent.newBuilder()
            .setCommunicationId(communicationId.getId())
            .setVersion(version)
            .build()
            .toByteArray();
    }

    /**
     * テスト用CommunicationIdRemovedEventのProtocolBuffersバイナリデータを生成する。
     */
    public static byte[] newNormalCommunicationIdRemovedEvent(CommunicationId communicationId, String version) {
        return proto.skysign.event.CommunicationIdRemovedEvent.newBuilder()
            .setCommunicationId(communicationId.getId())
            .setVersion(version)
            .build()
            .toByteArray();
    }
}