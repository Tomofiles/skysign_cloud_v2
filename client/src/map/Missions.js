import React, { useContext } from 'react';
import { AppContext } from '../context/Context';

import Mission from './Mission';

const Missions = () => {
  const { missions } = useContext(AppContext);

  return (
    <>
      {missions.map(mission => (
        <Mission key={mission.id} mission={mission} />
      ))}
    </>
  );
}

export default Missions;