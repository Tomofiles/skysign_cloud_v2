export const EDIT_MODE = {
  NONE: 0,
  MISSION: 1,
}

export const initialEditMode = EDIT_MODE.NONE;

export const editModeReducer = (state, action) => {
  switch (action.type) {
    case 'NONE':
      return EDIT_MODE.NONE;
    case 'MISSION':
      return EDIT_MODE.MISSION;
    default:
      return EDIT_MODE.NONE;
    }
  }
  