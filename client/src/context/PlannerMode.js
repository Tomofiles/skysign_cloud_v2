export const PLANNER_MODE = {
  NONE: 0,
  PLANNING: 1,
}

export const initialPlannerMode = PLANNER_MODE.NONE;

export const plannerModeReducer = (state, action) => {
  switch (action.type) {
    case 'NONE':
      return PLANNER_MODE.NONE;
    case 'PLANNING':
      return PLANNER_MODE.PLANNING;
    default:
      return PLANNER_MODE.NONE;
    }
  }
  