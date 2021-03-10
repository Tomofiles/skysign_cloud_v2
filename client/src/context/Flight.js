export const initialFlight = undefined;

export const flightReducer = (state, action) => {
  switch (action.type) {
    case 'ID': {
      return action.id;
    }
    default: {
      return initialFlight;
    }
  }
}
  