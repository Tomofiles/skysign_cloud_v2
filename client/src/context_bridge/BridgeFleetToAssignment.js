import React, { useContext, useEffect } from 'react';

import { AppContext } from '../context/Context';
import { getAssignments } from '../plans/flightplans/FlightplansUtils';

const BridgeFleetToAssignment = () => {
  const { fleet, dispatchAssignments } = useContext(AppContext);

  useEffect(() => {
    if (fleet) {
      getAssignments(fleet)
        .then(data => {
          dispatchAssignments({ type: 'ROWS', rows: data.assignments });
        })
    } else {
      dispatchAssignments({ type: 'NONE' });
    }
  }, [ fleet, dispatchAssignments ])

  return (<></>)
}

export default BridgeFleetToAssignment;