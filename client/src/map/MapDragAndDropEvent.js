import React, { useState, useContext } from 'react';
import { ScreenSpaceEvent, useCesium } from 'resium';
import { ScreenSpaceEventType, defined, Cartographic, Math, SceneMode } from 'cesium';
import { AppContext } from '../context/Context';
import { EDIT_MODE } from '../context/EditMode';

const MapDragAndDropEvent = () => {
  const context = useCesium();
  const { dispatchEditMission } = useContext(AppContext);
  const { editMode } = useContext(AppContext);
  const [ draggingWaypoint, setDraggingWaypoint ] = useState(undefined);

  const onLeftDown = event => {
    if (context.scene.mode !== SceneMode.SCENE2D) {
      return;
    }
    if (editMode !== EDIT_MODE.MISSION) {
      return;
    }
    let pickedObjects = context.scene.drillPick(event.position);
    for (let object of pickedObjects) {
      if (object.id.properties
            && object.id.properties.draggable.getValue()) {
        let index = object.id.properties.index.getValue();
        context.scene.screenSpaceCameraController.enableInputs = false;    
        setDraggingWaypoint(index);
      }
    }
  };

  const onMouseMove = event => {
    if (draggingWaypoint === undefined) {
      return;
    }
    let index = draggingWaypoint;
    let cartesian = context.camera.pickEllipsoid(event.endPosition);

    if (defined(cartesian)) {
      let cartographic = Cartographic.fromCartesian(cartesian);
      let longitude = Math.toDegrees(cartographic.longitude);
      let latitude = Math.toDegrees(cartographic.latitude);

      dispatchEditMission({
        type: 'CHANGE_POSITION',
        index: index,
        latitude: latitude,
        longitude: longitude,
      });
    }
  };

  const onLeftUp = event => {
    context.scene.screenSpaceCameraController.enableInputs = true;
    setDraggingWaypoint(undefined);
  };

  return (
    <>
      <ScreenSpaceEvent
        type={ScreenSpaceEventType.LEFT_DOWN}
        action={onLeftDown} />
      <ScreenSpaceEvent
        type={ScreenSpaceEventType.MOUSE_MOVE}
        action={onMouseMove} />
      <ScreenSpaceEvent
        type={ScreenSpaceEventType.LEFT_UP}
        action={onLeftUp} />
    </>
  );
}

export default MapDragAndDropEvent;