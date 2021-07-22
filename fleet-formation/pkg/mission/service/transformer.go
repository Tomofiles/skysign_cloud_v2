package service

import (
	m "fleet-formation/pkg/mission/domain/mission"
)

// NavigationTransformerFromCommand .
func NavigationTransformerFromCommand(
	command Mission,
) *m.Navigation {
	navigation := m.NewNavigation(command.GetNavigation().GetTakeoffPointGroundHeight())
	for _, w := range command.GetNavigation().GetWaypoints() {
		navigation.PushNextWaypoint(
			w.GetLatitude(),
			w.GetLongitude(),
			w.GetRelativeHeight(),
			w.GetSpeed(),
		)
	}
	return navigation
}
