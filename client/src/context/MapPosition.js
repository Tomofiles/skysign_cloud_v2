export const initialMapPosition = {
  longitude: 0,
  latitude: 0,
  height: 0,
};

export const mapPositionReducer = (state, action) => {
  switch (action.type) {
    case 'CURRENT': {
      return {
        ...initialMapPosition,
        longitude: action.longitude,
        latitude: action.latitude,
        height: action.height,
      };
    }
    default: {
      return {
        ...initialMapPosition,
      };
    }
  }
}
  