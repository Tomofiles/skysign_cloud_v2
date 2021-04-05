import { COMMAND_TYPE } from "../flights/flights/FlightControlUtils";

export const initialSteps = [];

export const stepsReducer = (state, action) => {
  switch (action.type) {
    case 'INIT': {
      const newSteps = [];
      action.ids
        .forEach(id => {
          newSteps
            .push({
              communication_id: id,
            })
        });
      return newSteps;
    }
    case 'CHANGE_STEP': {
      const newSteps = [ ...state ];
      newSteps
        .filter(step => step.communication_id === action.communication_id)
        .forEach(step => {
          step.step = action.step;
          step.command = getCommandByStep(action.step);
          step.mission = action.mission_id;
        });
      return newSteps;
    }
    default: {
      return initialSteps;
    }
  }
}

const getCommandByStep = step => {
  switch (step) {
    case 0:
      return COMMAND_TYPE.UPLOAD;
    case 1:
      return COMMAND_TYPE.TAKEOFF;
    case 2:
      return COMMAND_TYPE.START;
    case 3:
      return COMMAND_TYPE.PAUSE;
    case 4:
      return COMMAND_TYPE.LAND;
    case 5:
      return COMMAND_TYPE.RETURN;
    default:
      return COMMAND_TYPE.NONE;
  }
}