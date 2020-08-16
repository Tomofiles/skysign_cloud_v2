import React, { useState } from 'react';

import { Entity, ModelGraphics } from 'resium';
import useInterval from 'use-interval';

import { convertDroneData } from './CesiumHelper';
import { getTelemetry } from './MapUtils'

const Drone = (props) => {
  const [ data, setData ] = useState({});

  useInterval(() => {
    getTelemetry(props.data.id)
      .then(data => {
        setData(convertDroneData(props.data.vehicleId, data.telemetry));
      });
  },
  1000);

  return (
    <Entity
      position={data.position}
      orientation={data.orientation}
      description={data.description} >
        <ModelGraphics
          uri="CesiumDrone.gltf"
          scale={0.05}
          minimumPixelSize={100}
          show
          runAnimations={data.armed} />
    </Entity>
  );
}
  
export default Drone;