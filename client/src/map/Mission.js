import React, { useEffect, useReducer } from 'react';

import { getMission } from '../missions/missions/MissionUtils'
import Waypoint from './Waypoint';
import Path from './Path';
import { getPathsForDisplayToMap, getWaypointsForDisplayToMap } from '../missions/missions/MissionHelper';
import { editMissionReducer, initialEditMission } from '../context/EditMission';

const Mission = (props) => {
  const [ editMission, dispatchEditMission ] = useReducer(editMissionReducer, initialEditMission);

  useEffect(() => {
    getMission(props.data.missionId)
      .then(data => {
        dispatchEditMission({
          type: 'OPEN',
          mission: data,
        });
      });
  }, [ props.data.missionId, dispatchEditMission ]);

  return (
    <>
      {getWaypointsForDisplayToMap(editMission).map((waypoint, index) => (
        <Waypoint key={waypoint.id} index={index} waypoint={waypoint} />
      ))}
      {getPathsForDisplayToMap(editMission).map(path => (
        <Path key={path.id} path={path} />
      ))}
    </>
  );
}
  
export default Mission;