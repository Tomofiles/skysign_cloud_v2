import React, { useContext } from 'react';
import { ScreenSpaceEvent, useCesium } from 'resium';
import { ScreenSpaceEventType, defined, Cartographic, Math, SceneMode } from 'cesium';
import { AppContext } from '../context/Context';
import { EDIT_MODE } from '../context/EditMode';
import { getTakeoffHeight } from './MapUtils';

const MapDoubleClickEvent = () => {
  const context = useCesium();
  const { editMission, dispatchEditMission, dispatchMessage } = useContext(AppContext);
  const { editMode } = useContext(AppContext);

  const onDoubleClick = async (event) => {
    if (context.scene.mode === SceneMode.SCENE2D) {
      let cartesian = context.camera.pickEllipsoid(event.position);

      if (defined(cartesian)) {
        let cartographic = Cartographic.fromCartesian(cartesian);
        let longitude = Math.toDegrees(cartographic.longitude);
        let latitude = Math.toDegrees(cartographic.latitude);

        if (editMode === EDIT_MODE.MISSION) {
          if (editMission.navigation.waypoints.length === 0) {
            getTakeoffHeight(latitude, longitude)
              .then(height => {
                dispatchEditMission({
                  type: 'CHANGE_TAKEOFF_POINT_GROUND_HEIGHT',
                  height: height.height,
                });
                dispatchEditMission({
                  type: 'ADD_WAYPOINT',
                  latitude: latitude,
                  longitude: longitude,
                });
              })
              .catch(message => {
                dispatchMessage({ type: 'NOTIFY_ERROR', message: message });
              });
          } else {
            dispatchEditMission({
              type: 'ADD_WAYPOINT',
              latitude: latitude,
              longitude: longitude,
            });
          }
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