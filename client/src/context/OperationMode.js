export const OPERATION_MODE = {
  NONE: 0,
  OPERATION: 1,
}

export const initialOperationMode = OPERATION_MODE.NONE;

export const operationModeReducer = (state, action) => {
  switch (action.type) {
    case 'NONE':
      return OPERATION_MODE.NONE;
    case 'OPERATION':
      return OPERATION_MODE.OPERATION;
    default:
      return OPERATION_MODE.NONE;
    }
  }
  