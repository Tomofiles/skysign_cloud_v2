import React, { useContext } from 'react';
import { Scene } from 'resium';
import { AppContext } from '../context/Context';

const SceneMode = () => {
  const { mapMode } = useContext(AppContext);

  return (
    <Scene
      mode={mapMode}
      morphTime={0.0}
      morphDuration={0.0} />
  );
}

export default SceneMode;