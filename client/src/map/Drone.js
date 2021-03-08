import React, { useEffect, useState } from 'react';

import { Entity, ModelGraphics } from 'resium';

import { convertDroneData } from './CesiumHelper';

const Drone = (props) => {
  const [ data, setData ] = useState({});

  useEffect(() => {
    setData(convertDroneData(props.telemetry.id, props.telemetry.telemetry));
  }, [ props.telemetry, setData ])

  return (
    <Entity
      position={data.position}
      orientation={data.orientation}
    >
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