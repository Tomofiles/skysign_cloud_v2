package net.tomofiles.skysign.communication.domain.communication;

import java.util.List;
import java.util.stream.Collectors;

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

    @Getter(value = AccessLevel.PACKAGE)
    @Setter(value = AccessLevel.PACKAGE)
    private Telemetry telemetry = null;

    @Getter(value = AccessLevel.PACKAGE)
    private final List<Command> commands;

    @Getter
    @Setter(value = AccessLevel.PACKAGE)
    private Version version;

    public void pushTelemetry(
            double latitude,
            double longitude,
            double altitude,
            double speed,
            boolean armed,
            String flightMode,
            double orientationX,
            double orientationY,
            double orientationZ,
            double orientationW) {
        this.telemetry = Telemetry.newInstance()
                .setPosition(latitude, longitude, altitude, speed)
                .setArmed(armed)
                .setFlightMode(flightMode)
                .setOrientation(orientationX, orientationY, orientationZ, orientationW);
    }

    public TelemetrySnapshot pullTelemetry() {
        return new TelemetrySnapshot(
                this.telemetry.getPosition().getLatitude(),
                this.telemetry.getPosition().getLongitude(),
                this.telemetry.getPosition().getAltitude(),
                this.telemetry.getSpeed(),
                this.telemetry.isArmed(),
                this.telemetry.getFlightMode(),
                this.telemetry.getOrientation().getX(),
                this.telemetry.getOrientation().getY(),
                this.telemetry.getOrientation().getZ(),
                this.telemetry.getOrientation().getW()
        );
    }

    public List<CommandId> getCommandId() {
        return this.commands.stream()
                .map(Command::getId)
                .collect(Collectors.toList());
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

    public CommandType pullCommandById(CommandId id) {
        Command command = this.commands.stream()
                .filter(c -> c.equals(new Command(id, CommandType.NONE)))
                .findAny()
                .orElse(null);
        if (command == null) {
            return null;
        }
        this.commands.remove(command);
        return command.getType();
    }
}