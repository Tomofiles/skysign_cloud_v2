export const initialTrajectories = [];

export const trajectoriesReducer = (state, action) => {
  switch (action.type) {
    case 'ROWS': {
      return [
        ...action.rows,
      ];
    }
    default: {
      return [
        ...initialTrajectories
      ];
    }
  }
}
  