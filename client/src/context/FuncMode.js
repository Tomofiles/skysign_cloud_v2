export const FUNC_MODE = {
  NONE: 0,
  ASSETS: 1,
  MISSIONS: 2,
  PLANS: 3,
  FLIGHTS: 4,
  REPORTS: 5,
}

export const initialFuncMode = FUNC_MODE.REPORTS;

export const funcModeReducer = (state, action) => {
  switch (action.type) {
    case 'NONE':
      return FUNC_MODE.NONE;
    case 'ASSETS':
      return FUNC_MODE.ASSETS;
    case 'MISSIONS':
      return FUNC_MODE.MISSIONS;
    case 'PLANS':
      return FUNC_MODE.PLANS;
    case 'FLIGHTS':
      return FUNC_MODE.FLIGHTS;
    case 'REPORTS':
      return FUNC_MODE.REPORTS;
    default:
      return FUNC_MODE.NONE;
    }
  }
  