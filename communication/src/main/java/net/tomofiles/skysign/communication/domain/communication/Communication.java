package net.tomofiles.skysign.communication.domain.communication;

import java.util.List;

import lombok.AccessLevel;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.RequiredArgsConstructor;
import lombok.Setter;
import net.tomofiles.skysign.communication.domain.common.Version;

@RequiredArgsConstructor(access = AccessLevel.PACKAGE)
@EqualsAndHashCode(of = {"id"})
public class Communication {
    @Getter
    private final CommunicationId id;

    @Getter
    @Setter(value = AccessLevel.PACKAGE)
    private MissionId missionId;

    private Telemetry telemetry = null;

    private final List<Command> commands;

    @Getter
    @Setter(value = AccessLevel.PACKAGE)
    private Version version;

    public void pushTelemetry(double latitude, double longitude) {
        this.telemetry = new Telemetry(latitude, longitude);
    }

    public void standBy(MissionId missionId) {
        this.missionId = missionId;
    }

    public void cancel() {
        this.missionId = null;
    }

    public CommandId pushCommand(CommandType commandType) {
        CommandId id = CommandId.newId();
        this.commands.add(new Command(id, commandType));
        return id;
    }

    public Command pullCommandById(CommandId id) {
        return this.commands.stream()
                .filter(c -> c.equals(new Command(id, CommandType.NONE)))
                .findAny()
                .orElse(null);
    }
}