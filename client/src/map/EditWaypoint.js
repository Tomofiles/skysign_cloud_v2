import React from 'react';

import { Entity } from 'resium';
import { PolylineDashMaterialProperty, Color, Cartesian2 } from 'cesium';

const EditWaypoint = (props) => {
  return (
    <>
      <Entity
        position={props.waypoint.groundPosition}
        point={{
          pixelSize: 17.0,
          color: Color.WHITE
        }} />
      <Entity
        polyline={{
          positions: [
            props.waypoint.groundPosition,
            props.waypoint.airPosition
          ],
          material: new PolylineDashMaterialProperty({
            color: Color.fromCssColorString("#2bb3c0"),
            dashPattern: 3855,
          }),
          width: 1
        }} />
      <Entity
        position={props.waypoint.airPosition}
        point={{
          pixelSize: 10.0,
          color: Color.WHITE
        }}
        label={{
          text: "WP" + (props.index + 1),
          font: "8px",
          pixelOffset: new Cartesian2(0, -15.0)
        }} />
    </>
  );
}
  
export default EditWaypoint;