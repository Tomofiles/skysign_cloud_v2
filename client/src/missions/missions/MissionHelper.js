import { Cartesian3 } from 'cesium';

export const getWaypointsForDisplayToMap = (mission) => {
  return mission.items
      .map((item, index) => {
        return {
          id: "WP_" + index,
          groundPosition: Cartesian3.fromDegrees(
              item.longitude,
              item.latitude,
              mission.takeoffPointGroundHeight),
          airPosition: Cartesian3.fromDegrees(
              item.longitude,
              item.latitude,
              mission.takeoffPointGroundHeight + item.relativeHeight)
        }
      });
}

export const getPathsForDisplayToMap = (mission) => {

  const pairOfWaypoint = (paths, takeoffPointGroundHeight) => {
    return (prev, current, index) => {
      paths.push({
        id: "PT_" + index,
        prevPosition: Cartesian3.fromDegrees(
            prev.longitude,
            prev.latitude,
            takeoffPointGroundHeight + prev.relativeHeight),
        currentPosition: Cartesian3.fromDegrees(
            current.longitude,
            current.latitude,
            takeoffPointGroundHeight + current.relativeHeight),
      });
    };
  };

  const pairwise = (arr, func) => {
    for(let i = 0; i < arr.length - 1; i++){
        func(arr[i], arr[i + 1], i)
    }
  }

  const paths = [];
  pairwise(mission.items, pairOfWaypoint(paths, mission.takeoffPointGroundHeight));
  return paths;
}