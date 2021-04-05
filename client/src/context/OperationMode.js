export const OPERATION_MODE = {
  NONE: 0,
  OPERATION: 1,
  REPORT: 2
}

export const initialOperationMode = OPERATION_MODE.NONE;

export const operationModeReducer = (state, action) => {
  switch (action.type) {
    case 'NONE':
      return OPERATION_MODE.NONE;
    case 'OPERATION':
      return OPERATION_MODE.OPERATION;
    case 'REPORT':
      return OPERATION_MODE.REPORT;
    default:
      return OPERATION_MODE.NONE;
    }
  }
  