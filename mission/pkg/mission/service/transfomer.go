package service

import (
	m "mission/pkg/mission/domain/mission"
)

// NavigationTransfomerFromCommand .
func NavigationTransfomerFromCommand(
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
