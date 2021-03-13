import React, { useContext, useEffect } from 'react';

import { AppContext } from '../context/Context';
import { getFlight } from '../flights/flights/FlightUtils';
import { getAssignments, getFlightplan } from '../plans/flightplans/FlightplansUtils';

const BridgeFlightoperationToFlightplanAndAssignment = () => {
  const { flightoperation, dispatchFlightplan, dispatchAssignments } = useContext(AppContext);

  useEffect(() => {
    if (flightoperation) {
      getFlight(flightoperation)
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
  }, [ flightoperation, dispatchFlightplan, dispatchAssignments ])

  return (<></>)
}

export default BridgeFlightoperationToFlightplanAndAssignment;