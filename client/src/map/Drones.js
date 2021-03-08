import React, { useContext } from 'react';
import { AppContext } from '../context/Context';

import Drone from './Drone';

const Drones = () => {
  const { telemetries } = useContext(AppContext);

  return (
    <div>
      {telemetries.map(telemetry => (
        <Drone key={telemetry.id} telemetry={telemetry} />
      ))}
    </div>
  );
}

export default Drones;