import React, { useGlobal, useState } from 'reactn';
import { ScreenSpaceEvent, useCesium } from 'resium';
import { ScreenSpaceEventType, defined, Cartographic, Math, SceneMode } from 'cesium';
import { Mission } from '../plans/missions/MissionHelper';
import { EDIT_MODE } from '../App';

const MapDragAndDropEvent = (props) => {
  const context = useCesium();
  const [ mission, setMission ] = useGlobal("editMission");
  const editMode = useGlobal("editMode")[0];
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
      
      let newMission = Object.assign(Object.create(Mission.prototype), mission);
      newMission.changePosition(index, latitude, longitude);

      setMission(newMission);
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