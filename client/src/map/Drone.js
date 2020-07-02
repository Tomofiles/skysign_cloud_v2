import React from 'react';

import { Entity } from 'resium';

const Drone = (props) => {
  return (
    <Entity
      id={props.data.id}
      position={props.data.position}
      orientation={props.data.orientation}
      model={props.data.model} />
  );
}
  
export default Drone;