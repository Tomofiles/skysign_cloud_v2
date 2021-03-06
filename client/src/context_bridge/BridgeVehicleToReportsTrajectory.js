import React, { useContext, useEffect } from 'react';

import { AppContext } from '../context/Context';
import { OPERATION_MODE } from '../context/OperationMode';
import { getTrajectory } from '../reports/reports/ReportUtils';

const BridgeVehicleToReportsTrajectory = () => {
  const { vehicles, operationMode, dispatchTrajectories } = useContext(AppContext);

  useEffect(() => {
    if (operationMode !== OPERATION_MODE.REPORT) {
      dispatchTrajectories({ type: 'NONE' });
      return;
    }
    if (vehicles.length === 0) {
      dispatchTrajectories({ type: 'NONE' });
      return;
    }

    let trajectories = [];
    vehicles
      .forEach(vehicle => {
        trajectories.push(getTrajectory(vehicle.id));
      });

    Promise
      .all(trajectories)
      .then(data => {
        console.log(data);
        dispatchTrajectories({ type: 'ROWS', rows: data });
      });
  }, [ vehicles, operationMode, dispatchTrajectories ])

  return (<></>)
}

export default BridgeVehicleToReportsTrajectory;