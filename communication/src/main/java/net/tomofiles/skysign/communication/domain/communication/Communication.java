package net.tomofiles.skysign.communication.domain.communication;

import java.time.LocalDateTime;
import java.util.ArrayList;
import java.util.Comparator;
import java.util.List;
import java.util.stream.Collectors;

import lombok.AccessLevel;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.Setter;
import lombok.ToString;
import net.tomofiles.skysign.communication.domain.vehicle.VehicleId;
import net.tomofiles.skysign.communication.event.EmptyPublisher;
import net.tomofiles.skysign.communication.event.Publisher;

@EqualsAndHashCode(of = {"id"})
@ToString
public class Communication {
    @Getter
    private final CommunicationId id;

    private final Generator generator;

    @Getter
    private final VehicleId vehicleId;

    @Getter
    @Setter(value = AccessLevel.PACKAGE)
    private boolean controlled = false;

    @Getter
    @Setter(value = AccessLevel.PACKAGE)
    private MissionId missionId = null;

    @Getter(value = AccessLevel.PACKAGE)
    @Setter(value = AccessLevel.PACKAGE)
    private Telemetry telemetry = null;

    @Getter(value = AccessLevel.PACKAGE)
    private final List<Command> commands;

    @Setter
    private Publisher publisher = new EmptyPublisher();
    
    Communication(CommunicationId id, VehicleId vehicleId, Generator generator) {
        this.id = id;
        this.vehicleId = vehicleId;
        this.commands = new ArrayList<>();

        this.generator = generator;
    }

    public void pushTelemetry(TelemetrySnapshot snapshot) {
        LocalDateTime time = this.generator.newTime();
        this.telemetry = Telemetry.newInstance()
                .setPosition(
                        snapshot.getLatitude(),
                        snapshot.getLongitude(),
                        snapshot.getAltitude(),
                        snapshot.getRelativeAltitude(),
                        snapshot.getSpeed())
                .setArmed(snapshot.isArmed())
                .setFlightMode(snapshot.getFlightMode())
                .setOrientation(
                        snapshot.getX(),
                        snapshot.getY(),
                        snapshot.getZ(),
                        snapshot.getW());
        this.publisher
                .publish(
                        new TelemetryUpdatedEvent(snapshot, time));
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

    public void staging(MissionId missionId) {
        this.missionId = missionId;
    }

    public void cancel() {
        this.missionId = null;
    }

    public void control() {
        LocalDateTime time = this.generator.newTime();
        this.controlled = true;
        this.publisher
                .publish(
                        new CommunicationControlledEvent(
                                this.id,
                                this.vehicleId,
                                this.missionId,
                                time));
    }

    public void uncontrol() {
        this.controlled = false;
    }

    public CommandId pushCommand(CommandType commandType) {
        CommandId id = this.generator.newCommandId();
        LocalDateTime time = this.generator.newTime();
        this.commands.add(new Command(id, commandType, time));
        return id;
    }

    public CommandType pullCommandById(CommandId id) {
        Command command = this.commands.stream()
                .filter(Command.empty(id)::equals)
                .findAny()
                .orElse(null);
        if (command == null) {
            return null;
        }
        this.commands.remove(command);
        return command.getType();
    }
}