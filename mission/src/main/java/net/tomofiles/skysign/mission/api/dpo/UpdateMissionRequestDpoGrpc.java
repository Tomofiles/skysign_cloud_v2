package net.tomofiles.skysign.mission.api.dpo;

import lombok.RequiredArgsConstructor;
import net.tomofiles.skysign.mission.domain.mission.GeodesicCoordinates;
import net.tomofiles.skysign.mission.domain.mission.Height;
import net.tomofiles.skysign.mission.domain.mission.MissionId;
import net.tomofiles.skysign.mission.domain.mission.Navigation;
import net.tomofiles.skysign.mission.domain.mission.Speed;
import net.tomofiles.skysign.mission.service.dpo.UpdateMissionRequestDpo;
import proto.skysign.Mission;

@RequiredArgsConstructor
public class UpdateMissionRequestDpoGrpc implements UpdateMissionRequestDpo {

    private final Mission request;

    @Override
    public MissionId getMissionId() {
        return new MissionId(request.getId());
    }

    @Override
    public String getMissionName() {
        return request.getName();
    }

    @Override
    public Navigation getNavigation() {
        Navigation navigation = new Navigation();
        navigation.setTakeoffPointGroundHeight(Height.fromM(request.getTakeoffPointGroundHeight()));
        request.getItemsList()
                .forEach(item -> {
                    navigation.pushNextWaypoint(
                        new GeodesicCoordinates(item.getLatitude(), item.getLongitude()),
                        Height.fromM(item.getRelativeHeight()),
                        Speed.fromMS(item.getSpeed()));
                });
        return navigation;
    }
}