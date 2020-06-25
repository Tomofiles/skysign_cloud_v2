package net.tomofiles.skysign.communication.domain.vehicle;

import java.util.UUID;

import lombok.AccessLevel;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.RequiredArgsConstructor;

@Getter
@RequiredArgsConstructor(access = AccessLevel.PUBLIC)
@EqualsAndHashCode(of = {"id"})
public class VehicleId {
    private final String id;

    public static VehicleId newId() {
        return new VehicleId(UUID.randomUUID().toString());
    }
}