export const initialStagingRows = [];

export const stagingRowsReducer = (state, action) => {
  switch (action.type) {
    case 'ROWS': {
      return [
        ...action.rows,
      ];
    }
    default: {
      return [
        ...initialStagingRows
      ];
    }
  }
}
  