import React, { useGlobal } from 'reactn';
import { ScreenSpaceEvent, useCesium } from 'resium';
import { ScreenSpaceEventType, defined, Cartographic, Math, SceneMode } from 'cesium';
import { Mission } from '../plans/missions/MissionHelper';
import { EDIT_MODE } from '../App';

const MapDoubleClickEvent = (props) => {
  const context = useCesium();
  const [ mission, setMission ] = useGlobal("editMission");
  const editMode = useGlobal("editMode")[0];

  const onDoubleClick = async (event) => {
    if (context.scene.mode === SceneMode.SCENE2D) {
      let cartesian = context.camera.pickEllipsoid(event.position);

      if (defined(cartesian)) {
        let cartographic = Cartographic.fromCartesian(cartesian);
        let longitude = Math.toDegrees(cartographic.longitude);
        let latitude = Math.toDegrees(cartographic.latitude);

        if (editMode === EDIT_MODE.MISSION) {
          let newMission = Object.assign(Object.create(Mission.prototype), mission);
          await newMission.addWaypoint(latitude, longitude);
  
          setMission(newMission);
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