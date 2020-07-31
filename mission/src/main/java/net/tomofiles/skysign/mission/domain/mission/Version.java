package net.tomofiles.skysign.mission.domain.mission;

import java.util.UUID;

import lombok.AllArgsConstructor;
import lombok.EqualsAndHashCode;
import lombok.Getter;

@AllArgsConstructor
@Getter
@EqualsAndHashCode(of = "version")
public class Version {
    private final String version;

    public static Version newVersion() {
        return new Version(UUID.randomUUID().toString());
    }
}