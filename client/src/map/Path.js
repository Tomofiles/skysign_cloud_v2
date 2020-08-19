import React from 'react';

import { Entity } from 'resium';
import { PolylineArrowMaterialProperty, Color } from 'cesium';

const Path = (props) => {
  return (
    <Entity
      polyline={{
        positions: [
          props.path.prevPosition,
          props.path.currentPosition
        ],
        material: new PolylineArrowMaterialProperty(
          Color.fromCssColorString("#2bb3c0")
        ),
        width: 10.0,
      }} />
  );
}
  
export default Path;