import React, { useEffect, useState } from 'react';
import { Color, Cartesian3 } from 'cesium';
import { Entity, PolylineGraphics } from 'resium';

const Trajectory = (props) => {
  const [ data, setData ] = useState([]);

  useEffect(() => {
    const data = [];
    props.telemetries
      .forEach(telemetry => {
        data.push(Cartesian3.fromDegrees(
          telemetry.longitude,
          telemetry.latitude,
          telemetry.altitude));
      });
    setData(data);
  }, [ props.telemetries, setData ])

  return (
    <Entity>
      <PolylineGraphics
        positions={data}
        width={3.0}
        material={Color.fromCssColorString("#9600cd")}
      />
    </Entity>
  );
}
  
export default Trajectory;