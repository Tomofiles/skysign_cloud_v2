export const initialFleet = undefined;

export const fleetReducer = (state, action) => {
  switch (action.type) {
    case 'ID': {
      return action.id;
    }
    default: {
      return initialFleet;
    }
  }
}
  