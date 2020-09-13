import { getTakeoffHeight } from './MissionUtils'
import { Cartesian3 } from 'cesium';

const DEFAULT_HEIGHT = 10.0;
const DEFAULT_SPEED = 3.0;

export class Mission {
  constructor() {
    this.name = undefined;
    this.takeoffPointGroundHeight = undefined;
    this.items = [];
  }

  nameMission(name) {
    this.name = name;
  }

  async addWaypoint(latitude, longitude) {
    this.items.push({
      latitude: latitude,
      longitude: longitude,
      relativeHeight: DEFAULT_HEIGHT,
      speed: DEFAULT_SPEED,
    });
    if (!this.takeoffPointGroundHeight) {
      let height = await getTakeoffHeight(latitude, longitude);
      this.takeoffPointGroundHeight = height.height;
    }
  }

  changeRelativeHeight(index, height) {
    this.items[index].relativeHeight = height;
  }

  changeSpeed(index, speed) {
    this.items[index].speed = speed;
  }

  changePosition(index, latitude, longitude) {
    this.items[index].latitude = latitude;
    this.items[index].longitude = longitude;
  }

  removeWaypoint(index) {
    this.items.splice(index, 1);
    if (this.items.length === 0) {
      this.takeoffPointGroundHeight = undefined;
    }
  }

  getWaypointsForDisplayToMap() {
    return this.items
        .map((item, index) => {
          return {
            id: "WP_" + index,
            groundPosition: Cartesian3.fromDegrees(
                item.longitude,
                item.latitude,
                this.takeoffPointGroundHeight),
            airPosition: Cartesian3.fromDegrees(
                item.longitude,
                item.latitude,
                this.takeoffPointGroundHeight + item.relativeHeight)
          }
        }, this);
  }

  getPathsForDisplayToMap() {

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
    pairwise(this.items, pairOfWaypoint(paths, this.takeoffPointGroundHeight));
    return paths;
  }
}
