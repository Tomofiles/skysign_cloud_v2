export const initialFlightplan = {
  name: "-",
  description: "-",
};

export const flightplanReducer = (state, action) => {
  switch (action.type) {
    case 'DATA': {
      return {
        ...action.data,
      };
    }
    default: {
      return initialFlightplan;
    }
  }
}
  