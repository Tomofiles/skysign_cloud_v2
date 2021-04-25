import { Cartesian3 } from 'cesium';

export const getWaypointsForDisplayToMap = (mission) => {
  return mission.items
      .map((item, index) => {
        return {
          id: "WP_" + index,
          groundPosition: Cartesian3.fromDegrees(
              item.longitude,
              item.latitude,
              mission.takeoff_point_ground_height),
          airPosition: Cartesian3.fromDegrees(
              item.longitude,
              item.latitude,
              mission.takeoff_point_ground_height + item.relative_height)
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
  pairwise(mission.items, pairOfWaypoint(paths, mission.takeoff_point_ground_height));
  return paths;
}