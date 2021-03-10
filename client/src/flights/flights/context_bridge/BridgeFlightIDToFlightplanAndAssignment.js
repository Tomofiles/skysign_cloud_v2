import React, { useContext, useEffect } from 'react';

import { getAssignments, getFlightplan } from '../../../plans/flightplans/FlightplansUtils';
import { getFlight } from '../../../flights/flights/FlightUtils';
import { AppContext } from '../../../context/Context';

const BridgeFlightIDToFlightplanAndAssignment = () => {
  const { flight, dispatchFlightplan, dispatchAssignments } = useContext(AppContext);

  useEffect(() => {
    if (flight) {
      getFlight(flight)
      .then(data => {
        getFlightplan(data.flightplan_id)
          .then(data => {
            dispatchFlightplan({ type: 'DATA', data: data });
          })
        getAssignments(data.flightplan_id)
          .then(data => {
            dispatchAssignments({ type: 'ROWS', rows: data.assignments });
          })
      })
    } else {
      dispatchFlightplan({ type: 'NONE' });
      dispatchAssignments({ type: 'NONE' });
    }
  }, [ flight, dispatchFlightplan, dispatchAssignments ])

  return (<></>)
}

export default BridgeFlightIDToFlightplanAndAssignment;