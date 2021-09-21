import React, { useContext, useEffect } from 'react';
import { getVehicle } from '../assets/vehicles/VehicleUtils';

import { AppContext } from '../context/Context';

const BridgeAssignmentToVehicle = () => {
  const { assignments, dispatchVehicles, dispatchMessage } = useContext(AppContext);

  useEffect(() => {
    if (assignments.length === 0) {
      dispatchVehicles({ type: 'NONE' });
      return;
    }

    let vehicles = [];
    assignments
      .forEach(assignment => {
        vehicles.push(getVehicle(assignment.vehicle_id));
      });

    Promise
      .all(vehicles)
      .then(data => {
        dispatchVehicles({ type: 'ROWS', rows: data });
      })
      .catch(message => {
        dispatchMessage({ type: 'NOTIFY_ERROR', message: message });
      });
  }, [ assignments, dispatchVehicles, dispatchMessage ])

  return (<></>)
}

export default BridgeAssignmentToVehicle;