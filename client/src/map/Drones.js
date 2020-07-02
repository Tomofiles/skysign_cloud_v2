import React, { useState } from 'react';

import axios from 'axios';
import { convertDroneData } from './CesiumHelper';
import Drone from './Drone';
import useInterval from 'use-interval';

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
  
const Drones = () => {
  const [ data, setData ] = useState({id: ""});

  useInterval(() => {
    // getTelemetry("baa424fc-2710-42c0-b17d-d7cad938d0c3")
    //   .then(data => {
    //     setData(convertDroneData("baa424fc-2710-42c0-b17d-d7cad938d0c3", data));
    //   });
  },
  1000);

  return (
    <div>
      {/* {data.id !== "" &&
        <Drone data={data} />
      } */}
    </div>
  );
}

export default Drones;