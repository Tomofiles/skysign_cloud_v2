package net.tomofiles.skysign.mission.domain.mission;

import lombok.AllArgsConstructor;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.ToString;

@AllArgsConstructor
@Getter
@EqualsAndHashCode(of = "version")
@ToString
public class Version {
    private final String version;
}