import React, { useContext } from 'react';
import { ScreenSpaceEvent, useCesium } from 'resium';
import { ScreenSpaceEventType, defined, Cartographic, Math, SceneMode } from 'cesium';
import { AppContext } from '../context/Context';
import { EDIT_MODE } from '../context/EditMode';

const MapDoubleClickEvent = () => {
  const context = useCesium();
  const { dispatchEditMission } = useContext(AppContext);
  const { editMode } = useContext(AppContext);

  const onDoubleClick = async (event) => {
    if (context.scene.mode === SceneMode.SCENE2D) {
      let cartesian = context.camera.pickEllipsoid(event.position);

      if (defined(cartesian)) {
        let cartographic = Cartographic.fromCartesian(cartesian);
        let longitude = Math.toDegrees(cartographic.longitude);
        let latitude = Math.toDegrees(cartographic.latitude);

        if (editMode === EDIT_MODE.MISSION) {
          dispatchEditMission({
            type: 'ADD_WAYPOINT',
            latitude: latitude,
            longitude: longitude,
          });
        }
      }
    }
  };

  return (
    <ScreenSpaceEvent
      type={ScreenSpaceEventType.LEFT_DOUBLE_CLICK}
      action={onDoubleClick} />
  );
}

export default MapDoubleClickEvent;