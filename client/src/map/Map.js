import React from 'react';

import { hot } from "react-hot-loader/root";
import {
  Viewer,
  Camera,
  Clock,
  ScreenSpaceEventHandler,
  ImageryLayer,
} from "resium";
import { createWorldTerrain, IonImageryProvider } from "cesium"

import Drones from './Drones'
import EditMission from './EditMission';
import {} from './Key'
import MapDoubleClickEvent from './MapDoubleClickEvent';
import MapDragAndDropEvent from './MapDragAndDropEvent';
import SceneMode from './SceneMode'
import Missions from './Missions';
import CameraCurrentPosition from './CameraCurrentPosition';

const imageryProvider = new IonImageryProvider({ assetId: 2 });

const Map = (props) => {
  return (
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
        <CameraCurrentPosition />
      </Camera>
      <Drones />
      <Missions />
      <EditMission />
      <ScreenSpaceEventHandler >
        <MapDoubleClickEvent />
        <MapDragAndDropEvent />
      </ScreenSpaceEventHandler>
    </Viewer>
  );
}

export default hot(Map);
