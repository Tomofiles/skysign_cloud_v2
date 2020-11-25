import React, { useContext } from 'react';

import Waypoint from './Waypoint';
import Path from './Path';
import { AppContext } from '../context/Context';
import { getPathsForDisplayToMap, getWaypointsForDisplayToMap } from '../plans/missions/MissionHelper';

const EditMission = () => {
  const { editMission } = useContext(AppContext);

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

export default EditMission;