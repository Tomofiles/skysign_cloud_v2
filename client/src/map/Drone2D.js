import React, { useEffect, useState } from 'react';

import { BillboardGraphics, Entity } from 'resium';

import { convertDroneData } from './CesiumHelper';

const Drone2D = (props) => {
  const [ data, setData ] = useState({});

  useEffect(() => {
    setData(convertDroneData(props.telemetry.id, props.telemetry.telemetry));
  }, [ props.telemetry, setData ])

  return (
    <Entity
      position={data.position}
    >
      <BillboardGraphics
        image="drone_symbol_2d.png"
        scale={0.5}
        alignedAxis={data.alignedAxis}
        rotation={-data.heading}
        minimumPixelSize={100}
        show />
    </Entity>
  );
}
  
export default Drone2D;