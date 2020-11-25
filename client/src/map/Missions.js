import React, { useContext } from 'react';
import { AppContext } from '../context/Context';

import Mission from './Mission';

const Missions = () => {
  const { stagingRows } = useContext(AppContext);

  return (
    <div>
      {stagingRows.map(data => (
        data.isControlled && data.missionId !== "" &&
          <Mission key={data.id} data={data} />
      ))}
    </div>
  );
}

export default Missions;