import React, { useContext } from 'react';

import { SceneMode } from 'cesium';

import { AppContext } from '../context/Context';
import Drone3D from './Drone3D';
import Drone2D from './Drone2D';

const Drones = () => {
  const { mapMode, telemetries } = useContext(AppContext);

  return (
    <>
      {mapMode === SceneMode.SCENE2D ? (
        telemetries.map(telemetry => (
          <Drone2D key={telemetry.id} telemetry={telemetry} />
        ))
      ) : (
        telemetries.map(telemetry => (
          <Drone3D key={telemetry.id} telemetry={telemetry} />
        ))
      )}
    </>
  );
}

export default Drones;