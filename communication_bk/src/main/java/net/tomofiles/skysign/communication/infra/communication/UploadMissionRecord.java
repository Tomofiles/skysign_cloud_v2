package net.tomofiles.skysign.communication.infra.communication;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.EqualsAndHashCode;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
@EqualsAndHashCode(of = {"id"})
public class UploadMissionRecord {
    private String id;
    private String communicationId;
    private String missionId;
}