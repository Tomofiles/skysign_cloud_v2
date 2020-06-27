package net.tomofiles.skysign.communication.usecase.dto;

import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class ControlCommandDto {
    private ControlCommandType type;
    private String missionId;
}