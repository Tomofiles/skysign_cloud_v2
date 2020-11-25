import React, { useContext } from 'react';
import { AppContext } from '../context/Context';

import Drone from './Drone';

const Drones = () => {
  const { stagingRows } = useContext(AppContext);

  return (
    <div>
      {stagingRows.map(data => (
        data.isControlled &&
          <Drone key={data.id} data={data} />
      ))}
    </div>
  );
}

export default Drones;