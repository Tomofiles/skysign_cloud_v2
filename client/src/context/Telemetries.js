export const initialTelemetries = [];

export const telemetriesReducer = (state, action) => {
  switch (action.type) {
    case 'ROWS': {
      return [
        ...action.rows,
      ];
    }
    default: {
      return [
        ...initialTelemetries
      ];
    }
  }
}
  