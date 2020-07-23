package net.tomofiles.skysign.communication.domain.communication;

import java.time.LocalDateTime;
import java.util.Comparator;
import java.util.List;
import java.util.stream.Collectors;

import lombok.AccessLevel;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.RequiredArgsConstructor;
import lombok.Setter;

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

    public void pushTelemetry(
            double latitude,
            double longitude,
            double altitude,
            double relativeAltitude,
            double speed,
            boolean armed,
            String flightMode,
            double orientationX,
            double orientationY,
            double orientationZ,
            double orientationW) {
        this.telemetry = Telemetry.newInstance()
                .setPosition(latitude, longitude, altitude, relativeAltitude, speed)
                .setArmed(armed)
                .setFlightMode(flightMode)
                .setOrientation(orientationX, orientationY, orientationZ, orientationW);
    }

    public TelemetrySnapshot pullTelemetry() {
        return new TelemetrySnapshot(
                this.telemetry.getPosition().getLatitude(),
                this.telemetry.getPosition().getLongitude(),
                this.telemetry.getPosition().getAltitude(),
                this.telemetry.getPosition().getRelativeAltitude(),
                this.telemetry.getSpeed(),
                this.telemetry.isArmed(),
                this.telemetry.getFlightMode(),
                this.telemetry.getOrientation().getX(),
                this.telemetry.getOrientation().getY(),
                this.telemetry.getOrientation().getZ(),
                this.telemetry.getOrientation().getW()
        );
    }

    public List<CommandId> getCommandIds() {
        return this.commands.stream()
                .sorted(Comparator.comparing(Command::getTime))
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
        this.commands.add(new Command(id, commandType, LocalDateTime.now()));
        return id;
    }

    public CommandType pullCommandById(CommandId id) {
        Command command = this.commands.stream()
                .filter(c -> c.equals(Command.empty(id)))
                .findAny()
                .orElse(null);
        if (command == null) {
            return null;
        }
        this.commands.remove(command);
        return command.getType();
    }
}