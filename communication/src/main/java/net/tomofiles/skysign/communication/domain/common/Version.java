package net.tomofiles.skysign.communication.domain.common;

import lombok.AllArgsConstructor;
import lombok.Getter;

@AllArgsConstructor
@Getter
public class Version {
    private final int version;

    public Version nextVersion() {
        return new Version(this.version + 1);
    }
}