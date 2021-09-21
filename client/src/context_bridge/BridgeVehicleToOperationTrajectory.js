import React, { useContext, useEffect, useState } from 'react';

import useInterval from 'use-interval';

import { AppContext } from '../context/Context';
import { OPERATION_MODE } from '../context/OperationMode';
import { getTrajectory } from '../reports/reports/ReportUtils';

const BridgeVehicleToOperationTrajectory = () => {
  const { vehicles, operationMode, dispatchTrajectories, dispatchMessage } = useContext(AppContext);
  const [ isRender, setIsRender ] = useState(false);

  useEffect(() => {
    if (operationMode === OPERATION_MODE.OPERATION) {
      setIsRender(true);
    } else {
      setIsRender(false);
      dispatchTrajectories({ type: 'NONE' });
    }
  }, [ vehicles, operationMode, dispatchTrajectories ])

  useInterval(() => {
    if (!isRender) {
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
      })
      .catch(message => {
        dispatchMessage({ type: 'NOTIFY_ERROR', message: message });
      });
  },
  1000);

  return (<></>)
}

export default BridgeVehicleToOperationTrajectory;