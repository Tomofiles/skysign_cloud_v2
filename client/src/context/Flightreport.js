export const initialFlightreport = undefined;

export const flightreportReducer = (state, action) => {
  switch (action.type) {
    case 'ID': {
      return action.id;
    }
    default: {
      return initialFlightreport;
    }
  }
}
  