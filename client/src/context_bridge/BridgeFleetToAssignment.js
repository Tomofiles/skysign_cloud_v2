import React, { useContext, useEffect } from 'react';

import { AppContext } from '../context/Context';
import { getAssignments } from '../plans/flightplans/FlightplansUtils';

const BridgeFleetToAssignment = () => {
  const { fleet, dispatchAssignments, dispatchMessage } = useContext(AppContext);

  useEffect(() => {
    if (fleet) {
      getAssignments(fleet)
        .then(data => {
          dispatchAssignments({ type: 'ROWS', rows: data.assignments });
        })
        .catch(message => {
          dispatchMessage({ type: 'NOTIFY_ERROR', message: message });
        });
    } else {
      dispatchAssignments({ type: 'NONE' });
    }
  }, [ fleet, dispatchAssignments, dispatchMessage ])

  return (<></>)
}

export default BridgeFleetToAssignment;