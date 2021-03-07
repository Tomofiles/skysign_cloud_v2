export const initialMissions = [];

export const missionsReducer = (state, action) => {
  switch (action.type) {
    case 'ROWS': {
      return [
        ...action.rows,
      ];
    }
    default: {
      return [
        ...initialMissions
      ];
    }
  }
}
  