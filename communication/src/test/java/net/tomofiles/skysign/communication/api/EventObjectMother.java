package net.tomofiles.skysign.communication.api;

import net.tomofiles.skysign.communication.domain.communication.CommunicationId;

public class EventObjectMother {

    /**
     * テスト用CommunicationIdGaveEventのProtocolBuffersバイナリデータを生成する。
     */
    public static byte[] newNormalCommunicationIdGaveEvent(CommunicationId communicationId) {
        return proto.skysign.event.CommunicationIdGaveEvent.newBuilder()
            .setCommunicationId(communicationId.getId())
            .build()
            .toByteArray();
    }

    /**
     * テスト用CommunicationIdRemovedEventのProtocolBuffersバイナリデータを生成する。
     */
    public static byte[] newNormalCommunicationIdRemovedEvent(CommunicationId communicationId) {
        return proto.skysign.event.CommunicationIdRemovedEvent.newBuilder()
            .setCommunicationId(communicationId.getId())
            .build()
            .toByteArray();
    }
}