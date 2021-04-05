export const initialVehicles = [];

export const vehiclesReducer = (state, action) => {
  switch (action.type) {
    case 'ROWS': {
      return [
        ...action.rows,
      ];
    }
    default: {
      return [
        ...initialVehicles
      ];
    }
  }
}
  