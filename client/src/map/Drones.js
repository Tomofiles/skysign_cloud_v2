import React, { useGlobal } from 'reactn';

import Drone from './Drone';

const Drones = () => {
  const [ rows ] = useGlobal("stagingRows");

  return (
    <div>
      {rows.map(data => (
        data.isControlled &&
          <Drone key={data.id} data={data} />
      ))}
    </div>
  );
}

export default Drones;