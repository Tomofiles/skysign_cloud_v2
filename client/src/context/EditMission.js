const DEFAULT_HEIGHT = 10.0;
const DEFAULT_SPEED = 3.0;

export const initialEditMission = {
  name: undefined,
  takeoff_point_ground_height: undefined,
  items: [],
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
      const newItems = [ ...state.items ];
      newItems[action.index].relative_height = action.height;
      return {
        ...state,
        items: newItems,
      };
    }
    case 'CHANGE_TAKEOFF_POINT_GROUND_HEIGHT': {
      return {
        ...state,
        takeoff_point_ground_height: action.height,
      };
    }
    case 'CHANGE_SPEED': {
      const newItems = [ ...state.items ];
      newItems[action.index].speed = action.speed;
      return {
        ...state,
        items: newItems,
      };
    }
    case 'ADD_WAYPOINT': {
      const newItems = [ ...state.items ];
      newItems.push({
        latitude: action.latitude,
        longitude: action.longitude,
        relative_height: DEFAULT_HEIGHT,
        speed: DEFAULT_SPEED,
      });
      return {
        ...state,
        items: newItems,
      };
    }
    case 'REMOVE_WAYPOINT': {
      const newItems = [ ...state.items ];
      newItems.splice(action.index, 1);
      return {
        ...state,
        items: newItems,
      };
    }
    case 'CHANGE_POSITION': {
      const newItems = [ ...state.items ];
      newItems[action.index].latitude = action.latitude;
      newItems[action.index].longitude = action.longitude;
      return {
        ...state,
        items: newItems,
      };
    }
    case 'OPEN': {
      return {
        ...initialEditMission,
        name: action.mission.name,
        takeoff_point_ground_height: action.mission.takeoff_point_ground_height,
        items: action.mission.items,
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
  