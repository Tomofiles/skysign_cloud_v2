import { Cartesian3 } from 'cesium';

export const getWaypointsForDisplayToMap = (mission) => {
  return mission.navigation.waypoints
      .map((waypoint, index) => {
        return {
          id: "WP_" + index,
          groundPosition: Cartesian3.fromDegrees(
              waypoint.longitude,
              waypoint.latitude,
              mission.navigation.takeoff_point_ground_height),
          airPosition: Cartesian3.fromDegrees(
              waypoint.longitude,
              waypoint.latitude,
              mission.navigation.takeoff_point_ground_height + waypoint.relative_height)
        }
      });
}

export const getPathsForDisplayToMap = (mission) => {

  const pairOfWaypoint = (paths, takeoff_point_ground_height) => {
    return (prev, current, index) => {
      paths.push({
        id: "PT_" + index,
        prevPosition: Cartesian3.fromDegrees(
            prev.longitude,
            prev.latitude,
            takeoff_point_ground_height + prev.relative_height),
        currentPosition: Cartesian3.fromDegrees(
            current.longitude,
            current.latitude,
            takeoff_point_ground_height + current.relative_height),
      });
    };
  };

  const pairwise = (arr, func) => {
    for(let i = 0; i < arr.length - 1; i++){
        func(arr[i], arr[i + 1], i)
    }
  }

  const paths = [];
  pairwise(mission.navigation.waypoints, pairOfWaypoint(paths, mission.navigation.takeoff_point_ground_height));
  return paths;
}