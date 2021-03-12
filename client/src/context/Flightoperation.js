export const initialFlightoperation = undefined;

export const flightoperationReducer = (state, action) => {
  switch (action.type) {
    case 'ID': {
      return action.id;
    }
    default: {
      return initialFlightoperation;
    }
  }
}
  