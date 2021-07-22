const DEFAULT_HEIGHT = 10.0;
const DEFAULT_SPEED = 3.0;

export const initialEditNavigation = {
  takeoff_point_ground_height: undefined,
  waypoints: [],
}

export const initialEditMission = {
  name: undefined,
  navigation: initialEditNavigation,
}

export const editMissionReducer = (state, action) => {
  switch (action.type) {
    case 'CHANGE_NAME': {
      return {
        ...state,
        name: action.name,
      };
    }
    case 'CHANGE_RELATIVE_HEIGHT': {
      const newWaypoints = [ ...state.navigation.waypoints ];
      newWaypoints[action.index].relative_height = action.height;
      const newNavigation = {
        ...state.navigation,
        waypoints: newWaypoints,
      }
      return {
        ...state,
        navigation: newNavigation,
      };
    }
    case 'CHANGE_TAKEOFF_POINT_GROUND_HEIGHT': {
      const newNavigation = {
        ...state.navigation,
        takeoff_point_ground_height: action.height,
      }
      return {
        ...state,
        navigation: newNavigation,
      };
    }
    case 'CHANGE_SPEED': {
      const newWaypoints = [ ...state.navigation.waypoints ];
      newWaypoints[action.index].speed = action.speed;
      const newNavigation = {
        ...state.navigation,
        waypoints: newWaypoints,
      }
      return {
        ...state,
        navigation: newNavigation,
      };
    }
    case 'ADD_WAYPOINT': {
      const newWaypoints = [ ...state.navigation.waypoints ];
      newWaypoints.push({
        latitude: action.latitude,
        longitude: action.longitude,
        relative_height: DEFAULT_HEIGHT,
        speed: DEFAULT_SPEED,
      });
      const newNavigation = {
        ...state.navigation,
        waypoints: newWaypoints,
      }
      return {
        ...state,
        navigation: newNavigation,
      };
    }
    case 'REMOVE_WAYPOINT': {
      const newWaypoints = [ ...state.navigation.waypoints ];
      newWaypoints.splice(action.index, 1);
      const newNavigation = {
        ...state.navigation,
        waypoints: newWaypoints,
      }
      return {
        ...state,
        navigation: newNavigation,
      };
    }
    case 'CHANGE_POSITION': {
      const newWaypoints = [ ...state.navigation.waypoints ];
      newWaypoints[action.index].latitude = action.latitude;
      newWaypoints[action.index].longitude = action.longitude;
      const newNavigation = {
        ...state.navigation,
        waypoints: newWaypoints,
      }
      return {
        ...state,
        navigation: newNavigation,
      };
    }
    case 'OPEN': {
      return {
        ...initialEditMission,
        name: action.mission.name,
        navigation: {
          ...initialEditNavigation,
          takeoff_point_ground_height: action.mission.navigation.takeoff_point_ground_height,
          waypoints: action.mission.navigation.waypoints,
        }
      };
    }
    case 'CLEAR': {
      return {
        ...initialEditMission
      };
    }
    default: {
      return {
        ...initialEditMission
      };
    }
  }
}
  