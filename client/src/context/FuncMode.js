export const FUNC_MODE = {
  NONE: 0,
  ASSETS: 1,
  MISSIONS: 2,
  PLANS: 3,
  CONTROLS: 4,
  FLIGHTS: 5,
}

export const initialFuncMode = FUNC_MODE.FLIGHTS;

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
    case 'CONTROLS':
      return FUNC_MODE.CONTROLS;
    case 'FLIGHTS':
      return FUNC_MODE.FLIGHTS;
    default:
      return FUNC_MODE.NONE;
    }
  }
  