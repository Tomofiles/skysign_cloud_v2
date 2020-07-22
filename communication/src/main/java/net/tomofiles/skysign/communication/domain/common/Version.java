package net.tomofiles.skysign.communication.domain.common;

import lombok.AllArgsConstructor;
import lombok.EqualsAndHashCode;
import lombok.Getter;

@AllArgsConstructor
@Getter
@EqualsAndHashCode(of = "version")
public class Version {
    private final int version;

    public Version nextVersion() {
        return new Version(this.version + 1);
    }
}