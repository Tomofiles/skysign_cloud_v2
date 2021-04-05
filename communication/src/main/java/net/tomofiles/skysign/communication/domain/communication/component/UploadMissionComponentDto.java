package net.tomofiles.skysign.communication.domain.communication.component;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.EqualsAndHashCode;
import lombok.ToString;

@Data
@AllArgsConstructor
@EqualsAndHashCode
@ToString
public class UploadMissionComponentDto {
    private String id;
    private String missionId;
}