import React, { useState, useEffect } from 'react';

import { hot } from "react-hot-loader/root";
import {
  Viewer,
  Camera,
  CameraFlyTo
} from "resium";
import { Cartesian3 } from "cesium"

const Map = (props) => {
  const [position, setPosition] = useState({ cartesian3: Cartesian3.fromDegrees(145.5, 35.5, 1000) });

  useEffect(() => {
    getCurrentPosition();
  }, [])

  const getCurrentPosition = () => {
    navigator.geolocation.getCurrentPosition(p => {
      const { latitude, longitude } = p.coords;
      const pos = Cartesian3.fromDegrees(longitude, latitude, 1000);
      setPosition({ cartesian3: pos });
    });
  };

  return (
    <div >
      <Viewer
        full={false}
        scene3DOnly={true}
        selectionIndicator={false}
        baseLayerPicker={false}
        navigationHelpButton={false}
        homeButton={false}
        geocoder={false}
        animation={false}
        timeline={false}
        fullscreenButton={false}
        className={props.classes.mapArea}
        >
          <Camera>
            <CameraFlyTo duration={0} destination={position.cartesian3} />
          </Camera>
      </Viewer>
    </div>
  );
}

export default hot(Map);
