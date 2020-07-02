import React, { useState, useEffect } from 'react';

import { hot } from "react-hot-loader/root";
import {
  Viewer,
  Camera,
  CameraFlyTo
} from "resium";
import { Cartesian3, createWorldTerrain, IonImageryProvider } from "cesium"

import Drones from './Drones'
import {} from './Key'

const Map = (props) => {
  const [position, setPosition] = useState({ cartesian3: Cartesian3.fromDegrees(-73.7578307, 45.467115299999996, 1000) });

  useEffect(() => {
    getCurrentPosition();
  }, [])

  const getCurrentPosition = () => {
    navigator.geolocation.getCurrentPosition(p => {
      const { latitude, longitude } = p.coords;
      const pos = Cartesian3.fromDegrees(longitude, latitude, 2000);
      // setPosition({ cartesian3: pos });
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
        imageryProvider={new IonImageryProvider({ assetId: 2 })}
        terrainProvider={createWorldTerrain()}
        >
          <Camera>
            <CameraFlyTo duration={0} destination={position.cartesian3} />
          </Camera>
          <Drones />
      </Viewer>
    </div>
  );
}

export default hot(Map);
