import React, { useContext, useEffect } from 'react';

import { AppContext } from '../context/Context';

const BridgeVehicleToStep = () => {
  const { vehicles, dispatchSteps } = useContext(AppContext);

  useEffect(() => {
    if (vehicles.length === 0) {
      dispatchSteps({ type: 'NONE' });
      return;
    }

    let communicationIds = [];
    vehicles
      .forEach(vehicle => {
        communicationIds.push(vehicle.commId);
      });

    dispatchSteps({ type: 'INIT', ids: communicationIds });
  }, [ vehicles, dispatchSteps ])

  return (<></>)
}

export default BridgeVehicleToStep;