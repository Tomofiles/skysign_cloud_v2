import React, { useState, useEffect } from 'react';

import { Entity } from 'resium';
import axios from 'axios';
import useInterval from 'use-interval';

import { convertDroneData } from './CesiumHelper';

export async function getTelemetry(id) {
  try {
    const res = await axios
      .get(`/api/v1/vehicles/${id}/telemetries`, {
        params: {}
      })
    return res.data;
  } catch(error) {
    console.log(error);
  }
}
  
const Drone = (props) => {
  const [ data, setData ] = useState({});

  useInterval(() => {
    getTelemetry(props.data.vehicle)
      .then(data => {
        setData(convertDroneData(props.data.vehicle, data));
      });
  },
  1000);

  return (
    <Entity
      id={data.id}
      position={data.position}
      orientation={data.orientation}
      model={data.model} />
  );
}
  
export default Drone;