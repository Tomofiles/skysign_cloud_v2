package net.tomofiles.skysign.communication.domain.communication;

import java.util.UUID;

import lombok.AccessLevel;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.RequiredArgsConstructor;

@Getter
@RequiredArgsConstructor(access = AccessLevel.PUBLIC)
@EqualsAndHashCode(of = {"id"})
public class CommunicationId {
    private final String id;
    
    public static CommunicationId newId() {
        return new CommunicationId(UUID.randomUUID().toString());
    }
}