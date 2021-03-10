export const initialAssignments = [];

export const assignmentsReducer = (state, action) => {
  switch (action.type) {
    case 'ROWS': {
      return [
        ...action.rows,
      ];
    }
    default: {
      return [
        ...initialAssignments
      ];
    }
  }
}
  