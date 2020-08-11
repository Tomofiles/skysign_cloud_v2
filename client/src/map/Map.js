import React, { useState, useEffect } from 'react';

import { hot } from "react-hot-loader/root";
import {
  Viewer,
  Camera,
  CameraFlyTo,
  Clock,
  ScreenSpaceEventHandler,
  ImageryLayer,
} from "resium";
import { Cartesian3, createWorldTerrain, IonImageryProvider } from "cesium"

import Drones from './Drones'
import EditMission from './EditMission';
import {} from './Key'
import MapDoubleClickEvent from './MapDoubleClickEvent';
import SceneMode from './SceneMode'
import Missions from './Missions';

const imageryProvider = new IonImageryProvider({ assetId: 2 });

const Map = (props) => {
  const [ position, setPosition ] = useState({ cartesian3: Cartesian3.fromDegrees(-73.7578307, 45.467115299999996, 1000) });

  useEffect(() => {
    getCurrentPosition();
  }, [])

  const getCurrentPosition = () => {
    navigator.geolocation.getCurrentPosition(p => {
      const { latitude, longitude } = p.coords;
      const pos = Cartesian3.fromDegrees(longitude, latitude, 2000);
      setPosition({ cartesian3: pos });
    });
  };

  return (
    <div >
      <Viewer
        full={false}
        sceneModePicker={false}
        selectionIndicator={false}
        baseLayerPicker={false}
        navigationHelpButton={false}
        homeButton={false}
        geocoder={false}
        animation={false}
        timeline={false}
        fullscreenButton={false}
        className={props.classes.mapArea}
        terrainProvider={createWorldTerrain()}
        >
          <ImageryLayer
            imageryProvider={imageryProvider} />
          <SceneMode />
          <Clock shouldAnimate />
          <Camera>
            <CameraFlyTo duration={0} destination={position.cartesian3} />
          </Camera>
          <Drones />
          <Missions />
          <EditMission />
          <ScreenSpaceEventHandler >
            <MapDoubleClickEvent />
          </ScreenSpaceEventHandler>
      </Viewer>
    </div>
  );
}

export default hot(Map);
