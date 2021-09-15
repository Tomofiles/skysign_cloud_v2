package service

import (
	m "fleet-formation/pkg/mission/domain/mission"
)

// NavigationTransformerFromCommand .
func NavigationTransformerFromCommand(
	command Mission,
) *m.Navigation {
	navigation := m.NewNavigation(command.GetNavigation().GetTakeoffPointGroundAltitudeM())
	for _, w := range command.GetNavigation().GetWaypoints() {
		navigation.PushNextWaypoint(
			w.GetLatitudeDegree(),
			w.GetLongitudeDegree(),
			w.GetRelativeAltitudeM(),
			w.GetSpeedMS(),
		)
	}
	return navigation
}
